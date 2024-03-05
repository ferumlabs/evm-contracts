import "tsconfig-paths/register";
import { env, ETHERNAL_NODE, NODE, DEV } from '@/env';
if (env === ETHERNAL_NODE && !process.env.ETHERNAL_PASSWORD) {
  throw new Error("Missing ETHERNAL_PASSWORD. Set it as an environment variable.");
}
import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import "@openzeppelin/hardhat-upgrades";
import "@symblox/hardhat-abi-gen";
import "hardhat-deploy";
import "hardhat-ethernal";

const namedAccounts = ((): HardhatUserConfig['namedAccounts'] => {
  switch (env) {
    case NODE:
    case ETHERNAL_NODE:
      return {
        deployer: 0,
        localVaultDeployer: {
          hardhat: "0x33791bc5ee582D483AEA73E84c2aEAf3b7D5577f",
        },
      };
    case DEV:
      return {
        deployer: 0,
      };
    default:
      throw new Error("Unsupported environment: " + env);
  }
})();

const config: HardhatUserConfig = {
  solidity: "0.8.23",
  networks: {
    hardhat: {
      chainId: 1337,
      initialBaseFeePerGas: 0,
    },
  },
  paths: {
    artifacts: "./hardhat_artifacts",
    cache: "./hardhat_cache",
    tests: "./contracts/evm/tests",
    deploy: './contracts/evm/deploy',
    deployments: './hardhat_deployments',
  },
  typechain: {
    outDir: './hardhat_artifacts/types',
    target: "ethers-v6",
  },
  abiExporter: {
    path: './hardhat_abi/',
  },
  namedAccounts,
  ethernal: {
    disableSync: false,
    disableTrace: false,
    uploadAst: true,
    disabled: env !== ETHERNAL_NODE,
    skipFirstBlock: false,
    verbose: false,
  },
};

export default config;
