// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractAllocationManager

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IAllocationManagerTypesAllocateParams is an auto generated low-level Go binding around an user-defined struct.
type IAllocationManagerTypesAllocateParams struct {
	OperatorSet   OperatorSet
	Strategies    []common.Address
	NewMagnitudes []uint64
}

// IAllocationManagerTypesAllocation is an auto generated low-level Go binding around an user-defined struct.
type IAllocationManagerTypesAllocation struct {
	CurrentMagnitude uint64
	PendingDiff      *big.Int
	EffectBlock      uint32
}

// IAllocationManagerTypesCreateSetParams is an auto generated low-level Go binding around an user-defined struct.
type IAllocationManagerTypesCreateSetParams struct {
	OperatorSetId uint32
	Strategies    []common.Address
}

// IAllocationManagerTypesDeregisterParams is an auto generated low-level Go binding around an user-defined struct.
type IAllocationManagerTypesDeregisterParams struct {
	Operator       common.Address
	Avs            common.Address
	OperatorSetIds []uint32
}

// IAllocationManagerTypesRegisterParams is an auto generated low-level Go binding around an user-defined struct.
type IAllocationManagerTypesRegisterParams struct {
	Avs            common.Address
	OperatorSetIds []uint32
	Data           []byte
}

// IAllocationManagerTypesSlashingParams is an auto generated low-level Go binding around an user-defined struct.
type IAllocationManagerTypesSlashingParams struct {
	Operator      common.Address
	OperatorSetId uint32
	WadToSlash    *big.Int
	Description   string
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// ContractAllocationManagerMetaData contains all meta data concerning the ContractAllocationManager contract.
var ContractAllocationManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_delegation\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_DEALLOCATION_DELAY\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_ALLOCATION_CONFIGURATION_DELAY\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ALLOCATION_CONFIGURATION_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEALLOCATION_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addStrategiesToOperatorSet\",\"inputs\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"clearDeallocationQueue\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"numToClear\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createOperatorSets\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.CreateSetParams[]\",\"components\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterFromOperatorSets\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.DeregisterParams\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"encumberedMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAVSRegistrar\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAVSRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatableMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatedSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocation\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.Allocation\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocationDelay\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocations\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.Allocation[]\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudes\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudes\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudesAtBlock\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMemberCount\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMembers\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinimumSlashableStake\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"futureBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"slashableStake\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRegisteredSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStrategiesInOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStrategyAllocations\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.Allocation[]\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"modifyAllocations\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.AllocateParams[]\",\"components\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"newMagnitudes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerForOperatorSets\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.RegisterParams\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeStrategiesFromOperatorSet\",\"inputs\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAVSRegistrar\",\"inputs\":[{\"name\":\"registrar\",\"type\":\"address\",\"internalType\":\"contractIAVSRegistrar\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAllocationDelay\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAllocationDelay\",\"inputs\":[{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashOperator\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.SlashingParams\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"wadToSlash\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AVSMetadataURIUpdated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AVSRegistrarSet\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"registrar\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIAVSRegistrar\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllocationDelaySet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllocationUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"magnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EncumberedMagnitudeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"encumberedMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxMagnitudeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"maxMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetCreated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSlashed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategies\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"contractIStrategy[]\"},{\"name\":\"wadSlashed\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"description\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyAddedToOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyRemovedFromOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyMemberOfSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Empty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientMagnitude\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBlockNumber\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCaller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidStrategy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWadToSlash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ModificationAlreadyPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotMemberOfSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyDelegationManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotSlashable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OutOfBounds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SameMagnitude\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyAlreadyInOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotInOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UninitializedAllocationDelay\",\"inputs\":[]}]",
	Bin: "0x61010060405234801561001157600080fd5b5060405161559d38038061559d83398101604081905261003091610180565b838282856001600160a01b03811661005b576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b039081166080529290921660a05263ffffffff90811660c0521660e052610087610090565b505050506101d9565b600054610100900460ff16156100fc5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161461014d576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b038116811461016457600080fd5b50565b805163ffffffff8116811461017b57600080fd5b919050565b6000806000806080858703121561019657600080fd5b84516101a18161014f565b60208601519094506101b28161014f565b92506101c060408601610167565b91506101ce60608601610167565b905092959194509250565b60805160a05160c05160e05161533e61025f6000396000818161051501526139530152600081816103250152818161112201526120970152600081816106a201528181610b400152818161143e01528181611ada01528181611be901526128bd01526000818161054f01528181610ca301528181611b370152612fb0015261533e6000f3fe608060405234801561001057600080fd5b50600436106102745760003560e01c8063715018a611610151578063a98fb355116100c3578063cd6dc68711610087578063cd6dc68714610677578063ce7b5e4b1461068a578063df5cf7231461069d578063f25f1610146106c4578063f2fde38b146106d7578063fabc1cbc146106ea57600080fd5b8063a98fb355146105fc578063b2447af71461060f578063b92f60a514610622578063b9fbaed114610635578063c221d8ae1461066457600080fd5b8063886f119511610115578063886f11951461054a5780638ce64854146105715780638da5cb5b1461059157806394d7d00c146105a2578063a9333ec8146105b5578063a984eb3a146105c857600080fd5b8063715018a6146104e257806376999342146104ea57806379ae50cd146104fd5780637bc1ef6114610510578063847d634f1461053757600080fd5b80634a10ffe5116101ea5780635ac86ab7116101ae5780635ac86ab7146104495780635c489bb51461046c5780635c975abb1461047f5780636cfb4481146104915780636e3492b5146104bc5780636e875dba146104cf57600080fd5b80634a10ffe5146103e85780634b5046ef14610408578063547afb871461041b57806356c483e61461042e578063595c6a671461044157600080fd5b8063260dc7581161023c578063260dc758146102fd5780632981eb77146103205780632bab2c4a1461035c578063304c10cd1461037c57806340120dab146103a75780634177a87c146103c857600080fd5b80630ea43e431461027957806310e1b9b81461028e578063136439dd146102b757806315ea7917146102ca57806315fe5028146102dd575b600080fd5b61028c610287366004614330565b6106fd565b005b6102a161029c366004614438565b610c58565b6040516102ae9190614482565b60405180910390f35b61028c6102c53660046144b5565b610c8e565b61028c6102d8366004614512565b610d65565b6102f06102eb366004614553565b611294565b6040516102ae91906145d5565b61031061030b3660046145e8565b6113b1565b60405190151581526020016102ae565b6103477f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016102ae565b61036f61036a36600461469d565b6113e9565b6040516102ae9190614758565b61038f61038a366004614553565b6116e5565b6040516001600160a01b0390911681526020016102ae565b6103ba6103b53660046147bd565b611715565b6040516102ae929190614859565b6103db6103d63660046145e8565b611898565b6040516102ae91906148b9565b6103fb6103f63660046148cc565b6118bf565b6040516102ae9190614912565b61028c61041636600461495e565b611969565b6103fb6104293660046149e4565b611a25565b61028c61043c366004614a33565b611acf565b61028c611b22565b610310610457366004614a68565b606654600160ff9092169190911b9081161490565b61028c61047a366004614a8b565b611bd4565b6066545b6040519081526020016102ae565b6104a461049f3660046147bd565b611c86565b6040516001600160401b0390911681526020016102ae565b61028c6104ca366004614abe565b611df7565b6103db6104dd3660046145e8565b6121da565b61028c6121ec565b61028c6104f8366004614af2565b6121fe565b6102f061050b366004614553565b61232c565b6103477f000000000000000000000000000000000000000000000000000000000000000081565b61028c610545366004614512565b61240b565b61038f7f000000000000000000000000000000000000000000000000000000000000000081565b61058461057f366004614b44565b61263d565b6040516102ae9190614b8b565b6033546001600160a01b031661038f565b6103fb6105b0366004614b9e565b612706565b6104a46105c33660046147bd565b6127f6565b6104a46105d63660046147bd565b60a26020908152600092835260408084209091529082529020546001600160401b031681565b61028c61060a366004614bfd565b612826565b61048361061d3660046145e8565b61286d565b61028c610630366004614abe565b61287f565b610648610643366004614553565b612ba2565b60408051921515835263ffffffff9091166020830152016102ae565b6103db610672366004614c6f565b612c3d565b61028c610685366004614c9c565b612c6f565b61028c610698366004614af2565b612d91565b61038f7f000000000000000000000000000000000000000000000000000000000000000081565b61028c6106d2366004614553565b612ebf565b61028c6106e5366004614553565b612f38565b61028c6106f83660046144b5565b612fae565b6066546001906002908116036107265760405163840a48d560e01b815260040160405180910390fd5b816040013560001080156107465750670de0b6b3a7640000604083013511155b61076357604051631353603160e01b815260040160405180910390fd5b60006040518060400160405280336001600160a01b031681526020018460200160208101906107929190614a8b565b63ffffffff169052905060006107b46107ae6020860186614553565b836130bf565b60208084015184516001600160a01b03166000908152609890925260409091209192506107eb919063ffffffff9081169061313616565b61080857604051631fb1705560e21b815260040160405180910390fd5b80610826576040516325131d4f60e01b815260040160405180910390fd5b600061084b609960006108388661314e565b81526020019081526020016000206131ae565b90506000816001600160401b038111156108675761086761437f565b604051908082528060200260200182016040528015610890578160200160208202803683370190505b5090506000826001600160401b038111156108ad576108ad61437f565b6040519080825280602002602001820160405280156108d6578160200160208202803683370190505b50905060005b83811015610bf657600061091382609960006108f78b61314e565b81526020019081526020016000206131b890919063ffffffff16565b905060008061093761092860208d018d614553565b6109318b61314e565b856131c4565b915091508286858151811061094e5761094e614cc8565b6001600160a01b039092166020928302919091019091015280516001600160401b031660000361098057505050610bee565b805160009061099c906001600160401b031660408e0135613335565b83519091506000906109ba906001600160401b03808516911661334c565b9050808787815181106109cf576109cf614cc8565b60200260200101818152505081836000018181516109ed9190614cf4565b6001600160401b0316905250835182908590610a0a908390614cf4565b6001600160401b0316905250602084018051839190610a2a908390614cf4565b6001600160401b031690525060208301516000600f9190910b1215610b04576000610a708e604001358560200151610a6190614d13565b6001600160801b031690613335565b9050806001600160401b031684602001818151610a8d9190614d39565b915090600f0b9081600f0b815250507f1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd8e6000016020810190610ad09190614553565b8d88610ae488600001518960200151613361565b8860400151604051610afa959493929190614d66565b60405180910390a1505b610b25610b1460208f018f614553565b610b1d8d61314e565b878787613376565b610b3e610b3560208f018f614553565b855187906135eb565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663b6f73bdf8e6000016020810190610b809190614553565b6040516001600160e01b031960e084901b1681526001600160a01b039182166004820152908816602482015260448101849052606401600060405180830381600087803b158015610bd057600080fd5b505af1158015610be4573d6000803e3d6000fd5b5050505050505050505b6001016108dc565b507f80969ad29428d6797ee7aad084f9e4a42a82fc506dcd2ca3b6fb431f85ccebe5610c256020890189614553565b868484610c3560608d018d614db7565b604051610c4796959493929190614e26565b60405180910390a150505050505050565b6040805160608101825260008082526020820181905291810182905290610c82856109318661314e565b925050505b9392505050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610cf2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d169190614e89565b610d3357604051631d77d47760e21b815260040160405180910390fd5b6066548181168114610d585760405163c61dca5d60e01b815260040160405180910390fd5b610d6182613670565b5050565b606654600090600190811603610d8e5760405163840a48d560e01b815260040160405180910390fd5b600080610d9a33612ba2565b9150915081610dbc5760405163fa55fc8160e01b815260040160405180910390fd5b60005b8481101561128c57858582818110610dd957610dd9614cc8565b9050602002810190610deb9190614eab565b610df9906060810190614ecb565b9050868683818110610e0d57610e0d614cc8565b9050602002810190610e1f9190614eab565b610e2d906040810190614ecb565b905014610e4d576040516343714afd60e01b815260040160405180910390fd5b36868683818110610e6057610e60614cc8565b9050602002810190610e729190614eab565b90506000610e8e33610e89368590038501856145e8565b6130bf565b9050610ed9610ea36040840160208501614a8b565b63ffffffff1660986000610eba6020870187614553565b6001600160a01b03168152602081019190915260400160002090613136565b610ef657604051631fb1705560e21b815260040160405180910390fd5b60005b888885818110610f0b57610f0b614cc8565b9050602002810190610f1d9190614eab565b610f2b906040810190614ecb565b9050811015611281576000898986818110610f4857610f48614cc8565b9050602002810190610f5a9190614eab565b610f68906040810190614ecb565b83818110610f7857610f78614cc8565b9050602002016020810190610f8d9190614553565b9050610f9c338261ffff6136ad565b600080610fba33610931610fb5368a90038a018a6145e8565b61314e565b915091508060200151600f0b600014610fe657604051630d8fcbe360e41b815260040160405180910390fd5b6000611002610ffa368990038901896145e8565b8584896137b6565b905061106882600001518e8e8b81811061101e5761101e614cc8565b90506020028101906110309190614eab565b61103e906060810190614ecb565b8881811061104e5761104e614cc8565b90506020020160208101906110639190614f14565b613823565b600f0b6020830181905260000361109257604051634606179360e11b815260040160405180910390fd5b60008260200151600f0b12156111e55780156111575761111d6110bd610fb5368a90038a018a6145e8565b33600090815260a3602090815260408083206001600160a01b038a1684529091529020908154600160801b90819004600f0b6000818152600180860160205260409091209390935583546001600160801b03908116939091011602179055565b6111477f000000000000000000000000000000000000000000000000000000000000000043614f3d565b63ffffffff166040830152611253565b61116983602001518360200151613361565b6001600160401b031660208401528c8c8981811061118957611189614cc8565b905060200281019061119b9190614eab565b6111a9906060810190614ecb565b868181106111b9576111b9614cc8565b90506020020160208101906111ce9190614f14565b6001600160401b0316825260006020830152611253565b60008260200151600f0b13156112535761120783602001518360200151613361565b6001600160401b03908116602085018190528451909116101561123d57604051636c9be0bf60e01b815260040160405180910390fd5b6112478943614f3d565b63ffffffff1660408301525b61127133611269610fb5368b90038b018b6145e8565b868686613376565b505060019092019150610ef99050565b505050600101610dbf565b505050505050565b6001600160a01b0381166000908152609d60205260408120606091906112b9906131ae565b90506000816001600160401b038111156112d5576112d561437f565b60405190808252806020026020018201604052801561131a57816020015b60408051808201909152600080825260208201528152602001906001900390816112f35790505b50905060005b828110156113a9576001600160a01b0385166000908152609d602052604090206113849061134e90836131b8565b60408051808201909152600080825260208201525060408051808201909152606082901c815263ffffffff909116602082015290565b82828151811061139657611396614cc8565b6020908102919091010152600101611320565b509392505050565b60208082015182516001600160a01b031660009081526098909252604082206113e39163ffffffff9081169061313616565b92915050565b606083516001600160401b038111156114045761140461437f565b60405190808252806020026020018201604052801561143757816020015b60608152602001906001900390816114225790505b50905060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f0e0e67686866040518363ffffffff1660e01b815260040161148a929190614f59565b600060405180830381865afa1580156114a7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526114cf9190810190614f7e565b905060005b85518110156116db5760008682815181106114f1576114f1614cc8565b6020026020010151905085516001600160401b038111156115145761151461437f565b60405190808252806020026020018201604052801561153d578160200160208202803683370190505b5084838151811061155057611550614cc8565b602002602001018190525060005b86518110156116d157600087828151811061157b5761157b614cc8565b6020908102919091018101516001600160a01b03808616600090815260a18452604080822092841682529190935282209092506115b79061383b565b9050806001600160401b03166000036115d15750506116c9565b60006115de858d85610c58565b90508863ffffffff16816040015163ffffffff1611158015611607575060008160200151600f0b125b1561162a5761161e81600001518260200151613361565b6001600160401b031681525b8051600090611646906001600160401b0390811690851661334c565b905061168d8189898151811061165e5761165e614cc8565b6020026020010151878151811061167757611677614cc8565b602002602001015161388990919063ffffffff16565b89888151811061169f5761169f614cc8565b602002602001015186815181106116b8576116b8614cc8565b602002602001018181525050505050505b60010161155e565b50506001016114d4565b5050949350505050565b6001600160a01b03808216600090815260976020526040812054909116801561170e5780610c87565b5090919050565b6001600160a01b0382166000908152609d60205260408120606091829161173b906131ae565b90506000816001600160401b038111156117575761175761437f565b60405190808252806020026020018201604052801561179c57816020015b60408051808201909152600080825260208201528152602001906001900390816117755790505b5090506000826001600160401b038111156117b9576117b961437f565b60405190808252806020026020018201604052801561180457816020015b60408051606081018252600080825260208083018290529282015282526000199092019101816117d75790505b50905060005b83811015611889576001600160a01b0388166000908152609d602052604081206118389061134e90846131b8565b90508084838151811061184d5761184d614cc8565b602002602001018190525061186389828a610c58565b83838151811061187557611875614cc8565b60209081029190910101525060010161180a565b509093509150505b9250929050565b60606000610c87609960006118ac8661314e565b815260200190815260200160002061389e565b6060600083516001600160401b038111156118dc576118dc61437f565b604051908082528060200260200182016040528015611905578160200160208202803683370190505b50905060005b84518110156113a95761193785828151811061192957611929614cc8565b6020026020010151856127f6565b82828151811061194957611949614cc8565b6001600160401b039092166020928302919091019091015260010161190b565b6066546000906001908116036119925760405163840a48d560e01b815260040160405180910390fd5b8382146119b2576040516343714afd60e01b815260040160405180910390fd5b60005b84811015611a1c57611a14878787848181106119d3576119d3614cc8565b90506020020160208101906119e89190614553565b8686858181106119fa576119fa614cc8565b9050602002016020810190611a0f9190615092565b6136ad565b6001016119b5565b50505050505050565b6060600082516001600160401b03811115611a4257611a4261437f565b604051908082528060200260200182016040528015611a6b578160200160208202803683370190505b50905060005b83518110156113a957611a9d85858381518110611a9057611a90614cc8565b60200260200101516127f6565b828281518110611aaf57611aaf614cc8565b6001600160401b0390921660209283029190910190910152600101611a71565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611b185760405163f739589b60e01b815260040160405180910390fd5b610d6182826138ab565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015611b86573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611baa9190614e89565b611bc757604051631d77d47760e21b815260040160405180910390fd5b611bd2600019613670565b565b6040516336b87bd760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690636d70f7ae90602401602060405180830381865afa158015611c38573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c5c9190614e89565b611c79576040516325ec6c1f60e01b815260040160405180910390fd5b611c8333826138ab565b50565b6001600160a01b03828116600081815260a2602090815260408083209486168084529482528083205493835260a38252808320948352939052918220546001600160401b0390911690600f81810b600160801b909204900b03825b81811015611db3576001600160a01b03808716600090815260a3602090815260408083209389168352929052908120611d1a9083613a4e565b6001600160a01b03888116600090815260a0602090815260408083208584528252808320938b16835292815290829020825160608101845290546001600160401b0381168252600160401b8104600f0b92820192909252600160c01b90910463ffffffff16918101829052919250431015611d96575050611db3565b611da4858260200151613361565b94505050806001019050611ce1565b506001600160a01b03808616600090815260a1602090815260408083209388168352929052208290611de49061383b565b611dee9190614cf4565b95945050505050565b606654600290600490811603611e205760405163840a48d560e01b815260040160405180910390fd5b611e2d6020830183614553565b6001600160a01b0316336001600160a01b03161480611e6c5750611e576040830160208401614553565b6001600160a01b0316336001600160a01b0316145b611e89576040516348f5c3ed60e01b815260040160405180910390fd5b60005b611e996040840184614ecb565b905081101561214c5760006040518060400160405280856020016020810190611ec29190614553565b6001600160a01b03168152602001611edd6040870187614ecb565b85818110611eed57611eed614cc8565b9050602002016020810190611f029190614a8b565b63ffffffff168152509050611f32816020015163ffffffff1660986000876020016020810190610eba9190614553565b611f4f57604051631fb1705560e21b815260040160405180910390fd5b609e6000611f606020870187614553565b6001600160a01b03166001600160a01b031681526020019081526020016000206000611f8b8361314e565b815260208101919091526040016000205460ff16611fbc576040516325131d4f60e01b815260040160405180910390fd5b611ff8611fc88261314e565b609c6000611fd96020890189614553565b6001600160a01b03168152602081019190915260400160002090613abf565b506120326120096020860186614553565b609a60006120168561314e565b8152602001908152602001600020613acb90919063ffffffff16565b506120406020850185614553565b6001600160a01b03167fad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe8260405161207891906150b6565b60405180910390a26040805180820190915260008152602081016120bc7f000000000000000000000000000000000000000000000000000000000000000043614f3d565b63ffffffff169052609e60006120d56020880188614553565b6001600160a01b03166001600160a01b0316815260200190815260200160002060006121008461314e565b815260208082019290925260400160002082518154939092015163ffffffff166101000264ffffffff00199215159290921664ffffffffff199093169290921717905550600101611e8c565b5061216061038a6040840160208501614553565b6001600160a01b0316639d8e0c2361217b6020850185614553565b6121886040860186614ecb565b6040518463ffffffff1660e01b81526004016121a6939291906150ff565b600060405180830381600087803b1580156121c057600080fd5b505af19250505080156121d1575060015b15610d61575050565b60606113e3609a60006118ac8561314e565b6121f4613ae0565b611bd26000613b3a565b6040805180820182523380825263ffffffff8087166020808501829052600093845260989052939091209192612235929161313616565b61225257604051631fb1705560e21b815260040160405180910390fd5b600061225d8261314e565b905060005b8381101561128c576122a885858381811061227f5761227f614cc8565b90506020020160208101906122949190614553565b600084815260996020526040902090613b8c565b6122c55760405163585cfb2f60e01b815260040160405180910390fd5b7f7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b838686848181106122f9576122f9614cc8565b905060200201602081019061230e9190614553565b60405161231c929190615124565b60405180910390a1600101612262565b6001600160a01b0381166000908152609c6020526040812060609190612351906131ae565b90506000816001600160401b0381111561236d5761236d61437f565b6040519080825280602002602001820160405280156123b257816020015b604080518082019091526000808252602082015281526020019060019003908161238b5790505b50905060005b828110156113a9576001600160a01b0385166000908152609c602052604090206123e69061134e90836131b8565b8282815181106123f8576123f8614cc8565b60209081029190910101526001016123b8565b60005b818110156126385760006040518060400160405280336001600160a01b0316815260200185858581811061244457612444614cc8565b9050602002810190612456919061514a565b612464906020810190614a8b565b63ffffffff90811690915260208083015133600090815260989092526040909120929350612497929190811690613ba116565b6124b457604051631fb1705560e21b815260040160405180910390fd5b60408051808201825233815260208381015163ffffffff169082015290517f31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c916124fd916150b6565b60405180910390a160006125108261314e565b905060005b85858581811061252757612527614cc8565b9050602002810190612539919061514a565b612547906020810190614ecb565b905081101561262d5761259586868681811061256557612565614cc8565b9050602002810190612577919061514a565b612585906020810190614ecb565b8381811061227f5761227f614cc8565b507f7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b838787878181106125ca576125ca614cc8565b90506020028101906125dc919061514a565b6125ea906020810190614ecb565b848181106125fa576125fa614cc8565b905060200201602081019061260f9190614553565b60405161261d929190615124565b60405180910390a1600101612515565b50505060010161240e565b505050565b6060600084516001600160401b0381111561265a5761265a61437f565b6040519080825280602002602001820160405280156126a557816020015b60408051606081018252600080825260208083018290529282015282526000199092019101816126785790505b50905060005b85518110156126fd576126d88682815181106126c9576126c9614cc8565b60200260200101518686610c58565b8282815181106126ea576126ea614cc8565b60209081029190910101526001016126ab565b50949350505050565b6060600083516001600160401b038111156127235761272361437f565b60405190808252806020026020018201604052801561274c578160200160208202803683370190505b50905060005b84518110156126fd576001600160a01b038616600090815260a16020526040812086516127c49287929189908690811061278e5761278e614cc8565b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020613bad90919063ffffffff16565b8282815181106127d6576127d6614cc8565b6001600160401b0390921660209283029190910190910152600101612752565b6001600160a01b03808316600090815260a1602090815260408083209385168352929052908120610c879061383b565b336001600160a01b03167fa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c9437138383604051612861929190615160565b60405180910390a25050565b60006113e3609a60006108388561314e565b6066546002906004908116036128a85760405163840a48d560e01b815260040160405180910390fd5b6040516336b87bd760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690636d70f7ae90602401602060405180830381865afa15801561290c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129309190614e89565b61294d5760405163ccea9e6f60e01b815260040160405180910390fd5b60005b61295d6020840184614ecb565b9050811015612b1a5760408051808201909152600090806129816020870187614553565b6001600160a01b0316815260200185806020019061299f9190614ecb565b858181106129af576129af614cc8565b90506020020160208101906129c49190614a8b565b63ffffffff90811690915260208083015183516001600160a01b0316600090815260989092526040909120929350612a0192919081169061313616565b612a1e57604051631fb1705560e21b815260040160405180910390fd5b612a2833826130bf565b15612a4657604051636c6c6e2760e11b815260040160405180910390fd5b612a67612a528261314e565b336000908152609c6020526040902090613ba1565b50612a9533609a6000612a798561314e565b8152602001908152602001600020613b8c90919063ffffffff16565b50336001600160a01b03167f43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e82604051612acf91906150b6565b60405180910390a2336000908152609e60205260408120600191612af28461314e565b81526020810191909152604001600020805460ff191691151591909117905550600101612950565b50612b2b61038a6020840184614553565b6001600160a01b031663adcf73f733612b476020860186614ecb565b612b546040880188614db7565b6040518663ffffffff1660e01b8152600401612b74959493929190615174565b600060405180830381600087803b158015612b8e57600080fd5b505af115801561128c573d6000803e3d6000fd5b6001600160a01b0381166000908152609b602090815260408083208151608081018352905463ffffffff80821680845260ff600160201b8404161515958401869052650100000000008304821694840194909452600160481b909104166060820181905284939192919015801590612c245750826060015163ffffffff164310155b15612c33575050604081015160015b9590945092505050565b6001600160a01b0382166000908152609f6020526040812060609190612c6790826118ac8661314e565b949350505050565b600054610100900460ff1615808015612c8f5750600054600160ff909116105b80612ca95750303b158015612ca9575060005460ff166001145b612d115760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff191660011790558015612d34576000805461ff0019166101001790555b612d3d82613670565b612d4683613b3a565b8015612638576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a1505050565b6040805180820182523380825263ffffffff8087166020808501829052600093845260989052939091209192612dc8929161313616565b612de557604051631fb1705560e21b815260040160405180910390fd5b6000612df08261314e565b905060005b8381101561128c57612e3b858583818110612e1257612e12614cc8565b9050602002016020810190612e279190614553565b600084815260996020526040902090613acb565b612e58576040516331bc342760e11b815260040160405180910390fd5b7f7b4b073d80dcac55a11177d8459ad9f664ceeb91f71f27167bb14f8152a7eeee83868684818110612e8c57612e8c614cc8565b9050602002016020810190612ea19190614553565b604051612eaf929190615124565b60405180910390a1600101612df5565b33600081815260976020526040902080546001600160a01b0319166001600160a01b0384161790557f2ae945c40c44dc0ec263f95609c3fdc6952e0aefa22d6374e44f2c997acedf8590612f12816116e5565b604080516001600160a01b0393841681529290911660208301520160405180910390a150565b612f40613ae0565b6001600160a01b038116612fa55760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401612d08565b611c8381613b3a565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561300c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061303091906151b8565b6001600160a01b0316336001600160a01b0316146130615760405163794821ff60e01b815260040160405180910390fd5b606654801982198116146130885760405163c61dca5d60e01b815260040160405180910390fd5b606682905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001612861565b6001600160a01b0382166000908152609e602052604081208190816130e38561314e565b8152602080820192909252604090810160002081518083019092525460ff8116151580835261010090910463ffffffff1692820192909252915080612c6757506020015163ffffffff1643109392505050565b60008181526001830160205260408120541515610c87565b60008160000151826020015163ffffffff1660405160200161319692919060609290921b6001600160601b031916825260a01b6001600160a01b031916601482015260200190565b6040516020818303038152906040526113e3906151d5565b60006113e3825490565b6000610c878383613bfe565b604080518082018252600080825260208083018290528351606081018552828152808201839052808501839052845180860186526001600160a01b03898116855260a18452868520908816855290925293822092939281906132259061383b565b6001600160401b0390811682526001600160a01b03898116600081815260a260209081526040808320948c168084529482528083205486169682019690965291815260a082528481208b8252825284812092815291815290839020835160608101855290549283168152600160401b8304600f0b91810191909152600160c01b90910463ffffffff169181018290529192504310156132c857909250905061332d565b6132da81600001518260200151613361565b6001600160401b0316815260208101516000600f9190910b12156133195761330a82602001518260200151613361565b6001600160401b031660208301525b600060408201819052602082015290925090505b935093915050565b6000610c878383670de0b6b3a76400006001613c28565b6000610c8783670de0b6b3a764000084613c83565b6000610c87826001600160401b038516614d39565b602082810180516001600160a01b03888116600081815260a286526040808220938a1680835293875290819020805467ffffffffffffffff19166001600160401b0395861617905593518451918252948101919091529216908201527facf9095feb3a370c9cf692421c69ef320d4db5c66e6a7d29c7694eb02364fc559060600160405180910390a16001600160a01b03858116600090815260a060209081526040808320888452825280832093871683529281528282208451815486840151878701516001600160401b039093166001600160c01b031990921691909117600160401b6001600160801b03909216919091021763ffffffff60c01b1916600160c01b63ffffffff9283160217909155835180850185528381528201929092528251808401909352606087901c8352908616908201527f1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd90869083516040516134e493929188914390614d66565b60405180910390a16020810151600f0b1561354f576001600160a01b0385166000908152609f6020908152604080832087845290915290206135269084613b8c565b506001600160a01b0385166000908152609d602052604090206135499085613ba1565b506135e4565b80516001600160401b03166000036135e4576001600160a01b0385166000908152609f60209081526040808320878452909152902061358e9084613acb565b506001600160a01b0385166000908152609f6020908152604080832087845290915290206135bb906131ae565b6000036135e4576001600160a01b0385166000908152609d6020526040902061128c9085613abf565b5050505050565b6001600160a01b03808416600090815260a16020908152604080832093861683529290522061361b904383613d6d565b5050604080516001600160a01b038086168252841660208201526001600160401b038316918101919091527f1c6458079a41077d003c11faf9bf097e693bd67979e4e6500bac7b29db779b5c90606001612d84565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b6001600160a01b03838116600090815260a360209081526040808320938616835292905290812054600f81810b600160801b909204900b035b6000811180156136f957508261ffff1682105b156135e4576001600160a01b03808616600090815260a360209081526040808320938816835292905290812061372e90613d87565b905060008061373e3384896131c4565b91509150806040015163ffffffff1643101561375c575050506135e4565b6137698884898585613376565b6001600160a01b03808916600090815260a360209081526040808320938b1683529290522061379790613ddb565b506137a1856151f9565b94506137ac84615212565b93505050506136e6565b60006137e584609960006137c98961314e565b8152602001908152602001600020613e5a90919063ffffffff16565b6137f157506000612c67565b816137fe57506000612c67565b82516001600160401b031660000361381857506000612c67565b506001949350505050565b6000610c876001600160401b03808516908416615229565b805460009080156138795761386383613855600184615256565b600091825260209091200190565b54600160201b90046001600160401b0316610c87565b670de0b6b3a76400009392505050565b6000610c878383670de0b6b3a7640000613c83565b60606000610c8783613e7c565b6001600160a01b0382166000908152609b60209081526040918290208251608081018452905463ffffffff808216835260ff600160201b830416151593830193909352650100000000008104831693820193909352600160481b9092041660608201819052158015906139285750806060015163ffffffff164310155b1561394257604081015163ffffffff168152600160208201525b63ffffffff821660408201526139787f000000000000000000000000000000000000000000000000000000000000000043614f3d565b63ffffffff90811660608381019182526001600160a01b0386166000818152609b602090815260409182902087518154838a0151858b01519851928a1664ffffffffff1990921691909117600160201b91151591909102176cffffffffffffffff0000000000191665010000000000978916979097026cffffffff000000000000000000191696909617600160481b968816968702179055815192835294871694820194909452928301919091527f4e85751d6331506c6c62335f207eb31f12a61e570f34f5c17640308785c6d4db9101612d84565b600080613a71613a5d84613ed8565b8554613a6c9190600f0b615269565b613f46565b8454909150600160801b9004600f90810b9082900b12613aa457604051632d0483c560e21b815260040160405180910390fd5b600f0b60009081526001939093016020525050604090205490565b6000610c878383613faf565b6000610c87836001600160a01b038416613faf565b6033546001600160a01b03163314611bd25760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401612d08565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000610c87836001600160a01b0384166140a2565b6000610c8783836140a2565b815460009081613bbf858583856140f1565b90508015613bec57613bd685613855600184615256565b54600160201b90046001600160401b0316611dee565b50670de0b6b3a7640000949350505050565b6000826000018281548110613c1557613c15614cc8565b9060005260206000200154905092915050565b600080613c36868686613c83565b90506001836002811115613c4c57613c4c615291565b148015613c69575060008480613c6457613c646152a7565b868809115b15611dee57613c796001826152bd565b9695505050505050565b6000808060001985870985870292508281108382030391505080600003613cbd57838281613cb357613cb36152a7565b0492505050610c87565b808411613d045760405162461bcd60e51b81526020600482015260156024820152744d6174683a206d756c446976206f766572666c6f7760581b6044820152606401612d08565b60008486880960026001871981018816978890046003810283188082028403028082028403028082028403028082028403028082028403029081029092039091026000889003889004909101858311909403939093029303949094049190911702949350505050565b600080613d7b858585614147565b91509150935093915050565b6000613da28254600f81810b600160801b909204900b131590565b15613dc057604051631ed9509560e11b815260040160405180910390fd5b508054600f0b60009081526001909101602052604090205490565b6000613df68254600f81810b600160801b909204900b131590565b15613e1457604051631ed9509560e11b815260040160405180910390fd5b508054600f0b6000818152600180840160205260408220805492905583546fffffffffffffffffffffffffffffffff191692016001600160801b03169190911790915590565b6001600160a01b03811660009081526001830160205260408120541515610c87565b606081600001805480602002602001604051908101604052809291908181526020018280548015613ecc57602002820191906000526020600020905b815481526020019060010190808311613eb8575b50505050509050919050565b60006001600160ff1b03821115613f425760405162461bcd60e51b815260206004820152602860248201527f53616665436173743a2076616c756520646f65736e27742066697420696e2061604482015267371034b73a191a9b60c11b6064820152608401612d08565b5090565b80600f81900b8114613faa5760405162461bcd60e51b815260206004820152602760248201527f53616665436173743a2076616c756520646f65736e27742066697420696e20316044820152663238206269747360c81b6064820152608401612d08565b919050565b60008181526001830160205260408120548015614098576000613fd3600183615256565b8554909150600090613fe790600190615256565b905081811461404c57600086600001828154811061400757614007614cc8565b906000526020600020015490508087600001848154811061402a5761402a614cc8565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061405d5761405d6152d0565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506113e3565b60009150506113e3565b60008181526001830160205260408120546140e9575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556113e3565b5060006113e3565b60005b818310156113a95760006141088484614315565b60008781526020902090915063ffffffff86169082015463ffffffff16111561413357809250614141565b61413e8160016152bd565b93505b506140f4565b8254600090819080156142a857600061416587613855600185615256565b60408051808201909152905463ffffffff808216808452600160201b9092046001600160401b0316602084015291925090871610156141e65760405162461bcd60e51b815260206004820152601960248201527f536e617073686f743a2064656372656173696e67206b657973000000000000006044820152606401612d08565b805163ffffffff808816911603614237578461420788613855600186615256565b80546001600160401b0392909216600160201b026bffffffffffffffff0000000019909216919091179055614298565b6040805180820190915263ffffffff80881682526001600160401b0380881660208085019182528b54600181018d5560008d8152919091209451940180549151909216600160201b026001600160601b031990911693909216929092171790555b60200151925083915061332d9050565b50506040805180820190915263ffffffff80851682526001600160401b0380851660208085019182528854600181018a5560008a81529182209551950180549251909316600160201b026001600160601b031990921694909316939093179290921790915590508161332d565b600061432460028484186152e6565b610c87908484166152bd565b60006020828403121561434257600080fd5b81356001600160401b0381111561435857600080fd5b820160808185031215610c8757600080fd5b6001600160a01b0381168114611c8357600080fd5b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b03811182821017156143bd576143bd61437f565b604052919050565b803563ffffffff81168114613faa57600080fd5b6000604082840312156143eb57600080fd5b604080519081016001600160401b038111828210171561440d5761440d61437f565b604052905080823561441e8161436a565b815261442c602084016143c5565b60208201525092915050565b60008060006080848603121561444d57600080fd5b83356144588161436a565b925061446785602086016143d9565b915060608401356144778161436a565b809150509250925092565b81516001600160401b03168152602080830151600f0b9082015260408083015163ffffffff1690820152606081016113e3565b6000602082840312156144c757600080fd5b5035919050565b60008083601f8401126144e057600080fd5b5081356001600160401b038111156144f757600080fd5b6020830191508360208260051b850101111561189157600080fd5b6000806020838503121561452557600080fd5b82356001600160401b0381111561453b57600080fd5b614547858286016144ce565b90969095509350505050565b60006020828403121561456557600080fd5b8135610c878161436a565b80516001600160a01b0316825260209081015163ffffffff16910152565b600081518084526020840193506020830160005b828110156145cb576145b5868351614570565b60409590950194602091909101906001016145a2565b5093949350505050565b602081526000610c87602083018461458e565b6000604082840312156145fa57600080fd5b610c8783836143d9565b60006001600160401b0382111561461d5761461d61437f565b5060051b60200190565b600082601f83011261463857600080fd5b813561464b61464682614604565b614395565b8082825260208201915060208360051b86010192508583111561466d57600080fd5b602085015b838110156146935780356146858161436a565b835260209283019201614672565b5095945050505050565b60008060008060a085870312156146b357600080fd5b6146bd86866143d9565b935060408501356001600160401b038111156146d857600080fd5b6146e487828801614627565b93505060608501356001600160401b0381111561470057600080fd5b61470c87828801614627565b92505061471b608086016143c5565b905092959194509250565b600081518084526020840193506020830160005b828110156145cb57815186526020958601959091019060010161473a565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b828110156147b157603f1987860301845261479c858351614726565b94506020938401939190910190600101614780565b50929695505050505050565b600080604083850312156147d057600080fd5b82356147db8161436a565b915060208301356147eb8161436a565b809150509250929050565b600081518084526020840193506020830160005b828110156145cb5761484386835180516001600160401b03168252602080820151600f0b9083015260409081015163ffffffff16910152565b606095909501946020919091019060010161480a565b60408152600061486c604083018561458e565b8281036020840152611dee81856147f6565b600081518084526020840193506020830160005b828110156145cb5781516001600160a01b0316865260209586019590910190600101614892565b602081526000610c87602083018461487e565b600080604083850312156148df57600080fd5b82356001600160401b038111156148f557600080fd5b61490185828601614627565b92505060208301356147eb8161436a565b602080825282518282018190526000918401906040840190835b818110156149535783516001600160401b031683526020938401939092019160010161492c565b509095945050505050565b60008060008060006060868803121561497657600080fd5b85356149818161436a565b945060208601356001600160401b0381111561499c57600080fd5b6149a8888289016144ce565b90955093505060408601356001600160401b038111156149c757600080fd5b6149d3888289016144ce565b969995985093965092949392505050565b600080604083850312156149f757600080fd5b8235614a028161436a565b915060208301356001600160401b03811115614a1d57600080fd5b614a2985828601614627565b9150509250929050565b60008060408385031215614a4657600080fd5b8235614a518161436a565b9150614a5f602084016143c5565b90509250929050565b600060208284031215614a7a57600080fd5b813560ff81168114610c8757600080fd5b600060208284031215614a9d57600080fd5b610c87826143c5565b600060608284031215614ab857600080fd5b50919050565b600060208284031215614ad057600080fd5b81356001600160401b03811115614ae657600080fd5b612c6784828501614aa6565b600080600060408486031215614b0757600080fd5b614b10846143c5565b925060208401356001600160401b03811115614b2b57600080fd5b614b37868287016144ce565b9497909650939450505050565b600080600060808486031215614b5957600080fd5b83356001600160401b03811115614b6f57600080fd5b614b7b86828701614627565b93505061446785602086016143d9565b602081526000610c8760208301846147f6565b600080600060608486031215614bb357600080fd5b8335614bbe8161436a565b925060208401356001600160401b03811115614bd957600080fd5b614be586828701614627565b925050614bf4604085016143c5565b90509250925092565b60008060208385031215614c1057600080fd5b82356001600160401b03811115614c2657600080fd5b8301601f81018513614c3757600080fd5b80356001600160401b03811115614c4d57600080fd5b856020828401011115614c5f57600080fd5b6020919091019590945092505050565b60008060608385031215614c8257600080fd5b8235614c8d8161436a565b9150614a5f84602085016143d9565b60008060408385031215614caf57600080fd5b8235614cba8161436a565b946020939093013593505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6001600160401b0382811682821603908111156113e3576113e3614cde565b600081600f0b60016001607f1b03198103614d3057614d30614cde565b60000392915050565b600f81810b9083900b0160016001607f1b03811360016001607f1b0319821217156113e3576113e3614cde565b6001600160a01b038616815260c08101614d836020830187614570565b6001600160a01b039490941660608201526001600160401b0392909216608083015263ffffffff1660a09091015292915050565b6000808335601e19843603018112614dce57600080fd5b8301803591506001600160401b03821115614de857600080fd5b60200191503681900382131561189157600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b0387168152614e3f6020820187614570565b60c060608201526000614e5560c083018761487e565b8281036080840152614e678187614726565b905082810360a0840152614e7c818587614dfd565b9998505050505050505050565b600060208284031215614e9b57600080fd5b81518015158114610c8757600080fd5b60008235607e19833603018112614ec157600080fd5b9190910192915050565b6000808335601e19843603018112614ee257600080fd5b8301803591506001600160401b03821115614efc57600080fd5b6020019150600581901b360382131561189157600080fd5b600060208284031215614f2657600080fd5b81356001600160401b0381168114610c8757600080fd5b63ffffffff81811683821601908111156113e3576113e3614cde565b604081526000614f6c604083018561487e565b8281036020840152611dee818561487e565b600060208284031215614f9057600080fd5b81516001600160401b03811115614fa657600080fd5b8201601f81018413614fb757600080fd5b8051614fc561464682614604565b8082825260208201915060208360051b850101925086831115614fe757600080fd5b602084015b838110156150875780516001600160401b0381111561500a57600080fd5b8501603f8101891361501b57600080fd5b602081015161502c61464682614604565b808282526020820191506020808460051b8601010192508b83111561505057600080fd5b6040840193505b82841015615072578351825260209384019390910190615057565b86525050602093840193919091019050614fec565b509695505050505050565b6000602082840312156150a457600080fd5b813561ffff81168114610c8757600080fd5b604081016113e38284614570565b81835260208301925060008160005b848110156145cb5763ffffffff6150e9836143c5565b16865260209586019591909101906001016150d3565b6001600160a01b0384168152604060208201819052600090611dee90830184866150c4565b606081016151328285614570565b6001600160a01b039290921660409190910152919050565b60008235603e19833603018112614ec157600080fd5b602081526000612c67602083018486614dfd565b6001600160a01b038616815260606020820181905260009061519990830186886150c4565b82810360408401526151ac818587614dfd565b98975050505050505050565b6000602082840312156151ca57600080fd5b8151610c878161436a565b80516020808301519190811015614ab85760001960209190910360031b1b16919050565b60006001820161520b5761520b614cde565b5060010190565b60008161522157615221614cde565b506000190190565b600f82810b9082900b0360016001607f1b0319811260016001607f1b03821317156113e3576113e3614cde565b818103818111156113e3576113e3614cde565b808201828112600083128015821682158216171561528957615289614cde565b505092915050565b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b808201808211156113e3576113e3614cde565b634e487b7160e01b600052603160045260246000fd5b60008261530357634e487b7160e01b600052601260045260246000fd5b50049056fea2646970667358221220b6d280e0e6895995aaac91a384ae2068c66dc78fe2e7b4ad9e628fbe5193645064736f6c634300081b0033",
}

// ContractAllocationManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractAllocationManagerMetaData.ABI instead.
var ContractAllocationManagerABI = ContractAllocationManagerMetaData.ABI

// ContractAllocationManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractAllocationManagerMetaData.Bin instead.
var ContractAllocationManagerBin = ContractAllocationManagerMetaData.Bin

// DeployContractAllocationManager deploys a new Ethereum contract, binding an instance of ContractAllocationManager to it.
func DeployContractAllocationManager(auth *bind.TransactOpts, backend bind.ContractBackend, _delegation common.Address, _pauserRegistry common.Address, _DEALLOCATION_DELAY uint32, _ALLOCATION_CONFIGURATION_DELAY uint32) (common.Address, *types.Transaction, *ContractAllocationManager, error) {
	parsed, err := ContractAllocationManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractAllocationManagerBin), backend, _delegation, _pauserRegistry, _DEALLOCATION_DELAY, _ALLOCATION_CONFIGURATION_DELAY)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractAllocationManager{ContractAllocationManagerCaller: ContractAllocationManagerCaller{contract: contract}, ContractAllocationManagerTransactor: ContractAllocationManagerTransactor{contract: contract}, ContractAllocationManagerFilterer: ContractAllocationManagerFilterer{contract: contract}}, nil
}

