package elcontracts_test

import (
	"context"
	"math/big"
	"os"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/testutils/testclients"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	allocationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"
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

func TestRegisterForOperatorSets(t *testing.T) {
	eigenClients, anvilHttpEndpoint := testclients.BuildTestClients(t)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	avsAddress := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	operatorSetId := uint32(1)
	erc20MockStrategyAddr := contractAddrs.Erc20MockStrategy

	err := createOperatorSet(eigenClients, avsAddress, operatorSetId, erc20MockStrategyAddr)
	require.NoError(t, err)

	operatorAddressHex := "70997970C51812dc3A010C7d01b50e0d17dc79C8"
	operatorPrivateKeyHex := "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"

	// Create operator clients
	operatorClients, err := newTestClients(anvilHttpEndpoint, operatorPrivateKeyHex)
	require.NoError(t, err)

	operatorAddress := common.HexToAddress(operatorAddressHex)
	keypair, err := bls.NewKeyPairFromString("0x01")
	require.NoError(t, err)

	request := elcontracts.RegistrationRequest{
		OperatorAddress: operatorAddress,
		AVSAddress:      avsAddress,
		OperatorSetIds:  []uint32{operatorSetId},
		WaitForReceipt:  true,
		Socket:          "socket",
		BlsKeyPair:      keypair,
	}

	operatorSet := allocationmanager.OperatorSet{
		Avs: avsAddress,
		Id:  uint32(operatorSetId),
	}

	registryCoordinatorAddress := contractAddrs.RegistryCoordinator
	receipt, err := operatorClients.ElChainWriter.RegisterForOperatorSets(
		context.Background(),
		registryCoordinatorAddress,
		request,
	)

	require.NoError(t, err)
	const RECEIPT_SUCCESS_STATUS = uint64(1)
	require.Equal(t, RECEIPT_SUCCESS_STATUS, receipt.Status)

	isRegistered, err := operatorClients.ElChainReader.IsOperatorRegisteredWithOperatorSet(
		context.Background(),
		operatorAddress,
		operatorSet,
	)
	require.NoError(t, err)
	require.Equal(t, isRegistered, true)
}

func createOperatorSet(client *clients.Clients, avsAddress common.Address, operatorSetId uint32, erc20MockStrategyAddr common.Address) error {
	allocationManagerAddress := client.EigenlayerContractBindings.AllocationManagerAddr
	allocationManager, err := allocationmanager.NewContractAllocationManager(allocationManagerAddress, client.EthHttpClient)
	if err != nil {
		return err
	}
	registryCoordinatorAddress := client.AvsRegistryContractBindings.RegistryCoordinatorAddr
	registryCoordinator, err := regcoord.NewContractRegistryCoordinator(registryCoordinatorAddress, client.EthHttpClient)
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
