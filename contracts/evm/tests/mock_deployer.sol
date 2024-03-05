// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.23;

import { MockCoin } from "./mock_coin.sol";
import { IDeployer } from "../launch_vault/deployer/deployer_interface.sol";
import { IERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

struct PoolInfo {
  bool isValue;
  MockCoin vaultCoin;
  MockCoin asset;
}

contract MockDeployer is IDeployer {
  mapping(IERC20 => PoolInfo) pools;

  event Deployed(IERC20 asset, uint amount);

  event Removed(IERC20 asset, uint amount);

  function registerPool(MockCoin asset, MockCoin vaultCoin) external {
    require(!pools[asset].isValue);
    pools[asset].isValue = true;
    pools[asset].vaultCoin = vaultCoin;
    pools[asset].asset = asset;
  }

  function deploy(IERC20 asset, uint amount) external {
    require(pools[asset].isValue);
    pools[asset].asset.burn(address(this), amount);
    pools[asset].vaultCoin.mint(msg.sender, amount);
    require(pools[asset].vaultCoin.balanceOf(address(this)) == 0, "unexpected vault token amount");
    require(pools[asset].asset.balanceOf(address(this)) == 0, "unexpected asset amount");
    emit Deployed(asset, amount);
  }

  function remove(IERC20 asset, uint amount) external {
    require(pools[asset].isValue);
    pools[asset].vaultCoin.burn(address(this), amount);
    pools[asset].asset.mint(msg.sender, amount);
    require(pools[asset].vaultCoin.balanceOf(address(this)) == 0, "unexpected vault token amount");
    require(pools[asset].asset.balanceOf(address(this)) == 0, "unexpected asset amount");
    emit Removed(asset, amount);
  }

  function totalDeployedAmount(IERC20 asset) external view returns (uint) {
    require(pools[asset].isValue);
    return pools[asset].vaultCoin.balanceOf(msg.sender);
  }

  function poolToken(IERC20 asset) external view returns (address) {
    return address(pools[asset].vaultCoin);
  }

  function mockPositiveYield(MockCoin asset, address vault, uint amount) external {
    require(pools[asset].isValue);
    pools[asset].vaultCoin.mint(vault, amount);
  }

  function mockNegativeYield(MockCoin asset, address vault, uint amount) external {
    require(pools[asset].isValue);
    pools[asset].vaultCoin.burn(vault, amount);
  }
}