// ContractAllocationManagerMethods is an auto generated interface around an Ethereum contract.
type ContractAllocationManagerMethods interface {
	ContractAllocationManagerCalls
	ContractAllocationManagerTransacts
	ContractAllocationManagerFilters
}

// ContractAllocationManagerCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractAllocationManagerCalls interface {
	ALLOCATIONCONFIGURATIONDELAY(opts *bind.CallOpts) (uint32, error)

	DEALLOCATIONDELAY(opts *bind.CallOpts) (uint32, error)

	Delegation(opts *bind.CallOpts) (common.Address, error)

	EncumberedMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error)

	GetAVSRegistrar(opts *bind.CallOpts, avs common.Address) (common.Address, error)

	GetAllocatableMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error)

	GetAllocatedSets(opts *bind.CallOpts, operator common.Address) ([]OperatorSet, error)

	GetAllocatedStrategies(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet) ([]common.Address, error)

	GetAllocation(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet, strategy common.Address) (IAllocationManagerTypesAllocation, error)

	GetAllocationDelay(opts *bind.CallOpts, operator common.Address) (bool, uint32, error)

	GetAllocations(opts *bind.CallOpts, operators []common.Address, operatorSet OperatorSet, strategy common.Address) ([]IAllocationManagerTypesAllocation, error)

	GetMaxMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error)

	GetMaxMagnitudes(opts *bind.CallOpts, operators []common.Address, strategy common.Address) ([]uint64, error)

	GetMaxMagnitudes0(opts *bind.CallOpts, operator common.Address, strategies []common.Address) ([]uint64, error)

	GetMaxMagnitudesAtBlock(opts *bind.CallOpts, operator common.Address, strategies []common.Address, blockNumber uint32) ([]uint64, error)

	GetMemberCount(opts *bind.CallOpts, operatorSet OperatorSet) (*big.Int, error)

	GetMembers(opts *bind.CallOpts, operatorSet OperatorSet) ([]common.Address, error)

	GetMinimumSlashableStake(opts *bind.CallOpts, operatorSet OperatorSet, operators []common.Address, strategies []common.Address, futureBlock uint32) ([][]*big.Int, error)

	GetRegisteredSets(opts *bind.CallOpts, operator common.Address) ([]OperatorSet, error)

	GetStrategiesInOperatorSet(opts *bind.CallOpts, operatorSet OperatorSet) ([]common.Address, error)

	GetStrategyAllocations(opts *bind.CallOpts, operator common.Address, strategy common.Address) ([]OperatorSet, []IAllocationManagerTypesAllocation, error)

	IsOperatorSet(opts *bind.CallOpts, operatorSet OperatorSet) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts, index uint8) (bool, error)

	Paused0(opts *bind.CallOpts) (*big.Int, error)

	PauserRegistry(opts *bind.CallOpts) (common.Address, error)
}

// ContractAllocationManagerTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractAllocationManagerTransacts interface {
	AddStrategiesToOperatorSet(opts *bind.TransactOpts, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error)

	ClearDeallocationQueue(opts *bind.TransactOpts, operator common.Address, strategies []common.Address, numToClear []uint16) (*types.Transaction, error)

	CreateOperatorSets(opts *bind.TransactOpts, params []IAllocationManagerTypesCreateSetParams) (*types.Transaction, error)

	DeregisterFromOperatorSets(opts *bind.TransactOpts, params IAllocationManagerTypesDeregisterParams) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialPausedStatus *big.Int) (*types.Transaction, error)

	ModifyAllocations(opts *bind.TransactOpts, params []IAllocationManagerTypesAllocateParams) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	PauseAll(opts *bind.TransactOpts) (*types.Transaction, error)

	RegisterForOperatorSets(opts *bind.TransactOpts, params IAllocationManagerTypesRegisterParams) (*types.Transaction, error)

	RemoveStrategiesFromOperatorSet(opts *bind.TransactOpts, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetAVSRegistrar(opts *bind.TransactOpts, registrar common.Address) (*types.Transaction, error)

	SetAllocationDelay(opts *bind.TransactOpts, operator common.Address, delay uint32) (*types.Transaction, error)

	SetAllocationDelay0(opts *bind.TransactOpts, delay uint32) (*types.Transaction, error)

	SlashOperator(opts *bind.TransactOpts, params IAllocationManagerTypesSlashingParams) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	UpdateAVSMetadataURI(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error)
}

// ContractAllocationManagerFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractAllocationManagerFilters interface {
	FilterAVSMetadataURIUpdated(opts *bind.FilterOpts, avs []common.Address) (*ContractAllocationManagerAVSMetadataURIUpdatedIterator, error)
	WatchAVSMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerAVSMetadataURIUpdated, avs []common.Address) (event.Subscription, error)
	ParseAVSMetadataURIUpdated(log types.Log) (*ContractAllocationManagerAVSMetadataURIUpdated, error)

	FilterAVSRegistrarSet(opts *bind.FilterOpts) (*ContractAllocationManagerAVSRegistrarSetIterator, error)
	WatchAVSRegistrarSet(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerAVSRegistrarSet) (event.Subscription, error)
	ParseAVSRegistrarSet(log types.Log) (*ContractAllocationManagerAVSRegistrarSet, error)

	FilterAllocationDelaySet(opts *bind.FilterOpts) (*ContractAllocationManagerAllocationDelaySetIterator, error)
	WatchAllocationDelaySet(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerAllocationDelaySet) (event.Subscription, error)
	ParseAllocationDelaySet(log types.Log) (*ContractAllocationManagerAllocationDelaySet, error)

	FilterAllocationUpdated(opts *bind.FilterOpts) (*ContractAllocationManagerAllocationUpdatedIterator, error)
	WatchAllocationUpdated(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerAllocationUpdated) (event.Subscription, error)
	ParseAllocationUpdated(log types.Log) (*ContractAllocationManagerAllocationUpdated, error)

	FilterEncumberedMagnitudeUpdated(opts *bind.FilterOpts) (*ContractAllocationManagerEncumberedMagnitudeUpdatedIterator, error)
	WatchEncumberedMagnitudeUpdated(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerEncumberedMagnitudeUpdated) (event.Subscription, error)
	ParseEncumberedMagnitudeUpdated(log types.Log) (*ContractAllocationManagerEncumberedMagnitudeUpdated, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractAllocationManagerInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractAllocationManagerInitialized, error)

	FilterMaxMagnitudeUpdated(opts *bind.FilterOpts) (*ContractAllocationManagerMaxMagnitudeUpdatedIterator, error)
	WatchMaxMagnitudeUpdated(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerMaxMagnitudeUpdated) (event.Subscription, error)
	ParseMaxMagnitudeUpdated(log types.Log) (*ContractAllocationManagerMaxMagnitudeUpdated, error)

	FilterOperatorAddedToOperatorSet(opts *bind.FilterOpts, operator []common.Address) (*ContractAllocationManagerOperatorAddedToOperatorSetIterator, error)
	WatchOperatorAddedToOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerOperatorAddedToOperatorSet, operator []common.Address) (event.Subscription, error)
	ParseOperatorAddedToOperatorSet(log types.Log) (*ContractAllocationManagerOperatorAddedToOperatorSet, error)

	FilterOperatorRemovedFromOperatorSet(opts *bind.FilterOpts, operator []common.Address) (*ContractAllocationManagerOperatorRemovedFromOperatorSetIterator, error)
	WatchOperatorRemovedFromOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerOperatorRemovedFromOperatorSet, operator []common.Address) (event.Subscription, error)
	ParseOperatorRemovedFromOperatorSet(log types.Log) (*ContractAllocationManagerOperatorRemovedFromOperatorSet, error)

	FilterOperatorSetCreated(opts *bind.FilterOpts) (*ContractAllocationManagerOperatorSetCreatedIterator, error)
	WatchOperatorSetCreated(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerOperatorSetCreated) (event.Subscription, error)
	ParseOperatorSetCreated(log types.Log) (*ContractAllocationManagerOperatorSetCreated, error)

	FilterOperatorSlashed(opts *bind.FilterOpts) (*ContractAllocationManagerOperatorSlashedIterator, error)
	WatchOperatorSlashed(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerOperatorSlashed) (event.Subscription, error)
	ParseOperatorSlashed(log types.Log) (*ContractAllocationManagerOperatorSlashed, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractAllocationManagerOwnershipTransferredIterator, error)
	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)
	ParseOwnershipTransferred(log types.Log) (*ContractAllocationManagerOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractAllocationManagerPausedIterator, error)
	WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerPaused, account []common.Address) (event.Subscription, error)
	ParsePaused(log types.Log) (*ContractAllocationManagerPaused, error)

	FilterStrategyAddedToOperatorSet(opts *bind.FilterOpts) (*ContractAllocationManagerStrategyAddedToOperatorSetIterator, error)
	WatchStrategyAddedToOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerStrategyAddedToOperatorSet) (event.Subscription, error)
	ParseStrategyAddedToOperatorSet(log types.Log) (*ContractAllocationManagerStrategyAddedToOperatorSet, error)

	FilterStrategyRemovedFromOperatorSet(opts *bind.FilterOpts) (*ContractAllocationManagerStrategyRemovedFromOperatorSetIterator, error)
	WatchStrategyRemovedFromOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerStrategyRemovedFromOperatorSet) (event.Subscription, error)
	ParseStrategyRemovedFromOperatorSet(log types.Log) (*ContractAllocationManagerStrategyRemovedFromOperatorSet, error)

	FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractAllocationManagerUnpausedIterator, error)
	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerUnpaused, account []common.Address) (event.Subscription, error)
	ParseUnpaused(log types.Log) (*ContractAllocationManagerUnpaused, error)
}

// ContractAllocationManager is an auto generated Go binding around an Ethereum contract.
type ContractAllocationManager struct {
	ContractAllocationManagerCaller     // Read-only binding to the contract
	ContractAllocationManagerTransactor // Write-only binding to the contract
	ContractAllocationManagerFilterer   // Log filterer for contract events
}

// ContractAllocationManager implements the ContractAllocationManagerMethods interface.
var _ ContractAllocationManagerMethods = (*ContractAllocationManager)(nil)

// ContractAllocationManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractAllocationManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAllocationManagerCaller implements the ContractAllocationManagerCalls interface.
var _ ContractAllocationManagerCalls = (*ContractAllocationManagerCaller)(nil)

// ContractAllocationManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractAllocationManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAllocationManagerTransactor implements the ContractAllocationManagerTransacts interface.
var _ ContractAllocationManagerTransacts = (*ContractAllocationManagerTransactor)(nil)

// ContractAllocationManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractAllocationManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAllocationManagerFilterer implements the ContractAllocationManagerFilters interface.
var _ ContractAllocationManagerFilters = (*ContractAllocationManagerFilterer)(nil)

// ContractAllocationManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractAllocationManagerSession struct {
	Contract     *ContractAllocationManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ContractAllocationManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractAllocationManagerCallerSession struct {
	Contract *ContractAllocationManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// ContractAllocationManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractAllocationManagerTransactorSession struct {
	Contract     *ContractAllocationManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ContractAllocationManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractAllocationManagerRaw struct {
	Contract *ContractAllocationManager // Generic contract binding to access the raw methods on
}

// ContractAllocationManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractAllocationManagerCallerRaw struct {
	Contract *ContractAllocationManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractAllocationManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractAllocationManagerTransactorRaw struct {
	Contract *ContractAllocationManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractAllocationManager creates a new instance of ContractAllocationManager, bound to a specific deployed contract.
func NewContractAllocationManager(address common.Address, backend bind.ContractBackend) (*ContractAllocationManager, error) {
	contract, err := bindContractAllocationManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManager{ContractAllocationManagerCaller: ContractAllocationManagerCaller{contract: contract}, ContractAllocationManagerTransactor: ContractAllocationManagerTransactor{contract: contract}, ContractAllocationManagerFilterer: ContractAllocationManagerFilterer{contract: contract}}, nil
}

// NewContractAllocationManagerCaller creates a new read-only instance of ContractAllocationManager, bound to a specific deployed contract.
func NewContractAllocationManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractAllocationManagerCaller, error) {
	contract, err := bindContractAllocationManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerCaller{contract: contract}, nil
}

// NewContractAllocationManagerTransactor creates a new write-only instance of ContractAllocationManager, bound to a specific deployed contract.
func NewContractAllocationManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractAllocationManagerTransactor, error) {
	contract, err := bindContractAllocationManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerTransactor{contract: contract}, nil
}

// NewContractAllocationManagerFilterer creates a new log filterer instance of ContractAllocationManager, bound to a specific deployed contract.
func NewContractAllocationManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractAllocationManagerFilterer, error) {
	contract, err := bindContractAllocationManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerFilterer{contract: contract}, nil
}

// bindContractAllocationManager binds a generic wrapper to an already deployed contract.
func bindContractAllocationManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractAllocationManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAllocationManager *ContractAllocationManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAllocationManager.Contract.ContractAllocationManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAllocationManager *ContractAllocationManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.ContractAllocationManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAllocationManager *ContractAllocationManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.ContractAllocationManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAllocationManager *ContractAllocationManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAllocationManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAllocationManager *ContractAllocationManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAllocationManager *ContractAllocationManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.contract.Transact(opts, method, params...)
}

// ALLOCATIONCONFIGURATIONDELAY is a free data retrieval call binding the contract method 0x7bc1ef61.
//
// Solidity: function ALLOCATION_CONFIGURATION_DELAY() view returns(uint32)
func (_ContractAllocationManager *ContractAllocationManagerCaller) ALLOCATIONCONFIGURATIONDELAY(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "ALLOCATION_CONFIGURATION_DELAY")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ALLOCATIONCONFIGURATIONDELAY is a free data retrieval call binding the contract method 0x7bc1ef61.
//
// Solidity: function ALLOCATION_CONFIGURATION_DELAY() view returns(uint32)
func (_ContractAllocationManager *ContractAllocationManagerSession) ALLOCATIONCONFIGURATIONDELAY() (uint32, error) {
	return _ContractAllocationManager.Contract.ALLOCATIONCONFIGURATIONDELAY(&_ContractAllocationManager.CallOpts)
}

// ALLOCATIONCONFIGURATIONDELAY is a free data retrieval call binding the contract method 0x7bc1ef61.
//
// Solidity: function ALLOCATION_CONFIGURATION_DELAY() view returns(uint32)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) ALLOCATIONCONFIGURATIONDELAY() (uint32, error) {
	return _ContractAllocationManager.Contract.ALLOCATIONCONFIGURATIONDELAY(&_ContractAllocationManager.CallOpts)
}

// DEALLOCATIONDELAY is a free data retrieval call binding the contract method 0x2981eb77.
//
// Solidity: function DEALLOCATION_DELAY() view returns(uint32)
func (_ContractAllocationManager *ContractAllocationManagerCaller) DEALLOCATIONDELAY(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "DEALLOCATION_DELAY")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// DEALLOCATIONDELAY is a free data retrieval call binding the contract method 0x2981eb77.
//
// Solidity: function DEALLOCATION_DELAY() view returns(uint32)
func (_ContractAllocationManager *ContractAllocationManagerSession) DEALLOCATIONDELAY() (uint32, error) {
	return _ContractAllocationManager.Contract.DEALLOCATIONDELAY(&_ContractAllocationManager.CallOpts)
}

// DEALLOCATIONDELAY is a free data retrieval call binding the contract method 0x2981eb77.
//
// Solidity: function DEALLOCATION_DELAY() view returns(uint32)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) DEALLOCATIONDELAY() (uint32, error) {
	return _ContractAllocationManager.Contract.DEALLOCATIONDELAY(&_ContractAllocationManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractAllocationManager *ContractAllocationManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractAllocationManager *ContractAllocationManagerSession) Delegation() (common.Address, error) {
	return _ContractAllocationManager.Contract.Delegation(&_ContractAllocationManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractAllocationManager.Contract.Delegation(&_ContractAllocationManager.CallOpts)
}

// EncumberedMagnitude is a free data retrieval call binding the contract method 0xa984eb3a.
//
// Solidity: function encumberedMagnitude(address operator, address strategy) view returns(uint64)
func (_ContractAllocationManager *ContractAllocationManagerCaller) EncumberedMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "encumberedMagnitude", operator, strategy)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// EncumberedMagnitude is a free data retrieval call binding the contract method 0xa984eb3a.
//
// Solidity: function encumberedMagnitude(address operator, address strategy) view returns(uint64)
func (_ContractAllocationManager *ContractAllocationManagerSession) EncumberedMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _ContractAllocationManager.Contract.EncumberedMagnitude(&_ContractAllocationManager.CallOpts, operator, strategy)
}

// EncumberedMagnitude is a free data retrieval call binding the contract method 0xa984eb3a.
//
// Solidity: function encumberedMagnitude(address operator, address strategy) view returns(uint64)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) EncumberedMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _ContractAllocationManager.Contract.EncumberedMagnitude(&_ContractAllocationManager.CallOpts, operator, strategy)
}

// GetAVSRegistrar is a free data retrieval call binding the contract method 0x304c10cd.
//
// Solidity: function getAVSRegistrar(address avs) view returns(address)
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetAVSRegistrar(opts *bind.CallOpts, avs common.Address) (common.Address, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getAVSRegistrar", avs)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAVSRegistrar is a free data retrieval call binding the contract method 0x304c10cd.
//
// Solidity: function getAVSRegistrar(address avs) view returns(address)
func (_ContractAllocationManager *ContractAllocationManagerSession) GetAVSRegistrar(avs common.Address) (common.Address, error) {
	return _ContractAllocationManager.Contract.GetAVSRegistrar(&_ContractAllocationManager.CallOpts, avs)
}

// GetAVSRegistrar is a free data retrieval call binding the contract method 0x304c10cd.
//
// Solidity: function getAVSRegistrar(address avs) view returns(address)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetAVSRegistrar(avs common.Address) (common.Address, error) {
	return _ContractAllocationManager.Contract.GetAVSRegistrar(&_ContractAllocationManager.CallOpts, avs)
}

