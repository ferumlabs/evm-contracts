// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.23;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

interface IDeployer {
  function deploy(IERC20 asset, uint amount) external;
  function remove(IERC20 asset, uint amount) external;
  function totalDeployedAmount(IERC20 asset) external view returns (uint);
  function poolToken(IERC20 asset) external view returns (address);
}
