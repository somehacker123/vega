package node

import (
	"context"
	"sync/atomic"
	"time"

	api "code.vegaprotocol.io/vega/protos/vega/api/v1"
	commandspb "code.vegaprotocol.io/vega/protos/vega/commands/v1"
	"code.vegaprotocol.io/vega/wallet/network"

	"github.com/cenkalti/backoff/v4"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type Forwarder struct {
	log      *zap.Logger
	nodeCfgs network.HostConfig
	clts     []api.CoreServiceClient
	conns    []*grpc.ClientConn
	next     uint64
}

func NewForwarder(log *zap.Logger, nodeConfigs network.HostConfig) (*Forwarder, error) {
	if len(nodeConfigs.Hosts) == 0 {
		return nil, ErrNoHostSpecified
	}

	clts := make([]api.CoreServiceClient, 0, len(nodeConfigs.Hosts))
	conns := make([]*grpc.ClientConn, 0, len(nodeConfigs.Hosts))
	for _, v := range nodeConfigs.Hosts {
		conn, err := grpc.Dial(v, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Debug("Couldn't dial gRPC host", zap.String("address", v))
			return nil, err
		}
		conns = append(conns, conn)
		clts = append(clts, api.NewCoreServiceClient(conn))
	}

	return &Forwarder{
		log:      log,
		nodeCfgs: nodeConfigs,
		clts:     clts,
		conns:    conns,
	}, nil
}

func (n *Forwarder) Stop() {
	hasErrors := false
	for i, v := range n.nodeCfgs.Hosts {
		n.log.Debug("Closing gRPC client", zap.String("address", v))
		if err := n.conns[i].Close(); err != nil {
			n.log.Warn("Couldn't close gRPC client", zap.Error(err))
			hasErrors = true
		}
	}
	if hasErrors {
		n.log.Warn("gRPC clients successfully with errors")
	} else {
		n.log.Info("gRPC clients successfully closed")
	}
}

func (n *Forwarder) HealthCheck(ctx context.Context) error {
	req := api.GetVegaTimeRequest{}
	return backoff.Retry(
		func() error {
			clt := n.clts[n.nextClt()]
			resp, err := clt.GetVegaTime(ctx, &req)
			if err != nil {
				return err
			}
			n.log.Debug("Response from GetVegaTime", zap.Int64("timestamp", resp.Timestamp))
			return nil
		},
		backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 5),
	)
}

// LastBlockHeightAndHash returns information about the last block from vega and the node it used to fetch it.
func (n *Forwarder) LastBlockHeightAndHash(ctx context.Context) (*api.LastBlockHeightResponse, int, error) {
	req := api.LastBlockHeightRequest{}
	var resp *api.LastBlockHeightResponse
	clt := -1
	err := backoff.Retry(
		func() error {
			clt = n.nextClt()
			r, err := n.clts[clt].LastBlockHeight(ctx, &req)
			if err != nil {
				n.log.Debug("Couldn't get last block", zap.Error(err))
				return err
			}
			resp = r
			n.log.Info("", zap.Uint64("block-height", r.Height), zap.String("block-hash", r.Hash), zap.Uint32("difficulty", r.SpamPowDifficulty), zap.String("function", r.SpamPowHashFunction))
			return nil
		},
		backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 5),
	)

	if err != nil {
		n.log.Error("Couldn't get last block", zap.Error(err))
		clt = -1
	} else {
		n.log.Debug("Last block when sending transaction",
			zap.Time("request.time", time.Now()),
			zap.Uint64("block.height", resp.Height),
		)
	}

	return resp, clt, err
}

func (n *Forwarder) CheckTx(ctx context.Context, tx *commandspb.Transaction, cltIdx int) (*api.CheckTransactionResponse, error) {
	req := api.CheckTransactionRequest{
		Tx: tx,
	}
	var resp *api.CheckTransactionResponse
	if cltIdx < 0 {
		cltIdx = n.nextClt()
	}
	err := backoff.Retry(
		func() error {
			clt := n.clts[cltIdx]
			r, err := clt.CheckTransaction(ctx, &req)
			if err != nil {
				n.log.Error("Couldn't check transaction", zap.Error(err))
				return err
			}
			n.log.Debug("Response from CheckTransaction",
				zap.Bool("success", r.Success),
			)
			resp = r
			return nil
		},
		backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 5),
	)

	return resp, err
}

func (n *Forwarder) SpamStatistics(ctx context.Context, pubkey string) (*api.GetSpamStatisticsResponse, int, error) {
	req := api.GetSpamStatisticsRequest{
		PartyId: pubkey,
	}
	var resp *api.GetSpamStatisticsResponse
	clt := -1
	err := backoff.Retry(
		func() error {
			clt = n.nextClt()
			r, err := n.clts[clt].GetSpamStatistics(ctx, &req)
			if err != nil {
				n.log.Debug("Couldn't get spam statistics", zap.Error(err))
				return err
			}
			resp = r
			return nil
		},
		backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 5),
	)

	if err != nil {
		n.log.Error("Couldn't get spam statistics", zap.Error(err))
		clt = -1
	} else {
		n.log.Debug("Spam statistics",
			zap.Time("request.time", time.Now()),
		)
	}

	return resp, clt, err
}

func (n *Forwarder) SendTx(ctx context.Context, tx *commandspb.Transaction, ty api.SubmitTransactionRequest_Type, cltIdx int) (*api.SubmitTransactionResponse, error) {
	req := api.SubmitTransactionRequest{
		Tx:   tx,
		Type: ty,
	}
	var resp *api.SubmitTransactionResponse
	if cltIdx < 0 {
		cltIdx = n.nextClt()
	}
	if err := backoff.Retry(
		func() error {
			clt := n.clts[cltIdx]
			r, err := clt.SubmitTransaction(ctx, &req)
			if err != nil {
				return n.handleSubmissionError(err)
			}
			n.log.Debug("Transaction successfully submitted",
				zap.Bool("success", r.Success),
				zap.String("hash", r.TxHash),
			)
			resp = r
			return nil
		},
		backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 5),
	); err != nil {
		return nil, err
	}

	return resp, nil
}

func (n *Forwarder) handleSubmissionError(err error) error {
	statusErr := intoStatusError(err)

	if statusErr == nil {
		n.log.Error("couldn't submit transaction",
			zap.Error(err),
		)
		return err
	}

	if statusErr.Code == codes.InvalidArgument {
		n.log.Error(
			"transaction has been rejected because of an invalid argument or state, skipping retry...",
			zap.Error(statusErr),
		)
		// Returning a permanent error kills the retry loop.
		return backoff.Permanent(statusErr)
	}

	n.log.Error("couldn't submit transaction",
		zap.Error(statusErr),
	)
	return statusErr
}

func (n *Forwarder) nextClt() int {
	i := atomic.AddUint64(&n.next, 1)
	n.log.Info("Sending transaction to Vega node",
		zap.String("host", n.nodeCfgs.Hosts[(int(i)-1)%len(n.clts)]),
	)
	return (int(i) - 1) % len(n.clts)
}
