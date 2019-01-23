package blockchain

import (
	"github.com/pkg/errors"
	"vega/msg"
	"github.com/golang/protobuf/proto"
	"fmt"
)

type Processor interface {
	Process (payload []byte) error
	Validate (payload []byte) error
}

type abciProcessor struct {
	*Config
	
	blockchainService Service
	seenPayloads map[string]byte
}

func NewAbciProcessor(config *Config, blockchainService Service) Processor {
	return &abciProcessor{
		Config: config,
		blockchainService: blockchainService,
		seenPayloads: make(map[string]byte, 0),
	}
}

func (p *abciProcessor) getOrder(payload []byte) (*msg.Order, error) {
	order := msg.OrderPool.Get().(*msg.Order)
	err := proto.Unmarshal(payload, order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

// todo remove this in favour of raw Order above?
func (p *abciProcessor) getOrderAmendment(payload []byte) (*msg.Amendment, error) {
	amendment := &msg.Amendment{}
	err := proto.Unmarshal(payload, amendment)
	if err != nil {
		p.log.Errorf("Error: Decoding order to proto: ", err.Error())
		return nil, err
	}
	return amendment, nil
}

// Validate performs all validation on an incoming transaction payload.
func (p *abciProcessor) Validate(payload []byte) error {
	// Pre-validate (safety check)
 	if err, seen := p.hasSeen(payload); seen {
		p.log.Errorf("AbciProcessor: txn payload already exists, %s", err)
		return err
	}
	// Attempt to decode transaction payload
	_, cmd, err := txDecode(payload)
	if err != nil {
		p.log.Errorf("AbciProcessor: Decode failed when validating, %s", err)
		return err
	}
	// Ensure valid VEGA app command
	switch cmd {
		case
			SubmitOrderCommand,
			CancelOrderCommand,
			AmendOrderCommand:       // Note: Future valid VEGA commands here
			return nil
	}
	return errors.New("Unknown command when validating payload")

	// todo(cdm): More validation here using blockchain service methods
	//p.blockchainService.ValidateOrder()
}

// Process performs validation and then sends the command and data to
// the underlying blockchain service handlers e.g. submit order, etc.
func (p *abciProcessor) Process(payload []byte) error {
	// Pre-validate (safety check)
	if err, seen := p.hasSeen(payload); seen {
		p.log.Errorf("AbciProcessor: txn payload already exists, %s", err)
		return err
	}
	// Attempt to decode transaction payload
	data, cmd, err := txDecode(payload)
	if err != nil {
		p.log.Errorf("AbciProcessor: Decode failed when processing, %s", err)
		return err
	}
	// Process known command types
	switch cmd {
		case SubmitOrderCommand:
			order, err := p.getOrder(data)
			if err != nil {
				return err
			}
		    err = p.blockchainService.SubmitOrder(order)
		    if err != nil {
		    	return err
			}
		case CancelOrderCommand:
			order, err := p.getOrder(data)
			if err != nil {
				return err
			}
			err = p.blockchainService.CancelOrder(order)
			if err != nil {
				return err
			}
		case AmendOrderCommand:
			order, err := p.getOrderAmendment(data)
			if err != nil {
				return err
			}
			err = p.blockchainService.AmendOrder(order)
			if err != nil {
				return err
			}
		//case FutureVegaCommand
			// Note: Future valid VEGA commands here
		default:
			// todo structured logger
			p.log.Errorf("Unknown command received: %s", cmd)
			return errors.New(fmt.Sprintf("Unknown command received: %s", cmd))
	}
	return nil
}

// hasSeen helper performs duplicate checking on an incoming transaction payload.
func (p *abciProcessor) hasSeen(payload []byte) (error, bool) {
	// All vega transactions are prefixed with a unique hash, using
	// this means we do not have to re-compute each time for seen keys
	payloadHash, err := p.payloadHash(payload)
	if err != nil {
		return err, true
	}
	// Safety checks at business level to ensure duplicate transaction
	// payloads do not pass through to application core
	if err, exists := p.payloadExists(payloadHash); exists {
		return err, true
	}
	return nil, false
}

// payloadHash attempts to extract the unique hash at the start of all vega transactions.
// This unique hash is required to make all payloads unique. We return an error if we cannot
// extract this from the transaction payload or if we think it's malformed.
func (p *abciProcessor) payloadHash(payload []byte) (*string, error) {
	if len(payload) < 36 {
		return nil, errors.New("invalid length payload, must be greater than 37 bytes")
	}
	hash := string(payload[0:36])
	return &hash, nil
}

// payloadExists checks to see if a payload has been seen before in this batch
// recommended by tendermint team that an abci application has additional checking
// just like this to ensure duplicate transaction payloads do not pass through
// to the application core.
func (p *abciProcessor) payloadExists(payloadHash *string) (error, bool) {
	if _, exists := p.seenPayloads[*payloadHash]; exists {
		err := errors.New(fmt.Sprintf("payload exists: %s", payloadHash))
		p.log.Errorf("AbciProcessor: %s", err)
		return err, true
	}
	p.seenPayloads[*payloadHash] = 0xF
	return nil, false
}

// txDecode is takes the raw payload bytes and decodes the contents using a pre-defined
// strategy, we have a simple and efficient encoding at present. A partner encode function
// can be found in the blockchain client.
func txDecode (input []byte) (proto []byte, cmd Command, err error) {
	// Input is typically the bytes that arrive in raw format after consensus is reached.
	// Split the transaction dropping the unification bytes (uuid&pipe)
	var value []byte
	var cmdByte byte
	if len(input) > 37 {
		// obtain command from byte slice (0 indexed)
		cmdByte = input[36]
		// remaining bytes are payload
		value = input[37:]
	} else {
		return nil, 0, errors.New("payload size is incorrect, should be > 38 bytes")
	}
	return value, Command(cmdByte), nil
}
