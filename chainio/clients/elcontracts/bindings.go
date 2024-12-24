// Package elcontracts bindings.go contains functions that create contract bindings for the Eigenlayer Core contracts
// These functions are meant to be used by constructors of the chainio package.
package elcontracts

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	allocationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"
	delegationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/DelegationManager"
	avsdirectory "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IAVSDirectory"
	rewardscoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IRewardsCoordinator"
	strategymanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StrategyManager"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/utils"
)

// ContractBindings contains the contract bindings for the EigenLayer Core contracts
//
// Unclear why geth bindings don't store and expose the contract address,
// so we also store them here in case the different constructors that use this struct need them
type ContractBindings struct {
	StrategyManagerAddr       gethcommon.Address
	DelegationManagerAddr     gethcommon.Address
	AvsDirectoryAddr          gethcommon.Address
	AllocationManagerAddr     gethcommon.Address
	RewardsCoordinatorAddress gethcommon.Address
	DelegationManager         *delegationmanager.ContractDelegationManager
	StrategyManager           *strategymanager.ContractStrategyManager
	AvsDirectory              *avsdirectory.ContractIAVSDirectory
	AllocationManager         *allocationmanager.ContractAllocationManager
	RewardsCoordinator        *rewardscoordinator.ContractIRewardsCoordinator
}

func NewBindingsFromConfig(
	cfg Config,
	client eth.HttpBackend,
	logger logging.Logger,
) (*ContractBindings, error) {
	var (
		err error

		contractDelegationManager *delegationmanager.ContractDelegationManager
		contractStrategyManager   *strategymanager.ContractStrategyManager
		strategyManagerAddr       gethcommon.Address
		avsDirectory              *avsdirectory.ContractIAVSDirectory
		allocationManager         *allocationmanager.ContractAllocationManager
		rewardsCoordinator        *rewardscoordinator.ContractIRewardsCoordinator
	)

	if isZeroAddress(cfg.DelegationManagerAddress) {
		logger.Debug("DelegationManager address not provided, the calls to the contract will not work")
	} else {
		contractDelegationManager, err = delegationmanager.NewContractDelegationManager(cfg.DelegationManagerAddress, client)
		if err != nil {
			return nil, utils.WrapError("Failed to create DelegationManager contract", err)
		}

		strategyManagerAddr, err = contractDelegationManager.StrategyManager(&bind.CallOpts{})
		if err != nil {
			return nil, utils.WrapError("Failed to fetch StrategyManager address", err)
		}
		contractStrategyManager, err = strategymanager.NewContractStrategyManager(strategyManagerAddr, client)
		if err != nil {
			return nil, utils.WrapError("Failed to fetch StrategyManager contract", err)
		}
	}

	if isZeroAddress(cfg.AvsDirectoryAddress) {
		logger.Debug("AVSDirectory address not provided, the calls to the contract will not work")
	} else {
		avsDirectory, err = avsdirectory.NewContractIAVSDirectory(cfg.AvsDirectoryAddress, client)
		if err != nil {
			return nil, utils.WrapError("Failed to fetch AVSDirectory contract", err)
		}
	}

	if isZeroAddress(cfg.AllocationManagerAddress) {
		logger.Debug("AllocationManager address not provided, the calls made to the contract will not work")
	} else {
		allocationManager, err = allocationmanager.NewContractAllocationManager(cfg.AllocationManagerAddress, client)
		if err != nil {
			return nil, utils.WrapError("Failed to fetch AllocationManager contract", err)
		}
	}

	if isZeroAddress(cfg.RewardsCoordinatorAddress) {
		logger.Debug("RewardsCoordinator address not provided, the calls to the contract will not work")
	} else {
		rewardsCoordinator, err = rewardscoordinator.NewContractIRewardsCoordinator(cfg.RewardsCoordinatorAddress, client)
		if err != nil {
			return nil, utils.WrapError("Failed to fetch RewardsCoordinator contract", err)
		}
	}

	return &ContractBindings{
		StrategyManagerAddr:       strategyManagerAddr,
		DelegationManagerAddr:     cfg.DelegationManagerAddress,
		AvsDirectoryAddr:          cfg.AvsDirectoryAddress,
		AllocationManagerAddr:     cfg.AllocationManagerAddress,
		RewardsCoordinatorAddress: cfg.RewardsCoordinatorAddress,
		StrategyManager:           contractStrategyManager,
		DelegationManager:         contractDelegationManager,
		AvsDirectory:              avsDirectory,
		AllocationManager:         allocationManager,
		RewardsCoordinator:        rewardsCoordinator,
	}, nil
}
func isZeroAddress(address gethcommon.Address) bool {
	return address == gethcommon.Address{}
}

// NewEigenlayerContractBindings creates a new ContractBindings struct with the provided contract addresses
// Deprecated: Use NewBindingsFromConfig instead
func NewEigenlayerContractBindings(
	delegationManagerAddr gethcommon.Address,
	avsDirectoryAddr gethcommon.Address,
	allocationManagerAddr gethcommon.Address,
	ethclient eth.HttpBackend,
	logger logging.Logger,
) (*ContractBindings, error) {
	contractDelegationManager, err := delegationmanager.NewContractDelegationManager(delegationManagerAddr, ethclient)
	if err != nil {
		return nil, utils.WrapError("Failed to create DelegationManager contract", err)
	}

	strategyManagerAddr, err := contractDelegationManager.StrategyManager(&bind.CallOpts{})
	if err != nil {
		return nil, utils.WrapError("Failed to fetch StrategyManager address", err)
	}
	contractStrategyManager, err := strategymanager.NewContractStrategyManager(strategyManagerAddr, ethclient)
	if err != nil {
		return nil, utils.WrapError("Failed to fetch StrategyManager contract", err)
	}

	avsDirectory, err := avsdirectory.NewContractIAVSDirectory(avsDirectoryAddr, ethclient)
	if err != nil {
		return nil, utils.WrapError("Failed to fetch AVSDirectory contract", err)
	}

	allocationManager, err := allocationmanager.NewContractAllocationManager(allocationManagerAddr, ethclient)
	if err != nil {
		return nil, utils.WrapError("Failed to fetch AllocationManager  contract", err)
	}

	return &ContractBindings{
		StrategyManagerAddr:   strategyManagerAddr,
		DelegationManagerAddr: delegationManagerAddr,
		AvsDirectoryAddr:      avsDirectoryAddr,
		AllocationManagerAddr: allocationManagerAddr,
		StrategyManager:       contractStrategyManager,
		DelegationManager:     contractDelegationManager,
		AvsDirectory:          avsDirectory,
		AllocationManager:     allocationManager,
	}, nil
}
