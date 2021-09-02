package elrondapitest

import (
	"fmt"
	"math"
	"math/big"
	"testing"

	"github.com/ElrondNetwork/arwen-wasm-vm/v1_4/arwen"
	contextmock "github.com/ElrondNetwork/arwen-wasm-vm/v1_4/mock/context"
	test "github.com/ElrondNetwork/arwen-wasm-vm/v1_4/testcommon"
	"github.com/stretchr/testify/require"
	"gopkg.in/go-playground/assert.v1"
)

var repsArgument = []byte{0, 0, 0, byte(numberOfReps)}
var floatArgument1 = []byte{1, 10, 0, 0, 0, 100, 0, 0, 0, 108, 136, 217, 65, 19, 144, 71, 160, 0} // equal to 1.73476272346174595037472187482e+32
var floatArgument2 = []byte{1, 10, 0, 0, 0, 53, 0, 0, 0, 11, 190, 100, 79, 147, 188, 10, 8, 0}

func TestBigFloats_New(t *testing.T) {
	test.BuildInstanceCallTest(t).
		WithContracts(
			test.CreateInstanceContract(test.ParentAddress).
				WithCode(test.GetTestSCCode("big-floats", "../../"))).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithGasProvided(100000).
			WithFunction("BigFloatNewTest").
			WithArguments(repsArgument).
			Build()).
		AndAssertResults(func(host arwen.VMHost, stubBlockchainHook *contextmock.BlockchainHookStub, verify *test.VMOutputVerifier) {
			verify.
				Ok().
				ReturnData(
					[]byte{byte(numberOfReps - 1)})
		})
}

func TestBigFloats_NewFromFrac(t *testing.T) {
	test.BuildInstanceCallTest(t).
		WithContracts(
			test.CreateInstanceContract(test.ParentAddress).
				WithCode(test.GetTestSCCode("big-floats", "../../"))).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithGasProvided(100000).
			WithFunction("BigFloatNewFromFracTest").
			WithArguments(repsArgument).
			Build()).
		AndAssertResults(func(host arwen.VMHost, stubBlockchainHook *contextmock.BlockchainHookStub, verify *test.VMOutputVerifier) {
			verify.
				Ok().
				ReturnData(
					[]byte{byte(numberOfReps - 1)})
		})
}

func TestBigFloats_Add(t *testing.T) {
	test.BuildInstanceCallTest(t).
		WithContracts(
			test.CreateInstanceContract(test.ParentAddress).
				WithCode(test.GetTestSCCode("big-floats", "../../"))).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithGasProvided(100000).
			WithFunction("BigFloatAddTest").
			WithArguments(repsArgument,
				floatArgument1).
			Build()).
		AndAssertResults(func(host arwen.VMHost, stubBlockchainHook *contextmock.BlockchainHookStub, verify *test.VMOutputVerifier) {
			bigFloatValue := new(big.Float)
			_ = bigFloatValue.GobDecode(floatArgument1)
			for i := 0; i < numberOfReps; i++ {
				bigFloatValue.Add(bigFloatValue, bigFloatValue)
			}
			floatBuffer, _ := bigFloatValue.GobEncode()
			verify.
				Ok().
				ReturnData(
					floatBuffer)
		})
}

func TestBigFloats_Sub(t *testing.T) {
	test.BuildInstanceCallTest(t).
		WithContracts(
			test.CreateInstanceContract(test.ParentAddress).
				WithCode(test.GetTestSCCode("big-floats", "../../"))).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithGasProvided(100000).
			WithFunction("BigFloatSubTest").
			WithArguments(repsArgument,
				floatArgument1).
			Build()).
		AndAssertResults(func(host arwen.VMHost, stubBlockchainHook *contextmock.BlockchainHookStub, verify *test.VMOutputVerifier) {
			bigFloatValue := new(big.Float)
			_ = bigFloatValue.GobDecode(floatArgument1)
			for i := 0; i < numberOfReps; i++ {
				bigFloatValue.Sub(bigFloatValue, bigFloatValue)
			}
			floatBuffer, _ := bigFloatValue.GobEncode()
			verify.
				Ok().
				ReturnData(
					floatBuffer)
		})
}

