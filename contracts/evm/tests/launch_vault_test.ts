import chai, { expect } from "chai";
chai.use(require('chai-as-promised'));

import {
  MockCoin,
  MockDeployer,
  BlackwingVault,
  BlackwingVaultToken,
} from "@/hardhat_artifacts/types";
import {
  deployMockCoin,
  deployMockDeployer,
  deployBlackwingVaultContract,
  deployBlackwingVaultTokenContract,
} from "../utils";
import { ethers } from "hardhat";
import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { mine } from "@nomicfoundation/hardhat-network-helpers";

async function setup(): Promise<[
  MockCoin, // Mock asset
  MockCoin, // Deployer LP token
  MockDeployer, // Deployer
  BlackwingVault, // Vault
  BlackwingVaultToken, // Vault token for mock
]> {
  const mockAsset = await deployMockCoin();
  const mockAssetAddress = await mockAsset.getAddress();
  const deployerLPToken = await deployMockCoin();
  const deployerLPTokenAddress = await deployerLPToken.getAddress();
  const deployer = await deployMockDeployer();
  const deployerAddress = await deployer.getAddress();
  await deployer.registerPool(mockAssetAddress, deployerLPTokenAddress);

  const vault = await deployBlackwingVaultContract();
  await vault.enableWithdrawals();
  const vaultAddress = await vault.getAddress();
  const vaultToken = await deployBlackwingVaultTokenContract(vaultAddress);
  const vaultTokenAddress = await vaultToken.getAddress();
  await vault.registerAsset(mockAssetAddress, vaultTokenAddress, deployerAddress);

  return [mockAsset, deployerLPToken, deployer, vault, vaultToken];
}

async function initialDeposit(
  account: HardhatEthersSigner,
  vault: BlackwingVault,
  vaultToken: BlackwingVaultToken,
  asset: MockCoin,
  amount: number,
): Promise<number> {
  const initMultiplier = parseInt((await vault.INITIAL_LIQUIDITY_MULTIPLIER()).toString());
  const initialBalance = await asset.balanceOf(account);
  const tx = await vault.connect(account).deposit(asset, amount);
  expect(tx).to.emit(vault, "BalanceChange").withArgs(true, account, asset, amount);
  const mockAssetBalance = await asset.balanceOf(account);
  expect(mockAssetBalance).to.equal(parseInt(initialBalance.toString()) - amount);
  const vaultTokenBalance = await vaultToken.balanceOf(account);
  expect(vaultTokenBalance).to.equal(amount * initMultiplier);
  const totalSupply = await vaultToken.totalSupply();
  expect(totalSupply).to.equal(amount * initMultiplier);
  const lastBlockDeposited = await vault.lastDepositBlock(account);
  expect(lastBlockDeposited).to.equal(tx.blockNumber);
  return amount * initMultiplier;
}

async function mockYield(
  vault: BlackwingVault,
  asset: MockCoin,
  amount: number,
) {
  if (amount > 0) {
    await asset.mint(await vault.getAddress(), amount);
  } else {
    await asset.burn(await vault.getAddress(), -amount);
  }
}

async function mockDeployedYield(
  vault: BlackwingVault,
  deployer: MockDeployer,
  asset: MockCoin,
  amount: number,
) {
  if (amount > 0) {
    await deployer.mockPositiveYield(asset, vault, amount);
  } else {
    await deployer.mockNegativeYield(asset, vault, -amount);
  }
}

