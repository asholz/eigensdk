// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractDelegationManager

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

// IDelegationManagerTypesQueuedWithdrawalParams is an auto generated low-level Go binding around an user-defined struct.
type IDelegationManagerTypesQueuedWithdrawalParams struct {
	Strategies    []common.Address
	DepositShares []*big.Int
	Withdrawer    common.Address
}

// IDelegationManagerTypesWithdrawal is an auto generated low-level Go binding around an user-defined struct.
type IDelegationManagerTypesWithdrawal struct {
	Staker       common.Address
	DelegatedTo  common.Address
	Withdrawer   common.Address
	Nonce        *big.Int
	StartBlock   uint32
	Strategies   []common.Address
	ScaledShares []*big.Int
}

// ISignatureUtilsSignatureWithExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithExpiry struct {
	Signature []byte
	Expiry    *big.Int
}

// ContractDelegationManagerMetaData contains all meta data concerning the ContractDelegationManager contract.
var ContractDelegationManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"},{\"name\":\"_eigenPodManager\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"},{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"},{\"name\":\"_MIN_WITHDRAWAL_DELAY\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DELEGATION_APPROVAL_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beaconChainETHStrategy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burnOperatorShares\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"prevMaxMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newMaxMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateDelegationApprovalDigestHash\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateWithdrawalRoot\",\"inputs\":[{\"name\":\"withdrawal\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManagerTypes.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"completeQueuedWithdrawal\",\"inputs\":[{\"name\":\"withdrawal\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManagerTypes.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"contractIERC20[]\"},{\"name\":\"receiveAsTokens\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeQueuedWithdrawals\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[][]\",\"internalType\":\"contractIERC20[][]\"},{\"name\":\"receiveAsTokens\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"},{\"name\":\"numToComplete\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeQueuedWithdrawals\",\"inputs\":[{\"name\":\"withdrawals\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationManagerTypes.Withdrawal[]\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"tokens\",\"type\":\"address[][]\",\"internalType\":\"contractIERC20[][]\"},{\"name\":\"receiveAsTokens\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cumulativeWithdrawalsQueued\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"totalQueued\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseDelegatedShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"curDepositShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beaconChainSlashingFactorDecrease\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegateTo\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approverSignatureAndExpiry\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegatedTo\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationApprover\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationApproverSaltIsSpent\",\"inputs\":[{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"spent\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositScalingFactor\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eigenPodManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDepositedShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorShares\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorsShares\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQueuedWithdrawals\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"withdrawals\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationManagerTypes.Withdrawal[]\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"shares\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashableSharesInQueue\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWithdrawableShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"withdrawableShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"depositShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseDelegatedShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"prevDepositShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"addedShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isDelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minWithdrawalDelayBlocks\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"modifyOperatorDetails\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newDelegationApprover\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorShares\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingWithdrawals\",\"inputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"pending\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queueWithdrawals\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationManagerTypes.QueuedWithdrawalParams[]\",\"components\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"depositShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"queuedWithdrawals\",\"inputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"redelegate\",\"inputs\":[{\"name\":\"newOperator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newOperatorApproverSig\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"withdrawalRoots\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerAsOperator\",\"inputs\":[{\"name\":\"initDelegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allocationDelay\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"undelegate\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"withdrawalRoots\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorMetadataURI\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"DelegationApproverUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newDelegationApprover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositScalingFactorUpdated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"newDepositScalingFactor\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorMetadataURIUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSharesBurned\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSharesDecreased\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSharesIncreased\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlashingWithdrawalCompleted\",\"inputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlashingWithdrawalQueued\",\"inputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"withdrawal\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIDelegationManagerTypes.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"sharesToWithdraw\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakerDelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakerForceUndelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakerUndelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ActivelyDelegated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerCannotUndelegate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FullySlashed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPermissions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSnapshotOrdering\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotActivelyDelegated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyAllocationManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyEigenPodManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyStrategyManagerOrEigenPodManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorsCannotUndelegate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SaltSpent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalDelayNotElapsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalNotQueued\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawerNotCaller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawerNotStaker\",\"inputs\":[]}]",
	Bin: "0x61018060405234801561001157600080fd5b506040516162d33803806162d38339810160408190526100309161023b565b8186868684876001600160a01b03811661005d576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b0390811660805293841660a05291831660c05290911660e05263ffffffff166101005246610120526100946100b8565b610140526001600160a01b0316610160526100ad610164565b5050505050506102cb565b600061012051461461015c5750604080518082018252600a81526922b4b3b2b72630bcb2b960b11b60209182015281517f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a866818301527f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea81840152466060820152306080808301919091528351808303909101815260a0909101909252815191012090565b506101405190565b600054610100900460ff16156101d05760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff90811614610221576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b038116811461023857600080fd5b50565b60008060008060008060c0878903121561025457600080fd5b865161025f81610223565b602088015190965061027081610223565b604088015190955061028181610223565b606088015190945061029281610223565b60808801519093506102a381610223565b60a088015190925063ffffffff811681146102bd57600080fd5b809150509295509295509295565b60805160a05160c05160e05161010051610120516101405161016051615f096103ca600039600081816103fe015261351c0152600061415f0152600061409f015260008181610753015281816136b10152613b9a0152600081816107a301528181610d0201528181610eb9015281816117ae01528181611c4a0152818161255a01528181612a1b015261438501526000818161042501528181610e3e015281816117140152818161197c015281816132ff015261405d01526000818161035a01528181610e0c015281816118cb0152818161264a01526140370152600081816105a201528181610bf801528181610fdd01526128880152615f096000f3fe608060405234801561001057600080fd5b50600436106102bb5760003560e01c8063778e55f311610182578063bfae3fd2116100e9578063e4cc3f90116100a2578063f0e0e6761161007c578063f0e0e67614610824578063f2fde38b14610844578063f698da2514610857578063fabc1cbc1461085f57600080fd5b8063e4cc3f90146107eb578063ee74937f146107fe578063eea9064b1461081157600080fd5b8063bfae3fd214610736578063c448feb814610749578063c978f7ac1461077d578063ca8aa7c71461079e578063cd6dc687146107c5578063da8be864146107d857600080fd5b80639435bb431161013b5780639435bb431461061057806399f5371b14610623578063a1788484146106b2578063a33a3433146106d2578063b7f06ebe146106e5578063bb45fef21461070857600080fd5b8063778e55f31461055f57806378296ec51461058a578063886f11951461059d5780638da5cb5b146105c457806390041347146105d55780639104c319146105f557600080fd5b8063595c6a671161022657806360a0d1ce116101df57806360a0d1ce146104d457806365da1264146104e757806366d5ba93146105105780636d70f7ae146105315780636e17444814610544578063715018a61461055757600080fd5b8063595c6a671461045a578063597b36da146104625780635ac86ab7146104755780635c975abb146104985780635dd68579146104a05780635f48e667146104c157600080fd5b80633c651cf2116102785780633c651cf2146103945780633cdeb5e0146103a75780633e28391d146103d65780634657e26a146103f95780634665bcda1461042057806354b7c96c1461044757600080fd5b806304a4f979146102c05780630b9f487a146102fa5780630dd8dd021461030d578063136439dd1461032d5780632aa6d8881461034257806339b70e3814610355575b600080fd5b6102e77f14bde674c9f64b2ad00eaaee4a8bed1fabef35c7507e3c5b9cfc9436909a2dad81565b6040519081526020015b60405180910390f35b6102e7610308366004614d64565b610872565b61032061031b366004614e03565b6108fb565b6040516102f19190614e44565b61034061033b366004614e7c565b610be3565b005b610340610350366004614eea565b610cba565b61037c7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016102f1565b6103406103a2366004614f4c565b610e01565b61037c6103b5366004614f92565b6001600160a01b039081166000908152609960205260409020600101541690565b6103e96103e4366004614f92565b610f4c565b60405190151581526020016102f1565b61037c7f000000000000000000000000000000000000000000000000000000000000000081565b61037c7f000000000000000000000000000000000000000000000000000000000000000081565b610340610455366004614faf565b610f6c565b610340610fc8565b6102e761047036600461522f565b61107a565b6103e9610483366004615263565b606654600160ff9092169190911b9081161490565b6066546102e7565b6104b36104ae366004614f92565b6110aa565b6040516102f19291906153cc565b6103406104cf36600461543b565b61146b565b6103406104e23660046154c7565b611709565b61037c6104f5366004614f92565b609a602052600090815260409020546001600160a01b031681565b61052361051e366004614f92565b6118a2565b6040516102f1929190615509565b6103e961053f366004614f92565b611baf565b6102e7610552366004614faf565b611be9565b610340611cf8565b6102e761056d366004614faf565b609860209081526000928352604080842090915290825290205481565b61034061059836600461552e565b611d0a565b61037c7f000000000000000000000000000000000000000000000000000000000000000081565b6033546001600160a01b031661037c565b6105e86105e3366004615582565b611d92565b6040516102f191906155d1565b61037c73beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac081565b61034061061e3660046155e4565b611e6c565b610674610631366004614e7c565b60a460205260009081526040902080546001820154600283015460038401546004909401546001600160a01b039384169492841693909116919063ffffffff1685565b604080516001600160a01b03968716815294861660208601529290941691830191909152606082015263ffffffff909116608082015260a0016102f1565b6102e76106c0366004614f92565b609f6020526000908152604090205481565b6103206106e0366004615687565b611f23565b6103e96106f3366004614e7c565b609e6020526000908152604090205460ff1681565b6103e9610716366004615777565b609c60209081526000928352604080842090915290825290205460ff1681565b6102e7610744366004614faf565b611fc0565b60405163ffffffff7f00000000000000000000000000000000000000000000000000000000000000001681526020016102f1565b61079061078b366004615582565b611ffd565b6040516102f19291906157a3565b61037c7f000000000000000000000000000000000000000000000000000000000000000081565b6103406107d3366004615777565b612295565b6103206107e6366004614f92565b6123b6565b6103406107f93660046157c4565b6124f9565b61034061080c366004615848565b61254f565b61034061081f366004615687565b6126f4565b610837610832366004615899565b612757565b6040516102f1919061594c565b610340610852366004614f92565b6127fe565b6102e7612877565b61034061086d366004614e7c565b612886565b604080517f14bde674c9f64b2ad00eaaee4a8bed1fabef35c7507e3c5b9cfc9436909a2dad60208201526001600160a01b03808616928201929092528187166060820152908516608082015260a0810183905260c081018290526000906108f19060e0016040516020818303038152906040528051906020012061299f565b9695505050505050565b6066546060906001906002908116036109275760405163840a48d560e01b815260040160405180910390fd5b6000836001600160401b0381111561094157610941614fe8565b60405190808252806020026020018201604052801561096a578160200160208202803683370190505b50336000908152609a60205260408120549192506001600160a01b03909116905b85811015610bd8578686828181106109a5576109a561595f565b90506020028101906109b79190615975565b6109c5906020810190615995565b90508787838181106109d9576109d961595f565b90506020028101906109eb9190615975565b6109f59080615995565b905014610a15576040516343714afd60e01b815260040160405180910390fd5b33878783818110610a2857610a2861595f565b9050602002810190610a3a9190615975565b610a4b906060810190604001614f92565b6001600160a01b031614610a72576040516330c4716960e21b815260040160405180910390fd5b6000610ade33848a8a86818110610a8b57610a8b61595f565b9050602002810190610a9d9190615975565b610aa79080615995565b808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506129ce92505050565b9050610bb233848a8a86818110610af757610af761595f565b9050602002810190610b099190615975565b610b139080615995565b808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508e92508d9150889050818110610b5957610b5961595f565b9050602002810190610b6b9190615975565b610b79906020810190615995565b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250889250612b1c915050565b848381518110610bc457610bc461595f565b60209081029190910101525060010161098b565b509095945050505050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610c47573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c6b91906159de565b610c8857604051631d77d47760e21b815260040160405180910390fd5b6066548181168114610cad5760405163c61dca5d60e01b815260040160405180910390fd5b610cb68261310d565b5050565b610cc333610f4c565b15610ce157604051633bf2b50360e11b815260040160405180910390fd5b604051632b6241f360e11b815233600482015263ffffffff841660248201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906356c483e690604401600060405180830381600087803b158015610d4e57600080fd5b505af1158015610d62573d6000803e3d6000fd5b50505050610d70338561314a565b610d7a33336131ad565b6040516001600160a01b038516815233907fa453db612af59e5521d6ab9284dc3e2d06af286eb1b1b7b771fce4716c19f2c19060200160405180910390a2336001600160a01b03167f02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b67080908383604051610df39291906159fb565b60405180910390a250505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480610e605750336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016145b610e7d5760405163045206a560e21b815260040160405180910390fd5b6001600160a01b038481166000908152609a602052604080822054905163152667d960e31b8152908316600482018190528684166024830152927f0000000000000000000000000000000000000000000000000000000000000000169063a9333ec890604401602060405180830381865afa158015610f00573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f249190615a2a565b90506000610f338787846132b6565b9050610f4383888888888661339c565b50505050505050565b6001600160a01b039081166000908152609a602052604090205416151590565b81610f76816134dd565b610f935760405163932d94f760e01b815260040160405180910390fd5b610f9c83611baf565b610fb9576040516325ec6c1f60e01b815260040160405180910390fd5b610fc3838361314a565b505050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa15801561102c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061105091906159de565b61106d57604051631d77d47760e21b815260040160405180910390fd5b61107860001961310d565b565b60008160405160200161108d9190615a47565b604051602081830303815290604052805190602001209050919050565b6001600160a01b038116600090815260a36020526040812060609182916110d09061358b565b8051909150806001600160401b038111156110ed576110ed614fe8565b60405190808252806020026020018201604052801561112657816020015b611113614c12565b81526020019060019003908161110b5790505b509350806001600160401b0381111561114157611141614fe8565b60405190808252806020026020018201604052801561117457816020015b606081526020019060019003908161115f5790505b506001600160a01b038087166000908152609a60205260408120549295509116905b828110156114625760a460008583815181106111b4576111b461595f565b6020908102919091018101518252818101929092526040908101600020815160e08101835281546001600160a01b03908116825260018301548116828601526002830154168184015260038201546060820152600482015463ffffffff1660808201526005820180548451818702810187019095528085529194929360a086019390929083018282801561127157602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611253575b50505050508152602001600682018054806020026020016040519081016040528092919081815260200182805480156112c957602002820191906000526020600020905b8154815260200190600101908083116112b5575b5050505050815250508682815181106112e4576112e461595f565b60200260200101819052508581815181106113015761130161595f565b602002602001015160a00151516001600160401b0381111561132557611325614fe8565b60405190808252806020026020018201604052801561134e578160200160208202803683370190505b508582815181106113615761136161595f565b6020026020010181905250600061139688848985815181106113855761138561595f565b602002602001015160a001516129ce565b905060005b8783815181106113ad576113ad61595f565b602002602001015160a00151518110156114585761141a8884815181106113d6576113d661595f565b602002602001015160c0015182815181106113f3576113f361595f565b602002602001015183858151811061140d5761140d61595f565b6020026020010151613598565b87848151811061142c5761142c61595f565b602002602001015182815181106114455761144561595f565b602090810291909101015260010161139b565b5050600101611196565b50505050915091565b6066546002906004908116036114945760405163840a48d560e01b815260040160405180910390fd5b61149c6135a4565b33600090815260a360205260408120906114b5826135fd565b90508084116114c457836114c6565b805b93506000846001600160401b038111156114e2576114e2614fe8565b60405190808252806020026020018201604052801561151b57816020015b611508614c12565b8152602001906001900390816115005790505b50905060005b81518110156116715760a460006115388684613607565b81526020808201929092526040908101600020815160e08101835281546001600160a01b03908116825260018301548116828601526002830154168184015260038201546060820152600482015463ffffffff1660808201526005820180548451818702810187019095528085529194929360a08601939092908301828280156115eb57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116115cd575b505050505081526020016006820180548060200260200160405190810160405280929190818152602001828054801561164357602002820191906000526020600020905b81548152602001906001019080831161162f575b50505050508152505082828151811061165e5761165e61595f565b6020908102919091010152600101611521565b5060005b81518110156116f3576116eb8282815181106116935761169361595f565b60200260200101518b8b848181106116ad576116ad61595f565b90506020028101906116bf9190615995565b8b8b868181106116d1576116d161595f565b90506020020160208101906116e69190615a5a565b613613565b600101611675565b50505050611701600160c955565b505050505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461175257604051633213a66160e21b815260040160405180910390fd5b61175b83610f4c565b15610fc3576001600160a01b038381166000908152609a602052604080822054905163152667d960e31b81529083166004820181905273beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac06024830152927f0000000000000000000000000000000000000000000000000000000000000000169063a9333ec890604401602060405180830381865afa1580156117f5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118199190615a2a565b6001600160a01b038616600090815260a26020908152604080832073beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac0845282528083208151928301909152548152919250611880866118786001600160401b03808716908916613aa5565b849190613aba565b9050610f43848873beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac084613ae1565b6040516394f649dd60e01b81526001600160a01b038281166004830152606091829160009182917f000000000000000000000000000000000000000000000000000000000000000016906394f649dd90602401600060405180830381865afa158015611912573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261193a9190810190615ad5565b60405163fe243a1760e01b81526001600160a01b03888116600483015273beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac060248301529294509092506000917f0000000000000000000000000000000000000000000000000000000000000000169063fe243a1790604401602060405180830381865afa1580156119c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119e79190615b96565b9050806000036119fc57509094909350915050565b600083516001611a0c9190615bc5565b6001600160401b03811115611a2357611a23614fe8565b604051908082528060200260200182016040528015611a4c578160200160208202803683370190505b509050600084516001611a5f9190615bc5565b6001600160401b03811115611a7657611a76614fe8565b604051908082528060200260200182016040528015611a9f578160200160208202803683370190505b50905073beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac082865181518110611aca57611aca61595f565b60200260200101906001600160a01b031690816001600160a01b0316815250508281865181518110611afe57611afe61595f565b60200260200101818152505060005b8551811015611ba157858181518110611b2857611b2861595f565b6020026020010151838281518110611b4257611b4261595f565b60200260200101906001600160a01b031690816001600160a01b031681525050848181518110611b7457611b7461595f565b6020026020010151828281518110611b8e57611b8e61595f565b6020908102919091010152600101611b0d565b509097909650945050505050565b60006001600160a01b03821615801590611be357506001600160a01b038083166000818152609a6020526040902054909116145b92915050565b604080516001808252818301909252600091829190602080830190803683370190505090508281600081518110611c2257611c2261595f565b6001600160a01b03928316602091820292909201015260405163547afb8760e01b81526000917f0000000000000000000000000000000000000000000000000000000000000000169063547afb8790611c819088908690600401615bd8565b600060405180830381865afa158015611c9e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611cc69190810190615bfc565b600081518110611cd857611cd861595f565b60200260200101519050611cef8585836000613b5c565b95945050505050565b611d00613c11565b6110786000613c6b565b82611d14816134dd565b611d315760405163932d94f760e01b815260040160405180910390fd5b611d3a84611baf565b611d57576040516325ec6c1f60e01b815260040160405180910390fd5b836001600160a01b03167f02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b67080908484604051610df39291906159fb565b6060600082516001600160401b03811115611daf57611daf614fe8565b604051908082528060200260200182016040528015611dd8578160200160208202803683370190505b50905060005b8351811015611e64576001600160a01b03851660009081526098602052604081208551909190869084908110611e1657611e1661595f565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054828281518110611e5157611e5161595f565b6020908102919091010152600101611dde565b509392505050565b606654600290600490811603611e955760405163840a48d560e01b815260040160405180910390fd5b611e9d6135a4565b8560005b81811015611f1757611f0f898983818110611ebe57611ebe61595f565b9050602002810190611ed09190615c90565b611ed990615ca6565b888884818110611eeb57611eeb61595f565b9050602002810190611efd9190615995565b8888868181106116d1576116d161595f565b600101611ea1565b5050610f43600160c955565b6060611f2e33610f4c565b611f4b5760405163a5c7c44560e01b815260040160405180910390fd5b611f5433611baf565b15611f72576040516311ca333560e31b815260040160405180910390fd5b611f7b84611baf565b611f98576040516325ec6c1f60e01b815260040160405180910390fd5b611fa133613cbd565b9050611faf33858585613f27565b611fb933856131ad565b9392505050565b6001600160a01b03808316600090815260a260209081526040808320938516835292815282822083519182019093529154825290611fb990613fef565b60608082516001600160401b0381111561201957612019614fe8565b604051908082528060200260200182016040528015612042578160200160208202803683370190505b50915082516001600160401b0381111561205e5761205e614fe8565b604051908082528060200260200182016040528015612087578160200160208202803683370190505b506001600160a01b038086166000908152609a60205260408120549293509116906120b38683876129ce565b905060005b855181101561228a5760006120e58783815181106120d8576120d861595f565b602002602001015161400f565b9050806001600160a01b031663fe243a17898985815181106121095761210961595f565b60200260200101516040518363ffffffff1660e01b81526004016121439291906001600160a01b0392831681529116602082015260400190565b602060405180830381865afa158015612160573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121849190615b96565b8583815181106121965761219661595f565b602002602001018181525050600060a260008a6001600160a01b03166001600160a01b0316815260200190815260200160002060008985815181106121dd576121dd61595f565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060405180602001604052908160008201548152505090506122638684815181106122315761223161595f565b602002602001015185858151811061224b5761224b61595f565b602002602001015183613aba9092919063ffffffff16565b8784815181106122755761227561595f565b602090810291909101015250506001016120b8565b5050505b9250929050565b600054610100900460ff16158080156122b55750600054600160ff909116105b806122cf5750303b1580156122cf575060005460ff166001145b6123375760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff19166001179055801561235a576000805461ff0019166101001790555b6123638261310d565b61236c83613c6b565b8015610fc3576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050565b60606123c182610f4c565b6123de5760405163a5c7c44560e01b815260040160405180910390fd5b6123e782611baf565b15612405576040516311ca333560e31b815260040160405180910390fd5b6001600160a01b03821661242c576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b038083166000818152609a60205260409020549091169033148061245b575061245b816134dd565b8061248257506001600160a01b038181166000908152609960205260409020600101541633145b61249f57604051631e499a2360e11b815260040160405180910390fd5b336001600160a01b038416146124f057806001600160a01b0316836001600160a01b03167ff0eddf07e6ea14f388b47e1e94a0f464ecbd9eed4171130e0fc0e99fb4030a8a60405160405180910390a35b611fb983613cbd565b6066546002906004908116036125225760405163840a48d560e01b815260040160405180910390fd5b61252a6135a4565b61253e61253686615ca6565b858585613613565b612548600160c955565b5050505050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614612598576040516323d871a560e01b815260040160405180910390fd5b6001600160a01b0380851660009081526098602090815260408083209387168352929052908120546125d7906001600160401b03808616908516614082565b905060006125e786868686613b5c565b6125f19083615bc5565b90506126008660008785613ae1565b6001600160a01b03851673beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac01461170157604051633b9e9f0160e21b81526001600160a01b038681166004830152602482018390527f0000000000000000000000000000000000000000000000000000000000000000169063ee7a7c0490604401600060405180830381600087803b15801561268e57600080fd5b505af11580156126a2573d6000803e3d6000fd5b5050604080516001600160a01b038981168252602082018690528a1693507feff6aab2bc3f7c648896e1522eae71d6c22e3b0e218206b3f40af0e4d204716b92500160405180910390a2505050505050565b6126fd33610f4c565b1561271b57604051633bf2b50360e11b815260040160405180910390fd5b61272483611baf565b612741576040516325ec6c1f60e01b815260040160405180910390fd5b61274d33848484613f27565b610fc333846131ad565b6060600083516001600160401b0381111561277457612774614fe8565b6040519080825280602002602001820160405280156127a757816020015b60608152602001906001900390816127925790505b50905060005b8451811015611e64576127d98582815181106127cb576127cb61595f565b602002602001015185611d92565b8282815181106127eb576127eb61595f565b60209081029190910101526001016127ad565b612806613c11565b6001600160a01b03811661286b5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161232e565b61287481613c6b565b50565b600061288161409b565b905090565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156128e4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129089190615cb2565b6001600160a01b0316336001600160a01b0316146129395760405163794821ff60e01b815260040160405180910390fd5b606654801982198116146129605760405163c61dca5d60e01b815260040160405180910390fd5b606682905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020015b60405180910390a25050565b60006129a961409b565b60405161190160f01b602082015260228101919091526042810183905260620161108d565b6060600082516001600160401b038111156129eb576129eb614fe8565b604051908082528060200260200182016040528015612a14578160200160208202803683370190505b50905060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663547afb8786866040518363ffffffff1660e01b8152600401612a67929190615bd8565b600060405180830381865afa158015612a84573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612aac9190810190615bfc565b905060005b8451811015610bd857612af787868381518110612ad057612ad061595f565b6020026020010151848481518110612aea57612aea61595f565b60200260200101516132b6565b838281518110612b0957612b0961595f565b6020908102919091010152600101612ab1565b60006001600160a01b038616612b45576040516339b190bb60e11b815260040160405180910390fd5b8351600003612b675760405163796cc52560e01b815260040160405180910390fd5b600084516001600160401b03811115612b8257612b82614fe8565b604051908082528060200260200182016040528015612bab578160200160208202803683370190505b509050600085516001600160401b03811115612bc957612bc9614fe8565b604051908082528060200260200182016040528015612bf2578160200160208202803683370190505b50905060005b8651811015612f3b576000612c188883815181106120d8576120d861595f565b9050600060a260008c6001600160a01b03166001600160a01b0316815260200190815260200160002060008a8581518110612c5557612c5561595f565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000206040518060200160405290816000820154815250509050816001600160a01b031663fe243a178c8b8681518110612cb657612cb661595f565b60200260200101516040518363ffffffff1660e01b8152600401612cf09291906001600160a01b0392831681529116602082015260400190565b602060405180830381865afa158015612d0d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d319190615b96565b888481518110612d4357612d4361595f565b60200260200101511115612d6a5760405163f020e5b960e01b815260040160405180910390fd5b612d99888481518110612d7f57612d7f61595f565b602002602001015188858151811061224b5761224b61595f565b848481518110612dab57612dab61595f565b602002602001018181525050612df3848481518110612dcc57612dcc61595f565b6020026020010151888581518110612de657612de661595f565b6020026020010151614181565b858481518110612e0557612e0561595f565b60209081029190910101526001600160a01b038a1615612e9a57612e5c8a8a8581518110612e3557612e3561595f565b6020026020010151878681518110612e4f57612e4f61595f565b602002602001015161419d565b612e9a8a8c8b8681518110612e7357612e7361595f565b6020026020010151878781518110612e8d57612e8d61595f565b6020026020010151613ae1565b816001600160a01b031663724af4238c8b8681518110612ebc57612ebc61595f565b60200260200101518b8781518110612ed657612ed661595f565b60200260200101516040518463ffffffff1660e01b8152600401612efc93929190615ccf565b600060405180830381600087803b158015612f1657600080fd5b505af1158015612f2a573d6000803e3d6000fd5b505050505050806001019050612bf8565b506001600160a01b0388166000908152609f60205260408120805491829190612f6383615cf3565b919050555060006040518060e001604052808b6001600160a01b031681526020018a6001600160a01b031681526020018b6001600160a01b031681526020018381526020014363ffffffff1681526020018981526020018581525090506000612fcb8261107a565b6000818152609e602090815260408083208054600160ff19909116811790915560a4835292819020865181546001600160a01b03199081166001600160a01b039283161783558885015195830180548216968316969096179095559187015160028201805490951692169190911790925560608501516003830155608085015160048301805463ffffffff191663ffffffff90921691909117905560a0850151805193945085936130829260058501920190614c70565b5060c0820151805161309e916006840191602090910190614cd5565b5050506001600160a01b038b16600090815260a3602052604090206130c3908261422d565b507f26b2aae26516e8719ef50ea2f6831a2efbd4e37dccdf0f6936b27bc08e793e308183866040516130f793929190615d0c565b60405180910390a19a9950505050505050505050565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b6001600160a01b0382811660008181526099602090815260409182902060010180546001600160a01b0319169486169485179055905192835290917f773b54c04d756fcc5e678111f7d730de3be98192000799eee3d63716055a87c69101612993565b6066546000906001908116036131d65760405163840a48d560e01b815260040160405180910390fd5b6001600160a01b038381166000818152609a602052604080822080546001600160a01b0319169487169485179055517fc3ee9f2e5fda98e8066a1f745b2df9285f416fe98cf2559cd21484b3d87433049190a3600080613235856118a2565b9150915060006132468686856129ce565b905060005b8351811015610f43576132ae868886848151811061326b5761326b61595f565b602002602001015160008786815181106132875761328761595f565b60200260200101518787815181106132a1576132a161595f565b602002602001015161339c565b60010161324b565b600073beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeabf196001600160a01b0384160161338c5760405163a3d75e0960e01b81526001600160a01b0385811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063a3d75e0990602401602060405180830381865afa158015613348573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061336c9190615a2a565b90506133846001600160401b03848116908316613aa5565b915050611fb9565b506001600160401b031692915050565b806000036133bd57604051630a33bc6960e21b815260040160405180910390fd5b6001600160a01b03808616600090815260a2602090815260408083209388168352929052206133ee81858585614239565b6040805160208101909152815481527f8be932bac54561f27260f95463d9b8ab37e06b2842e5ee2404157cc13df6eb8f908790879061342c90613fef565b60405161343b93929190615ccf565b60405180910390a161344c86610f4c565b15610f43576001600160a01b03808816600090815260986020908152604080832093891683529290529081208054859290613488908490615bc5565b92505081905550866001600160a01b03167f1ec042c965e2edd7107b51188ee0f383e22e76179041ab3a9d18ff151405166c8787866040516134cc93929190615ccf565b60405180910390a250505050505050565b604051631beb2b9760e31b81526001600160a01b038281166004830152336024830152306044830152600080356001600160e01b0319166064840152917f00000000000000000000000000000000000000000000000000000000000000009091169063df595cb8906084016020604051808303816000875af1158015613567573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611be391906159de565b60606000611fb9836142b2565b6000611fb98383613aa5565b600260c954036135f65760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161232e565b600260c955565b6000611be3825490565b6000611fb9838361430e565b60a0840151518214613638576040516343714afd60e01b815260040160405180910390fd5b83604001516001600160a01b0316336001600160a01b03161461366e576040516316110d3560e21b815260040160405180910390fd5b60006136798561107a565b6000818152609e602052604090205490915060ff166136ab576040516387c9d21960e01b815260040160405180910390fd5b606060007f000000000000000000000000000000000000000000000000000000000000000087608001516136df9190615d37565b90504363ffffffff168163ffffffff16111561370e576040516378f67ae160e11b815260040160405180910390fd5b613726876000015188602001518960a0015184614338565b87516001600160a01b039081166000908152609a60205260408120548a5160a08c0151949650921693509161375d919084906129ce565b905060005b8860a00151518110156139bf57600061378a8a60a0015183815181106120d8576120d861595f565b905060006137c18b60c0015184815181106137a7576137a761595f565b602002602001015187858151811061140d5761140d61595f565b9050871561389757816001600160a01b0316632eae418c8c600001518d60a0015186815181106137f3576137f361595f565b60200260200101518d8d8881811061380d5761380d61595f565b90506020020160208101906138229190614f92565b60405160e085901b6001600160e01b03191681526001600160a01b0393841660048201529183166024830152909116604482015260648101849052608401600060405180830381600087803b15801561387a57600080fd5b505af115801561388e573d6000803e3d6000fd5b505050506139b5565b600080836001600160a01b031663c4623ea18e600001518f60a0015188815181106138c4576138c461595f565b60200260200101518f8f8a8181106138de576138de61595f565b90506020020160208101906138f39190614f92565b60405160e085901b6001600160e01b03191681526001600160a01b039384166004820152918316602483015290911660448201526064810186905260840160408051808303816000875af115801561394f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139739190615d53565b915091506139b2878e600001518f60a0015188815181106139965761399661595f565b602002602001015185858b8b815181106132a1576132a161595f565b50505b5050600101613762565b5087516001600160a01b0316600090815260a3602052604090206139e3908561446d565b50600084815260a46020526040812080546001600160a01b031990811682556001820180548216905560028201805490911690556003810182905560048101805463ffffffff1916905590613a3b6005830182614d10565b613a49600683016000614d10565b50506000848152609e602052604090819020805460ff19169055517f1f40400889274ed07b24845e5054a87a0cab969eb1277aafe61ae352e7c32a0090613a939086815260200190565b60405180910390a15050505050505050565b6000611fb98383670de0b6b3a7640000614479565b6000613ad982613ad3613acc87613fef565b8690613aa5565b90613aa5565b949350505050565b6001600160a01b03808516600090815260986020908152604080832093861683529290529081208054839290613b18908490615d77565b92505081905550836001600160a01b03167f6909600037b75d7b4733aedd815442b5ec018a827751c832aaff64eba5d6d2dd848484604051610df393929190615ccf565b6001600160a01b03808516600090815260a56020908152604080832093871683529290529081208190613b8e90614563565b90506000613bea613bbf7f000000000000000000000000000000000000000000000000000000000000000043615d8a565b6001600160a01b03808a16600090815260a560209081526040808320938c168352929052209061457e565b90506000613bf88284615d77565b9050613c0581878761459b565b98975050505050505050565b6033546001600160a01b031633146110785760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161232e565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b606654606090600190600290811603613ce95760405163840a48d560e01b815260040160405180910390fd5b6001600160a01b038084166000818152609a602052604080822080546001600160a01b0319811690915590519316928392917ffee30966a256b71e14bc0ebfc94315e28ef4a97a7131a9e2b7a310a73af4467691a3600080613d4a866118a2565b915091508151600003613d5f57505050613f21565b81516001600160401b03811115613d7857613d78614fe8565b604051908082528060200260200182016040528015613da1578160200160208202803683370190505b5094506000613db18785856129ce565b905060005b8351811015613f1b576040805160018082528183019092526000916020808301908036833750506040805160018082528183019092529293506000929150602080830190803683375050604080516001808252818301909252929350600092915060208083019080368337019050509050868481518110613e3957613e3961595f565b602002602001015183600081518110613e5457613e5461595f565b60200260200101906001600160a01b031690816001600160a01b031681525050858481518110613e8657613e8661595f565b602002602001015182600081518110613ea157613ea161595f565b602002602001018181525050848481518110613ebf57613ebf61595f565b602002602001015181600081518110613eda57613eda61595f565b602002602001018181525050613ef38b89858585612b1c565b8a8581518110613f0557613f0561595f565b6020908102919091010152505050600101613db6565b50505050505b50919050565b6001600160a01b038084166000908152609960205260409020600101541680613f505750613fe9565b6001600160a01b0381166000908152609c6020908152604080832085845290915290205460ff1615613f9557604051630d4c4c9160e21b815260040160405180910390fd5b6001600160a01b0381166000908152609c602090815260408083208584528252909120805460ff19166001179055830151612548908290613fdd908890889084908890610872565b855160208701516145ba565b50505050565b805160009015614000578151611be3565b670de0b6b3a764000092915050565b60006001600160a01b03821673beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac01461405b577f0000000000000000000000000000000000000000000000000000000000000000611be3565b7f000000000000000000000000000000000000000000000000000000000000000092915050565b6000614091848385600161460c565b613ad99085615d77565b60007f0000000000000000000000000000000000000000000000000000000000000000461461415c5750604080518082018252600a81526922b4b3b2b72630bcb2b960b11b60209182015281517f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a866818301527f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea81840152466060820152306080808301919091528351808303909101815260a0909101909252815191012090565b507f000000000000000000000000000000000000000000000000000000000000000090565b60008160000361419357506000611be3565b611fb9838361465d565b6001600160a01b03821673beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac014610fc3576001600160a01b03808416600090815260a56020908152604080832093861683529290529081206141f190614563565b9050613fe9436142018484615bc5565b6001600160a01b03808816600090815260a560209081526040808320938a168352929052209190614672565b6000611fb9838361467d565b8260000361425a57614253670de0b6b3a76400008261465d565b8455613fe9565b604080516020810190915284548152600090614277908584613aba565b905060006142858483615bc5565b905060006142a7846142a161429a888a615bc5565b859061465d565b9061465d565b875550505050505050565b60608160000180548060200260200160405190810160405280929190818152602001828054801561430257602002820191906000526020600020905b8154815260200190600101908083116142ee575b50505050509050919050565b60008260000182815481106143255761432561595f565b9060005260206000200154905092915050565b6060600083516001600160401b0381111561435557614355614fe8565b60405190808252806020026020018201604052801561437e578160200160208202803683370190505b50905060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166394d7d00c8787876040518463ffffffff1660e01b81526004016143d393929190615da6565b600060405180830381865afa1580156143f0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526144189190810190615bfc565b905060005b85518110156144615761443c88878381518110612ad057612ad061595f565b83828151811061444e5761444e61595f565b602090810291909101015260010161441d565b50909695505050505050565b6000611fb983836146cc565b60008080600019858709858702925082811083820303915050806000036144b3578382816144a9576144a9615de0565b0492505050611fb9565b8084116144fa5760405162461bcd60e51b81526020600482015260156024820152744d6174683a206d756c446976206f766572666c6f7760581b604482015260640161232e565b60008486880960026001871981018816978890046003810283188082028403028082028403028082028403028082028403028082028403029081029092039091026000889003889004909101858311909403939093029303949094049190911702949350505050565b600061456f82826147bf565b6001600160e01b031692915050565b600061458b838383614806565b6001600160e01b03169392505050565b6000613ad96145aa8385615df6565b85906001600160401b0316613aa5565b428110156145db57604051630819bdcd60e01b815260040160405180910390fd5b6145ef6001600160a01b0385168484614850565b613fe957604051638baa579f60e01b815260040160405180910390fd5b60008061461a868686614479565b9050600183600281111561463057614630615e15565b14801561464d57506000848061464857614648615de0565b868809115b15611cef576108f1600182615bc5565b6000611fb983670de0b6b3a764000084614479565b610fc38383836148a7565b60008181526001830160205260408120546146c457508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155611be3565b506000611be3565b600081815260018301602052604081205480156147b55760006146f0600183615d77565b855490915060009061470490600190615d77565b90508181146147695760008660000182815481106147245761472461595f565b90600052602060002001549050808760000184815481106147475761474761595f565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061477a5761477a615e2b565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050611be3565b6000915050611be3565b815460009080156147fe576147e7846147d9600184615d77565b600091825260209091200190565b5464010000000090046001600160e01b0316613ad9565b509092915050565b825460009081614818868683856149af565b905080156148465761482f866147d9600184615d77565b5464010000000090046001600160e01b03166108f1565b5091949350505050565b600080600061485f8585614a05565b9092509050600081600481111561487857614878615e15565b1480156148965750856001600160a01b0316826001600160a01b0316145b806108f157506108f1868686614a47565b825480156149605760006148c0856147d9600185615d77565b60408051808201909152905463ffffffff8082168084526401000000009092046001600160e01b0316602084015291925090851610156149135760405163151b8e3f60e11b815260040160405180910390fd5b805163ffffffff80861691160361495e5782614934866147d9600186615d77565b80546001600160e01b03929092166401000000000263ffffffff9092169190911790555050505050565b505b506040805180820190915263ffffffff92831681526001600160e01b03918216602080830191825285546001810187556000968752952091519051909216640100000000029190921617910155565b60005b81831015611e645760006149c68484614b33565b60008781526020902090915063ffffffff86169082015463ffffffff1611156149f1578092506149ff565b6149fc816001615bc5565b93505b506149b2565b6000808251604103614a3b5760208301516040840151606085015160001a614a2f87828585614b4e565b9450945050505061228e565b5060009050600261228e565b6000806000856001600160a01b0316631626ba7e60e01b8686604051602401614a71929190615e65565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092529051614aaf9190615e9f565b600060405180830381855afa9150503d8060008114614aea576040519150601f19603f3d011682016040523d82523d6000602084013e614aef565b606091505b5091509150818015614b0357506020815110155b80156108f157508051630b135d3f60e11b90614b289083016020908101908401615b96565b149695505050505050565b6000614b426002848418615eb1565b611fb990848416615bc5565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115614b855750600090506003614c09565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015614bd9573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116614c0257600060019250925050614c09565b9150600090505b94509492505050565b6040518060e0016040528060006001600160a01b0316815260200160006001600160a01b0316815260200160006001600160a01b0316815260200160008152602001600063ffffffff16815260200160608152602001606081525090565b828054828255906000526020600020908101928215614cc5579160200282015b82811115614cc557825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190614c90565b50614cd1929150614d2a565b5090565b828054828255906000526020600020908101928215614cc5579160200282015b82811115614cc5578251825591602001919060010190614cf5565b508054600082559060005260206000209081019061287491905b5b80821115614cd15760008155600101614d2b565b6001600160a01b038116811461287457600080fd5b8035614d5f81614d3f565b919050565b600080600080600060a08688031215614d7c57600080fd5b8535614d8781614d3f565b94506020860135614d9781614d3f565b93506040860135614da781614d3f565b94979396509394606081013594506080013592915050565b60008083601f840112614dd157600080fd5b5081356001600160401b03811115614de857600080fd5b6020830191508360208260051b850101111561228e57600080fd5b60008060208385031215614e1657600080fd5b82356001600160401b03811115614e2c57600080fd5b614e3885828601614dbf565b90969095509350505050565b602080825282518282018190526000918401906040840190835b81811015610bd8578351835260209384019390920191600101614e5e565b600060208284031215614e8e57600080fd5b5035919050565b803563ffffffff81168114614d5f57600080fd5b60008083601f840112614ebb57600080fd5b5081356001600160401b03811115614ed257600080fd5b60208301915083602082850101111561228e57600080fd5b60008060008060608587031215614f0057600080fd5b8435614f0b81614d3f565b9350614f1960208601614e95565b925060408501356001600160401b03811115614f3457600080fd5b614f4087828801614ea9565b95989497509550505050565b60008060008060808587031215614f6257600080fd5b8435614f6d81614d3f565b93506020850135614f7d81614d3f565b93969395505050506040820135916060013590565b600060208284031215614fa457600080fd5b8135611fb981614d3f565b60008060408385031215614fc257600080fd5b8235614fcd81614d3f565b91506020830135614fdd81614d3f565b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b60405160e081016001600160401b038111828210171561502057615020614fe8565b60405290565b604080519081016001600160401b038111828210171561502057615020614fe8565b604051601f8201601f191681016001600160401b038111828210171561507057615070614fe8565b604052919050565b60006001600160401b0382111561509157615091614fe8565b5060051b60200190565b600082601f8301126150ac57600080fd5b81356150bf6150ba82615078565b615048565b8082825260208201915060208360051b8601019250858311156150e157600080fd5b602085015b838110156151075780356150f981614d3f565b8352602092830192016150e6565b5095945050505050565b600082601f83011261512257600080fd5b81356151306150ba82615078565b8082825260208201915060208360051b86010192508583111561515257600080fd5b602085015b83811015615107578035835260209283019201615157565b600060e0828403121561518157600080fd5b615189614ffe565b905061519482614d54565b81526151a260208301614d54565b60208201526151b360408301614d54565b6040820152606082810135908201526151ce60808301614e95565b608082015260a08201356001600160401b038111156151ec57600080fd5b6151f88482850161509b565b60a08301525060c08201356001600160401b0381111561521757600080fd5b61522384828501615111565b60c08301525092915050565b60006020828403121561524157600080fd5b81356001600160401b0381111561525757600080fd5b613ad98482850161516f565b60006020828403121561527557600080fd5b813560ff81168114611fb957600080fd5b600081518084526020840193506020830160005b828110156152c15781516001600160a01b031686526020958601959091019060010161529a565b5093949350505050565b600081518084526020840193506020830160005b828110156152c15781518652602095860195909101906001016152df565b80516001600160a01b03908116835260208083015182169084015260408083015190911690830152606080820151908301526080808201516000916153499085018263ffffffff169052565b5060a082015160e060a085015261536360e0850182615286565b905060c083015184820360c0860152611cef82826152cb565b600082825180855260208501945060208160051b8301016020850160005b8381101561446157601f198584030188526153b68383516152cb565b602098890198909350919091019060010161539a565b6000604082016040835280855180835260608501915060608160051b86010192506020870160005b8281101561542557605f198786030184526154108583516152fd565b945060209384019391909101906001016153f4565b505050508281036020840152611cef818561537c565b60008060008060006060868803121561545357600080fd5b85356001600160401b0381111561546957600080fd5b61547588828901614dbf565b90965094505060208601356001600160401b0381111561549457600080fd5b6154a088828901614dbf565b96999598509660400135949350505050565b6001600160401b038116811461287457600080fd5b6000806000606084860312156154dc57600080fd5b83356154e781614d3f565b92506020840135915060408401356154fe816154b2565b809150509250925092565b60408152600061551c6040830185615286565b8281036020840152611cef81856152cb565b60008060006040848603121561554357600080fd5b833561554e81614d3f565b925060208401356001600160401b0381111561556957600080fd5b61557586828701614ea9565b9497909650939450505050565b6000806040838503121561559557600080fd5b82356155a081614d3f565b915060208301356001600160401b038111156155bb57600080fd5b6155c78582860161509b565b9150509250929050565b602081526000611fb960208301846152cb565b600080600080600080606087890312156155fd57600080fd5b86356001600160401b0381111561561357600080fd5b61561f89828a01614dbf565b90975095505060208701356001600160401b0381111561563e57600080fd5b61564a89828a01614dbf565b90955093505060408701356001600160401b0381111561566957600080fd5b61567589828a01614dbf565b979a9699509497509295939492505050565b60008060006060848603121561569c57600080fd5b83356156a781614d3f565b925060208401356001600160401b038111156156c257600080fd5b8401604081870312156156d457600080fd5b6156dc615026565b81356001600160401b038111156156f257600080fd5b8201601f8101881361570357600080fd5b80356001600160401b0381111561571c5761571c614fe8565b61572f601f8201601f1916602001615048565b81815289602083850101111561574457600080fd5b81602084016020830137600060209282018301528352928301359282019290925293969395505050506040919091013590565b6000806040838503121561578a57600080fd5b823561579581614d3f565b946020939093013593505050565b60408152600061551c60408301856152cb565b801515811461287457600080fd5b600080600080606085870312156157da57600080fd5b84356001600160401b038111156157f057600080fd5b850160e0818803121561580257600080fd5b935060208501356001600160401b0381111561581d57600080fd5b61582987828801614dbf565b909450925050604085013561583d816157b6565b939692955090935050565b6000806000806080858703121561585e57600080fd5b843561586981614d3f565b9350602085013561587981614d3f565b92506040850135615889816154b2565b9150606085013561583d816154b2565b600080604083850312156158ac57600080fd5b82356001600160401b038111156158c257600080fd5b8301601f810185136158d357600080fd5b80356158e16150ba82615078565b8082825260208201915060208360051b85010192508783111561590357600080fd5b6020840193505b8284101561592e57833561591d81614d3f565b82526020938401939091019061590a565b945050505060208301356001600160401b038111156155bb57600080fd5b602081526000611fb9602083018461537c565b634e487b7160e01b600052603260045260246000fd5b60008235605e1983360301811261598b57600080fd5b9190910192915050565b6000808335601e198436030181126159ac57600080fd5b8301803591506001600160401b038211156159c657600080fd5b6020019150600581901b360382131561228e57600080fd5b6000602082840312156159f057600080fd5b8151611fb9816157b6565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b600060208284031215615a3c57600080fd5b8151611fb9816154b2565b602081526000611fb960208301846152fd565b600060208284031215615a6c57600080fd5b8135611fb9816157b6565b600082601f830112615a8857600080fd5b8151615a966150ba82615078565b8082825260208201915060208360051b860101925085831115615ab857600080fd5b602085015b83811015615107578051835260209283019201615abd565b60008060408385031215615ae857600080fd5b82516001600160401b03811115615afe57600080fd5b8301601f81018513615b0f57600080fd5b8051615b1d6150ba82615078565b8082825260208201915060208360051b850101925087831115615b3f57600080fd5b6020840193505b82841015615b6a578351615b5981614d3f565b825260209384019390910190615b46565b8095505050505060208301516001600160401b03811115615b8a57600080fd5b6155c785828601615a77565b600060208284031215615ba857600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b80820180821115611be357611be3615baf565b6001600160a01b0383168152604060208201819052600090613ad990830184615286565b600060208284031215615c0e57600080fd5b81516001600160401b03811115615c2457600080fd5b8201601f81018413615c3557600080fd5b8051615c436150ba82615078565b8082825260208201915060208360051b850101925086831115615c6557600080fd5b6020840193505b828410156108f1578351615c7f816154b2565b825260209384019390910190615c6c565b6000823560de1983360301811261598b57600080fd5b6000611be3368361516f565b600060208284031215615cc457600080fd5b8151611fb981614d3f565b6001600160a01b039384168152919092166020820152604081019190915260600190565b600060018201615d0557615d05615baf565b5060010190565b838152606060208201526000615d2560608301856152fd565b82810360408401526108f181856152cb565b63ffffffff8181168382160190811115611be357611be3615baf565b60008060408385031215615d6657600080fd5b505080516020909101519092909150565b81810381811115611be357611be3615baf565b63ffffffff8281168282160390811115611be357611be3615baf565b6001600160a01b0384168152606060208201819052600090615dca90830185615286565b905063ffffffff83166040830152949350505050565b634e487b7160e01b600052601260045260246000fd5b6001600160401b038281168282160390811115611be357611be3615baf565b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052603160045260246000fd5b60005b83811015615e5c578181015183820152602001615e44565b50506000910152565b8281526040602082015260008251806040840152615e8a816060850160208701615e41565b601f01601f1916919091016060019392505050565b6000825161598b818460208701615e41565b600082615ece57634e487b7160e01b600052601260045260246000fd5b50049056fea2646970667358221220a92baf3d17088462dc768d30893e5f9f3cbef19a8b9cc08b887f6374dc50d1af64736f6c634300081b0033",
}

// ContractDelegationManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractDelegationManagerMetaData.ABI instead.
var ContractDelegationManagerABI = ContractDelegationManagerMetaData.ABI

// ContractDelegationManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractDelegationManagerMetaData.Bin instead.
var ContractDelegationManagerBin = ContractDelegationManagerMetaData.Bin

// DeployContractDelegationManager deploys a new Ethereum contract, binding an instance of ContractDelegationManager to it.
func DeployContractDelegationManager(auth *bind.TransactOpts, backend bind.ContractBackend, _strategyManager common.Address, _eigenPodManager common.Address, _allocationManager common.Address, _pauserRegistry common.Address, _permissionController common.Address, _MIN_WITHDRAWAL_DELAY uint32) (common.Address, *types.Transaction, *ContractDelegationManager, error) {
	parsed, err := ContractDelegationManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractDelegationManagerBin), backend, _strategyManager, _eigenPodManager, _allocationManager, _pauserRegistry, _permissionController, _MIN_WITHDRAWAL_DELAY)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractDelegationManager{ContractDelegationManagerCaller: ContractDelegationManagerCaller{contract: contract}, ContractDelegationManagerTransactor: ContractDelegationManagerTransactor{contract: contract}, ContractDelegationManagerFilterer: ContractDelegationManagerFilterer{contract: contract}}, nil
}