func TestBigFloats_Success_Mul(t *testing.T) {
	numberOfReps := 10
	repsArgument := []byte{0, 0, 0, byte(numberOfReps)}
	test.BuildInstanceCallTest(t).
		WithContracts(
			test.CreateInstanceContract(test.ParentAddress).
				WithCode(test.GetTestSCCode("big-floats", "../../"))).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithGasProvided(100000).
			WithFunction("BigFloatMulTest").
			WithArguments(repsArgument,
				floatArgument1).
			Build()).
		AndAssertResults(func(host arwen.VMHost, stubBlockchainHook *contextmock.BlockchainHookStub, verify *test.VMOutputVerifier) {
			bigFloatValue := new(big.Float)
			err := bigFloatValue.GobDecode(floatArgument1)
			require.Nil(t, err)
			for i := 0; i < numberOfReps; i++ {
				resultMul := new(big.Float).Mul(bigFloatValue, bigFloatValue)
				bigFloatValue.Set(resultMul)
			}
			floatBuffer, _ := bigFloatValue.GobEncode()
			fmt.Println(floatBuffer)
			verify.
				Ok().
				ReturnData(
					floatBuffer)
		})
}

func TestBigFloats_FailExecution_Mul(t *testing.T) {
	numberOfReps := 30
	repsArgument := []byte{0, 0, 0, byte(numberOfReps)}
	test.BuildInstanceCallTest(t).
		WithContracts(
			test.CreateInstanceContract(test.ParentAddress).
				WithCode(test.GetTestSCCode("big-floats", "../../"))).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithGasProvided(100000).
			WithFunction("BigFloatMulTest").
			WithArguments(repsArgument,
				floatArgument1).
			Build()).
		AndAssertResults(func(host arwen.VMHost, stubBlockchainHook *contextmock.BlockchainHookStub, verify *test.VMOutputVerifier) {
			verify.
				ReturnCode(10).
				ReturnData()
		})
}

func TestBigFloats_Div(t *testing.T) {
	test.BuildInstanceCallTest(t).
		WithContracts(
			test.CreateInstanceContract(test.ParentAddress).
				WithCode(test.GetTestSCCode("big-floats", "../../"))).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithGasProvided(100000).
			WithFunction("BigFloatDivTest").
			WithArguments(repsArgument,
				floatArgument1,
				floatArgument2).
			Build()).
		AndAssertResults(func(host arwen.VMHost, stubBlockchainHook *contextmock.BlockchainHookStub, verify *test.VMOutputVerifier) {
			numerator := new(big.Float)
			_ = numerator.GobDecode(floatArgument1)
			assert.Equal(t, uint(100), numerator.Prec())
			denominator := new(big.Float)
			_ = denominator.GobDecode(floatArgument2)
			assert.Equal(t, uint(53), denominator.Prec())
			for i := 0; i < numberOfReps; i++ {
				resultMul := new(big.Float).Quo(numerator, denominator)
				numerator.Set(resultMul)
			}
			assert.Equal(t, uint(100), numerator.Prec())
			floatBuffer, _ := numerator.GobEncode()
			verify.
				Ok().
				ReturnData(
					floatBuffer)
		})
}

func TestBigFloats_Truncate(t *testing.T) {
	test.BuildInstanceCallTest(t).
		WithContracts(
			test.CreateInstanceContract(test.ParentAddress).
				WithCode(test.GetTestSCCode("big-floats", "../../"))).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithGasProvided(100000).
			WithFunction("BigFloatTruncateTest").
			WithArguments(repsArgument,
				floatArgument1,
				floatArgument2).
			Build()).
		AndAssertResults(func(host arwen.VMHost, stubBlockchainHook *contextmock.BlockchainHookStub, verify *test.VMOutputVerifier) {
			numerator := new(big.Float)
			_ = numerator.GobDecode(floatArgument1)
			assert.Equal(t, uint(100), numerator.Prec())
			denominator := new(big.Float)
			_ = denominator.GobDecode(floatArgument2)
			assert.Equal(t, uint(53), denominator.Prec())
			for i := 0; i < numberOfReps; i++ {
				rDiv := big.NewInt(0)
				numerator.Int(rDiv)
				numerator.SetInt(rDiv)
				numerator.Sub(numerator, denominator)
			}
			assert.Equal(t, uint(100), numerator.Prec())
			floatBuffer, _ := numerator.GobEncode()
			verify.
				Ok().
				ReturnData(
					floatBuffer)
		})
}

