package elcontracts

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	avsdirectory "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AVSDirectory"
	allocationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"
	delegationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/DelegationManager"
	erc20 "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IERC20"
	rewardscoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IRewardsCoordinator"
	strategy "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IStrategy"
	strategymanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StrategyManager"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
)

type Config struct {
	DelegationManagerAddress  gethcommon.Address
	AvsDirectoryAddress       gethcommon.Address
	RewardsCoordinatorAddress gethcommon.Address
}

type ChainReader struct {
	logger             logging.Logger
	delegationManager  *delegationmanager.ContractDelegationManager
	strategyManager    *strategymanager.ContractStrategyManager
	avsDirectory       *avsdirectory.ContractAVSDirectory
	rewardsCoordinator *rewardscoordinator.ContractIRewardsCoordinator
	allocationManager  *allocationmanager.ContractAllocationManager
	ethClient          eth.HttpBackend
}

var errLegacyAVSsNotSupported = errors.New("method not supported for legacy AVSs")

func NewChainReader(
	delegationManager *delegationmanager.ContractDelegationManager,
	strategyManager *strategymanager.ContractStrategyManager,
	avsDirectory *avsdirectory.ContractAVSDirectory,
	rewardsCoordinator *rewardscoordinator.ContractIRewardsCoordinator,
	allocationManager *allocationmanager.ContractAllocationManager,
	logger logging.Logger,
	ethClient eth.HttpBackend,
) *ChainReader {
	logger = logger.With(logging.ComponentKey, "elcontracts/reader")

	return &ChainReader{
		delegationManager:  delegationManager,
		strategyManager:    strategyManager,
		avsDirectory:       avsDirectory,
		rewardsCoordinator: rewardsCoordinator,
		allocationManager:  allocationManager,
		logger:             logger,
		ethClient:          ethClient,
	}
}

// BuildELChainReader creates a new ELChainReader
// Deprecated: Use BuildFromConfig instead
func BuildELChainReader(
	delegationManagerAddr gethcommon.Address,
	avsDirectoryAddr gethcommon.Address,
	ethClient eth.HttpBackend,
	logger logging.Logger,
) (*ChainReader, error) {
	elContractBindings, err := NewEigenlayerContractBindings(
		delegationManagerAddr,
		avsDirectoryAddr,
		ethClient,
		logger,
	)
	if err != nil {
		return nil, err
	}
	return NewChainReader(
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.AvsDirectory,
		elContractBindings.RewardsCoordinator,
		elContractBindings.AllocationManager,
		logger,
		ethClient,
	), nil
}

func NewReaderFromConfig(
	cfg Config,
	ethClient eth.HttpBackend,
	logger logging.Logger,
) (*ChainReader, error) {
	elContractBindings, err := NewBindingsFromConfig(
		cfg,
		ethClient,
		logger,
	)
	if err != nil {
		return nil, err
	}
	return NewChainReader(
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.AvsDirectory,
		elContractBindings.RewardsCoordinator,
		elContractBindings.AllocationManager,
		logger,
		ethClient,
	), nil
}

func (r *ChainReader) IsOperatorRegistered(
	ctx context.Context,
	operator types.Operator,
) (bool, error) {
	if r.delegationManager == nil {
		return false, errors.New("DelegationManager contract not provided")
	}

	isOperator, err := r.delegationManager.IsOperator(
		&bind.CallOpts{Context: ctx},
		gethcommon.HexToAddress(operator.Address),
	)
	if err != nil {
		return false, err
	}

	return isOperator, nil
}

// GetStakerShares returns the amount of shares that a staker has in all of the strategies in which they have nonzero
// shares
func (r *ChainReader) GetStakerShares(
	ctx context.Context,
	stakerAddress gethcommon.Address,
	blockNumber *big.Int,
) ([]gethcommon.Address, []*big.Int, error) {
	return r.delegationManager.GetDepositedShares(&bind.CallOpts{Context: ctx}, stakerAddress)
}

// GetDelegatedOperator returns the operator that a staker has delegated to
func (r *ChainReader) GetDelegatedOperator(
	ctx context.Context,
	stakerAddress gethcommon.Address,
	blockNumber *big.Int,
) (gethcommon.Address, error) {
	return r.delegationManager.DelegatedTo(&bind.CallOpts{Context: ctx}, stakerAddress)
}

