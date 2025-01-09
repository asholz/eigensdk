package elcontracts_test

import (
	"context"
	"math/big"
	"os"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	allocationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"
	rewardscoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IRewardsCoordinator"
	strategy "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IStrategy"
	mockerc20 "github.com/Layr-Labs/eigensdk-go/contracts/bindings/MockERC20"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/testutils/testclients"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegisterOperator(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)
	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)
	anvilWsEndpoint, err := anvilC.Endpoint(context.Background(), "ws")
	require.NoError(t, err)
	logger := logging.NewTextSLogger(os.Stdout, &logging.SLoggerOptions{Level: testConfig.LogLevel})

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)
	require.NoError(t, err)

	chainioConfig := clients.BuildAllConfig{
		EthHttpUrl:                 anvilHttpEndpoint,
		EthWsUrl:                   anvilWsEndpoint,
		RegistryCoordinatorAddr:    contractAddrs.RegistryCoordinator.String(),
		OperatorStateRetrieverAddr: contractAddrs.OperatorStateRetriever.String(),
		AvsName:                    "exampleAvs",
		PromMetricsIpPortAddress:   ":9090",
	}

	t.Run("register as an operator", func(t *testing.T) {
		// Fund the new address with 5 ether
		fundedAccount := "0x408EfD9C90d59298A9b32F4441aC9Df6A2d8C3E1"
		fundedPrivateKeyHex := "3339854a8622364bcd5650fa92eac82d5dccf04089f5575a761c9b7d3c405b1c"
		richPrivateKeyHex := "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
		code, _, err := anvilC.Exec(
			context.Background(),
			[]string{"cast",
				"send",
				"0x408EfD9C90d59298A9b32F4441aC9Df6A2d8C3E1",
				"--value",
				"5ether",
				"--private-key",
				richPrivateKeyHex,
			},
		)
		assert.NoError(t, err)
		assert.Equal(t, 0, code)

		ecdsaPrivateKey, err := crypto.HexToECDSA(fundedPrivateKeyHex)
		require.NoError(t, err)

		clients, err := clients.BuildAll(
			chainioConfig,
			ecdsaPrivateKey,
			logger,
		)
		require.NoError(t, err)

		operator :=
			types.Operator{
				Address:                   fundedAccount,
				DelegationApproverAddress: "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				MetadataUrl:               "https://madhur-test-public.s3.us-east-2.amazonaws.com/metadata.json",
			}

		receipt, err := clients.ElChainWriter.RegisterAsOperator(context.Background(), operator, true)
		assert.NoError(t, err)
		assert.True(t, receipt.Status == 1)
	})

	t.Run("register as an operator already registered", func(t *testing.T) {
		operatorAddress := "0x408EfD9C90d59298A9b32F4441aC9Df6A2d8C3E1"
		operatorPrivateKeyHex := "3339854a8622364bcd5650fa92eac82d5dccf04089f5575a761c9b7d3c405b1c"

		ecdsaPrivateKey, err := crypto.HexToECDSA(operatorPrivateKeyHex)
		require.NoError(t, err)

		clients, err := clients.BuildAll(
			chainioConfig,
			ecdsaPrivateKey,
			logger,
		)
		require.NoError(t, err)

		operator :=
			types.Operator{
				Address:                   operatorAddress,
				DelegationApproverAddress: "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				MetadataUrl:               "https://madhur-test-public.s3.us-east-2.amazonaws.com/metadata.json",
			}

		_, err = clients.ElChainWriter.RegisterAsOperator(context.Background(), operator, true)
		assert.Error(t, err)
	})
}

func TestChainWriter(t *testing.T) {
	clients, anvilHttpEndpoint := testclients.BuildTestClients(t)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	t.Run("update operator details", func(t *testing.T) {
		walletModified, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
		assert.NoError(t, err)
		walletModifiedAddress := crypto.PubkeyToAddress(walletModified.PublicKey)

		operatorModified := types.Operator{
			Address:                   walletModifiedAddress.Hex(),
			DelegationApproverAddress: walletModifiedAddress.Hex(),
			MetadataUrl:               "eigensdk-go",
		}

		receipt, err := clients.ElChainWriter.UpdateOperatorDetails(context.Background(), operatorModified, true)
		assert.NoError(t, err)
		assert.True(t, receipt.Status == 1)
	})

	t.Run("update metadata URI", func(t *testing.T) {
		walletModified, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
		assert.NoError(t, err)
		walletModifiedAddress := crypto.PubkeyToAddress(walletModified.PublicKey)
		receipt, err := clients.ElChainWriter.UpdateMetadataURI(
			context.Background(),
			walletModifiedAddress,
			"https://0.0.0.0",
			true,
		)
		assert.NoError(t, err)
		assert.True(t, receipt.Status == 1)
	})

	t.Run("deposit ERC20 into strategy", func(t *testing.T) {
		amount := big.NewInt(1)
		receipt, err := clients.ElChainWriter.DepositERC20IntoStrategy(
			context.Background(),
			contractAddrs.Erc20MockStrategy,
			amount,
			true,
		)
		assert.NoError(t, err)
		assert.True(t, receipt.Status == 1)
	})
}

