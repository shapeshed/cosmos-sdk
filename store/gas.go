package store

import (
	"fmt"
	"math"
)

// Gas defines type alias of uint64 for gas consumption. Gas is measured by the
// SDK for store operations such as Get and Set calls. In addition, callers have
// the ability to explicitly charge gas for costly operations such as signature
// verification.
type Gas uint64

// Gas consumption descriptors.
const (
	GasIterNextCostFlatDesc = "IterNextFlat"
	GasValuePerByteDesc     = "ValuePerByte"
	GasWritePerByteDesc     = "WritePerByte"
	GasReadPerByteDesc      = "ReadPerByte"
	GasWriteCostFlatDesc    = "WriteFlat"
	GasReadCostFlatDesc     = "ReadFlat"
	GasHasDesc              = "Has"
	GasDeleteDesc           = "Delete"
)

type (
	// ErrorNegativeGasConsumed defines an error thrown when the amount of gas refunded
	// results in a negative gas consumed amount.
	ErrorNegativeGasConsumed struct {
		Descriptor string
	}

	// ErrorOutOfGas defines an error thrown when an action results in out of gas.
	ErrorOutOfGas struct {
		Descriptor string
	}

	// ErrorGasOverflow defines an error thrown when an action results gas consumption
	// unsigned integer overflow.
	ErrorGasOverflow struct {
		Descriptor string
	}
)

// GasMeter defines an interface for gas consumption tracking.
type GasMeter interface {
	// GasConsumed returns the amount of gas consumed so far.
	GasConsumed() Gas
	// GasConsumedToLimit returns the gas limit if gas consumed is past the limit,
	// otherwise it returns the consumed gas so far.
	GasConsumedToLimit() Gas
	// GasRemaining returns the gas left in the GasMeter.
	GasRemaining() Gas
	// Limit returns the gas limit (if any).
	Limit() Gas
	// ConsumeGas adds the given amount of gas to the gas consumed and should panic
	// if it overflows the gas limit (if any).
	ConsumeGas(amount Gas, descriptor string)
	// RefundGas will deduct the given amount from the gas consumed so far. If the
	// amount is greater than the gas consumed, the function should panic.
	RefundGas(amount Gas, descriptor string)
	// IsPastLimit returns <true> if the gas consumed so far is past the limit (if any),
	// otherwise it returns <false>.
	IsPastLimit() bool
	// IsOutOfGas returns <true> if the gas consumed so far is greater than or equal
	// to gas limit (if any), otherwise it returns <false>.
	IsOutOfGas() bool

	fmt.Stringer
}

// GasConfig defines gas cost for each operation on a KVStore.
type GasConfig struct {
	HasCost          Gas
	DeleteCost       Gas
	ReadCostFlat     Gas
	ReadCostPerByte  Gas
	WriteCostFlat    Gas
	WriteCostPerByte Gas
	IterNextCostFlat Gas
}

// DefaultGasConfig returns a default GasConfig for gas metering.
//
// Note, these values are essentially arbitrary. They are not based on any specific
// computation or measurements, but mainly reflect relative costs, i.e. writes
// should be more expensive than reads.
func DefaultGasConfig() GasConfig {
	return GasConfig{
		HasCost:          1000,
		DeleteCost:       1000,
		ReadCostFlat:     1000,
		ReadCostPerByte:  3,
		WriteCostFlat:    2000,
		WriteCostPerByte: 30,
		IterNextCostFlat: 30,
	}
}

// defaultGasMeter defines a default implementation of a GasMeter.
type defaultGasMeter struct {
	limit    Gas
	consumed Gas
}

// NewGasMeter returns a reference to a GasMeter with the provided limit.
func NewGasMeter(limit Gas) GasMeter {
	return &defaultGasMeter{
		limit: limit,
	}
}

func (gm *defaultGasMeter) GasConsumed() Gas {
	return gm.consumed
}

// NOTE: This behavior should only be called when recovering from a panic when
// BlockGasMeter consumes gas past the gas limit.
func (gm *defaultGasMeter) GasConsumedToLimit() Gas {
	if gm.IsPastLimit() {
		return gm.limit
	}

	return gm.consumed
}

func (gm *defaultGasMeter) GasRemaining() Gas {
	if gm.IsPastLimit() {
		return 0
	}

	return gm.limit - gm.consumed
}

func (gm *defaultGasMeter) Limit() Gas {
	return gm.limit
}

func (gm *defaultGasMeter) ConsumeGas(amount Gas, descriptor string) {
	var overflow bool

	gm.consumed, overflow = addGasOverflow(gm.consumed, amount)
	if overflow {
		gm.consumed = math.MaxUint64
		panic(ErrorGasOverflow{descriptor})
	}

	if gm.consumed > gm.limit {
		panic(ErrorOutOfGas{descriptor})
	}
}

func (gm *defaultGasMeter) RefundGas(amount Gas, descriptor string) {
	if gm.consumed < amount {
		panic(ErrorNegativeGasConsumed{Descriptor: descriptor})
	}

	gm.consumed -= amount
}

func (gm *defaultGasMeter) IsPastLimit() bool {
	return gm.consumed > gm.limit
}

func (gm *defaultGasMeter) IsOutOfGas() bool {
	return gm.consumed >= gm.limit
}

func (gm *defaultGasMeter) String() string {
	return fmt.Sprintf("GasMeter{limit: %d, consumed: %d}", gm.limit, gm.consumed)
}

// addGasOverflow performs the addition operation on two uint64 integers and
// returns a boolean on whether or not the result overflows.
func addGasOverflow(a, b Gas) (Gas, bool) {
	if math.MaxUint64-a < b {
		return 0, true
	}

	return a + b, false
}
