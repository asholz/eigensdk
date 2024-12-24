// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import "eigenlayer-contracts/src/contracts/permissions/PauserRegistry.sol";

import {IDelegationManager} from "eigenlayer-contracts/src/contracts/interfaces/IDelegationManager.sol";
import {IStrategyManager, IStrategy} from "eigenlayer-contracts/src/contracts/interfaces/IStrategyManager.sol";
import "eigenlayer-contracts/src/test/mocks/EmptyContract.sol";

import "eigenlayer-middleware/src/RegistryCoordinator.sol" as blsregcoord;
import {IServiceManager} from "eigenlayer-middleware/src/interfaces/IServiceManager.sol";
import {IBLSApkRegistry, IIndexRegistry, IStakeRegistry} from "eigenlayer-middleware/src/RegistryCoordinator.sol";
import {BLSApkRegistry} from "eigenlayer-middleware/src/BLSApkRegistry.sol";
import {IndexRegistry} from "eigenlayer-middleware/src/IndexRegistry.sol";
import {StakeRegistry} from "eigenlayer-middleware/src/StakeRegistry.sol";
import {OperatorStateRetriever} from "eigenlayer-middleware/src/OperatorStateRetriever.sol";

import {MockAvsContracts} from "./parsers/MockAvsContractsParser.sol";
import {EigenlayerContracts, EigenlayerContractsParser} from "./parsers/EigenlayerContractsParser.sol";
import {ConfigsReadWriter} from "./parsers/ConfigsReadWriter.sol";
import {MockAvsServiceManager} from "../src/MockAvsServiceManager.sol";
import {ContractsRegistry} from "../src/ContractsRegistry.sol";

import "forge-std/Script.sol";
import "forge-std/StdJson.sol";