const SUCCESS_STATUS = uint64(1)

func TestSetClaimerFor(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)
	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)

	privateKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
	config := elcontracts.Config{
		DelegationManagerAddress:  contractAddrs.DelegationManager,
		RewardsCoordinatorAddress: rewardsCoordinatorAddr,
	}

	// Create ChainWriter
	chainWriter, err := newTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	require.NoError(t, err)

	waitForReceipt := true
	claimer := common.HexToAddress("0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6")

	// call SetClaimerFor
	receipt, err := chainWriter.SetClaimerFor(context.Background(), claimer, waitForReceipt)
	require.NoError(t, err)
	require.True(t, receipt.Status == SUCCESS_STATUS)
}

func TestSetOperatorPISplit(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
	config := elcontracts.Config{
		DelegationManagerAddress:  contractAddrs.DelegationManager,
		RewardsCoordinatorAddress: rewardsCoordinatorAddr,
	}

	privateKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	operatorAddr := common.HexToAddress("f39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	waitForReceipt := true

	activationDelay := uint32(0)
	// Set activation delay to zero so that the new PI split can be retrieved immediately after setting it
	receipt, err := setTestRewardsCoordinatorActivationDelay(anvilHttpEndpoint, privateKeyHex, activationDelay)
	require.NoError(t, err)
	require.True(t, receipt.Status == SUCCESS_STATUS)

	// Create ChainWriter
	chainWriter, err := newTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	require.NoError(t, err)

	chainReader, err := newTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	startingSplit, err := chainReader.GetOperatorPISplit(context.Background(), operatorAddr)
	require.NoError(t, err)

	newSplit := startingSplit * 2
	// Set a new operator PI split
	receipt, err = chainWriter.SetOperatorPISplit(context.Background(), operatorAddr, newSplit, waitForReceipt)
	require.NoError(t, err)
	require.True(t, receipt.Status == SUCCESS_STATUS)

	// Retrieve the operator PI split to check it has been set
	updatedSplit, err := chainReader.GetOperatorPISplit(context.Background(), operatorAddr)
	require.NoError(t, err)
	require.Equal(t, newSplit, updatedSplit)
}

func TestSetOperatorAVSSplit(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
	avsAddr := contractAddrs.ServiceManager
	config := elcontracts.Config{
		DelegationManagerAddress:  contractAddrs.DelegationManager,
		RewardsCoordinatorAddress: rewardsCoordinatorAddr,
	}

	privateKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	operatorAddr := common.HexToAddress("f39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	waitForReceipt := true
	activationDelay := uint32(0)

	// Set activation delay to zero so that the new PI split can be retrieved immediately after setting it
	receipt, err := setTestRewardsCoordinatorActivationDelay(anvilHttpEndpoint, privateKeyHex, activationDelay)
	require.NoError(t, err)
	require.True(t, receipt.Status == SUCCESS_STATUS)

	// Create ChainWriter
	chainWriter, err := newTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	require.NoError(t, err)

	chainReader, err := newTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	startingSplit, err := chainReader.GetOperatorAVSSplit(context.Background(), operatorAddr, avsAddr)
	require.NoError(t, err)

	newSplit := startingSplit * 2
	// Set a new operator AVS split
	receipt, err = chainWriter.SetOperatorAVSSplit(
		context.Background(),
		operatorAddr,
		avsAddr,
		newSplit,
		waitForReceipt,
	)
	require.NoError(t, err)
	require.True(t, receipt.Status == SUCCESS_STATUS)

	// Retrieve the operator AVS split to check it has been set
	updatedSplit, err := chainReader.GetOperatorAVSSplit(context.Background(), operatorAddr, avsAddr)
	require.NoError(t, err)
	require.Equal(t, newSplit, updatedSplit)
}

func TestSetAllocationDelay(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
	config := elcontracts.Config{
		DelegationManagerAddress:  contractAddrs.DelegationManager,
		RewardsCoordinatorAddress: rewardsCoordinatorAddr,
	}

	privateKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	operatorAddr := common.HexToAddress("f39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	waitForReceipt := true

	// Create ChainWriter
	chainWriter, err := newTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	require.NoError(t, err)

	delay := uint32(10)
	receipt, err := chainWriter.SetAllocationDelay(context.Background(), operatorAddr, delay, waitForReceipt)
	require.NoError(t, err)
	require.True(t, receipt.Status == SUCCESS_STATUS)
}

func TestSetAndRemovePermission(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)
	permissionControllerAddr := common.HexToAddress("0x610178dA211FEF7D417bC0e6FeD39F05609AD788")

	privateKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	config := elcontracts.Config{
		DelegationManagerAddress:     contractAddrs.DelegationManager,
		PermissionsControllerAddress: permissionControllerAddr,
	}
	chainWriter, err := newTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	require.NoError(t, err)
	chainReader, err := newTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	accountAddress := common.HexToAddress("f39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	appointeeAddress := common.HexToAddress("009440d62dc85c73dbf889b7ad1f4da8b231d2ef")
	target := common.HexToAddress("14dC79964da2C08b23698B3D3cc7Ca32193d9955")
	selector := [4]byte{0, 1, 2, 3}
	waitForReceipt := true

	setPermissionRequest := elcontracts.SetPermissionRequest{
		AccountAddress:   accountAddress,
		AppointeeAddress: appointeeAddress,
		Target:           target,
		Selector:         selector,
		WaitForReceipt:   waitForReceipt,
	}

	removePermissionRequest := elcontracts.RemovePermissionRequest{
		AccountAddress:   accountAddress,
		AppointeeAddress: appointeeAddress,
		Target:           target,
		Selector:         selector,
		WaitForReceipt:   waitForReceipt,
	}
	receipt, err := chainWriter.SetPermission(context.Background(), setPermissionRequest)
	require.NoError(t, err)
	require.True(t, receipt.Status == SUCCESS_STATUS)

	canCall, err := chainReader.CanCall(context.Background(), accountAddress, appointeeAddress, target, selector)
	require.NoError(t, err)
	require.True(t, canCall)

	receipt, err = chainWriter.RemovePermission(context.Background(), removePermissionRequest)
	require.NoError(t, err)
	require.True(t, receipt.Status == SUCCESS_STATUS)

	canCall, err = chainReader.CanCall(context.Background(), accountAddress, appointeeAddress, target, selector)
	require.NoError(t, err)
	require.False(t, canCall)
}

// Sets the testing RewardsCoordinator contract's activationDelay.
// This is useful to test ChainWriter setter functions that depend on activationDelay.
func setTestRewardsCoordinatorActivationDelay(
	httpEndpoint string,
	privateKeyHex string,
	activationDelay uint32,
) (*gethtypes.Receipt, error) {
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(httpEndpoint)
	rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
	ethHttpClient, err := ethclient.Dial(httpEndpoint)
	if err != nil {
		return nil, utils.WrapError("Failed to create eth client", err)
	}

	rewardsCoordinator, err := rewardscoordinator.NewContractIRewardsCoordinator(rewardsCoordinatorAddr, ethHttpClient)
	if err != nil {
		return nil, utils.WrapError("Failed to create rewards coordinator", err)
	}

	txManager, err := newTestTxManager(httpEndpoint, privateKeyHex)
	if err != nil {
		return nil, utils.WrapError("Failed to create tx manager", err)
	}

	noSendOpts, err := txManager.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("Failed to get NoSend tx opts", err)
	}

	tx, err := rewardsCoordinator.SetActivationDelay(noSendOpts, activationDelay)
	if err != nil {
		return nil, utils.WrapError("Failed to create SetActivationDelay tx", err)
	}

	receipt, err := txManager.Send(context.Background(), tx, true)
	if err != nil {
		return nil, utils.WrapError("Failed to send SetActivationDelay tx", err)
	}
	return receipt, err
}

// TODO: consider moving this and the other utilities below to testutils
func newTestTxManager(httpEndpoint string, privateKeyHex string) (*txmgr.SimpleTxManager, error) {
	testConfig := testutils.GetDefaultTestConfig()
	ethHttpClient, err := ethclient.Dial(httpEndpoint)
	if err != nil {
		return nil, utils.WrapError("Failed to create eth client", err)
	}

	chainid, err := ethHttpClient.ChainID(context.Background())
	if err != nil {
		return nil, utils.WrapError("Failed to retrieve chain id", err)
	}
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, utils.WrapError("Failed to convert hex string to private key", err)
	}
	signerV2, addr, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: privateKey}, chainid)
	if err != nil {
		return nil, utils.WrapError("Failed to create signer", err)
	}

	logger := logging.NewTextSLogger(os.Stdout, &logging.SLoggerOptions{Level: testConfig.LogLevel})

	pkWallet, err := wallet.NewPrivateKeyWallet(ethHttpClient, signerV2, addr, logger)
	if err != nil {
		return nil, utils.WrapError("Failed to create wallet", err)
	}

	txManager := txmgr.NewSimpleTxManager(pkWallet, ethHttpClient, logger, addr)
	return txManager, nil
}

