// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractOperatorStateRetriever

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

// OperatorStateRetrieverCheckSignaturesIndices is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverCheckSignaturesIndices struct {
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// OperatorStateRetrieverOperator is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverOperator struct {
	Operator   common.Address
	OperatorId [32]byte
	Stake      *big.Int
}

// ContractOperatorStateRetrieverMetaData contains all meta data concerning the ContractOperatorStateRetriever contract.
var ContractOperatorStateRetrieverMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getBatchOperatorFromId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50611dcb8061001f6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806331b36bd9146100675780633563b0d1146100905780634d2b57fe146100b05780634f739f74146100d05780635c155662146100f0578063cefdc1d414610110575b600080fd5b61007a6100753660046113b8565b610131565b60405161008791906114ac565b60405180910390f35b6100a361009e3660046114e8565b610245565b6040516100879190611653565b6100c36100be3660046116c4565b6106c5565b6040516100879190611713565b6100e36100de3660046117aa565b6107d2565b60405161008791906118ab565b6101036100fe366004611963565b610eea565b60405161008791906119c6565b61012361011e3660046119fe565b6110a8565b604051610087929190611a35565b606081516001600160401b0381111561014c5761014c61134f565b604051908082528060200260200182016040528015610175578160200160208202803683370190505b50905060005b825181101561023e57836001600160a01b03166313542a4e8483815181106101a5576101a5611a56565b60200260200101516040518263ffffffff1660e01b81526004016101d891906001600160a01b0391909116815260200190565b602060405180830381865afa1580156101f5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102199190611a6c565b82828151811061022b5761022b611a56565b602090810291909101015260010161017b565b5092915050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610287573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102ab9190611a85565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa1580156102ed573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103119190611a85565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610353573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103779190611a85565b9050600086516001600160401b038111156103945761039461134f565b6040519080825280602002602001820160405280156103c757816020015b60608152602001906001900390816103b25790505b50905060005b87518110156106b95760008882815181106103ea576103ea611a56565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa15801561044b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526104739190810190611aa2565b905080516001600160401b0381111561048e5761048e61134f565b6040519080825280602002602001820160405280156104d957816020015b60408051606081018252600080825260208083018290529282015282526000199092019101816104ac5790505b508484815181106104ec576104ec611a56565b602002602001018190525060005b81518110156106ae576040518060600160405280876001600160a01b03166347b314e885858151811061052f5761052f611a56565b60200260200101516040518263ffffffff1660e01b815260040161055591815260200190565b602060405180830381865afa158015610572573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105969190611a85565b6001600160a01b031681526020018383815181106105b6576105b6611a56565b60200260200101518152602001896001600160a01b031663fa28c6278585815181106105e4576105e4611a56565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015610640573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106649190611b37565b6001600160601b031681525085858151811061068257610682611a56565b6020026020010151828151811061069b5761069b611a56565b60209081029190910101526001016104fa565b5050506001016103cd565b50979650505050505050565b606081516001600160401b038111156106e0576106e061134f565b604051908082528060200260200182016040528015610709578160200160208202803683370190505b50905060005b825181101561023e57836001600160a01b031663296bb06484838151811061073957610739611a56565b60200260200101516040518263ffffffff1660e01b815260040161075f91815260200190565b602060405180830381865afa15801561077c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107a09190611a85565b8282815181106107b2576107b2611a56565b6001600160a01b039092166020928302919091019091015260010161070f565b6107fd6040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa15801561083d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108619190611a85565b905061088e6040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e906108be908b9089908990600401611b60565b600060405180830381865afa1580156108db573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526109039190810190611ba7565b81526040516340e03a8160e11b81526001600160a01b038316906381c0750290610935908b908b908b90600401611c64565b600060405180830381865afa158015610952573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261097a9190810190611ba7565b6040820152856001600160401b038111156109975761099761134f565b6040519080825280602002602001820160405280156109ca57816020015b60608152602001906001900390816109b55790505b50606082015260005b60ff8116871115610dfb576000856001600160401b038111156109f8576109f861134f565b604051908082528060200260200182016040528015610a21578160200160208202803683370190505b5083606001518360ff1681518110610a3b57610a3b611a56565b602002602001018190525060005b86811015610d055760008c6001600160a01b03166304ec63518a8a85818110610a7457610a74611a56565b905060200201358e88600001518681518110610a9257610a92611a56565b60200260200101516040518463ffffffff1660e01b8152600401610acf9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015610aec573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b109190611c8d565b9050806001600160c01b0316600003610bbb5760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a40160405180910390fd5b8a8a8560ff16818110610bd057610bd0611a56565b60016001600160c01b038516919093013560f81c1c82169091039050610cfc57856001600160a01b031663dd9846b98a8a85818110610c1157610c11611a56565b905060200201358d8d8860ff16818110610c2d57610c2d611a56565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015610c83573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ca79190611cb6565b85606001518560ff1681518110610cc057610cc0611a56565b60200260200101518481518110610cd957610cd9611a56565b63ffffffff9092166020928302919091019091015282610cf881611ce9565b9350505b50600101610a49565b506000816001600160401b03811115610d2057610d2061134f565b604051908082528060200260200182016040528015610d49578160200160208202803683370190505b50905060005b82811015610dc05784606001518460ff1681518110610d7057610d70611a56565b60200260200101518181518110610d8957610d89611a56565b6020026020010151828281518110610da357610da3611a56565b63ffffffff90921660209283029190910190910152600101610d4f565b508084606001518460ff1681518110610ddb57610ddb611a56565b602002602001018190525050508080610df390611d02565b9150506109d3565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e3c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e609190611a85565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90610e93908b908b908e90600401611d21565b600060405180830381865afa158015610eb0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610ed89190810190611ba7565b60208301525098975050505050505050565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b8152600401610f1c929190611d4b565b600060405180830381865afa158015610f39573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610f619190810190611ba7565b9050600084516001600160401b03811115610f7e57610f7e61134f565b604051908082528060200260200182016040528015610fa7578160200160208202803683370190505b50905060005b855181101561109e57866001600160a01b03166304ec6351878381518110610fd757610fd7611a56565b602002602001015187868581518110610ff257610ff2611a56565b60200260200101516040518463ffffffff1660e01b815260040161102f9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa15801561104c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110709190611c8d565b6001600160c01b031682828151811061108b5761108b611a56565b6020908102919091010152600101610fad565b5095945050505050565b60408051600180825281830190925260009160609183916020808301908036833701905050905084816000815181106110e3576110e3611a56565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e9061111f9088908690600401611d4b565b600060405180830381865afa15801561113c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526111649190810190611ba7565b60008151811061117657611176611a56565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa1580156111e2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112069190611c8d565b6001600160c01b03169050600061121c8261123a565b90508161122a8a838a610245565b9550955050505050935093915050565b606060008061124884611306565b61ffff166001600160401b038111156112635761126361134f565b6040519080825280601f01601f19166020018201604052801561128d576020820181803683370190505b5090506000805b8251821080156112a5575061010081105b156112fc576001811b9350858416156112ec578060f81b8383815181106112ce576112ce611a56565b60200101906001600160f81b031916908160001a9053508160010191505b6112f581611ce9565b9050611294565b5090949350505050565b6000805b82156113315761131b600184611d6a565b909216918061132981611d7d565b91505061130a565b92915050565b6001600160a01b038116811461134c57600080fd5b50565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b038111828210171561138d5761138d61134f565b604052919050565b60006001600160401b038211156113ae576113ae61134f565b5060051b60200190565b600080604083850312156113cb57600080fd5b82356113d681611337565b915060208301356001600160401b038111156113f157600080fd5b8301601f8101851361140257600080fd5b803561141561141082611395565b611365565b8082825260208201915060208360051b85010192508783111561143757600080fd5b6020840193505b8284101561146257833561145181611337565b82526020938401939091019061143e565b809450505050509250929050565b600081518084526020840193506020830160005b828110156114a2578151865260209586019590910190600101611484565b5093949350505050565b6020815260006114bf6020830184611470565b9392505050565b63ffffffff8116811461134c57600080fd5b80356114e3816114c6565b919050565b6000806000606084860312156114fd57600080fd5b833561150881611337565b925060208401356001600160401b0381111561152357600080fd5b8401601f8101861361153457600080fd5b80356001600160401b0381111561154d5761154d61134f565b611560601f8201601f1916602001611365565b81815287602083850101111561157557600080fd5b8160208401602083013760006020838301015280945050505061159a604085016114d8565b90509250925092565b600082825180855260208501945060208160051b8301016020850160005b8381101561164757848303601f190188528151805180855260209182019185019060005b8181101561162e57835180516001600160a01b03168452602080820151818601526040918201516001600160601b031691850191909152909301926060909201916001016115e5565b50506020998a01999094509290920191506001016115c1565b50909695505050505050565b6020815260006114bf60208301846115a3565b600082601f83011261167757600080fd5b813561168561141082611395565b8082825260208201915060208360051b8601019250858311156116a757600080fd5b602085015b8381101561109e5780358352602092830192016116ac565b600080604083850312156116d757600080fd5b82356116e281611337565b915060208301356001600160401b038111156116fd57600080fd5b61170985828601611666565b9150509250929050565b602080825282518282018190526000918401906040840190835b818110156117545783516001600160a01b031683526020938401939092019160010161172d565b509095945050505050565b60008083601f84011261177157600080fd5b5081356001600160401b0381111561178857600080fd5b6020830191508360208260051b85010111156117a357600080fd5b9250929050565b600080600080600080608087890312156117c357600080fd5b86356117ce81611337565b955060208701356117de816114c6565b945060408701356001600160401b038111156117f957600080fd5b8701601f8101891361180a57600080fd5b80356001600160401b0381111561182057600080fd5b89602082840101111561183257600080fd5b6020919091019450925060608701356001600160401b0381111561185557600080fd5b61186189828a0161175f565b979a9699509497509295939492505050565b600081518084526020840193506020830160005b828110156114a257815163ffffffff16865260209586019590910190600101611887565b6020815260008251608060208401526118c760a0840182611873565b90506020840151601f198483030160408501526118e48282611873565b9150506040840151601f198483030160608501526119028282611873565b6060860151858203601f190160808701528051808352919350602090810192508084019190600582901b85010160005b828110156106b957601f1986830301845261194e828651611873565b60209586019594909401939150600101611932565b60008060006060848603121561197857600080fd5b833561198381611337565b925060208401356001600160401b0381111561199e57600080fd5b6119aa86828701611666565b92505060408401356119bb816114c6565b809150509250925092565b602080825282518282018190526000918401906040840190835b818110156117545783518352602093840193909201916001016119e0565b600080600060608486031215611a1357600080fd5b8335611a1e81611337565b92506020840135915060408401356119bb816114c6565b828152604060208201526000611a4e60408301846115a3565b949350505050565b634e487b7160e01b600052603260045260246000fd5b600060208284031215611a7e57600080fd5b5051919050565b600060208284031215611a9757600080fd5b81516114bf81611337565b600060208284031215611ab457600080fd5b81516001600160401b03811115611aca57600080fd5b8201601f81018413611adb57600080fd5b8051611ae961141082611395565b8082825260208201915060208360051b850101925086831115611b0b57600080fd5b6020840193505b82841015611b2d578351825260209384019390910190611b12565b9695505050505050565b600060208284031215611b4957600080fd5b81516001600160601b03811681146114bf57600080fd5b63ffffffff84168152604060208201819052810182905260006001600160fb1b03831115611b8d57600080fd5b8260051b8085606085013791909101606001949350505050565b600060208284031215611bb957600080fd5b81516001600160401b03811115611bcf57600080fd5b8201601f81018413611be057600080fd5b8051611bee61141082611395565b8082825260208201915060208360051b850101925086831115611c1057600080fd5b6020840193505b82841015611b2d578351611c2a816114c6565b825260209384019390910190611c17565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff84168152604060208201526000611c84604083018486611c3b565b95945050505050565b600060208284031215611c9f57600080fd5b81516001600160c01b03811681146114bf57600080fd5b600060208284031215611cc857600080fd5b81516114bf816114c6565b634e487b7160e01b600052601160045260246000fd5b600060018201611cfb57611cfb611cd3565b5060010190565b600060ff821660ff8103611d1857611d18611cd3565b60010192915050565b604081526000611d35604083018587611c3b565b905063ffffffff83166020830152949350505050565b63ffffffff83168152604060208201526000611a4e6040830184611470565b8181038181111561133157611331611cd3565b600061ffff821661ffff8103611d1857611d18611cd356fea2646970667358221220f0eae877712b1cfe17de1bcfe3a9adc3ec48726e24eeef8cf19d62ab080f932564736f6c634300081b0033",
}

// ContractOperatorStateRetrieverABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractOperatorStateRetrieverMetaData.ABI instead.
var ContractOperatorStateRetrieverABI = ContractOperatorStateRetrieverMetaData.ABI

// ContractOperatorStateRetrieverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractOperatorStateRetrieverMetaData.Bin instead.
var ContractOperatorStateRetrieverBin = ContractOperatorStateRetrieverMetaData.Bin

// DeployContractOperatorStateRetriever deploys a new Ethereum contract, binding an instance of ContractOperatorStateRetriever to it.
func DeployContractOperatorStateRetriever(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractOperatorStateRetriever, error) {
	parsed, err := ContractOperatorStateRetrieverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractOperatorStateRetrieverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractOperatorStateRetriever{ContractOperatorStateRetrieverCaller: ContractOperatorStateRetrieverCaller{contract: contract}, ContractOperatorStateRetrieverTransactor: ContractOperatorStateRetrieverTransactor{contract: contract}, ContractOperatorStateRetrieverFilterer: ContractOperatorStateRetrieverFilterer{contract: contract}}, nil
}

// ContractOperatorStateRetrieverMethods is an auto generated interface around an Ethereum contract.
type ContractOperatorStateRetrieverMethods interface {
	ContractOperatorStateRetrieverCalls
	ContractOperatorStateRetrieverTransacts
	ContractOperatorStateRetrieverFilters
}

// ContractOperatorStateRetrieverCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractOperatorStateRetrieverCalls interface {
	GetBatchOperatorFromId(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error)

	GetBatchOperatorId(opts *bind.CallOpts, registryCoordinator common.Address, operators []common.Address) ([][32]byte, error)

	GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error)

	GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error)

	GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error)

	GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error)
}

// ContractOperatorStateRetrieverTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractOperatorStateRetrieverTransacts interface {
}

// ContractOperatorStateRetrieverFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractOperatorStateRetrieverFilters interface {
}

// ContractOperatorStateRetriever is an auto generated Go binding around an Ethereum contract.
type ContractOperatorStateRetriever struct {
	ContractOperatorStateRetrieverCaller     // Read-only binding to the contract
	ContractOperatorStateRetrieverTransactor // Write-only binding to the contract
	ContractOperatorStateRetrieverFilterer   // Log filterer for contract events
}

// ContractOperatorStateRetriever implements the ContractOperatorStateRetrieverMethods interface.
var _ ContractOperatorStateRetrieverMethods = (*ContractOperatorStateRetriever)(nil)

// ContractOperatorStateRetrieverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractOperatorStateRetrieverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOperatorStateRetrieverCaller implements the ContractOperatorStateRetrieverCalls interface.
var _ ContractOperatorStateRetrieverCalls = (*ContractOperatorStateRetrieverCaller)(nil)

// ContractOperatorStateRetrieverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractOperatorStateRetrieverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOperatorStateRetrieverTransactor implements the ContractOperatorStateRetrieverTransacts interface.
var _ ContractOperatorStateRetrieverTransacts = (*ContractOperatorStateRetrieverTransactor)(nil)

// ContractOperatorStateRetrieverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractOperatorStateRetrieverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOperatorStateRetrieverFilterer implements the ContractOperatorStateRetrieverFilters interface.
var _ ContractOperatorStateRetrieverFilters = (*ContractOperatorStateRetrieverFilterer)(nil)

// ContractOperatorStateRetrieverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractOperatorStateRetrieverSession struct {
	Contract     *ContractOperatorStateRetriever // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                   // Call options to use throughout this session
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ContractOperatorStateRetrieverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractOperatorStateRetrieverCallerSession struct {
	Contract *ContractOperatorStateRetrieverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                         // Call options to use throughout this session
}

// ContractOperatorStateRetrieverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractOperatorStateRetrieverTransactorSession struct {
	Contract     *ContractOperatorStateRetrieverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// ContractOperatorStateRetrieverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractOperatorStateRetrieverRaw struct {
	Contract *ContractOperatorStateRetriever // Generic contract binding to access the raw methods on
}

// ContractOperatorStateRetrieverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractOperatorStateRetrieverCallerRaw struct {
	Contract *ContractOperatorStateRetrieverCaller // Generic read-only contract binding to access the raw methods on
}

// ContractOperatorStateRetrieverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractOperatorStateRetrieverTransactorRaw struct {
	Contract *ContractOperatorStateRetrieverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractOperatorStateRetriever creates a new instance of ContractOperatorStateRetriever, bound to a specific deployed contract.
func NewContractOperatorStateRetriever(address common.Address, backend bind.ContractBackend) (*ContractOperatorStateRetriever, error) {
	contract, err := bindContractOperatorStateRetriever(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractOperatorStateRetriever{ContractOperatorStateRetrieverCaller: ContractOperatorStateRetrieverCaller{contract: contract}, ContractOperatorStateRetrieverTransactor: ContractOperatorStateRetrieverTransactor{contract: contract}, ContractOperatorStateRetrieverFilterer: ContractOperatorStateRetrieverFilterer{contract: contract}}, nil
}

// NewContractOperatorStateRetrieverCaller creates a new read-only instance of ContractOperatorStateRetriever, bound to a specific deployed contract.
func NewContractOperatorStateRetrieverCaller(address common.Address, caller bind.ContractCaller) (*ContractOperatorStateRetrieverCaller, error) {
	contract, err := bindContractOperatorStateRetriever(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractOperatorStateRetrieverCaller{contract: contract}, nil
}

// NewContractOperatorStateRetrieverTransactor creates a new write-only instance of ContractOperatorStateRetriever, bound to a specific deployed contract.
func NewContractOperatorStateRetrieverTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractOperatorStateRetrieverTransactor, error) {
	contract, err := bindContractOperatorStateRetriever(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractOperatorStateRetrieverTransactor{contract: contract}, nil
}

// NewContractOperatorStateRetrieverFilterer creates a new log filterer instance of ContractOperatorStateRetriever, bound to a specific deployed contract.
func NewContractOperatorStateRetrieverFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractOperatorStateRetrieverFilterer, error) {
	contract, err := bindContractOperatorStateRetriever(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractOperatorStateRetrieverFilterer{contract: contract}, nil
}

// bindContractOperatorStateRetriever binds a generic wrapper to an already deployed contract.
func bindContractOperatorStateRetriever(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractOperatorStateRetrieverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractOperatorStateRetriever.Contract.ContractOperatorStateRetrieverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOperatorStateRetriever.Contract.ContractOperatorStateRetrieverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractOperatorStateRetriever.Contract.ContractOperatorStateRetrieverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractOperatorStateRetriever.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOperatorStateRetriever.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractOperatorStateRetriever.Contract.contract.Transact(opts, method, params...)
}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCaller) GetBatchOperatorFromId(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _ContractOperatorStateRetriever.contract.Call(opts, &out, "getBatchOperatorFromId", registryCoordinator, operatorIds)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverSession) GetBatchOperatorFromId(registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	return _ContractOperatorStateRetriever.Contract.GetBatchOperatorFromId(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operatorIds)
}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerSession) GetBatchOperatorFromId(registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	return _ContractOperatorStateRetriever.Contract.GetBatchOperatorFromId(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operatorIds)
}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCaller) GetBatchOperatorId(opts *bind.CallOpts, registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _ContractOperatorStateRetriever.contract.Call(opts, &out, "getBatchOperatorId", registryCoordinator, operators)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverSession) GetBatchOperatorId(registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	return _ContractOperatorStateRetriever.Contract.GetBatchOperatorId(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operators)
}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerSession) GetBatchOperatorId(registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	return _ContractOperatorStateRetriever.Contract.GetBatchOperatorId(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operators)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractOperatorStateRetriever.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractOperatorStateRetriever.Contract.GetCheckSignaturesIndices(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractOperatorStateRetriever.Contract.GetCheckSignaturesIndices(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractOperatorStateRetriever.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractOperatorStateRetriever.Contract.GetOperatorState(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractOperatorStateRetriever.Contract.GetOperatorState(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractOperatorStateRetriever.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

	if err != nil {
		return *new(*big.Int), *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, out1, err

}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractOperatorStateRetriever.Contract.GetOperatorState0(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractOperatorStateRetriever.Contract.GetOperatorState0(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCaller) GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractOperatorStateRetriever.contract.Call(opts, &out, "getQuorumBitmapsAtBlockNumber", registryCoordinator, operatorIds, blockNumber)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractOperatorStateRetriever.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractOperatorStateRetriever.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operatorIds, blockNumber)
}