// ContractDelegationManagerMethods is an auto generated interface around an Ethereum contract.
type ContractDelegationManagerMethods interface {
	ContractDelegationManagerCalls
	ContractDelegationManagerTransacts
	ContractDelegationManagerFilters
}

// ContractDelegationManagerCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractDelegationManagerCalls interface {
	DELEGATIONAPPROVALTYPEHASH(opts *bind.CallOpts) ([32]byte, error)

	AllocationManager(opts *bind.CallOpts) (common.Address, error)

	BeaconChainETHStrategy(opts *bind.CallOpts) (common.Address, error)

	CalculateDelegationApprovalDigestHash(opts *bind.CallOpts, staker common.Address, operator common.Address, approver common.Address, approverSalt [32]byte, expiry *big.Int) ([32]byte, error)

	CalculateWithdrawalRoot(opts *bind.CallOpts, withdrawal IDelegationManagerTypesWithdrawal) ([32]byte, error)

	CumulativeWithdrawalsQueued(opts *bind.CallOpts, staker common.Address) (*big.Int, error)

	DelegatedTo(opts *bind.CallOpts, staker common.Address) (common.Address, error)

	DelegationApprover(opts *bind.CallOpts, operator common.Address) (common.Address, error)

	DelegationApproverSaltIsSpent(opts *bind.CallOpts, delegationApprover common.Address, salt [32]byte) (bool, error)

	DepositScalingFactor(opts *bind.CallOpts, staker common.Address, strategy common.Address) (*big.Int, error)

	DomainSeparator(opts *bind.CallOpts) ([32]byte, error)

	EigenPodManager(opts *bind.CallOpts) (common.Address, error)

	GetDepositedShares(opts *bind.CallOpts, staker common.Address) ([]common.Address, []*big.Int, error)

	GetOperatorShares(opts *bind.CallOpts, operator common.Address, strategies []common.Address) ([]*big.Int, error)

	GetOperatorsShares(opts *bind.CallOpts, operators []common.Address, strategies []common.Address) ([][]*big.Int, error)

	GetQueuedWithdrawals(opts *bind.CallOpts, staker common.Address) (struct {
		Withdrawals []IDelegationManagerTypesWithdrawal
		Shares      [][]*big.Int
	}, error)

	GetSlashableSharesInQueue(opts *bind.CallOpts, operator common.Address, strategy common.Address) (*big.Int, error)

	GetWithdrawableShares(opts *bind.CallOpts, staker common.Address, strategies []common.Address) (struct {
		WithdrawableShares []*big.Int
		DepositShares      []*big.Int
	}, error)

	IsDelegated(opts *bind.CallOpts, staker common.Address) (bool, error)

	IsOperator(opts *bind.CallOpts, operator common.Address) (bool, error)

	MinWithdrawalDelayBlocks(opts *bind.CallOpts) (uint32, error)

	OperatorShares(opts *bind.CallOpts, operator common.Address, strategy common.Address) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts, index uint8) (bool, error)

	Paused0(opts *bind.CallOpts) (*big.Int, error)

	PauserRegistry(opts *bind.CallOpts) (common.Address, error)

	PendingWithdrawals(opts *bind.CallOpts, withdrawalRoot [32]byte) (bool, error)

	PermissionController(opts *bind.CallOpts) (common.Address, error)

	QueuedWithdrawals(opts *bind.CallOpts, withdrawalRoot [32]byte) (struct {
		Staker      common.Address
		DelegatedTo common.Address
		Withdrawer  common.Address
		Nonce       *big.Int
		StartBlock  uint32
	}, error)

	StrategyManager(opts *bind.CallOpts) (common.Address, error)
}

// ContractDelegationManagerTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractDelegationManagerTransacts interface {
	BurnOperatorShares(opts *bind.TransactOpts, operator common.Address, strategy common.Address, prevMaxMagnitude uint64, newMaxMagnitude uint64) (*types.Transaction, error)

	CompleteQueuedWithdrawal(opts *bind.TransactOpts, withdrawal IDelegationManagerTypesWithdrawal, tokens []common.Address, receiveAsTokens bool) (*types.Transaction, error)

	CompleteQueuedWithdrawals(opts *bind.TransactOpts, tokens [][]common.Address, receiveAsTokens []bool, numToComplete *big.Int) (*types.Transaction, error)

	CompleteQueuedWithdrawals0(opts *bind.TransactOpts, withdrawals []IDelegationManagerTypesWithdrawal, tokens [][]common.Address, receiveAsTokens []bool) (*types.Transaction, error)

	DecreaseDelegatedShares(opts *bind.TransactOpts, staker common.Address, curDepositShares *big.Int, beaconChainSlashingFactorDecrease uint64) (*types.Transaction, error)

	DelegateTo(opts *bind.TransactOpts, operator common.Address, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error)

	IncreaseDelegatedShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, prevDepositShares *big.Int, addedShares *big.Int) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialPausedStatus *big.Int) (*types.Transaction, error)

	ModifyOperatorDetails(opts *bind.TransactOpts, operator common.Address, newDelegationApprover common.Address) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	PauseAll(opts *bind.TransactOpts) (*types.Transaction, error)

	QueueWithdrawals(opts *bind.TransactOpts, params []IDelegationManagerTypesQueuedWithdrawalParams) (*types.Transaction, error)

	Redelegate(opts *bind.TransactOpts, newOperator common.Address, newOperatorApproverSig ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error)

	RegisterAsOperator(opts *bind.TransactOpts, initDelegationApprover common.Address, allocationDelay uint32, metadataURI string) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	Undelegate(opts *bind.TransactOpts, staker common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	UpdateOperatorMetadataURI(opts *bind.TransactOpts, operator common.Address, metadataURI string) (*types.Transaction, error)
}

// ContractDelegationManagerFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractDelegationManagerFilters interface {
	FilterDelegationApproverUpdated(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerDelegationApproverUpdatedIterator, error)
	WatchDelegationApproverUpdated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerDelegationApproverUpdated, operator []common.Address) (event.Subscription, error)
	ParseDelegationApproverUpdated(log types.Log) (*ContractDelegationManagerDelegationApproverUpdated, error)

	FilterDepositScalingFactorUpdated(opts *bind.FilterOpts) (*ContractDelegationManagerDepositScalingFactorUpdatedIterator, error)
	WatchDepositScalingFactorUpdated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerDepositScalingFactorUpdated) (event.Subscription, error)
	ParseDepositScalingFactorUpdated(log types.Log) (*ContractDelegationManagerDepositScalingFactorUpdated, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractDelegationManagerInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractDelegationManagerInitialized, error)

	FilterOperatorMetadataURIUpdated(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerOperatorMetadataURIUpdatedIterator, error)
	WatchOperatorMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOperatorMetadataURIUpdated, operator []common.Address) (event.Subscription, error)
	ParseOperatorMetadataURIUpdated(log types.Log) (*ContractDelegationManagerOperatorMetadataURIUpdated, error)

	FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerOperatorRegisteredIterator, error)
	WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOperatorRegistered, operator []common.Address) (event.Subscription, error)
	ParseOperatorRegistered(log types.Log) (*ContractDelegationManagerOperatorRegistered, error)

	FilterOperatorSharesBurned(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerOperatorSharesBurnedIterator, error)
	WatchOperatorSharesBurned(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOperatorSharesBurned, operator []common.Address) (event.Subscription, error)
	ParseOperatorSharesBurned(log types.Log) (*ContractDelegationManagerOperatorSharesBurned, error)

	FilterOperatorSharesDecreased(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerOperatorSharesDecreasedIterator, error)
	WatchOperatorSharesDecreased(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOperatorSharesDecreased, operator []common.Address) (event.Subscription, error)
	ParseOperatorSharesDecreased(log types.Log) (*ContractDelegationManagerOperatorSharesDecreased, error)

	FilterOperatorSharesIncreased(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerOperatorSharesIncreasedIterator, error)
	WatchOperatorSharesIncreased(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOperatorSharesIncreased, operator []common.Address) (event.Subscription, error)
	ParseOperatorSharesIncreased(log types.Log) (*ContractDelegationManagerOperatorSharesIncreased, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractDelegationManagerOwnershipTransferredIterator, error)
	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)
	ParseOwnershipTransferred(log types.Log) (*ContractDelegationManagerOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractDelegationManagerPausedIterator, error)
	WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerPaused, account []common.Address) (event.Subscription, error)
	ParsePaused(log types.Log) (*ContractDelegationManagerPaused, error)

	FilterSlashingWithdrawalCompleted(opts *bind.FilterOpts) (*ContractDelegationManagerSlashingWithdrawalCompletedIterator, error)
	WatchSlashingWithdrawalCompleted(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerSlashingWithdrawalCompleted) (event.Subscription, error)
	ParseSlashingWithdrawalCompleted(log types.Log) (*ContractDelegationManagerSlashingWithdrawalCompleted, error)

	FilterSlashingWithdrawalQueued(opts *bind.FilterOpts) (*ContractDelegationManagerSlashingWithdrawalQueuedIterator, error)
	WatchSlashingWithdrawalQueued(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerSlashingWithdrawalQueued) (event.Subscription, error)
	ParseSlashingWithdrawalQueued(log types.Log) (*ContractDelegationManagerSlashingWithdrawalQueued, error)

	FilterStakerDelegated(opts *bind.FilterOpts, staker []common.Address, operator []common.Address) (*ContractDelegationManagerStakerDelegatedIterator, error)
	WatchStakerDelegated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerStakerDelegated, staker []common.Address, operator []common.Address) (event.Subscription, error)
	ParseStakerDelegated(log types.Log) (*ContractDelegationManagerStakerDelegated, error)

	FilterStakerForceUndelegated(opts *bind.FilterOpts, staker []common.Address, operator []common.Address) (*ContractDelegationManagerStakerForceUndelegatedIterator, error)
	WatchStakerForceUndelegated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerStakerForceUndelegated, staker []common.Address, operator []common.Address) (event.Subscription, error)
	ParseStakerForceUndelegated(log types.Log) (*ContractDelegationManagerStakerForceUndelegated, error)

	FilterStakerUndelegated(opts *bind.FilterOpts, staker []common.Address, operator []common.Address) (*ContractDelegationManagerStakerUndelegatedIterator, error)
	WatchStakerUndelegated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerStakerUndelegated, staker []common.Address, operator []common.Address) (event.Subscription, error)
	ParseStakerUndelegated(log types.Log) (*ContractDelegationManagerStakerUndelegated, error)

	FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractDelegationManagerUnpausedIterator, error)
	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerUnpaused, account []common.Address) (event.Subscription, error)
	ParseUnpaused(log types.Log) (*ContractDelegationManagerUnpaused, error)
}

// ContractDelegationManager is an auto generated Go binding around an Ethereum contract.
type ContractDelegationManager struct {
	ContractDelegationManagerCaller     // Read-only binding to the contract
	ContractDelegationManagerTransactor // Write-only binding to the contract
	ContractDelegationManagerFilterer   // Log filterer for contract events
}

// ContractDelegationManager implements the ContractDelegationManagerMethods interface.
var _ ContractDelegationManagerMethods = (*ContractDelegationManager)(nil)

// ContractDelegationManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractDelegationManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractDelegationManagerCaller implements the ContractDelegationManagerCalls interface.
var _ ContractDelegationManagerCalls = (*ContractDelegationManagerCaller)(nil)

// ContractDelegationManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractDelegationManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractDelegationManagerTransactor implements the ContractDelegationManagerTransacts interface.
var _ ContractDelegationManagerTransacts = (*ContractDelegationManagerTransactor)(nil)

// ContractDelegationManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractDelegationManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractDelegationManagerFilterer implements the ContractDelegationManagerFilters interface.
var _ ContractDelegationManagerFilters = (*ContractDelegationManagerFilterer)(nil)

// ContractDelegationManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractDelegationManagerSession struct {
	Contract     *ContractDelegationManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ContractDelegationManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractDelegationManagerCallerSession struct {
	Contract *ContractDelegationManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// ContractDelegationManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractDelegationManagerTransactorSession struct {
	Contract     *ContractDelegationManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ContractDelegationManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractDelegationManagerRaw struct {
	Contract *ContractDelegationManager // Generic contract binding to access the raw methods on
}

// ContractDelegationManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractDelegationManagerCallerRaw struct {
	Contract *ContractDelegationManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractDelegationManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractDelegationManagerTransactorRaw struct {
	Contract *ContractDelegationManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractDelegationManager creates a new instance of ContractDelegationManager, bound to a specific deployed contract.
func NewContractDelegationManager(address common.Address, backend bind.ContractBackend) (*ContractDelegationManager, error) {
	contract, err := bindContractDelegationManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManager{ContractDelegationManagerCaller: ContractDelegationManagerCaller{contract: contract}, ContractDelegationManagerTransactor: ContractDelegationManagerTransactor{contract: contract}, ContractDelegationManagerFilterer: ContractDelegationManagerFilterer{contract: contract}}, nil
}

// NewContractDelegationManagerCaller creates a new read-only instance of ContractDelegationManager, bound to a specific deployed contract.
func NewContractDelegationManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractDelegationManagerCaller, error) {
	contract, err := bindContractDelegationManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerCaller{contract: contract}, nil
}

// NewContractDelegationManagerTransactor creates a new write-only instance of ContractDelegationManager, bound to a specific deployed contract.
func NewContractDelegationManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractDelegationManagerTransactor, error) {
	contract, err := bindContractDelegationManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerTransactor{contract: contract}, nil
}

// NewContractDelegationManagerFilterer creates a new log filterer instance of ContractDelegationManager, bound to a specific deployed contract.
func NewContractDelegationManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractDelegationManagerFilterer, error) {
	contract, err := bindContractDelegationManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerFilterer{contract: contract}, nil
}

// bindContractDelegationManager binds a generic wrapper to an already deployed contract.
func bindContractDelegationManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractDelegationManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractDelegationManager *ContractDelegationManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractDelegationManager.Contract.ContractDelegationManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractDelegationManager *ContractDelegationManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.ContractDelegationManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractDelegationManager *ContractDelegationManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.ContractDelegationManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractDelegationManager *ContractDelegationManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractDelegationManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractDelegationManager *ContractDelegationManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractDelegationManager *ContractDelegationManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.contract.Transact(opts, method, params...)
}

// DELEGATIONAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0x04a4f979.
//
// Solidity: function DELEGATION_APPROVAL_TYPEHASH() view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCaller) DELEGATIONAPPROVALTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "DELEGATION_APPROVAL_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATIONAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0x04a4f979.
//
// Solidity: function DELEGATION_APPROVAL_TYPEHASH() view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerSession) DELEGATIONAPPROVALTYPEHASH() ([32]byte, error) {
	return _ContractDelegationManager.Contract.DELEGATIONAPPROVALTYPEHASH(&_ContractDelegationManager.CallOpts)
}

// DELEGATIONAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0x04a4f979.
//
// Solidity: function DELEGATION_APPROVAL_TYPEHASH() view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) DELEGATIONAPPROVALTYPEHASH() ([32]byte, error) {
	return _ContractDelegationManager.Contract.DELEGATIONAPPROVALTYPEHASH(&_ContractDelegationManager.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerSession) AllocationManager() (common.Address, error) {
	return _ContractDelegationManager.Contract.AllocationManager(&_ContractDelegationManager.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) AllocationManager() (common.Address, error) {
	return _ContractDelegationManager.Contract.AllocationManager(&_ContractDelegationManager.CallOpts)
}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCaller) BeaconChainETHStrategy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "beaconChainETHStrategy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerSession) BeaconChainETHStrategy() (common.Address, error) {
	return _ContractDelegationManager.Contract.BeaconChainETHStrategy(&_ContractDelegationManager.CallOpts)
}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) BeaconChainETHStrategy() (common.Address, error) {
	return _ContractDelegationManager.Contract.BeaconChainETHStrategy(&_ContractDelegationManager.CallOpts)
}

// CalculateDelegationApprovalDigestHash is a free data retrieval call binding the contract method 0x0b9f487a.
//
// Solidity: function calculateDelegationApprovalDigestHash(address staker, address operator, address approver, bytes32 approverSalt, uint256 expiry) view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCaller) CalculateDelegationApprovalDigestHash(opts *bind.CallOpts, staker common.Address, operator common.Address, approver common.Address, approverSalt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "calculateDelegationApprovalDigestHash", staker, operator, approver, approverSalt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateDelegationApprovalDigestHash is a free data retrieval call binding the contract method 0x0b9f487a.
//
// Solidity: function calculateDelegationApprovalDigestHash(address staker, address operator, address approver, bytes32 approverSalt, uint256 expiry) view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerSession) CalculateDelegationApprovalDigestHash(staker common.Address, operator common.Address, approver common.Address, approverSalt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractDelegationManager.Contract.CalculateDelegationApprovalDigestHash(&_ContractDelegationManager.CallOpts, staker, operator, approver, approverSalt, expiry)
}

// CalculateDelegationApprovalDigestHash is a free data retrieval call binding the contract method 0x0b9f487a.
//
// Solidity: function calculateDelegationApprovalDigestHash(address staker, address operator, address approver, bytes32 approverSalt, uint256 expiry) view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) CalculateDelegationApprovalDigestHash(staker common.Address, operator common.Address, approver common.Address, approverSalt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractDelegationManager.Contract.CalculateDelegationApprovalDigestHash(&_ContractDelegationManager.CallOpts, staker, operator, approver, approverSalt, expiry)
}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0x597b36da.
//
// Solidity: function calculateWithdrawalRoot((address,address,address,uint256,uint32,address[],uint256[]) withdrawal) pure returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCaller) CalculateWithdrawalRoot(opts *bind.CallOpts, withdrawal IDelegationManagerTypesWithdrawal) ([32]byte, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "calculateWithdrawalRoot", withdrawal)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0x597b36da.
//
// Solidity: function calculateWithdrawalRoot((address,address,address,uint256,uint32,address[],uint256[]) withdrawal) pure returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerSession) CalculateWithdrawalRoot(withdrawal IDelegationManagerTypesWithdrawal) ([32]byte, error) {
	return _ContractDelegationManager.Contract.CalculateWithdrawalRoot(&_ContractDelegationManager.CallOpts, withdrawal)
}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0x597b36da.
//
// Solidity: function calculateWithdrawalRoot((address,address,address,uint256,uint32,address[],uint256[]) withdrawal) pure returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) CalculateWithdrawalRoot(withdrawal IDelegationManagerTypesWithdrawal) ([32]byte, error) {
	return _ContractDelegationManager.Contract.CalculateWithdrawalRoot(&_ContractDelegationManager.CallOpts, withdrawal)
}

