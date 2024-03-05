// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.23;

import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {IPool} from "@aave/core-v3/contracts/interfaces/IPool.sol";
import {IPoolAddressesProvider} from "@aave/core-v3/contracts/interfaces/IPoolAddressesProvider.sol";
import {IPoolAddressesProviderRegistry} from "@aave/core-v3/contracts/interfaces/IPoolAddressesProviderRegistry.sol";
import {IAToken} from "@aave/core-v3/contracts/interfaces/IAToken.sol";
import {IDeployer} from "./deployer_interface.sol";

struct PoolInfo {
  IAToken aToken;
  IPoolAddressesProvider pap;
  bool isValue;
}

// Implements the IDeployer interface to deploy assets by supplying them to Aave lending pools.

contract BlackwingAaveDeployer is AccessControl, IDeployer {
  string public constant UNREGISTERED_ASSET_ERR = '1'; // Asset not registered
  string public constant UNAUTHORIZED_ERR = '2'; // Not authorized to perform function
  string public constant INVALID_ATOKEN_ERR = '3'; // Invalid Aave aToken.
  string public constant INVALID_POOL_ADDRESS_PROVIDER_ERR = '4'; // Invalid Aave pool address provider.
  string public constant NOT_ENOUGH_ATOKENS_TO_BURN = '5'; // Deployer not given enough aTokens to burn.

  bytes32 public constant OWNER_ROLE = keccak256("OWNER_ROLE");
  bytes32 public constant VAULT_ROLE = keccak256("VAULT_ROLE");

  IPoolAddressesProviderRegistry aaveRegistry;
  mapping(IERC20 => PoolInfo) pools;

  constructor(IPoolAddressesProviderRegistry _aaveRegistry) AccessControl() {
    AccessControl._grantRole(OWNER_ROLE, msg.sender);
    aaveRegistry = _aaveRegistry;
  }

  function registerVault(address vault) public {
    require(hasRole(OWNER_ROLE, msg.sender), UNAUTHORIZED_ERR);
    AccessControl._grantRole(VAULT_ROLE, vault);
  }

  function whitelistAsset(IERC20 asset, IPoolAddressesProvider pap) public {
    require(hasRole(OWNER_ROLE, msg.sender), UNAUTHORIZED_ERR);
    require(
      aaveRegistry.getAddressesProviderIdByAddress(address(pap)) > 0,
      INVALID_POOL_ADDRESS_PROVIDER_ERR
    );
    IAToken aToken = IAToken(IPool(pap.getPool()).getReserveData(address(asset)).aTokenAddress);
    require(aToken.UNDERLYING_ASSET_ADDRESS() == address(asset), INVALID_ATOKEN_ERR);
    require(asset.approve(pap.getPool(), type(uint256).max));
    pools[asset] = PoolInfo(aToken, pap, true);
  }

  function deploy(IERC20 asset, uint amount) public {
    require(hasRole(VAULT_ROLE, msg.sender), UNAUTHORIZED_ERR);
    requireAssetWhitelisted(asset);
    IPool pool = IPool(pools[asset].pap.getPool());
    pool.supply(address(asset), amount, msg.sender, 0);
  }

  function remove(IERC20 asset, uint amount) public {
    require(hasRole(VAULT_ROLE, msg.sender), UNAUTHORIZED_ERR);
    requireAssetWhitelisted(asset);
    PoolInfo memory info = pools[asset];
    require(info.aToken.balanceOf(address(this)) >= amount, NOT_ENOUGH_ATOKENS_TO_BURN);
    IPool pool = IPool(info.pap.getPool());
    pool.withdraw(address(asset), amount, msg.sender);
  }

  function poolToken(IERC20 asset) public view override returns (address) {
    requireAssetWhitelisted(asset);
    return address(pools[asset].aToken);
  }

  function totalDeployedAmount(IERC20 asset) public view returns (uint) {
    requireAssetWhitelisted(asset);
    return pools[asset].aToken.balanceOf(msg.sender);
  }

  function requireAssetWhitelisted(IERC20 asset) internal view {
    require(pools[asset].isValue, UNREGISTERED_ASSET_ERR);
  }
}