func (r *ChainReader) GetOperatorDetails(
	ctx context.Context,
	operator types.Operator,
) (types.Operator, error) {
	if r.delegationManager == nil {
		return types.Operator{}, errors.New("DelegationManager contract not provided")
	}

	delegationManagerAddress, err := r.delegationManager.DelegationApprover(
		&bind.CallOpts{Context: ctx},
		gethcommon.HexToAddress(operator.Address),
	)
	if err != nil {
		return types.Operator{}, err
	}

	isSet, delay, err := r.allocationManager.GetAllocationDelay(
		&bind.CallOpts{
			Context: ctx,
		},
		gethcommon.HexToAddress(operator.Address),
	)
	if err == nil {
		return types.Operator{}, err
	}

	var allocationDelay uint32
	if isSet {
		allocationDelay = delay
	} else {
		allocationDelay = 0
	}

	return types.Operator{
		Address:                   operator.Address,
		DelegationApproverAddress: delegationManagerAddress.Hex(),
		AllocationDelay:           allocationDelay,
	}, nil
}

// GetStrategyAndUnderlyingToken returns the strategy contract and the underlying token address
func (r *ChainReader) GetStrategyAndUnderlyingToken(
	ctx context.Context,
	strategyAddr gethcommon.Address,
) (*strategy.ContractIStrategy, gethcommon.Address, error) {
	contractStrategy, err := strategy.NewContractIStrategy(strategyAddr, r.ethClient)
	if err != nil {
		return nil, gethcommon.Address{}, utils.WrapError("Failed to fetch strategy contract", err)
	}
	underlyingTokenAddr, err := contractStrategy.UnderlyingToken(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, gethcommon.Address{}, utils.WrapError("Failed to fetch token contract", err)
	}
	return contractStrategy, underlyingTokenAddr, nil
}

// GetStrategyAndUnderlyingERC20Token returns the strategy contract, the erc20 bindings for the underlying token
// and the underlying token address
func (r *ChainReader) GetStrategyAndUnderlyingERC20Token(
	ctx context.Context,
	strategyAddr gethcommon.Address,
) (*strategy.ContractIStrategy, erc20.ContractIERC20Methods, gethcommon.Address, error) {
	contractStrategy, err := strategy.NewContractIStrategy(strategyAddr, r.ethClient)
	if err != nil {
		return nil, nil, gethcommon.Address{}, utils.WrapError("Failed to fetch strategy contract", err)
	}
	underlyingTokenAddr, err := contractStrategy.UnderlyingToken(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, nil, gethcommon.Address{}, utils.WrapError("Failed to fetch token contract", err)
	}
	contractUnderlyingToken, err := erc20.NewContractIERC20(underlyingTokenAddr, r.ethClient)
	if err != nil {
		return nil, nil, gethcommon.Address{}, utils.WrapError("Failed to fetch token contract", err)
	}
	return contractStrategy, contractUnderlyingToken, underlyingTokenAddr, nil
}

func (r *ChainReader) GetOperatorSharesInStrategy(
	ctx context.Context,
	operatorAddr gethcommon.Address,
	strategyAddr gethcommon.Address,
) (*big.Int, error) {
	if r.delegationManager == nil {
		return &big.Int{}, errors.New("DelegationManager contract not provided")
	}

	return r.delegationManager.OperatorShares(
		&bind.CallOpts{Context: ctx},
		operatorAddr,
		strategyAddr,
	)
}

func (r *ChainReader) CalculateDelegationApprovalDigestHash(
	ctx context.Context,
	staker gethcommon.Address,
	operator gethcommon.Address,
	delegationApprover gethcommon.Address,
	approverSalt [32]byte,
	expiry *big.Int,
) ([32]byte, error) {
	if r.delegationManager == nil {
		return [32]byte{}, errors.New("DelegationManager contract not provided")
	}

	return r.delegationManager.CalculateDelegationApprovalDigestHash(
		&bind.CallOpts{Context: ctx},
		staker,
		operator,
		delegationApprover,
		approverSalt,
		expiry,
	)
}

func (r *ChainReader) CalculateOperatorAVSRegistrationDigestHash(
	ctx context.Context,
	operator gethcommon.Address,
	avs gethcommon.Address,
	salt [32]byte,
	expiry *big.Int,
) ([32]byte, error) {
	if r.avsDirectory == nil {
		return [32]byte{}, errors.New("AVSDirectory contract not provided")
	}

	return r.avsDirectory.CalculateOperatorAVSRegistrationDigestHash(
		&bind.CallOpts{Context: ctx},
		operator,
		avs,
		salt,
		expiry,
	)
}