describe("Launch vaults", function () {
  it("Vault token address", async function () {
    const [mockAsset, _deployerLPToken, _deployer, vault, vaultToken] = await setup();
    expect(await vault.vaultTokenAddress(mockAsset)).to.equal(await vaultToken.getAddress());
  });

  it("Only owner allowed to update deployer", async function () {
    const [mockAsset, _deployerLPToken, _deployer, vault, _vaultToken] = await setup();

    const newDeployerLPToken = await deployMockCoin();
    const newDeployer = await deployMockDeployer();
    await newDeployer.registerPool(await mockAsset.getAddress(), await newDeployerLPToken.getAddress());

    const [_owner, addr1] = await ethers.getSigners();

    await expect(vault.connect(addr1).updateDeployer(mockAsset, newDeployer)).to.be.rejectedWith(await vault.UNAUTHORIZED_ERR());

    await vault.updateDeployer(mockAsset, newDeployer)
  });

  it("Only owner allowed to register asset", async function () {
    const [_mockAsset, _deployerLPToken, deployer, vault, _vaultToken] = await setup();
    const newAsset = await deployMockCoin();

    const newVaultToken = await deployBlackwingVaultTokenContract(await vault.getAddress());

    const [_owner, addr1] = await ethers.getSigners();

    await expect(vault.connect(addr1).registerAsset(newAsset, newVaultToken, deployer)).to.be.rejectedWith(await vault.UNAUTHORIZED_ERR());

    await vault.registerAsset(newAsset, newVaultToken, deployer);
  });

  it("Only owner allowed to set min blocks since last deposit", async function () {
    const [_mockAsset, _deployerLPToken, deployer, vault, _vaultToken] = await setup();

    const [_owner, addr1] = await ethers.getSigners();

    await expect(vault.connect(addr1).setMinBlocksSinceLastDeposit(100)).to.be.rejectedWith(await vault.UNAUTHORIZED_ERR());

    await vault.setMinBlocksSinceLastDeposit(100);
  });

  it("Should fail when trying to register an already registered asset", async function () {
    const [_mockAsset, _deployerLPToken, deployer, vault, _vaultToken] = await setup();
    const newAsset = await deployMockCoin();

    const newVaultToken = await deployBlackwingVaultTokenContract(await vault.getAddress());

    const [_owner, addr1] = await ethers.getSigners();

    await vault.registerAsset(newAsset, newVaultToken, deployer);

    await expect(vault.registerAsset(newAsset, newVaultToken, deployer)).to.be.rejectedWith(await vault.REGISTERED_ASSET_ERR());
  });

  it("Should fail when trying to deposit using an unregistered asset", async function () {
    const [mockAsset, _deployerLPToken, _deployer, vault, vaultToken] = await setup();
    const vaultAddress = await vault.getAddress();
    const mockAsset2 = await deployMockCoin();
    const [owner] = await ethers.getSigners();

    await mockAsset.mint(owner, 10_000);
    await mockAsset.approve(vaultAddress, 10_000);
    await mockAsset2.mint(owner, 10_000);
    await mockAsset2.approve(vaultAddress, 10_000);

    await initialDeposit(owner, vault, vaultToken, mockAsset, 1_000);

    await expect(vault.deposit(mockAsset2, 1_000)).to.be.rejectedWith(await vault.UNREGISTERED_ASSET_ERR());
  });

  it("Balance with no deployed assets", async function () {
    const [mockAsset, _deployerLPToken, _deployer, vault, _vaultToken] = await setup();
    const vaultAddress = await vault.getAddress();
    const [owner] = await ethers.getSigners();

    await mockAsset.mint(owner, 10_000);
    await mockAsset.approve(vaultAddress, 10_000);

    expect(await vault.balance(mockAsset, owner)).to.equal(0);
  });

  it("Initial liquidity with only 1 token", async function () {
    const [mockAsset, _deployerLPToken, _deployer, vault, vaultToken] = await setup();
    const vaultAddress = await vault.getAddress();
    const [owner] = await ethers.getSigners();

    await mockAsset.mint(owner, 10_000);
    await mockAsset.approve(vaultAddress, 10_000);

    await initialDeposit(owner, vault, vaultToken, mockAsset, 1);
  });

  it("Should update last deposit block", async function () {
    const [mockAsset, _deployerLPToken, _deployer, vault, vaultToken] = await setup();
    const vaultAddress = await vault.getAddress();
    const [_owner, addr1] = await ethers.getSigners();

    await mockAsset.mint(addr1, 10_000);
    await mockAsset.connect(addr1).approve(vaultAddress, 10_000);
    await initialDeposit(addr1, vault, vaultToken, mockAsset, 1_000);
    await mine(1000)

    const initialLastBlockDeposited = await vault.lastDepositBlock(addr1);

    const tx = await vault.connect(addr1).deposit(mockAsset, 2_000);
    const lastBlockDeposited = await vault.lastDepositBlock(addr1);
    expect(lastBlockDeposited).to.equal(tx.blockNumber);
    expect(lastBlockDeposited - initialLastBlockDeposited).to.be.greaterThanOrEqual(1000)
  });

  it("Should emit deposit withdrawal events", async function () {
    const [mockAsset, _deployerLPToken, _deployer, vault, vaultToken] = await setup();
    const vaultAddress = await vault.getAddress();

    const [_owner, addr1, addr2] = await ethers.getSigners();

    await mockAsset.mint(addr1, 10_000);
    await mockAsset.mint(addr2, 10_000);
    await mockAsset.connect(addr1).approve(vaultAddress, 10_000);
    await mockAsset.connect(addr2).approve(vaultAddress, 10_000);

    const totalInitialSupply = await initialDeposit(addr1, vault, vaultToken, mockAsset, 1_000);

    await mockYield(vault, mockAsset, 1_000);

    const depositTx = await vault.connect(addr2).deposit(mockAsset, 200);
    expect(depositTx).to.emit(vault, "BalanceChange").withArgs(true, addr2, mockAsset, 200);

    expect(await vaultToken.totalSupply()).to.equal(10_000 + totalInitialSupply);
    expect(await vaultToken.balanceOf(addr2)).to.equal(10_000);
    expect(await mockAsset.balanceOf(addr2)).to.equal(9_800);
    expect(await vault.balance(mockAsset, addr2)).to.equal(200);
    expect(await mockAsset.balanceOf(vault)).to.equal(2_200);

    await mockYield(vault, mockAsset, 1_000);

    expect(await vault.balance(mockAsset, addr2)).to.equal(290);
    expect(await vaultToken.totalSupply()).to.equal(10_000 + totalInitialSupply);
    expect(await vaultToken.balanceOf(addr2)).to.equal(10_000);
    expect(await mockAsset.balanceOf(vault)).to.equal(3_200);

    await expect(vault.connect(addr2).withdraw(mockAsset, 10_001)).to.be.rejected;
    const withdrawTx = await vault.connect(addr2).withdraw(mockAsset, 10_000);
    expect(withdrawTx).to.emit(vault, "BalanceChange").withArgs(false, addr2, mockAsset, 90);
    expect(await mockAsset.balanceOf(addr2)).to.equal(10_090);
    expect(await vaultToken.totalSupply()).to.equal(totalInitialSupply);
    expect(await mockAsset.balanceOf(vault)).to.equal(2_910);
  });

  it("Should fail a withdrawal when withdrawals are disabled", async function () {
    const [mockAsset, _deployerLPToken, _deployer, vault, vaultToken] = await setup();
    await vault.disableWithdrawals();
    const vaultAddress = await vault.getAddress();

    const [_owner, addr1] = await ethers.getSigners();

    await mockAsset.mint(addr1, 10_000);
    await mockAsset.connect(addr1).approve(vaultAddress, 10_000);

    const totalInitialSupply = await initialDeposit(addr1, vault, vaultToken, mockAsset, 1_000);
    await expect(vault.connect(addr1).withdraw(mockAsset, totalInitialSupply)).to.be.rejectedWith(await vault.WITHDRAWS_DISABLED_ERR());
    await vault.enableWithdrawals();
    await vault.connect(addr1).withdraw(mockAsset, totalInitialSupply);
  });

  it("Should fail a withdrawal when not enough blocks have been passed", async function () {
    const [mockAsset, _deployerLPToken, _deployer, vault, vaultToken] = await setup();
    const vaultAddress = await vault.getAddress();

    const [_owner, addr1] = await ethers.getSigners();

    await vault.setMinBlocksSinceLastDeposit(5_000)
    await mockAsset.mint(addr1, 10_000);
    await mockAsset.connect(addr1).approve(vaultAddress, 10_000);

    const totalInitialSupply = await initialDeposit(addr1, vault, vaultToken, mockAsset, 1_000);
    await expect(vault.connect(addr1).withdraw(mockAsset, totalInitialSupply)).to.be.rejectedWith(await vault.MIN_BLOCKS_SINCE_LAST_DEPOSIT_ERR());
    await mine(5_000)

    await vault.connect(addr1).withdraw(mockAsset, totalInitialSupply);
  });

  it("Should mint/burn the correct amount of LP tokens when depositing/withdrawing increasing yield", async function () {
    const [mockAsset, _deployerLPToken, _deployer, vault, vaultToken] = await setup();
    const vaultAddress = await vault.getAddress();

    const [_owner, addr1, addr2] = await ethers.getSigners();

    await mockAsset.mint(addr1, 10_000);
    await mockAsset.mint(addr2, 10_000);
    await mockAsset.connect(addr1).approve(vaultAddress, 10_000);
    await mockAsset.connect(addr2).approve(vaultAddress, 10_000);

    const totalInitialSupply = await initialDeposit(addr1, vault, vaultToken, mockAsset, 1_000);

    await mockYield(vault, mockAsset, 1_000);

    await vault.connect(addr2).deposit(mockAsset, 200);
    expect(await vaultToken.totalSupply()).to.equal(10_000 + totalInitialSupply);
    expect(await vaultToken.balanceOf(addr2)).to.equal(10_000);
    expect(await mockAsset.balanceOf(addr2)).to.equal(9_800);
    expect(await vault.balance(mockAsset, addr2)).to.equal(200);
    expect(await mockAsset.balanceOf(vault)).to.equal(2_200);

    await mockYield(vault, mockAsset, 1_000);

    expect(await vault.balance(mockAsset, addr2)).to.equal(290);
    expect(await vaultToken.totalSupply()).to.equal(10_000 + totalInitialSupply);
    expect(await vaultToken.balanceOf(addr2)).to.equal(10_000);
    expect(await mockAsset.balanceOf(vault)).to.equal(3_200);

    await expect(vault.connect(addr2).withdraw(mockAsset, 10_001)).to.be.rejected;
    await vault.connect(addr2).withdraw(mockAsset, 10_000);
    expect(await mockAsset.balanceOf(addr2)).to.equal(10_090);
    expect(await vaultToken.totalSupply()).to.equal(totalInitialSupply);
    expect(await mockAsset.balanceOf(vault)).to.equal(2_910);
  });

  it("Should mint/burn the correct amount of LP tokens when depositing/withdrawing decreasing yield", async function () {
    const [mockAsset, _deployerLPToken, _deployer, vault, vaultToken] = await setup();
    const vaultAddress = await vault.getAddress();

    const [_owner, addr1, addr2] = await ethers.getSigners();

    await mockAsset.mint(addr1, 10_000);
    await mockAsset.mint(addr2, 10_000);
    await mockAsset.connect(addr1).approve(vaultAddress, 10_000);
    await mockAsset.connect(addr2).approve(vaultAddress, 10_000);

    const totalInitialSupply = await initialDeposit(addr1, vault, vaultToken, mockAsset, 1_000);

    await mockYield(vault, mockAsset, -500);

    await vault.connect(addr2).deposit(mockAsset, 200);
    expect(await vaultToken.totalSupply()).to.equal(40_000 + totalInitialSupply);
    expect(await vaultToken.balanceOf(addr2)).to.equal(40_000);
    expect(await mockAsset.balanceOf(addr2)).to.equal(9_800);
    expect(await vault.balance(mockAsset, addr2)).to.equal(200);
    expect(await mockAsset.balanceOf(vault)).to.equal(700);

    await mockYield(vault, mockAsset, -250);

    expect(await vault.balance(mockAsset, addr2)).to.equal(128);
    expect(await vaultToken.totalSupply()).to.equal(40_000 + totalInitialSupply);
    expect(await vaultToken.balanceOf(addr2)).to.equal(40_000);
    expect(await mockAsset.balanceOf(vault)).to.equal(450);

    await expect(vault.connect(addr2).withdraw(mockAsset, 40_001)).to.be.rejected;
    await vault.connect(addr2).withdraw(mockAsset, 40_000)
    expect(await mockAsset.balanceOf(addr2)).to.equal(9_928);
    expect(await vaultToken.totalSupply()).to.equal(totalInitialSupply);
    expect(await mockAsset.balanceOf(vault)).to.equal(322);
  });

  it("Should mint/burn the correct amount of LP tokens when depositing/withdrawing increasing deployed yield", async function () {
    const [mockAsset, _deployerLPToken, deployer, vault, vaultToken] = await setup();
    const vaultAddress = await vault.getAddress();

    const [_owner, addr1, addr2] = await ethers.getSigners();

    await mockAsset.mint(addr1, 10_000);
    await mockAsset.mint(addr2, 10_000);
    await mockAsset.connect(addr1).approve(vaultAddress, 10_000);
    await mockAsset.connect(addr2).approve(vaultAddress, 10_000);

    const totalInitialSupply = await initialDeposit(addr1, vault, vaultToken, mockAsset, 1_000);

    await vault.deployAssets(mockAsset, 700);
    await mockDeployedYield(vault, deployer, mockAsset, 1_000);

    await vault.connect(addr2).deposit(mockAsset, 200);
    expect(await vaultToken.totalSupply()).to.equal(10_000 + totalInitialSupply);
    expect(await vaultToken.balanceOf(addr2)).to.equal(10_000);
    expect(await mockAsset.balanceOf(addr2)).to.equal(9_800);
    expect(await vault.balance(mockAsset, addr2)).to.equal(200);
    expect(await mockAsset.balanceOf(vault)).to.equal(500);

    await mockDeployedYield(vault, deployer, mockAsset, 1_000);

    expect(await vault.balance(mockAsset, addr2)).to.equal(290);
    expect(await vaultToken.totalSupply()).to.equal(10_000 + totalInitialSupply);
    expect(await vaultToken.balanceOf(addr2)).to.equal(10_000);
    expect(await mockAsset.balanceOf(vault)).to.equal(500);

    await expect(vault.connect(addr2).withdraw(mockAsset, 10_001)).to.be.rejected;
    await vault.connect(addr2).withdraw(mockAsset, 10_000);
    expect(await mockAsset.balanceOf(addr2)).to.equal(10_090);
    expect(await vaultToken.totalSupply()).to.equal(totalInitialSupply);
    expect(await mockAsset.balanceOf(vault)).to.equal(210);
  });

  it("Should mint/burn the correct amount of LP tokens when depositing/withdrawing decreasing deployed yield", async function () {
    const [mockAsset, _deployerLPToken, deployer, vault, vaultToken] = await setup();
    const vaultAddress = await vault.getAddress();

    const [_owner, addr1, addr2] = await ethers.getSigners();

    await mockAsset.mint(addr1, 10_000);
    await mockAsset.mint(addr2, 10_000);
    await mockAsset.connect(addr1).approve(vaultAddress, 10_000);
    await mockAsset.connect(addr2).approve(vaultAddress, 10_000);

    const totalInitialSupply = await initialDeposit(addr1, vault, vaultToken, mockAsset, 1_000);

    vault.deployAssets(mockAsset, 750);

    await mockDeployedYield(vault, deployer, mockAsset, -500);

    await vault.connect(addr2).deposit(mockAsset, 200);
    expect(await vaultToken.totalSupply()).to.equal(40_000 + totalInitialSupply);
    expect(await vaultToken.balanceOf(addr2)).to.equal(40_000);
    expect(await mockAsset.balanceOf(addr2)).to.equal(9_800);
    expect(await vault.balance(mockAsset, addr2)).to.equal(200);
    expect(await mockAsset.balanceOf(vault)).to.equal(450);

    await mockDeployedYield(vault, deployer, mockAsset, -250);

    expect(await vault.balance(mockAsset, addr2)).to.equal(128);
    expect(await vaultToken.totalSupply()).to.equal(40_000 + totalInitialSupply);
    expect(await vaultToken.balanceOf(addr2)).to.equal(40_000);
    expect(await mockAsset.balanceOf(vault)).to.equal(450);

    await expect(vault.connect(addr2).withdraw(mockAsset, 40_001)).to.be.rejected;
    await vault.connect(addr2).withdraw(mockAsset, 40_000);
    expect(await mockAsset.balanceOf(addr2)).to.equal(9_928);
    expect(await vaultToken.totalSupply()).to.equal(totalInitialSupply);
    expect(await mockAsset.balanceOf(vault)).to.equal(322);
  });

  it("Vault token granularity on deposit - excessive yield", async function () {
    const [mockAsset, _deployerLPToken, _deployer, vault, vaultToken] = await setup();
    const vaultAddress = await vault.getAddress();

    const [_owner, addr1, addr2] = await ethers.getSigners();

    await mockAsset.mint(addr1, 1_000);
    await mockAsset.mint(addr2, 1_000);
    await mockAsset.connect(addr1).approve(vaultAddress, 1_000);
    await mockAsset.connect(addr2).approve(vaultAddress, 1_000);

    const initialTotalSupply = await initialDeposit(addr1, vault, vaultToken, mockAsset, 1);

    await mockYield(vault, mockAsset, initialTotalSupply);

    // This call fails because the value of an LP token is too high for 1 unit of the underlying to equal
    // at least one unit of the LP token.
    await expect(vault.connect(addr2).deposit(mockAsset, 1)).to.be.rejectedWith(await vault.VAULT_TOKEN_GRANULARITY_ERR());
    // 10 units of the underlying is enough to mint at least 1 LP token.
    vault.connect(addr2).deposit(mockAsset, 10)
  });

  it("Vault token granularity on withdraw - negative yield", async function () {
    const [mockAsset, _deployerLPToken, _deployer, vault, vaultToken] = await setup();
    const vaultAddress = await vault.getAddress();

    const [_owner, addr1, addr2] = await ethers.getSigners();

    await mockAsset.mint(addr1, 1_000);
    await mockAsset.mint(addr2, 1_000);
    await mockAsset.connect(addr1).approve(vaultAddress, 1_000);
    await mockAsset.connect(addr2).approve(vaultAddress, 1_000);

    const initialTotalSupply = await initialDeposit(addr1, vault, vaultToken, mockAsset, 10);

    await mockYield(vault, mockAsset, -9);

    // This call fails because the value of an LP token is too low for 1 unit of the LP token to equal at least
    // 1 unit of the underlying asset.
    await expect(vault.connect(addr2).withdraw(mockAsset, 1)).to.be.rejectedWith(await vault.VAULT_TOKEN_GRANULARITY_ERR());
    // The entire balance of the LP token is enough to withdraw at least 1 unit of the underlying.
    vault.connect(addr2).withdraw(mockAsset, initialTotalSupply)
  });
});

