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
	Strategies    []common.Address
	WadsToSlash   []*big.Int
	Description   string
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// ContractAllocationManagerMetaData contains all meta data concerning the ContractAllocationManager contract.
var ContractAllocationManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_delegation\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"},{\"name\":\"_DEALLOCATION_DELAY\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_ALLOCATION_CONFIGURATION_DELAY\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ALLOCATION_CONFIGURATION_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEALLOCATION_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addStrategiesToOperatorSet\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"clearDeallocationQueue\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"numToClear\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createOperatorSets\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.CreateSetParams[]\",\"components\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterFromOperatorSets\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.DeregisterParams\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"encumberedMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAVSRegistrar\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAVSRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatableMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatedSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocation\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.Allocation\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocationDelay\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocations\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.Allocation[]\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudes\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudes\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudesAtBlock\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMemberCount\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMembers\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinimumSlashableStake\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"futureBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"slashableStake\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetCount\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRegisteredSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStrategiesInOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStrategyAllocations\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.Allocation[]\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isMemberOfOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"modifyAllocations\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.AllocateParams[]\",\"components\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"newMagnitudes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerForOperatorSets\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.RegisterParams\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeStrategiesFromOperatorSet\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAVSRegistrar\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registrar\",\"type\":\"address\",\"internalType\":\"contractIAVSRegistrar\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAllocationDelay\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashOperator\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.SlashingParams\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"wadsToSlash\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AVSMetadataURIUpdated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AVSRegistrarSet\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"registrar\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIAVSRegistrar\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllocationDelaySet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllocationUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"magnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"effectBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EncumberedMagnitudeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"encumberedMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxMagnitudeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"maxMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromOperatorSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetCreated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSlashed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategies\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"contractIStrategy[]\"},{\"name\":\"wadSlashed\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"description\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyAddedToOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyRemovedFromOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyMemberOfSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Empty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientMagnitude\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCaller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPermissions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSnapshotOrdering\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWadToSlash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxStrategiesExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ModificationAlreadyPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotMemberOfSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotSlashable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OutOfBounds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SameMagnitude\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategiesMustBeInAscendingOrder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyAlreadyInOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotInOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UninitializedAllocationDelay\",\"inputs\":[]}]",
	Bin: "0x61012060405234801561001157600080fd5b50604051615c93380380615c9383398101604081905261003091610186565b82858383876001600160a01b03811661005c576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b0390811660805292831660a05263ffffffff91821660c0521660e052166101005261008c610096565b50505050506101f3565b600054610100900460ff16156101025760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff90811614610153576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b038116811461016a57600080fd5b50565b805163ffffffff8116811461018157600080fd5b919050565b600080600080600060a0868803121561019e57600080fd5b85516101a981610155565b60208701519095506101ba81610155565b60408701519094506101cb81610155565b92506101d96060870161016d565b91506101e76080870161016d565b90509295509295909350565b60805160a05160c05160e05161010051615a0861028b60003960008181610403015261380f0152600081816105520152613e2d015260008181610328015281816122fd0152612a1b01526000818161070501528181610d090152818161167201528181611d2a01528181611d940152612cdb0152600081816105790152818161079e01528181611e3b01526134790152615a086000f3fe608060405234801561001057600080fd5b506004361061028a5760003560e01c80636e875dba1161015c578063a984eb3a116100ce578063c221d8ae11610087578063c221d8ae146106c7578063cd6dc687146106da578063d3d96ff4146106ed578063df5cf72314610700578063f2fde38b14610727578063fabc1cbc1461073a57600080fd5b8063a984eb3a14610618578063adc2e3d91461064c578063b2447af71461065f578063b66bd98914610672578063b9fbaed114610685578063ba1a84e5146106b457600080fd5b80638ce64854116101205780638ce648541461059b5780638da5cb5b146105bb57806394d7d00c146105cc578063952899ee146105df578063a9333ec8146105f2578063a98218211461060557600080fd5b80636e875dba1461051f578063715018a61461053257806379ae50cd1461053a5780637bc1ef611461054d578063886f11951461057457600080fd5b80634657e26a11610200578063595c6a67116101b9578063595c6a67146104915780635ac86ab7146104995780635c975abb146104bc578063670d3ba2146104ce5780636cfb4481146104e15780636e3492b51461050c57600080fd5b80634657e26a146103fe5780634a10ffe5146104255780634b5046ef1461044557806350feea2014610458578063547afb871461046b57806356c483e61461047e57600080fd5b80632981eb77116102525780632981eb77146103235780632bab2c4a1461035f578063304c10cd1461037f57806336352057146103aa57806340120dab146103bd5780634177a87c146103de57600080fd5b806310e1b9b81461028f578063136439dd146102b857806315fe5028146102cd578063260dc758146102ed578063261f84e014610310575b600080fd5b6102a261029d366004614891565b61074d565b6040516102af91906148db565b60405180910390f35b6102cb6102c636600461490e565b610789565b005b6102e06102db366004614927565b610860565b6040516102af91906149a9565b6103006102fb3660046149bc565b61097d565b60405190151581526020016102af565b6102cb61031e366004614a1c565b6109b5565b61034a7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016102af565b61037261036d366004614b09565b610cb4565b6040516102af9190614bc4565b61039261038d366004614927565b610fb0565b6040516001600160a01b0390911681526020016102af565b6102cb6103b8366004614c29565b610fe0565b6103d06103cb366004614c7f565b6117cc565b6040516102af929190614d10565b6103f16103ec3660046149bc565b61194f565b6040516102af9190614d70565b6103927f000000000000000000000000000000000000000000000000000000000000000081565b610438610433366004614d83565b611976565b6040516102af9190614dc9565b6102cb610453366004614e15565b611a20565b6102cb610466366004614e9b565b611adc565b610438610479366004614efd565b611c75565b6102cb61048c366004614f4c565b611d1f565b6102cb611e26565b6103006104a7366004614f81565b606654600160ff9092169190911b9081161490565b6066545b6040519081526020016102af565b6103006104dc366004614fa4565b611ed8565b6104f46104ef366004614c7f565b611eeb565b6040516001600160401b0390911681526020016102af565b6102cb61051a366004614fe9565b61205c565b6103f161052d3660046149bc565b612440565b6102cb612452565b6102e0610548366004614927565b612464565b61034a7f000000000000000000000000000000000000000000000000000000000000000081565b6103927f000000000000000000000000000000000000000000000000000000000000000081565b6105ae6105a936600461501d565b612543565b6040516102af9190615064565b6033546001600160a01b0316610392565b6104386105da366004615077565b61260c565b6102cb6105ed3660046150d6565b6126fc565b6104f4610600366004614c7f565b612bcc565b6102cb61061336600461528b565b612bfc565b6104f4610626366004614c7f565b60a26020908152600092835260408084209091529082529020546001600160401b031681565b6102cb61065a366004615310565b612c6c565b6104c061066d3660046149bc565b612fca565b6102cb610680366004614e9b565b612fef565b610698610693366004614927565b61314d565b60408051921515835263ffffffff9091166020830152016102af565b6104c06106c2366004614927565b6131e8565b6103f16106d5366004614fa4565b613209565b6102cb6106e8366004615355565b61323b565b6102cb6106fb366004614c7f565b61335e565b6103927f000000000000000000000000000000000000000000000000000000000000000081565b6102cb610735366004614927565b6133fe565b6102cb61074836600461490e565b613477565b604080516060810182526000808252602082018190529181018290529061077d856107778661358f565b856135f4565b925050505b9392505050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa1580156107ed573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108119190615381565b61082e57604051631d77d47760e21b815260040160405180910390fd5b60665481811681146108535760405163c61dca5d60e01b815260040160405180910390fd5b61085c82613765565b5050565b6001600160a01b0381166000908152609d6020526040812060609190610885906137a2565b90506000816001600160401b038111156108a1576108a16147b0565b6040519080825280602002602001820160405280156108e657816020015b60408051808201909152600080825260208201528152602001906001900390816108bf5790505b50905060005b82811015610975576001600160a01b0385166000908152609d602052604090206109509061091a90836137ac565b60408051808201909152600080825260208201525060408051808201909152606082901c815263ffffffff909116602082015290565b828281518110610962576109626153a3565b60209081029190910101526001016108ec565b509392505050565b60208082015182516001600160a01b031660009081526098909252604082206109af9163ffffffff908116906137b816565b92915050565b826109bf816137d0565b6109dc5760405163932d94f760e01b815260040160405180910390fd5b60005b82811015610cad5760218484838181106109fb576109fb6153a3565b9050602002810190610a0d91906153b9565b610a1b9060208101906153d9565b90501115610a3c576040516301a1443960e31b815260040160405180910390fd5b60006040518060400160405280876001600160a01b03168152602001868685818110610a6a57610a6a6153a3565b9050602002810190610a7c91906153b9565b610a8a906020810190615422565b63ffffffff168152509050610ad6816020015163ffffffff1660986000896001600160a01b03166001600160a01b0316815260200190815260200160002061387e90919063ffffffff16565b610af357604051631fb1705560e21b815260040160405180910390fd5b7f31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c6040518060400160405280886001600160a01b03168152602001836020015163ffffffff16815250604051610b49919061543d565b60405180910390a16000610b5c8261358f565b905060005b868685818110610b7357610b736153a3565b9050602002810190610b8591906153b9565b610b939060208101906153d9565b9050811015610ca257610c0a878786818110610bb157610bb16153a3565b9050602002810190610bc391906153b9565b610bd19060208101906153d9565b83818110610be157610be16153a3565b9050602002016020810190610bf69190614927565b60008481526099602052604090209061388a565b507f7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b83888887818110610c3f57610c3f6153a3565b9050602002810190610c5191906153b9565b610c5f9060208101906153d9565b84818110610c6f57610c6f6153a3565b9050602002016020810190610c849190614927565b604051610c9292919061544b565b60405180910390a1600101610b61565b5050506001016109df565b5050505050565b606083516001600160401b03811115610ccf57610ccf6147b0565b604051908082528060200260200182016040528015610d0257816020015b6060815260200190600190039081610ced5790505b50905060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f0e0e67686866040518363ffffffff1660e01b8152600401610d55929190615471565b600060405180830381865afa158015610d72573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610d9a9190810190615496565b905060005b8551811015610fa6576000868281518110610dbc57610dbc6153a3565b6020026020010151905085516001600160401b03811115610ddf57610ddf6147b0565b604051908082528060200260200182016040528015610e08578160200160208202803683370190505b50848381518110610e1b57610e1b6153a3565b602002602001018190525060005b8651811015610f9c576000878281518110610e4657610e466153a3565b6020908102919091018101516001600160a01b03808616600090815260a1845260408082209284168252919093528220909250610e829061389f565b9050806001600160401b0316600003610e9c575050610f94565b6000610ea9858d8561074d565b90508863ffffffff16816040015163ffffffff1611158015610ed2575060008160200151600f0b125b15610ef557610ee9816000015182602001516138b3565b6001600160401b031681525b8051600090610f11906001600160401b039081169085166138c8565b9050610f5881898981518110610f2957610f296153a3565b60200260200101518781518110610f4257610f426153a3565b60200260200101516138dd90919063ffffffff16565b898881518110610f6a57610f6a6153a3565b60200260200101518681518110610f8357610f836153a3565b602002602001018181525050505050505b600101610e29565b5050600101610d9f565b5050949350505050565b6001600160a01b038082166000908152609760205260408120549091168015610fd95780610782565b5090919050565b6066546001906002908116036110095760405163840a48d560e01b815260040160405180910390fd5b82611013816137d0565b6110305760405163932d94f760e01b815260040160405180910390fd5b60006040518060400160405280866001600160a01b0316815260200185602001602081019061105f9190615422565b63ffffffff1690529050600061108161107b6020870187614927565b836138f2565b60208084015184516001600160a01b03166000908152609890925260409091209192506110b8919063ffffffff908116906137b816565b6110d557604051631fb1705560e21b815260040160405180910390fd5b806110f35760405163ebbff49760e01b815260040160405180910390fd5b600061110260408701876153d9565b90506001600160401b0381111561111b5761111b6147b0565b604051908082528060200260200182016040528015611144578160200160208202803683370190505b50905060005b61115760408801886153d9565b905081101561175d578015806111ea575061117560408801886153d9565b6111806001846155c0565b81811061118f5761118f6153a3565b90506020020160208101906111a49190614927565b6001600160a01b03166111ba60408901896153d9565b838181106111ca576111ca6153a3565b90506020020160208101906111df9190614927565b6001600160a01b0316115b61120757604051639f1c805360e01b815260040160405180910390fd5b61121460608801886153d9565b82818110611224576112246153a3565b9050602002013560001080156112655750670de0b6b3a764000061124b60608901896153d9565b8381811061125b5761125b6153a3565b9050602002013511155b61128257604051631353603160e01b815260040160405180910390fd5b6112e061129260408901896153d9565b838181106112a2576112a26153a3565b90506020020160208101906112b79190614927565b609960006112c48861358f565b815260200190815260200160002061396990919063ffffffff16565b6112fd576040516331bc342760e11b815260040160405180910390fd5b60008061135061131060208b018b614927565b6113198861358f565b61132660408d018d6153d9565b87818110611336576113366153a3565b905060200201602081019061134b9190614927565b6135f4565b805191935091506001600160401b031660000361136e575050611755565b60006113aa61138060608c018c6153d9565b86818110611390576113906153a3565b85516001600160401b03169260209091020135905061398b565b83519091506113c56001600160401b038084169083166138c8565b8686815181106113d7576113d76153a3565b60200260200101818152505081836000018181516113f591906155d3565b6001600160401b03169052508351829085906114129083906155d3565b6001600160401b03169052506020840180518391906114329083906155d3565b6001600160401b031690525060208301516000600f9190910b121561155057600061149761146360608e018e6153d9565b88818110611473576114736153a3565b905060200201358560200151611488906155f2565b6001600160801b03169061398b565b9050806001600160401b0316846020018181516114b49190615618565b600f0b9052507f1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd6114e860208e018e614927565b8a8e80604001906114f991906153d9565b8a818110611509576115096153a3565b905060200201602081019061151e9190614927565b611530886000015189602001516138b3565b8860400151604051611546959493929190615645565b60405180910390a1505b6115a261156060208d018d614927565b6115698a61358f565b61157660408f018f6153d9565b89818110611586576115866153a3565b905060200201602081019061159b9190614927565b87876139a2565b7f1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd6115d060208d018d614927565b896115de60408f018f6153d9565b898181106115ee576115ee6153a3565b90506020020160208101906116039190614927565b865160405161161794939291904390615645565b60405180910390a161166861162f60208d018d614927565b61163c60408e018e6153d9565b8881811061164c5761164c6153a3565b90506020020160208101906116619190614927565b8651613bec565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001663ee74937f6116a460208e018e614927565b6116b160408f018f6153d9565b898181106116c1576116c16153a3565b90506020020160208101906116d69190614927565b875160405160e085901b6001600160e01b03191681526001600160a01b0393841660048201529290911660248301526001600160401b038086166044840152166064820152608401600060405180830381600087803b15801561173857600080fd5b505af115801561174c573d6000803e3d6000fd5b50505050505050505b60010161114a565b507f80969ad29428d6797ee7aad084f9e4a42a82fc506dcd2ca3b6fb431f85ccebe561178c6020880188614927565b8461179a60408a018a6153d9565b856117a860808d018d615696565b6040516117bb9796959493929190615705565b60405180910390a150505050505050565b6001600160a01b0382166000908152609d6020526040812060609182916117f2906137a2565b90506000816001600160401b0381111561180e5761180e6147b0565b60405190808252806020026020018201604052801561185357816020015b604080518082019091526000808252602082015281526020019060019003908161182c5790505b5090506000826001600160401b03811115611870576118706147b0565b6040519080825280602002602001820160405280156118bb57816020015b604080516060810182526000808252602080830182905292820152825260001990920191018161188e5790505b50905060005b83811015611940576001600160a01b0388166000908152609d602052604081206118ef9061091a90846137ac565b905080848381518110611904576119046153a3565b602002602001018190525061191a89828a61074d565b83838151811061192c5761192c6153a3565b6020908102919091010152506001016118c1565b509093509150505b9250929050565b60606000610782609960006119638661358f565b8152602001908152602001600020613c6f565b6060600083516001600160401b03811115611993576119936147b0565b6040519080825280602002602001820160405280156119bc578160200160208202803683370190505b50905060005b8451811015610975576119ee8582815181106119e0576119e06153a3565b602002602001015185612bcc565b828281518110611a0057611a006153a3565b6001600160401b03909216602092830291909101909101526001016119c2565b606654600090600190811603611a495760405163840a48d560e01b815260040160405180910390fd5b838214611a69576040516343714afd60e01b815260040160405180910390fd5b60005b84811015611ad357611acb87878784818110611a8a57611a8a6153a3565b9050602002016020810190611a9f9190614927565b868685818110611ab157611ab16153a3565b9050602002016020810190611ac6919061579d565b613c7c565b600101611a6c565b50505050505050565b83611ae6816137d0565b611b035760405163932d94f760e01b815260040160405180910390fd5b604080518082019091526001600160a01b038616815263ffffffff851660208201526000611b308261358f565b60008181526099602052604090209091506021908590611b4f906137a2565b611b5991906157c1565b1115611b78576040516301a1443960e31b815260040160405180910390fd5b6020808301516001600160a01b038916600090815260989092526040909120611baa9163ffffffff908116906137b816565b611bc757604051631fb1705560e21b815260040160405180910390fd5b60005b84811015611c6b57611be7868683818110610be157610be16153a3565b611c045760405163585cfb2f60e01b815260040160405180910390fd5b7f7ab260fe0af193db5f4986770d831bda4ea46099dc817e8b6716dcae8af8e88b83878784818110611c3857611c386153a3565b9050602002016020810190611c4d9190614927565b604051611c5b92919061544b565b60405180910390a1600101611bca565b5050505050505050565b6060600082516001600160401b03811115611c9257611c926147b0565b604051908082528060200260200182016040528015611cbb578160200160208202803683370190505b50905060005b835181101561097557611ced85858381518110611ce057611ce06153a3565b6020026020010151612bcc565b828281518110611cff57611cff6153a3565b6001600160401b0390921660209283029190910190910152600101611cc1565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611e1c57611d58826137d0565b611d75576040516348f5c3ed60e01b815260040160405180910390fd5b6040516336b87bd760e11b81526001600160a01b0383811660048301527f00000000000000000000000000000000000000000000000000000000000000001690636d70f7ae90602401602060405180830381865afa158015611ddb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611dff9190615381565b611e1c5760405163ccea9e6f60e01b815260040160405180910390fd5b61085c8282613d85565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015611e8a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611eae9190615381565b611ecb57604051631d77d47760e21b815260040160405180910390fd5b611ed6600019613765565b565b600061078283609a60006112c48661358f565b6001600160a01b03828116600081815260a2602090815260408083209486168084529482528083205493835260a38252808320948352939052918220546001600160401b0390911690600f81810b600160801b909204900b03825b81811015612018576001600160a01b03808716600090815260a3602090815260408083209389168352929052908120611f7f9083613f28565b6001600160a01b03888116600090815260a0602090815260408083208584528252808320938b16835292815290829020825160608101845290546001600160401b0381168252600160401b8104600f0b92820192909252600160c01b90910463ffffffff16918101829052919250431015611ffb575050612018565b6120098582602001516138b3565b94505050806001019050611f46565b506001600160a01b03808616600090815260a16020908152604080832093881683529290522082906120499061389f565b61205391906155d3565b95945050505050565b6066546002906004908116036120855760405163840a48d560e01b815260040160405180910390fd5b61209a6120956020840184614927565b6137d0565b806120b357506120b36120956040840160208501614927565b6120d0576040516348f5c3ed60e01b815260040160405180910390fd5b60005b6120e060408401846153d9565b90508110156123b257600060405180604001604052808560200160208101906121099190614927565b6001600160a01b0316815260200161212460408701876153d9565b85818110612134576121346153a3565b90506020020160208101906121499190615422565b63ffffffff168152509050612198816020015163ffffffff16609860008760200160208101906121799190614927565b6001600160a01b031681526020810191909152604001600020906137b8565b6121b557604051631fb1705560e21b815260040160405180910390fd5b609e60006121c66020870187614927565b6001600160a01b03166001600160a01b0316815260200190815260200160002060006121f18361358f565b815260208101919091526040016000205460ff16612222576040516325131d4f60e01b815260040160405180910390fd5b61225e61222e8261358f565b609c600061223f6020890189614927565b6001600160a01b03168152602081019190915260400160002090613f99565b5061229861226f6020860186614927565b609a600061227c8561358f565b8152602001908152602001600020613fa590919063ffffffff16565b506122a66020850185614927565b6001600160a01b03167fad34c3070be1dffbcaa499d000ba2b8d9848aefcac3059df245dd95c4ece14fe826040516122de919061543d565b60405180910390a26040805180820190915260008152602081016123227f0000000000000000000000000000000000000000000000000000000000000000436157d4565b63ffffffff169052609e600061233b6020880188614927565b6001600160a01b03166001600160a01b0316815260200190815260200160002060006123668461358f565b815260208082019290925260400160002082518154939092015163ffffffff166101000264ffffffff00199215159290921664ffffffffff1990931692909217179055506001016120d3565b506123c661038d6040840160208501614927565b6001600160a01b0316639d8e0c236123e16020850185614927565b6123ee60408601866153d9565b6040518463ffffffff1660e01b815260040161240c9392919061582b565b600060405180830381600087803b15801561242657600080fd5b505af1925050508015612437575060015b1561085c575050565b60606109af609a60006119638561358f565b61245a613fba565b611ed66000614014565b6001600160a01b0381166000908152609c6020526040812060609190612489906137a2565b90506000816001600160401b038111156124a5576124a56147b0565b6040519080825280602002602001820160405280156124ea57816020015b60408051808201909152600080825260208201528152602001906001900390816124c35790505b50905060005b82811015610975576001600160a01b0385166000908152609c6020526040902061251e9061091a90836137ac565b828281518110612530576125306153a3565b60209081029190910101526001016124f0565b6060600084516001600160401b03811115612560576125606147b0565b6040519080825280602002602001820160405280156125ab57816020015b604080516060810182526000808252602080830182905292820152825260001990920191018161257e5790505b50905060005b8551811015612603576125de8682815181106125cf576125cf6153a3565b6020026020010151868661074d565b8282815181106125f0576125f06153a3565b60209081029190910101526001016125b1565b50949350505050565b6060600083516001600160401b03811115612629576126296147b0565b604051908082528060200260200182016040528015612652578160200160208202803683370190505b50905060005b8451811015612603576001600160a01b038616600090815260a16020526040812086516126ca92879291899086908110612694576126946153a3565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002061406690919063ffffffff16565b8282815181106126dc576126dc6153a3565b6001600160401b0390921660209283029190910190910152600101612658565b6066546000906001908116036127255760405163840a48d560e01b815260040160405180910390fd5b61272e836137d0565b61274b576040516348f5c3ed60e01b815260040160405180910390fd5b60008060006127598661314d565b915091508161277b5760405163fa55fc8160e01b815260040160405180910390fd5b9150600090505b8351811015610cad5783818151811061279d5761279d6153a3565b602002602001015160400151518482815181106127bc576127bc6153a3565b60200260200101516020015151146127e7576040516343714afd60e01b815260040160405180910390fd5b60008482815181106127fb576127fb6153a3565b602090810291909101810151518082015181516001600160a01b031660009081526098909352604090922090925061283c9163ffffffff908116906137b816565b61285957604051631fb1705560e21b815260040160405180910390fd5b600061286587836138f2565b905060005b86848151811061287c5761287c6153a3565b60200260200101516020015151811015612bc15760008785815181106128a4576128a46153a3565b60200260200101516020015182815181106128c1576128c16153a3565b602002602001015190506128d8898261ffff613c7c565b6000806128e88b6107778861358f565b915091508060200151600f0b60001461291457604051630d8fcbe360e41b815260040160405180910390fd5b60006129228785848961407b565b905061296882600001518c8a8151811061293e5761293e6153a3565b602002602001015160400151878151811061295b5761295b6153a3565b60200260200101516140b3565b600f0b6020830181905260000361299257604051634606179360e11b815260040160405180910390fd5b60008260200151600f0b1215612acf578015612a5057612a166129b48861358f565b6001600160a01b03808f16600090815260a360209081526040808320938a16835292905220908154600160801b90819004600f0b6000818152600180860160205260409091209390935583546001600160801b03908116939091011602179055565b612a407f0000000000000000000000000000000000000000000000000000000000000000436157d4565b63ffffffff166040830152612b3d565b612a62836020015183602001516138b3565b6001600160401b031660208401528a518b9089908110612a8457612a846153a3565b6020026020010151604001518581518110612aa157612aa16153a3565b6020908102919091018101516001600160401b0316835260009083015263ffffffff43166040830152612b3d565b60008260200151600f0b1315612b3d57612af1836020015183602001516138b3565b6001600160401b039081166020850181905284519091161015612b2757604051636c9be0bf60e01b815260040160405180910390fd5b612b3189436157d4565b63ffffffff1660408301525b612b528c612b4a8961358f565b8686866139a2565b7f1487af5418c47ee5ea45ef4a93398668120890774a9e13487e61e9dc3baf76dd8c612b8061091a8a61358f565b86612b93866000015187602001516138b3565b8660400151604051612ba9959493929190615645565b60405180910390a150506001909201915061286a9050565b505050600101612782565b6001600160a01b03808316600090815260a16020908152604080832093851683529290529081206107829061389f565b82612c06816137d0565b612c235760405163932d94f760e01b815260040160405180910390fd5b836001600160a01b03167fa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c9437138484604051612c5e929190615850565b60405180910390a250505050565b606654600290600490811603612c955760405163840a48d560e01b815260040160405180910390fd5b82612c9f816137d0565b612cbc5760405163932d94f760e01b815260040160405180910390fd5b6040516336b87bd760e11b81526001600160a01b0385811660048301527f00000000000000000000000000000000000000000000000000000000000000001690636d70f7ae90602401602060405180830381865afa158015612d22573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d469190615381565b612d635760405163ccea9e6f60e01b815260040160405180910390fd5b60005b612d7360208501856153d9565b9050811015612f42576040805180820190915260009080612d976020880188614927565b6001600160a01b03168152602001868060200190612db591906153d9565b85818110612dc557612dc56153a3565b9050602002016020810190612dda9190615422565b63ffffffff90811690915260208083015183516001600160a01b0316600090815260989092526040909120929350612e179291908116906137b816565b612e3457604051631fb1705560e21b815260040160405180910390fd5b612e3e86826138f2565b15612e5c57604051636c6c6e2760e11b815260040160405180910390fd5b612e86612e688261358f565b6001600160a01b0388166000908152609c602052604090209061387e565b50612eb486609a6000612e988561358f565b815260200190815260200160002061388a90919063ffffffff16565b50856001600160a01b03167f43232edf9071753d2321e5fa7e018363ee248e5f2142e6c08edd3265bfb4895e82604051612eee919061543d565b60405180910390a26001600160a01b0386166000908152609e60205260408120600191612f1a8461358f565b81526020810191909152604001600020805460ff191691151591909117905550600101612d66565b50612f5361038d6020850185614927565b6001600160a01b031663adcf73f785612f6f60208701876153d9565b612f7c6040890189615696565b6040518663ffffffff1660e01b8152600401612f9c959493929190615864565b600060405180830381600087803b158015612fb657600080fd5b505af1158015611c6b573d6000803e3d6000fd5b60006109af609a6000612fdc8561358f565b81526020019081526020016000206137a2565b83612ff9816137d0565b6130165760405163932d94f760e01b815260040160405180910390fd5b6040805180820182526001600160a01b03871680825263ffffffff808816602080850182905260009384526098905293909120919261305692916137b816565b61307357604051631fb1705560e21b815260040160405180910390fd5b600061307e8261358f565b905060005b84811015611c6b576130c98686838181106130a0576130a06153a3565b90506020020160208101906130b59190614927565b600084815260996020526040902090613fa5565b6130e6576040516331bc342760e11b815260040160405180910390fd5b7f7b4b073d80dcac55a11177d8459ad9f664ceeb91f71f27167bb14f8152a7eeee8387878481811061311a5761311a6153a3565b905060200201602081019061312f9190614927565b60405161313d92919061544b565b60405180910390a1600101613083565b6001600160a01b0381166000908152609b602090815260408083208151608081018352905463ffffffff80821680845260ff600160201b8404161515958401869052650100000000008304821694840194909452600160481b9091041660608201819052849391929190158015906131cf5750826060015163ffffffff164310155b156131de575050604081015160015b9590945092505050565b6001600160a01b03811660009081526098602052604081206109af906137a2565b6001600160a01b0382166000908152609f602052604081206060919061323390826119638661358f565b949350505050565b600054610100900460ff161580801561325b5750600054600160ff909116105b806132755750303b158015613275575060005460ff166001145b6132dd5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff191660011790558015613300576000805461ff0019166101001790555b61330982613765565b61331283614014565b8015613359576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a15b505050565b81613368816137d0565b6133855760405163932d94f760e01b815260040160405180910390fd5b6001600160a01b03838116600090815260976020526040902080546001600160a01b0319169184169190911790557f2ae945c40c44dc0ec263f95609c3fdc6952e0aefa22d6374e44f2c997acedf85836133de81610fb0565b604080516001600160a01b03938416815292909116602083015201613350565b613406613fba565b6001600160a01b03811661346b5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016132d4565b61347481614014565b50565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156134d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134f991906158a8565b6001600160a01b0316336001600160a01b03161461352a5760405163794821ff60e01b815260040160405180910390fd5b606654801982198116146135515760405163c61dca5d60e01b815260040160405180910390fd5b606682905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200160405180910390a25050565b60008160000151826020015163ffffffff166040516020016135dc92919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b6040516020818303038152906040526109af906158c5565b604080518082018252600080825260208083018290528351606081018552828152808201839052808501839052845180860186526001600160a01b03898116855260a18452868520908816855290925293822092939281906136559061389f565b6001600160401b0390811682526001600160a01b03898116600081815260a260209081526040808320948c168084529482528083205486169682019690965291815260a082528481208b8252825284812092815291815290839020835160608101855290549283168152600160401b8304600f0b91810191909152600160c01b90910463ffffffff169181018290529192504310156136f857909250905061375d565b61370a816000015182602001516138b3565b6001600160401b0316815260208101516000600f9190910b12156137495761373a826020015182602001516138b3565b6001600160401b031660208301525b600060408201819052602082015290925090505b935093915050565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b60006109af825490565b600061078283836140cb565b60008181526001830160205260408120541515610782565b604051631beb2b9760e31b81526001600160a01b038281166004830152336024830152306044830152600080356001600160e01b0319166064840152917f00000000000000000000000000000000000000000000000000000000000000009091169063df595cb8906084016020604051808303816000875af115801561385a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109af9190615381565b600061078283836140f5565b6000610782836001600160a01b0384166140f5565b60006109af82670de0b6b3a7640000614144565b6000610782826001600160401b038516615618565b600061078283670de0b6b3a76400008461418a565b60006107828383670de0b6b3a764000061418a565b6001600160a01b0382166000908152609e602052604081208190816139168561358f565b8152602080820192909252604090810160002081518083019092525460ff8116151580835261010090910463ffffffff169282019290925291508061323357506020015163ffffffff1643109392505050565b6001600160a01b03811660009081526001830160205260408120541515610782565b60006107828383670de0b6b3a76400006001614274565b6020808301516001600160a01b03808816600090815260a284526040808220928816825291909352909120546001600160401b03908116911614613a6a57602082810180516001600160a01b03888116600081815260a286526040808220938a1680835293875290819020805467ffffffffffffffff19166001600160401b0395861617905593518451918252948101919091529216908201527facf9095feb3a370c9cf692421c69ef320d4db5c66e6a7d29c7694eb02364fc559060600160405180910390a15b6001600160a01b03808616600090815260a060209081526040808320888452825280832093871683529281529082902083518154928501519385015163ffffffff16600160c01b0263ffffffff60c01b196001600160801b038616600160401b026001600160c01b03199095166001600160401b03909316929092179390931716919091179055600f0b15613b4f576001600160a01b0385166000908152609f602090815260408083208784529091529020613b26908461388a565b506001600160a01b0385166000908152609d60205260409020613b49908561387e565b50610cad565b80516001600160401b0316600003610cad576001600160a01b0385166000908152609f602090815260408083208784529091529020613b8e9084613fa5565b506001600160a01b0385166000908152609f602090815260408083208784529091529020613bbb906137a2565b600003610cad576001600160a01b0385166000908152609d60205260409020613be49085613f99565b505050505050565b6001600160a01b03808416600090815260a160209081526040808320938616835292905220613c1c9043836142cf565b604080516001600160a01b038086168252841660208201526001600160401b038316918101919091527f1c6458079a41077d003c11faf9bf097e693bd67979e4e6500bac7b29db779b5c90606001613350565b60606000610782836142e3565b6001600160a01b03838116600090815260a360209081526040808320938616835292905290812054600f81810b600160801b909204900b035b600081118015613cc857508261ffff1682105b15610cad576001600160a01b03808616600090815260a3602090815260408083209388168352929052908120613cfd9061433f565b9050600080613d0d8884896135f4565b91509150806040015163ffffffff16431015613d2b57505050610cad565b613d3888848985856139a2565b6001600160a01b03808916600090815260a360209081526040808320938b16835292905220613d6690614393565b50613d70856158e9565b9450613d7b84615902565b9350505050613cb5565b6001600160a01b0382166000908152609b60209081526040918290208251608081018452905463ffffffff808216835260ff600160201b830416151593830193909352650100000000008104831693820193909352600160481b909204166060820181905215801590613e025750806060015163ffffffff164310155b15613e1c57604081015163ffffffff168152600160208201525b63ffffffff82166040820152613e527f0000000000000000000000000000000000000000000000000000000000000000436157d4565b63ffffffff90811660608381019182526001600160a01b0386166000818152609b602090815260409182902087518154838a0151858b01519851928a1664ffffffffff1990921691909117600160201b91151591909102176cffffffffffffffff0000000000191665010000000000978916979097026cffffffff000000000000000000191696909617600160481b968816968702179055815192835294871694820194909452928301919091527f4e85751d6331506c6c62335f207eb31f12a61e570f34f5c17640308785c6d4db9101613350565b600080613f4b613f3784614412565b8554613f469190600f0b615919565b614480565b8454909150600160801b9004600f90810b9082900b12613f7e57604051632d0483c560e21b815260040160405180910390fd5b600f0b60009081526001939093016020525050604090205490565b600061078283836144e9565b6000610782836001600160a01b0384166144e9565b6033546001600160a01b03163314611ed65760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016132d4565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60006107828383670de0b6b3a76400006145dc565b600061408e84609960006112c48961358f565b80156140975750815b801561205357505090516001600160401b031615159392505050565b60006107826001600160401b03808516908416615941565b60008260000182815481106140e2576140e26153a3565b9060005260206000200154905092915050565b600081815260018301602052604081205461413c575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556109af565b5060006109af565b815460009080156141825761416c8461415e6001846155c0565b600091825260209091200190565b54600160201b90046001600160e01b0316613233565b509092915050565b60008080600019858709858702925082811083820303915050806000036141c4578382816141ba576141ba61596e565b0492505050610782565b80841161420b5760405162461bcd60e51b81526020600482015260156024820152744d6174683a206d756c446976206f766572666c6f7760581b60448201526064016132d4565b60008486880960026001871981018816978890046003810283188082028403028082028403028082028403028082028403028082028403029081029092039091026000889003889004909101858311909403939093029303949094049190911702949350505050565b60008061428286868661418a565b9050600183600281111561429857614298615984565b1480156142b55750600084806142b0576142b061596e565b868809115b15612053576142c56001826157c1565b9695505050505050565b61335983836001600160401b038416614625565b60608160000180548060200260200160405190810160405280929190818152602001828054801561433357602002820191906000526020600020905b81548152602001906001019080831161431f575b50505050509050919050565b600061435a8254600f81810b600160801b909204900b131590565b1561437857604051631ed9509560e11b815260040160405180910390fd5b508054600f0b60009081526001909101602052604090205490565b60006143ae8254600f81810b600160801b909204900b131590565b156143cc57604051631ed9509560e11b815260040160405180910390fd5b508054600f0b6000818152600180840160205260408220805492905583546fffffffffffffffffffffffffffffffff191692016001600160801b03169190911790915590565b60006001600160ff1b0382111561447c5760405162461bcd60e51b815260206004820152602860248201527f53616665436173743a2076616c756520646f65736e27742066697420696e2061604482015267371034b73a191a9b60c11b60648201526084016132d4565b5090565b80600f81900b81146144e45760405162461bcd60e51b815260206004820152602760248201527f53616665436173743a2076616c756520646f65736e27742066697420696e20316044820152663238206269747360c81b60648201526084016132d4565b919050565b600081815260018301602052604081205480156145d257600061450d6001836155c0565b8554909150600090614521906001906155c0565b9050818114614586576000866000018281548110614541576145416153a3565b9060005260206000200154905080876000018481548110614564576145646153a3565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806145975761459761599a565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506109af565b60009150506109af565b8254600090816145ee8686838561472a565b9050801561461b576146058661415e6001846155c0565b54600160201b90046001600160e01b031661077d565b5091949350505050565b825480156146dc57600061463e8561415e6001856155c0565b60408051808201909152905463ffffffff808216808452600160201b9092046001600160e01b0316602084015291925090851610156146905760405163151b8e3f60e11b815260040160405180910390fd5b805163ffffffff8086169116036146da57826146b18661415e6001866155c0565b80546001600160e01b0392909216600160201b0263ffffffff9092169190911790555050505050565b505b506040805180820190915263ffffffff92831681526001600160e01b03918216602080830191825285546001810187556000968752952091519051909216600160201b029190921617910155565b60005b818310156109755760006147418484614780565b60008781526020902090915063ffffffff86169082015463ffffffff16111561476c5780925061477a565b6147778160016157c1565b93505b5061472d565b600061478f60028484186159b0565b610782908484166157c1565b6001600160a01b038116811461347457600080fd5b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b03811182821017156147e8576147e86147b0565b60405290565b604051601f8201601f191681016001600160401b0381118282101715614816576148166147b0565b604052919050565b803563ffffffff811681146144e457600080fd5b60006040828403121561484457600080fd5b604080519081016001600160401b0381118282101715614866576148666147b0565b60405290508082356148778161479b565b81526148856020840161481e565b60208201525092915050565b6000806000608084860312156148a657600080fd5b83356148b18161479b565b92506148c08560208601614832565b915060608401356148d08161479b565b809150509250925092565b81516001600160401b03168152602080830151600f0b9082015260408083015163ffffffff1690820152606081016109af565b60006020828403121561492057600080fd5b5035919050565b60006020828403121561493957600080fd5b81356107828161479b565b80516001600160a01b0316825260209081015163ffffffff16910152565b600081518084526020840193506020830160005b8281101561499f57614989868351614944565b6040959095019460209190910190600101614976565b5093949350505050565b6020815260006107826020830184614962565b6000604082840312156149ce57600080fd5b6107828383614832565b60008083601f8401126149ea57600080fd5b5081356001600160401b03811115614a0157600080fd5b6020830191508360208260051b850101111561194857600080fd5b600080600060408486031215614a3157600080fd5b8335614a3c8161479b565b925060208401356001600160401b03811115614a5757600080fd5b614a63868287016149d8565b9497909650939450505050565b60006001600160401b03821115614a8957614a896147b0565b5060051b60200190565b600082601f830112614aa457600080fd5b8135614ab7614ab282614a70565b6147ee565b8082825260208201915060208360051b860101925085831115614ad957600080fd5b602085015b83811015614aff578035614af18161479b565b835260209283019201614ade565b5095945050505050565b60008060008060a08587031215614b1f57600080fd5b614b298686614832565b935060408501356001600160401b03811115614b4457600080fd5b614b5087828801614a93565b93505060608501356001600160401b03811115614b6c57600080fd5b614b7887828801614a93565b925050614b876080860161481e565b905092959194509250565b600081518084526020840193506020830160005b8281101561499f578151865260209586019590910190600101614ba6565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015614c1d57603f19878603018452614c08858351614b92565b94506020938401939190910190600101614bec565b50929695505050505050565b60008060408385031215614c3c57600080fd5b8235614c478161479b565b915060208301356001600160401b03811115614c6257600080fd5b830160a08186031215614c7457600080fd5b809150509250929050565b60008060408385031215614c9257600080fd5b8235614c9d8161479b565b91506020830135614c748161479b565b600081518084526020840193506020830160005b8281101561499f57614cfa86835180516001600160401b03168252602080820151600f0b9083015260409081015163ffffffff16910152565b6060959095019460209190910190600101614cc1565b604081526000614d236040830185614962565b82810360208401526120538185614cad565b600081518084526020840193506020830160005b8281101561499f5781516001600160a01b0316865260209586019590910190600101614d49565b6020815260006107826020830184614d35565b60008060408385031215614d9657600080fd5b82356001600160401b03811115614dac57600080fd5b614db885828601614a93565b9250506020830135614c748161479b565b602080825282518282018190526000918401906040840190835b81811015614e0a5783516001600160401b0316835260209384019390920191600101614de3565b509095945050505050565b600080600080600060608688031215614e2d57600080fd5b8535614e388161479b565b945060208601356001600160401b03811115614e5357600080fd5b614e5f888289016149d8565b90955093505060408601356001600160401b03811115614e7e57600080fd5b614e8a888289016149d8565b969995985093965092949392505050565b60008060008060608587031215614eb157600080fd5b8435614ebc8161479b565b9350614eca6020860161481e565b925060408501356001600160401b03811115614ee557600080fd5b614ef1878288016149d8565b95989497509550505050565b60008060408385031215614f1057600080fd5b8235614f1b8161479b565b915060208301356001600160401b03811115614f3657600080fd5b614f4285828601614a93565b9150509250929050565b60008060408385031215614f5f57600080fd5b8235614f6a8161479b565b9150614f786020840161481e565b90509250929050565b600060208284031215614f9357600080fd5b813560ff8116811461078257600080fd5b60008060608385031215614fb757600080fd5b8235614fc28161479b565b9150614f788460208501614832565b600060608284031215614fe357600080fd5b50919050565b600060208284031215614ffb57600080fd5b81356001600160401b0381111561501157600080fd5b61323384828501614fd1565b60008060006080848603121561503257600080fd5b83356001600160401b0381111561504857600080fd5b61505486828701614a93565b9350506148c08560208601614832565b6020815260006107826020830184614cad565b60008060006060848603121561508c57600080fd5b83356150978161479b565b925060208401356001600160401b038111156150b257600080fd5b6150be86828701614a93565b9250506150cd6040850161481e565b90509250925092565b600080604083850312156150e957600080fd5b82356150f48161479b565b915060208301356001600160401b0381111561510f57600080fd5b8301601f8101851361512057600080fd5b803561512e614ab282614a70565b8082825260208201915060208360051b85010192508783111561515057600080fd5b602084015b8381101561527c5780356001600160401b0381111561517357600080fd5b85016080818b03601f1901121561518957600080fd5b6151916147c6565b61519e8b60208401614832565b815260608201356001600160401b038111156151b957600080fd5b6151c88c602083860101614a93565b60208301525060808201356001600160401b038111156151e757600080fd5b6020818401019250508a601f8301126151ff57600080fd5b813561520d614ab282614a70565b8082825260208201915060208360051b86010192508d83111561522f57600080fd5b6020850194505b828510156152665784356001600160401b038116811461525557600080fd5b825260209485019490910190615236565b6040840152505084525060209283019201615155565b50809450505050509250929050565b6000806000604084860312156152a057600080fd5b83356152ab8161479b565b925060208401356001600160401b038111156152c657600080fd5b8401601f810186136152d757600080fd5b80356001600160401b038111156152ed57600080fd5b8660208284010111156152ff57600080fd5b939660209190910195509293505050565b6000806040838503121561532357600080fd5b823561532e8161479b565b915060208301356001600160401b0381111561534957600080fd5b614f4285828601614fd1565b6000806040838503121561536857600080fd5b82356153738161479b565b946020939093013593505050565b60006020828403121561539357600080fd5b8151801515811461078257600080fd5b634e487b7160e01b600052603260045260246000fd5b60008235603e198336030181126153cf57600080fd5b9190910192915050565b6000808335601e198436030181126153f057600080fd5b8301803591506001600160401b0382111561540a57600080fd5b6020019150600581901b360382131561194857600080fd5b60006020828403121561543457600080fd5b6107828261481e565b604081016109af8284614944565b606081016154598285614944565b6001600160a01b039290921660409190910152919050565b6040815260006154846040830185614d35565b82810360208401526120538185614d35565b6000602082840312156154a857600080fd5b81516001600160401b038111156154be57600080fd5b8201601f810184136154cf57600080fd5b80516154dd614ab282614a70565b8082825260208201915060208360051b8501019250868311156154ff57600080fd5b602084015b8381101561559f5780516001600160401b0381111561552257600080fd5b8501603f8101891361553357600080fd5b6020810151615544614ab282614a70565b808282526020820191506020808460051b8601010192508b83111561556857600080fd5b6040840193505b8284101561558a57835182526020938401939091019061556f565b86525050602093840193919091019050615504565b509695505050505050565b634e487b7160e01b600052601160045260246000fd5b818103818111156109af576109af6155aa565b6001600160401b0382811682821603908111156109af576109af6155aa565b600081600f0b60016001607f1b0319810361560f5761560f6155aa565b60000392915050565b600f81810b9083900b0160016001607f1b03811360016001607f1b0319821217156109af576109af6155aa565b6001600160a01b038616815260c081016156626020830187614944565b6001600160a01b039490941660608201526001600160401b0392909216608083015263ffffffff1660a09091015292915050565b6000808335601e198436030181126156ad57600080fd5b8301803591506001600160401b038211156156c757600080fd5b60200191503681900382131561194857600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b0388168152600060c08201615724602084018a614944565b60c060608401528690528660e0830160005b888110156157665782356157498161479b565b6001600160a01b0316825260209283019290910190600101615736565b5083810360808501526157798188614b92565b91505082810360a084015261578f8185876156dc565b9a9950505050505050505050565b6000602082840312156157af57600080fd5b813561ffff8116811461078257600080fd5b808201808211156109af576109af6155aa565b63ffffffff81811683821601908111156109af576109af6155aa565b81835260208301925060008160005b8481101561499f5763ffffffff6158158361481e565b16865260209586019591909101906001016157ff565b6001600160a01b038416815260406020820181905260009061205390830184866157f0565b6020815260006132336020830184866156dc565b6001600160a01b038616815260606020820181905260009061588990830186886157f0565b828103604084015261589c8185876156dc565b98975050505050505050565b6000602082840312156158ba57600080fd5b81516107828161479b565b80516020808301519190811015614fe35760001960209190910360031b1b16919050565b6000600182016158fb576158fb6155aa565b5060010190565b600081615911576159116155aa565b506000190190565b8082018281126000831280158216821582161715615939576159396155aa565b505092915050565b600f82810b9082900b0360016001607f1b0319811260016001607f1b03821317156109af576109af6155aa565b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052603160045260246000fd5b6000826159cd57634e487b7160e01b600052601260045260246000fd5b50049056fea26469706673582212209fefd9029bfc14b1c009b24f56cf4ce0dc0b1da7c81fc9e5633e446a037f5e5964736f6c634300081b0033",
}

// ContractAllocationManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractAllocationManagerMetaData.ABI instead.
var ContractAllocationManagerABI = ContractAllocationManagerMetaData.ABI

// ContractAllocationManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractAllocationManagerMetaData.Bin instead.
var ContractAllocationManagerBin = ContractAllocationManagerMetaData.Bin

// DeployContractAllocationManager deploys a new Ethereum contract, binding an instance of ContractAllocationManager to it.
func DeployContractAllocationManager(auth *bind.TransactOpts, backend bind.ContractBackend, _delegation common.Address, _pauserRegistry common.Address, _permissionController common.Address, _DEALLOCATION_DELAY uint32, _ALLOCATION_CONFIGURATION_DELAY uint32) (common.Address, *types.Transaction, *ContractAllocationManager, error) {
	parsed, err := ContractAllocationManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractAllocationManagerBin), backend, _delegation, _pauserRegistry, _permissionController, _DEALLOCATION_DELAY, _ALLOCATION_CONFIGURATION_DELAY)
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

	GetOperatorSetCount(opts *bind.CallOpts, avs common.Address) (*big.Int, error)

	GetRegisteredSets(opts *bind.CallOpts, operator common.Address) ([]OperatorSet, error)

	GetStrategiesInOperatorSet(opts *bind.CallOpts, operatorSet OperatorSet) ([]common.Address, error)

	GetStrategyAllocations(opts *bind.CallOpts, operator common.Address, strategy common.Address) ([]OperatorSet, []IAllocationManagerTypesAllocation, error)

	IsMemberOfOperatorSet(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet) (bool, error)

	IsOperatorSet(opts *bind.CallOpts, operatorSet OperatorSet) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts, index uint8) (bool, error)

	Paused0(opts *bind.CallOpts) (*big.Int, error)

	PauserRegistry(opts *bind.CallOpts) (common.Address, error)

	PermissionController(opts *bind.CallOpts) (common.Address, error)
}

// ContractAllocationManagerTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractAllocationManagerTransacts interface {
	AddStrategiesToOperatorSet(opts *bind.TransactOpts, avs common.Address, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error)

	ClearDeallocationQueue(opts *bind.TransactOpts, operator common.Address, strategies []common.Address, numToClear []uint16) (*types.Transaction, error)

	CreateOperatorSets(opts *bind.TransactOpts, avs common.Address, params []IAllocationManagerTypesCreateSetParams) (*types.Transaction, error)

	DeregisterFromOperatorSets(opts *bind.TransactOpts, params IAllocationManagerTypesDeregisterParams) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialPausedStatus *big.Int) (*types.Transaction, error)

	ModifyAllocations(opts *bind.TransactOpts, operator common.Address, params []IAllocationManagerTypesAllocateParams) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	PauseAll(opts *bind.TransactOpts) (*types.Transaction, error)

	RegisterForOperatorSets(opts *bind.TransactOpts, operator common.Address, params IAllocationManagerTypesRegisterParams) (*types.Transaction, error)

	RemoveStrategiesFromOperatorSet(opts *bind.TransactOpts, avs common.Address, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetAVSRegistrar(opts *bind.TransactOpts, avs common.Address, registrar common.Address) (*types.Transaction, error)

	SetAllocationDelay(opts *bind.TransactOpts, operator common.Address, delay uint32) (*types.Transaction, error)

	SlashOperator(opts *bind.TransactOpts, avs common.Address, params IAllocationManagerTypesSlashingParams) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	UpdateAVSMetadataURI(opts *bind.TransactOpts, avs common.Address, metadataURI string) (*types.Transaction, error)
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

// GetOperatorSetCount is a free data retrieval call binding the contract method 0xba1a84e5.
//
// Solidity: function getOperatorSetCount(address avs) view returns(uint256)
func (_ContractAllocationManager *ContractAllocationManagerCaller) GetOperatorSetCount(opts *bind.CallOpts, avs common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "getOperatorSetCount", avs)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorSetCount is a free data retrieval call binding the contract method 0xba1a84e5.
//
// Solidity: function getOperatorSetCount(address avs) view returns(uint256)
func (_ContractAllocationManager *ContractAllocationManagerSession) GetOperatorSetCount(avs common.Address) (*big.Int, error) {
	return _ContractAllocationManager.Contract.GetOperatorSetCount(&_ContractAllocationManager.CallOpts, avs)
}

// GetOperatorSetCount is a free data retrieval call binding the contract method 0xba1a84e5.
//
// Solidity: function getOperatorSetCount(address avs) view returns(uint256)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) GetOperatorSetCount(avs common.Address) (*big.Int, error) {
	return _ContractAllocationManager.Contract.GetOperatorSetCount(&_ContractAllocationManager.CallOpts, avs)
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

// IsMemberOfOperatorSet is a free data retrieval call binding the contract method 0x670d3ba2.
//
// Solidity: function isMemberOfOperatorSet(address operator, (address,uint32) operatorSet) view returns(bool)
func (_ContractAllocationManager *ContractAllocationManagerCaller) IsMemberOfOperatorSet(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet) (bool, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "isMemberOfOperatorSet", operator, operatorSet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMemberOfOperatorSet is a free data retrieval call binding the contract method 0x670d3ba2.
//
// Solidity: function isMemberOfOperatorSet(address operator, (address,uint32) operatorSet) view returns(bool)
func (_ContractAllocationManager *ContractAllocationManagerSession) IsMemberOfOperatorSet(operator common.Address, operatorSet OperatorSet) (bool, error) {
	return _ContractAllocationManager.Contract.IsMemberOfOperatorSet(&_ContractAllocationManager.CallOpts, operator, operatorSet)
}

// IsMemberOfOperatorSet is a free data retrieval call binding the contract method 0x670d3ba2.
//
// Solidity: function isMemberOfOperatorSet(address operator, (address,uint32) operatorSet) view returns(bool)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) IsMemberOfOperatorSet(operator common.Address, operatorSet OperatorSet) (bool, error) {
	return _ContractAllocationManager.Contract.IsMemberOfOperatorSet(&_ContractAllocationManager.CallOpts, operator, operatorSet)
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

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ContractAllocationManager *ContractAllocationManagerCaller) PermissionController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAllocationManager.contract.Call(opts, &out, "permissionController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ContractAllocationManager *ContractAllocationManagerSession) PermissionController() (common.Address, error) {
	return _ContractAllocationManager.Contract.PermissionController(&_ContractAllocationManager.CallOpts)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ContractAllocationManager *ContractAllocationManagerCallerSession) PermissionController() (common.Address, error) {
	return _ContractAllocationManager.Contract.PermissionController(&_ContractAllocationManager.CallOpts)
}

// AddStrategiesToOperatorSet is a paid mutator transaction binding the contract method 0x50feea20.
//
// Solidity: function addStrategiesToOperatorSet(address avs, uint32 operatorSetId, address[] strategies) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) AddStrategiesToOperatorSet(opts *bind.TransactOpts, avs common.Address, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "addStrategiesToOperatorSet", avs, operatorSetId, strategies)
}