func (r *ChainReader) GetDistributionRootsLength(ctx context.Context) (*big.Int, error) {
	if r.rewardsCoordinator == nil {
		return nil, errors.New("RewardsCoordinator contract not provided")
	}

	return r.rewardsCoordinator.GetDistributionRootsLength(&bind.CallOpts{Context: ctx})
}

func (r *ChainReader) CurrRewardsCalculationEndTimestamp(ctx context.Context) (uint32, error) {
	if r.rewardsCoordinator == nil {
		return 0, errors.New("RewardsCoordinator contract not provided")
	}

	return r.rewardsCoordinator.CurrRewardsCalculationEndTimestamp(&bind.CallOpts{Context: ctx})
}

func (r *ChainReader) GetCurrentClaimableDistributionRoot(
	ctx context.Context,
) (rewardscoordinator.IRewardsCoordinatorTypesDistributionRoot, error) {
	if r.rewardsCoordinator == nil {
		return rewardscoordinator.IRewardsCoordinatorTypesDistributionRoot{}, errors.New(
			"RewardsCoordinator contract not provided",
		)
	}

	return r.rewardsCoordinator.GetCurrentClaimableDistributionRoot(&bind.CallOpts{Context: ctx})
}

func (r *ChainReader) GetRootIndexFromHash(
	ctx context.Context,
	rootHash [32]byte,
) (uint32, error) {
	if r.rewardsCoordinator == nil {
		return 0, errors.New("RewardsCoordinator contract not provided")
	}

	return r.rewardsCoordinator.GetRootIndexFromHash(&bind.CallOpts{Context: ctx}, rootHash)
}

func (r *ChainReader) GetCumulativeClaimed(
	ctx context.Context,
	earner gethcommon.Address,
	token gethcommon.Address,
) (*big.Int, error) {
	if r.rewardsCoordinator == nil {
		return nil, errors.New("RewardsCoordinator contract not provided")
	}

	return r.rewardsCoordinator.CumulativeClaimed(&bind.CallOpts{Context: ctx}, earner, token)
}

func (r *ChainReader) CheckClaim(
	ctx context.Context,
	claim rewardscoordinator.IRewardsCoordinatorTypesRewardsMerkleClaim,
) (bool, error) {
	if r.rewardsCoordinator == nil {
		return false, errors.New("RewardsCoordinator contract not provided")
	}

	return r.rewardsCoordinator.CheckClaim(&bind.CallOpts{Context: ctx}, claim)
}

func (r *ChainReader) GetAllocatableMagnitude(
	ctx context.Context,
	operatorAddress gethcommon.Address,
	strategyAddress gethcommon.Address,
) (uint64, error) {
	if r.allocationManager == nil {
		return 0, errors.New("AllocationManager contract not provided")
	}

	return r.allocationManager.GetAllocatableMagnitude(&bind.CallOpts{Context: ctx}, operatorAddress, strategyAddress)
}

func (r *ChainReader) GetMaxMagnitudes(
	ctx context.Context,
	operatorAddress gethcommon.Address,
	strategyAddresses []gethcommon.Address,
) ([]uint64, error) {
	if r.allocationManager == nil {
		return []uint64{}, errors.New("AllocationManager contract not provided")
	}

	return r.allocationManager.GetMaxMagnitudes0(&bind.CallOpts{Context: ctx}, operatorAddress, strategyAddresses)
}

func (r *ChainReader) GetAllocationInfo(
	ctx context.Context,
	operatorAddress gethcommon.Address,
	strategyAddress gethcommon.Address,
) ([]AllocationInfo, error) {
	if r.allocationManager == nil {
		return nil, errors.New("AllocationManager contract not provided")
	}

	opSets, allocationInfo, err := r.allocationManager.GetStrategyAllocations(
		&bind.CallOpts{Context: ctx},
		operatorAddress,
		strategyAddress,
	)
	if err != nil {
		return nil, err
	}

	allocationsInfo := make([]AllocationInfo, len(opSets))
	for i, opSet := range opSets {
		allocationsInfo[i] = AllocationInfo{
			OperatorSetId:    opSet.Id,
			AvsAddress:       opSet.Avs,
			CurrentMagnitude: big.NewInt(int64(allocationInfo[i].CurrentMagnitude)),
			PendingDiff:      allocationInfo[i].PendingDiff,
			EffectBlock:      allocationInfo[i].EffectBlock,
		}
	}

	return allocationsInfo, nil
}

