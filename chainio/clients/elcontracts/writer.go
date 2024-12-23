package elcontracts

import (
	"context"
	"errors"

	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	allocationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"
	delegationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/DelegationManager"
	avsdirectory "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IAVSDirectory"
	erc20 "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IERC20"
	rewardscoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IRewardsCoordinator"
	strategy "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IStrategy"
	strategymanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StrategyManager"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
)

type Reader interface {
	GetStrategyAndUnderlyingERC20Token(
		ctx context.Context, strategyAddr gethcommon.Address,
	) (*strategy.ContractIStrategy, erc20.ContractIERC20Methods, gethcommon.Address, error)
}

type ChainWriter struct {
	delegationManager   *delegationmanager.ContractDelegationManager
	strategyManager     *strategymanager.ContractStrategyManager
	rewardsCoordinator  *rewardscoordinator.ContractIRewardsCoordinator
	avsDirectory        *avsdirectory.ContractIAVSDirectory
	allocationManager   *allocationmanager.ContractAllocationManager
	strategyManagerAddr gethcommon.Address
	elChainReader       Reader
	ethClient           eth.HttpBackend
	logger              logging.Logger
	txMgr               txmgr.TxManager
}

func NewChainWriter(
	delegationManager *delegationmanager.ContractDelegationManager,
	strategyManager *strategymanager.ContractStrategyManager,
	rewardsCoordinator *rewardscoordinator.ContractIRewardsCoordinator,
	avsDirectory *avsdirectory.ContractIAVSDirectory,
	allocationManager *allocationmanager.ContractAllocationManager,
	strategyManagerAddr gethcommon.Address,
	elChainReader Reader,
	ethClient eth.HttpBackend,
	logger logging.Logger,
	eigenMetrics metrics.Metrics,
	txMgr txmgr.TxManager,
) *ChainWriter {
	logger = logger.With(logging.ComponentKey, "elcontracts/writer")

	return &ChainWriter{
		delegationManager:   delegationManager,
		strategyManager:     strategyManager,
		strategyManagerAddr: strategyManagerAddr,
		rewardsCoordinator:  rewardsCoordinator,
		avsDirectory:        avsDirectory,
		allocationManager:   allocationManager,
		elChainReader:       elChainReader,
		logger:              logger,
		ethClient:           ethClient,
		txMgr:               txMgr,
	}
}

// BuildELChainWriter builds an ChainWriter instance.
// Deprecated: Use NewWriterFromConfig instead.
func BuildELChainWriter(
	delegationManagerAddr gethcommon.Address,
	avsDirectoryAddr gethcommon.Address,
	allocationManagerAddr gethcommon.Address,
	ethClient eth.HttpBackend,
	logger logging.Logger,
	eigenMetrics metrics.Metrics,
	txMgr txmgr.TxManager,
) (*ChainWriter, error) {
	elContractBindings, err := NewEigenlayerContractBindings(
		delegationManagerAddr,
		avsDirectoryAddr,
		allocationManagerAddr,
		ethClient,
		logger,
	)
	if err != nil {
		return nil, err
	}
	elChainReader := NewChainReader(
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.AvsDirectory,
		elContractBindings.AllocationManager,
		elContractBindings.RewardsCoordinator,
		logger,
		ethClient,
	)
	return NewChainWriter(
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.RewardsCoordinator,
		elContractBindings.AvsDirectory,
		elContractBindings.AllocationManager,
		elContractBindings.StrategyManagerAddr,
		elChainReader,
		ethClient,
		logger,
		eigenMetrics,
		txMgr,
	), nil
}

func NewWriterFromConfig(
	cfg Config,
	ethClient eth.HttpBackend,
	logger logging.Logger,
	eigenMetrics metrics.Metrics,
	txMgr txmgr.TxManager,
) (*ChainWriter, error) {
	elContractBindings, err := NewBindingsFromConfig(
		cfg,
		ethClient,
		logger,
	)
	if err != nil {
		return nil, err
	}
	elChainReader := NewChainReader(
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.AvsDirectory,
		elContractBindings.AllocationManager,
		elContractBindings.RewardsCoordinator,
		logger,
		ethClient,
	)
	return NewChainWriter(
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.RewardsCoordinator,
		elContractBindings.AvsDirectory,
		elContractBindings.AllocationManager,
		elContractBindings.StrategyManagerAddr,
		elChainReader,
		ethClient,
		logger,
		eigenMetrics,
		txMgr,
	), nil
}

func (w *ChainWriter) RegisterAsOperator(
	ctx context.Context,
	operator types.Operator,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.delegationManager == nil {
		return nil, errors.New("DelegationManager contract not provided")
	}

	w.logger.Infof("registering operator %s to EigenLayer", operator.Address)

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.delegationManager.RegisterAsOperator(noSendTxOpts, gethcommon.HexToAddress(operator.DelegationApproverAddress), operator.AllocationDelay, operator.MetadataUrl)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Info("tx successfully included", "txHash", receipt.TxHash.String())

	return receipt, nil
}