// CumulativeWithdrawalsQueued is a free data retrieval call binding the contract method 0xa1788484.
//
// Solidity: function cumulativeWithdrawalsQueued(address staker) view returns(uint256 totalQueued)
func (_ContractDelegationManager *ContractDelegationManagerCaller) CumulativeWithdrawalsQueued(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "cumulativeWithdrawalsQueued", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeWithdrawalsQueued is a free data retrieval call binding the contract method 0xa1788484.
//
// Solidity: function cumulativeWithdrawalsQueued(address staker) view returns(uint256 totalQueued)
func (_ContractDelegationManager *ContractDelegationManagerSession) CumulativeWithdrawalsQueued(staker common.Address) (*big.Int, error) {
	return _ContractDelegationManager.Contract.CumulativeWithdrawalsQueued(&_ContractDelegationManager.CallOpts, staker)
}

// CumulativeWithdrawalsQueued is a free data retrieval call binding the contract method 0xa1788484.
//
// Solidity: function cumulativeWithdrawalsQueued(address staker) view returns(uint256 totalQueued)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) CumulativeWithdrawalsQueued(staker common.Address) (*big.Int, error) {
	return _ContractDelegationManager.Contract.CumulativeWithdrawalsQueued(&_ContractDelegationManager.CallOpts, staker)
}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(address staker) view returns(address operator)
func (_ContractDelegationManager *ContractDelegationManagerCaller) DelegatedTo(opts *bind.CallOpts, staker common.Address) (common.Address, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "delegatedTo", staker)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(address staker) view returns(address operator)
func (_ContractDelegationManager *ContractDelegationManagerSession) DelegatedTo(staker common.Address) (common.Address, error) {
	return _ContractDelegationManager.Contract.DelegatedTo(&_ContractDelegationManager.CallOpts, staker)
}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(address staker) view returns(address operator)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) DelegatedTo(staker common.Address) (common.Address, error) {
	return _ContractDelegationManager.Contract.DelegatedTo(&_ContractDelegationManager.CallOpts, staker)
}

// DelegationApprover is a free data retrieval call binding the contract method 0x3cdeb5e0.
//
// Solidity: function delegationApprover(address operator) view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCaller) DelegationApprover(opts *bind.CallOpts, operator common.Address) (common.Address, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "delegationApprover", operator)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationApprover is a free data retrieval call binding the contract method 0x3cdeb5e0.
//
// Solidity: function delegationApprover(address operator) view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerSession) DelegationApprover(operator common.Address) (common.Address, error) {
	return _ContractDelegationManager.Contract.DelegationApprover(&_ContractDelegationManager.CallOpts, operator)
}

// DelegationApprover is a free data retrieval call binding the contract method 0x3cdeb5e0.
//
// Solidity: function delegationApprover(address operator) view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) DelegationApprover(operator common.Address) (common.Address, error) {
	return _ContractDelegationManager.Contract.DelegationApprover(&_ContractDelegationManager.CallOpts, operator)
}

// DelegationApproverSaltIsSpent is a free data retrieval call binding the contract method 0xbb45fef2.
//
// Solidity: function delegationApproverSaltIsSpent(address delegationApprover, bytes32 salt) view returns(bool spent)
func (_ContractDelegationManager *ContractDelegationManagerCaller) DelegationApproverSaltIsSpent(opts *bind.CallOpts, delegationApprover common.Address, salt [32]byte) (bool, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "delegationApproverSaltIsSpent", delegationApprover, salt)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DelegationApproverSaltIsSpent is a free data retrieval call binding the contract method 0xbb45fef2.
//
// Solidity: function delegationApproverSaltIsSpent(address delegationApprover, bytes32 salt) view returns(bool spent)
func (_ContractDelegationManager *ContractDelegationManagerSession) DelegationApproverSaltIsSpent(delegationApprover common.Address, salt [32]byte) (bool, error) {
	return _ContractDelegationManager.Contract.DelegationApproverSaltIsSpent(&_ContractDelegationManager.CallOpts, delegationApprover, salt)
}

// DelegationApproverSaltIsSpent is a free data retrieval call binding the contract method 0xbb45fef2.
//
// Solidity: function delegationApproverSaltIsSpent(address delegationApprover, bytes32 salt) view returns(bool spent)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) DelegationApproverSaltIsSpent(delegationApprover common.Address, salt [32]byte) (bool, error) {
	return _ContractDelegationManager.Contract.DelegationApproverSaltIsSpent(&_ContractDelegationManager.CallOpts, delegationApprover, salt)
}

// DepositScalingFactor is a free data retrieval call binding the contract method 0xbfae3fd2.
//
// Solidity: function depositScalingFactor(address staker, address strategy) view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerCaller) DepositScalingFactor(opts *bind.CallOpts, staker common.Address, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "depositScalingFactor", staker, strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositScalingFactor is a free data retrieval call binding the contract method 0xbfae3fd2.
//
// Solidity: function depositScalingFactor(address staker, address strategy) view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerSession) DepositScalingFactor(staker common.Address, strategy common.Address) (*big.Int, error) {
	return _ContractDelegationManager.Contract.DepositScalingFactor(&_ContractDelegationManager.CallOpts, staker, strategy)
}

// DepositScalingFactor is a free data retrieval call binding the contract method 0xbfae3fd2.
//
// Solidity: function depositScalingFactor(address staker, address strategy) view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) DepositScalingFactor(staker common.Address, strategy common.Address) (*big.Int, error) {
	return _ContractDelegationManager.Contract.DepositScalingFactor(&_ContractDelegationManager.CallOpts, staker, strategy)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerSession) DomainSeparator() ([32]byte, error) {
	return _ContractDelegationManager.Contract.DomainSeparator(&_ContractDelegationManager.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) DomainSeparator() ([32]byte, error) {
	return _ContractDelegationManager.Contract.DomainSeparator(&_ContractDelegationManager.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCaller) EigenPodManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "eigenPodManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerSession) EigenPodManager() (common.Address, error) {
	return _ContractDelegationManager.Contract.EigenPodManager(&_ContractDelegationManager.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) EigenPodManager() (common.Address, error) {
	return _ContractDelegationManager.Contract.EigenPodManager(&_ContractDelegationManager.CallOpts)
}

// GetDepositedShares is a free data retrieval call binding the contract method 0x66d5ba93.
//
// Solidity: function getDepositedShares(address staker) view returns(address[], uint256[])
func (_ContractDelegationManager *ContractDelegationManagerCaller) GetDepositedShares(opts *bind.CallOpts, staker common.Address) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "getDepositedShares", staker)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetDepositedShares is a free data retrieval call binding the contract method 0x66d5ba93.
//
// Solidity: function getDepositedShares(address staker) view returns(address[], uint256[])
func (_ContractDelegationManager *ContractDelegationManagerSession) GetDepositedShares(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _ContractDelegationManager.Contract.GetDepositedShares(&_ContractDelegationManager.CallOpts, staker)
}

// GetDepositedShares is a free data retrieval call binding the contract method 0x66d5ba93.
//
// Solidity: function getDepositedShares(address staker) view returns(address[], uint256[])
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) GetDepositedShares(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _ContractDelegationManager.Contract.GetDepositedShares(&_ContractDelegationManager.CallOpts, staker)
}

// GetOperatorShares is a free data retrieval call binding the contract method 0x90041347.
//
// Solidity: function getOperatorShares(address operator, address[] strategies) view returns(uint256[])
func (_ContractDelegationManager *ContractDelegationManagerCaller) GetOperatorShares(opts *bind.CallOpts, operator common.Address, strategies []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "getOperatorShares", operator, strategies)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetOperatorShares is a free data retrieval call binding the contract method 0x90041347.
//
// Solidity: function getOperatorShares(address operator, address[] strategies) view returns(uint256[])
func (_ContractDelegationManager *ContractDelegationManagerSession) GetOperatorShares(operator common.Address, strategies []common.Address) ([]*big.Int, error) {
	return _ContractDelegationManager.Contract.GetOperatorShares(&_ContractDelegationManager.CallOpts, operator, strategies)
}

// GetOperatorShares is a free data retrieval call binding the contract method 0x90041347.
//
// Solidity: function getOperatorShares(address operator, address[] strategies) view returns(uint256[])
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) GetOperatorShares(operator common.Address, strategies []common.Address) ([]*big.Int, error) {
	return _ContractDelegationManager.Contract.GetOperatorShares(&_ContractDelegationManager.CallOpts, operator, strategies)
}

// GetOperatorsShares is a free data retrieval call binding the contract method 0xf0e0e676.
//
// Solidity: function getOperatorsShares(address[] operators, address[] strategies) view returns(uint256[][])
func (_ContractDelegationManager *ContractDelegationManagerCaller) GetOperatorsShares(opts *bind.CallOpts, operators []common.Address, strategies []common.Address) ([][]*big.Int, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "getOperatorsShares", operators, strategies)

	if err != nil {
		return *new([][]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][]*big.Int)).(*[][]*big.Int)

	return out0, err

}

// GetOperatorsShares is a free data retrieval call binding the contract method 0xf0e0e676.
//
// Solidity: function getOperatorsShares(address[] operators, address[] strategies) view returns(uint256[][])
func (_ContractDelegationManager *ContractDelegationManagerSession) GetOperatorsShares(operators []common.Address, strategies []common.Address) ([][]*big.Int, error) {
	return _ContractDelegationManager.Contract.GetOperatorsShares(&_ContractDelegationManager.CallOpts, operators, strategies)
}

// GetOperatorsShares is a free data retrieval call binding the contract method 0xf0e0e676.
//
// Solidity: function getOperatorsShares(address[] operators, address[] strategies) view returns(uint256[][])
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) GetOperatorsShares(operators []common.Address, strategies []common.Address) ([][]*big.Int, error) {
	return _ContractDelegationManager.Contract.GetOperatorsShares(&_ContractDelegationManager.CallOpts, operators, strategies)
}

// GetQueuedWithdrawals is a free data retrieval call binding the contract method 0x5dd68579.
//
// Solidity: function getQueuedWithdrawals(address staker) view returns((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, uint256[][] shares)
func (_ContractDelegationManager *ContractDelegationManagerCaller) GetQueuedWithdrawals(opts *bind.CallOpts, staker common.Address) (struct {
	Withdrawals []IDelegationManagerTypesWithdrawal
	Shares      [][]*big.Int
}, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "getQueuedWithdrawals", staker)

	outstruct := new(struct {
		Withdrawals []IDelegationManagerTypesWithdrawal
		Shares      [][]*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Withdrawals = *abi.ConvertType(out[0], new([]IDelegationManagerTypesWithdrawal)).(*[]IDelegationManagerTypesWithdrawal)
	outstruct.Shares = *abi.ConvertType(out[1], new([][]*big.Int)).(*[][]*big.Int)

	return *outstruct, err

}

// GetQueuedWithdrawals is a free data retrieval call binding the contract method 0x5dd68579.
//
// Solidity: function getQueuedWithdrawals(address staker) view returns((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, uint256[][] shares)
func (_ContractDelegationManager *ContractDelegationManagerSession) GetQueuedWithdrawals(staker common.Address) (struct {
	Withdrawals []IDelegationManagerTypesWithdrawal
	Shares      [][]*big.Int
}, error) {
	return _ContractDelegationManager.Contract.GetQueuedWithdrawals(&_ContractDelegationManager.CallOpts, staker)
}

// GetQueuedWithdrawals is a free data retrieval call binding the contract method 0x5dd68579.
//
// Solidity: function getQueuedWithdrawals(address staker) view returns((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, uint256[][] shares)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) GetQueuedWithdrawals(staker common.Address) (struct {
	Withdrawals []IDelegationManagerTypesWithdrawal
	Shares      [][]*big.Int
}, error) {
	return _ContractDelegationManager.Contract.GetQueuedWithdrawals(&_ContractDelegationManager.CallOpts, staker)
}

// GetSlashableSharesInQueue is a free data retrieval call binding the contract method 0x6e174448.
//
// Solidity: function getSlashableSharesInQueue(address operator, address strategy) view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerCaller) GetSlashableSharesInQueue(opts *bind.CallOpts, operator common.Address, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "getSlashableSharesInQueue", operator, strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSlashableSharesInQueue is a free data retrieval call binding the contract method 0x6e174448.
//
// Solidity: function getSlashableSharesInQueue(address operator, address strategy) view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerSession) GetSlashableSharesInQueue(operator common.Address, strategy common.Address) (*big.Int, error) {
	return _ContractDelegationManager.Contract.GetSlashableSharesInQueue(&_ContractDelegationManager.CallOpts, operator, strategy)
}

// GetSlashableSharesInQueue is a free data retrieval call binding the contract method 0x6e174448.
//
// Solidity: function getSlashableSharesInQueue(address operator, address strategy) view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) GetSlashableSharesInQueue(operator common.Address, strategy common.Address) (*big.Int, error) {
	return _ContractDelegationManager.Contract.GetSlashableSharesInQueue(&_ContractDelegationManager.CallOpts, operator, strategy)
}

// GetWithdrawableShares is a free data retrieval call binding the contract method 0xc978f7ac.
//
// Solidity: function getWithdrawableShares(address staker, address[] strategies) view returns(uint256[] withdrawableShares, uint256[] depositShares)
func (_ContractDelegationManager *ContractDelegationManagerCaller) GetWithdrawableShares(opts *bind.CallOpts, staker common.Address, strategies []common.Address) (struct {
	WithdrawableShares []*big.Int
	DepositShares      []*big.Int
}, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "getWithdrawableShares", staker, strategies)

	outstruct := new(struct {
		WithdrawableShares []*big.Int
		DepositShares      []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.WithdrawableShares = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.DepositShares = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetWithdrawableShares is a free data retrieval call binding the contract method 0xc978f7ac.
//
// Solidity: function getWithdrawableShares(address staker, address[] strategies) view returns(uint256[] withdrawableShares, uint256[] depositShares)
func (_ContractDelegationManager *ContractDelegationManagerSession) GetWithdrawableShares(staker common.Address, strategies []common.Address) (struct {
	WithdrawableShares []*big.Int
	DepositShares      []*big.Int
}, error) {
	return _ContractDelegationManager.Contract.GetWithdrawableShares(&_ContractDelegationManager.CallOpts, staker, strategies)
}

// GetWithdrawableShares is a free data retrieval call binding the contract method 0xc978f7ac.
//
// Solidity: function getWithdrawableShares(address staker, address[] strategies) view returns(uint256[] withdrawableShares, uint256[] depositShares)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) GetWithdrawableShares(staker common.Address, strategies []common.Address) (struct {
	WithdrawableShares []*big.Int
	DepositShares      []*big.Int
}, error) {
	return _ContractDelegationManager.Contract.GetWithdrawableShares(&_ContractDelegationManager.CallOpts, staker, strategies)
}

// IsDelegated is a free data retrieval call binding the contract method 0x3e28391d.
//
// Solidity: function isDelegated(address staker) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerCaller) IsDelegated(opts *bind.CallOpts, staker common.Address) (bool, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "isDelegated", staker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDelegated is a free data retrieval call binding the contract method 0x3e28391d.
//
// Solidity: function isDelegated(address staker) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerSession) IsDelegated(staker common.Address) (bool, error) {
	return _ContractDelegationManager.Contract.IsDelegated(&_ContractDelegationManager.CallOpts, staker)
}

// IsDelegated is a free data retrieval call binding the contract method 0x3e28391d.
//
// Solidity: function isDelegated(address staker) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) IsDelegated(staker common.Address) (bool, error) {
	return _ContractDelegationManager.Contract.IsDelegated(&_ContractDelegationManager.CallOpts, staker)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerCaller) IsOperator(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "isOperator", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerSession) IsOperator(operator common.Address) (bool, error) {
	return _ContractDelegationManager.Contract.IsOperator(&_ContractDelegationManager.CallOpts, operator)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) IsOperator(operator common.Address) (bool, error) {
	return _ContractDelegationManager.Contract.IsOperator(&_ContractDelegationManager.CallOpts, operator)
}

// MinWithdrawalDelayBlocks is a free data retrieval call binding the contract method 0xc448feb8.
//
// Solidity: function minWithdrawalDelayBlocks() view returns(uint32)
func (_ContractDelegationManager *ContractDelegationManagerCaller) MinWithdrawalDelayBlocks(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "minWithdrawalDelayBlocks")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MinWithdrawalDelayBlocks is a free data retrieval call binding the contract method 0xc448feb8.
//
// Solidity: function minWithdrawalDelayBlocks() view returns(uint32)
func (_ContractDelegationManager *ContractDelegationManagerSession) MinWithdrawalDelayBlocks() (uint32, error) {
	return _ContractDelegationManager.Contract.MinWithdrawalDelayBlocks(&_ContractDelegationManager.CallOpts)
}

