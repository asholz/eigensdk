package avsregistry_test

import (
	"context"
	"math/big"
	"testing"

	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/testutils/testclients"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

func TestReaderMethods(t *testing.T) {
	clients, _ := testclients.BuildTestClients(t)
	chainReader := clients.ReadClients.AvsRegistryChainReader

	quorumNumbers := types.QuorumNums{0}

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
}

func TestIsOperatorSetQuorum(t *testing.T) {
	// Test set up
	clients, anvilHttpEndpoint := testclients.BuildTestClients(t)

	chainWriter := clients.AvsRegistryChainWriter
	chainReader := clients.ReadClients.AvsRegistryChainReader

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	// By default, quorums are not operator sets
	isOperatorSet, err := chainReader.IsOperatorSetQuorum(
		&bind.CallOpts{},
		0,
	)
	require.NoError(t, err)
	require.False(t, isOperatorSet)

	// Enabling operator sets to create slashable stake quorums
	registryCoordinatorAddress := contractAddrs.RegistryCoordinator
	registryCoordinator, err := regcoord.NewContractRegistryCoordinator(
		registryCoordinatorAddress,
		clients.EthHttpClient,
	)
	require.NoError(t, err)

	txManager := clients.TxManager
	noSendTxOpts, err := txManager.GetNoSendTxOpts()
	require.NoError(t, err)

	tx, err := registryCoordinator.EnableOperatorSets(noSendTxOpts)
	require.NoError(t, err)

	_, err = txManager.Send(context.Background(), tx, true)
	require.NoError(t, err)

	// Create a new slashable stake quorum
	operatorSetParams := regcoord.ISlashingRegistryCoordinatorTypesOperatorSetParam{
		MaxOperatorCount: 5,
	}
	minimumStakeNeeded := big.NewInt(0)

	strategyAddr := contractAddrs.Erc20MockStrategy
	strategyParam := regcoord.IStakeRegistryTypesStrategyParams{
		Strategy:   strategyAddr,
		Multiplier: big.NewInt(1e18),
	}

	lookAheadPeriod := uint32(0)

	receipt, err := chainWriter.CreateSlashableStakeQuorum(
		context.Background(),
		operatorSetParams,
		minimumStakeNeeded,
		[]regcoord.IStakeRegistryTypesStrategyParams{strategyParam},
		lookAheadPeriod,
		true,
	)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, gethtypes.ReceiptStatusSuccessful)

	// After inserting a new flow quorum, requesting if is an operator set returns true
	isOperatorSet, err = chainReader.IsOperatorSetQuorum(
		&bind.CallOpts{},
		1,
	)
	require.NoError(t, err)
	require.True(t, isOperatorSet)
}