// Creates a testing ChainWriter from an httpEndpoint, private key and config.
// This is needed because the existing testclients.BuildTestClients returns a
// ChainWriter with a null rewardsCoordinator, which is required for some of the tests.
func newTestChainWriterFromConfig(
	httpEndpoint string,
	privateKeyHex string,
	config elcontracts.Config,
) (*elcontracts.ChainWriter, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, utils.WrapError("Failed convert hex string to ecdsa private key", err)
	}
	testConfig := testutils.GetDefaultTestConfig()
	logger := logging.NewTextSLogger(os.Stdout, &logging.SLoggerOptions{Level: testConfig.LogLevel})
	ethHttpClient, err := ethclient.Dial(httpEndpoint)
	if err != nil {
		return nil, utils.WrapError("Failed to create eth client", err)
	}
	chainid, err := ethHttpClient.ChainID(context.Background())
	if err != nil {
		return nil, utils.WrapError("Failed to get chain id", err)
	}
	promReg := prometheus.NewRegistry()
	eigenMetrics := metrics.NewEigenMetrics("", "", promReg, logger)
	signerV2, addr, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: privateKey}, chainid)
	if err != nil {
		return nil, utils.WrapError("Failed to create the signer from the given config", err)
	}

	pkWallet, err := wallet.NewPrivateKeyWallet(ethHttpClient, signerV2, addr, logger)
	if err != nil {
		return nil, utils.WrapError("Failed to create wallet", err)
	}
	txManager := txmgr.NewSimpleTxManager(pkWallet, ethHttpClient, logger, addr)
	testWriter, err := elcontracts.NewWriterFromConfig(
		config,
		ethHttpClient,
		logger,
		eigenMetrics,
		txManager,
	)
	if err != nil {
		return nil, err
	}
	return testWriter, nil
}