// MinWithdrawalDelayBlocks is a free data retrieval call binding the contract method 0xc448feb8.
//
// Solidity: function minWithdrawalDelayBlocks() view returns(uint32)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) MinWithdrawalDelayBlocks() (uint32, error) {
	return _ContractDelegationManager.Contract.MinWithdrawalDelayBlocks(&_ContractDelegationManager.CallOpts)
}

// OperatorShares is a free data retrieval call binding the contract method 0x778e55f3.
//
// Solidity: function operatorShares(address operator, address strategy) view returns(uint256 shares)
func (_ContractDelegationManager *ContractDelegationManagerCaller) OperatorShares(opts *bind.CallOpts, operator common.Address, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "operatorShares", operator, strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorShares is a free data retrieval call binding the contract method 0x778e55f3.
//
// Solidity: function operatorShares(address operator, address strategy) view returns(uint256 shares)
func (_ContractDelegationManager *ContractDelegationManagerSession) OperatorShares(operator common.Address, strategy common.Address) (*big.Int, error) {
	return _ContractDelegationManager.Contract.OperatorShares(&_ContractDelegationManager.CallOpts, operator, strategy)
}

// OperatorShares is a free data retrieval call binding the contract method 0x778e55f3.
//
// Solidity: function operatorShares(address operator, address strategy) view returns(uint256 shares)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) OperatorShares(operator common.Address, strategy common.Address) (*big.Int, error) {
	return _ContractDelegationManager.Contract.OperatorShares(&_ContractDelegationManager.CallOpts, operator, strategy)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerSession) Owner() (common.Address, error) {
	return _ContractDelegationManager.Contract.Owner(&_ContractDelegationManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) Owner() (common.Address, error) {
	return _ContractDelegationManager.Contract.Owner(&_ContractDelegationManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerSession) Paused(index uint8) (bool, error) {
	return _ContractDelegationManager.Contract.Paused(&_ContractDelegationManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractDelegationManager.Contract.Paused(&_ContractDelegationManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerSession) Paused0() (*big.Int, error) {
	return _ContractDelegationManager.Contract.Paused0(&_ContractDelegationManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractDelegationManager.Contract.Paused0(&_ContractDelegationManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractDelegationManager.Contract.PauserRegistry(&_ContractDelegationManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractDelegationManager.Contract.PauserRegistry(&_ContractDelegationManager.CallOpts)
}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xb7f06ebe.
//
// Solidity: function pendingWithdrawals(bytes32 withdrawalRoot) view returns(bool pending)
func (_ContractDelegationManager *ContractDelegationManagerCaller) PendingWithdrawals(opts *bind.CallOpts, withdrawalRoot [32]byte) (bool, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "pendingWithdrawals", withdrawalRoot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xb7f06ebe.
//
// Solidity: function pendingWithdrawals(bytes32 withdrawalRoot) view returns(bool pending)
func (_ContractDelegationManager *ContractDelegationManagerSession) PendingWithdrawals(withdrawalRoot [32]byte) (bool, error) {
	return _ContractDelegationManager.Contract.PendingWithdrawals(&_ContractDelegationManager.CallOpts, withdrawalRoot)
}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xb7f06ebe.
//
// Solidity: function pendingWithdrawals(bytes32 withdrawalRoot) view returns(bool pending)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) PendingWithdrawals(withdrawalRoot [32]byte) (bool, error) {
	return _ContractDelegationManager.Contract.PendingWithdrawals(&_ContractDelegationManager.CallOpts, withdrawalRoot)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCaller) PermissionController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "permissionController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerSession) PermissionController() (common.Address, error) {
	return _ContractDelegationManager.Contract.PermissionController(&_ContractDelegationManager.CallOpts)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) PermissionController() (common.Address, error) {
	return _ContractDelegationManager.Contract.PermissionController(&_ContractDelegationManager.CallOpts)
}

// QueuedWithdrawals is a free data retrieval call binding the contract method 0x99f5371b.
//
// Solidity: function queuedWithdrawals(bytes32 withdrawalRoot) view returns(address staker, address delegatedTo, address withdrawer, uint256 nonce, uint32 startBlock)
func (_ContractDelegationManager *ContractDelegationManagerCaller) QueuedWithdrawals(opts *bind.CallOpts, withdrawalRoot [32]byte) (struct {
	Staker      common.Address
	DelegatedTo common.Address
	Withdrawer  common.Address
	Nonce       *big.Int
	StartBlock  uint32
}, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "queuedWithdrawals", withdrawalRoot)

	outstruct := new(struct {
		Staker      common.Address
		DelegatedTo common.Address
		Withdrawer  common.Address
		Nonce       *big.Int
		StartBlock  uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Staker = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.DelegatedTo = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Withdrawer = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Nonce = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.StartBlock = *abi.ConvertType(out[4], new(uint32)).(*uint32)

	return *outstruct, err

}

// QueuedWithdrawals is a free data retrieval call binding the contract method 0x99f5371b.
//
// Solidity: function queuedWithdrawals(bytes32 withdrawalRoot) view returns(address staker, address delegatedTo, address withdrawer, uint256 nonce, uint32 startBlock)
func (_ContractDelegationManager *ContractDelegationManagerSession) QueuedWithdrawals(withdrawalRoot [32]byte) (struct {
	Staker      common.Address
	DelegatedTo common.Address
	Withdrawer  common.Address
	Nonce       *big.Int
	StartBlock  uint32
}, error) {
	return _ContractDelegationManager.Contract.QueuedWithdrawals(&_ContractDelegationManager.CallOpts, withdrawalRoot)
}

// QueuedWithdrawals is a free data retrieval call binding the contract method 0x99f5371b.
//
// Solidity: function queuedWithdrawals(bytes32 withdrawalRoot) view returns(address staker, address delegatedTo, address withdrawer, uint256 nonce, uint32 startBlock)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) QueuedWithdrawals(withdrawalRoot [32]byte) (struct {
	Staker      common.Address
	DelegatedTo common.Address
	Withdrawer  common.Address
	Nonce       *big.Int
	StartBlock  uint32
}, error) {
	return _ContractDelegationManager.Contract.QueuedWithdrawals(&_ContractDelegationManager.CallOpts, withdrawalRoot)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerSession) StrategyManager() (common.Address, error) {
	return _ContractDelegationManager.Contract.StrategyManager(&_ContractDelegationManager.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) StrategyManager() (common.Address, error) {
	return _ContractDelegationManager.Contract.StrategyManager(&_ContractDelegationManager.CallOpts)
}

// BurnOperatorShares is a paid mutator transaction binding the contract method 0xee74937f.
//
// Solidity: function burnOperatorShares(address operator, address strategy, uint64 prevMaxMagnitude, uint64 newMaxMagnitude) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) BurnOperatorShares(opts *bind.TransactOpts, operator common.Address, strategy common.Address, prevMaxMagnitude uint64, newMaxMagnitude uint64) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "burnOperatorShares", operator, strategy, prevMaxMagnitude, newMaxMagnitude)
}

// BurnOperatorShares is a paid mutator transaction binding the contract method 0xee74937f.
//
// Solidity: function burnOperatorShares(address operator, address strategy, uint64 prevMaxMagnitude, uint64 newMaxMagnitude) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) BurnOperatorShares(operator common.Address, strategy common.Address, prevMaxMagnitude uint64, newMaxMagnitude uint64) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.BurnOperatorShares(&_ContractDelegationManager.TransactOpts, operator, strategy, prevMaxMagnitude, newMaxMagnitude)
}

// BurnOperatorShares is a paid mutator transaction binding the contract method 0xee74937f.
//
// Solidity: function burnOperatorShares(address operator, address strategy, uint64 prevMaxMagnitude, uint64 newMaxMagnitude) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) BurnOperatorShares(operator common.Address, strategy common.Address, prevMaxMagnitude uint64, newMaxMagnitude uint64) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.BurnOperatorShares(&_ContractDelegationManager.TransactOpts, operator, strategy, prevMaxMagnitude, newMaxMagnitude)
}

// CompleteQueuedWithdrawal is a paid mutator transaction binding the contract method 0xe4cc3f90.
//
// Solidity: function completeQueuedWithdrawal((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, address[] tokens, bool receiveAsTokens) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) CompleteQueuedWithdrawal(opts *bind.TransactOpts, withdrawal IDelegationManagerTypesWithdrawal, tokens []common.Address, receiveAsTokens bool) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "completeQueuedWithdrawal", withdrawal, tokens, receiveAsTokens)
}

// CompleteQueuedWithdrawal is a paid mutator transaction binding the contract method 0xe4cc3f90.
//
// Solidity: function completeQueuedWithdrawal((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, address[] tokens, bool receiveAsTokens) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) CompleteQueuedWithdrawal(withdrawal IDelegationManagerTypesWithdrawal, tokens []common.Address, receiveAsTokens bool) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.CompleteQueuedWithdrawal(&_ContractDelegationManager.TransactOpts, withdrawal, tokens, receiveAsTokens)
}

// CompleteQueuedWithdrawal is a paid mutator transaction binding the contract method 0xe4cc3f90.
//
// Solidity: function completeQueuedWithdrawal((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, address[] tokens, bool receiveAsTokens) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) CompleteQueuedWithdrawal(withdrawal IDelegationManagerTypesWithdrawal, tokens []common.Address, receiveAsTokens bool) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.CompleteQueuedWithdrawal(&_ContractDelegationManager.TransactOpts, withdrawal, tokens, receiveAsTokens)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0x5f48e667.
//
// Solidity: function completeQueuedWithdrawals(address[][] tokens, bool[] receiveAsTokens, uint256 numToComplete) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) CompleteQueuedWithdrawals(opts *bind.TransactOpts, tokens [][]common.Address, receiveAsTokens []bool, numToComplete *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "completeQueuedWithdrawals", tokens, receiveAsTokens, numToComplete)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0x5f48e667.
//
// Solidity: function completeQueuedWithdrawals(address[][] tokens, bool[] receiveAsTokens, uint256 numToComplete) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) CompleteQueuedWithdrawals(tokens [][]common.Address, receiveAsTokens []bool, numToComplete *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.CompleteQueuedWithdrawals(&_ContractDelegationManager.TransactOpts, tokens, receiveAsTokens, numToComplete)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0x5f48e667.
//
// Solidity: function completeQueuedWithdrawals(address[][] tokens, bool[] receiveAsTokens, uint256 numToComplete) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) CompleteQueuedWithdrawals(tokens [][]common.Address, receiveAsTokens []bool, numToComplete *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.CompleteQueuedWithdrawals(&_ContractDelegationManager.TransactOpts, tokens, receiveAsTokens, numToComplete)
}

// CompleteQueuedWithdrawals0 is a paid mutator transaction binding the contract method 0x9435bb43.
//
// Solidity: function completeQueuedWithdrawals((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, bool[] receiveAsTokens) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) CompleteQueuedWithdrawals0(opts *bind.TransactOpts, withdrawals []IDelegationManagerTypesWithdrawal, tokens [][]common.Address, receiveAsTokens []bool) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "completeQueuedWithdrawals0", withdrawals, tokens, receiveAsTokens)
}

// CompleteQueuedWithdrawals0 is a paid mutator transaction binding the contract method 0x9435bb43.
//
// Solidity: function completeQueuedWithdrawals((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, bool[] receiveAsTokens) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) CompleteQueuedWithdrawals0(withdrawals []IDelegationManagerTypesWithdrawal, tokens [][]common.Address, receiveAsTokens []bool) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.CompleteQueuedWithdrawals0(&_ContractDelegationManager.TransactOpts, withdrawals, tokens, receiveAsTokens)
}

// CompleteQueuedWithdrawals0 is a paid mutator transaction binding the contract method 0x9435bb43.
//
// Solidity: function completeQueuedWithdrawals((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, bool[] receiveAsTokens) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) CompleteQueuedWithdrawals0(withdrawals []IDelegationManagerTypesWithdrawal, tokens [][]common.Address, receiveAsTokens []bool) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.CompleteQueuedWithdrawals0(&_ContractDelegationManager.TransactOpts, withdrawals, tokens, receiveAsTokens)
}

// DecreaseDelegatedShares is a paid mutator transaction binding the contract method 0x60a0d1ce.
//
// Solidity: function decreaseDelegatedShares(address staker, uint256 curDepositShares, uint64 beaconChainSlashingFactorDecrease) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) DecreaseDelegatedShares(opts *bind.TransactOpts, staker common.Address, curDepositShares *big.Int, beaconChainSlashingFactorDecrease uint64) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "decreaseDelegatedShares", staker, curDepositShares, beaconChainSlashingFactorDecrease)
}

// DecreaseDelegatedShares is a paid mutator transaction binding the contract method 0x60a0d1ce.
//
// Solidity: function decreaseDelegatedShares(address staker, uint256 curDepositShares, uint64 beaconChainSlashingFactorDecrease) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) DecreaseDelegatedShares(staker common.Address, curDepositShares *big.Int, beaconChainSlashingFactorDecrease uint64) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.DecreaseDelegatedShares(&_ContractDelegationManager.TransactOpts, staker, curDepositShares, beaconChainSlashingFactorDecrease)
}

// DecreaseDelegatedShares is a paid mutator transaction binding the contract method 0x60a0d1ce.
//
// Solidity: function decreaseDelegatedShares(address staker, uint256 curDepositShares, uint64 beaconChainSlashingFactorDecrease) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) DecreaseDelegatedShares(staker common.Address, curDepositShares *big.Int, beaconChainSlashingFactorDecrease uint64) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.DecreaseDelegatedShares(&_ContractDelegationManager.TransactOpts, staker, curDepositShares, beaconChainSlashingFactorDecrease)
}

// DelegateTo is a paid mutator transaction binding the contract method 0xeea9064b.
//
// Solidity: function delegateTo(address operator, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) DelegateTo(opts *bind.TransactOpts, operator common.Address, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "delegateTo", operator, approverSignatureAndExpiry, approverSalt)
}

// DelegateTo is a paid mutator transaction binding the contract method 0xeea9064b.
//
// Solidity: function delegateTo(address operator, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) DelegateTo(operator common.Address, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.DelegateTo(&_ContractDelegationManager.TransactOpts, operator, approverSignatureAndExpiry, approverSalt)
}

// DelegateTo is a paid mutator transaction binding the contract method 0xeea9064b.
//
// Solidity: function delegateTo(address operator, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) DelegateTo(operator common.Address, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.DelegateTo(&_ContractDelegationManager.TransactOpts, operator, approverSignatureAndExpiry, approverSalt)
}

// IncreaseDelegatedShares is a paid mutator transaction binding the contract method 0x3c651cf2.
//
// Solidity: function increaseDelegatedShares(address staker, address strategy, uint256 prevDepositShares, uint256 addedShares) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) IncreaseDelegatedShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, prevDepositShares *big.Int, addedShares *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "increaseDelegatedShares", staker, strategy, prevDepositShares, addedShares)
}

// IncreaseDelegatedShares is a paid mutator transaction binding the contract method 0x3c651cf2.
//
// Solidity: function increaseDelegatedShares(address staker, address strategy, uint256 prevDepositShares, uint256 addedShares) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) IncreaseDelegatedShares(staker common.Address, strategy common.Address, prevDepositShares *big.Int, addedShares *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.IncreaseDelegatedShares(&_ContractDelegationManager.TransactOpts, staker, strategy, prevDepositShares, addedShares)
}

// IncreaseDelegatedShares is a paid mutator transaction binding the contract method 0x3c651cf2.
//
// Solidity: function increaseDelegatedShares(address staker, address strategy, uint256 prevDepositShares, uint256 addedShares) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) IncreaseDelegatedShares(staker common.Address, strategy common.Address, prevDepositShares *big.Int, addedShares *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.IncreaseDelegatedShares(&_ContractDelegationManager.TransactOpts, staker, strategy, prevDepositShares, addedShares)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "initialize", initialOwner, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) Initialize(initialOwner common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.Initialize(&_ContractDelegationManager.TransactOpts, initialOwner, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) Initialize(initialOwner common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.Initialize(&_ContractDelegationManager.TransactOpts, initialOwner, initialPausedStatus)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0x54b7c96c.
//
// Solidity: function modifyOperatorDetails(address operator, address newDelegationApprover) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) ModifyOperatorDetails(opts *bind.TransactOpts, operator common.Address, newDelegationApprover common.Address) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "modifyOperatorDetails", operator, newDelegationApprover)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0x54b7c96c.
//
// Solidity: function modifyOperatorDetails(address operator, address newDelegationApprover) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) ModifyOperatorDetails(operator common.Address, newDelegationApprover common.Address) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.ModifyOperatorDetails(&_ContractDelegationManager.TransactOpts, operator, newDelegationApprover)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0x54b7c96c.
//
// Solidity: function modifyOperatorDetails(address operator, address newDelegationApprover) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) ModifyOperatorDetails(operator common.Address, newDelegationApprover common.Address) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.ModifyOperatorDetails(&_ContractDelegationManager.TransactOpts, operator, newDelegationApprover)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.Pause(&_ContractDelegationManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.Pause(&_ContractDelegationManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.PauseAll(&_ContractDelegationManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.PauseAll(&_ContractDelegationManager.TransactOpts)
}