func (r *ChainReader) GetOperatorShares(
	ctx context.Context,
	operatorAddress gethcommon.Address,
	strategyAddresses []gethcommon.Address,
) ([]*big.Int, error) {
	if r.delegationManager == nil {
		return nil, errors.New("DelegationManager contract not provided")
	}

	return r.delegationManager.GetOperatorShares(&bind.CallOpts{
		Context: ctx,
	}, operatorAddress, strategyAddresses)
}

func (r *ChainReader) GetOperatorsShares(
	ctx context.Context,
	operatorAddress []gethcommon.Address,
	strategyAddresses []gethcommon.Address,
) ([][]*big.Int, error) {
	if r.delegationManager == nil {
		return nil, errors.New("DelegationManager contract not provided")
	}

	return r.delegationManager.GetOperatorsShares(&bind.CallOpts{Context: ctx}, operatorAddress, strategyAddresses)
}

// GetNumOperatorSetsForOperator returns the number of operator sets that an operator is part of
// Doesn't include M2 AVSs
func (r *ChainReader) GetNumOperatorSetsForOperator(
	ctx context.Context,
	operatorAddress gethcommon.Address,
) (*big.Int, error) {
	opSets, err := r.allocationManager.GetAllocatedSets(&bind.CallOpts{Context: ctx}, operatorAddress)
	if err != nil {
		return nil, err
	}
	return big.NewInt(int64(len(opSets))), nil
}

// GetOperatorSetsForOperator returns the list of operator sets that an operator is part of
// Doesn't include M2 AVSs
func (r *ChainReader) GetOperatorSetsForOperator(
	ctx context.Context,
	operatorAddress gethcommon.Address,
) ([]allocationmanager.OperatorSet, error) {
	// TODO: we're fetching max int64 operatorSets here. What's the practical limit for timeout by RPC? do we need to
	// paginate?
	return r.allocationManager.GetAllocatedSets(&bind.CallOpts{Context: ctx}, operatorAddress)
}

// IsOperatorRegisteredWithOperatorSet returns if an operator is registered with a specific operator set
func (r *ChainReader) IsOperatorRegisteredWithOperatorSet(
	ctx context.Context,
	operatorAddress gethcommon.Address,
	operatorSet allocationmanager.OperatorSet,
) (bool, error) {
	if operatorSet.Id == 0 {
		// this is an M2 AVS
		status, err := r.avsDirectory.AvsOperatorStatus(&bind.CallOpts{Context: ctx}, operatorSet.Avs, operatorAddress)
		if err != nil {
			return false, err
		}

		return status == 1, nil
	} else {
		registeredOperatorSets, err := r.allocationManager.GetRegisteredSets(&bind.CallOpts{Context: ctx}, operatorAddress)
		if err != nil {
			return false, err
		}
		for _, registeredOperatorSet := range registeredOperatorSets {
			if registeredOperatorSet.Id == operatorSet.Id && registeredOperatorSet.Avs == operatorSet.Avs {
				return true, nil
			}
		}

		return false, nil
	}
}

// GetOperatorsForOperatorSet returns the list of operators in a specific operator set
// Not supported for M2 AVSs
func (r *ChainReader) GetOperatorsForOperatorSet(
	ctx context.Context,
	operatorSet allocationmanager.OperatorSet,
) ([]gethcommon.Address, error) {
	if operatorSet.Id == 0 {
		return nil, errLegacyAVSsNotSupported
	} else {
		return r.allocationManager.GetMembers(&bind.CallOpts{Context: ctx}, operatorSet)
	}
}

// GetNumOperatorsForOperatorSet returns the number of operators in a specific operator set
func (r *ChainReader) GetNumOperatorsForOperatorSet(
	ctx context.Context,
	operatorSet allocationmanager.OperatorSet,
) (*big.Int, error) {
	if operatorSet.Id == 0 {
		return nil, errLegacyAVSsNotSupported
	} else {
		return r.allocationManager.GetMemberCount(&bind.CallOpts{Context: ctx}, operatorSet)
	}
}

