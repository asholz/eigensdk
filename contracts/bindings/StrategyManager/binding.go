// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractStrategyManager

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

// ContractStrategyManagerMetaData contains all meta data concerning the ContractStrategyManager contract.
var ContractStrategyManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_delegation\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEPOSIT_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addStrategiesToDepositWhitelist\",\"inputs\":[{\"name\":\"strategiesToWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateStrategyDepositDigestHash\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositIntoStrategy\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"depositedShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositIntoStrategyWithSignature\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"depositedShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDeposits\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakerStrategyList\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialStrategyWhitelister\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeDepositShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"depositSharesToRemove\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeStrategiesFromDepositWhitelist\",\"inputs\":[{\"name\":\"strategiesToRemoveFromWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStrategyWhitelister\",\"inputs\":[{\"name\":\"newStrategyWhitelister\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakerDepositShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerStrategyList\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"strategies\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerStrategyListLength\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyIsWhitelistedForDeposit\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"whitelisted\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyWhitelister\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawSharesAsTokens\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyAddedToDepositWhitelist\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyRemovedFromDepositWhitelist\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyWhitelisterChanged\",\"inputs\":[{\"name\":\"previousAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxStrategiesExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyDelegationManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyStrategyWhitelister\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SharesAmountTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SharesAmountZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StakerAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotWhitelisted\",\"inputs\":[]}]",
	Bin: "0x61010060405234801561001157600080fd5b506040516126fa3803806126fa8339810160408190526100309161020a565b81816001600160a01b038116610059576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b039081166080521660a0524660c052610077610089565b60e052610082610133565b5050610244565b600060c051461461012c5750604080518082018252600a81526922b4b3b2b72630bcb2b960b11b60209182015281517f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a866818301527f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea81840152466060820152306080808301919091528351808303909101815260a0909101909252815191012090565b5060e05190565b600054610100900460ff161561019f5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116146101f0576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b038116811461020757600080fd5b50565b6000806040838503121561021d57600080fd5b8251610228816101f2565b6020840151909250610239816101f2565b809150509250929050565b60805160a05160c05160e05161244c6102ae60003960006117e401526000611724015260008181610461015281816106fd01528181610a6501528181610e1301526113cf01526000818161032001528181610504015281816108670152611007015261244c6000f3fe608060405234801561001057600080fd5b50600436106101da5760003560e01c80638da5cb5b11610104578063cbc2bd62116100a2578063f2fde38b11610071578063f2fde38b14610496578063f698da25146104a9578063fabc1cbc146104b1578063fe243a17146104c457600080fd5b8063cbc2bd6214610429578063de44acb61461043c578063df5cf7231461045c578063e7a050aa1461048357600080fd5b80639ac01d61116100de5780639ac01d61146103c8578063b5d8b5b8146103db578063c4623ea1146103ee578063c66567021461041657600080fd5b80638da5cb5b1461038357806394f649dd14610394578063967fc0d2146103b557600080fd5b80635c975abb1161017c578063724af4231161014b578063724af423146102e85780637ecebe00146102fb578063886f11951461031b5780638b8aac3c1461035a57600080fd5b80635c975abb146102a25780635de08ff2146102aa578063663c1de4146102bd578063715018a6146102e057600080fd5b806332e89ace116101b857806332e89ace1461021a57806348825e9414610240578063595c6a67146102675780635ac86ab71461026f57600080fd5b8063136439dd146101df5780631794bb3c146101f45780632eae418c14610207575b600080fd5b6101f26101ed366004611e96565b6104ef565b005b6101f2610202366004611ec4565b6105c6565b6101f2610215366004611f05565b6106f2565b61022d610228366004611f6c565b6107ab565b6040519081526020015b60405180910390f35b61022d7f4337f82d142e41f2a8c10547cd8c859bddb92262a61058e77842e24d9dea922481565b6101f2610852565b61029261027d36600461206f565b609854600160ff9092169190911b9081161490565b6040519015158152602001610237565b60985461022d565b6101f26102b8366004612092565b610904565b6102926102cb366004612109565b60d16020526000908152604090205460ff1681565b6101f2610a48565b6101f26102f6366004611ec4565b610a5a565b61022d610309366004612109565b60ca6020526000908152604090205481565b6103427f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610237565b61022d610368366004612109565b6001600160a01b0316600090815260ce602052604090205490565b6033546001600160a01b0316610342565b6103a76103a2366004612109565b610aae565b60405161023792919061216b565b60cb54610342906001600160a01b031681565b61022d6103d63660046121c5565b610c2e565b6101f26103e9366004612092565b610cc0565b6104016103fc366004611f05565b610e05565b60408051928352602083019190915201610237565b6101f2610424366004612109565b610e6b565b61034261043736600461222a565b610e7f565b61044f61044a366004612109565b610eb7565b6040516102379190612256565b6103427f000000000000000000000000000000000000000000000000000000000000000081565b61022d610491366004611ec4565b610f2d565b6101f26104a4366004612109565b610f80565b61022d610ff6565b6101f26104bf366004611e96565b611005565b61022d6104d2366004612269565b60cd60209081526000928352604080842090915290825290205481565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610553573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061057791906122a2565b61059457604051631d77d47760e21b815260040160405180910390fd5b60985481811681146105b95760405163c61dca5d60e01b815260040160405180910390fd5b6105c28261111d565b5050565b600054610100900460ff16158080156105e65750600054600160ff909116105b806106005750303b158015610600575060005460ff166001145b6106685760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff19166001179055801561068b576000805461ff0019166101001790555b6106948261111d565b61069d8461115a565b6106a6836111ac565b80156106ec576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461073b5760405163f739589b60e01b815260040160405180910390fd5b604051636ce5768960e11b81526001600160a01b03858116600483015283811660248301526044820183905284169063d9caed1290606401600060405180830381600087803b15801561078d57600080fd5b505af11580156107a1573d6000803e3d6000fd5b5050505050505050565b60985460009081906001908116036107d65760405163840a48d560e01b815260040160405180910390fd5b6107de611215565b6001600160a01b038516600090815260ca602052604090205461081086610809818c8c8c878c610c2e565b868861126e565b6001600160a01b038616600090815260ca6020526040902060018201905561083a868a8a8a6112c0565b9250506108476001606555565b509695505050505050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa1580156108b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108da91906122a2565b6108f757604051631d77d47760e21b815260040160405180910390fd5b61090260001961111d565b565b60cb546001600160a01b0316331461092f576040516320ba3ff960e21b815260040160405180910390fd5b8060005b818110156106ec5760d16000858584818110610951576109516122c4565b90506020020160208101906109669190612109565b6001600160a01b0316815260208101919091526040016000205460ff16610a4057600160d1600086868581811061099f5761099f6122c4565b90506020020160208101906109b49190612109565b6001600160a01b031681526020810191909152604001600020805460ff19169115159190911790557f0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe848483818110610a0f57610a0f6122c4565b9050602002016020810190610a249190612109565b6040516001600160a01b03909116815260200160405180910390a15b600101610933565b610a50611438565b610902600061115a565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610aa35760405163f739589b60e01b815260040160405180910390fd5b6106ec838383611492565b6001600160a01b038116600090815260ce60205260408120546060918291908167ffffffffffffffff811115610ae657610ae6611f56565b604051908082528060200260200182016040528015610b0f578160200160208202803683370190505b50905060005b82811015610ba0576001600160a01b038616600090815260cd6020908152604080832060ce9092528220805491929184908110610b5457610b546122c4565b60009182526020808320909101546001600160a01b031683528201929092526040019020548251839083908110610b8d57610b8d6122c4565b6020908102919091010152600101610b15565b5060ce6000866001600160a01b03166001600160a01b031681526020019081526020016000208181805480602002602001604051908101604052809291908181526020018280548015610c1c57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610bfe575b50505050509150935093505050915091565b604080517f4337f82d142e41f2a8c10547cd8c859bddb92262a61058e77842e24d9dea922460208201526001600160a01b03808916928201929092528187166060820152908516608082015260a0810184905260c0810183905260e08101829052600090610cb590610100016040516020818303038152906040528051906020012061155b565b979650505050505050565b60cb546001600160a01b03163314610ceb576040516320ba3ff960e21b815260040160405180910390fd5b8060005b818110156106ec5760d16000858584818110610d0d57610d0d6122c4565b9050602002016020810190610d229190612109565b6001600160a01b0316815260208101919091526040016000205460ff1615610dfd57600060d16000868685818110610d5c57610d5c6122c4565b9050602002016020810190610d719190612109565b6001600160a01b031681526020810191909152604001600020805460ff19169115159190911790557f4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030848483818110610dcc57610dcc6122c4565b9050602002016020810190610de19190612109565b6040516001600160a01b03909116815260200160405180910390a15b600101610cef565b600080336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610e515760405163f739589b60e01b815260040160405180910390fd5b610e5d868587866115a2565b915091505b94509492505050565b610e73611438565b610e7c816111ac565b50565b60ce6020528160005260406000208181548110610e9b57600080fd5b6000918252602090912001546001600160a01b03169150829050565b6001600160a01b038116600090815260ce6020908152604091829020805483518184028101840190945280845260609392830182828015610f2157602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610f03575b50505050509050919050565b6098546000908190600190811603610f585760405163840a48d560e01b815260040160405180910390fd5b610f60611215565b610f6c338686866112c0565b9150610f786001606555565b509392505050565b610f88611438565b6001600160a01b038116610fed5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161065f565b610e7c8161115a565b6000611000611720565b905090565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611063573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061108791906122da565b6001600160a01b0316336001600160a01b0316146110b85760405163794821ff60e01b815260040160405180910390fd5b609854801982198116146110df5760405163c61dca5d60e01b815260040160405180910390fd5b609882905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200160405180910390a25050565b609881905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60cb54604080516001600160a01b03928316815291831660208301527f4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29910160405180910390a160cb80546001600160a01b0319166001600160a01b0392909216919091179055565b6002606554036112675760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161065f565b6002606555565b4281101561128f57604051630819bdcd60e01b815260040160405180910390fd5b6112a36001600160a01b0385168484611806565b6106ec57604051638baa579f60e01b815260040160405180910390fd5b6001600160a01b038316600090815260d16020526040812054849060ff166112fb57604051632efd965160e11b815260040160405180910390fd5b6113106001600160a01b038516338786611867565b6040516311f9fbc960e21b81526001600160a01b038581166004830152602482018590528616906347e7ef24906044016020604051808303816000875af115801561135f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061138391906122f7565b9150600080611394888789876115a2565b604051631e328e7960e11b81526001600160a01b038b811660048301528a8116602483015260448201849052606482018390529294509092507f000000000000000000000000000000000000000000000000000000000000000090911690633c651cf290608401600060405180830381600087803b15801561141557600080fd5b505af1158015611429573d6000803e3d6000fd5b50505050505050949350505050565b6033546001600160a01b031633146109025760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161065f565b6000816000036114b5576040516342061b2560e11b815260040160405180910390fd5b6001600160a01b03808516600090815260cd6020908152604080832093871683529290522054808311156114fc57604051634b18b19360e01b815260040160405180910390fd5b6115068382612326565b6001600160a01b03808716600090815260cd60209081526040808320938916835292905290812082905590915081900361154e5761154485856118c1565b6001915050611554565b60009150505b9392505050565b6000611565611720565b60405161190160f01b6020820152602281019190915260428101839052606201604051602081830303815290604052805190602001209050919050565b6000806001600160a01b0386166115cc576040516316f2ccc960e01b815260040160405180910390fd5b826000036115ed576040516342061b2560e11b815260040160405180910390fd5b6001600160a01b03808716600090815260cd6020908152604080832093881683529290529081205490819003611696576001600160a01b038716600090815260ce602090815260409091205410611657576040516301a1443960e31b815260040160405180910390fd5b6001600160a01b03878116600090815260ce602090815260408220805460018101825590835291200180546001600160a01b0319169187169190911790555b6116a0848261233f565b6001600160a01b03888116600081815260cd602090815260408083208b861680855290835292819020959095558451928352928a169282019290925291820152606081018590527f7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a969060800160405180910390a196929550919350505050565b60007f000000000000000000000000000000000000000000000000000000000000000046146117e15750604080518082018252600a81526922b4b3b2b72630bcb2b960b11b60209182015281517f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a866818301527f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea81840152466060820152306080808301919091528351808303909101815260a0909101909252815191012090565b507f000000000000000000000000000000000000000000000000000000000000000090565b60008060006118158585611a4a565b9092509050600081600481111561182e5761182e612352565b14801561184c5750856001600160a01b0316826001600160a01b0316145b8061185d575061185d868686611a8f565b9695505050505050565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b1790526106ec908590611b7b565b6001600160a01b038216600090815260ce6020526040812054905b818110156119db576001600160a01b03848116600090815260ce6020526040902080549185169183908110611913576119136122c4565b6000918252602090912001546001600160a01b0316036119d3576001600160a01b038416600090815260ce60205260409020805461195390600190612326565b81548110611963576119636122c4565b60009182526020808320909101546001600160a01b03878116845260ce90925260409092208054919092169190839081106119a0576119a06122c4565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506119db565b6001016118dc565b8181036119fb57604051632df15a4160e11b815260040160405180910390fd5b6001600160a01b038416600090815260ce60205260409020805480611a2257611a22612368565b600082815260209020810160001990810180546001600160a01b031916905501905550505050565b6000808251604103611a805760208301516040840151606085015160001a611a7487828585611c55565b94509450505050611a88565b506000905060025b9250929050565b6000806000856001600160a01b0316631626ba7e60e01b8686604051602401611ab99291906123ce565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092529051611af791906123e7565b600060405180830381855afa9150503d8060008114611b32576040519150601f19603f3d011682016040523d82523d6000602084013e611b37565b606091505b5091509150818015611b4b57506020815110155b801561185d57508051630b135d3f60e11b90611b7090830160209081019084016122f7565b149695505050505050565b6000611bd0826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316611d169092919063ffffffff16565b9050805160001480611bf1575080806020019051810190611bf191906122a2565b611c505760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840161065f565b505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115611c8c5750600090506003610e62565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611ce0573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116611d0957600060019250925050610e62565b9660009650945050505050565b6060611d258484600085611d2d565b949350505050565b606082471015611d8e5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840161065f565b600080866001600160a01b03168587604051611daa91906123e7565b60006040518083038185875af1925050503d8060008114611de7576040519150601f19603f3d011682016040523d82523d6000602084013e611dec565b606091505b5091509150610cb58783838760608315611e67578251600003611e60576001600160a01b0385163b611e605760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161065f565b5081611d25565b611d258383815115611e7c5781518083602001fd5b8060405162461bcd60e51b815260040161065f9190612403565b600060208284031215611ea857600080fd5b5035919050565b6001600160a01b0381168114610e7c57600080fd5b600080600060608486031215611ed957600080fd5b8335611ee481611eaf565b92506020840135611ef481611eaf565b929592945050506040919091013590565b60008060008060808587031215611f1b57600080fd5b8435611f2681611eaf565b93506020850135611f3681611eaf565b92506040850135611f4681611eaf565b9396929550929360600135925050565b634e487b7160e01b600052604160045260246000fd5b60008060008060008060c08789031215611f8557600080fd5b8635611f9081611eaf565b95506020870135611fa081611eaf565b9450604087013593506060870135611fb781611eaf565b92506080870135915060a087013567ffffffffffffffff811115611fda57600080fd5b8701601f81018913611feb57600080fd5b803567ffffffffffffffff81111561200557612005611f56565b604051601f8201601f19908116603f0116810167ffffffffffffffff8111828210171561203457612034611f56565b6040528181528282016020018b101561204c57600080fd5b816020840160208301376000602083830101528093505050509295509295509295565b60006020828403121561208157600080fd5b813560ff8116811461155457600080fd5b600080602083850312156120a557600080fd5b823567ffffffffffffffff8111156120bc57600080fd5b8301601f810185136120cd57600080fd5b803567ffffffffffffffff8111156120e457600080fd5b8560208260051b84010111156120f957600080fd5b6020919091019590945092505050565b60006020828403121561211b57600080fd5b813561155481611eaf565b600081518084526020840193506020830160005b828110156121615781516001600160a01b031686526020958601959091019060010161213a565b5093949350505050565b60408152600061217e6040830185612126565b828103602084015280845180835260208301915060208601925060005b818110156121b957835183526020938401939092019160010161219b565b50909695505050505050565b60008060008060008060c087890312156121de57600080fd5b86356121e981611eaf565b955060208701356121f981611eaf565b9450604087013561220981611eaf565b959894975094956060810135955060808101359460a0909101359350915050565b6000806040838503121561223d57600080fd5b823561224881611eaf565b946020939093013593505050565b6020815260006115546020830184612126565b6000806040838503121561227c57600080fd5b823561228781611eaf565b9150602083013561229781611eaf565b809150509250929050565b6000602082840312156122b457600080fd5b8151801515811461155457600080fd5b634e487b7160e01b600052603260045260246000fd5b6000602082840312156122ec57600080fd5b815161155481611eaf565b60006020828403121561230957600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b8181038181111561233957612339612310565b92915050565b8082018082111561233957612339612310565b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052603160045260246000fd5b60005b83811015612399578181015183820152602001612381565b50506000910152565b600081518084526123ba81602086016020860161237e565b601f01601f19169290920160200192915050565b828152604060208201526000611d2560408301846123a2565b600082516123f981846020870161237e565b9190910192915050565b60208152600061155460208301846123a256fea2646970667358221220beafd05d74355d3f422c307870f70963ada2f2b92377a9364aed71fc0ecd000264736f6c634300081b0033",
}

// ContractStrategyManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractStrategyManagerMetaData.ABI instead.
var ContractStrategyManagerABI = ContractStrategyManagerMetaData.ABI

// ContractStrategyManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractStrategyManagerMetaData.Bin instead.
var ContractStrategyManagerBin = ContractStrategyManagerMetaData.Bin

// DeployContractStrategyManager deploys a new Ethereum contract, binding an instance of ContractStrategyManager to it.
func DeployContractStrategyManager(auth *bind.TransactOpts, backend bind.ContractBackend, _delegation common.Address, _pauserRegistry common.Address) (common.Address, *types.Transaction, *ContractStrategyManager, error) {
	parsed, err := ContractStrategyManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractStrategyManagerBin), backend, _delegation, _pauserRegistry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractStrategyManager{ContractStrategyManagerCaller: ContractStrategyManagerCaller{contract: contract}, ContractStrategyManagerTransactor: ContractStrategyManagerTransactor{contract: contract}, ContractStrategyManagerFilterer: ContractStrategyManagerFilterer{contract: contract}}, nil
}

// ContractStrategyManagerMethods is an auto generated interface around an Ethereum contract.
type ContractStrategyManagerMethods interface {
	ContractStrategyManagerCalls
	ContractStrategyManagerTransacts
	ContractStrategyManagerFilters
}

// ContractStrategyManagerCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractStrategyManagerCalls interface {
	DEPOSITTYPEHASH(opts *bind.CallOpts) ([32]byte, error)

	CalculateStrategyDepositDigestHash(opts *bind.CallOpts, staker common.Address, strategy common.Address, token common.Address, amount *big.Int, nonce *big.Int, expiry *big.Int) ([32]byte, error)

	Delegation(opts *bind.CallOpts) (common.Address, error)

	DomainSeparator(opts *bind.CallOpts) ([32]byte, error)

	GetDeposits(opts *bind.CallOpts, staker common.Address) ([]common.Address, []*big.Int, error)

	GetStakerStrategyList(opts *bind.CallOpts, staker common.Address) ([]common.Address, error)

	Nonces(opts *bind.CallOpts, signer common.Address) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts, index uint8) (bool, error)

	Paused0(opts *bind.CallOpts) (*big.Int, error)

	PauserRegistry(opts *bind.CallOpts) (common.Address, error)

	StakerDepositShares(opts *bind.CallOpts, staker common.Address, strategy common.Address) (*big.Int, error)

	StakerStrategyList(opts *bind.CallOpts, staker common.Address, arg1 *big.Int) (common.Address, error)

	StakerStrategyListLength(opts *bind.CallOpts, staker common.Address) (*big.Int, error)

	StrategyIsWhitelistedForDeposit(opts *bind.CallOpts, strategy common.Address) (bool, error)

	StrategyWhitelister(opts *bind.CallOpts) (common.Address, error)
}

// ContractStrategyManagerTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractStrategyManagerTransacts interface {
	AddShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, token common.Address, shares *big.Int) (*types.Transaction, error)

	AddStrategiesToDepositWhitelist(opts *bind.TransactOpts, strategiesToWhitelist []common.Address) (*types.Transaction, error)

	DepositIntoStrategy(opts *bind.TransactOpts, strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error)

	DepositIntoStrategyWithSignature(opts *bind.TransactOpts, strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialStrategyWhitelister common.Address, initialPausedStatus *big.Int) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	PauseAll(opts *bind.TransactOpts) (*types.Transaction, error)

	RemoveDepositShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, depositSharesToRemove *big.Int) (*types.Transaction, error)

	RemoveStrategiesFromDepositWhitelist(opts *bind.TransactOpts, strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetStrategyWhitelister(opts *bind.TransactOpts, newStrategyWhitelister common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	WithdrawSharesAsTokens(opts *bind.TransactOpts, staker common.Address, strategy common.Address, token common.Address, shares *big.Int) (*types.Transaction, error)
}

// ContractStrategyManagerFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractStrategyManagerFilters interface {
	FilterDeposit(opts *bind.FilterOpts) (*ContractStrategyManagerDepositIterator, error)
	WatchDeposit(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerDeposit) (event.Subscription, error)
	ParseDeposit(log types.Log) (*ContractStrategyManagerDeposit, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractStrategyManagerInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractStrategyManagerInitialized, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractStrategyManagerOwnershipTransferredIterator, error)
	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)
	ParseOwnershipTransferred(log types.Log) (*ContractStrategyManagerOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractStrategyManagerPausedIterator, error)
	WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerPaused, account []common.Address) (event.Subscription, error)
	ParsePaused(log types.Log) (*ContractStrategyManagerPaused, error)

	FilterStrategyAddedToDepositWhitelist(opts *bind.FilterOpts) (*ContractStrategyManagerStrategyAddedToDepositWhitelistIterator, error)
	WatchStrategyAddedToDepositWhitelist(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerStrategyAddedToDepositWhitelist) (event.Subscription, error)
	ParseStrategyAddedToDepositWhitelist(log types.Log) (*ContractStrategyManagerStrategyAddedToDepositWhitelist, error)

	FilterStrategyRemovedFromDepositWhitelist(opts *bind.FilterOpts) (*ContractStrategyManagerStrategyRemovedFromDepositWhitelistIterator, error)
	WatchStrategyRemovedFromDepositWhitelist(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerStrategyRemovedFromDepositWhitelist) (event.Subscription, error)
	ParseStrategyRemovedFromDepositWhitelist(log types.Log) (*ContractStrategyManagerStrategyRemovedFromDepositWhitelist, error)

	FilterStrategyWhitelisterChanged(opts *bind.FilterOpts) (*ContractStrategyManagerStrategyWhitelisterChangedIterator, error)
	WatchStrategyWhitelisterChanged(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerStrategyWhitelisterChanged) (event.Subscription, error)
	ParseStrategyWhitelisterChanged(log types.Log) (*ContractStrategyManagerStrategyWhitelisterChanged, error)

	FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractStrategyManagerUnpausedIterator, error)
	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerUnpaused, account []common.Address) (event.Subscription, error)
	ParseUnpaused(log types.Log) (*ContractStrategyManagerUnpaused, error)
}