func TestBigFloats_Mod(t *testing.T) {
	test.BuildInstanceCallTest(t).
		WithContracts(
			test.CreateInstanceContract(test.ParentAddress).
				WithCode(test.GetTestSCCode("big-floats", "../../"))).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithGasProvided(100000).
			WithFunction("BigFloatModTest").
			WithArguments(repsArgument,
				floatArgument1,
				floatArgument2).
			Build()).
		AndAssertResults(func(host arwen.VMHost, stubBlockchainHook *contextmock.BlockchainHookStub, verify *test.VMOutputVerifier) {
			numerator := big.NewFloat(0)
			denominator := big.NewFloat(0)
			numerator.SetPrec(0)
			_ = numerator.GobDecode(floatArgument1)
			assert.Equal(t, uint(100), numerator.Prec())
			denominator.SetPrec(0)
			_ = denominator.GobDecode(floatArgument2)
			assert.Equal(t, uint(53), denominator.Prec())
			result := big.NewFloat(0)
			result.SetPrec(0)
			for i := 0; i < numberOfReps; i++ {
				result.Quo(numerator, denominator)
				rdiv := big.NewInt(0)
				result.Int(rdiv)
				result.Sub(result, new(big.Float).SetInt(rdiv))
				numerator.Sub(numerator, denominator)
			}
			assert.Equal(t, uint(100), numerator.Prec())
			assert.Equal(t, uint(100), result.Prec())
			floatBuffer, _ := result.GobEncode()
			verify.
				Ok().
				ReturnData(
					floatBuffer)
		})
}

func TestBigFloats_Abs(t *testing.T) {
	bigFloatArguments := make([][]byte, numberOfReps+1)
	for i := range bigFloatArguments {
		bigFloatArguments[i] = make([]byte, 0)
	}
	bigFloatArguments[0] = repsArgument
	for i := 0; i < numberOfReps; i++ {
		floatValue := big.NewFloat(-float64(i) - 1)
		encodedFloat, _ := floatValue.GobEncode()
		bigFloatArguments[i+1] = encodedFloat
	}

	test.BuildInstanceCallTest(t).
		WithContracts(
			test.CreateInstanceContract(test.ParentAddress).
				WithCode(test.GetTestSCCode("big-floats", "../../"))).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithGasProvided(100000).
			WithFunction("BigFloatAbsTest").
			WithArguments(bigFloatArguments...).
			Build()).
		AndAssertResults(func(host arwen.VMHost, stubBlockchainHook *contextmock.BlockchainHookStub, verify *test.VMOutputVerifier) {
			encodedLastFloat := bigFloatArguments[numberOfReps]
			lastFloat := new(big.Float)
			_ = lastFloat.GobDecode(encodedLastFloat)
			absLastFloat := new(big.Float)
			absLastFloat.Abs(lastFloat)
			encodedAbsFloat, _ := absLastFloat.GobEncode()
			verify.
				ReturnMessage("").
				Ok().
				ReturnData(encodedAbsFloat)

		})
}

func TestBigFloats_Neg(t *testing.T) {
	bigFloatArguments := make([][]byte, numberOfReps+1)
	for i := range bigFloatArguments {
		bigFloatArguments[i] = make([]byte, 0)
	}
	bigFloatArguments[0] = repsArgument
	for i := 0; i < numberOfReps; i++ {
		floatValue := big.NewFloat(-float64(i) - 1)
		encodedFloat, _ := floatValue.GobEncode()
		bigFloatArguments[i+1] = encodedFloat
	}

	test.BuildInstanceCallTest(t).
		WithContracts(
			test.CreateInstanceContract(test.ParentAddress).
				WithCode(test.GetTestSCCode("big-floats", "../../"))).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithGasProvided(100000).
			WithFunction("BigFloatNegTest").
			WithArguments(bigFloatArguments...).
			Build()).
		AndAssertResults(func(host arwen.VMHost, stubBlockchainHook *contextmock.BlockchainHookStub, verify *test.VMOutputVerifier) {
			encodedLastFloat := bigFloatArguments[numberOfReps]
			lastFloat := new(big.Float)
			_ = lastFloat.GobDecode(encodedLastFloat)
			negLastFloat := new(big.Float)
			negLastFloat.Neg(lastFloat)
			encodedNegFloat, _ := negLastFloat.GobEncode()
			verify.
				Ok().
				ReturnData(encodedNegFloat)
		})
}