func TestModifyAllocations(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	operatorAddr := common.HexToAddress("f39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	privateKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	config := elcontracts.Config{
		DelegationManagerAddress: contractAddrs.DelegationManager,
	}

	chainWriter, err := newTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	require.NoError(t, err)

	chainReader, err := newTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	strategyAddr := contractAddrs.Erc20MockStrategy
	avsAddr := common.HexToAddress("f39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	operatorSetId := uint32(1)

	waitForReceipt := true
	delay := uint32(1)
	// The allocation delay must be initialized before modifying the allocations
	receipt, err := chainWriter.SetAllocationDelay(context.Background(), operatorAddr, delay, waitForReceipt)
	require.NoError(t, err)
	require.True(t, receipt.Status == SUCCESS_STATUS)

	allocationConfigurationDelay := 1200
	// Advance the chain by the required number of blocks
	// (ALLOCATION_CONFIGURATION_DELAY) to apply the allocation delay
	testutils.AdvanceChainByNBlocksExecInContainer(context.Background(), allocationConfigurationDelay+1, anvilC)

	// Retrieve the allocation delay so that the delay is applied
	_, err = chainReader.GetAllocationDelay(context.Background(), operatorAddr)
	require.NoError(t, err)

	clients, _ := newTestClients(anvilHttpEndpoint, privateKeyHex)
	err = createOperatorSet(clients, avsAddr, operatorSetId, strategyAddr)
	require.NoError(t, err)

	operatorSet := allocationmanager.OperatorSet{
		Avs: avsAddr,
		Id:  operatorSetId,
	}
	newAllocation := uint64(100)
	allocateParams := []allocationmanager.IAllocationManagerTypesAllocateParams{
		{
			OperatorSet:   operatorSet,
			Strategies:    []common.Address{strategyAddr},
			NewMagnitudes: []uint64{newAllocation},
		},
	}

	receipt, err = chainWriter.ModifyAllocations(context.Background(), operatorAddr, allocateParams, waitForReceipt)
	require.NoError(t, err)
	require.True(t, receipt.Status == SUCCESS_STATUS)

	// Check that the new allocation is pending
	allocationInfo, err := chainReader.GetAllocationInfo(context.Background(), operatorAddr, strategyAddr)
	require.NoError(t, err)
	pendingDiff := allocationInfo[0].PendingDiff
	require.Equal(t, big.NewInt(int64(newAllocation)), pendingDiff)

	// Retrieve the allocation delay and advance the chain
	allocationDelay, err := chainReader.GetAllocationDelay(context.Background(), operatorAddr)
	require.NoError(t, err)
	testutils.AdvanceChainByNBlocksExecInContainer(context.Background(), int(allocationDelay), anvilC)

	// Check the new allocation has been updated after the delay
	allocationInfo, err = chainReader.GetAllocationInfo(context.Background(), operatorAddr, strategyAddr)
	require.NoError(t, err)

	currentMagnitude := allocationInfo[0].CurrentMagnitude
	require.Equal(t, big.NewInt(int64(newAllocation)), currentMagnitude)
}

func TestAddAndRemovePendingAdmin(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)
	// TODO: unhardcode permissionControllerAddr
	permissionControllerAddr := common.HexToAddress("0x610178dA211FEF7D417bC0e6FeD39F05609AD788")

	operatorAddr := common.HexToAddress("f39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	privateKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	config := elcontracts.Config{
		DelegationManagerAddress:     contractAddrs.DelegationManager,
		PermissionsControllerAddress: permissionControllerAddr,
	}
	chainWriter, err := newTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	require.NoError(t, err)
	chainReader, err := newTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	pendingAdmin := common.HexToAddress("009440d62dc85c73dbf889b7ad1f4da8b231d2ef")
	request := elcontracts.AddPendingAdminRequest{
		AccountAddress: operatorAddr,
		AdminAddress:   pendingAdmin,
		WaitForReceipt: true,
	}

	removePendingAdminRequest := elcontracts.RemovePendingAdminRequest{
		AccountAddress: operatorAddr,
		AdminAddress:   pendingAdmin,
		WaitForReceipt: true,
	}
	t.Run("add pending admin", func(t *testing.T) {
		receipt, err := chainWriter.AddPendingAdmin(context.Background(), request)
		require.NoError(t, err)
		require.True(t, receipt.Status == SUCCESS_STATUS)

		isPendingAdmin, err := chainReader.IsPendingAdmin(context.Background(), operatorAddr, pendingAdmin)
		require.NoError(t, err)
		require.True(t, isPendingAdmin)
	})
	t.Run("remove pending admin", func(t *testing.T) {
		receipt, err := chainWriter.RemovePendingAdmin(context.Background(), removePendingAdminRequest)
		require.NoError(t, err)
		require.True(t, receipt.Status == SUCCESS_STATUS)

		isPendingAdmin, err := chainReader.IsPendingAdmin(context.Background(), operatorAddr, pendingAdmin)
		require.NoError(t, err)
		require.False(t, isPendingAdmin)
	})
}