contract DeployMockAvsRegistries is
    Script,
    ConfigsReadWriter,
    EigenlayerContractsParser
{
    // MockAvs contracts
    ProxyAdmin public mockAvsProxyAdmin;
    PauserRegistry public mockAvsPauserReg;
    blsregcoord.RegistryCoordinator public registryCoordinator;
    blsregcoord.IRegistryCoordinator public registryCoordinatorImplementation;
    IBLSApkRegistry public blsApkRegistry;
    IBLSApkRegistry public blsApkRegistryImplementation;
    IIndexRegistry public indexRegistry;
    IIndexRegistry public indexRegistryImplementation;
    IStakeRegistry public stakeRegistry;
    IStakeRegistry public stakeRegistryImplementation;
    OperatorStateRetriever public operatorStateRetriever;
    EmptyContract public emptyContract;

    struct MockAvsOpsAddresses {
        address communityMultisig;
        address pauser;
        address churner;
        address ejector;
    }

    function _loadAvsOpsAddresses(
        string memory opsAddressesFileName
    ) internal view returns (MockAvsOpsAddresses memory) {
        string memory opsAddresses = readInput(opsAddressesFileName);
        MockAvsOpsAddresses memory addressConfig;
        addressConfig.communityMultisig = stdJson.readAddress(
            opsAddresses,
            ".communityMultisig"
        );
        addressConfig.pauser = stdJson.readAddress(opsAddresses, ".pauser");
        addressConfig.churner = stdJson.readAddress(opsAddresses, ".churner");
        addressConfig.ejector = stdJson.readAddress(opsAddresses, ".ejector");
        return addressConfig;
    }

    function _deploymockAvsRegistryContracts(
        EigenlayerContracts memory eigenlayerContracts,
        MockAvsOpsAddresses memory addressConfig,
        MockAvsServiceManager mockAvsServiceManager,
        MockAvsServiceManager mockAvsServiceManagerImplementation
    ) internal returns (MockAvsContracts memory) {
        {   // Pauser registry scope
            address[] memory pausers = new address[](2);
            pausers[0] = addressConfig.pauser;
            pausers[1] = addressConfig.communityMultisig;
            mockAvsPauserReg = new PauserRegistry(
                pausers,
                addressConfig.communityMultisig
            );
        }

        {   // Proxy deployments scope
            registryCoordinator = blsregcoord.RegistryCoordinator(
                address(
                    new TransparentUpgradeableProxy(
                        address(emptyContract),
                        address(mockAvsProxyAdmin),
                        ""
                    )
                )
            );
            blsApkRegistry = IBLSApkRegistry(
                address(
                    new TransparentUpgradeableProxy(
                        address(emptyContract),
                        address(mockAvsProxyAdmin),
                        ""
                    )
                )
            );
            indexRegistry = IIndexRegistry(
                address(
                    new TransparentUpgradeableProxy(
                        address(emptyContract),
                        address(mockAvsProxyAdmin),
                        ""
                    )
                )
            );
            stakeRegistry = IStakeRegistry(
                address(
                    new TransparentUpgradeableProxy(
                        address(emptyContract),
                        address(mockAvsProxyAdmin),
                        ""
                    )
                )
            );
        }

        operatorStateRetriever = new OperatorStateRetriever();

        {   // Implementation deployments scope
            blsApkRegistryImplementation = new BLSApkRegistry(registryCoordinator);

            mockAvsProxyAdmin.upgrade(
                TransparentUpgradeableProxy(payable(address(blsApkRegistry))),
                address(blsApkRegistryImplementation)
            );

            indexRegistryImplementation = new IndexRegistry(registryCoordinator);

            mockAvsProxyAdmin.upgrade(
                TransparentUpgradeableProxy(payable(address(indexRegistry))),
                address(indexRegistryImplementation)
            );

            {   // Stake registry scope
                stakeRegistryImplementation = new StakeRegistry(
                    registryCoordinator,
                    eigenlayerContracts.delegationManager,
                    eigenlayerContracts.avsDirectory,
                    mockAvsServiceManager
                );

                mockAvsProxyAdmin.upgrade(
                    TransparentUpgradeableProxy(payable(address(stakeRegistry))),
                    address(stakeRegistryImplementation)
                );
            }

            registryCoordinatorImplementation = new blsregcoord.RegistryCoordinator(
                blsregcoord.IServiceManager(address(mockAvsServiceManager)),
                blsregcoord.IStakeRegistry(address(stakeRegistry)),
                blsregcoord.IBLSApkRegistry(address(blsApkRegistry)),
                blsregcoord.IIndexRegistry(address(indexRegistry)),
                eigenlayerContracts.avsDirectory,
                mockAvsPauserReg
            );
        }

        {   // Registry coordinator initialization scope
            uint numQuorums = 0;
            blsregcoord.RegistryCoordinator.OperatorSetParam[]
                memory quorumsOperatorSetParams = new blsregcoord.RegistryCoordinator.OperatorSetParam[](
                    numQuorums
                );
            for (uint i = 0; i < numQuorums; i++) {
                quorumsOperatorSetParams[i] = blsregcoord
                    .IRegistryCoordinator
                    .OperatorSetParam({
                        maxOperatorCount: 10000,
                        kickBIPsOfOperatorStake: 15000,
                        kickBIPsOfTotalStake: 100
                    });
            }
            uint96[] memory quorumsMinimumStake = new uint96[](numQuorums);
            IStakeRegistry.StrategyParams[][]
                memory quorumsStrategyParams = new IStakeRegistry.StrategyParams[][](
                    numQuorums
                );

            mockAvsProxyAdmin.upgradeAndCall(
                TransparentUpgradeableProxy(
                    payable(address(registryCoordinator))
                ),
                address(registryCoordinatorImplementation),
                abi.encodeWithSelector(
                    blsregcoord.RegistryCoordinator.initialize.selector,
                    addressConfig.communityMultisig,
                    addressConfig.churner,
                    addressConfig.ejector,
                    addressConfig.pauser,
                    0,
                    quorumsOperatorSetParams,
                    quorumsMinimumStake,
                    quorumsStrategyParams
                )
            );
        }

        require(
            Ownable(address(registryCoordinator)).owner() != address(0),
            "Owner uninitialized"
        );

        {   // JSON output scope
            string memory parent_object = "parent object";
            string memory deployed_addresses = "addresses";
            vm.serializeAddress(
                deployed_addresses,
                "proxyAdmin",
                address(mockAvsProxyAdmin)
            );
            vm.serializeAddress(
                deployed_addresses,
                "mockAvsServiceManager",
                address(mockAvsServiceManager)
            );
            vm.serializeAddress(
                deployed_addresses,
                "mockAvsServiceManagerImplementation",
                address(mockAvsServiceManagerImplementation)
            );
            vm.serializeAddress(
                deployed_addresses,
                "registryCoordinator",
                address(registryCoordinator)
            );
            vm.serializeAddress(
                deployed_addresses,
                "registryCoordinatorImplementation",
                address(registryCoordinatorImplementation)
            );
            string memory deployed_addresses_output = vm.serializeAddress(
                deployed_addresses,
                "operatorStateRetriever",
                address(operatorStateRetriever)
            );

            string memory finalJson = vm.serializeString(
                parent_object,
                deployed_addresses,
                deployed_addresses_output
            );

            writeOutput(finalJson, "mockavs_deployment_output");
        }

        return
            MockAvsContracts(
                mockAvsServiceManager,
                registryCoordinator,
                operatorStateRetriever
            );
    }
    
    function _writeContractsToRegistry(
        ContractsRegistry contractsRegistry,
        EigenlayerContracts memory eigenlayerContracts,
        MockAvsContracts memory mockAvsContracts
    ) internal {
        contractsRegistry.registerContract(
            "mockAvsServiceManager",
            address(mockAvsContracts.mockAvsServiceManager)
        );
        contractsRegistry.registerContract(
            "mockAvsRegistryCoordinator",
            address(mockAvsContracts.registryCoordinator)
        );
        contractsRegistry.registerContract(
            "mockAvsOperatorStateRetriever",
            address(mockAvsContracts.operatorStateRetriever)
        );
        contractsRegistry.registerContract(
            "delegationManager",
            address(eigenlayerContracts.delegationManager)
        );
        contractsRegistry.registerContract(
            "strategyManager",
            address(eigenlayerContracts.strategyManager)
        );
        contractsRegistry.registerContract(
            "rewardsCoordinator",
            address(eigenlayerContracts.rewardsCoordinator)
        );
    }
}
