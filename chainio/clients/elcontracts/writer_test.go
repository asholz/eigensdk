package elcontracts_test

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/testutils/testclients"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	allocationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"
	contractreg "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ContractsRegistry"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
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

func getAllocationManagerAddress(ethHttpUrl string) common.Address {
	ethHttpClient, err := ethclient.Dial(ethHttpUrl)
	if err != nil {
		// TODO: handle error
		panic(err)
	}
	// The ContractsRegistry contract should always be deployed at this address on anvil
	// it's the address of the contract created at nonce 0 by the first anvil account
	// (0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266)
	contractsRegistry, err := contractreg.NewContractContractsRegistry(
		common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3"),
		ethHttpClient,
	)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		contractNumber := big.NewInt(int64(i))
		names, err := contractsRegistry.ContractNames(&bind.CallOpts{}, contractNumber)
		if err != nil {
			panic(err)
		}
		fmt.Println("ContractNames(0) = ", names)
	}
	allocationManagerAddr, err := contractsRegistry.Contracts(&bind.CallOpts{}, "allocationManager")
	if err != nil {
		panic(err)
	}
	return allocationManagerAddr
}

func TestRegisterForOperatorSets(t *testing.T) {
	const RECEIPT_SUCCESS_STATUS = uint64(1)
	// TODO: consider replacing all this setup with testclients.BuildTestClients
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)

	anvilWsEndpoint, err := anvilC.Endpoint(context.Background(), "ws")
	require.NoError(t, err)

	ethHttpClient, err := ethclient.Dial(anvilHttpEndpoint)
	require.NoError(t, err)

	// allocationManagerAddress := getAllocationManagerAddress(anvilHttpEndpoint)
	// fmt.Println("getAllocationManagerAddress = ", allocationManagerAddress)

	// TODO: fetch this address using the auxiliary function
	allocationManagerAddress := common.HexToAddress("0x8A791620dd6260079BF849Dc5567aDC3F2FdC318")
	fmt.Println("getAllocationManagerAddress = ", allocationManagerAddress)

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)
	registryCoordinatorAddress := contractAddrs.RegistryCoordinator
	allocationManager, err := allocationmanager.NewContractAllocationManager(allocationManagerAddress, ethHttpClient)
	require.NoError(t, err)

	registryCoordinator, err := regcoord.NewContractRegistryCoordinator(registryCoordinatorAddress, ethHttpClient)
	require.NoError(t, err)

	avsAddress := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

	chainioConfig := clients.BuildAllConfig{
		EthHttpUrl:                 anvilHttpEndpoint,
		EthWsUrl:                   anvilWsEndpoint,
		RegistryCoordinatorAddr:    contractAddrs.RegistryCoordinator.String(),
		OperatorStateRetrieverAddr: contractAddrs.OperatorStateRetriever.String(),
		AvsName:                    "exampleAvs",
		PromMetricsIpPortAddress:   ":9090",
	}
	fundedPrivateKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

	ecdsaPrivateKey, err := crypto.HexToECDSA(fundedPrivateKeyHex)
	require.NoError(t, err)
	logger := logging.NewTextSLogger(os.Stdout, &logging.SLoggerOptions{Level: testConfig.LogLevel})

	eigenClients, err := clients.BuildAll(
		chainioConfig,
		ecdsaPrivateKey,
		logger,
	)
	require.NoError(t, err)

	noSendTxOpts, err := eigenClients.TxManager.GetNoSendTxOpts()
	require.NoError(t, err)

	tx, err := allocationManager.SetAVSRegistrar(noSendTxOpts, avsAddress, registryCoordinatorAddress)
	require.NoError(t, err)

	waitForReceipt := true

	receipt, err := eigenClients.TxManager.Send(context.Background(), tx, waitForReceipt)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, RECEIPT_SUCCESS_STATUS)

	tx, err = registryCoordinator.EnableOperatorSets(noSendTxOpts)
	require.NoError(t, err)

	receipt, err = eigenClients.TxManager.Send(context.Background(), tx, waitForReceipt)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, RECEIPT_SUCCESS_STATUS)

	operatorSetParam := regcoord.IRegistryCoordinatorOperatorSetParam{
		MaxOperatorCount:        10,
		KickBIPsOfOperatorStake: 100,
		KickBIPsOfTotalStake:    1000,
	}
	minimumStake := big.NewInt(0)
	strategyParams := regcoord.IStakeRegistryStrategyParams{
		Strategy:   contractAddrs.Erc20MockStrategy,
		Multiplier: big.NewInt(1),
	}
	strategyParamsArray := []regcoord.IStakeRegistryStrategyParams{strategyParams}
	lookAheadPeriod := uint32(0)
	tx, err = registryCoordinator.CreateSlashableStakeQuorum(noSendTxOpts, operatorSetParam, minimumStake, strategyParamsArray, lookAheadPeriod)
	require.NoError(t, err)

	receipt, err = eigenClients.TxManager.Send(context.Background(), tx, waitForReceipt)
	require.Equal(t, receipt.Status, RECEIPT_SUCCESS_STATUS)
	require.NoError(t, err)

	operatorSetId := uint32(1)
	strategies := []common.Address{contractAddrs.Erc20MockStrategy}
	operatorSetParams := allocationmanager.IAllocationManagerTypesCreateSetParams{
		OperatorSetId: operatorSetId,
		Strategies:    strategies,
	}
	operatorSetParamsArray := []allocationmanager.IAllocationManagerTypesCreateSetParams{operatorSetParams}
	tx, err = allocationManager.CreateOperatorSets(noSendTxOpts, avsAddress, operatorSetParamsArray)
	require.NoError(t, err)

	receipt, err = eigenClients.TxManager.Send(context.Background(), tx, waitForReceipt)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, RECEIPT_SUCCESS_STATUS)

	operatorAddress := common.HexToAddress("70997970C51812dc3A010C7d01b50e0d17dc79C8")
	operatorPrivateKey, err := crypto.HexToECDSA("59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d")
	require.NoError(t, err)
	request := elcontracts.RegistrationRequest{
		OperatorAddress: operatorAddress,
		AVSAddress:      avsAddress,
		OperatorSetIds:  []uint32{operatorSetId},
		WaitForReceipt:  true,
	}
	operatorClients, err := clients.BuildAll(
		chainioConfig,
		operatorPrivateKey,
		logger,
	)
	require.NoError(t, err)

	receipt, err = operatorClients.ElChainWriter.RegisterForOperatorSets(context.Background(), request)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, RECEIPT_SUCCESS_STATUS)

	operatorSet := allocationmanager.OperatorSet{
		Avs: avsAddress,
		Id:  operatorSetId,
	}
	isRegistered, err := operatorClients.ElChainReader.IsOperatorRegisteredWithOperatorSet(context.Background(), operatorAddress, operatorSet)
	require.NoError(t, err)

	require.Equal(t, isRegistered, true)
}