func TestAcceptAdmin(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)
	// TODO: unhardcode permissionControllerAddr
	permissionControllerAddr := common.HexToAddress("0x610178dA211FEF7D417bC0e6FeD39F05609AD788")

	accountAddr := common.HexToAddress("f39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	accountPrivateKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	config := elcontracts.Config{
		DelegationManagerAddress:     contractAddrs.DelegationManager,
		PermissionsControllerAddress: permissionControllerAddr,
	}
	accountChainWriter, err := newTestChainWriterFromConfig(anvilHttpEndpoint, accountPrivateKeyHex, config)
	require.NoError(t, err)

	pendingAdminPrivateKeyHex := "4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356"
	adminChainWriter, err := newTestChainWriterFromConfig(anvilHttpEndpoint, pendingAdminPrivateKeyHex, config)
	require.NoError(t, err)

	chainReader, err := newTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	pendingAdminAddr := common.HexToAddress("14dC79964da2C08b23698B3D3cc7Ca32193d9955")
	request := elcontracts.AddPendingAdminRequest{
		AccountAddress: accountAddr,
		AdminAddress:   pendingAdminAddr,
		WaitForReceipt: true,
	}
	receipt, err := accountChainWriter.AddPendingAdmin(context.Background(), request)
	require.NoError(t, err)
	require.True(t, receipt.Status == SUCCESS_STATUS)

	acceptAdminRequest := elcontracts.AcceptAdminRequest{
		AccountAddress: accountAddr,
		WaitForReceipt: true,
	}
	t.Run("accept admin", func(t *testing.T) {
		receipt, err = adminChainWriter.AcceptAdmin(context.Background(), acceptAdminRequest)
		require.NoError(t, err)
		require.True(t, receipt.Status == SUCCESS_STATUS)

		isAdmin, err := chainReader.IsAdmin(context.Background(), accountAddr, pendingAdminAddr)
		require.NoError(t, err)
		require.True(t, isAdmin)
	})
}

func TestRemoveAdmin(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)
	// TODO: unhardcode permissionControllerAddr
	permissionControllerAddr := common.HexToAddress("0x610178dA211FEF7D417bC0e6FeD39F05609AD788")

	accountAddr := common.HexToAddress("f39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	accountPrivateKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	config := elcontracts.Config{
		DelegationManagerAddress:     contractAddrs.DelegationManager,
		PermissionsControllerAddress: permissionControllerAddr,
	}
	accountChainWriter, err := newTestChainWriterFromConfig(anvilHttpEndpoint, accountPrivateKeyHex, config)
	require.NoError(t, err)

	// Adding two admins and removing one. Cannot remove the last admin, so one must remain
	admin1 := common.HexToAddress("14dC79964da2C08b23698B3D3cc7Ca32193d9955")
	admin1PrivateKeyHex := "4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356"

	admin2 := common.HexToAddress("23618e81E3f5cdF7f54C3d65f7FBc0aBf5B21E8f")
	admin2PrivateKeyHex := "dbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97"

	admin1ChainWriter, err := newTestChainWriterFromConfig(anvilHttpEndpoint, admin1PrivateKeyHex, config)
	require.NoError(t, err)

	admin2ChainWriter, err := newTestChainWriterFromConfig(anvilHttpEndpoint, admin2PrivateKeyHex, config)
	require.NoError(t, err)

	chainReader, err := newTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	addAdmin1Request := elcontracts.AddPendingAdminRequest{
		AccountAddress: accountAddr,
		AdminAddress:   admin1,
		WaitForReceipt: true,
	}
	addAdmin2Request := elcontracts.AddPendingAdminRequest{
		AccountAddress: accountAddr,
		AdminAddress:   admin2,
		WaitForReceipt: true,
	}
	acceptAdminRequest := elcontracts.AcceptAdminRequest{
		AccountAddress: accountAddr,
		WaitForReceipt: true,
	}
	// Add and accept admin 1
	receipt, err := accountChainWriter.AddPendingAdmin(context.Background(), addAdmin1Request)
	require.NoError(t, err)
	require.True(t, receipt.Status == SUCCESS_STATUS)

	receipt, err = admin1ChainWriter.AcceptAdmin(context.Background(), acceptAdminRequest)
	require.NoError(t, err)
	require.True(t, receipt.Status == SUCCESS_STATUS)

	// Add and accept admin 2
	receipt, err = admin1ChainWriter.AddPendingAdmin(context.Background(), addAdmin2Request)
	require.NoError(t, err)
	require.True(t, receipt.Status == SUCCESS_STATUS)

	receipt, err = admin2ChainWriter.AcceptAdmin(context.Background(), acceptAdminRequest)
	require.NoError(t, err)
	require.True(t, receipt.Status == SUCCESS_STATUS)

	removeAdminRequest := elcontracts.RemoveAdminRequest{
		AccountAddress: accountAddr,
		AdminAddress:   admin2,
		WaitForReceipt: true,
	}
	t.Run("remove admin 2", func(t *testing.T) {
		receipt, err = admin1ChainWriter.RemoveAdmin(context.Background(), removeAdminRequest)
		require.NoError(t, err)
		require.True(t, receipt.Status == SUCCESS_STATUS)

		isAdmin, err := chainReader.IsAdmin(context.Background(), accountAddr, admin2)
		require.NoError(t, err)
		require.False(t, isAdmin)
	})
}

