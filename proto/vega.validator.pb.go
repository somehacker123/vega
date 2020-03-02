// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/vega.proto

package proto

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Amount) Validate() error {
	return nil
}
func (this *Party) Validate() error {
	return nil
}
func (this *RiskFactor) Validate() error {
	return nil
}
func (this *RiskResult) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *Order) Validate() error {
	return nil
}
func (this *OrderCancellationConfirmation) Validate() error {
	if this.Order != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Order); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Order", err)
		}
	}
	return nil
}
func (this *OrderConfirmation) Validate() error {
	if this.Order != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Order); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Order", err)
		}
	}
	for _, item := range this.Trades {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Trades", err)
			}
		}
	}
	for _, item := range this.PassiveOrdersAffected {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PassiveOrdersAffected", err)
			}
		}
	}
	return nil
}
func (this *Trade) Validate() error {
	return nil
}
func (this *TradeSet) Validate() error {
	for _, item := range this.Trades {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Trades", err)
			}
		}
	}
	return nil
}
func (this *Candle) Validate() error {
	return nil
}
func (this *PriceLevel) Validate() error {
	return nil
}
func (this *MarketDepth) Validate() error {
	for _, item := range this.Buy {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Buy", err)
			}
		}
	}
	for _, item := range this.Sell {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Sell", err)
			}
		}
	}
	return nil
}
func (this *Position) Validate() error {
	return nil
}
func (this *PositionTrade) Validate() error {
	return nil
}
func (this *Statistics) Validate() error {
	return nil
}
func (this *PendingOrder) Validate() error {
	return nil
}
func (this *NotifyTraderAccount) Validate() error {
	return nil
}
func (this *Withdraw) Validate() error {
	return nil
}
func (this *OrderAmendment) Validate() error {
	if this.OrderID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("OrderID", fmt.Errorf(`value '%v' must not be an empty string`, this.OrderID))
	}
	if this.PartyID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartyID", fmt.Errorf(`value '%v' must not be an empty string`, this.PartyID))
	}
	if !(this.Price > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Price", fmt.Errorf(`value '%v' must be greater than '0'`, this.Price))
	}
	if !(this.Size > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Size_", fmt.Errorf(`value '%v' must be greater than '0'`, this.Size))
	}
	return nil
}
func (this *OrderSubmission) Validate() error {
	if this.MarketID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MarketID", fmt.Errorf(`value '%v' must not be an empty string`, this.MarketID))
	}
	if this.PartyID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartyID", fmt.Errorf(`value '%v' must not be an empty string`, this.PartyID))
	}
	if !(this.Size > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Size_", fmt.Errorf(`value '%v' must be greater than '0'`, this.Size))
	}
	if _, ok := Side_name[int32(this.Side)]; !ok {
		return github_com_mwitkow_go_proto_validators.FieldError("Side", fmt.Errorf(`value '%v' must be a valid Side field`, this.Side))
	}
	if _, ok := Order_TimeInForce_name[int32(this.TimeInForce)]; !ok {
		return github_com_mwitkow_go_proto_validators.FieldError("TimeInForce", fmt.Errorf(`value '%v' must be a valid Order_TimeInForce field`, this.TimeInForce))
	}
	if _, ok := Order_Type_name[int32(this.Type)]; !ok {
		return github_com_mwitkow_go_proto_validators.FieldError("Type", fmt.Errorf(`value '%v' must be a valid Order_Type field`, this.Type))
	}
	return nil
}
func (this *OrderCancellation) Validate() error {
	if this.OrderID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("OrderID", fmt.Errorf(`value '%v' must not be an empty string`, this.OrderID))
	}
	if this.MarketID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MarketID", fmt.Errorf(`value '%v' must not be an empty string`, this.MarketID))
	}
	if this.PartyID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartyID", fmt.Errorf(`value '%v' must not be an empty string`, this.PartyID))
	}
	return nil
}
func (this *Account) Validate() error {
	return nil
}
func (this *FinancialAmount) Validate() error {
	return nil
}
func (this *Transfer) Validate() error {
	if this.Amount != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Amount); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Amount", err)
		}
	}
	return nil
}
func (this *TransferRequest) Validate() error {
	for _, item := range this.FromAccount {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("FromAccount", err)
			}
		}
	}
	for _, item := range this.ToAccount {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ToAccount", err)
			}
		}
	}
	return nil
}
func (this *LedgerEntry) Validate() error {
	return nil
}
func (this *TransferBalance) Validate() error {
	if this.Account != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Account); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Account", err)
		}
	}
	return nil
}
func (this *TransferResponse) Validate() error {
	for _, item := range this.Transfers {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Transfers", err)
			}
		}
	}
	for _, item := range this.Balances {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Balances", err)
			}
		}
	}
	return nil
}
func (this *MarginLevels) Validate() error {
	return nil
}
func (this *MarketData) Validate() error {
	return nil
}
func (this *ErrorDetail) Validate() error {
	return nil
}
func (this *SignedBundle) Validate() error {
	return nil
}
