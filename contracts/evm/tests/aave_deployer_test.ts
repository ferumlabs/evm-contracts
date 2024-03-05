import chai, { expect } from "chai";
chai.use(require('chai-as-promised'));

import {
  deployMockCoin,
  deployBlackwingVaultContract,
  deployAaveDeployer,
  deployAave,
} from "../utils";
import { ethers } from "hardhat";

describe("Aave deployer", function () {
  it("Only owner allowed to register vault", async function () {
    const [_mockAsset, _aToken, _pool, registry, _provider] = await deployAave();
    const deployer = await deployAaveDeployer(registry);
    const vault = await deployBlackwingVaultContract();

    const [owner, addr] = await ethers.getSigners();

    await expect(
      deployer
        .connect(addr)
        .registerVault(await vault.getAddress()))
      .to.be
      .rejectedWith(await deployer.UNAUTHORIZED_ERR());
    await deployer
      .connect(owner)
      .registerVault(await vault.getAddress());
  });

  it("Only owner allowed to whitelist asset", async function () {
    const [mockAsset, _aToken, _pool, registry, provider] = await deployAave();
    const deployer = await deployAaveDeployer(registry);
    const vault = await deployBlackwingVaultContract();
    await deployer.registerVault(await vault.getAddress());

    const [_owner, addr] = await ethers.getSigners();

    await expect(
      deployer.connect(addr).whitelistAsset(await mockAsset.getAddress(), provider)
    ).to.be.rejectedWith(await deployer.UNAUTHORIZED_ERR());
    await deployer.whitelistAsset(await mockAsset.getAddress(), provider);
  });

  it("Invalid asset provider pair", async function () {
    const [_mockAsset, _aToken, _pool, registry, provider] = await deployAave();
    const otherMockAsset = await deployMockCoin();
    const deployer = await deployAaveDeployer(registry);
    const vault = await deployBlackwingVaultContract();
    await deployer.registerVault(await vault.getAddress());

    const [_owner, _addr] = await ethers.getSigners();

    await expect(deployer.whitelistAsset(await otherMockAsset.getAddress(), provider)).to.be.
    rejectedWith(await deployer.INVALID_ATOKEN_ERR());
  });

  it("Pool token", async function () {
    const [mockAsset, aToken, _pool, registry, provider] = await deployAave();
    const otherMockAsset = await deployMockCoin();
    const deployer = await deployAaveDeployer(registry);
    const vault = await deployBlackwingVaultContract();
    await deployer.registerVault(await vault.getAddress());
    await deployer.whitelistAsset(await mockAsset.getAddress(), provider);
    expect(await deployer.poolToken(mockAsset)).to.equal(await aToken.getAddress());
    await expect(deployer.poolToken(otherMockAsset)).to.be.rejectedWith(await deployer.UNREGISTERED_ASSET_ERR());
  });

  it("Deploy / withdraw unregistered asset", async function () {
    const [mockAsset, _aToken, _pool, registry, provider] = await deployAave();
    const otherMockAsset = await deployMockCoin();
    const deployer = await deployAaveDeployer(registry);

    const [owner, vault] = await ethers.getSigners();
    await otherMockAsset.mint(owner, 10_000);

    await deployer.registerVault(await vault.getAddress());
    await deployer.whitelistAsset(await mockAsset.getAddress(), provider);

    await expect(deployer.connect(vault).deploy(await otherMockAsset.getAddress(), 1_000)).to.be.
    rejectedWith(await deployer.UNREGISTERED_ASSET_ERR());
    await expect(deployer.connect(vault).remove(await otherMockAsset.getAddress(), 1_000)).to.be.
    rejectedWith(await deployer.UNREGISTERED_ASSET_ERR());
  });

  it("Deploy and withdraw", async function () {
    const [mockAsset, aToken, _pool, registry, provider] = await deployAave();
    const mockAssetAddress = await mockAsset.getAddress();
    const deployer = await deployAaveDeployer(registry);
    const deployerAddress = await deployer.getAddress();
    const [_owner, vault] = await ethers.getSigners();
    await deployer.registerVault(await vault.getAddress()); // 2nd signer is the vault.
    await deployer.whitelistAsset(mockAssetAddress, provider);

    await mockAsset.mint(deployerAddress, 10_000);

    // Only allowed to deploy from registered vault address;
    expect(deployer.deploy(mockAssetAddress, 1_000)).to.be.rejectedWith(await deployer.UNAUTHORIZED_ERR());
    await deployer.connect(vault).deploy(mockAssetAddress, 1_000);
    expect(await deployer.connect(vault).totalDeployedAmount(mockAssetAddress)).to.equal(1_000);

    // Deploy should transfer the aTokens to the caller of the deploy method.
    expect(await aToken.balanceOf(vault)).to.equal(1_000);
    expect(await aToken.totalSupply()).to.equal(1_000);

    // Remove requires caller to first transfer the aTokens to the deployer.
    await expect(deployer.connect(vault).remove(mockAssetAddress, 500)).to.be.
    rejectedWith(await deployer.NOT_ENOUGH_ATOKENS_TO_BURN());
    await aToken.connect(vault).transfer(deployerAddress, 500);
    expect(await aToken.balanceOf(vault)).to.equal(500);
    await deployer.connect(vault).remove(mockAssetAddress, 500);
    expect(await aToken.balanceOf(vault)).to.equal(500);
    expect(await aToken.totalSupply()).to.equal(500);

    // Only the vault is allowed to remove.
    await expect(deployer.remove(mockAssetAddress, 200)).to.be.rejectedWith(await deployer.UNAUTHORIZED_ERR());
  });
});
