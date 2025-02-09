package avsregistry_test

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/testutils/testclients"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestReaderMethods(t *testing.T) {
	clients, anvilHttpEndpoint := testclients.BuildTestClients(t)
	chainReader := clients.ReadClients.AvsRegistryChainReader
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)
	strategy := contractAddrs.Erc20MockStrategy
	quorumNumber := types.QuorumNum(0)
	quorumNumbers := types.QuorumNums{0}

	operatorPrivateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY

	config := avsregistry.Config{
		RegistryCoordinatorAddress:    contractAddrs.RegistryCoordinator,
		OperatorStateRetrieverAddress: contractAddrs.OperatorStateRetriever,
		ServiceManagerAddress:         contractAddrs.ServiceManager,
	}

	chainWriter, err := testclients.NewTestAvsRegistryWriterFromConfig(anvilHttpEndpoint, operatorPrivateKeyHex, config)
	require.NoError(t, err)

	keypair, err := bls.NewKeyPairFromString("0x01")
	require.NoError(t, err)

	operatorAddress := gethcommon.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
	operatorPrivateKey, err := crypto.HexToECDSA(testutils.ANVIL_FIRST_PRIVATE_KEY)
	require.NoError(t, err)

	t.Run("get quorum state", func(t *testing.T) {
		count, err := chainReader.GetQuorumCount(&bind.CallOpts{})
		require.NoError(t, err)
		require.NotNil(t, count)
	})

	t.Run("get operator stake in quorums at current block", func(t *testing.T) {
		stake, err := chainReader.GetOperatorsStakeInQuorumsAtCurrentBlock(&bind.CallOpts{}, quorumNumbers)
		require.NoError(t, err)
		require.NotNil(t, stake)
	})

	t.Run("get operator stake in quorums at block", func(t *testing.T) {
		stake, err := chainReader.GetOperatorsStakeInQuorumsAtBlock(&bind.CallOpts{}, quorumNumbers, 100)
		require.NoError(t, err)
		require.NotNil(t, stake)
	})

	t.Run("get operator address in quorums at current block", func(t *testing.T) {
		addresses, err := chainReader.GetOperatorAddrsInQuorumsAtCurrentBlock(&bind.CallOpts{}, quorumNumbers)
		require.NoError(t, err)
		require.NotNil(t, addresses)
	})

	t.Run(
		"get operators stake in quorums of operator at block returns error for non-registered operator",
		func(t *testing.T) {
			operatorAddress := common.Address{0x1}
			operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
			require.NoError(t, err)

			_, _, err = chainReader.GetOperatorsStakeInQuorumsOfOperatorAtBlock(&bind.CallOpts{}, operatorId, 100)
			require.Error(t, err)
			require.Contains(t, err.Error(), "Failed to get operators state")
		})

	t.Run(
		"get single operator stake in quorums of operator at current block returns error for non-registered operator",
		func(t *testing.T) {
			operatorAddress := common.Address{0x1}
			operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
			require.NoError(t, err)

			stakes, err := chainReader.GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock(&bind.CallOpts{}, operatorId)
			require.NoError(t, err)
			require.Equal(t, 0, len(stakes))
		})

	t.Run("get check signatures indices returns error for non-registered operator", func(t *testing.T) {
		operatorAddress := common.Address{0x1}
		operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)

		_, err = chainReader.GetCheckSignaturesIndices(
			&bind.CallOpts{},
			100,
			quorumNumbers,
			[]types.OperatorId{operatorId},
		)
		require.Contains(t, err.Error(), "Failed to get check signatures indices")
	})

	t.Run("get operator id", func(t *testing.T) {
		operatorAddress := common.Address{0x1}
		operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)
		require.NotNil(t, operatorId)
	})

	t.Run("get operator from id returns zero address for non-registered operator", func(t *testing.T) {
		operatorAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")
		operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)

		retrievedAddress, err := chainReader.GetOperatorFromId(&bind.CallOpts{}, operatorId)
		require.NoError(t, err)
		require.Equal(t, retrievedAddress, common.Address{0x0})
	})

	t.Run("query registration detail", func(t *testing.T) {
		operatorAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")
		quorums, err := chainReader.QueryRegistrationDetail(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)
		require.Equal(t, 1, len(quorums))
	})

	t.Run("is operator registered", func(t *testing.T) {
		operatorAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")
		isRegistered, err := chainReader.IsOperatorRegistered(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)
		require.False(t, isRegistered)
	})

	t.Run(
		"query existing registered operator pub keys", func(t *testing.T) {
			addresses, pubKeys, err := chainReader.QueryExistingRegisteredOperatorPubKeys(
				context.Background(),
				big.NewInt(0),
				nil,
				nil,
			)
			require.NoError(t, err)
			require.Equal(t, 0, len(pubKeys))
			require.Equal(t, 0, len(addresses))
		})

	t.Run(
		"query existing registered operator sockets", func(t *testing.T) {
			address_to_sockets, err := chainReader.QueryExistingRegisteredOperatorSockets(
				context.Background(),
				big.NewInt(0),
				nil,
				nil,
			)
			require.NoError(t, err)
			require.Equal(t, 0, len(address_to_sockets))
		})

	t.Run("Is operator set quorum", func(t *testing.T) {
		// A quorum registered in the old workflow should return false
		isOperatorSet, err := chainReader.IsOperatorSetQuorum(
			&bind.CallOpts{},
			0,
		)
		require.NoError(t, err)
		require.False(t, isOperatorSet)

		// TODO: Make a test with the new workflow testing a case returning true
	})

	t.Run("Get weight of operator for quorum", func(t *testing.T) {
		operatorAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")
		// A quorum registered in the old workflow should return 0
		weight, err := chainReader.WeightOfOperatorForQuorum(
			&bind.CallOpts{},
			0,
			operatorAddress,
		)
		require.NoError(t, err)
		require.Equal(t, int64(0), weight.Int64())
	})

	t.Run("Get strategy params length", func(t *testing.T) {
		quorumNumber := types.QuorumNum(0)
		length, err := chainReader.StrategyParamsLength(&bind.CallOpts{}, quorumNumber)
		require.NoError(t, err)
		require.Equal(t, int64(1), length.Int64())
	})

	t.Run("Get strategy params by index", func(t *testing.T) {
		params, err := chainReader.StrategyParamsByIndex(&bind.CallOpts{}, quorumNumber, big.NewInt(0))
		require.NoError(t, err)
		require.Equal(t, strategy, params.Strategy)
	})

	t.Run("Get stakeHistory length", func(t *testing.T) {
		operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)
		length, err := chainReader.GetStakeHistoryLength(&bind.CallOpts{}, operatorId, quorumNumber)
		require.NoError(t, err)
		require.Equal(t, int64(1), length.Int64())
	})

	t.Run("Get Stake History", func(t *testing.T) {
		operatorAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")
		operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)
		stakeHistory, err := chainReader.GetStakeHistory(&bind.CallOpts{}, operatorId, quorumNumber)
		require.NoError(t, err)
		require.Equal(t, 0, len(stakeHistory))
	})

	t.Run("Get latest stake update", func(t *testing.T) {

		// stakeUpdate, err := chainReader.GetLatestStakeUpdate(&bind.CallOpts{}, operatorId, quorumNumber)

		//REGISTER OPERATOR
		receipt, err := chainWriter.RegisterOperator(
			context.Background(),
			operatorPrivateKey,
			keypair,
			quorumNumbers,
			"",
			true,
		)
		require.NoError(t, err)
		require.NotNil(t, receipt)

		operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)

		stakeUpdate2, err := chainReader.GetLatestStakeUpdate(&bind.CallOpts{}, operatorId, quorumNumber)
		fmt.Println("STAKE", stakeUpdate2.Stake)
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), stakeUpdate2.Stake)
		require.Equal(t, uint32(0), stakeUpdate2.NextUpdateBlockNumber)
	})

	// t.Run("Get stake at block number", func(t *testing.T) {
	// 	operatorAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")
	// 	operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
	// 	require.NoError(t, err)
	// 	stake, err := chainReader.GetStakeAtBlockNumber(&bind.CallOpts{}, operatorId, quorumNumber, 0)
	// 	require.NoError(t, err)
	// 	require.Equal(t, int64(0), stake.Int64())
	// })

}