// AddStrategiesToOperatorSet is a paid mutator transaction binding the contract method 0x50feea20.
//
// Solidity: function addStrategiesToOperatorSet(address avs, uint32 operatorSetId, address[] strategies) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) AddStrategiesToOperatorSet(avs common.Address, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.AddStrategiesToOperatorSet(&_ContractAllocationManager.TransactOpts, avs, operatorSetId, strategies)
}

// AddStrategiesToOperatorSet is a paid mutator transaction binding the contract method 0x50feea20.
//
// Solidity: function addStrategiesToOperatorSet(address avs, uint32 operatorSetId, address[] strategies) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) AddStrategiesToOperatorSet(avs common.Address, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.AddStrategiesToOperatorSet(&_ContractAllocationManager.TransactOpts, avs, operatorSetId, strategies)
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

// CreateOperatorSets is a paid mutator transaction binding the contract method 0x261f84e0.
//
// Solidity: function createOperatorSets(address avs, (uint32,address[])[] params) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) CreateOperatorSets(opts *bind.TransactOpts, avs common.Address, params []IAllocationManagerTypesCreateSetParams) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "createOperatorSets", avs, params)
}

// CreateOperatorSets is a paid mutator transaction binding the contract method 0x261f84e0.
//
// Solidity: function createOperatorSets(address avs, (uint32,address[])[] params) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) CreateOperatorSets(avs common.Address, params []IAllocationManagerTypesCreateSetParams) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.CreateOperatorSets(&_ContractAllocationManager.TransactOpts, avs, params)
}