describe("Launch vaults asset deployments", function () {
  it("Deployment of unregistered asset", async function () {
    const [mockAsset, _deployerLPToken, _deployer, vault, vaultToken] = await setup();
    const vaultAddress = await vault.getAddress();
    const mockAsset2 = await deployMockCoin();
    const [owner] = await ethers.getSigners();

    await mockAsset.mint(owner, 10_000);
    await mockAsset.approve(vaultAddress, 10_000);

    await initialDeposit(owner, vault, vaultToken, mockAsset, 1_000);

    await expect(vault.deployAssets(mockAsset2, 1_000)).to.be.rejectedWith(await vault.UNREGISTERED_ASSET_ERR());
  });

  it("Only owner allowed to deploy/remove", async function () {
    const [mockAsset, _deployerLPToken, _deployer, vault, vaultToken] = await setup();
    const vaultAddress = await vault.getAddress();
    const [owner, addr1] = await ethers.getSigners();

    await mockAsset.mint(owner, 10_000);
    await mockAsset.approve(vaultAddress, 10_000);

    await initialDeposit(owner, vault, vaultToken, mockAsset, 1_000);

    await expect(vault.connect(addr1).deployAssets(mockAsset, 500)).to.be.rejectedWith(await vault.UNAUTHORIZED_ERR());
    await vault.deployAssets(mockAsset, 1_000);
    await expect(vault.connect(addr1).removeAssets(mockAsset, 500)).to.be.rejectedWith(await vault.UNAUTHORIZED_ERR());
  });

  it("Deployment fails when trying to deploy too much", async function () {
    const [mockAsset, _deployerLPToken, _deployer, vault, vaultToken] = await setup();
    const vaultAddress = await vault.getAddress();
    const [owner] = await ethers.getSigners();

    await mockAsset.mint(owner, 10_000);
    await mockAsset.approve(vaultAddress, 10_000);

    await initialDeposit(owner, vault, vaultToken, mockAsset, 1_000);

    await expect(vault.deployAssets(mockAsset, 1_001)).to.be.rejected;
  });

  it("Removal fails when trying to remove too much", async function () {
    const [mockAsset, _deployerLPToken, _deployer, vault, vaultToken] = await setup();
    const vaultAddress = await vault.getAddress();
    const [owner] = await ethers.getSigners();

    await mockAsset.mint(owner, 10_000);
    await mockAsset.approve(vaultAddress, 10_000);

    await initialDeposit(owner, vault, vaultToken, mockAsset, 1_000);

    await expect(vault.removeAssets(mockAsset, 1_001)).to.be.rejected;
  });

  it("Deployment", async function () {
    // The mock deployer contract expects ERC20 tokens to be transferred to it when deploying and
    // its LP tokens to be transferred when removing.

    const [mockAsset, deployerLPToken, deployer, vault, vaultToken] = await setup();
    const vaultAddress = await vault.getAddress();

    expect(await deployer.poolToken(mockAsset)).to.equal(await deployerLPToken.getAddress())

    const [_owner, addr1, addr2] = await ethers.getSigners();

    await mockAsset.mint(addr1, 10_000);
    await mockAsset.mint(addr2, 10_000);
    await mockAsset.connect(addr1).approve(vaultAddress, 10_000);
    await mockAsset.connect(addr2).approve(vaultAddress, 10_000);

    await initialDeposit(addr1, vault, vaultToken, mockAsset, 1_000);

    expect(await deployerLPToken.totalSupply()).to.equal(0);
    expect(await vault.deployAssets(mockAsset, 700)).to.emit(deployer, "Deployed").withArgs(mockAsset, 700);
    expect(await mockAsset.balanceOf(vault)).to.equal(300);

    expect(await deployerLPToken.balanceOf(vault)).to.equal(700);
    await mockDeployedYield(vault, deployer, mockAsset, 1_000);
    expect(await vault.removeAssets(mockAsset, 800)).to.emit(deployer, "Removed").withArgs(mockAsset, 800);
    expect(await deployerLPToken.balanceOf(vault)).to.equal(900);
    expect(await mockAsset.balanceOf(vault)).to.equal(1_100);

    expect(await vault.deployAssets(mockAsset, 200)).to.emit(deployer, "Deployed").withArgs(mockAsset, 200);
    expect(await mockAsset.balanceOf(vault)).to.equal(900);
    expect(await vault.removeAssets(mockAsset, 1_100)).to.emit(deployer, "Removed").withArgs(mockAsset, 1_100);
    expect(await mockAsset.balanceOf(vault)).to.equal(2_000);
  });

  it("Withdraw when funds deployed", async function () {
    const [mockAsset, deployerLPToken, deployer, vault, vaultToken] = await setup();
    const vaultAddress = await vault.getAddress();

    expect(await deployer.poolToken(mockAsset)).to.equal(await deployerLPToken.getAddress())

    const [_owner, addr1, addr2] = await ethers.getSigners();

    await mockAsset.mint(addr1, 10_000);
    await mockAsset.mint(addr2, 10_000);
    await mockAsset.connect(addr1).approve(vaultAddress, 10_000);
    await mockAsset.connect(addr2).approve(vaultAddress, 10_000);

    await initialDeposit(addr1, vault, vaultToken, mockAsset, 1_000);

    expect(await deployerLPToken.totalSupply()).to.equal(0);
    expect(await vault.deployAssets(mockAsset, 700)).to.emit(deployer, "Deployed").withArgs(mockAsset, 700);
    expect(await deployerLPToken.balanceOf(vault)).to.equal(700);
    expect(await mockAsset.balanceOf(vault)).to.equal(300);

    expect(await vaultToken.balanceOf(addr1)).to.equal(100_000);
    expect(await vault.connect(addr1).withdraw(mockAsset, 50_000)).to.emit(deployer, "Removed").withArgs(mockAsset, 200);
    expect(await deployerLPToken.balanceOf(vault)).to.equal(500);
    expect(await mockAsset.balanceOf(vault)).to.equal(0);
  });
});