// ContractStrategyManager is an auto generated Go binding around an Ethereum contract.
type ContractStrategyManager struct {
	ContractStrategyManagerCaller     // Read-only binding to the contract
	ContractStrategyManagerTransactor // Write-only binding to the contract
	ContractStrategyManagerFilterer   // Log filterer for contract events
}

// ContractStrategyManager implements the ContractStrategyManagerMethods interface.
var _ ContractStrategyManagerMethods = (*ContractStrategyManager)(nil)

// ContractStrategyManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractStrategyManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractStrategyManagerCaller implements the ContractStrategyManagerCalls interface.
var _ ContractStrategyManagerCalls = (*ContractStrategyManagerCaller)(nil)

// ContractStrategyManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractStrategyManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractStrategyManagerTransactor implements the ContractStrategyManagerTransacts interface.
var _ ContractStrategyManagerTransacts = (*ContractStrategyManagerTransactor)(nil)

// ContractStrategyManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractStrategyManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractStrategyManagerFilterer implements the ContractStrategyManagerFilters interface.
var _ ContractStrategyManagerFilters = (*ContractStrategyManagerFilterer)(nil)

// ContractStrategyManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractStrategyManagerSession struct {
	Contract     *ContractStrategyManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ContractStrategyManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractStrategyManagerCallerSession struct {
	Contract *ContractStrategyManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// ContractStrategyManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractStrategyManagerTransactorSession struct {
	Contract     *ContractStrategyManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ContractStrategyManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractStrategyManagerRaw struct {
	Contract *ContractStrategyManager // Generic contract binding to access the raw methods on
}

// ContractStrategyManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractStrategyManagerCallerRaw struct {
	Contract *ContractStrategyManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractStrategyManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractStrategyManagerTransactorRaw struct {
	Contract *ContractStrategyManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractStrategyManager creates a new instance of ContractStrategyManager, bound to a specific deployed contract.
func NewContractStrategyManager(address common.Address, backend bind.ContractBackend) (*ContractStrategyManager, error) {
	contract, err := bindContractStrategyManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManager{ContractStrategyManagerCaller: ContractStrategyManagerCaller{contract: contract}, ContractStrategyManagerTransactor: ContractStrategyManagerTransactor{contract: contract}, ContractStrategyManagerFilterer: ContractStrategyManagerFilterer{contract: contract}}, nil
}

// NewContractStrategyManagerCaller creates a new read-only instance of ContractStrategyManager, bound to a specific deployed contract.
func NewContractStrategyManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractStrategyManagerCaller, error) {
	contract, err := bindContractStrategyManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerCaller{contract: contract}, nil
}

// NewContractStrategyManagerTransactor creates a new write-only instance of ContractStrategyManager, bound to a specific deployed contract.
func NewContractStrategyManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractStrategyManagerTransactor, error) {
	contract, err := bindContractStrategyManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerTransactor{contract: contract}, nil
}

// NewContractStrategyManagerFilterer creates a new log filterer instance of ContractStrategyManager, bound to a specific deployed contract.
func NewContractStrategyManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractStrategyManagerFilterer, error) {
	contract, err := bindContractStrategyManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerFilterer{contract: contract}, nil
}

// bindContractStrategyManager binds a generic wrapper to an already deployed contract.
func bindContractStrategyManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractStrategyManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractStrategyManager *ContractStrategyManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractStrategyManager.Contract.ContractStrategyManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractStrategyManager *ContractStrategyManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.ContractStrategyManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractStrategyManager *ContractStrategyManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.ContractStrategyManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractStrategyManager *ContractStrategyManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractStrategyManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractStrategyManager *ContractStrategyManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractStrategyManager *ContractStrategyManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITTYPEHASH is a free data retrieval call binding the contract method 0x48825e94.
//
// Solidity: function DEPOSIT_TYPEHASH() view returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerCaller) DEPOSITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "DEPOSIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEPOSITTYPEHASH is a free data retrieval call binding the contract method 0x48825e94.
//
// Solidity: function DEPOSIT_TYPEHASH() view returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerSession) DEPOSITTYPEHASH() ([32]byte, error) {
	return _ContractStrategyManager.Contract.DEPOSITTYPEHASH(&_ContractStrategyManager.CallOpts)
}

// DEPOSITTYPEHASH is a free data retrieval call binding the contract method 0x48825e94.
//
// Solidity: function DEPOSIT_TYPEHASH() view returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) DEPOSITTYPEHASH() ([32]byte, error) {
	return _ContractStrategyManager.Contract.DEPOSITTYPEHASH(&_ContractStrategyManager.CallOpts)
}

// CalculateStrategyDepositDigestHash is a free data retrieval call binding the contract method 0x9ac01d61.
//
// Solidity: function calculateStrategyDepositDigestHash(address staker, address strategy, address token, uint256 amount, uint256 nonce, uint256 expiry) view returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerCaller) CalculateStrategyDepositDigestHash(opts *bind.CallOpts, staker common.Address, strategy common.Address, token common.Address, amount *big.Int, nonce *big.Int, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "calculateStrategyDepositDigestHash", staker, strategy, token, amount, nonce, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateStrategyDepositDigestHash is a free data retrieval call binding the contract method 0x9ac01d61.
//
// Solidity: function calculateStrategyDepositDigestHash(address staker, address strategy, address token, uint256 amount, uint256 nonce, uint256 expiry) view returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerSession) CalculateStrategyDepositDigestHash(staker common.Address, strategy common.Address, token common.Address, amount *big.Int, nonce *big.Int, expiry *big.Int) ([32]byte, error) {
	return _ContractStrategyManager.Contract.CalculateStrategyDepositDigestHash(&_ContractStrategyManager.CallOpts, staker, strategy, token, amount, nonce, expiry)
}

