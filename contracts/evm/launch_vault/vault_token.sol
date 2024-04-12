// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.23;

import {AccessControlUpgradeable} from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";

// LP token representing share of the assets in the vault. Should be non transferrable unless configured
// to be via `transferable`.

contract BlackwingVaultToken is ERC20Upgradeable, AccessControlUpgradeable {
  string public constant TRANSFERS_DISABLED_ERR = '1'; // Transfers are disabled
  string public constant UNAUTHORIZED_ERR = '2'; // Not authorized to perform function

  bytes32 public constant VAULT_ROLE = keccak256("VAULT_ROLE");

  address private vault;
  bool private transferable;

  /// @custom:oz-upgrades-unsafe-allow constructor
  constructor() {
    _disableInitializers();
  }

  function initialize(address _vault, string memory name, string memory symbol) public initializer {
    ERC20Upgradeable.__ERC20_init(name, symbol);
    AccessControlUpgradeable.__AccessControl_init();
    AccessControlUpgradeable._grantRole(VAULT_ROLE, _vault);
    vault = _vault;
    transferable = false;
  }

  function getVaultAddress() public view returns (address) {
    return vault;
  }

  function transfer(address recipient, uint256 amount) public override returns (bool) {
    require(transferable, TRANSFERS_DISABLED_ERR);
    return super.transfer(recipient, amount);
  }

  function transferFrom(address sender, address recipient, uint256 amount) public override returns (bool) {
    require(transferable, TRANSFERS_DISABLED_ERR);
    return super.transferFrom(sender, recipient, amount);
  }

  function mint(address to, uint256 amount) external {
    require(hasRole(VAULT_ROLE, msg.sender), UNAUTHORIZED_ERR);
    _mint(to, amount);
  }

  function burn(address from, uint256 amount) external {
    require(hasRole(VAULT_ROLE, msg.sender), UNAUTHORIZED_ERR);
    _burn(from, amount);
  }

  function enableTransfers() public {
    require(hasRole(VAULT_ROLE, msg.sender), UNAUTHORIZED_ERR);
    transferable = true;
  }

  function disableTransfers() public {
    require(hasRole(VAULT_ROLE, msg.sender), UNAUTHORIZED_ERR);
    transferable = false;
  }
}