describe("Vault tokens", function () {
  it("Vault token vault address", async function () {
    const [mockAsset, _deployerLPToken, _deployer, vault, vaultToken] = await setup();
    const vaultAddress = await vault.getAddress();
    expect(await vaultToken.getVaultAddress()).to.equal(vaultAddress);
  });

  it("Vault token transfer only allowed to change via vault and owner", async function () {
    const [mockAsset, _deployerLPToken, _deployer, vault, vaultToken] = await setup();

    const [_owner, addr1] = await ethers.getSigners();

    await expect(vaultToken.enableTransfers()).to.be.rejectedWith(await vaultToken.UNAUTHORIZED_ERR());
    await expect(vault.connect(addr1).enableLPTransfers(mockAsset)).to.be.rejectedWith(await vault.UNAUTHORIZED_ERR());
    await vault.enableLPTransfers(mockAsset);
    await vault.disableLPTransfers(mockAsset);
    await vault.enableLPTransfers(mockAsset);
    await expect(vaultToken.disableTransfers()).to.be.rejectedWith(await vaultToken.UNAUTHORIZED_ERR());
    await expect(vault.connect(addr1).disableLPTransfers(mockAsset)).to.be.rejectedWith(await vault.UNAUTHORIZED_ERR());
  });

  it("Vault token transfer fails for unregistered asset", async function () {
    const [_mockAsset, _deployerLPToken, _deployer, vault, _vaultToken] = await setup();
    const mockAsset2 = await deployMockCoin();

    await expect(vault.disableLPTransfers(mockAsset2)).to.be.rejectedWith(await vault.UNREGISTERED_ASSET_ERR());
  });

  it("Vault token mint/burn only allowed to change via vault", async function () {
    const [mockAsset, _deployerLPToken, _deployer, vault, vaultToken] = await setup();
    const vaultAddress = await vault.getAddress();

    const [owner] = await ethers.getSigners();

    await expect(vaultToken.mint(owner, 100)).to.be.rejectedWith(await vaultToken.UNAUTHORIZED_ERR());

    await mockAsset.approve(vaultAddress, 1_000);
    await mockAsset.mint(owner, 1_000);
    await initialDeposit(owner, vault, vaultToken, mockAsset, 100);

    await expect(vaultToken.burn(owner, 10)).to.be.rejectedWith(await vaultToken.UNAUTHORIZED_ERR());
  });

  it("Vault token not transferrable unless enabled", async function () {
    const [mockAsset, _deployerLPToken, _deployer, vault, vaultToken] = await setup();
    const vaultAddress = await vault.getAddress();

    const [_owner, addr1, addr2] = await ethers.getSigners();

    await mockAsset.mint(addr1, 1_000);
    await mockAsset.mint(addr2, 1_000);
    await mockAsset.connect(addr1).approve(vaultAddress, 1_000);
    await mockAsset.connect(addr2).approve(vaultAddress, 1_000);

    await initialDeposit(addr1, vault, vaultToken, mockAsset, 100);

    await expect(vaultToken.connect(addr1).transfer(addr2, 100)).to.be.rejectedWith(await vaultToken.TRANSFERS_DISABLED_ERR());
    await expect(vaultToken.connect(addr1).transferFrom(addr1, addr2, 100)).to.be.rejectedWith(await vaultToken.TRANSFERS_DISABLED_ERR());

    await vault.enableLPTransfers(mockAsset);

    await vaultToken.connect(addr1).transfer(addr2, 150);
    vaultToken.connect(addr2).approve(addr1, 50);
    await vaultToken.connect(addr1).transferFrom(addr2, addr1, 50);
    expect(await vaultToken.balanceOf(addr2)).to.equal(100);

    await vault.disableLPTransfers(mockAsset);

    await expect(vaultToken.connect(addr1).transfer(addr2, 100)).to.be.rejectedWith(await vaultToken.TRANSFERS_DISABLED_ERR());
  });
});
