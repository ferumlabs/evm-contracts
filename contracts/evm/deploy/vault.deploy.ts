import { env, NODE } from '@/env';
import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { DeployFunction, TxOptions } from 'hardhat-deploy/types';
import {
  VaultContractName,
  VaultTokenContractName,
  MockCoinContractName,
} from '../constants';
import {
  deployMockDeployer
} from '../utils';
import { MockCoin } from '@/hardhat_artifacts/types';
import ProxyAdmin from 'hardhat-deploy/extendedArtifacts/ProxyAdmin.json';

const deploy: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  switch (env) {
    case NODE:
      return await nodeDeploy(hre);
    default:
      throw new Error("Unsupported environment: " + env);
  }
}

const nodeDeploy: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  if (env !== NODE) {
    throw new Error("Trying to deploy dev artifacts in non dev environment: " + env);
  }

  const { deploy } = hre.deployments;
  const { ethers } = hre;
  const { localVaultDeployer: localVaultDeployerName } = await hre.getNamedAccounts();
  await fundAccount(hre, localVaultDeployerName, "1000");
  const localVaultDeployer = await ethers.getSigner(localVaultDeployerName);

  console.log("Starting deploy...");

  const deployerContract = await deployMockDeployer();
  const deployerAddress = await deployerContract.getAddress();

  const wethContract = await deployMockCoin(hre, localVaultDeployerName, "WETH", "WETH", 18);
  const wethAddress = await wethContract.getAddress();
  console.log("Depolyed WETH to", wethAddress);
  const usdcContract = await deployMockCoin(hre, localVaultDeployerName, "USDC", "USDC", 6);
  const usdcAddress = await usdcContract.getAddress();
  console.log("Depolyed USDC to", usdcAddress);

  const vaultDeployment = await deploy(VaultContractName, {
    from: localVaultDeployerName,
    log: true,
    proxy: {
      proxyContract: "OpenZeppelinTransparentProxy",
      owner: localVaultDeployerName,
      viaAdminContract: {
        name: "LaunchVaultProxyAdmin",
        artifact: ProxyAdmin,
      },
      execute: {
        init: {
          methodName: "initialize",
          args: [0],
        },
      },
    },
  });
  const vaultContract = await ethers.getContractAt(VaultContractName, vaultDeployment.address);
  console.log("Deployed vault to", await vaultContract.getAddress());

  const vaultTokenWETHDeployment = await deploy(VaultTokenContractName, {
    from: localVaultDeployerName,
    log: true,
    proxy: {
      proxyContract: "OpenZeppelinTransparentProxy",
      owner: localVaultDeployerName,
      viaAdminContract: {
        name: "LaunchVaultProxyAdmin",
        artifact: ProxyAdmin,
      },
      execute: {
        init: {
          methodName: "initialize",
          args: [vaultDeployment.address, "Blackwing Launch Vault Token - WETH", "blvtWETH"],
        },
      },
    },
  });
  const vaultTokenUSDCDeployment = await deploy(VaultTokenContractName, {
    from: localVaultDeployerName,
    log: true,
    proxy: {
      proxyContract: "OpenZeppelinTransparentProxy",
      viaAdminContract: {
        name: "LaunchVaultProxyAdmin",
        artifact: ProxyAdmin,
      },
      execute: {
        init: {
          methodName: "initialize",
          args: [vaultDeployment.address, "Blackwing Launch Vault Token - USDC", "blvtUSDC"],
        },
      },
    },
  });

  await vaultContract.
    connect(localVaultDeployer).
    registerAsset(wethAddress, vaultTokenWETHDeployment.address, deployerAddress);
  await vaultContract.
    connect(localVaultDeployer).
    registerAsset(usdcAddress, vaultTokenUSDCDeployment.address, deployerAddress);
  const usdcDeployerToken = await deployMockCoin(hre, localVaultDeployerName, "USDC Deployer Token", "deployerUSDC", 18);
  await deployerContract.registerPool(
    usdcAddress,
    await usdcDeployerToken.getAddress(),
  );
  const wethDeployerToken = await deployMockCoin(hre, localVaultDeployerName, "WETH Deployer Token", "deployerWETH", 18);
  await deployerContract.registerPool(
    wethAddress,
    await wethDeployerToken.getAddress(),
  );

  console.log("Deploy finished");
};

async function deployProxyAdmin(hre: HardhatRuntimeEnvironment, deployer: string) {
  const { deploy } = hre.deployments;
  const deployment = await deploy("DefaultProxyAdmin", {
    from: deployer,
    log: true,
  });
  return await hre.ethers.getContractAt("DefaultProxyAdmin", deployment.address);
}

async function deployMockCoin(
  hre: HardhatRuntimeEnvironment,
  from: TxOptions['from'],
  name: string,
  symbol: string,
  decimals: number
 ): Promise<MockCoin> {
  const { deploy } = hre.deployments;

  const deployment = await deploy(MockCoinContractName, {
    from,
    log: true,
    args: [name, symbol, decimals],
  });
  return await hre.ethers.getContractAt(MockCoinContractName, deployment.address);
}

async function fundAccount(hre: HardhatRuntimeEnvironment, address: string, amount: string) {
  const addr4 = (await hre.ethers.getSigners())[4];
  const txResponse = await addr4.sendTransaction({
    from: addr4,
    to: address,
    value: hre.ethers.parseUnits(
      amount,
      "ether"
    ),
  });
  await txResponse.wait();
}

export default deploy;