// CalculateStrategyDepositDigestHash is a free data retrieval call binding the contract method 0x9ac01d61.
//
// Solidity: function calculateStrategyDepositDigestHash(address staker, address strategy, address token, uint256 amount, uint256 nonce, uint256 expiry) view returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) CalculateStrategyDepositDigestHash(staker common.Address, strategy common.Address, token common.Address, amount *big.Int, nonce *big.Int, expiry *big.Int) ([32]byte, error) {
	return _ContractStrategyManager.Contract.CalculateStrategyDepositDigestHash(&_ContractStrategyManager.CallOpts, staker, strategy, token, amount, nonce, expiry)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerSession) Delegation() (common.Address, error) {
	return _ContractStrategyManager.Contract.Delegation(&_ContractStrategyManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractStrategyManager.Contract.Delegation(&_ContractStrategyManager.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerSession) DomainSeparator() ([32]byte, error) {
	return _ContractStrategyManager.Contract.DomainSeparator(&_ContractStrategyManager.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) DomainSeparator() ([32]byte, error) {
	return _ContractStrategyManager.Contract.DomainSeparator(&_ContractStrategyManager.CallOpts)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address staker) view returns(address[], uint256[])
func (_ContractStrategyManager *ContractStrategyManagerCaller) GetDeposits(opts *bind.CallOpts, staker common.Address) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "getDeposits", staker)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address staker) view returns(address[], uint256[])
func (_ContractStrategyManager *ContractStrategyManagerSession) GetDeposits(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _ContractStrategyManager.Contract.GetDeposits(&_ContractStrategyManager.CallOpts, staker)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address staker) view returns(address[], uint256[])
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) GetDeposits(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _ContractStrategyManager.Contract.GetDeposits(&_ContractStrategyManager.CallOpts, staker)
}

// GetStakerStrategyList is a free data retrieval call binding the contract method 0xde44acb6.
//
// Solidity: function getStakerStrategyList(address staker) view returns(address[])
func (_ContractStrategyManager *ContractStrategyManagerCaller) GetStakerStrategyList(opts *bind.CallOpts, staker common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "getStakerStrategyList", staker)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetStakerStrategyList is a free data retrieval call binding the contract method 0xde44acb6.
//
// Solidity: function getStakerStrategyList(address staker) view returns(address[])
func (_ContractStrategyManager *ContractStrategyManagerSession) GetStakerStrategyList(staker common.Address) ([]common.Address, error) {
	return _ContractStrategyManager.Contract.GetStakerStrategyList(&_ContractStrategyManager.CallOpts, staker)
}

// GetStakerStrategyList is a free data retrieval call binding the contract method 0xde44acb6.
//
// Solidity: function getStakerStrategyList(address staker) view returns(address[])
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) GetStakerStrategyList(staker common.Address) ([]common.Address, error) {
	return _ContractStrategyManager.Contract.GetStakerStrategyList(&_ContractStrategyManager.CallOpts, staker)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address signer) view returns(uint256 nonce)