func TestProcessClaim(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)
	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)

	privateKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
	config := elcontracts.Config{
		DelegationManagerAddress:  contractAddrs.DelegationManager,
		RewardsCoordinatorAddress: rewardsCoordinatorAddr,
	}

	// Create ChainWriter
	chainWriter, err := newTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	require.NoError(t, err)

	chainReader, err := newTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	activationDelay := uint32(0)
	// Set activation delay to zero so that the earnings can be claimed right after submitting the root
	receipt, err := setTestRewardsCoordinatorActivationDelay(anvilHttpEndpoint, privateKeyHex, activationDelay)
	require.NoError(t, err)
	require.True(t, receipt.Status == SUCCESS_STATUS)

	waitForReceipt := true
	cumulativeEarnings := int64(42)
	earner := common.HexToAddress("0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6")
	_, claim, err := newTestClaim(chainReader, anvilHttpEndpoint, cumulativeEarnings, privateKeyHex)
	require.NoError(t, err)

	receipt, err = chainWriter.ProcessClaim(context.Background(), *claim, earner, waitForReceipt)
	require.NoError(t, err)
	require.True(t, receipt.Status == SUCCESS_STATUS)
}

func TestProcessClaims(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)
	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)

	privateKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
	config := elcontracts.Config{
		DelegationManagerAddress:  contractAddrs.DelegationManager,
		RewardsCoordinatorAddress: rewardsCoordinatorAddr,
	}

	// Create ChainWriter
	chainWriter, err := newTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	require.NoError(t, err)

	chainReader, err := newTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	activationDelay := uint32(0)
	// Set activation delay to zero so that the earnings can be claimed right after submitting the root
	receipt, err := setTestRewardsCoordinatorActivationDelay(anvilHttpEndpoint, privateKeyHex, activationDelay)
	require.NoError(t, err)
	require.True(t, receipt.Status == SUCCESS_STATUS)

	earner := common.HexToAddress("0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6")

	waitForReceipt := true
	cumulativeEarnings1 := int64(42)
	cumulativeEarnings2 := int64(4256)

	// Generate 2 claims
	_, claim1, err := newTestClaim(chainReader, anvilHttpEndpoint, cumulativeEarnings1, privateKeyHex)
	require.NoError(t, err)

	_, claim2, err := newTestClaim(chainReader, anvilHttpEndpoint, cumulativeEarnings2, privateKeyHex)
	require.NoError(t, err)
	claims := []rewardscoordinator.IRewardsCoordinatorTypesRewardsMerkleClaim{
		*claim1, *claim2,
	}
	receipt, err = chainWriter.ProcessClaims(context.Background(), claims, earner, waitForReceipt)
	require.NoError(t, err)
	require.True(t, receipt.Status == SUCCESS_STATUS)
}