// GetStrategiesForOperatorSet returns the list of strategies that an operator set takes into account
// Not supported for M2 AVSs
func (r *ChainReader) GetStrategiesForOperatorSet(
	ctx context.Context,
	operatorSet allocationmanager.OperatorSet,
) ([]gethcommon.Address, error) {
	if operatorSet.Id == 0 {
		return nil, errLegacyAVSsNotSupported
	} else {
		return r.allocationManager.GetStrategiesInOperatorSet(&bind.CallOpts{Context: ctx}, operatorSet)
	}
}

// GetSlashableSharesForOperatorSets returns the strategies the operatorSets take into account, their
// operators, and the minimum amount of shares that are slashable by the operatorSets.
// Not supported for M2 AVSs
func (r *ChainReader) GetSlashableSharesForOperatorSets(
	ctx context.Context,
	operatorSets []allocationmanager.OperatorSet,
) ([]OperatorSetStakes, error) {
	operatorSetStakes := make([]OperatorSetStakes, len(operatorSets))
	currentBlock, err := r.ethClient.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}
	for i, operatorSet := range operatorSets {
		operators, err := r.GetOperatorsForOperatorSet(ctx, operatorSet)
		if err != nil {
			return nil, err
		}

		strategies, err := r.GetStrategiesForOperatorSet(ctx, operatorSet)
		if err != nil {
			return nil, err
		}

		slashableShares, err := r.allocationManager.GetMinimumSlashableStake(
			&bind.CallOpts{Context: ctx},
			allocationmanager.OperatorSet{
				Id:  operatorSet.Id,
				Avs: operatorSet.Avs,
			},
			operators,
			strategies,
			uint32(currentBlock),
		)
		if err != nil {
			return nil, err
		}

		operatorSetStakes[i] = OperatorSetStakes{
			OperatorSet:     operatorSet,
			Strategies:      strategies,
			Operators:       operators,
			SlashableStakes: slashableShares,
		}
	}

	return operatorSetStakes, nil
}

// GetDelegatedAndSlashableSharesForOperatorSetsBefore returns the strategies the operatorSets take into account, their
// operators, and the minimum amount of shares that multiple operators delegated to them and slashable by the
// operatorSets before a given timestamp.
// Timestamp must be in the future. Used to underestimate future slashable stake.
// Not supported for M2 AVSs
func (r *ChainReader) GetDelegatedAndSlashableSharesForOperatorSetsBefore(
	ctx context.Context,
	operatorSets []allocationmanager.OperatorSet,
	futureBlock uint32,
) ([]OperatorSetStakes, error) {
	operatorSetStakes := make([]OperatorSetStakes, len(operatorSets))
	for i, operatorSet := range operatorSets {
		operators, err := r.GetOperatorsForOperatorSet(ctx, operatorSet)
		if err != nil {
			return nil, err
		}

		strategies, err := r.GetStrategiesForOperatorSet(ctx, operatorSet)
		if err != nil {
			return nil, err
		}

		slashableShares, err := r.allocationManager.GetMinimumSlashableStake(
			&bind.CallOpts{Context: ctx},
			allocationmanager.OperatorSet{
				Id:  operatorSet.Id,
				Avs: operatorSet.Avs,
			},
			operators,
			strategies,
			futureBlock,
		)
		if err != nil {
			return nil, err
		}

		operatorSetStakes[i] = OperatorSetStakes{
			OperatorSet:     operatorSet,
			Strategies:      strategies,
			Operators:       operators,
			SlashableStakes: slashableShares,
		}
	}

	return operatorSetStakes, nil
}

func (r *ChainReader) GetAllocationDelay(
	ctx context.Context,
	operatorAddress gethcommon.Address,
) (uint32, error) {
	if r.allocationManager == nil {
		return 0, errors.New("AllocationManager contract not provided")
	}
	isSet, delay, err := r.allocationManager.GetAllocationDelay(&bind.CallOpts{Context: ctx}, operatorAddress)
	if err != nil {
		return 0, err
	}
	if !isSet {
		return 0, errors.New("allocation delay not set")
	}
	return delay, nil
}

func (r *ChainReader) GetRegisteredSets(
	ctx context.Context,
	operatorAddress gethcommon.Address,
) ([]allocationmanager.OperatorSet, error) {
	if r.allocationManager == nil {
		return nil, errors.New("AllocationManager contract not provided")
	}
	return r.allocationManager.GetRegisteredSets(&bind.CallOpts{Context: ctx}, operatorAddress)
}