// CreateOperatorSets is a paid mutator transaction binding the contract method 0x261f84e0.
//
// Solidity: function createOperatorSets(address avs, (uint32,address[])[] params) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) CreateOperatorSets(avs common.Address, params []IAllocationManagerTypesCreateSetParams) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.CreateOperatorSets(&_ContractAllocationManager.TransactOpts, avs, params)
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

// ModifyAllocations is a paid mutator transaction binding the contract method 0x952899ee.
//
// Solidity: function modifyAllocations(address operator, ((address,uint32),address[],uint64[])[] params) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) ModifyAllocations(opts *bind.TransactOpts, operator common.Address, params []IAllocationManagerTypesAllocateParams) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "modifyAllocations", operator, params)
}

// ModifyAllocations is a paid mutator transaction binding the contract method 0x952899ee.
//
// Solidity: function modifyAllocations(address operator, ((address,uint32),address[],uint64[])[] params) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) ModifyAllocations(operator common.Address, params []IAllocationManagerTypesAllocateParams) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.ModifyAllocations(&_ContractAllocationManager.TransactOpts, operator, params)
}

// ModifyAllocations is a paid mutator transaction binding the contract method 0x952899ee.
//
// Solidity: function modifyAllocations(address operator, ((address,uint32),address[],uint64[])[] params) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) ModifyAllocations(operator common.Address, params []IAllocationManagerTypesAllocateParams) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.ModifyAllocations(&_ContractAllocationManager.TransactOpts, operator, params)
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