func newTestClaim(
	chainReader *elcontracts.ChainReader,
	httpEndpoint string,
	cumulativeEarnings int64,
	privateKeyHex string,
) ([32]byte, *rewardscoordinator.IRewardsCoordinatorTypesRewardsMerkleClaim, error) {
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(httpEndpoint)
	mockStrategyAddr := contractAddrs.Erc20MockStrategy
	rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
	emptyRoot := [32]byte{}
	waitForReceipt := true
	ethClient, err := ethclient.Dial(httpEndpoint)

	if err != nil {
		return emptyRoot, nil, utils.WrapError("Failed to create eth client", err)
	}
	txManager, err := newTestTxManager(httpEndpoint, privateKeyHex)
	if err != nil {
		return emptyRoot, nil, utils.WrapError("Failed to create tx manager", err)
	}
	contractStrategy, err := strategy.NewContractIStrategy(mockStrategyAddr, ethClient)
	if err != nil {
		return emptyRoot, nil, utils.WrapError("Failed to fetch strategy contract", err)
	}
	tokenAddr, err := contractStrategy.UnderlyingToken(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		return emptyRoot, nil, utils.WrapError("Failed to fetch token address", err)
	}

	token, err := mockerc20.NewContractMockERC20(tokenAddr, ethClient)
	if err != nil {
		return emptyRoot, nil, utils.WrapError("Failed to create token contract", err)
	}

	noSendTxOpts, err := txManager.GetNoSendTxOpts()
	if err != nil {
		return emptyRoot, nil, utils.WrapError("Failed to get NoSend tx opts", err)
	}
	// Mint tokens for the RewardsCoordinator
	tx, err := token.Mint(noSendTxOpts, rewardsCoordinatorAddr, big.NewInt(cumulativeEarnings))
	if err != nil {
		return emptyRoot, nil, utils.WrapError("Failed to create MockERC20.Mint tx", err)
	}
	_, err = txManager.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return emptyRoot, nil, utils.WrapError("Failed to mint tokens for RewardsCoordinator", err)
	}

	// Generate token tree leaf
	// For the tree structure, see
	// https://github.com/Layr-Labs/eigenlayer-contracts/blob/a888a1cd1479438dda4b138245a69177b125a973/docs/core/RewardsCoordinator.md#rewards-merkle-tree-structure
	earnerAddr := common.HexToAddress("f39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	tokenLeaf := rewardscoordinator.IRewardsCoordinatorTypesTokenTreeMerkleLeaf{
		Token:              tokenAddr,
		CumulativeEarnings: big.NewInt(cumulativeEarnings),
	}
	// Hash token tree leaf to get root
	encodedTokenLeaf := []byte{}
	tokenLeafSalt := uint8(1)

	// Write the *big.Int to a 32-byte sized buffer to match the uint256 length
	cumulativeEarningsBytes := [32]byte{}
	tokenLeaf.CumulativeEarnings.FillBytes(cumulativeEarningsBytes[:])

	encodedTokenLeaf = append(encodedTokenLeaf, tokenLeafSalt)
	encodedTokenLeaf = append(encodedTokenLeaf, tokenLeaf.Token.Bytes()...)
	encodedTokenLeaf = append(encodedTokenLeaf, cumulativeEarningsBytes[:]...)

	// Hash token tree leaf to get root
	earnerTokenRoot := crypto.Keccak256(encodedTokenLeaf)

	// Generate earner tree leaf
	earnerLeaf := rewardscoordinator.IRewardsCoordinatorTypesEarnerTreeMerkleLeaf{
		Earner:          earnerAddr,
		EarnerTokenRoot: [32]byte(earnerTokenRoot),
	}
	encodedEarnerLeaf := []byte{}
	earnerLeafSalt := uint8(0)
	encodedEarnerLeaf = append(encodedEarnerLeaf, earnerLeafSalt)
	encodedEarnerLeaf = append(encodedEarnerLeaf, earnerLeaf.Earner.Bytes()...)
	encodedEarnerLeaf = append(encodedEarnerLeaf, earnerTokenRoot...)

	// Hash earner tree leaf to get root
	earnerTreeRoot := crypto.Keccak256(encodedEarnerLeaf)

	// Fetch the next root index from contract
	nextRootIndex, err := chainReader.GetDistributionRootsLength(context.Background())
	if err != nil {
		return emptyRoot, nil, utils.WrapError("Failed to call GetDistributionRootsLength", err)
	}
	tokenLeaves := []rewardscoordinator.IRewardsCoordinatorTypesTokenTreeMerkleLeaf{tokenLeaf}
	// Construct the claim
	claim := rewardscoordinator.IRewardsCoordinatorTypesRewardsMerkleClaim{
		RootIndex:   uint32(nextRootIndex.Uint64()),
		EarnerIndex: 0,
		// Empty proof because leaf == root
		EarnerTreeProof: []byte{},
		EarnerLeaf:      earnerLeaf,
		TokenIndices:    []uint32{0},
		// Empty proof because leaf == root
		TokenTreeProofs: [][]byte{{}},
		TokenLeaves:     tokenLeaves,
	}

	root := earnerTreeRoot
	// Fetch the current timestamp to increase it
	currRewardsCalculationEndTimestamp, err := chainReader.CurrRewardsCalculationEndTimestamp(context.Background())
	if err != nil {
		return emptyRoot, nil, utils.WrapError("Failed to call CurrRewardsCalculationEndTimestamp", err)
	}

	rewardsCoordinator, err := rewardscoordinator.NewContractIRewardsCoordinator(rewardsCoordinatorAddr, ethClient)
	if err != nil {
		return emptyRoot, nil, utils.WrapError("Failed to create rewards coordinator contract", err)
	}

	rewardsUpdater := common.HexToAddress("f39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	tx, err = rewardsCoordinator.SetRewardsUpdater(noSendTxOpts, rewardsUpdater)
	if err != nil {
		return emptyRoot, nil, utils.WrapError("Failed to create SetRewardsUpdater tx", err)
	}

	_, err = txManager.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return emptyRoot, nil, utils.WrapError("Failed to setRewardsUpdate", err)
	}

	tx, err = rewardsCoordinator.SubmitRoot(noSendTxOpts, [32]byte(root), currRewardsCalculationEndTimestamp+1)
	if err != nil {
		return emptyRoot, nil, utils.WrapError("Failed to create SubmitRoot tx", err)
	}

	_, err = txManager.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return emptyRoot, nil, utils.WrapError("Failed to submit root", err)
	}

	return [32]byte(root), &claim, nil
}