func (_ContractStrategyManager *ContractStrategyManagerCaller) Nonces(opts *bind.CallOpts, signer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "nonces", signer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address signer) view returns(uint256 nonce)
func (_ContractStrategyManager *ContractStrategyManagerSession) Nonces(signer common.Address) (*big.Int, error) {
	return _ContractStrategyManager.Contract.Nonces(&_ContractStrategyManager.CallOpts, signer)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address signer) view returns(uint256 nonce)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) Nonces(signer common.Address) (*big.Int, error) {
	return _ContractStrategyManager.Contract.Nonces(&_ContractStrategyManager.CallOpts, signer)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerSession) Owner() (common.Address, error) {
	return _ContractStrategyManager.Contract.Owner(&_ContractStrategyManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) Owner() (common.Address, error) {
	return _ContractStrategyManager.Contract.Owner(&_ContractStrategyManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractStrategyManager *ContractStrategyManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractStrategyManager *ContractStrategyManagerSession) Paused(index uint8) (bool, error) {
	return _ContractStrategyManager.Contract.Paused(&_ContractStrategyManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractStrategyManager.Contract.Paused(&_ContractStrategyManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerSession) Paused0() (*big.Int, error) {
	return _ContractStrategyManager.Contract.Paused0(&_ContractStrategyManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractStrategyManager.Contract.Paused0(&_ContractStrategyManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractStrategyManager.Contract.PauserRegistry(&_ContractStrategyManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractStrategyManager.Contract.PauserRegistry(&_ContractStrategyManager.CallOpts)
}

// StakerDepositShares is a free data retrieval call binding the contract method 0xfe243a17.
//
// Solidity: function stakerDepositShares(address staker, address strategy) view returns(uint256 shares)
func (_ContractStrategyManager *ContractStrategyManagerCaller) StakerDepositShares(opts *bind.CallOpts, staker common.Address, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "stakerDepositShares", staker, strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerDepositShares is a free data retrieval call binding the contract method 0xfe243a17.
//
// Solidity: function stakerDepositShares(address staker, address strategy) view returns(uint256 shares)
func (_ContractStrategyManager *ContractStrategyManagerSession) StakerDepositShares(staker common.Address, strategy common.Address) (*big.Int, error) {
	return _ContractStrategyManager.Contract.StakerDepositShares(&_ContractStrategyManager.CallOpts, staker, strategy)
}

// StakerDepositShares is a free data retrieval call binding the contract method 0xfe243a17.
//
// Solidity: function stakerDepositShares(address staker, address strategy) view returns(uint256 shares)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) StakerDepositShares(staker common.Address, strategy common.Address) (*big.Int, error) {
	return _ContractStrategyManager.Contract.StakerDepositShares(&_ContractStrategyManager.CallOpts, staker, strategy)
}

// StakerStrategyList is a free data retrieval call binding the contract method 0xcbc2bd62.
//
// Solidity: function stakerStrategyList(address staker, uint256 ) view returns(address strategies)
func (_ContractStrategyManager *ContractStrategyManagerCaller) StakerStrategyList(opts *bind.CallOpts, staker common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "stakerStrategyList", staker, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakerStrategyList is a free data retrieval call binding the contract method 0xcbc2bd62.
//
// Solidity: function stakerStrategyList(address staker, uint256 ) view returns(address strategies)
func (_ContractStrategyManager *ContractStrategyManagerSession) StakerStrategyList(staker common.Address, arg1 *big.Int) (common.Address, error) {
	return _ContractStrategyManager.Contract.StakerStrategyList(&_ContractStrategyManager.CallOpts, staker, arg1)
}

// StakerStrategyList is a free data retrieval call binding the contract method 0xcbc2bd62.
//
// Solidity: function stakerStrategyList(address staker, uint256 ) view returns(address strategies)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) StakerStrategyList(staker common.Address, arg1 *big.Int) (common.Address, error) {
	return _ContractStrategyManager.Contract.StakerStrategyList(&_ContractStrategyManager.CallOpts, staker, arg1)
}

// StakerStrategyListLength is a free data retrieval call binding the contract method 0x8b8aac3c.
//
// Solidity: function stakerStrategyListLength(address staker) view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerCaller) StakerStrategyListLength(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "stakerStrategyListLength", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerStrategyListLength is a free data retrieval call binding the contract method 0x8b8aac3c.
//
// Solidity: function stakerStrategyListLength(address staker) view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerSession) StakerStrategyListLength(staker common.Address) (*big.Int, error) {
	return _ContractStrategyManager.Contract.StakerStrategyListLength(&_ContractStrategyManager.CallOpts, staker)
}

// StakerStrategyListLength is a free data retrieval call binding the contract method 0x8b8aac3c.
//
// Solidity: function stakerStrategyListLength(address staker) view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) StakerStrategyListLength(staker common.Address) (*big.Int, error) {
	return _ContractStrategyManager.Contract.StakerStrategyListLength(&_ContractStrategyManager.CallOpts, staker)
}

// StrategyIsWhitelistedForDeposit is a free data retrieval call binding the contract method 0x663c1de4.
//
// Solidity: function strategyIsWhitelistedForDeposit(address strategy) view returns(bool whitelisted)
func (_ContractStrategyManager *ContractStrategyManagerCaller) StrategyIsWhitelistedForDeposit(opts *bind.CallOpts, strategy common.Address) (bool, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "strategyIsWhitelistedForDeposit", strategy)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StrategyIsWhitelistedForDeposit is a free data retrieval call binding the contract method 0x663c1de4.
//
// Solidity: function strategyIsWhitelistedForDeposit(address strategy) view returns(bool whitelisted)
func (_ContractStrategyManager *ContractStrategyManagerSession) StrategyIsWhitelistedForDeposit(strategy common.Address) (bool, error) {
	return _ContractStrategyManager.Contract.StrategyIsWhitelistedForDeposit(&_ContractStrategyManager.CallOpts, strategy)
}

// StrategyIsWhitelistedForDeposit is a free data retrieval call binding the contract method 0x663c1de4.
//
// Solidity: function strategyIsWhitelistedForDeposit(address strategy) view returns(bool whitelisted)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) StrategyIsWhitelistedForDeposit(strategy common.Address) (bool, error) {
	return _ContractStrategyManager.Contract.StrategyIsWhitelistedForDeposit(&_ContractStrategyManager.CallOpts, strategy)
}

// StrategyWhitelister is a free data retrieval call binding the contract method 0x967fc0d2.
//
// Solidity: function strategyWhitelister() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCaller) StrategyWhitelister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "strategyWhitelister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyWhitelister is a free data retrieval call binding the contract method 0x967fc0d2.
//
// Solidity: function strategyWhitelister() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerSession) StrategyWhitelister() (common.Address, error) {
	return _ContractStrategyManager.Contract.StrategyWhitelister(&_ContractStrategyManager.CallOpts)
}

// StrategyWhitelister is a free data retrieval call binding the contract method 0x967fc0d2.
//
// Solidity: function strategyWhitelister() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) StrategyWhitelister() (common.Address, error) {
	return _ContractStrategyManager.Contract.StrategyWhitelister(&_ContractStrategyManager.CallOpts)
}

// AddShares is a paid mutator transaction binding the contract method 0xc4623ea1.
//
// Solidity: function addShares(address staker, address strategy, address token, uint256 shares) returns(uint256, uint256)
func (_ContractStrategyManager *ContractStrategyManagerTransactor) AddShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, token common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "addShares", staker, strategy, token, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0xc4623ea1.
//
// Solidity: function addShares(address staker, address strategy, address token, uint256 shares) returns(uint256, uint256)
func (_ContractStrategyManager *ContractStrategyManagerSession) AddShares(staker common.Address, strategy common.Address, token common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.AddShares(&_ContractStrategyManager.TransactOpts, staker, strategy, token, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0xc4623ea1.
//
// Solidity: function addShares(address staker, address strategy, address token, uint256 shares) returns(uint256, uint256)
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) AddShares(staker common.Address, strategy common.Address, token common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.AddShares(&_ContractStrategyManager.TransactOpts, staker, strategy, token, shares)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0x5de08ff2.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) AddStrategiesToDepositWhitelist(opts *bind.TransactOpts, strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "addStrategiesToDepositWhitelist", strategiesToWhitelist)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0x5de08ff2.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) AddStrategiesToDepositWhitelist(strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.AddStrategiesToDepositWhitelist(&_ContractStrategyManager.TransactOpts, strategiesToWhitelist)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0x5de08ff2.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) AddStrategiesToDepositWhitelist(strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.AddStrategiesToDepositWhitelist(&_ContractStrategyManager.TransactOpts, strategiesToWhitelist)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xe7a050aa.
//
// Solidity: function depositIntoStrategy(address strategy, address token, uint256 amount) returns(uint256 depositedShares)
func (_ContractStrategyManager *ContractStrategyManagerTransactor) DepositIntoStrategy(opts *bind.TransactOpts, strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "depositIntoStrategy", strategy, token, amount)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xe7a050aa.
//
// Solidity: function depositIntoStrategy(address strategy, address token, uint256 amount) returns(uint256 depositedShares)
func (_ContractStrategyManager *ContractStrategyManagerSession) DepositIntoStrategy(strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.DepositIntoStrategy(&_ContractStrategyManager.TransactOpts, strategy, token, amount)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xe7a050aa.
//
// Solidity: function depositIntoStrategy(address strategy, address token, uint256 amount) returns(uint256 depositedShares)
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) DepositIntoStrategy(strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.DepositIntoStrategy(&_ContractStrategyManager.TransactOpts, strategy, token, amount)
}

// DepositIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x32e89ace.
//
// Solidity: function depositIntoStrategyWithSignature(address strategy, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 depositedShares)
func (_ContractStrategyManager *ContractStrategyManagerTransactor) DepositIntoStrategyWithSignature(opts *bind.TransactOpts, strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "depositIntoStrategyWithSignature", strategy, token, amount, staker, expiry, signature)
}

// DepositIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x32e89ace.
//
// Solidity: function depositIntoStrategyWithSignature(address strategy, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 depositedShares)
func (_ContractStrategyManager *ContractStrategyManagerSession) DepositIntoStrategyWithSignature(strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.DepositIntoStrategyWithSignature(&_ContractStrategyManager.TransactOpts, strategy, token, amount, staker, expiry, signature)
}

// DepositIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x32e89ace.
//
// Solidity: function depositIntoStrategyWithSignature(address strategy, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 depositedShares)
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) DepositIntoStrategyWithSignature(strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.DepositIntoStrategyWithSignature(&_ContractStrategyManager.TransactOpts, strategy, token, amount, staker, expiry, signature)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address initialStrategyWhitelister, uint256 initialPausedStatus) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialStrategyWhitelister common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "initialize", initialOwner, initialStrategyWhitelister, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address initialStrategyWhitelister, uint256 initialPausedStatus) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) Initialize(initialOwner common.Address, initialStrategyWhitelister common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.Initialize(&_ContractStrategyManager.TransactOpts, initialOwner, initialStrategyWhitelister, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address initialStrategyWhitelister, uint256 initialPausedStatus) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) Initialize(initialOwner common.Address, initialStrategyWhitelister common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.Initialize(&_ContractStrategyManager.TransactOpts, initialOwner, initialStrategyWhitelister, initialPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.Pause(&_ContractStrategyManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.Pause(&_ContractStrategyManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.PauseAll(&_ContractStrategyManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.PauseAll(&_ContractStrategyManager.TransactOpts)
}

// RemoveDepositShares is a paid mutator transaction binding the contract method 0x724af423.
//
// Solidity: function removeDepositShares(address staker, address strategy, uint256 depositSharesToRemove) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) RemoveDepositShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, depositSharesToRemove *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "removeDepositShares", staker, strategy, depositSharesToRemove)
}

// RemoveDepositShares is a paid mutator transaction binding the contract method 0x724af423.
//
// Solidity: function removeDepositShares(address staker, address strategy, uint256 depositSharesToRemove) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) RemoveDepositShares(staker common.Address, strategy common.Address, depositSharesToRemove *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.RemoveDepositShares(&_ContractStrategyManager.TransactOpts, staker, strategy, depositSharesToRemove)
}

// RemoveDepositShares is a paid mutator transaction binding the contract method 0x724af423.
//
// Solidity: function removeDepositShares(address staker, address strategy, uint256 depositSharesToRemove) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) RemoveDepositShares(staker common.Address, strategy common.Address, depositSharesToRemove *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.RemoveDepositShares(&_ContractStrategyManager.TransactOpts, staker, strategy, depositSharesToRemove)
}

// RemoveStrategiesFromDepositWhitelist is a paid mutator transaction binding the contract method 0xb5d8b5b8.
//
// Solidity: function removeStrategiesFromDepositWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) RemoveStrategiesFromDepositWhitelist(opts *bind.TransactOpts, strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "removeStrategiesFromDepositWhitelist", strategiesToRemoveFromWhitelist)
}