func TestBigFloats_Cmp(t *testing.T) {
	bigFloatArguments := make([][]byte, 2*numberOfReps+1)
	for i := range bigFloatArguments {
		bigFloatArguments[i] = make([]byte, 0)
	}
	bigFloatArguments[0] = repsArgument
	argsCounter := 1
	for i := 0; i < numberOfReps; i++ {
		floatValue := big.NewFloat(-float64(i) - 1)
		encodedFloat, _ := floatValue.GobEncode()
		bigFloatArguments[argsCounter] = encodedFloat
		absFloatValue := new(big.Float).Neg(floatValue)
		encodedAbsFloat, _ := absFloatValue.GobEncode()
		bigFloatArguments[argsCounter+1] = encodedAbsFloat
		argsCounter += 2
	}

	test.BuildInstanceCallTest(t).
		WithContracts(
			test.CreateInstanceContract(test.ParentAddress).
				WithCode(test.GetTestSCCode("big-floats", "../../"))).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithGasProvided(100000).
			WithFunction("BigFloatCmpTest").
			WithArguments(bigFloatArguments...).
			Build()).
		AndAssertResults(func(host arwen.VMHost, stubBlockchainHook *contextmock.BlockchainHookStub, verify *test.VMOutputVerifier) {
			encodedLastFloat := bigFloatArguments[numberOfReps*2]
			lastFloat := new(big.Float)
			_ = lastFloat.GobDecode(encodedLastFloat)
			encodedPreviousLastFloat := bigFloatArguments[numberOfReps*2-1]
			previousLastFloat := new(big.Float)
			_ = previousLastFloat.GobDecode(encodedPreviousLastFloat)
			cmpResult := previousLastFloat.Cmp(lastFloat)
			verify.
				Ok().
				ReturnData([]byte{byte(cmpResult)})
		})
}

func TestBigFloats_Sign(t *testing.T) {
	bigFloatArguments := make([][]byte, numberOfReps+1)
	for i := range bigFloatArguments {
		bigFloatArguments[i] = make([]byte, 0)
	}
	bigFloatArguments[0] = repsArgument
	for i := 0; i < numberOfReps; i++ {
		floatValue := big.NewFloat(-float64(i) - 1)
		encodedFloat, _ := floatValue.GobEncode()
		bigFloatArguments[i+1] = encodedFloat
	}

	test.BuildInstanceCallTest(t).
		WithContracts(
			test.CreateInstanceContract(test.ParentAddress).
				WithCode(test.GetTestSCCode("big-floats", "../../"))).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithGasProvided(100000).
			WithFunction("BigFloatSignTest").
			WithArguments(bigFloatArguments...).
			Build()).
		AndAssertResults(func(host arwen.VMHost, stubBlockchainHook *contextmock.BlockchainHookStub, verify *test.VMOutputVerifier) {
			encodedLastFloat := bigFloatArguments[numberOfReps]
			lastFloat := new(big.Float)
			_ = lastFloat.GobDecode(encodedLastFloat)
			verify.
				Ok().
				ReturnData([]byte{byte(lastFloat.Sign())})
		})
}

func TestBigFloats_Clone(t *testing.T) {
	bigFloatArguments := make([][]byte, numberOfReps+1)
	for i := range bigFloatArguments {
		bigFloatArguments[i] = make([]byte, 0)
	}
	bigFloatArguments[0] = repsArgument
	for i := 0; i < numberOfReps; i++ {
		floatValue := big.NewFloat(-float64(i) - 1)
		encodedFloat, _ := floatValue.GobEncode()
		bigFloatArguments[i+1] = encodedFloat
	}

	test.BuildInstanceCallTest(t).
		WithContracts(
			test.CreateInstanceContract(test.ParentAddress).
				WithCode(test.GetTestSCCode("big-floats", "../../"))).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithGasProvided(100000).
			WithFunction("BigFloatCloneTest").
			WithArguments(bigFloatArguments...).
			Build()).
		AndAssertResults(func(host arwen.VMHost, stubBlockchainHook *contextmock.BlockchainHookStub, verify *test.VMOutputVerifier) {
			encodedLastFloat := bigFloatArguments[numberOfReps]
			lastFloat := new(big.Float)
			_ = lastFloat.GobDecode(encodedLastFloat)
			verify.
				Ok().
				ReturnData(encodedLastFloat)
		})
}