func createOperatorSet(
	client *clients.Clients,
	avsAddress common.Address,
	operatorSetId uint32,
	erc20MockStrategyAddr common.Address,
) error {
	allocationManagerAddress := client.EigenlayerContractBindings.AllocationManagerAddr
	allocationManager, err := allocationmanager.NewContractAllocationManager(
		allocationManagerAddress,
		client.EthHttpClient,
	)
	if err != nil {
		return err
	}
	registryCoordinatorAddress := client.AvsRegistryContractBindings.RegistryCoordinatorAddr
	registryCoordinator, err := regcoord.NewContractRegistryCoordinator(
		registryCoordinatorAddress,
		client.EthHttpClient,
	)
	if err != nil {
		return err
	}

	noSendTxOpts, err := client.TxManager.GetNoSendTxOpts()
	if err != nil {
		return err
	}

	tx, err := allocationManager.SetAVSRegistrar(noSendTxOpts, avsAddress, registryCoordinatorAddress)
	if err != nil {
		return err
	}

	waitForReceipt := true

	_, err = client.TxManager.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return err
	}

	tx, err = registryCoordinator.EnableOperatorSets(noSendTxOpts)
	if err != nil {
		return err
	}

	_, err = client.TxManager.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return err
	}

	operatorSetParam := regcoord.IRegistryCoordinatorOperatorSetParam{
		MaxOperatorCount:        10,
		KickBIPsOfOperatorStake: 100,
		KickBIPsOfTotalStake:    1000,
	}
	minimumStake := big.NewInt(0)

	strategyParams := regcoord.IStakeRegistryStrategyParams{
		Strategy:   erc20MockStrategyAddr,
		Multiplier: big.NewInt(1),
	}
	strategyParamsArray := []regcoord.IStakeRegistryStrategyParams{strategyParams}
	lookAheadPeriod := uint32(0)
	tx, err = registryCoordinator.CreateSlashableStakeQuorum(
		noSendTxOpts,
		operatorSetParam,
		minimumStake,
		strategyParamsArray,
		lookAheadPeriod,
	)
	if err != nil {
		return err
	}

	_, err = client.TxManager.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return err
	}

	strategies := []common.Address{erc20MockStrategyAddr}
	operatorSetParams := allocationmanager.IAllocationManagerTypesCreateSetParams{
		OperatorSetId: operatorSetId,
		Strategies:    strategies,
	}
	operatorSetParamsArray := []allocationmanager.IAllocationManagerTypesCreateSetParams{operatorSetParams}
	tx, err = allocationManager.CreateOperatorSets(noSendTxOpts, avsAddress, operatorSetParamsArray)
	if err != nil {
		return err
	}

	_, err = client.TxManager.Send(context.Background(), tx, waitForReceipt)
	return err
}

func newTestClients(httpEndpoint string, privateKeyHex string) (*clients.Clients, error) {
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(httpEndpoint)
	chainioConfig := clients.BuildAllConfig{
		EthHttpUrl:                 httpEndpoint,
		EthWsUrl:                   httpEndpoint,
		RegistryCoordinatorAddr:    contractAddrs.RegistryCoordinator.String(),
		OperatorStateRetrieverAddr: contractAddrs.OperatorStateRetriever.String(),
		AvsName:                    "exampleAvs",
		PromMetricsIpPortAddress:   ":9090",
	}
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}
	testConfig := testutils.GetDefaultTestConfig()
	logger := logging.NewTextSLogger(os.Stdout, &logging.SLoggerOptions{Level: testConfig.LogLevel})

	testClients, err := clients.BuildAll(
		chainioConfig,
		privateKey,
		logger,
	)
	if err != nil {
		return nil, err
	}
	return testClients, nil
}

func newTestChainReaderFromConfig(
	httpEndpoint string,
	config elcontracts.Config,
) (*elcontracts.ChainReader, error) {
	testConfig := testutils.GetDefaultTestConfig()
	logger := logging.NewTextSLogger(os.Stdout, &logging.SLoggerOptions{Level: testConfig.LogLevel})
	ethHttpClient, err := ethclient.Dial(httpEndpoint)
	if err != nil {
		return nil, utils.WrapError("Failed to create eth client", err)
	}

	testReader, err := elcontracts.NewReaderFromConfig(
		config,
		ethHttpClient,
		logger,
	)
	if err != nil {
		return nil, utils.WrapError("Failed to create chain reader from config", err)
	}
	return testReader, nil
}