// RemoveStrategiesFromDepositWhitelist is a paid mutator transaction binding the contract method 0xb5d8b5b8.
//
// Solidity: function removeStrategiesFromDepositWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) RemoveStrategiesFromDepositWhitelist(strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.RemoveStrategiesFromDepositWhitelist(&_ContractStrategyManager.TransactOpts, strategiesToRemoveFromWhitelist)
}

// RemoveStrategiesFromDepositWhitelist is a paid mutator transaction binding the contract method 0xb5d8b5b8.
//
// Solidity: function removeStrategiesFromDepositWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) RemoveStrategiesFromDepositWhitelist(strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.RemoveStrategiesFromDepositWhitelist(&_ContractStrategyManager.TransactOpts, strategiesToRemoveFromWhitelist)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.RenounceOwnership(&_ContractStrategyManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.RenounceOwnership(&_ContractStrategyManager.TransactOpts)
}

// SetStrategyWhitelister is a paid mutator transaction binding the contract method 0xc6656702.
//
// Solidity: function setStrategyWhitelister(address newStrategyWhitelister) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) SetStrategyWhitelister(opts *bind.TransactOpts, newStrategyWhitelister common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "setStrategyWhitelister", newStrategyWhitelister)
}

// SetStrategyWhitelister is a paid mutator transaction binding the contract method 0xc6656702.
//
// Solidity: function setStrategyWhitelister(address newStrategyWhitelister) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) SetStrategyWhitelister(newStrategyWhitelister common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.SetStrategyWhitelister(&_ContractStrategyManager.TransactOpts, newStrategyWhitelister)
}

// SetStrategyWhitelister is a paid mutator transaction binding the contract method 0xc6656702.
//
// Solidity: function setStrategyWhitelister(address newStrategyWhitelister) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) SetStrategyWhitelister(newStrategyWhitelister common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.SetStrategyWhitelister(&_ContractStrategyManager.TransactOpts, newStrategyWhitelister)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.TransferOwnership(&_ContractStrategyManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.TransferOwnership(&_ContractStrategyManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.Unpause(&_ContractStrategyManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.Unpause(&_ContractStrategyManager.TransactOpts, newPausedStatus)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x2eae418c.
//
// Solidity: function withdrawSharesAsTokens(address staker, address strategy, address token, uint256 shares) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) WithdrawSharesAsTokens(opts *bind.TransactOpts, staker common.Address, strategy common.Address, token common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "withdrawSharesAsTokens", staker, strategy, token, shares)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x2eae418c.
//
// Solidity: function withdrawSharesAsTokens(address staker, address strategy, address token, uint256 shares) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) WithdrawSharesAsTokens(staker common.Address, strategy common.Address, token common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.WithdrawSharesAsTokens(&_ContractStrategyManager.TransactOpts, staker, strategy, token, shares)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x2eae418c.
//
// Solidity: function withdrawSharesAsTokens(address staker, address strategy, address token, uint256 shares) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) WithdrawSharesAsTokens(staker common.Address, strategy common.Address, token common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.WithdrawSharesAsTokens(&_ContractStrategyManager.TransactOpts, staker, strategy, token, shares)
}

// ContractStrategyManagerDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ContractStrategyManager contract.
type ContractStrategyManagerDepositIterator struct {
	Event *ContractStrategyManagerDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractStrategyManagerDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyManagerDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractStrategyManagerDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractStrategyManagerDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyManagerDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyManagerDeposit represents a Deposit event raised by the ContractStrategyManager contract.
type ContractStrategyManagerDeposit struct {
	Staker   common.Address
	Token    common.Address
	Strategy common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address staker, address token, address strategy, uint256 shares)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) FilterDeposit(opts *bind.FilterOpts) (*ContractStrategyManagerDepositIterator, error) {

	logs, sub, err := _ContractStrategyManager.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerDepositIterator{contract: _ContractStrategyManager.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address staker, address token, address strategy, uint256 shares)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerDeposit) (event.Subscription, error) {

	logs, sub, err := _ContractStrategyManager.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyManagerDeposit)
				if err := _ContractStrategyManager.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address staker, address token, address strategy, uint256 shares)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) ParseDeposit(log types.Log) (*ContractStrategyManagerDeposit, error) {
	event := new(ContractStrategyManagerDeposit)
	if err := _ContractStrategyManager.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStrategyManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractStrategyManager contract.
type ContractStrategyManagerInitializedIterator struct {
	Event *ContractStrategyManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractStrategyManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractStrategyManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractStrategyManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyManagerInitialized represents a Initialized event raised by the ContractStrategyManager contract.
type ContractStrategyManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractStrategyManagerInitializedIterator, error) {

	logs, sub, err := _ContractStrategyManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerInitializedIterator{contract: _ContractStrategyManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractStrategyManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyManagerInitialized)
				if err := _ContractStrategyManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) ParseInitialized(log types.Log) (*ContractStrategyManagerInitialized, error) {
	event := new(ContractStrategyManagerInitialized)
	if err := _ContractStrategyManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStrategyManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractStrategyManager contract.
type ContractStrategyManagerOwnershipTransferredIterator struct {
	Event *ContractStrategyManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractStrategyManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyManagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractStrategyManagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractStrategyManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractStrategyManager contract.
type ContractStrategyManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractStrategyManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractStrategyManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerOwnershipTransferredIterator{contract: _ContractStrategyManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractStrategyManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyManagerOwnershipTransferred)
				if err := _ContractStrategyManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractStrategyManagerOwnershipTransferred, error) {
	event := new(ContractStrategyManagerOwnershipTransferred)
	if err := _ContractStrategyManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStrategyManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractStrategyManager contract.
type ContractStrategyManagerPausedIterator struct {
	Event *ContractStrategyManagerPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractStrategyManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyManagerPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractStrategyManagerPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractStrategyManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyManagerPaused represents a Paused event raised by the ContractStrategyManager contract.
type ContractStrategyManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractStrategyManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractStrategyManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerPausedIterator{contract: _ContractStrategyManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractStrategyManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyManagerPaused)
				if err := _ContractStrategyManager.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) ParsePaused(log types.Log) (*ContractStrategyManagerPaused, error) {
	event := new(ContractStrategyManagerPaused)
	if err := _ContractStrategyManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStrategyManagerStrategyAddedToDepositWhitelistIterator is returned from FilterStrategyAddedToDepositWhitelist and is used to iterate over the raw logs and unpacked data for StrategyAddedToDepositWhitelist events raised by the ContractStrategyManager contract.
type ContractStrategyManagerStrategyAddedToDepositWhitelistIterator struct {
	Event *ContractStrategyManagerStrategyAddedToDepositWhitelist // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractStrategyManagerStrategyAddedToDepositWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyManagerStrategyAddedToDepositWhitelist)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractStrategyManagerStrategyAddedToDepositWhitelist)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractStrategyManagerStrategyAddedToDepositWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyManagerStrategyAddedToDepositWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyManagerStrategyAddedToDepositWhitelist represents a StrategyAddedToDepositWhitelist event raised by the ContractStrategyManager contract.
type ContractStrategyManagerStrategyAddedToDepositWhitelist struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyAddedToDepositWhitelist is a free log retrieval operation binding the contract event 0x0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe.
//
// Solidity: event StrategyAddedToDepositWhitelist(address strategy)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) FilterStrategyAddedToDepositWhitelist(opts *bind.FilterOpts) (*ContractStrategyManagerStrategyAddedToDepositWhitelistIterator, error) {

	logs, sub, err := _ContractStrategyManager.contract.FilterLogs(opts, "StrategyAddedToDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerStrategyAddedToDepositWhitelistIterator{contract: _ContractStrategyManager.contract, event: "StrategyAddedToDepositWhitelist", logs: logs, sub: sub}, nil
}

// WatchStrategyAddedToDepositWhitelist is a free log subscription operation binding the contract event 0x0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe.
//
// Solidity: event StrategyAddedToDepositWhitelist(address strategy)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) WatchStrategyAddedToDepositWhitelist(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerStrategyAddedToDepositWhitelist) (event.Subscription, error) {

	logs, sub, err := _ContractStrategyManager.contract.WatchLogs(opts, "StrategyAddedToDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyManagerStrategyAddedToDepositWhitelist)
				if err := _ContractStrategyManager.contract.UnpackLog(event, "StrategyAddedToDepositWhitelist", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStrategyAddedToDepositWhitelist is a log parse operation binding the contract event 0x0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe.
//
// Solidity: event StrategyAddedToDepositWhitelist(address strategy)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) ParseStrategyAddedToDepositWhitelist(log types.Log) (*ContractStrategyManagerStrategyAddedToDepositWhitelist, error) {
	event := new(ContractStrategyManagerStrategyAddedToDepositWhitelist)
	if err := _ContractStrategyManager.contract.UnpackLog(event, "StrategyAddedToDepositWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStrategyManagerStrategyRemovedFromDepositWhitelistIterator is returned from FilterStrategyRemovedFromDepositWhitelist and is used to iterate over the raw logs and unpacked data for StrategyRemovedFromDepositWhitelist events raised by the ContractStrategyManager contract.
type ContractStrategyManagerStrategyRemovedFromDepositWhitelistIterator struct {
	Event *ContractStrategyManagerStrategyRemovedFromDepositWhitelist // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractStrategyManagerStrategyRemovedFromDepositWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyManagerStrategyRemovedFromDepositWhitelist)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractStrategyManagerStrategyRemovedFromDepositWhitelist)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractStrategyManagerStrategyRemovedFromDepositWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyManagerStrategyRemovedFromDepositWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyManagerStrategyRemovedFromDepositWhitelist represents a StrategyRemovedFromDepositWhitelist event raised by the ContractStrategyManager contract.
type ContractStrategyManagerStrategyRemovedFromDepositWhitelist struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyRemovedFromDepositWhitelist is a free log retrieval operation binding the contract event 0x4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030.
//
// Solidity: event StrategyRemovedFromDepositWhitelist(address strategy)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) FilterStrategyRemovedFromDepositWhitelist(opts *bind.FilterOpts) (*ContractStrategyManagerStrategyRemovedFromDepositWhitelistIterator, error) {

	logs, sub, err := _ContractStrategyManager.contract.FilterLogs(opts, "StrategyRemovedFromDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerStrategyRemovedFromDepositWhitelistIterator{contract: _ContractStrategyManager.contract, event: "StrategyRemovedFromDepositWhitelist", logs: logs, sub: sub}, nil
}

// WatchStrategyRemovedFromDepositWhitelist is a free log subscription operation binding the contract event 0x4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030.
//
// Solidity: event StrategyRemovedFromDepositWhitelist(address strategy)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) WatchStrategyRemovedFromDepositWhitelist(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerStrategyRemovedFromDepositWhitelist) (event.Subscription, error) {

	logs, sub, err := _ContractStrategyManager.contract.WatchLogs(opts, "StrategyRemovedFromDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyManagerStrategyRemovedFromDepositWhitelist)
				if err := _ContractStrategyManager.contract.UnpackLog(event, "StrategyRemovedFromDepositWhitelist", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStrategyRemovedFromDepositWhitelist is a log parse operation binding the contract event 0x4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030.
//
// Solidity: event StrategyRemovedFromDepositWhitelist(address strategy)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) ParseStrategyRemovedFromDepositWhitelist(log types.Log) (*ContractStrategyManagerStrategyRemovedFromDepositWhitelist, error) {
	event := new(ContractStrategyManagerStrategyRemovedFromDepositWhitelist)
	if err := _ContractStrategyManager.contract.UnpackLog(event, "StrategyRemovedFromDepositWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStrategyManagerStrategyWhitelisterChangedIterator is returned from FilterStrategyWhitelisterChanged and is used to iterate over the raw logs and unpacked data for StrategyWhitelisterChanged events raised by the ContractStrategyManager contract.
type ContractStrategyManagerStrategyWhitelisterChangedIterator struct {
	Event *ContractStrategyManagerStrategyWhitelisterChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractStrategyManagerStrategyWhitelisterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyManagerStrategyWhitelisterChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractStrategyManagerStrategyWhitelisterChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractStrategyManagerStrategyWhitelisterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyManagerStrategyWhitelisterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyManagerStrategyWhitelisterChanged represents a StrategyWhitelisterChanged event raised by the ContractStrategyManager contract.
type ContractStrategyManagerStrategyWhitelisterChanged struct {
	PreviousAddress common.Address
	NewAddress      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStrategyWhitelisterChanged is a free log retrieval operation binding the contract event 0x4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29.
//
// Solidity: event StrategyWhitelisterChanged(address previousAddress, address newAddress)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) FilterStrategyWhitelisterChanged(opts *bind.FilterOpts) (*ContractStrategyManagerStrategyWhitelisterChangedIterator, error) {

	logs, sub, err := _ContractStrategyManager.contract.FilterLogs(opts, "StrategyWhitelisterChanged")
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerStrategyWhitelisterChangedIterator{contract: _ContractStrategyManager.contract, event: "StrategyWhitelisterChanged", logs: logs, sub: sub}, nil
}

// WatchStrategyWhitelisterChanged is a free log subscription operation binding the contract event 0x4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29.
//
// Solidity: event StrategyWhitelisterChanged(address previousAddress, address newAddress)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) WatchStrategyWhitelisterChanged(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerStrategyWhitelisterChanged) (event.Subscription, error) {

	logs, sub, err := _ContractStrategyManager.contract.WatchLogs(opts, "StrategyWhitelisterChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyManagerStrategyWhitelisterChanged)
				if err := _ContractStrategyManager.contract.UnpackLog(event, "StrategyWhitelisterChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStrategyWhitelisterChanged is a log parse operation binding the contract event 0x4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29.
//
// Solidity: event StrategyWhitelisterChanged(address previousAddress, address newAddress)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) ParseStrategyWhitelisterChanged(log types.Log) (*ContractStrategyManagerStrategyWhitelisterChanged, error) {
	event := new(ContractStrategyManagerStrategyWhitelisterChanged)
	if err := _ContractStrategyManager.contract.UnpackLog(event, "StrategyWhitelisterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStrategyManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractStrategyManager contract.
type ContractStrategyManagerUnpausedIterator struct {
	Event *ContractStrategyManagerUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractStrategyManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyManagerUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractStrategyManagerUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractStrategyManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyManagerUnpaused represents a Unpaused event raised by the ContractStrategyManager contract.
type ContractStrategyManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractStrategyManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractStrategyManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerUnpausedIterator{contract: _ContractStrategyManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractStrategyManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyManagerUnpaused)
				if err := _ContractStrategyManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) ParseUnpaused(log types.Log) (*ContractStrategyManagerUnpaused, error) {
	event := new(ContractStrategyManagerUnpaused)
	if err := _ContractStrategyManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