// GetAllocatableMagnitude is a free data retrieval call binding the contract method 0x6cfb4481.
//
// Solidity: function getAllocatableMagnitude(address operator, address strategy) view returns(uint64)
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetAllocatableMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getAllocatableMagnitude", operator, strategy)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetAllocatableMagnitude is a free data retrieval call binding the contract method 0x6cfb4481.
//
// Solidity: function getAllocatableMagnitude(address operator, address strategy) view returns(uint64)
func (_ContractAllocationManager *ContractAllocationManagerSession) GetAllocatableMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _ContractAllocationManager.Contract.GetAllocatableMagnitude(&_ContractAllocationManager.CallOpts, operator, strategy)
}

// GetAllocatableMagnitude is a free data retrieval call binding the contract method 0x6cfb4481.
//
// Solidity: function getAllocatableMagnitude(address operator, address strategy) view returns(uint64)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetAllocatableMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _ContractAllocationManager.Contract.GetAllocatableMagnitude(&_ContractAllocationManager.CallOpts, operator, strategy)
}

// GetAllocatedSets is a free data retrieval call binding the contract method 0x15fe5028.
//
// Solidity: function getAllocatedSets(address operator) view returns((address,uint32)[])
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetAllocatedSets(opts *bind.CallOpts, operator common.Address) ([]OperatorSet, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getAllocatedSets", operator)

	if err != nil {
		return *new([]OperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new([]OperatorSet)).(*[]OperatorSet)

	return out0, err

}

// GetAllocatedSets is a free data retrieval call binding the contract method 0x15fe5028.
//
// Solidity: function getAllocatedSets(address operator) view returns((address,uint32)[])
func (_ContractAllocationManager *ContractAllocationManagerSession) GetAllocatedSets(operator common.Address) ([]OperatorSet, error) {
	return _ContractAllocationManager.Contract.GetAllocatedSets(&_ContractAllocationManager.CallOpts, operator)
}

// GetAllocatedSets is a free data retrieval call binding the contract method 0x15fe5028.
//
// Solidity: function getAllocatedSets(address operator) view returns((address,uint32)[])
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetAllocatedSets(operator common.Address) ([]OperatorSet, error) {
	return _ContractAllocationManager.Contract.GetAllocatedSets(&_ContractAllocationManager.CallOpts, operator)
}

// GetAllocatedStrategies is a free data retrieval call binding the contract method 0xc221d8ae.
//
// Solidity: function getAllocatedStrategies(address operator, (address,uint32) operatorSet) view returns(address[])
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetAllocatedStrategies(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet) ([]common.Address, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getAllocatedStrategies", operator, operatorSet)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllocatedStrategies is a free data retrieval call binding the contract method 0xc221d8ae.
//
// Solidity: function getAllocatedStrategies(address operator, (address,uint32) operatorSet) view returns(address[])
func (_ContractAllocationManager *ContractAllocationManagerSession) GetAllocatedStrategies(operator common.Address, operatorSet OperatorSet) ([]common.Address, error) {
	return _ContractAllocationManager.Contract.GetAllocatedStrategies(&_ContractAllocationManager.CallOpts, operator, operatorSet)
}

// GetAllocatedStrategies is a free data retrieval call binding the contract method 0xc221d8ae.
//
// Solidity: function getAllocatedStrategies(address operator, (address,uint32) operatorSet) view returns(address[])
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetAllocatedStrategies(operator common.Address, operatorSet OperatorSet) ([]common.Address, error) {
	return _ContractAllocationManager.Contract.GetAllocatedStrategies(&_ContractAllocationManager.CallOpts, operator, operatorSet)
}

// GetAllocation is a free data retrieval call binding the contract method 0x10e1b9b8.
//
// Solidity: function getAllocation(address operator, (address,uint32) operatorSet, address strategy) view returns((uint64,int128,uint32))
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetAllocation(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet, strategy common.Address) (IAllocationManagerTypesAllocation, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getAllocation", operator, operatorSet, strategy)

	if err != nil {
		return *new(IAllocationManagerTypesAllocation), err
	}

	out0 := *abi.ConvertType(out[0], new(IAllocationManagerTypesAllocation)).(*IAllocationManagerTypesAllocation)

	return out0, err

}

// GetAllocation is a free data retrieval call binding the contract method 0x10e1b9b8.
//
// Solidity: function getAllocation(address operator, (address,uint32) operatorSet, address strategy) view returns((uint64,int128,uint32))
func (_ContractAllocationManager *ContractAllocationManagerSession) GetAllocation(operator common.Address, operatorSet OperatorSet, strategy common.Address) (IAllocationManagerTypesAllocation, error) {
	return _ContractAllocationManager.Contract.GetAllocation(&_ContractAllocationManager.CallOpts, operator, operatorSet, strategy)
}

// GetAllocation is a free data retrieval call binding the contract method 0x10e1b9b8.
//
// Solidity: function getAllocation(address operator, (address,uint32) operatorSet, address strategy) view returns((uint64,int128,uint32))
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetAllocation(operator common.Address, operatorSet OperatorSet, strategy common.Address) (IAllocationManagerTypesAllocation, error) {
	return _ContractAllocationManager.Contract.GetAllocation(&_ContractAllocationManager.CallOpts, operator, operatorSet, strategy)
}

// GetAllocationDelay is a free data retrieval call binding the contract method 0xb9fbaed1.
//
// Solidity: function getAllocationDelay(address operator) view returns(bool, uint32)
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetAllocationDelay(opts *bind.CallOpts, operator common.Address) (bool, uint32, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getAllocationDelay", operator)

	if err != nil {
		return *new(bool), *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return out0, out1, err

}

// GetAllocationDelay is a free data retrieval call binding the contract method 0xb9fbaed1.
//
// Solidity: function getAllocationDelay(address operator) view returns(bool, uint32)
func (_ContractAllocationManager *ContractAllocationManagerSession) GetAllocationDelay(operator common.Address) (bool, uint32, error) {
	return _ContractAllocationManager.Contract.GetAllocationDelay(&_ContractAllocationManager.CallOpts, operator)
}

// GetAllocationDelay is a free data retrieval call binding the contract method 0xb9fbaed1.
//
// Solidity: function getAllocationDelay(address operator) view returns(bool, uint32)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetAllocationDelay(operator common.Address) (bool, uint32, error) {
	return _ContractAllocationManager.Contract.GetAllocationDelay(&_ContractAllocationManager.CallOpts, operator)
}

// GetAllocations is a free data retrieval call binding the contract method 0x8ce64854.
//
// Solidity: function getAllocations(address[] operators, (address,uint32) operatorSet, address strategy) view returns((uint64,int128,uint32)[])
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetAllocations(opts *bind.CallOpts, operators []common.Address, operatorSet OperatorSet, strategy common.Address) ([]IAllocationManagerTypesAllocation, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getAllocations", operators, operatorSet, strategy)

	if err != nil {
		return *new([]IAllocationManagerTypesAllocation), err
	}

	out0 := *abi.ConvertType(out[0], new([]IAllocationManagerTypesAllocation)).(*[]IAllocationManagerTypesAllocation)

	return out0, err

}

// GetAllocations is a free data retrieval call binding the contract method 0x8ce64854.
//
// Solidity: function getAllocations(address[] operators, (address,uint32) operatorSet, address strategy) view returns((uint64,int128,uint32)[])
func (_ContractAllocationManager *ContractAllocationManagerSession) GetAllocations(operators []common.Address, operatorSet OperatorSet, strategy common.Address) ([]IAllocationManagerTypesAllocation, error) {
	return _ContractAllocationManager.Contract.GetAllocations(&_ContractAllocationManager.CallOpts, operators, operatorSet, strategy)
}

// GetAllocations is a free data retrieval call binding the contract method 0x8ce64854.
//
// Solidity: function getAllocations(address[] operators, (address,uint32) operatorSet, address strategy) view returns((uint64,int128,uint32)[])
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetAllocations(operators []common.Address, operatorSet OperatorSet, strategy common.Address) ([]IAllocationManagerTypesAllocation, error) {
	return _ContractAllocationManager.Contract.GetAllocations(&_ContractAllocationManager.CallOpts, operators, operatorSet, strategy)
}

// GetMaxMagnitude is a free data retrieval call binding the contract method 0xa9333ec8.
//
// Solidity: function getMaxMagnitude(address operator, address strategy) view returns(uint64)
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetMaxMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getMaxMagnitude", operator, strategy)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetMaxMagnitude is a free data retrieval call binding the contract method 0xa9333ec8.
//
// Solidity: function getMaxMagnitude(address operator, address strategy) view returns(uint64)
func (_ContractAllocationManager *ContractAllocationManagerSession) GetMaxMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _ContractAllocationManager.Contract.GetMaxMagnitude(&_ContractAllocationManager.CallOpts, operator, strategy)
}

// GetMaxMagnitude is a free data retrieval call binding the contract method 0xa9333ec8.
//
// Solidity: function getMaxMagnitude(address operator, address strategy) view returns(uint64)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetMaxMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _ContractAllocationManager.Contract.GetMaxMagnitude(&_ContractAllocationManager.CallOpts, operator, strategy)
}

