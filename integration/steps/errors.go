package steps

import (
	"fmt"
	"strconv"
	"strings"

	types "code.vegaprotocol.io/protos/vega"
	commandspb "code.vegaprotocol.io/protos/vega/commands/v1"
	"code.vegaprotocol.io/vega/integration/stubs"
	"code.vegaprotocol.io/vega/logging"
)

type ErroneousRow interface {
	ExpectError() bool
	Error() string
	Reference() string
}

func DebugTxErrors(broker *stubs.BrokerStub, log *logging.Logger) {
	log.Info("DEBUGGING ALL TRANSACTION ERRORS")
	data := broker.GetTxErrors()
	for _, e := range data {
		p := e.Proto()
		log.Infof("TxError: %s\n", p.String())
	}
}

func DebugLPSTxErrors(broker *stubs.BrokerStub, log *logging.Logger) {
	log.Info("DEBUGGING LP SUBMISSION ERRORS")
	data := broker.GetLPSErrors()
	for _, e := range data {
		p := e.Proto()
		log.Infof("LP Submission error: %s - LP: %#v\n", p.String(), p.GetLiquidityProvisionSubmission)
	}
}

func checkExpectedError(row ErroneousRow, returnedErr error) error {
	if row.ExpectError() && returnedErr == nil {
		return fmt.Errorf("\"%s\" should have fail", row.Reference())
	}

	if returnedErr != nil {
		if !row.ExpectError() {
			return fmt.Errorf("\"%s\" has failed: %s", row.Reference(), returnedErr.Error())
		}

		if row.Error() != returnedErr.Error() {
			return formatDiff(fmt.Sprintf("\"%s\" is failing as expected but not with the expected error message", row.Reference()),
				map[string]string{
					"error": row.Error(),
				},
				map[string]string{
					"error": returnedErr.Error(),
				},
			)
		}
	}
	return nil
}

func formatDiff(msg string, expected, got map[string]string) error {
	var expectedStr strings.Builder
	var gotStr strings.Builder
	formatStr := "\n\t%s\t(%s)"
	for name, value := range expected {
		_, _ = fmt.Fprintf(&expectedStr, formatStr, name, value)
		_, _ = fmt.Fprintf(&gotStr, formatStr, name, got[name])
	}

	return fmt.Errorf("\n%s\nexpected:%s\ngot:%s",
		msg,
		expectedStr.String(),
		gotStr.String(),
	)
}

func u64ToS(n uint64) string {
	return strconv.FormatUint(n, 10)
}

func u64SToS(ns []uint64) string {
	ss := []string{}
	for _, n := range ns {
		ss = append(ss, u64ToS(n))
	}
	return strings.Join(ss, " ")
}

func i64ToS(n int64) string {
	return strconv.FormatInt(n, 10)
}

func errOrderNotFound(reference string, party string, err error) error {
	return fmt.Errorf("order not found for party(%s) with reference(%s): %v", party, reference, err)
}

func errMarketDataNotFound(marketID string, err error) error {
	return fmt.Errorf("market data not found for market(%v): %s", marketID, err.Error())
}

type CancelOrderError struct {
	reference string
	request   commandspb.OrderCancellation
	Err       error
}

func (c CancelOrderError) Error() string {
	return fmt.Sprintf("failed to cancel order [%v] with reference %s: %v", c.request, c.reference, c.Err)
}

func (c *CancelOrderError) Unwrap() error { return c.Err }

type SubmitOrderError struct {
	reference string
	request   commandspb.OrderSubmission
	Err       error
}

func (s SubmitOrderError) Error() string {
	return fmt.Sprintf("failed to submit order [%v] with reference %s: %v", s.request, s.reference, s.Err)
}

func (s *SubmitOrderError) Unwrap() error { return s.Err }

func errOrderEventsNotFound(party, marketID string, side types.Side, size, price uint64) error {
	return fmt.Errorf("no matching order event found %v, %v, %v, %v, %v", party, marketID, side.String(), size, price)
}