func TestBigFloats_Sqrt(t *testing.T) {
	bigFloatArguments := make([][]byte, numberOfReps+1)
	for i := range bigFloatArguments {
		bigFloatArguments[i] = make([]byte, 0)
	}
	bigFloatArguments[0] = repsArgument
	for i := 0; i < numberOfReps; i++ {
		floatValue := big.NewFloat(float64(i) + 1)
		encodedFloat, _ := floatValue.GobEncode()
		bigFloatArguments[i+1] = encodedFloat
	}

	test.BuildInstanceCallTest(t).
		WithContracts(
			test.CreateInstanceContract(test.ParentAddress).
				WithCode(test.GetTestSCCode("big-floats", "../../"))).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithGasProvided(100000).
			WithFunction("BigFloatSqrtTest").
			WithArguments(bigFloatArguments...).
			Build()).
		AndAssertResults(func(host arwen.VMHost, stubBlockchainHook *contextmock.BlockchainHookStub, verify *test.VMOutputVerifier) {
			encodedLastFloat := bigFloatArguments[numberOfReps]
			lastFloat := new(big.Float)
			_ = lastFloat.GobDecode(encodedLastFloat)
			sqrtFloat := new(big.Float).Sqrt(lastFloat)
			encodedSqrtFloat, _ := sqrtFloat.GobEncode()
			verify.
				Ok().
				ReturnData(encodedSqrtFloat)
		})
}

func TestBigFloats_Log2(t *testing.T) {
	bigFloatArguments := make([][]byte, numberOfReps+1)
	for i := range bigFloatArguments {
		bigFloatArguments[i] = make([]byte, 0)
	}
	bigFloatArguments[0] = repsArgument
	for i := 0; i < numberOfReps; i++ {
		floatValue := big.NewFloat(float64(i) + 1)
		encodedFloat, _ := floatValue.GobEncode()
		bigFloatArguments[i+1] = encodedFloat
	}

	test.BuildInstanceCallTest(t).
		WithContracts(
			test.CreateInstanceContract(test.ParentAddress).
				WithCode(test.GetTestSCCode("big-floats", "../../"))).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithGasProvided(100000).
			WithFunction("BigFloatLog2Test").
			WithArguments(bigFloatArguments...).
			Build()).
		AndAssertResults(func(host arwen.VMHost, stubBlockchainHook *contextmock.BlockchainHookStub, verify *test.VMOutputVerifier) {
			encodedLastFloat := bigFloatArguments[numberOfReps]
			lastFloat := new(big.Float)
			_ = lastFloat.GobDecode(encodedLastFloat)
			bigIntOp := new(big.Int)
			lastFloat.Int(bigIntOp)
			verify.
				Ok().
				ReturnData([]byte{byte(bigIntOp.BitLen() - 1)})
		})
}

func TestBigFloats_Pow(t *testing.T) {
	bigFloatArguments := make([][]byte, numberOfReps+1)
	for i := range bigFloatArguments {
		bigFloatArguments[i] = make([]byte, 0)
	}
	bigFloatArguments[0] = []byte{0, 0, 0, byte(3)}
	for i := 0; i < numberOfReps; i++ {
		floatValue := big.NewFloat(1.6)
		encodedFloat, _ := floatValue.GobEncode()
		bigFloatArguments[i+1] = encodedFloat
	}

	test.BuildInstanceCallTest(t).
		WithContracts(
			test.CreateInstanceContract(test.ParentAddress).
				WithCode(test.GetTestSCCode("big-floats", "../../"))).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithGasProvided(100000).
			WithFunction("BigFloatPowTest").
			WithArguments(bigFloatArguments...).
			Build()).
		AndAssertResults(func(host arwen.VMHost, stubBlockchainHook *contextmock.BlockchainHookStub, verify *test.VMOutputVerifier) {
			resultFloat := big.NewFloat(1.6)
			intermediaryFloat := new(big.Float).Mul(resultFloat, resultFloat)
			resultFloat.Set(intermediaryFloat)
			encodedResult, _ := resultFloat.GobEncode()
			verify.
				Ok().
				ReturnData(encodedResult)
		})
}