// RegisterForOperatorSets is a paid mutator transaction binding the contract method 0xadc2e3d9.
//
// Solidity: function registerForOperatorSets(address operator, (address,uint32[],bytes) params) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) RegisterForOperatorSets(opts *bind.TransactOpts, operator common.Address, params IAllocationManagerTypesRegisterParams) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "registerForOperatorSets", operator, params)
}

// RegisterForOperatorSets is a paid mutator transaction binding the contract method 0xadc2e3d9.
//
// Solidity: function registerForOperatorSets(address operator, (address,uint32[],bytes) params) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) RegisterForOperatorSets(operator common.Address, params IAllocationManagerTypesRegisterParams) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.RegisterForOperatorSets(&_ContractAllocationManager.TransactOpts, operator, params)
}

// RegisterForOperatorSets is a paid mutator transaction binding the contract method 0xadc2e3d9.
//
// Solidity: function registerForOperatorSets(address operator, (address,uint32[],bytes) params) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) RegisterForOperatorSets(operator common.Address, params IAllocationManagerTypesRegisterParams) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.RegisterForOperatorSets(&_ContractAllocationManager.TransactOpts, operator, params)
}

// RemoveStrategiesFromOperatorSet is a paid mutator transaction binding the contract method 0xb66bd989.
//
// Solidity: function removeStrategiesFromOperatorSet(address avs, uint32 operatorSetId, address[] strategies) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) RemoveStrategiesFromOperatorSet(opts *bind.TransactOpts, avs common.Address, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "removeStrategiesFromOperatorSet", avs, operatorSetId, strategies)
}