// GetMaxMagnitudes is a free data retrieval call binding the contract method 0x4a10ffe5.
//
// Solidity: function getMaxMagnitudes(address[] operators, address strategy) view returns(uint64[])
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetMaxMagnitudes(opts *bind.CallOpts, operators []common.Address, strategy common.Address) ([]uint64, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getMaxMagnitudes", operators, strategy)

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// GetMaxMagnitudes is a free data retrieval call binding the contract method 0x4a10ffe5.
//
// Solidity: function getMaxMagnitudes(address[] operators, address strategy) view returns(uint64[])
func (_ContractAllocationManager *ContractAllocationManagerSession) GetMaxMagnitudes(operators []common.Address, strategy common.Address) ([]uint64, error) {
	return _ContractAllocationManager.Contract.GetMaxMagnitudes(&_ContractAllocationManager.CallOpts, operators, strategy)
}

// GetMaxMagnitudes is a free data retrieval call binding the contract method 0x4a10ffe5.
//
// Solidity: function getMaxMagnitudes(address[] operators, address strategy) view returns(uint64[])
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetMaxMagnitudes(operators []common.Address, strategy common.Address) ([]uint64, error) {
	return _ContractAllocationManager.Contract.GetMaxMagnitudes(&_ContractAllocationManager.CallOpts, operators, strategy)
}

// GetMaxMagnitudes0 is a free data retrieval call binding the contract method 0x547afb87.
//
// Solidity: function getMaxMagnitudes(address operator, address[] strategies) view returns(uint64[])
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetMaxMagnitudes0(opts *bind.CallOpts, operator common.Address, strategies []common.Address) ([]uint64, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getMaxMagnitudes0", operator, strategies)

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// GetMaxMagnitudes0 is a free data retrieval call binding the contract method 0x547afb87.
//
// Solidity: function getMaxMagnitudes(address operator, address[] strategies) view returns(uint64[])
func (_ContractAllocationManager *ContractAllocationManagerSession) GetMaxMagnitudes0(operator common.Address, strategies []common.Address) ([]uint64, error) {
	return _ContractAllocationManager.Contract.GetMaxMagnitudes0(&_ContractAllocationManager.CallOpts, operator, strategies)
}

// GetMaxMagnitudes0 is a free data retrieval call binding the contract method 0x547afb87.
//
// Solidity: function getMaxMagnitudes(address operator, address[] strategies) view returns(uint64[])
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetMaxMagnitudes0(operator common.Address, strategies []common.Address) ([]uint64, error) {
	return _ContractAllocationManager.Contract.GetMaxMagnitudes0(&_ContractAllocationManager.CallOpts, operator, strategies)
}

// GetMaxMagnitudesAtBlock is a free data retrieval call binding the contract method 0x94d7d00c.
//
// Solidity: function getMaxMagnitudesAtBlock(address operator, address[] strategies, uint32 blockNumber) view returns(uint64[])
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetMaxMagnitudesAtBlock(opts *bind.CallOpts, operator common.Address, strategies []common.Address, blockNumber uint32) ([]uint64, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getMaxMagnitudesAtBlock", operator, strategies, blockNumber)

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// GetMaxMagnitudesAtBlock is a free data retrieval call binding the contract method 0x94d7d00c.
//
// Solidity: function getMaxMagnitudesAtBlock(address operator, address[] strategies, uint32 blockNumber) view returns(uint64[])
func (_ContractAllocationManager *ContractAllocationManagerSession) GetMaxMagnitudesAtBlock(operator common.Address, strategies []common.Address, blockNumber uint32) ([]uint64, error) {
	return _ContractAllocationManager.Contract.GetMaxMagnitudesAtBlock(&_ContractAllocationManager.CallOpts, operator, strategies, blockNumber)
}

// GetMaxMagnitudesAtBlock is a free data retrieval call binding the contract method 0x94d7d00c.
//
// Solidity: function getMaxMagnitudesAtBlock(address operator, address[] strategies, uint32 blockNumber) view returns(uint64[])
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetMaxMagnitudesAtBlock(operator common.Address, strategies []common.Address, blockNumber uint32) ([]uint64, error) {
	return _ContractAllocationManager.Contract.GetMaxMagnitudesAtBlock(&_ContractAllocationManager.CallOpts, operator, strategies, blockNumber)
}

// GetMemberCount is a free data retrieval call binding the contract method 0xb2447af7.
//
// Solidity: function getMemberCount((address,uint32) operatorSet) view returns(uint256)
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetMemberCount(opts *bind.CallOpts, operatorSet OperatorSet) (*big.Int, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getMemberCount", operatorSet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMemberCount is a free data retrieval call binding the contract method 0xb2447af7.
//
// Solidity: function getMemberCount((address,uint32) operatorSet) view returns(uint256)
func (_ContractAllocationManager *ContractAllocationManagerSession) GetMemberCount(operatorSet OperatorSet) (*big.Int, error) {
	return _ContractAllocationManager.Contract.GetMemberCount(&_ContractAllocationManager.CallOpts, operatorSet)
}

// GetMemberCount is a free data retrieval call binding the contract method 0xb2447af7.
//
// Solidity: function getMemberCount((address,uint32) operatorSet) view returns(uint256)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetMemberCount(operatorSet OperatorSet) (*big.Int, error) {
	return _ContractAllocationManager.Contract.GetMemberCount(&_ContractAllocationManager.CallOpts, operatorSet)
}

// GetMembers is a free data retrieval call binding the contract method 0x6e875dba.
//
// Solidity: function getMembers((address,uint32) operatorSet) view returns(address[])
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetMembers(opts *bind.CallOpts, operatorSet OperatorSet) ([]common.Address, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getMembers", operatorSet)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMembers is a free data retrieval call binding the contract method 0x6e875dba.
//
// Solidity: function getMembers((address,uint32) operatorSet) view returns(address[])
func (_ContractAllocationManager *ContractAllocationManagerSession) GetMembers(operatorSet OperatorSet) ([]common.Address, error) {
	return _ContractAllocationManager.Contract.GetMembers(&_ContractAllocationManager.CallOpts, operatorSet)
}

// GetMembers is a free data retrieval call binding the contract method 0x6e875dba.
//
// Solidity: function getMembers((address,uint32) operatorSet) view returns(address[])
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetMembers(operatorSet OperatorSet) ([]common.Address, error) {
	return _ContractAllocationManager.Contract.GetMembers(&_ContractAllocationManager.CallOpts, operatorSet)
}

// GetMinimumSlashableStake is a free data retrieval call binding the contract method 0x2bab2c4a.
//
// Solidity: function getMinimumSlashableStake((address,uint32) operatorSet, address[] operators, address[] strategies, uint32 futureBlock) view returns(uint256[][] slashableStake)
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetMinimumSlashableStake(opts *bind.CallOpts, operatorSet OperatorSet, operators []common.Address, strategies []common.Address, futureBlock uint32) ([][]*big.Int, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getMinimumSlashableStake", operatorSet, operators, strategies, futureBlock)

	if err != nil {
		return *new([][]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][]*big.Int)).(*[][]*big.Int)

	return out0, err

}

// GetMinimumSlashableStake is a free data retrieval call binding the contract method 0x2bab2c4a.
//
// Solidity: function getMinimumSlashableStake((address,uint32) operatorSet, address[] operators, address[] strategies, uint32 futureBlock) view returns(uint256[][] slashableStake)
func (_ContractAllocationManager *ContractAllocationManagerSession) GetMinimumSlashableStake(operatorSet OperatorSet, operators []common.Address, strategies []common.Address, futureBlock uint32) ([][]*big.Int, error) {
	return _ContractAllocationManager.Contract.GetMinimumSlashableStake(&_ContractAllocationManager.CallOpts, operatorSet, operators, strategies, futureBlock)
}

// GetMinimumSlashableStake is a free data retrieval call binding the contract method 0x2bab2c4a.
//
// Solidity: function getMinimumSlashableStake((address,uint32) operatorSet, address[] operators, address[] strategies, uint32 futureBlock) view returns(uint256[][] slashableStake)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetMinimumSlashableStake(operatorSet OperatorSet, operators []common.Address, strategies []common.Address, futureBlock uint32) ([][]*big.Int, error) {
	return _ContractAllocationManager.Contract.GetMinimumSlashableStake(&_ContractAllocationManager.CallOpts, operatorSet, operators, strategies, futureBlock)
}

// GetRegisteredSets is a free data retrieval call binding the contract method 0x79ae50cd.
//
// Solidity: function getRegisteredSets(address operator) view returns((address,uint32)[])
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetRegisteredSets(opts *bind.CallOpts, operator common.Address) ([]OperatorSet, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getRegisteredSets", operator)

	if err != nil {
		return *new([]OperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new([]OperatorSet)).(*[]OperatorSet)

	return out0, err

}

// GetRegisteredSets is a free data retrieval call binding the contract method 0x79ae50cd.
//
// Solidity: function getRegisteredSets(address operator) view returns((address,uint32)[])
func (_ContractAllocationManager *ContractAllocationManagerSession) GetRegisteredSets(operator common.Address) ([]OperatorSet, error) {
	return _ContractAllocationManager.Contract.GetRegisteredSets(&_ContractAllocationManager.CallOpts, operator)
}

// GetRegisteredSets is a free data retrieval call binding the contract method 0x79ae50cd.
//
// Solidity: function getRegisteredSets(address operator) view returns((address,uint32)[])
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetRegisteredSets(operator common.Address) ([]OperatorSet, error) {
	return _ContractAllocationManager.Contract.GetRegisteredSets(&_ContractAllocationManager.CallOpts, operator)
}

// GetStrategiesInOperatorSet is a free data retrieval call binding the contract method 0x4177a87c.
//
// Solidity: function getStrategiesInOperatorSet((address,uint32) operatorSet) view returns(address[])
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetStrategiesInOperatorSet(opts *bind.CallOpts, operatorSet OperatorSet) ([]common.Address, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getStrategiesInOperatorSet", operatorSet)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetStrategiesInOperatorSet is a free data retrieval call binding the contract method 0x4177a87c.
//
// Solidity: function getStrategiesInOperatorSet((address,uint32) operatorSet) view returns(address[])
func (_ContractAllocationManager *ContractAllocationManagerSession) GetStrategiesInOperatorSet(operatorSet OperatorSet) ([]common.Address, error) {
	return _ContractAllocationManager.Contract.GetStrategiesInOperatorSet(&_ContractAllocationManager.CallOpts, operatorSet)
}

// GetStrategiesInOperatorSet is a free data retrieval call binding the contract method 0x4177a87c.
//
// Solidity: function getStrategiesInOperatorSet((address,uint32) operatorSet) view returns(address[])
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetStrategiesInOperatorSet(operatorSet OperatorSet) ([]common.Address, error) {
	return _ContractAllocationManager.Contract.GetStrategiesInOperatorSet(&_ContractAllocationManager.CallOpts, operatorSet)
}

// GetStrategyAllocations is a free data retrieval call binding the contract method 0x40120dab.
//
// Solidity: function getStrategyAllocations(address operator, address strategy) view returns((address,uint32)[], (uint64,int128,uint32)[])
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetStrategyAllocations(opts *bind.CallOpts, operator common.Address, strategy common.Address) ([]OperatorSet, []IAllocationManagerTypesAllocation, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getStrategyAllocations", operator, strategy)

	if err != nil {
		return *new([]OperatorSet), *new([]IAllocationManagerTypesAllocation), err
	}

	out0 := *abi.ConvertType(out[0], new([]OperatorSet)).(*[]OperatorSet)
	out1 := *abi.ConvertType(out[1], new([]IAllocationManagerTypesAllocation)).(*[]IAllocationManagerTypesAllocation)

	return out0, out1, err

}

// GetStrategyAllocations is a free data retrieval call binding the contract method 0x40120dab.
//
// Solidity: function getStrategyAllocations(address operator, address strategy) view returns((address,uint32)[], (uint64,int128,uint32)[])
func (_ContractAllocationManager *ContractAllocationManagerSession) GetStrategyAllocations(operator common.Address, strategy common.Address) ([]OperatorSet, []IAllocationManagerTypesAllocation, error) {
	return _ContractAllocationManager.Contract.GetStrategyAllocations(&_ContractAllocationManager.CallOpts, operator, strategy)
}

// GetStrategyAllocations is a free data retrieval call binding the contract method 0x40120dab.
//
// Solidity: function getStrategyAllocations(address operator, address strategy) view returns((address,uint32)[], (uint64,int128,uint32)[])
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetStrategyAllocations(operator common.Address, strategy common.Address) ([]OperatorSet, []IAllocationManagerTypesAllocation, error) {
	return _ContractAllocationManager.Contract.GetStrategyAllocations(&_ContractAllocationManager.CallOpts, operator, strategy)
}

// IsOperatorSet is a free data retrieval call binding the contract method 0x260dc758.
//
// Solidity: function isOperatorSet((address,uint32) operatorSet) view returns(bool)
func (_ContractAllocationManager *ContractAllocationManagerCaller) IsOperatorSet(opts *bind.CallOpts, operatorSet OperatorSet) (bool, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "isOperatorSet", operatorSet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorSet is a free data retrieval call binding the contract method 0x260dc758.
//
// Solidity: function isOperatorSet((address,uint32) operatorSet) view returns(bool)
func (_ContractAllocationManager *ContractAllocationManagerSession) IsOperatorSet(operatorSet OperatorSet) (bool, error) {
	return _ContractAllocationManager.Contract.IsOperatorSet(&_ContractAllocationManager.CallOpts, operatorSet)
}

// IsOperatorSet is a free data retrieval call binding the contract method 0x260dc758.
//
// Solidity: function isOperatorSet((address,uint32) operatorSet) view returns(bool)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) IsOperatorSet(operatorSet OperatorSet) (bool, error) {
	return _ContractAllocationManager.Contract.IsOperatorSet(&_ContractAllocationManager.CallOpts, operatorSet)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAllocationManager *ContractAllocationManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAllocationManager *ContractAllocationManagerSession) Owner() (common.Address, error) {
	return _ContractAllocationManager.Contract.Owner(&_ContractAllocationManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) Owner() (common.Address, error) {
	return _ContractAllocationManager.Contract.Owner(&_ContractAllocationManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractAllocationManager *ContractAllocationManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractAllocationManager *ContractAllocationManagerSession) Paused(index uint8) (bool, error) {
	return _ContractAllocationManager.Contract.Paused(&_ContractAllocationManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractAllocationManager.Contract.Paused(&_ContractAllocationManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractAllocationManager *ContractAllocationManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractAllocationManager *ContractAllocationManagerSession) Paused0() (*big.Int, error) {
	return _ContractAllocationManager.Contract.Paused0(&_ContractAllocationManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractAllocationManager.Contract.Paused0(&_ContractAllocationManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractAllocationManager *ContractAllocationManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractAllocationManager *ContractAllocationManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractAllocationManager.Contract.PauserRegistry(&_ContractAllocationManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractAllocationManager.Contract.PauserRegistry(&_ContractAllocationManager.CallOpts)
}

// AddStrategiesToOperatorSet is a paid mutator transaction binding the contract method 0x76999342.
//
// Solidity: function addStrategiesToOperatorSet(uint32 operatorSetId, address[] strategies) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) AddStrategiesToOperatorSet(opts *bind.TransactOpts, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "addStrategiesToOperatorSet", operatorSetId, strategies)
}

// AddStrategiesToOperatorSet is a paid mutator transaction binding the contract method 0x76999342.
//
// Solidity: function addStrategiesToOperatorSet(uint32 operatorSetId, address[] strategies) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) AddStrategiesToOperatorSet(operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.AddStrategiesToOperatorSet(&_ContractAllocationManager.TransactOpts, operatorSetId, strategies)
}

// AddStrategiesToOperatorSet is a paid mutator transaction binding the contract method 0x76999342.
//
// Solidity: function addStrategiesToOperatorSet(uint32 operatorSetId, address[] strategies) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) AddStrategiesToOperatorSet(operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.AddStrategiesToOperatorSet(&_ContractAllocationManager.TransactOpts, operatorSetId, strategies)
}

// ClearDeallocationQueue is a paid mutator transaction binding the contract method 0x4b5046ef.
//
// Solidity: function clearDeallocationQueue(address operator, address[] strategies, uint16[] numToClear) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) ClearDeallocationQueue(opts *bind.TransactOpts, operator common.Address, strategies []common.Address, numToClear []uint16) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "clearDeallocationQueue", operator, strategies, numToClear)
}

// ClearDeallocationQueue is a paid mutator transaction binding the contract method 0x4b5046ef.
//
// Solidity: function clearDeallocationQueue(address operator, address[] strategies, uint16[] numToClear) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) ClearDeallocationQueue(operator common.Address, strategies []common.Address, numToClear []uint16) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.ClearDeallocationQueue(&_ContractAllocationManager.TransactOpts, operator, strategies, numToClear)
}

// ClearDeallocationQueue is a paid mutator transaction binding the contract method 0x4b5046ef.
//
// Solidity: function clearDeallocationQueue(address operator, address[] strategies, uint16[] numToClear) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) ClearDeallocationQueue(operator common.Address, strategies []common.Address, numToClear []uint16) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.ClearDeallocationQueue(&_ContractAllocationManager.TransactOpts, operator, strategies, numToClear)
}

// CreateOperatorSets is a paid mutator transaction binding the contract method 0x847d634f.
//
// Solidity: function createOperatorSets((uint32,address[])[] params) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) CreateOperatorSets(opts *bind.TransactOpts, params []IAllocationManagerTypesCreateSetParams) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "createOperatorSets", params)
}

// CreateOperatorSets is a paid mutator transaction binding the contract method 0x847d634f.
//
// Solidity: function createOperatorSets((uint32,address[])[] params) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) CreateOperatorSets(params []IAllocationManagerTypesCreateSetParams) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.CreateOperatorSets(&_ContractAllocationManager.TransactOpts, params)
}

// CreateOperatorSets is a paid mutator transaction binding the contract method 0x847d634f.
//
// Solidity: function createOperatorSets((uint32,address[])[] params) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) CreateOperatorSets(params []IAllocationManagerTypesCreateSetParams) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.CreateOperatorSets(&_ContractAllocationManager.TransactOpts, params)
}

// DeregisterFromOperatorSets is a paid mutator transaction binding the contract method 0x6e3492b5.
//
// Solidity: function deregisterFromOperatorSets((address,address,uint32[]) params) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) DeregisterFromOperatorSets(opts *bind.TransactOpts, params IAllocationManagerTypesDeregisterParams) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "deregisterFromOperatorSets", params)
}

// DeregisterFromOperatorSets is a paid mutator transaction binding the contract method 0x6e3492b5.
//
// Solidity: function deregisterFromOperatorSets((address,address,uint32[]) params) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) DeregisterFromOperatorSets(params IAllocationManagerTypesDeregisterParams) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.DeregisterFromOperatorSets(&_ContractAllocationManager.TransactOpts, params)
}

// DeregisterFromOperatorSets is a paid mutator transaction binding the contract method 0x6e3492b5.
//
// Solidity: function deregisterFromOperatorSets((address,address,uint32[]) params) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) DeregisterFromOperatorSets(params IAllocationManagerTypesDeregisterParams) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.DeregisterFromOperatorSets(&_ContractAllocationManager.TransactOpts, params)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "initialize", initialOwner, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) Initialize(initialOwner common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.Initialize(&_ContractAllocationManager.TransactOpts, initialOwner, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) Initialize(initialOwner common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.Initialize(&_ContractAllocationManager.TransactOpts, initialOwner, initialPausedStatus)
}

// ModifyAllocations is a paid mutator transaction binding the contract method 0x15ea7917.
//
// Solidity: function modifyAllocations(((address,uint32),address[],uint64[])[] params) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) ModifyAllocations(opts *bind.TransactOpts, params []IAllocationManagerTypesAllocateParams) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "modifyAllocations", params)
}

// ModifyAllocations is a paid mutator transaction binding the contract method 0x15ea7917.
//
// Solidity: function modifyAllocations(((address,uint32),address[],uint64[])[] params) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) ModifyAllocations(params []IAllocationManagerTypesAllocateParams) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.ModifyAllocations(&_ContractAllocationManager.TransactOpts, params)
}

// ModifyAllocations is a paid mutator transaction binding the contract method 0x15ea7917.
//
// Solidity: function modifyAllocations(((address,uint32),address[],uint64[])[] params) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) ModifyAllocations(params []IAllocationManagerTypesAllocateParams) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.ModifyAllocations(&_ContractAllocationManager.TransactOpts, params)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.Pause(&_ContractAllocationManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.Pause(&_ContractAllocationManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.PauseAll(&_ContractAllocationManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.PauseAll(&_ContractAllocationManager.TransactOpts)
}

// RegisterForOperatorSets is a paid mutator transaction binding the contract method 0xb92f60a5.
//
// Solidity: function registerForOperatorSets((address,uint32[],bytes) params) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) RegisterForOperatorSets(opts *bind.TransactOpts, params IAllocationManagerTypesRegisterParams) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "registerForOperatorSets", params)
}

// RegisterForOperatorSets is a paid mutator transaction binding the contract method 0xb92f60a5.
//
// Solidity: function registerForOperatorSets((address,uint32[],bytes) params) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) RegisterForOperatorSets(params IAllocationManagerTypesRegisterParams) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.RegisterForOperatorSets(&_ContractAllocationManager.TransactOpts, params)
}

// RegisterForOperatorSets is a paid mutator transaction binding the contract method 0xb92f60a5.
//
// Solidity: function registerForOperatorSets((address,uint32[],bytes) params) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) RegisterForOperatorSets(params IAllocationManagerTypesRegisterParams) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.RegisterForOperatorSets(&_ContractAllocationManager.TransactOpts, params)
}

// RemoveStrategiesFromOperatorSet is a paid mutator transaction binding the contract method 0xce7b5e4b.
//
// Solidity: function removeStrategiesFromOperatorSet(uint32 operatorSetId, address[] strategies) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) RemoveStrategiesFromOperatorSet(opts *bind.TransactOpts, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "removeStrategiesFromOperatorSet", operatorSetId, strategies)
}

// RemoveStrategiesFromOperatorSet is a paid mutator transaction binding the contract method 0xce7b5e4b.
//
// Solidity: function removeStrategiesFromOperatorSet(uint32 operatorSetId, address[] strategies) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) RemoveStrategiesFromOperatorSet(operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.RemoveStrategiesFromOperatorSet(&_ContractAllocationManager.TransactOpts, operatorSetId, strategies)
}

// RemoveStrategiesFromOperatorSet is a paid mutator transaction binding the contract method 0xce7b5e4b.
//
// Solidity: function removeStrategiesFromOperatorSet(uint32 operatorSetId, address[] strategies) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) RemoveStrategiesFromOperatorSet(operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.RemoveStrategiesFromOperatorSet(&_ContractAllocationManager.TransactOpts, operatorSetId, strategies)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.RenounceOwnership(&_ContractAllocationManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.RenounceOwnership(&_ContractAllocationManager.TransactOpts)
}