func TestBigFloats_Floor(t *testing.T) {
	bigFloatArguments := make([][]byte, numberOfReps+1)
	for i := range bigFloatArguments {
		bigFloatArguments[i] = make([]byte, 0)
	}
	bigFloatArguments[0] = repsArgument
	for i := 0; i < numberOfReps; i++ {
		floatValue := big.NewFloat((float64(i) + 2) / (float64(i) + 1))
		encodedFloat, _ := floatValue.GobEncode()
		bigFloatArguments[i+1] = encodedFloat
	}

	test.BuildInstanceCallTest(t).
		WithContracts(
			test.CreateInstanceContract(test.ParentAddress).
				WithCode(test.GetTestSCCode("big-floats", "../../"))).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithGasProvided(100000).
			WithFunction("BigFloatFloorTest").
			WithArguments(bigFloatArguments...).
			Build()).
		AndAssertResults(func(host arwen.VMHost, stubBlockchainHook *contextmock.BlockchainHookStub, verify *test.VMOutputVerifier) {
			encodedLastFloat := bigFloatArguments[numberOfReps]
			lastFloat := new(big.Float)
			_ = lastFloat.GobDecode(encodedLastFloat)
			bigIntOp := new(big.Int)
			lastFloat.Int(bigIntOp)
			verify.
				Ok().
				ReturnData(bigIntOp.Bytes())
		})
}

func TestBigFloats_Ceil(t *testing.T) {
	bigFloatArguments := make([][]byte, numberOfReps+1)
	for i := range bigFloatArguments {
		bigFloatArguments[i] = make([]byte, 0)
	}
	bigFloatArguments[0] = repsArgument
	for i := 0; i < numberOfReps; i++ {
		floatValue := big.NewFloat((float64(i) + 2) / (float64(i) + 1))
		encodedFloat, _ := floatValue.GobEncode()
		bigFloatArguments[i+1] = encodedFloat
	}

	test.BuildInstanceCallTest(t).
		WithContracts(
			test.CreateInstanceContract(test.ParentAddress).
				WithCode(test.GetTestSCCode("big-floats", "../../"))).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithGasProvided(100000).
			WithFunction("BigFloatCeilTest").
			WithArguments(bigFloatArguments...).
			Build()).
		AndAssertResults(func(host arwen.VMHost, stubBlockchainHook *contextmock.BlockchainHookStub, verify *test.VMOutputVerifier) {
			encodedLastFloat := bigFloatArguments[numberOfReps]
			lastFloat := new(big.Float)
			_ = lastFloat.GobDecode(encodedLastFloat)
			bigIntOp := new(big.Int)
			lastFloat.Int(bigIntOp)
			bigIntOp.Add(bigIntOp, big.NewInt(1))
			verify.
				Ok().
				ReturnData(bigIntOp.Bytes())
		})
}

func TestBigFloats_IsInt(t *testing.T) {
	bigFloatArguments := make([][]byte, numberOfReps+1)
	for i := range bigFloatArguments {
		bigFloatArguments[i] = make([]byte, 0)
	}
	bigFloatArguments[0] = repsArgument
	for i := 0; i < numberOfReps; i++ {
		floatValue := big.NewFloat((float64(i) + 2))
		encodedFloat, _ := floatValue.GobEncode()
		bigFloatArguments[i+1] = encodedFloat
	}

	test.BuildInstanceCallTest(t).
		WithContracts(
			test.CreateInstanceContract(test.ParentAddress).
				WithCode(test.GetTestSCCode("big-floats", "../../"))).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithGasProvided(100000).
			WithFunction("BigFloatIsIntTest").
			WithArguments(bigFloatArguments...).
			Build()).
		AndAssertResults(func(host arwen.VMHost, stubBlockchainHook *contextmock.BlockchainHookStub, verify *test.VMOutputVerifier) {
			encodedLastFloat := bigFloatArguments[numberOfReps]
			lastFloat := new(big.Float)
			_ = lastFloat.GobDecode(encodedLastFloat)
			isInt := -2
			if lastFloat.IsInt() {
				isInt = 1
			} else {
				isInt = 0
			}
			verify.
				Ok().
				ReturnData([]byte{byte(isInt)})
		})
}

