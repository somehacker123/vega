package statevar

import (
	"code.vegaprotocol.io/vega/types/num"
)

// value is an interface for representing differnet types of floating point scalars vectors and matrices.
type value interface {
	Equals(other value) bool
	WithinTolerance(other value, tolerance float64) bool
	ToDecimal() DecimalValue
}

// KeyValueBundle is a slice of key value and their expected tolerances.
type KeyValueBundle struct {
	KVT []KeyValueTol
}

// ToDecimal converts a key value bundle to its decimal counterpart.
func (kvb *KeyValueBundle) ToDecimal() *KeyValueResult {
	res := &KeyValueResult{
		KeyDecimalValue: map[string]DecimalValue{},
	}
	for _, kv := range kvb.KVT {
		res.KeyDecimalValue[kv.Key] = kv.Val.ToDecimal()
	}
	return res
}

// WithinTolerance returns true if the two bundles have the same keys, same tolerances and the values at the same index are with the tolerance of each other
func (kvb *KeyValueBundle) WithinTolerance(other *KeyValueBundle) bool {
	if len(kvb.KVT) != len(other.KVT) {
		return false
	}
	for i, kv := range kvb.KVT {
		if kv.Key != other.KVT[i].Key {
			return false
		}
		if kv.Tolerance != other.KVT[i].Tolerance {
			return false
		}

		if !kv.Val.WithinTolerance(other.KVT[i].Val, kv.Tolerance) {
			return false
		}
	}
	return true
}

// Equals returns true of the two bundles have the same keys in the same order and the values in the same index are equal
func (kvb *KeyValueBundle) Equals(other *KeyValueBundle) bool {
	if len(kvb.KVT) != len(other.KVT) {
		return false
	}
	for i, kv := range kvb.KVT {
		if kv.Key != other.KVT[i].Key {
			return false
		}
		if !kv.Val.Equals(other.KVT[i].Val) {
			return false
		}
	}
	return true
}

type KeyValueTol struct {
	Key       string  // the name of the key
	Val       value   // the floating point value (scalar, vector, matrix)
	Tolerance float64 // the tolerance to use in comparison
}

// the result of a state variable is keyed by the name and the value is a decimal value (scalar/vector/matrix)
type KeyValueResult struct {
	KeyDecimalValue map[string]DecimalValue
}

type DecimalValue interface{}

type DecimalScalarValue struct {
	Value num.Decimal
}

type DecimalVectorValue struct {
	Value []num.Decimal
}

type DecimalMatrixValue struct {
	Value [][]num.Decimal
}