// SetAVSRegistrar is a paid mutator transaction binding the contract method 0xf25f1610.
//
// Solidity: function setAVSRegistrar(address registrar) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) SetAVSRegistrar(opts *bind.TransactOpts, registrar common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "setAVSRegistrar", registrar)
}

// SetAVSRegistrar is a paid mutator transaction binding the contract method 0xf25f1610.
//
// Solidity: function setAVSRegistrar(address registrar) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) SetAVSRegistrar(registrar common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.SetAVSRegistrar(&_ContractAllocationManager.TransactOpts, registrar)
}

// SetAVSRegistrar is a paid mutator transaction binding the contract method 0xf25f1610.
//
// Solidity: function setAVSRegistrar(address registrar) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) SetAVSRegistrar(registrar common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.SetAVSRegistrar(&_ContractAllocationManager.TransactOpts, registrar)
}

// SetAllocationDelay is a paid mutator transaction binding the contract method 0x56c483e6.
//
// Solidity: function setAllocationDelay(address operator, uint32 delay) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) SetAllocationDelay(opts *bind.TransactOpts, operator common.Address, delay uint32) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "setAllocationDelay", operator, delay)
}

// SetAllocationDelay is a paid mutator transaction binding the contract method 0x56c483e6.
//
// Solidity: function setAllocationDelay(address operator, uint32 delay) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) SetAllocationDelay(operator common.Address, delay uint32) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.SetAllocationDelay(&_ContractAllocationManager.TransactOpts, operator, delay)
}

// SetAllocationDelay is a paid mutator transaction binding the contract method 0x56c483e6.
//
// Solidity: function setAllocationDelay(address operator, uint32 delay) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) SetAllocationDelay(operator common.Address, delay uint32) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.SetAllocationDelay(&_ContractAllocationManager.TransactOpts, operator, delay)
}

// SetAllocationDelay0 is a paid mutator transaction binding the contract method 0x5c489bb5.
//
// Solidity: function setAllocationDelay(uint32 delay) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) SetAllocationDelay0(opts *bind.TransactOpts, delay uint32) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "setAllocationDelay0", delay)
}

// SetAllocationDelay0 is a paid mutator transaction binding the contract method 0x5c489bb5.
//
// Solidity: function setAllocationDelay(uint32 delay) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) SetAllocationDelay0(delay uint32) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.SetAllocationDelay0(&_ContractAllocationManager.TransactOpts, delay)
}

// SetAllocationDelay0 is a paid mutator transaction binding the contract method 0x5c489bb5.
//
// Solidity: function setAllocationDelay(uint32 delay) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) SetAllocationDelay0(delay uint32) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.SetAllocationDelay0(&_ContractAllocationManager.TransactOpts, delay)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x0ea43e43.
//
// Solidity: function slashOperator((address,uint32,uint256,string) params) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) SlashOperator(opts *bind.TransactOpts, params IAllocationManagerTypesSlashingParams) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "slashOperator", params)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x0ea43e43.
//
// Solidity: function slashOperator((address,uint32,uint256,string) params) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) SlashOperator(params IAllocationManagerTypesSlashingParams) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.SlashOperator(&_ContractAllocationManager.TransactOpts, params)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x0ea43e43.
//
// Solidity: function slashOperator((address,uint32,uint256,string) params) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) SlashOperator(params IAllocationManagerTypesSlashingParams) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.SlashOperator(&_ContractAllocationManager.TransactOpts, params)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.TransferOwnership(&_ContractAllocationManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.TransferOwnership(&_ContractAllocationManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.Unpause(&_ContractAllocationManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.Unpause(&_ContractAllocationManager.TransactOpts, newPausedStatus)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "updateAVSMetadataURI", metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) UpdateAVSMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.UpdateAVSMetadataURI(&_ContractAllocationManager.TransactOpts, metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) UpdateAVSMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.UpdateAVSMetadataURI(&_ContractAllocationManager.TransactOpts, metadataURI)
}

