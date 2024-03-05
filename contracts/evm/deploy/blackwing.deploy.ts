import { env, ETHERNAL_NODE } from '@/env';
import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { DeployFunction } from 'hardhat-deploy/types';
import { BlackwingContractName } from '../constants';
import ProxyAdmin from 'hardhat-deploy/extendedArtifacts/ProxyAdmin.json';

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const { deploy } = hre.deployments;
  const { deployer } = await hre.getNamedAccounts();

  console.log("Starting deploy...");

  const blackwing = await deploy(BlackwingContractName, {
    contract: BlackwingContractName,
    from: deployer,
    log: true,
    proxy: {
      proxyContract: "OpenZeppelinTransparentProxy",
      owner: deployer,
      viaAdminContract: {
        name: "BlackwingProxyAdmin",
        artifact: ProxyAdmin,
      },
      execute: {
        init: {
          methodName: "initialize",
          args: [],
        },
      },
    },
  });

  if (env === ETHERNAL_NODE) {
    console.log("Pushing to ethernal");
    await hre.ethernal.push({
      name: BlackwingContractName,
      address: blackwing.address,
    });
  }

  console.log("Deploy finished");
};

export default func;