// QueueWithdrawals is a paid mutator transaction binding the contract method 0x0dd8dd02.
//
// Solidity: function queueWithdrawals((address[],uint256[],address)[] params) returns(bytes32[])
func (_ContractDelegationManager *ContractDelegationManagerTransactor) QueueWithdrawals(opts *bind.TransactOpts, params []IDelegationManagerTypesQueuedWithdrawalParams) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "queueWithdrawals", params)
}

// QueueWithdrawals is a paid mutator transaction binding the contract method 0x0dd8dd02.
//
// Solidity: function queueWithdrawals((address[],uint256[],address)[] params) returns(bytes32[])
func (_ContractDelegationManager *ContractDelegationManagerSession) QueueWithdrawals(params []IDelegationManagerTypesQueuedWithdrawalParams) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.QueueWithdrawals(&_ContractDelegationManager.TransactOpts, params)
}

// QueueWithdrawals is a paid mutator transaction binding the contract method 0x0dd8dd02.
//
// Solidity: function queueWithdrawals((address[],uint256[],address)[] params) returns(bytes32[])
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) QueueWithdrawals(params []IDelegationManagerTypesQueuedWithdrawalParams) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.QueueWithdrawals(&_ContractDelegationManager.TransactOpts, params)
}

// Redelegate is a paid mutator transaction binding the contract method 0xa33a3433.
//
// Solidity: function redelegate(address newOperator, (bytes,uint256) newOperatorApproverSig, bytes32 approverSalt) returns(bytes32[] withdrawalRoots)
func (_ContractDelegationManager *ContractDelegationManagerTransactor) Redelegate(opts *bind.TransactOpts, newOperator common.Address, newOperatorApproverSig ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "redelegate", newOperator, newOperatorApproverSig, approverSalt)
}

// Redelegate is a paid mutator transaction binding the contract method 0xa33a3433.
//
// Solidity: function redelegate(address newOperator, (bytes,uint256) newOperatorApproverSig, bytes32 approverSalt) returns(bytes32[] withdrawalRoots)
func (_ContractDelegationManager *ContractDelegationManagerSession) Redelegate(newOperator common.Address, newOperatorApproverSig ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.Redelegate(&_ContractDelegationManager.TransactOpts, newOperator, newOperatorApproverSig, approverSalt)
}

// Redelegate is a paid mutator transaction binding the contract method 0xa33a3433.
//
// Solidity: function redelegate(address newOperator, (bytes,uint256) newOperatorApproverSig, bytes32 approverSalt) returns(bytes32[] withdrawalRoots)
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) Redelegate(newOperator common.Address, newOperatorApproverSig ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.Redelegate(&_ContractDelegationManager.TransactOpts, newOperator, newOperatorApproverSig, approverSalt)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x2aa6d888.
//
// Solidity: function registerAsOperator(address initDelegationApprover, uint32 allocationDelay, string metadataURI) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) RegisterAsOperator(opts *bind.TransactOpts, initDelegationApprover common.Address, allocationDelay uint32, metadataURI string) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "registerAsOperator", initDelegationApprover, allocationDelay, metadataURI)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x2aa6d888.
//
// Solidity: function registerAsOperator(address initDelegationApprover, uint32 allocationDelay, string metadataURI) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) RegisterAsOperator(initDelegationApprover common.Address, allocationDelay uint32, metadataURI string) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.RegisterAsOperator(&_ContractDelegationManager.TransactOpts, initDelegationApprover, allocationDelay, metadataURI)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x2aa6d888.
//
// Solidity: function registerAsOperator(address initDelegationApprover, uint32 allocationDelay, string metadataURI) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) RegisterAsOperator(initDelegationApprover common.Address, allocationDelay uint32, metadataURI string) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.RegisterAsOperator(&_ContractDelegationManager.TransactOpts, initDelegationApprover, allocationDelay, metadataURI)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.RenounceOwnership(&_ContractDelegationManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.RenounceOwnership(&_ContractDelegationManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.TransferOwnership(&_ContractDelegationManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.TransferOwnership(&_ContractDelegationManager.TransactOpts, newOwner)
}

// Undelegate is a paid mutator transaction binding the contract method 0xda8be864.
//
// Solidity: function undelegate(address staker) returns(bytes32[] withdrawalRoots)
func (_ContractDelegationManager *ContractDelegationManagerTransactor) Undelegate(opts *bind.TransactOpts, staker common.Address) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "undelegate", staker)
}

// Undelegate is a paid mutator transaction binding the contract method 0xda8be864.
//
// Solidity: function undelegate(address staker) returns(bytes32[] withdrawalRoots)
func (_ContractDelegationManager *ContractDelegationManagerSession) Undelegate(staker common.Address) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.Undelegate(&_ContractDelegationManager.TransactOpts, staker)
}

// Undelegate is a paid mutator transaction binding the contract method 0xda8be864.
//
// Solidity: function undelegate(address staker) returns(bytes32[] withdrawalRoots)
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) Undelegate(staker common.Address) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.Undelegate(&_ContractDelegationManager.TransactOpts, staker)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.Unpause(&_ContractDelegationManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.Unpause(&_ContractDelegationManager.TransactOpts, newPausedStatus)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x78296ec5.
//
// Solidity: function updateOperatorMetadataURI(address operator, string metadataURI) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) UpdateOperatorMetadataURI(opts *bind.TransactOpts, operator common.Address, metadataURI string) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "updateOperatorMetadataURI", operator, metadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x78296ec5.
//
// Solidity: function updateOperatorMetadataURI(address operator, string metadataURI) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) UpdateOperatorMetadataURI(operator common.Address, metadataURI string) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.UpdateOperatorMetadataURI(&_ContractDelegationManager.TransactOpts, operator, metadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x78296ec5.
//
// Solidity: function updateOperatorMetadataURI(address operator, string metadataURI) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) UpdateOperatorMetadataURI(operator common.Address, metadataURI string) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.UpdateOperatorMetadataURI(&_ContractDelegationManager.TransactOpts, operator, metadataURI)
}

