package api_test

import (
	"context"
	"io"
	"testing"

	apipb "code.vegaprotocol.io/protos/data-node/api/v1"
	pb "code.vegaprotocol.io/protos/vega"
	eventspb "code.vegaprotocol.io/protos/vega/events/v1"
	"code.vegaprotocol.io/vega/events"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestObserveTransferResponses(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimout)
	defer cancel()

	server := NewTestServer(t, ctx, true)
	defer server.ctrl.Finish()

	client := apipb.NewTradingDataServiceClient(server.clientConn)
	require.NotNil(t, client)

	// we need to subscribe to the stream prior to publishing the events
	stream, err := client.TransferResponsesSubscribe(ctx, &apipb.TransferResponsesSubscribeRequest{})
	assert.NoError(t, err)

	PublishEvents(t, ctx, server.broker, func(be *eventspb.BusEvent) (events.Event, error) {
		tr := be.GetTransferResponses()
		require.NotNil(t, tr)
		var responses []*pb.TransferResponse
		for _, resp := range tr.Responses {
			responses = append(responses, &pb.TransferResponse{
				Transfers: resp.Transfers,
				Balances:  resp.Balances,
			})
		}
		e := events.NewTransferResponse(ctx, TransferResponsesFromProto(responses))
		return e, nil
	}, "transfer-responses-events.golden")

	// The stream contains a timeout from the context we gave it at creation so we don't
	// have to worry about this blocking forever
	resp, err := stream.Recv()
	if err != io.EOF {
		require.NoError(t, err)
	}

	require.NotNil(t, resp)
	require.Equal(t, "076BB86A5AA41E3E*6d9d35f657589e40ddfb448b7ad4a7463b66efb307527fedd2aa7df1bbd5ea616", resp.Response.Transfers[0].FromAccount)
	require.Equal(t, "076BB86A5AA41E3E0f3d86044f8e7efff27131227235fb6db82574e24f788c30723d67f888b51d616d9d35f657589e40ddfb448b7ad4a7463b66efb307527fedd2aa7df1bbd5ea613", resp.Response.Transfers[0].ToAccount)
	require.Equal(t, "10412267", resp.Response.Transfers[0].Amount)
	require.Equal(t, "TRANSFER_TYPE_LIQUIDITY_FEE_DISTRIBUTE", resp.Response.Transfers[0].Reference)
	require.Equal(t, "settlement", resp.Response.Transfers[0].Type)
	require.Equal(t, int64(1622563663355188728), resp.Response.Transfers[0].Timestamp)

	require.Equal(t, "076BB86A5AA41E3E0f3d86044f8e7efff27131227235fb6db82574e24f788c30723d67f888b51d616d9d35f657589e40ddfb448b7ad4a7463b66efb307527fedd2aa7df1bbd5ea613", resp.Response.Balances[0].Account.Id)
}