// ContractAllocationManagerAVSMetadataURIUpdatedIterator is returned from FilterAVSMetadataURIUpdated and is used to iterate over the raw logs and unpacked data for AVSMetadataURIUpdated events raised by the ContractAllocationManager contract.
type ContractAllocationManagerAVSMetadataURIUpdatedIterator struct {
	Event *ContractAllocationManagerAVSMetadataURIUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAllocationManagerAVSMetadataURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerAVSMetadataURIUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAllocationManagerAVSMetadataURIUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAllocationManagerAVSMetadataURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerAVSMetadataURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerAVSMetadataURIUpdated represents a AVSMetadataURIUpdated event raised by the ContractAllocationManager contract.
type ContractAllocationManagerAVSMetadataURIUpdated struct {
	Avs         common.Address
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAVSMetadataURIUpdated is a free log retrieval operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterAVSMetadataURIUpdated(opts *bind.FilterOpts, avs []common.Address) (*ContractAllocationManagerAVSMetadataURIUpdatedIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "AVSMetadataURIUpdated", avsRule)
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerAVSMetadataURIUpdatedIterator{contract: _ContractAllocationManager.contract, event: "AVSMetadataURIUpdated", logs: logs, sub: sub}, nil
}

// WatchAVSMetadataURIUpdated is a free log subscription operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchAVSMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerAVSMetadataURIUpdated, avs []common.Address) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "AVSMetadataURIUpdated", avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerAVSMetadataURIUpdated)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "AVSMetadataURIUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAVSMetadataURIUpdated is a log parse operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParseAVSMetadataURIUpdated(log types.Log) (*ContractAllocationManagerAVSMetadataURIUpdated, error) {
	event := new(ContractAllocationManagerAVSMetadataURIUpdated)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "AVSMetadataURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllocationManagerAVSRegistrarSetIterator is returned from FilterAVSRegistrarSet and is used to iterate over the raw logs and unpacked data for AVSRegistrarSet events raised by the ContractAllocationManager contract.
type ContractAllocationManagerAVSRegistrarSetIterator struct {
	Event *ContractAllocationManagerAVSRegistrarSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAllocationManagerAVSRegistrarSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerAVSRegistrarSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAllocationManagerAVSRegistrarSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAllocationManagerAVSRegistrarSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerAVSRegistrarSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerAVSRegistrarSet represents a AVSRegistrarSet event raised by the ContractAllocationManager contract.
type ContractAllocationManagerAVSRegistrarSet struct {
	Avs       common.Address
	Registrar common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAVSRegistrarSet is a free log retrieval operation binding the contract event 0x2ae945c40c44dc0ec263f95609c3fdc6952e0aefa22d6374e44f2c997acedf85.
//
// Solidity: event AVSRegistrarSet(address avs, address registrar)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterAVSRegistrarSet(opts *bind.FilterOpts) (*ContractAllocationManagerAVSRegistrarSetIterator, error) {

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "AVSRegistrarSet")
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerAVSRegistrarSetIterator{contract: _ContractAllocationManager.contract, event: "AVSRegistrarSet", logs: logs, sub: sub}, nil
}

// WatchAVSRegistrarSet is a free log subscription operation binding the contract event 0x2ae945c40c44dc0ec263f95609c3fdc6952e0aefa22d6374e44f2c997acedf85.
//
// Solidity: event AVSRegistrarSet(address avs, address registrar)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchAVSRegistrarSet(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerAVSRegistrarSet) (event.Subscription, error) {

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "AVSRegistrarSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerAVSRegistrarSet)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "AVSRegistrarSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAVSRegistrarSet is a log parse operation binding the contract event 0x2ae945c40c44dc0ec263f95609c3fdc6952e0aefa22d6374e44f2c997acedf85.
//
// Solidity: event AVSRegistrarSet(address avs, address registrar)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParseAVSRegistrarSet(log types.Log) (*ContractAllocationManagerAVSRegistrarSet, error) {
	event := new(ContractAllocationManagerAVSRegistrarSet)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "AVSRegistrarSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllocationManagerAllocationDelaySetIterator is returned from FilterAllocationDelaySet and is used to iterate over the raw logs and unpacked data for AllocationDelaySet events raised by the ContractAllocationManager contract.
type ContractAllocationManagerAllocationDelaySetIterator struct {
	Event *ContractAllocationManagerAllocationDelaySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAllocationManagerAllocationDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerAllocationDelaySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAllocationManagerAllocationDelaySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAllocationManagerAllocationDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerAllocationDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerAllocationDelaySet represents a AllocationDelaySet event raised by the ContractAllocationManager contract.
type ContractAllocationManagerAllocationDelaySet struct {
	Operator    common.Address
	Delay       uint32
	EffectBlock uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAllocationDelaySet is a free log retrieval operation binding the contract event 0x4e85751d6331506c6c62335f207eb31f12a61e570f34f5c17640308785c6d4db.
//
// Solidity: event AllocationDelaySet(address operator, uint32 delay, uint32 effectBlock)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterAllocationDelaySet(opts *bind.FilterOpts) (*ContractAllocationManagerAllocationDelaySetIterator, error) {

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "AllocationDelaySet")
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerAllocationDelaySetIterator{contract: _ContractAllocationManager.contract, event: "AllocationDelaySet", logs: logs, sub: sub}, nil
}

// WatchAllocationDelaySet is a free log subscription operation binding the contract event 0x4e85751d6331506c6c62335f207eb31f12a61e570f34f5c17640308785c6d4db.
//
// Solidity: event AllocationDelaySet(address operator, uint32 delay, uint32 effectBlock)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchAllocationDelaySet(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerAllocationDelaySet) (event.Subscription, error) {

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "AllocationDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerAllocationDelaySet)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "AllocationDelaySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAllocationDelaySet is a log parse operation binding the contract event 0x4e85751d6331506c6c62335f207eb31f12a61e570f34f5c17640308785c6d4db.
//
// Solidity: event AllocationDelaySet(address operator, uint32 delay, uint32 effectBlock)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParseAllocationDelaySet(log types.Log) (*ContractAllocationManagerAllocationDelaySet, error) {
	event := new(ContractAllocationManagerAllocationDelaySet)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "AllocationDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllocationManagerAllocationUpdatedIterator is returned from FilterAllocationUpdated and is used to iterate over the raw logs and unpacked data for AllocationUpdated events raised by the ContractAllocationManager contract.
type ContractAllocationManagerAllocationUpdatedIterator struct {
	Event *ContractAllocationManagerAllocationUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAllocationManagerAllocationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerAllocationUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAllocationManagerAllocationUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAllocationManagerAllocationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerAllocationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerAllocationUpdated represents a AllocationUpdated event raised by the ContractAllocationManager contract.
type ContractAllocationManagerAllocationUpdated struct {
	Operator    common.Address
	OperatorSet OperatorSet
	Strategy    common.Address
	Magnitude   uint64
	EffectBlock uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAllocationUpdated is a free log retrieval operation binding the contract event 0x1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd.
//
// Solidity: event AllocationUpdated(address operator, (address,uint32) operatorSet, address strategy, uint64 magnitude, uint32 effectBlock)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterAllocationUpdated(opts *bind.FilterOpts) (*ContractAllocationManagerAllocationUpdatedIterator, error) {

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "AllocationUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerAllocationUpdatedIterator{contract: _ContractAllocationManager.contract, event: "AllocationUpdated", logs: logs, sub: sub}, nil
}

// WatchAllocationUpdated is a free log subscription operation binding the contract event 0x1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd.
//
// Solidity: event AllocationUpdated(address operator, (address,uint32) operatorSet, address strategy, uint64 magnitude, uint32 effectBlock)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchAllocationUpdated(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerAllocationUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "AllocationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerAllocationUpdated)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "AllocationUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAllocationUpdated is a log parse operation binding the contract event 0x1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd.
//
// Solidity: event AllocationUpdated(address operator, (address,uint32) operatorSet, address strategy, uint64 magnitude, uint32 effectBlock)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParseAllocationUpdated(log types.Log) (*ContractAllocationManagerAllocationUpdated, error) {
	event := new(ContractAllocationManagerAllocationUpdated)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "AllocationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllocationManagerEncumberedMagnitudeUpdatedIterator is returned from FilterEncumberedMagnitudeUpdated and is used to iterate over the raw logs and unpacked data for EncumberedMagnitudeUpdated events raised by the ContractAllocationManager contract.
type ContractAllocationManagerEncumberedMagnitudeUpdatedIterator struct {
	Event *ContractAllocationManagerEncumberedMagnitudeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAllocationManagerEncumberedMagnitudeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerEncumberedMagnitudeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAllocationManagerEncumberedMagnitudeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAllocationManagerEncumberedMagnitudeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerEncumberedMagnitudeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerEncumberedMagnitudeUpdated represents a EncumberedMagnitudeUpdated event raised by the ContractAllocationManager contract.
type ContractAllocationManagerEncumberedMagnitudeUpdated struct {
	Operator            common.Address
	Strategy            common.Address
	EncumberedMagnitude uint64
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterEncumberedMagnitudeUpdated is a free log retrieval operation binding the contract event 0xacf9095feb3a370c9cf692421c69ef320d4db5c66e6a7d29c7694eb02364fc55.
//
// Solidity: event EncumberedMagnitudeUpdated(address operator, address strategy, uint64 encumberedMagnitude)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterEncumberedMagnitudeUpdated(opts *bind.FilterOpts) (*ContractAllocationManagerEncumberedMagnitudeUpdatedIterator, error) {

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "EncumberedMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerEncumberedMagnitudeUpdatedIterator{contract: _ContractAllocationManager.contract, event: "EncumberedMagnitudeUpdated", logs: logs, sub: sub}, nil
}

// WatchEncumberedMagnitudeUpdated is a free log subscription operation binding the contract event 0xacf9095feb3a370c9cf692421c69ef320d4db5c66e6a7d29c7694eb02364fc55.
//
// Solidity: event EncumberedMagnitudeUpdated(address operator, address strategy, uint64 encumberedMagnitude)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchEncumberedMagnitudeUpdated(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerEncumberedMagnitudeUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "EncumberedMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerEncumberedMagnitudeUpdated)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "EncumberedMagnitudeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEncumberedMagnitudeUpdated is a log parse operation binding the contract event 0xacf9095feb3a370c9cf692421c69ef320d4db5c66e6a7d29c7694eb02364fc55.
//
// Solidity: event EncumberedMagnitudeUpdated(address operator, address strategy, uint64 encumberedMagnitude)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParseEncumberedMagnitudeUpdated(log types.Log) (*ContractAllocationManagerEncumberedMagnitudeUpdated, error) {
	event := new(ContractAllocationManagerEncumberedMagnitudeUpdated)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "EncumberedMagnitudeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllocationManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractAllocationManager contract.
type ContractAllocationManagerInitializedIterator struct {
	Event *ContractAllocationManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAllocationManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAllocationManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAllocationManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerInitialized represents a Initialized event raised by the ContractAllocationManager contract.
type ContractAllocationManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractAllocationManagerInitializedIterator, error) {

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerInitializedIterator{contract: _ContractAllocationManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerInitialized)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParseInitialized(log types.Log) (*ContractAllocationManagerInitialized, error) {
	event := new(ContractAllocationManagerInitialized)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllocationManagerMaxMagnitudeUpdatedIterator is returned from FilterMaxMagnitudeUpdated and is used to iterate over the raw logs and unpacked data for MaxMagnitudeUpdated events raised by the ContractAllocationManager contract.
type ContractAllocationManagerMaxMagnitudeUpdatedIterator struct {
	Event *ContractAllocationManagerMaxMagnitudeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAllocationManagerMaxMagnitudeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerMaxMagnitudeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAllocationManagerMaxMagnitudeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAllocationManagerMaxMagnitudeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerMaxMagnitudeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerMaxMagnitudeUpdated represents a MaxMagnitudeUpdated event raised by the ContractAllocationManager contract.
type ContractAllocationManagerMaxMagnitudeUpdated struct {
	Operator     common.Address
	Strategy     common.Address
	MaxMagnitude uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMaxMagnitudeUpdated is a free log retrieval operation binding the contract event 0x1c6458079a41077d003c11faf9bf097e693bd67979e4e6500bac7b29db779b5c.
//
// Solidity: event MaxMagnitudeUpdated(address operator, address strategy, uint64 maxMagnitude)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterMaxMagnitudeUpdated(opts *bind.FilterOpts) (*ContractAllocationManagerMaxMagnitudeUpdatedIterator, error) {

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "MaxMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerMaxMagnitudeUpdatedIterator{contract: _ContractAllocationManager.contract, event: "MaxMagnitudeUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxMagnitudeUpdated is a free log subscription operation binding the contract event 0x1c6458079a41077d003c11faf9bf097e693bd67979e4e6500bac7b29db779b5c.
//
// Solidity: event MaxMagnitudeUpdated(address operator, address strategy, uint64 maxMagnitude)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchMaxMagnitudeUpdated(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerMaxMagnitudeUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "MaxMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerMaxMagnitudeUpdated)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "MaxMagnitudeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMaxMagnitudeUpdated is a log parse operation binding the contract event 0x1c6458079a41077d003c11faf9bf097e693bd67979e4e6500bac7b29db779b5c.
//
// Solidity: event MaxMagnitudeUpdated(address operator, address strategy, uint64 maxMagnitude)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParseMaxMagnitudeUpdated(log types.Log) (*ContractAllocationManagerMaxMagnitudeUpdated, error) {
	event := new(ContractAllocationManagerMaxMagnitudeUpdated)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "MaxMagnitudeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllocationManagerOperatorAddedToOperatorSetIterator is returned from FilterOperatorAddedToOperatorSet and is used to iterate over the raw logs and unpacked data for OperatorAddedToOperatorSet events raised by the ContractAllocationManager contract.
type ContractAllocationManagerOperatorAddedToOperatorSetIterator struct {
	Event *ContractAllocationManagerOperatorAddedToOperatorSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAllocationManagerOperatorAddedToOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerOperatorAddedToOperatorSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAllocationManagerOperatorAddedToOperatorSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAllocationManagerOperatorAddedToOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerOperatorAddedToOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerOperatorAddedToOperatorSet represents a OperatorAddedToOperatorSet event raised by the ContractAllocationManager contract.
type ContractAllocationManagerOperatorAddedToOperatorSet struct {
	Operator    common.Address
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorAddedToOperatorSet is a free log retrieval operation binding the contract event 0x43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e.
//
// Solidity: event OperatorAddedToOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterOperatorAddedToOperatorSet(opts *bind.FilterOpts, operator []common.Address) (*ContractAllocationManagerOperatorAddedToOperatorSetIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "OperatorAddedToOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerOperatorAddedToOperatorSetIterator{contract: _ContractAllocationManager.contract, event: "OperatorAddedToOperatorSet", logs: logs, sub: sub}, nil
}

// WatchOperatorAddedToOperatorSet is a free log subscription operation binding the contract event 0x43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e.
//
// Solidity: event OperatorAddedToOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchOperatorAddedToOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerOperatorAddedToOperatorSet, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "OperatorAddedToOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerOperatorAddedToOperatorSet)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "OperatorAddedToOperatorSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorAddedToOperatorSet is a log parse operation binding the contract event 0x43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e.
//
// Solidity: event OperatorAddedToOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParseOperatorAddedToOperatorSet(log types.Log) (*ContractAllocationManagerOperatorAddedToOperatorSet, error) {
	event := new(ContractAllocationManagerOperatorAddedToOperatorSet)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "OperatorAddedToOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllocationManagerOperatorRemovedFromOperatorSetIterator is returned from FilterOperatorRemovedFromOperatorSet and is used to iterate over the raw logs and unpacked data for OperatorRemovedFromOperatorSet events raised by the ContractAllocationManager contract.
type ContractAllocationManagerOperatorRemovedFromOperatorSetIterator struct {
	Event *ContractAllocationManagerOperatorRemovedFromOperatorSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAllocationManagerOperatorRemovedFromOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerOperatorRemovedFromOperatorSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAllocationManagerOperatorRemovedFromOperatorSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAllocationManagerOperatorRemovedFromOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerOperatorRemovedFromOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerOperatorRemovedFromOperatorSet represents a OperatorRemovedFromOperatorSet event raised by the ContractAllocationManager contract.
type ContractAllocationManagerOperatorRemovedFromOperatorSet struct {
	Operator    common.Address
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemovedFromOperatorSet is a free log retrieval operation binding the contract event 0xad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe.
//
// Solidity: event OperatorRemovedFromOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterOperatorRemovedFromOperatorSet(opts *bind.FilterOpts, operator []common.Address) (*ContractAllocationManagerOperatorRemovedFromOperatorSetIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "OperatorRemovedFromOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerOperatorRemovedFromOperatorSetIterator{contract: _ContractAllocationManager.contract, event: "OperatorRemovedFromOperatorSet", logs: logs, sub: sub}, nil
}

// WatchOperatorRemovedFromOperatorSet is a free log subscription operation binding the contract event 0xad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe.
//
// Solidity: event OperatorRemovedFromOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchOperatorRemovedFromOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerOperatorRemovedFromOperatorSet, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "OperatorRemovedFromOperatorSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerOperatorRemovedFromOperatorSet)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "OperatorRemovedFromOperatorSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorRemovedFromOperatorSet is a log parse operation binding the contract event 0xad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe.
//
// Solidity: event OperatorRemovedFromOperatorSet(address indexed operator, (address,uint32) operatorSet)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParseOperatorRemovedFromOperatorSet(log types.Log) (*ContractAllocationManagerOperatorRemovedFromOperatorSet, error) {
	event := new(ContractAllocationManagerOperatorRemovedFromOperatorSet)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "OperatorRemovedFromOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllocationManagerOperatorSetCreatedIterator is returned from FilterOperatorSetCreated and is used to iterate over the raw logs and unpacked data for OperatorSetCreated events raised by the ContractAllocationManager contract.
type ContractAllocationManagerOperatorSetCreatedIterator struct {
	Event *ContractAllocationManagerOperatorSetCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAllocationManagerOperatorSetCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerOperatorSetCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAllocationManagerOperatorSetCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAllocationManagerOperatorSetCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerOperatorSetCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerOperatorSetCreated represents a OperatorSetCreated event raised by the ContractAllocationManager contract.
type ContractAllocationManagerOperatorSetCreated struct {
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetCreated is a free log retrieval operation binding the contract event 0x31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c.
//
// Solidity: event OperatorSetCreated((address,uint32) operatorSet)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterOperatorSetCreated(opts *bind.FilterOpts) (*ContractAllocationManagerOperatorSetCreatedIterator, error) {

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "OperatorSetCreated")
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerOperatorSetCreatedIterator{contract: _ContractAllocationManager.contract, event: "OperatorSetCreated", logs: logs, sub: sub}, nil
}

// WatchOperatorSetCreated is a free log subscription operation binding the contract event 0x31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c.
//
// Solidity: event OperatorSetCreated((address,uint32) operatorSet)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchOperatorSetCreated(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerOperatorSetCreated) (event.Subscription, error) {

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "OperatorSetCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerOperatorSetCreated)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "OperatorSetCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorSetCreated is a log parse operation binding the contract event 0x31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c.
//
// Solidity: event OperatorSetCreated((address,uint32) operatorSet)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParseOperatorSetCreated(log types.Log) (*ContractAllocationManagerOperatorSetCreated, error) {
	event := new(ContractAllocationManagerOperatorSetCreated)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "OperatorSetCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllocationManagerOperatorSlashedIterator is returned from FilterOperatorSlashed and is used to iterate over the raw logs and unpacked data for OperatorSlashed events raised by the ContractAllocationManager contract.
type ContractAllocationManagerOperatorSlashedIterator struct {
	Event *ContractAllocationManagerOperatorSlashed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAllocationManagerOperatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerOperatorSlashed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAllocationManagerOperatorSlashed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAllocationManagerOperatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerOperatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerOperatorSlashed represents a OperatorSlashed event raised by the ContractAllocationManager contract.
type ContractAllocationManagerOperatorSlashed struct {
	Operator    common.Address
	OperatorSet OperatorSet
	Strategies  []common.Address
	WadSlashed  []*big.Int
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSlashed is a free log retrieval operation binding the contract event 0x80969ad29428d6797ee7aad084f9e4a42a82fc506dcd2ca3b6fb431f85ccebe5.
//
// Solidity: event OperatorSlashed(address operator, (address,uint32) operatorSet, address[] strategies, uint256[] wadSlashed, string description)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterOperatorSlashed(opts *bind.FilterOpts) (*ContractAllocationManagerOperatorSlashedIterator, error) {

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerOperatorSlashedIterator{contract: _ContractAllocationManager.contract, event: "OperatorSlashed", logs: logs, sub: sub}, nil
}

// WatchOperatorSlashed is a free log subscription operation binding the contract event 0x80969ad29428d6797ee7aad084f9e4a42a82fc506dcd2ca3b6fb431f85ccebe5.
//
// Solidity: event OperatorSlashed(address operator, (address,uint32) operatorSet, address[] strategies, uint256[] wadSlashed, string description)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchOperatorSlashed(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerOperatorSlashed) (event.Subscription, error) {

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerOperatorSlashed)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorSlashed is a log parse operation binding the contract event 0x80969ad29428d6797ee7aad084f9e4a42a82fc506dcd2ca3b6fb431f85ccebe5.
//
// Solidity: event OperatorSlashed(address operator, (address,uint32) operatorSet, address[] strategies, uint256[] wadSlashed, string description)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParseOperatorSlashed(log types.Log) (*ContractAllocationManagerOperatorSlashed, error) {
	event := new(ContractAllocationManagerOperatorSlashed)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllocationManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractAllocationManager contract.
type ContractAllocationManagerOwnershipTransferredIterator struct {
	Event *ContractAllocationManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAllocationManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAllocationManagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAllocationManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractAllocationManager contract.
type ContractAllocationManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractAllocationManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerOwnershipTransferredIterator{contract: _ContractAllocationManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerOwnershipTransferred)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractAllocationManagerOwnershipTransferred, error) {
	event := new(ContractAllocationManagerOwnershipTransferred)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllocationManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractAllocationManager contract.
type ContractAllocationManagerPausedIterator struct {
	Event *ContractAllocationManagerPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAllocationManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAllocationManagerPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAllocationManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerPaused represents a Paused event raised by the ContractAllocationManager contract.
type ContractAllocationManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractAllocationManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerPausedIterator{contract: _ContractAllocationManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerPaused)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParsePaused(log types.Log) (*ContractAllocationManagerPaused, error) {
	event := new(ContractAllocationManagerPaused)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllocationManagerStrategyAddedToOperatorSetIterator is returned from FilterStrategyAddedToOperatorSet and is used to iterate over the raw logs and unpacked data for StrategyAddedToOperatorSet events raised by the ContractAllocationManager contract.
type ContractAllocationManagerStrategyAddedToOperatorSetIterator struct {
	Event *ContractAllocationManagerStrategyAddedToOperatorSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAllocationManagerStrategyAddedToOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerStrategyAddedToOperatorSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAllocationManagerStrategyAddedToOperatorSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAllocationManagerStrategyAddedToOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerStrategyAddedToOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerStrategyAddedToOperatorSet represents a StrategyAddedToOperatorSet event raised by the ContractAllocationManager contract.
type ContractAllocationManagerStrategyAddedToOperatorSet struct {
	OperatorSet OperatorSet
	Strategy    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStrategyAddedToOperatorSet is a free log retrieval operation binding the contract event 0x7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b.
//
// Solidity: event StrategyAddedToOperatorSet((address,uint32) operatorSet, address strategy)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterStrategyAddedToOperatorSet(opts *bind.FilterOpts) (*ContractAllocationManagerStrategyAddedToOperatorSetIterator, error) {

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "StrategyAddedToOperatorSet")
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerStrategyAddedToOperatorSetIterator{contract: _ContractAllocationManager.contract, event: "StrategyAddedToOperatorSet", logs: logs, sub: sub}, nil
}

// WatchStrategyAddedToOperatorSet is a free log subscription operation binding the contract event 0x7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b.
//
// Solidity: event StrategyAddedToOperatorSet((address,uint32) operatorSet, address strategy)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchStrategyAddedToOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerStrategyAddedToOperatorSet) (event.Subscription, error) {

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "StrategyAddedToOperatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerStrategyAddedToOperatorSet)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "StrategyAddedToOperatorSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStrategyAddedToOperatorSet is a log parse operation binding the contract event 0x7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b.
//
// Solidity: event StrategyAddedToOperatorSet((address,uint32) operatorSet, address strategy)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParseStrategyAddedToOperatorSet(log types.Log) (*ContractAllocationManagerStrategyAddedToOperatorSet, error) {
	event := new(ContractAllocationManagerStrategyAddedToOperatorSet)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "StrategyAddedToOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllocationManagerStrategyRemovedFromOperatorSetIterator is returned from FilterStrategyRemovedFromOperatorSet and is used to iterate over the raw logs and unpacked data for StrategyRemovedFromOperatorSet events raised by the ContractAllocationManager contract.
type ContractAllocationManagerStrategyRemovedFromOperatorSetIterator struct {
	Event *ContractAllocationManagerStrategyRemovedFromOperatorSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAllocationManagerStrategyRemovedFromOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerStrategyRemovedFromOperatorSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAllocationManagerStrategyRemovedFromOperatorSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAllocationManagerStrategyRemovedFromOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerStrategyRemovedFromOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerStrategyRemovedFromOperatorSet represents a StrategyRemovedFromOperatorSet event raised by the ContractAllocationManager contract.
type ContractAllocationManagerStrategyRemovedFromOperatorSet struct {
	OperatorSet OperatorSet
	Strategy    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStrategyRemovedFromOperatorSet is a free log retrieval operation binding the contract event 0x7b4b073d80dcac55a11177d8459ad9f664ceeb91f71f27167bb14f8152a7eeee.
//
// Solidity: event StrategyRemovedFromOperatorSet((address,uint32) operatorSet, address strategy)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterStrategyRemovedFromOperatorSet(opts *bind.FilterOpts) (*ContractAllocationManagerStrategyRemovedFromOperatorSetIterator, error) {

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "StrategyRemovedFromOperatorSet")
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerStrategyRemovedFromOperatorSetIterator{contract: _ContractAllocationManager.contract, event: "StrategyRemovedFromOperatorSet", logs: logs, sub: sub}, nil
}

// WatchStrategyRemovedFromOperatorSet is a free log subscription operation binding the contract event 0x7b4b073d80dcac55a11177d8459ad9f664ceeb91f71f27167bb14f8152a7eeee.
//
// Solidity: event StrategyRemovedFromOperatorSet((address,uint32) operatorSet, address strategy)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchStrategyRemovedFromOperatorSet(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerStrategyRemovedFromOperatorSet) (event.Subscription, error) {

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "StrategyRemovedFromOperatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerStrategyRemovedFromOperatorSet)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "StrategyRemovedFromOperatorSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStrategyRemovedFromOperatorSet is a log parse operation binding the contract event 0x7b4b073d80dcac55a11177d8459ad9f664ceeb91f71f27167bb14f8152a7eeee.
//
// Solidity: event StrategyRemovedFromOperatorSet((address,uint32) operatorSet, address strategy)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParseStrategyRemovedFromOperatorSet(log types.Log) (*ContractAllocationManagerStrategyRemovedFromOperatorSet, error) {
	event := new(ContractAllocationManagerStrategyRemovedFromOperatorSet)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "StrategyRemovedFromOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAllocationManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractAllocationManager contract.
type ContractAllocationManagerUnpausedIterator struct {
	Event *ContractAllocationManagerUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAllocationManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllocationManagerUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAllocationManagerUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAllocationManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllocationManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllocationManagerUnpaused represents a Unpaused event raised by the ContractAllocationManager contract.
type ContractAllocationManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractAllocationManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractAllocationManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractAllocationManagerUnpausedIterator{contract: _ContractAllocationManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractAllocationManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractAllocationManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllocationManagerUnpaused)
				if err := _ContractAllocationManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractAllocationManager *ContractAllocationManagerFilterer) ParseUnpaused(log types.Log) (*ContractAllocationManagerUnpaused, error) {
	event := new(ContractAllocationManagerUnpaused)
	if err := _ContractAllocationManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