// RemoveStrategiesFromOperatorSet is a paid mutator transaction binding the contract method 0xb66bd989.
//
// Solidity: function removeStrategiesFromOperatorSet(address avs, uint32 operatorSetId, address[] strategies) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) RemoveStrategiesFromOperatorSet(avs common.Address, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.RemoveStrategiesFromOperatorSet(&_ContractAllocationManager.TransactOpts, avs, operatorSetId, strategies)
}

// RemoveStrategiesFromOperatorSet is a paid mutator transaction binding the contract method 0xb66bd989.
//
// Solidity: function removeStrategiesFromOperatorSet(address avs, uint32 operatorSetId, address[] strategies) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) RemoveStrategiesFromOperatorSet(avs common.Address, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.RemoveStrategiesFromOperatorSet(&_ContractAllocationManager.TransactOpts, avs, operatorSetId, strategies)
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

// SetAVSRegistrar is a paid mutator transaction binding the contract method 0xd3d96ff4.
//
// Solidity: function setAVSRegistrar(address avs, address registrar) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) SetAVSRegistrar(opts *bind.TransactOpts, avs common.Address, registrar common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "setAVSRegistrar", avs, registrar)
}

// SetAVSRegistrar is a paid mutator transaction binding the contract method 0xd3d96ff4.
//
// Solidity: function setAVSRegistrar(address avs, address registrar) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) SetAVSRegistrar(avs common.Address, registrar common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.SetAVSRegistrar(&_ContractAllocationManager.TransactOpts, avs, registrar)
}

