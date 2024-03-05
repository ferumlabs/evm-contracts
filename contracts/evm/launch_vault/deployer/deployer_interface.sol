// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.23;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

interface IDeployer {
  // Deploys the specified asset and amount. Should transfer rebasing tokens representing the deployed assets
  // to the caller of the function.
  function deploy(IERC20 asset, uint amount) external;
  // Reoves the specified asset and amount from deployment. Expects `amount` rebasing tokens to be transfered 
  // to the deployer contract before the call is made. Should transfer the removed assets back to the caller of
  // the function. 
  function remove(IERC20 asset, uint amount) external;
  // Returns the total amount deployed for the specified asset.
  function totalDeployedAmount(IERC20 asset) external view returns (uint);
  // Returns the address of the rebasing token representing the specified asset.
  function poolToken(IERC20 asset) external view returns (address);
}