func TestBigFloats_SetInt64(t *testing.T) {
	bigFloatArguments := make([][]byte, numberOfReps+1)
	for i := range bigFloatArguments {
		bigFloatArguments[i] = make([]byte, 0)
	}
	bigFloatArguments[0] = repsArgument
	for i := 0; i < numberOfReps; i++ {
		bigFloatArguments[i+1] = []byte{0, 0, 0, byte(i)}
	}

	test.BuildInstanceCallTest(t).
		WithContracts(
			test.CreateInstanceContract(test.ParentAddress).
				WithCode(test.GetTestSCCode("big-floats", "../../"))).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithGasProvided(100000).
			WithFunction("BigFloatSetInt64Test").
			WithArguments(bigFloatArguments...).
			Build()).
		AndAssertResults(func(host arwen.VMHost, stubBlockchainHook *contextmock.BlockchainHookStub, verify *test.VMOutputVerifier) {
			floatValue := big.NewFloat(0)
			floatValue.SetInt64(int64(numberOfReps - 1))
			encodedFloatValue, _ := floatValue.GobEncode()
			verify.
				Ok().
				ReturnData(encodedFloatValue)
		})
}

func TestBigFloats_SetBigInt(t *testing.T) {
	bigFloatArguments := make([][]byte, numberOfReps+1)
	for i := range bigFloatArguments {
		bigFloatArguments[i] = make([]byte, 0)
	}
	bigFloatArguments[0] = repsArgument
	for i := 0; i < numberOfReps; i++ {
		bigIntValue := big.NewInt(int64(i))
		bigFloatArguments[i+1] = bigIntValue.Bytes()
	}

	test.BuildInstanceCallTest(t).
		WithContracts(
			test.CreateInstanceContract(test.ParentAddress).
				WithCode(test.GetTestSCCode("big-floats", "../../"))).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithGasProvided(100000).
			WithFunction("BigFloatSetBigIntTest").
			WithArguments(bigFloatArguments...).
			Build()).
		AndAssertResults(func(host arwen.VMHost, stubBlockchainHook *contextmock.BlockchainHookStub, verify *test.VMOutputVerifier) {
			floatValue := big.NewFloat(0)
			floatValue.SetInt(big.NewInt(int64(numberOfReps) - 1))
			encodedFloatValue, _ := floatValue.GobEncode()
			verify.
				Ok().
				ReturnData(encodedFloatValue)
		})
}

func TestBigFloats_GetConstPi(t *testing.T) {
	test.BuildInstanceCallTest(t).
		WithContracts(
			test.CreateInstanceContract(test.ParentAddress).
				WithCode(test.GetTestSCCode("big-floats", "../../"))).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithGasProvided(100000).
			WithFunction("BigFloatGetConstPiTest").
			WithArguments([]byte{0, 0, 0, byte(numberOfReps)}).
			Build()).
		AndAssertResults(func(host arwen.VMHost, stubBlockchainHook *contextmock.BlockchainHookStub, verify *test.VMOutputVerifier) {
			piValue := math.Pi
			bigFloatValue := big.NewFloat(0).SetFloat64(piValue)
			encodedFloat, _ := bigFloatValue.GobEncode()
			verify.
				Ok().
				ReturnData(encodedFloat)
		})
}

func TestBigFloats_GetConstE(t *testing.T) {
	test.BuildInstanceCallTest(t).
		WithContracts(
			test.CreateInstanceContract(test.ParentAddress).
				WithCode(test.GetTestSCCode("big-floats", "../../"))).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithGasProvided(100000).
			WithFunction("BigFloatGetConstETest").
			WithArguments([]byte{0, 0, 0, byte(numberOfReps)}).
			Build()).
		AndAssertResults(func(host arwen.VMHost, stubBlockchainHook *contextmock.BlockchainHookStub, verify *test.VMOutputVerifier) {
			piValue := math.E
			bigFloatValue := big.NewFloat(0).SetFloat64(piValue)
			encodedFloat, _ := bigFloatValue.GobEncode()
			verify.
				Ok().
				ReturnData(encodedFloat)
		})
}