// SetAVSRegistrar is a paid mutator transaction binding the contract method 0xd3d96ff4.
//
// Solidity: function setAVSRegistrar(address avs, address registrar) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) SetAVSRegistrar(avs common.Address, registrar common.Address) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.SetAVSRegistrar(&_ContractAllocationManager.TransactOpts, avs, registrar)
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

// SlashOperator is a paid mutator transaction binding the contract method 0x36352057.
//
// Solidity: function slashOperator(address avs, (address,uint32,address[],uint256[],string) params) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) SlashOperator(opts *bind.TransactOpts, avs common.Address, params IAllocationManagerTypesSlashingParams) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "slashOperator", avs, params)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x36352057.
//
// Solidity: function slashOperator(address avs, (address,uint32,address[],uint256[],string) params) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) SlashOperator(avs common.Address, params IAllocationManagerTypesSlashingParams) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.SlashOperator(&_ContractAllocationManager.TransactOpts, avs, params)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x36352057.
//
// Solidity: function slashOperator(address avs, (address,uint32,address[],uint256[],string) params) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) SlashOperator(avs common.Address, params IAllocationManagerTypesSlashingParams) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.SlashOperator(&_ContractAllocationManager.TransactOpts, avs, params)
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

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa9821821.
//
// Solidity: function updateAVSMetadataURI(address avs, string metadataURI) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, avs common.Address, metadataURI string) (*types.Transaction, error) {
	return _ContractAllocationManager.contract.Transact(opts, "updateAVSMetadataURI", avs, metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa9821821.
//
// Solidity: function updateAVSMetadataURI(address avs, string metadataURI) returns()
func (_ContractAllocationManager *ContractAllocationManagerSession) UpdateAVSMetadataURI(avs common.Address, metadataURI string) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.UpdateAVSMetadataURI(&_ContractAllocationManager.TransactOpts, avs, metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa9821821.
//
// Solidity: function updateAVSMetadataURI(address avs, string metadataURI) returns()
func (_ContractAllocationManager *ContractAllocationManagerTransactorSession) UpdateAVSMetadataURI(avs common.Address, metadataURI string) (*types.Transaction, error) {
	return _ContractAllocationManager.Contract.UpdateAVSMetadataURI(&_ContractAllocationManager.TransactOpts, avs, metadataURI)
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
