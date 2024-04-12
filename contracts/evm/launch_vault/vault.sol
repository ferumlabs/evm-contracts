// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.23;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {AccessControlUpgradeable} from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import {IERC20} from "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {BlackwingVaultToken} from "./vault_token.sol";
import {IDeployer} from "./deployer/deployer_interface.sol";

struct PoolInfo {
  bool isValue;
  BlackwingVaultToken vaultToken;
  IDeployer deployer;
}

struct UserInfo {
  bool isValue;
  uint lastDepositBlock;
}

// The vault contract takes in user deposits and mints LP tokens for the user corresponding to their share
// of the assets in the vault. The assets the vault accepts is registered by the owner of the vault contract.
//
// Each vault asset is configured with an IDeployer which can deploy the assets to earn yield. Deployment and
// removal ("undeployment") of assets is triggered by the owner of the vault contract. The vault expects the
// IDeployer to mint rebasing tokens representing the assets deployed.

contract BlackwingVault is Initializable, AccessControlUpgradeable {
  using SafeERC20 for IERC20;

  event BalanceChange(bool isDeposit, address asset, address user, uint amount);

  string public constant VAULT_TOKEN_GRANULARITY_ERR = '1'; // Not enough assets provided to mint one vault token unit
  string public constant UNAUTHORIZED_ERR = '2'; // Not authorized to perform function
  string public constant UNREGISTERED_ASSET_ERR = '3'; // Asset not registered
  string public constant REGISTERED_ASSET_ERR = '4'; // Asset already registered
  string public constant ASSET_DEPLOYMENT_ERR = '5'; // Asset deployment failed
  string public constant ASSET_REMOVAL_ERR = '6'; // Deployed asset removal failed
  string public constant WITHDRAWS_DISABLED_ERR = '7'; // Withdraws are disabled
  string public constant MIN_BLOCKS_SINCE_LAST_DEPOSIT_ERR = '8'; // Min number of blocks since last deposit not reached for withdrawal

  uint public constant INITIAL_LIQUIDITY_MULTIPLIER = 100;
  bytes32 public constant OWNER_ROLE = keccak256("OWNER_ROLE");

  mapping(IERC20 => PoolInfo) private pools;
  mapping(address => UserInfo) private users;
  bool private withdrawsEnabled;
  uint private minBlocksSinceLastDeposit;

  /// @custom:oz-upgrades-unsafe-allow constructor
  constructor() {
    _disableInitializers();
  }

  function initialize(uint _minBlocksSinceLastDeposit) public initializer {
    AccessControlUpgradeable.__AccessControl_init();
    AccessControlUpgradeable._grantRole(OWNER_ROLE, msg.sender);
    withdrawsEnabled = false;
    minBlocksSinceLastDeposit = _minBlocksSinceLastDeposit;
  }

  function registerAsset(IERC20 asset, BlackwingVaultToken vaultToken, IDeployer deployer) public {
    require(hasRole(OWNER_ROLE, msg.sender), UNAUTHORIZED_ERR);
    require(!pools[asset].isValue, REGISTERED_ASSET_ERR);
    pools[asset] = PoolInfo({
      isValue: true,
      vaultToken: vaultToken,
      deployer: deployer
    });
  }

  function updateDeployer(IERC20 asset, IDeployer deployer) public {
    require(hasRole(OWNER_ROLE, msg.sender), UNAUTHORIZED_ERR);
    requireAssetRegistered(asset);
    pools[asset].deployer = deployer;
  }

  function enableWithdrawals() public {
    require(hasRole(OWNER_ROLE, msg.sender), UNAUTHORIZED_ERR);
    withdrawsEnabled = true;
  }

  function disableWithdrawals() public {
    require(hasRole(OWNER_ROLE, msg.sender), UNAUTHORIZED_ERR);
    withdrawsEnabled = false;
  }

  function setMinBlocksSinceLastDeposit(uint _minBlocksSinceLastDeposit) public {
    require(hasRole(OWNER_ROLE, msg.sender), UNAUTHORIZED_ERR);
    minBlocksSinceLastDeposit = _minBlocksSinceLastDeposit;
  }

  function deposit(IERC20 asset, uint amount) public {
    requireAssetRegistered(asset);

    PoolInfo memory pool = pools[asset];
    uint totalVaultTokenSupply = pool.vaultToken.totalSupply();
    uint vaultTokensToMint = 0;
    if (totalVaultTokenSupply == 0) {
      vaultTokensToMint = amount * INITIAL_LIQUIDITY_MULTIPLIER;
    } else {
      uint totalAssetBalance = getTotalAssetAmount(asset);
      // - totalAssetBalance can get large emough because of yield that it prevents users from depositing. 
      //   We mitigate this via INITIAL_LIQUIDITY_MULTIPLIER.
      // - totalAssetBalance could drop to 0 because of negative yield, in which case this will throw.
      // - totalVaultTokenSupply could increase to a large enough value that it overflows, preventing users 
      //   from depositing. This is fine as it can be resolved by withdrawing assets.
      vaultTokensToMint = amount * totalVaultTokenSupply / totalAssetBalance;
    }
    require(vaultTokensToMint > 0, VAULT_TOKEN_GRANULARITY_ERR);
    asset.safeTransferFrom(msg.sender, address(this), amount);
    pool.vaultToken.mint(msg.sender, vaultTokensToMint);
    if (!users[msg.sender].isValue) {
      users[msg.sender] = UserInfo({isValue: true, lastDepositBlock: block.number});
    } else {
      users[msg.sender].lastDepositBlock = block.number;
    }

    emit BalanceChange(true, address(asset), msg.sender, amount);
  }

  function withdraw(IERC20 asset, uint vaultTokensToBurn) public {
    requireAssetRegistered(asset);
    require(withdrawsEnabled, WITHDRAWS_DISABLED_ERR);

    UserInfo memory user = users[msg.sender];
    require(block.number > user.lastDepositBlock + minBlocksSinceLastDeposit, MIN_BLOCKS_SINCE_LAST_DEPOSIT_ERR);

    PoolInfo memory pool = pools[asset];
    uint undeployedAmount = getUndeployedAmount(asset);

    uint totalAssetBalance = getTotalAssetAmount(asset);
    uint totalVaultTokenSupply = pool.vaultToken.totalSupply();
    // - totalAssetBalance can get large enough (because of yield) that it causes overflows. This is pretty
    //   unlikely to happen though - uint256 is a gigantic number.
    uint amountReturned = vaultTokensToBurn * totalAssetBalance / totalVaultTokenSupply;
    require(amountReturned > 0, VAULT_TOKEN_GRANULARITY_ERR);
    if (undeployedAmount < amountReturned) {
      address poolToken = pool.deployer.poolToken(asset);
      uint diff = amountReturned - undeployedAmount;
      if (poolToken != address(0)) {
        IERC20(poolToken).safeTransfer(address(pool.deployer), diff);
      }
      pool.deployer.remove(asset, diff);
    }
    pool.vaultToken.burn(msg.sender, vaultTokensToBurn);
    asset.safeTransfer(msg.sender, amountReturned);

    emit BalanceChange(false, address(asset), msg.sender, amountReturned);
  }

  function vaultTokenAddress(IERC20 asset) public view returns (address) {
    requireAssetRegistered(asset);
    return address(pools[asset].vaultToken);
  }

  function balance(IERC20 asset, address user) public view returns (uint) {
    requireAssetRegistered(asset);

    PoolInfo memory pool = pools[asset];
    uint vaultTokenBalance = pool.vaultToken.balanceOf(user);
    uint vaultTokenSupply = pool.vaultToken.totalSupply();
    if (vaultTokenSupply == 0) {
      return 0;
    }
    uint totalAssetBalance = getTotalAssetAmount(asset);
    return vaultTokenBalance * totalAssetBalance / vaultTokenSupply;
  }

  function deployAssets(IERC20 asset, uint amount) public {
    require(hasRole(OWNER_ROLE, msg.sender), UNAUTHORIZED_ERR);
    requireAssetRegistered(asset);

    PoolInfo memory pool = pools[asset];
    asset.safeTransfer(address(pool.deployer), amount);
    pool.deployer.deploy(asset, amount);
  }

  function removeAssets(IERC20 asset, uint amount) public {
    require(hasRole(OWNER_ROLE, msg.sender), UNAUTHORIZED_ERR);
    requireAssetRegistered(asset);

    IDeployer deployer = pools[asset].deployer;
    address poolToken = deployer.poolToken(asset);
    if (poolToken != address(0)) {
      IERC20(poolToken).safeTransfer(address(deployer), amount);
    }
    deployer.remove(asset, amount);
  }

  function requireAssetRegistered(IERC20 asset) internal view {
    require(pools[asset].isValue, UNREGISTERED_ASSET_ERR);
  }

  function enableLPTransfers(IERC20 asset) public {
    require(hasRole(OWNER_ROLE, msg.sender), UNAUTHORIZED_ERR);
    requireAssetRegistered(asset);
    pools[asset].vaultToken.enableTransfers();
  }

  function disableLPTransfers(IERC20 asset) public {
    require(hasRole(OWNER_ROLE, msg.sender), UNAUTHORIZED_ERR);
    requireAssetRegistered(asset);
    pools[asset].vaultToken.disableTransfers();
  }

  function getTotalAssetAmount(IERC20 asset) internal view returns (uint) {
    return getUndeployedAmount(asset) + getDeployedAmount(asset);
  }

  function getDeployedAmount(IERC20 asset) internal view returns (uint) {
    return pools[asset].deployer.totalDeployedAmount(asset);
  }

  function getUndeployedAmount(IERC20 asset) internal view returns (uint) {
    return asset.balanceOf(address(this));
  }

  function lastDepositBlock(address user) public view returns (uint) {
    return users[user].lastDepositBlock;
  }
}
