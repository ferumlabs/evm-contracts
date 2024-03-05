// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.23;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract MockAToken is ERC20 {
  IERC20 underlying;

  constructor(IERC20 _underlying) ERC20("MockRebaseCoin", "MCOIN") {
    underlying = _underlying;
  }

  function mint(address to, uint256 amount) external {
    _mint(to, amount);
  }

  function burn(address from, uint256 amount) external {
    _burn(from, amount);
  }

  function UNDERLYING_ASSET_ADDRESS() external view returns (address) {
    return address(underlying);
  }

  function mint(
    address,
    address,
    uint256,
    uint256
  ) external pure returns (bool) {
    revert();
  }

  function burn(address, address, uint256, uint256) external pure {
    revert();
  }


  function mintToTreasury(uint256, uint256) external pure {
    revert();
  }

  function transferOnLiquidation(address, address, uint256) external pure {
    revert();
  }

  function transferUnderlyingTo(address, uint256) external pure {
    revert();
  }

  function handleRepayment(address, address, uint256) external pure {
    revert();
  }

  function permit(
    address,
    address,
    uint256,
    uint256,
    uint8,
    bytes32,
    bytes32
  ) external pure {
    revert();
  }

  function RESERVE_TREASURY_ADDRESS() external pure returns (address) {
    revert();
  }

  function DOMAIN_SEPARATOR() external pure returns (bytes32) {
    revert();
  }

  function nonces(address) external pure returns (uint256) {
    revert();
  }

  function rescueTokens(address, address, uint256) external pure {
    revert();
  }
}