package node

import (
	"context"
	"sync/atomic"
	"time"

	api "code.vegaprotocol.io/protos/vega/api/v1"
	commandspb "code.vegaprotocol.io/protos/vega/commands/v1"
	"code.vegaprotocol.io/vega/wallet/network"

	"github.com/cenkalti/backoff/v4"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type Forwarder struct {
	log      *zap.Logger
	nodeCfgs network.GRPCConfig
	clts     []api.CoreServiceClient
	conns    []*grpc.ClientConn
	next     uint64
}

func NewForwarder(log *zap.Logger, nodeConfigs network.GRPCConfig) (*Forwarder, error) {
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

func (n *Forwarder) Stop() error {
	for i, v := range n.nodeCfgs.Hosts {
		n.log.Debug("Closing gRPC client", zap.String("address", v))
		if err := n.conns[i].Close(); err != nil {
			n.log.Warn("Couldn't close gRPC client", zap.Error(err))
			return err
		}
	}
	n.log.Info("gRPC clients successfully closed")
	return nil
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
		backoff.WithMaxRetries(backoff.NewExponentialBackOff(), n.nodeCfgs.Retries),
	)
}

func (n *Forwarder) GetNetworkChainID(ctx context.Context) (string, error) {
	chainID := ""
	req := api.StatisticsRequest{}
	err := backoff.Retry(
		func() error {
			clt := n.clts[n.nextClt()]
			resp, err := clt.Statistics(ctx, &req)
			if err != nil {
				return err
			}
			chainID = resp.Statistics.ChainId
			n.log.Debug("Response from Statistics", zap.String("chainID", chainID))
			return nil
		},
		backoff.WithMaxRetries(backoff.NewExponentialBackOff(), n.nodeCfgs.Retries),
	)
	if err != nil {
		n.log.Error("Couldn't get chainID", zap.Error(err))
	}

	return chainID, nil
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
		backoff.WithMaxRetries(backoff.NewExponentialBackOff(), n.nodeCfgs.Retries),
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
		backoff.WithMaxRetries(backoff.NewExponentialBackOff(), n.nodeCfgs.Retries),
	)

	return resp, err
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
		backoff.WithMaxRetries(backoff.NewExponentialBackOff(), n.nodeCfgs.Retries),
	); err != nil {
		return resp, err
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
