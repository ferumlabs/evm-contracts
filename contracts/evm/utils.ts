import { ethers, upgrades } from "hardhat";
import { contracts } from "../../hardhat_artifacts/types";
import { v4 as uuidv4 } from 'uuid';
import { PoolAddressesProviderRegistry } from '@ferum/aave-deploy-v3';
import { BaseContract } from "ethers";
import {
  VaultContractName,
  VaultTokenContractName,
  AaveDeployerContractName,
  MockContractName,
  MockCoinContractName,
  MockATokenContractName,
  MockDeplyerContractName,
  MockAavePool,
  MockAaveAddressProvider,
  MockAaveRegistry,
} from "@/contracts/evm/constants"

export function uuid(): Buffer {
  let id: string = uuidv4();
  return Buffer.from(id.replaceAll('-', ''), 'hex')
}

export async function deployBlackwingVaultContract(): Promise<contracts.evm.launchVault.vaultSol.BlackwingVault> {
  const factory = await ethers.getContractFactory(VaultContractName);
  const contract = await upgrades.deployProxy(factory, [0]);
  await contract.waitForDeployment();
  return await ethers.getContractAt(VaultContractName, await contract.getAddress());
}

export async function deployBlackwingVaultTokenContract(vaultAddress: string): Promise<contracts.evm.launchVault.vaultTokenSol.BlackwingVaultToken> {
  const factory = await ethers.getContractFactory(VaultTokenContractName);
  const contract = await upgrades.deployProxy(factory, [vaultAddress, "BlackwingLP", "BLP"]);
  await contract.waitForDeployment();
  return await ethers.getContractAt(VaultTokenContractName, await contract.getAddress());
}

export async function deployMockContract(): Promise<contracts.evm.tests.mockContractSol.MockContract> {
  const contract = await ethers.deployContract(MockContractName);
  await contract.waitForDeployment();
  return await ethers.getContractAt(MockContractName, await contract.getAddress());
}

export async function deployMockCoin(decimals: number = 18): Promise<contracts.evm.tests.mockCoinSol.MockCoin> {
  const contract = await ethers.deployContract(MockCoinContractName, ["MockCoin", "MCOIN", decimals]);
  await contract.waitForDeployment();
  return await ethers.getContractAt(MockCoinContractName, await contract.getAddress());
}

export async function deployMockDeployer(): Promise<contracts.evm.tests.mockDeployerSol.MockDeployer> {
  const contract = await ethers.deployContract(MockDeplyerContractName);
  await contract.waitForDeployment();
  return await ethers.getContractAt(MockDeplyerContractName, await contract.getAddress());
}

export async function deployAaveDeployer(registryAddress: PoolAddressesProviderRegistry): Promise<contracts.evm.launchVault.deployer.aaveDeployerSol.BlackwingAaveDeployer> {
  const contract = await ethers.deployContract(AaveDeployerContractName, [registryAddress]);
  await contract.waitForDeployment();
  return await ethers.getContractAt(AaveDeployerContractName, await contract.getAddress());
}

export async function deployAave(): Promise<[
  contracts.evm.tests.mockCoinSol.MockCoin,
  contracts.evm.tests.mockAtokenSol.MockAToken,
  contracts.evm.tests.mockAaveSol.MockAavePool,
  contracts.evm.tests.mockAaveSol.MockAaveRegistry,
  contracts.evm.tests.mockAaveSol.MockAaveAddressProvider,
]> {
  const mockToken = await deployMockCoin();
  const mockAToken = await ethers.deployContract(MockATokenContractName, [await mockToken.getAddress()]);
  await mockAToken.waitForDeployment();
  const mockPool = await ethers.deployContract(
    MockAavePool,
    [
      await mockAToken.getAddress(),
      await mockToken.getAddress(),
    ],
  );
  await mockPool.waitForDeployment();

  const mockProvider = await ethers.deployContract(MockAaveAddressProvider);
  await mockProvider.waitForDeployment();
  const providerAddress = await (mockProvider as BaseContract).getAddress();
  await mockProvider.setMarketId("test");
  await mockProvider.setPoolImpl(await mockPool.getAddress());

  const mockRegistryContract = await ethers.deployContract(MockAaveRegistry);
  await mockRegistryContract.waitForDeployment();
  await mockRegistryContract.registerAddressesProvider(providerAddress, 1);

  return [
    mockToken,
    mockAToken,
    mockPool,
    mockRegistryContract,
    mockProvider,
  ];
}