func (w *ChainWriter) UpdateOperatorDetails(
	ctx context.Context,
	operator types.Operator,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.delegationManager == nil {
		return nil, errors.New("DelegationManager contract not provided")
	}

	w.logger.Infof("updating operator details of operator %s to EigenLayer", operator.Address)

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}

	operatorAddress := gethcommon.HexToAddress(operator.Address)
	delegationApprover := gethcommon.HexToAddress(operator.DelegationApproverAddress)

	tx, err := w.delegationManager.ModifyOperatorDetails(noSendTxOpts, operatorAddress, delegationApprover)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Info(
		"successfully updated operator details",
		"txHash",
		receipt.TxHash.String(),
		"operator",
		operator.Address,
	)

	return receipt, nil
}

func (w *ChainWriter) UpdateMetadataURI(ctx context.Context, operator types.Operator, uri string, waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.delegationManager == nil {
		return nil, errors.New("DelegationManager contract not provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}

	operatorAddress := gethcommon.HexToAddress(operator.Address)

	tx, err := w.delegationManager.UpdateOperatorMetadataURI(noSendTxOpts, operatorAddress, uri)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Info(
		"successfully updated operator metadata uri",
		"txHash",
		receipt.TxHash.String(),
	)

	return receipt, nil
}

func (w *ChainWriter) DepositERC20IntoStrategy(
	ctx context.Context,
	strategyAddr gethcommon.Address,
	amount *big.Int,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.strategyManager == nil {
		return nil, errors.New("StrategyManager contract not provided")
	}

	w.logger.Infof("depositing %s tokens into strategy %s", amount.String(), strategyAddr)
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	_, underlyingTokenContract, underlyingTokenAddr, err := w.elChainReader.GetStrategyAndUnderlyingERC20Token(
		ctx,
		strategyAddr,
	)
	if err != nil {
		return nil, err
	}

	tx, err := underlyingTokenContract.Approve(noSendTxOpts, w.strategyManagerAddr, amount)
	if err != nil {
		return nil, errors.Join(errors.New("failed to approve token transfer"), err)
	}
	_, err = w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}

	tx, err = w.strategyManager.DepositIntoStrategy(noSendTxOpts, strategyAddr, underlyingTokenAddr, amount)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}

	w.logger.Infof("deposited %s into strategy %s", amount.String(), strategyAddr)
	return receipt, nil
}

func (w *ChainWriter) SetClaimerFor(
	ctx context.Context,
	claimer gethcommon.Address,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.rewardsCoordinator == nil {
		return nil, errors.New("RewardsCoordinator contract not provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}

	tx, err := w.rewardsCoordinator.SetClaimerFor(noSendTxOpts, claimer)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

func (w *ChainWriter) ProcessClaim(
	ctx context.Context,
	claim rewardscoordinator.IRewardsCoordinatorTypesRewardsMerkleClaim,
	earnerAddress gethcommon.Address,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.rewardsCoordinator == nil {
		return nil, errors.New("RewardsCoordinator contract not provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	tx, err := w.rewardsCoordinator.ProcessClaim(noSendTxOpts, claim, earnerAddress)
	if err != nil {
		return nil, utils.WrapError("failed to create ProcessClaim tx", err)
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

func (w *ChainWriter) SetOperatorAVSSplit(
	ctx context.Context,
	operator gethcommon.Address,
	avs gethcommon.Address,
	split uint16,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.rewardsCoordinator == nil {
		return nil, errors.New("RewardsCoordinator contract not provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	tx, err := w.rewardsCoordinator.SetOperatorAVSSplit(noSendTxOpts, operator, avs, split)
	if err != nil {
		return nil, utils.WrapError("failed to create SetOperatorAVSSplit tx", err)
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

func (w *ChainWriter) SetOperatorPISplit(
	ctx context.Context,
	operator gethcommon.Address,
	split uint16,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.rewardsCoordinator == nil {
		return nil, errors.New("RewardsCoordinator contract not provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	tx, err := w.rewardsCoordinator.SetOperatorPISplit(noSendTxOpts, operator, split)
	if err != nil {
		return nil, utils.WrapError("failed to create SetOperatorAVSSplit tx", err)
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

func (w *ChainWriter) ProcessClaims(
	ctx context.Context,
	claims []rewardscoordinator.IRewardsCoordinatorTypesRewardsMerkleClaim,
	earnerAddress gethcommon.Address,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.rewardsCoordinator == nil {
		return nil, errors.New("RewardsCoordinator contract not provided")
	}

	if len(claims) == 0 {
		return nil, errors.New("claims is empty, at least one claim must be provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	tx, err := w.rewardsCoordinator.ProcessClaims(noSendTxOpts, claims, earnerAddress)
	if err != nil {
		return nil, utils.WrapError("failed to create ProcessClaims tx", err)
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}
