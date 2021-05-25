package contexts

import (
	"crypto/elliptic"
	basicMath "math"
	"math/big"

	"github.com/ElrondNetwork/arwen-wasm-vm/arwen"
	"github.com/ElrondNetwork/arwen-wasm-vm/math"
)

const maxBigIntByteLenForNormalCost = 32

type bigIntMap map[int32]*big.Int
type ellipticCurveMap map[int32]*elliptic.CurveParams

type managedTypeContext struct {
	host             arwen.VMHost
	bigIntValues     bigIntMap
	ecValues         ellipticCurveMap
	ecStateStack     []ellipticCurveMap
	bigIntStateStack []bigIntMap
}

// NewBigIntContext creates a new bigIntContext
func NewManagedTypeContext(host arwen.VMHost) (*managedTypeContext, error) {
	context := &managedTypeContext{
		host:             host,
		bigIntValues:     make(bigIntMap),
		ecValues:         make(ellipticCurveMap),
		ecStateStack:     make([]ellipticCurveMap, 0),
		bigIntStateStack: make([]bigIntMap, 0),
	}

	return context, nil
}

// InitState initializes the underlying values map
func (context *managedTypeContext) InitState() {
	context.bigIntValues = make(bigIntMap)
	context.ecValues = make(ellipticCurveMap)
}

// PushState appends the values map to the state stack
func (context *managedTypeContext) PushState() {
	newBigIntState, newEcState := context.clone()
	context.bigIntStateStack = append(context.bigIntStateStack, newBigIntState)
	context.ecStateStack = append(context.ecStateStack, newEcState)
}

// PopSetActiveState removes the latest entry from the state stack and sets it as the current values map
func (context *managedTypeContext) PopSetActiveState() {
	bigIntStateStackLen := len(context.bigIntStateStack)
	ecStateStackLen := len(context.ecStateStack)
	if bigIntStateStackLen == 0 && ecStateStackLen == 0 {
		return
	}
	prevBigIntValues := context.bigIntStateStack[bigIntStateStackLen-1]
	context.bigIntStateStack = context.bigIntStateStack[:bigIntStateStackLen-1]
	context.bigIntValues = prevBigIntValues

	prevEcValues := context.ecStateStack[ecStateStackLen-1]
	context.ecStateStack = context.ecStateStack[:ecStateStackLen-1]
	context.ecValues = prevEcValues
}

// PopDiscard removes the latest entry from the state stack
func (context *managedTypeContext) PopDiscard() {
	bigIntStateStackLen := len(context.bigIntStateStack)
	ecStateStackLen := len(context.ecStateStack)
	if bigIntStateStackLen == 0 && ecStateStackLen == 0 {
		return
	}

	context.ecStateStack = context.ecStateStack[:ecStateStackLen-1]
	context.bigIntStateStack = context.bigIntStateStack[:bigIntStateStackLen-1]
}

// ClearStateStack initializes the state stack
func (context *managedTypeContext) ClearStateStack() {
	context.bigIntStateStack = make([]bigIntMap, 0)
	context.ecStateStack = make([]ellipticCurveMap, 0)
}

func (context *managedTypeContext) clone() (bigIntMap, ellipticCurveMap) {
	newBigIntState := make(bigIntMap, len(context.bigIntValues))
	newEcState := make(ellipticCurveMap, len(context.ecValues))
	for bigIntHandle, bigInt := range context.bigIntValues {
		newBigIntState[bigIntHandle] = big.NewInt(0).Set(bigInt)
	}
	for ecHandle, ec := range context.ecValues {
		newEcState[ecHandle] = ec
	}
	return newBigIntState, newEcState
}

// IsInterfaceNil returns true if there is no value under the interface
func (context *managedTypeContext) IsInterfaceNil() bool {
	return context == nil
}

// ConsumeGasForBigIntCopy uses gas for Copy operations
func (context *managedTypeContext) ConsumeGasForBigIntCopy(values ...*big.Int) {
	for _, val := range values {
		byteLen := val.BitLen() / 8
		context.ConsumeGasForThisIntNumberOfBytes(byteLen)
	}
}

// ConsumeGasForThisIntNumberOfBytes uses gas for the number of bytes given.
func (context *managedTypeContext) ConsumeGasForThisIntNumberOfBytes(byteLen int) {
	metering := context.host.Metering()
	if byteLen > maxBigIntByteLenForNormalCost {
		metering.UseGas(math.MulUint64(uint64(byteLen), metering.GasSchedule().BaseOperationCost.DataCopyPerByte))
	}
}

// ConsumeGasForThisBigIntNumberOfBytes uses gas for the number of bytes given that are being copied.
func (context *managedTypeContext) ConsumeGasForThisBigIntNumberOfBytes(byteLen *big.Int) {
	metering := context.host.Metering()
	DataCopyPerByte := metering.GasSchedule().BaseOperationCost.DataCopyPerByte

	gasToUseBigInt := big.NewInt(0).Mul(byteLen, big.NewInt(int64(DataCopyPerByte)))
	maxGasBigInt := big.NewInt(0).SetUint64(basicMath.MaxUint64)
	gasToUse := uint64(basicMath.MaxUint64)
	if gasToUseBigInt.Cmp(maxGasBigInt) < 0 {
		gasToUse = gasToUseBigInt.Uint64()
	}
	metering.UseGas(gasToUse)
}

// BIGINT

// GetOneOrCreate returns the value at the given handle. If there is no value under that value, it will set a new on with value 0.
func (context *managedTypeContext) GetBigIntOrCreate(handle int32) *big.Int {
	value, ok := context.bigIntValues[handle]
	if !ok {
		value = big.NewInt(0)
		context.bigIntValues[handle] = value
	}
	return value
}

// GetOne returns the value at the given handle. If there is no value under that handle, it will return error
func (context *managedTypeContext) GetBigInt(handle int32) (*big.Int, error) {
	value, ok := context.bigIntValues[handle]
	if !ok {
		return nil, arwen.ErrNoBigIntUnderThisHandle
	}
	return value, nil
}

// PutBigInt adds the given value to the current values map and returns the handle
func (context *managedTypeContext) PutBigInt(value int64) int32 {
	newHandle := int32(len(context.bigIntValues))
	for {
		if _, ok := context.bigIntValues[newHandle]; !ok {
			break
		}
		newHandle++
	}
	context.bigIntValues[newHandle] = big.NewInt(value)
	return newHandle
}

// ELLIPTIC CURVES

// GetOneEllipticCurve returns the elliptic curve under the given handle. If there is no value under that handle, it will return error
func (context *managedTypeContext) GetOneEllipticCurve(handle int32) (*elliptic.CurveParams, error) {
	curve, ok := context.ecValues[handle]
	if !ok {
		return nil, arwen.ErrNoEllipticCurveUnderThisHandle
	}
	return curve, nil
}

// PutEllipticCurve adds the given elliptic curve to the current ecValues map and returns the handle
func (context *managedTypeContext) PutEllipticCurve(curve *elliptic.CurveParams) int32 {
	newHandle := int32(len(context.ecValues))
	for {
		if _, ok := context.ecValues[newHandle]; !ok {
			break
		}
		newHandle++
	}
	context.ecValues[newHandle] = &elliptic.CurveParams{P: curve.P, N: curve.N, B: curve.B, Gx: curve.Gx, Gy: curve.Gy, BitSize: curve.BitSize, Name: curve.Name}
	return newHandle
}
