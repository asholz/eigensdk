package elcontracts_test

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
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

	contractreg "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ContractsRegistry"
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

	allocationManagerAddr, err := contractsRegistry.Contracts(&bind.CallOpts{}, "allocationManager")
	if err != nil {
		panic(err)
	}
	return allocationManagerAddr
}

func TestRegisterForOperatorSets(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)
	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)
	allocationManagerAddress := getAllocationManagerAddress(anvilHttpEndpoint)
	fmt.Println("getAllocationManagerAddress = ", allocationManagerAddress)
	// // anvilWsEndpoint, err := anvilC.Endpoint(context.Background(), "ws")
	// require.NoError(t, err)
	// logger := logging.NewTextSLogger(os.Stdout, &logging.SLoggerOptions{Level: testConfig.LogLevel})

	// code taken from Rust SDK test
	/*
			    async fn test_register_for_operator_sets() {
		        // let (_container, http_endpoint, _ws_endpoint) = start_anvil_container().await;
		        let http_endpoint = "http://localhost:8545".to_string();

		        let avs_address = OPERATOR_ADDRESS;

		        // set registrar - otherwise i dont know how to get permissions for registry coordinator address
		        // check if this is needed (should be done on deploy) or is already set
		        let allocation_manager_addr = get_allocation_manager_address(http_endpoint.clone()).await;
		        let signer = get_signer(OPERATOR_PRIVATE_KEY, &http_endpoint);
		        let allocation_manager = AllocationManager::new(allocation_manager_addr, signer.clone());
		        let registry_coordinator_addr =
		            get_registry_coordinator_address(http_endpoint.clone()).await;
		        allocation_manager
		            .setAVSRegistrar(avs_address, registry_coordinator_addr)
		            .send()
		            .await
		            .unwrap()
		            .get_receipt()
		            .await
		            .unwrap();

		        // enable operator sets in Registry Coordinator
		        let registry_coordinator =
		            RegistryCoordinator::new(registry_coordinator_addr, signer.clone());
		        registry_coordinator
		            .enableOperatorSets()
		            .send()
		            .await
		            .unwrap()
		            .get_receipt()
		            .await
		            .unwrap();

		        // create slashable quorum
		        let contract_registry_coordinator =
		            RegistryCoordinator::new(registry_coordinator_addr, signer.clone());
		        let operator_set_params = OperatorSetParam {
		            maxOperatorCount: 10,
		            kickBIPsOfOperatorStake: 100,
		            kickBIPsOfTotalStake: 1000,
		        };
		        let strategy_params = StrategyParams {
		            strategy: get_erc20_mock_strategy(http_endpoint.to_string()).await,
		            multiplier: U96::from(1),
		        };
		        contract_registry_coordinator
		            .createSlashableStakeQuorum(operator_set_params, U96::from(0), vec![strategy_params], 0)
		            .send()
		            .await
		            .unwrap()
		            .get_receipt()
		            .await
		            .unwrap();

		        // create operator set
		        let operator_set_id = 1;
		        let params = IAllocationManagerTypes::CreateSetParams {
		            operatorSetId: operator_set_id,
		            strategies: vec![get_erc20_mock_strategy(http_endpoint.clone()).await],
		        };
		        allocation_manager
		            .createOperatorSets(avs_address, vec![params])
		            .send()
		            .await
		            .unwrap()
		            .get_receipt()
		            .await
		            .unwrap();

		        // register to operator set
		        let operator_addr = address!("70997970C51812dc3A010C7d01b50e0d17dc79C8");
		        let operator_private_key =
		            "0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d";
		        let el_chain_writer =
		            new_test_writer(http_endpoint.clone(), operator_private_key.to_string()).await;
		        let tx_hash = el_chain_writer
		            .register_for_operator_sets(operator_addr, avs_address, vec![operator_set_id])
		            .await
		            .unwrap();

		        let receipt = wait_transaction(&http_endpoint, tx_hash).await.unwrap();
		        assert!(receipt.status());
		    }
	*/
}