// ContractDelegationManagerDelegationApproverUpdatedIterator is returned from FilterDelegationApproverUpdated and is used to iterate over the raw logs and unpacked data for DelegationApproverUpdated events raised by the ContractDelegationManager contract.
type ContractDelegationManagerDelegationApproverUpdatedIterator struct {
	Event *ContractDelegationManagerDelegationApproverUpdated // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerDelegationApproverUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerDelegationApproverUpdated)
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
		it.Event = new(ContractDelegationManagerDelegationApproverUpdated)
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
func (it *ContractDelegationManagerDelegationApproverUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerDelegationApproverUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerDelegationApproverUpdated represents a DelegationApproverUpdated event raised by the ContractDelegationManager contract.
type ContractDelegationManagerDelegationApproverUpdated struct {
	Operator              common.Address
	NewDelegationApprover common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterDelegationApproverUpdated is a free log retrieval operation binding the contract event 0x773b54c04d756fcc5e678111f7d730de3be98192000799eee3d63716055a87c6.
//
// Solidity: event DelegationApproverUpdated(address indexed operator, address newDelegationApprover)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterDelegationApproverUpdated(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerDelegationApproverUpdatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "DelegationApproverUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerDelegationApproverUpdatedIterator{contract: _ContractDelegationManager.contract, event: "DelegationApproverUpdated", logs: logs, sub: sub}, nil
}

// WatchDelegationApproverUpdated is a free log subscription operation binding the contract event 0x773b54c04d756fcc5e678111f7d730de3be98192000799eee3d63716055a87c6.
//
// Solidity: event DelegationApproverUpdated(address indexed operator, address newDelegationApprover)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchDelegationApproverUpdated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerDelegationApproverUpdated, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "DelegationApproverUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerDelegationApproverUpdated)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "DelegationApproverUpdated", log); err != nil {
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

// ParseDelegationApproverUpdated is a log parse operation binding the contract event 0x773b54c04d756fcc5e678111f7d730de3be98192000799eee3d63716055a87c6.
//
// Solidity: event DelegationApproverUpdated(address indexed operator, address newDelegationApprover)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseDelegationApproverUpdated(log types.Log) (*ContractDelegationManagerDelegationApproverUpdated, error) {
	event := new(ContractDelegationManagerDelegationApproverUpdated)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "DelegationApproverUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerDepositScalingFactorUpdatedIterator is returned from FilterDepositScalingFactorUpdated and is used to iterate over the raw logs and unpacked data for DepositScalingFactorUpdated events raised by the ContractDelegationManager contract.
type ContractDelegationManagerDepositScalingFactorUpdatedIterator struct {
	Event *ContractDelegationManagerDepositScalingFactorUpdated // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerDepositScalingFactorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerDepositScalingFactorUpdated)
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
		it.Event = new(ContractDelegationManagerDepositScalingFactorUpdated)
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
func (it *ContractDelegationManagerDepositScalingFactorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerDepositScalingFactorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerDepositScalingFactorUpdated represents a DepositScalingFactorUpdated event raised by the ContractDelegationManager contract.
type ContractDelegationManagerDepositScalingFactorUpdated struct {
	Staker                  common.Address
	Strategy                common.Address
	NewDepositScalingFactor *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterDepositScalingFactorUpdated is a free log retrieval operation binding the contract event 0x8be932bac54561f27260f95463d9b8ab37e06b2842e5ee2404157cc13df6eb8f.
//
// Solidity: event DepositScalingFactorUpdated(address staker, address strategy, uint256 newDepositScalingFactor)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterDepositScalingFactorUpdated(opts *bind.FilterOpts) (*ContractDelegationManagerDepositScalingFactorUpdatedIterator, error) {

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "DepositScalingFactorUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerDepositScalingFactorUpdatedIterator{contract: _ContractDelegationManager.contract, event: "DepositScalingFactorUpdated", logs: logs, sub: sub}, nil
}

// WatchDepositScalingFactorUpdated is a free log subscription operation binding the contract event 0x8be932bac54561f27260f95463d9b8ab37e06b2842e5ee2404157cc13df6eb8f.
//
// Solidity: event DepositScalingFactorUpdated(address staker, address strategy, uint256 newDepositScalingFactor)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchDepositScalingFactorUpdated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerDepositScalingFactorUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "DepositScalingFactorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerDepositScalingFactorUpdated)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "DepositScalingFactorUpdated", log); err != nil {
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

// ParseDepositScalingFactorUpdated is a log parse operation binding the contract event 0x8be932bac54561f27260f95463d9b8ab37e06b2842e5ee2404157cc13df6eb8f.
//
// Solidity: event DepositScalingFactorUpdated(address staker, address strategy, uint256 newDepositScalingFactor)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseDepositScalingFactorUpdated(log types.Log) (*ContractDelegationManagerDepositScalingFactorUpdated, error) {
	event := new(ContractDelegationManagerDepositScalingFactorUpdated)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "DepositScalingFactorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractDelegationManager contract.
type ContractDelegationManagerInitializedIterator struct {
	Event *ContractDelegationManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerInitialized)
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
		it.Event = new(ContractDelegationManagerInitialized)
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
func (it *ContractDelegationManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerInitialized represents a Initialized event raised by the ContractDelegationManager contract.
type ContractDelegationManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractDelegationManagerInitializedIterator, error) {

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerInitializedIterator{contract: _ContractDelegationManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerInitialized)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseInitialized(log types.Log) (*ContractDelegationManagerInitialized, error) {
	event := new(ContractDelegationManagerInitialized)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerOperatorMetadataURIUpdatedIterator is returned from FilterOperatorMetadataURIUpdated and is used to iterate over the raw logs and unpacked data for OperatorMetadataURIUpdated events raised by the ContractDelegationManager contract.
type ContractDelegationManagerOperatorMetadataURIUpdatedIterator struct {
	Event *ContractDelegationManagerOperatorMetadataURIUpdated // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerOperatorMetadataURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerOperatorMetadataURIUpdated)
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
		it.Event = new(ContractDelegationManagerOperatorMetadataURIUpdated)
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
func (it *ContractDelegationManagerOperatorMetadataURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerOperatorMetadataURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerOperatorMetadataURIUpdated represents a OperatorMetadataURIUpdated event raised by the ContractDelegationManager contract.
type ContractDelegationManagerOperatorMetadataURIUpdated struct {
	Operator    common.Address
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorMetadataURIUpdated is a free log retrieval operation binding the contract event 0x02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b6708090.
//
// Solidity: event OperatorMetadataURIUpdated(address indexed operator, string metadataURI)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterOperatorMetadataURIUpdated(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerOperatorMetadataURIUpdatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "OperatorMetadataURIUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerOperatorMetadataURIUpdatedIterator{contract: _ContractDelegationManager.contract, event: "OperatorMetadataURIUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorMetadataURIUpdated is a free log subscription operation binding the contract event 0x02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b6708090.
//
// Solidity: event OperatorMetadataURIUpdated(address indexed operator, string metadataURI)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchOperatorMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOperatorMetadataURIUpdated, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "OperatorMetadataURIUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerOperatorMetadataURIUpdated)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "OperatorMetadataURIUpdated", log); err != nil {
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

// ParseOperatorMetadataURIUpdated is a log parse operation binding the contract event 0x02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b6708090.
//
// Solidity: event OperatorMetadataURIUpdated(address indexed operator, string metadataURI)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseOperatorMetadataURIUpdated(log types.Log) (*ContractDelegationManagerOperatorMetadataURIUpdated, error) {
	event := new(ContractDelegationManagerOperatorMetadataURIUpdated)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "OperatorMetadataURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the ContractDelegationManager contract.
type ContractDelegationManagerOperatorRegisteredIterator struct {
	Event *ContractDelegationManagerOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerOperatorRegistered)
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
		it.Event = new(ContractDelegationManagerOperatorRegistered)
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
func (it *ContractDelegationManagerOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerOperatorRegistered represents a OperatorRegistered event raised by the ContractDelegationManager contract.
type ContractDelegationManagerOperatorRegistered struct {
	Operator           common.Address
	DelegationApprover common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0xa453db612af59e5521d6ab9284dc3e2d06af286eb1b1b7b771fce4716c19f2c1.
//
// Solidity: event OperatorRegistered(address indexed operator, address delegationApprover)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerOperatorRegisteredIterator{contract: _ContractDelegationManager.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0xa453db612af59e5521d6ab9284dc3e2d06af286eb1b1b7b771fce4716c19f2c1.
//
// Solidity: event OperatorRegistered(address indexed operator, address delegationApprover)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOperatorRegistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerOperatorRegistered)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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

// ParseOperatorRegistered is a log parse operation binding the contract event 0xa453db612af59e5521d6ab9284dc3e2d06af286eb1b1b7b771fce4716c19f2c1.
//
// Solidity: event OperatorRegistered(address indexed operator, address delegationApprover)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseOperatorRegistered(log types.Log) (*ContractDelegationManagerOperatorRegistered, error) {
	event := new(ContractDelegationManagerOperatorRegistered)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerOperatorSharesBurnedIterator is returned from FilterOperatorSharesBurned and is used to iterate over the raw logs and unpacked data for OperatorSharesBurned events raised by the ContractDelegationManager contract.
type ContractDelegationManagerOperatorSharesBurnedIterator struct {
	Event *ContractDelegationManagerOperatorSharesBurned // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerOperatorSharesBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerOperatorSharesBurned)
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
		it.Event = new(ContractDelegationManagerOperatorSharesBurned)
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
func (it *ContractDelegationManagerOperatorSharesBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerOperatorSharesBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerOperatorSharesBurned represents a OperatorSharesBurned event raised by the ContractDelegationManager contract.
type ContractDelegationManagerOperatorSharesBurned struct {
	Operator common.Address
	Strategy common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSharesBurned is a free log retrieval operation binding the contract event 0xeff6aab2bc3f7c648896e1522eae71d6c22e3b0e218206b3f40af0e4d204716b.
//
// Solidity: event OperatorSharesBurned(address indexed operator, address strategy, uint256 shares)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterOperatorSharesBurned(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerOperatorSharesBurnedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "OperatorSharesBurned", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerOperatorSharesBurnedIterator{contract: _ContractDelegationManager.contract, event: "OperatorSharesBurned", logs: logs, sub: sub}, nil
}

// WatchOperatorSharesBurned is a free log subscription operation binding the contract event 0xeff6aab2bc3f7c648896e1522eae71d6c22e3b0e218206b3f40af0e4d204716b.
//
// Solidity: event OperatorSharesBurned(address indexed operator, address strategy, uint256 shares)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchOperatorSharesBurned(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOperatorSharesBurned, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "OperatorSharesBurned", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerOperatorSharesBurned)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "OperatorSharesBurned", log); err != nil {
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

// ParseOperatorSharesBurned is a log parse operation binding the contract event 0xeff6aab2bc3f7c648896e1522eae71d6c22e3b0e218206b3f40af0e4d204716b.
//
// Solidity: event OperatorSharesBurned(address indexed operator, address strategy, uint256 shares)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseOperatorSharesBurned(log types.Log) (*ContractDelegationManagerOperatorSharesBurned, error) {
	event := new(ContractDelegationManagerOperatorSharesBurned)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "OperatorSharesBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerOperatorSharesDecreasedIterator is returned from FilterOperatorSharesDecreased and is used to iterate over the raw logs and unpacked data for OperatorSharesDecreased events raised by the ContractDelegationManager contract.
type ContractDelegationManagerOperatorSharesDecreasedIterator struct {
	Event *ContractDelegationManagerOperatorSharesDecreased // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerOperatorSharesDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerOperatorSharesDecreased)
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
		it.Event = new(ContractDelegationManagerOperatorSharesDecreased)
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
func (it *ContractDelegationManagerOperatorSharesDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerOperatorSharesDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerOperatorSharesDecreased represents a OperatorSharesDecreased event raised by the ContractDelegationManager contract.
type ContractDelegationManagerOperatorSharesDecreased struct {
	Operator common.Address
	Staker   common.Address
	Strategy common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSharesDecreased is a free log retrieval operation binding the contract event 0x6909600037b75d7b4733aedd815442b5ec018a827751c832aaff64eba5d6d2dd.
//
// Solidity: event OperatorSharesDecreased(address indexed operator, address staker, address strategy, uint256 shares)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterOperatorSharesDecreased(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerOperatorSharesDecreasedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "OperatorSharesDecreased", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerOperatorSharesDecreasedIterator{contract: _ContractDelegationManager.contract, event: "OperatorSharesDecreased", logs: logs, sub: sub}, nil
}

// WatchOperatorSharesDecreased is a free log subscription operation binding the contract event 0x6909600037b75d7b4733aedd815442b5ec018a827751c832aaff64eba5d6d2dd.
//
// Solidity: event OperatorSharesDecreased(address indexed operator, address staker, address strategy, uint256 shares)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchOperatorSharesDecreased(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOperatorSharesDecreased, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "OperatorSharesDecreased", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerOperatorSharesDecreased)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "OperatorSharesDecreased", log); err != nil {
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

// ParseOperatorSharesDecreased is a log parse operation binding the contract event 0x6909600037b75d7b4733aedd815442b5ec018a827751c832aaff64eba5d6d2dd.
//
// Solidity: event OperatorSharesDecreased(address indexed operator, address staker, address strategy, uint256 shares)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseOperatorSharesDecreased(log types.Log) (*ContractDelegationManagerOperatorSharesDecreased, error) {
	event := new(ContractDelegationManagerOperatorSharesDecreased)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "OperatorSharesDecreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerOperatorSharesIncreasedIterator is returned from FilterOperatorSharesIncreased and is used to iterate over the raw logs and unpacked data for OperatorSharesIncreased events raised by the ContractDelegationManager contract.
type ContractDelegationManagerOperatorSharesIncreasedIterator struct {
	Event *ContractDelegationManagerOperatorSharesIncreased // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerOperatorSharesIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerOperatorSharesIncreased)
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
		it.Event = new(ContractDelegationManagerOperatorSharesIncreased)
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
func (it *ContractDelegationManagerOperatorSharesIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerOperatorSharesIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerOperatorSharesIncreased represents a OperatorSharesIncreased event raised by the ContractDelegationManager contract.
type ContractDelegationManagerOperatorSharesIncreased struct {
	Operator common.Address
	Staker   common.Address
	Strategy common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSharesIncreased is a free log retrieval operation binding the contract event 0x1ec042c965e2edd7107b51188ee0f383e22e76179041ab3a9d18ff151405166c.
//
// Solidity: event OperatorSharesIncreased(address indexed operator, address staker, address strategy, uint256 shares)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterOperatorSharesIncreased(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerOperatorSharesIncreasedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "OperatorSharesIncreased", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerOperatorSharesIncreasedIterator{contract: _ContractDelegationManager.contract, event: "OperatorSharesIncreased", logs: logs, sub: sub}, nil
}

// WatchOperatorSharesIncreased is a free log subscription operation binding the contract event 0x1ec042c965e2edd7107b51188ee0f383e22e76179041ab3a9d18ff151405166c.
//
// Solidity: event OperatorSharesIncreased(address indexed operator, address staker, address strategy, uint256 shares)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchOperatorSharesIncreased(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOperatorSharesIncreased, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "OperatorSharesIncreased", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerOperatorSharesIncreased)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "OperatorSharesIncreased", log); err != nil {
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

// ParseOperatorSharesIncreased is a log parse operation binding the contract event 0x1ec042c965e2edd7107b51188ee0f383e22e76179041ab3a9d18ff151405166c.
//
// Solidity: event OperatorSharesIncreased(address indexed operator, address staker, address strategy, uint256 shares)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseOperatorSharesIncreased(log types.Log) (*ContractDelegationManagerOperatorSharesIncreased, error) {
	event := new(ContractDelegationManagerOperatorSharesIncreased)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "OperatorSharesIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractDelegationManager contract.
type ContractDelegationManagerOwnershipTransferredIterator struct {
	Event *ContractDelegationManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerOwnershipTransferred)
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
		it.Event = new(ContractDelegationManagerOwnershipTransferred)
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
func (it *ContractDelegationManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractDelegationManager contract.
type ContractDelegationManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractDelegationManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerOwnershipTransferredIterator{contract: _ContractDelegationManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerOwnershipTransferred)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractDelegationManagerOwnershipTransferred, error) {
	event := new(ContractDelegationManagerOwnershipTransferred)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractDelegationManager contract.
type ContractDelegationManagerPausedIterator struct {
	Event *ContractDelegationManagerPaused // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerPaused)
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
		it.Event = new(ContractDelegationManagerPaused)
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
func (it *ContractDelegationManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerPaused represents a Paused event raised by the ContractDelegationManager contract.
type ContractDelegationManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractDelegationManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerPausedIterator{contract: _ContractDelegationManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerPaused)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParsePaused(log types.Log) (*ContractDelegationManagerPaused, error) {
	event := new(ContractDelegationManagerPaused)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerSlashingWithdrawalCompletedIterator is returned from FilterSlashingWithdrawalCompleted and is used to iterate over the raw logs and unpacked data for SlashingWithdrawalCompleted events raised by the ContractDelegationManager contract.
type ContractDelegationManagerSlashingWithdrawalCompletedIterator struct {
	Event *ContractDelegationManagerSlashingWithdrawalCompleted // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerSlashingWithdrawalCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerSlashingWithdrawalCompleted)
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
		it.Event = new(ContractDelegationManagerSlashingWithdrawalCompleted)
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
func (it *ContractDelegationManagerSlashingWithdrawalCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerSlashingWithdrawalCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerSlashingWithdrawalCompleted represents a SlashingWithdrawalCompleted event raised by the ContractDelegationManager contract.
type ContractDelegationManagerSlashingWithdrawalCompleted struct {
	WithdrawalRoot [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSlashingWithdrawalCompleted is a free log retrieval operation binding the contract event 0x1f40400889274ed07b24845e5054a87a0cab969eb1277aafe61ae352e7c32a00.
//
// Solidity: event SlashingWithdrawalCompleted(bytes32 withdrawalRoot)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterSlashingWithdrawalCompleted(opts *bind.FilterOpts) (*ContractDelegationManagerSlashingWithdrawalCompletedIterator, error) {

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "SlashingWithdrawalCompleted")
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerSlashingWithdrawalCompletedIterator{contract: _ContractDelegationManager.contract, event: "SlashingWithdrawalCompleted", logs: logs, sub: sub}, nil
}

// WatchSlashingWithdrawalCompleted is a free log subscription operation binding the contract event 0x1f40400889274ed07b24845e5054a87a0cab969eb1277aafe61ae352e7c32a00.
//
// Solidity: event SlashingWithdrawalCompleted(bytes32 withdrawalRoot)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchSlashingWithdrawalCompleted(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerSlashingWithdrawalCompleted) (event.Subscription, error) {

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "SlashingWithdrawalCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerSlashingWithdrawalCompleted)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "SlashingWithdrawalCompleted", log); err != nil {
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

// ParseSlashingWithdrawalCompleted is a log parse operation binding the contract event 0x1f40400889274ed07b24845e5054a87a0cab969eb1277aafe61ae352e7c32a00.
//
// Solidity: event SlashingWithdrawalCompleted(bytes32 withdrawalRoot)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseSlashingWithdrawalCompleted(log types.Log) (*ContractDelegationManagerSlashingWithdrawalCompleted, error) {
	event := new(ContractDelegationManagerSlashingWithdrawalCompleted)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "SlashingWithdrawalCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerSlashingWithdrawalQueuedIterator is returned from FilterSlashingWithdrawalQueued and is used to iterate over the raw logs and unpacked data for SlashingWithdrawalQueued events raised by the ContractDelegationManager contract.
type ContractDelegationManagerSlashingWithdrawalQueuedIterator struct {
	Event *ContractDelegationManagerSlashingWithdrawalQueued // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerSlashingWithdrawalQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerSlashingWithdrawalQueued)
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
		it.Event = new(ContractDelegationManagerSlashingWithdrawalQueued)
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
func (it *ContractDelegationManagerSlashingWithdrawalQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerSlashingWithdrawalQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerSlashingWithdrawalQueued represents a SlashingWithdrawalQueued event raised by the ContractDelegationManager contract.
type ContractDelegationManagerSlashingWithdrawalQueued struct {
	WithdrawalRoot   [32]byte
	Withdrawal       IDelegationManagerTypesWithdrawal
	SharesToWithdraw []*big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSlashingWithdrawalQueued is a free log retrieval operation binding the contract event 0x26b2aae26516e8719ef50ea2f6831a2efbd4e37dccdf0f6936b27bc08e793e30.
//
// Solidity: event SlashingWithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal, uint256[] sharesToWithdraw)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterSlashingWithdrawalQueued(opts *bind.FilterOpts) (*ContractDelegationManagerSlashingWithdrawalQueuedIterator, error) {

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "SlashingWithdrawalQueued")
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerSlashingWithdrawalQueuedIterator{contract: _ContractDelegationManager.contract, event: "SlashingWithdrawalQueued", logs: logs, sub: sub}, nil
}

// WatchSlashingWithdrawalQueued is a free log subscription operation binding the contract event 0x26b2aae26516e8719ef50ea2f6831a2efbd4e37dccdf0f6936b27bc08e793e30.
//
// Solidity: event SlashingWithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal, uint256[] sharesToWithdraw)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchSlashingWithdrawalQueued(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerSlashingWithdrawalQueued) (event.Subscription, error) {

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "SlashingWithdrawalQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerSlashingWithdrawalQueued)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "SlashingWithdrawalQueued", log); err != nil {
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

// ParseSlashingWithdrawalQueued is a log parse operation binding the contract event 0x26b2aae26516e8719ef50ea2f6831a2efbd4e37dccdf0f6936b27bc08e793e30.
//
// Solidity: event SlashingWithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal, uint256[] sharesToWithdraw)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseSlashingWithdrawalQueued(log types.Log) (*ContractDelegationManagerSlashingWithdrawalQueued, error) {
	event := new(ContractDelegationManagerSlashingWithdrawalQueued)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "SlashingWithdrawalQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerStakerDelegatedIterator is returned from FilterStakerDelegated and is used to iterate over the raw logs and unpacked data for StakerDelegated events raised by the ContractDelegationManager contract.
type ContractDelegationManagerStakerDelegatedIterator struct {
	Event *ContractDelegationManagerStakerDelegated // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerStakerDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerStakerDelegated)
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
		it.Event = new(ContractDelegationManagerStakerDelegated)
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
func (it *ContractDelegationManagerStakerDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerStakerDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerStakerDelegated represents a StakerDelegated event raised by the ContractDelegationManager contract.
type ContractDelegationManagerStakerDelegated struct {
	Staker   common.Address
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakerDelegated is a free log retrieval operation binding the contract event 0xc3ee9f2e5fda98e8066a1f745b2df9285f416fe98cf2559cd21484b3d8743304.
//
// Solidity: event StakerDelegated(address indexed staker, address indexed operator)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterStakerDelegated(opts *bind.FilterOpts, staker []common.Address, operator []common.Address) (*ContractDelegationManagerStakerDelegatedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "StakerDelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerStakerDelegatedIterator{contract: _ContractDelegationManager.contract, event: "StakerDelegated", logs: logs, sub: sub}, nil
}

// WatchStakerDelegated is a free log subscription operation binding the contract event 0xc3ee9f2e5fda98e8066a1f745b2df9285f416fe98cf2559cd21484b3d8743304.
//
// Solidity: event StakerDelegated(address indexed staker, address indexed operator)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchStakerDelegated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerStakerDelegated, staker []common.Address, operator []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "StakerDelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerStakerDelegated)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "StakerDelegated", log); err != nil {
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

// ParseStakerDelegated is a log parse operation binding the contract event 0xc3ee9f2e5fda98e8066a1f745b2df9285f416fe98cf2559cd21484b3d8743304.
//
// Solidity: event StakerDelegated(address indexed staker, address indexed operator)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseStakerDelegated(log types.Log) (*ContractDelegationManagerStakerDelegated, error) {
	event := new(ContractDelegationManagerStakerDelegated)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "StakerDelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerStakerForceUndelegatedIterator is returned from FilterStakerForceUndelegated and is used to iterate over the raw logs and unpacked data for StakerForceUndelegated events raised by the ContractDelegationManager contract.
type ContractDelegationManagerStakerForceUndelegatedIterator struct {
	Event *ContractDelegationManagerStakerForceUndelegated // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerStakerForceUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerStakerForceUndelegated)
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
		it.Event = new(ContractDelegationManagerStakerForceUndelegated)
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
func (it *ContractDelegationManagerStakerForceUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerStakerForceUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerStakerForceUndelegated represents a StakerForceUndelegated event raised by the ContractDelegationManager contract.
type ContractDelegationManagerStakerForceUndelegated struct {
	Staker   common.Address
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakerForceUndelegated is a free log retrieval operation binding the contract event 0xf0eddf07e6ea14f388b47e1e94a0f464ecbd9eed4171130e0fc0e99fb4030a8a.
//
// Solidity: event StakerForceUndelegated(address indexed staker, address indexed operator)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterStakerForceUndelegated(opts *bind.FilterOpts, staker []common.Address, operator []common.Address) (*ContractDelegationManagerStakerForceUndelegatedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "StakerForceUndelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerStakerForceUndelegatedIterator{contract: _ContractDelegationManager.contract, event: "StakerForceUndelegated", logs: logs, sub: sub}, nil
}

// WatchStakerForceUndelegated is a free log subscription operation binding the contract event 0xf0eddf07e6ea14f388b47e1e94a0f464ecbd9eed4171130e0fc0e99fb4030a8a.
//
// Solidity: event StakerForceUndelegated(address indexed staker, address indexed operator)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchStakerForceUndelegated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerStakerForceUndelegated, staker []common.Address, operator []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "StakerForceUndelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerStakerForceUndelegated)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "StakerForceUndelegated", log); err != nil {
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

// ParseStakerForceUndelegated is a log parse operation binding the contract event 0xf0eddf07e6ea14f388b47e1e94a0f464ecbd9eed4171130e0fc0e99fb4030a8a.
//
// Solidity: event StakerForceUndelegated(address indexed staker, address indexed operator)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseStakerForceUndelegated(log types.Log) (*ContractDelegationManagerStakerForceUndelegated, error) {
	event := new(ContractDelegationManagerStakerForceUndelegated)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "StakerForceUndelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerStakerUndelegatedIterator is returned from FilterStakerUndelegated and is used to iterate over the raw logs and unpacked data for StakerUndelegated events raised by the ContractDelegationManager contract.
type ContractDelegationManagerStakerUndelegatedIterator struct {
	Event *ContractDelegationManagerStakerUndelegated // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerStakerUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerStakerUndelegated)
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
		it.Event = new(ContractDelegationManagerStakerUndelegated)
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
func (it *ContractDelegationManagerStakerUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerStakerUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerStakerUndelegated represents a StakerUndelegated event raised by the ContractDelegationManager contract.
type ContractDelegationManagerStakerUndelegated struct {
	Staker   common.Address
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakerUndelegated is a free log retrieval operation binding the contract event 0xfee30966a256b71e14bc0ebfc94315e28ef4a97a7131a9e2b7a310a73af44676.
//
// Solidity: event StakerUndelegated(address indexed staker, address indexed operator)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterStakerUndelegated(opts *bind.FilterOpts, staker []common.Address, operator []common.Address) (*ContractDelegationManagerStakerUndelegatedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "StakerUndelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerStakerUndelegatedIterator{contract: _ContractDelegationManager.contract, event: "StakerUndelegated", logs: logs, sub: sub}, nil
}

// WatchStakerUndelegated is a free log subscription operation binding the contract event 0xfee30966a256b71e14bc0ebfc94315e28ef4a97a7131a9e2b7a310a73af44676.
//
// Solidity: event StakerUndelegated(address indexed staker, address indexed operator)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchStakerUndelegated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerStakerUndelegated, staker []common.Address, operator []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "StakerUndelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerStakerUndelegated)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "StakerUndelegated", log); err != nil {
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

// ParseStakerUndelegated is a log parse operation binding the contract event 0xfee30966a256b71e14bc0ebfc94315e28ef4a97a7131a9e2b7a310a73af44676.
//
// Solidity: event StakerUndelegated(address indexed staker, address indexed operator)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseStakerUndelegated(log types.Log) (*ContractDelegationManagerStakerUndelegated, error) {
	event := new(ContractDelegationManagerStakerUndelegated)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "StakerUndelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractDelegationManager contract.
type ContractDelegationManagerUnpausedIterator struct {
	Event *ContractDelegationManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractDelegationManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerUnpaused)
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
		it.Event = new(ContractDelegationManagerUnpaused)
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
func (it *ContractDelegationManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerUnpaused represents a Unpaused event raised by the ContractDelegationManager contract.
type ContractDelegationManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractDelegationManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerUnpausedIterator{contract: _ContractDelegationManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerUnpaused)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseUnpaused(log types.Log) (*ContractDelegationManagerUnpaused, error) {
	event := new(ContractDelegationManagerUnpaused)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
