// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.23;

import {IPool} from "@aave/core-v3/contracts/interfaces/IPool.sol";
import {IPoolAddressesProvider} from "@aave/core-v3/contracts/interfaces/IPoolAddressesProvider.sol";
import {IPoolAddressesProviderRegistry} from "@aave/core-v3/contracts/interfaces/IPoolAddressesProviderRegistry.sol";
import {DataTypes} from '@aave/core-v3/contracts/protocol/libraries/types/DataTypes.sol';
import {MockAToken} from './mock_atoken.sol';
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract MockAaveRegistry is IPoolAddressesProviderRegistry {
  mapping(address => uint) ids;

  function getAddressesProvidersList() external pure returns (address[] memory) {
    // Not implemented.
    revert();
  }

  function getAddressesProviderIdByAddress(address addressesProvider) external view returns (uint256) {
    return ids[addressesProvider];
  }

  function getAddressesProviderAddressById(uint256) external pure returns (address) {
    // Not implemented.
    revert();
  }

  function registerAddressesProvider(address provider, uint256 id) external {
    ids[provider] = id;
  }

  function unregisterAddressesProvider(address provider) external {
    ids[provider] = 0;
  }
}

contract MockAaveAddressProvider is IPoolAddressesProvider {
  string marketId;
  address pool;

  function getMarketId() external view returns (string memory) {
    return marketId;
  }

  function setMarketId(string calldata newMarketId) external {
    marketId = newMarketId;
  }

  function getAddress(bytes32) external pure returns (address) {
    revert();
  }

  function setAddressAsProxy(bytes32, address) external pure {
    revert();
  }

  function setAddress(bytes32, address) external pure {
    revert();
  }

  function getPool() external view returns (address) {
    return pool;
  }

  function setPoolImpl(address newPoolImpl) external {
    pool = newPoolImpl;
  }

  function getPoolConfigurator() external pure returns (address) {
    revert();
  }

  function setPoolConfiguratorImpl(address) external pure {
    revert();
  }

  function getPriceOracle() external pure returns (address) {
    revert();
  }

  function setPriceOracle(address) external pure {
    revert();
  }

  function getACLManager() external pure returns (address) {
    revert();
  }

  function setACLManager(address) external pure {
    revert();
  }

  function getACLAdmin() external pure returns (address) {
    revert();
  }

  function setACLAdmin(address) external pure {
    revert();
  }

  function getPriceOracleSentinel() external pure returns (address) {
    revert();
  }

  function setPriceOracleSentinel(address) external pure {
    revert();
  }

  function getPoolDataProvider() external pure returns (address) {
    revert();
  }

  function setPoolDataProvider(address) external pure {
    revert();
  }
}

contract MockAavePool is IPool {
  MockAToken mockAToken;
  IERC20 registeredToken;

  constructor(MockAToken _mockAToken, IERC20 _token) {
    mockAToken = _mockAToken;
    registeredToken = _token;
  }

  function supply(address token, uint256 amount, address onBehalfOf, uint16) external {
    require(token == address(registeredToken), "Invalid token");
    if (onBehalfOf == address(0)) {
      mockAToken.mint(msg.sender, amount);
    } else {
      mockAToken.mint(onBehalfOf, amount);
    }
    IERC20(token).transferFrom(msg.sender, address(this), amount);
  }

  function withdraw(address token, uint256 amount, address to) external returns (uint256) {
    require(token == address(registeredToken), "Invalid token");
    mockAToken.burn(msg.sender, amount);
    if (to == address(0)) {
      IERC20(token).transfer(msg.sender, amount);
    } else {
      IERC20(token).transfer(to, amount);
    }
    return amount;
  }

  function getReserveData(address) external view returns (DataTypes.ReserveData memory) {
    return DataTypes.ReserveData({
      aTokenAddress: address(mockAToken),
      configuration: DataTypes.ReserveConfigurationMap({
        data: 0
      }),
      liquidityIndex: 0,
      currentLiquidityRate: 0,
      variableBorrowIndex: 0,
      currentVariableBorrowRate: 0,
      currentStableBorrowRate: 0,
      lastUpdateTimestamp: 0,
      id: 0,
      stableDebtTokenAddress: address(0),
      variableDebtTokenAddress: address(0),
      interestRateStrategyAddress: address(0),
      accruedToTreasury: 0,
      unbacked: 0,
      isolationModeTotalDebt: 0
    });
  }

  function mintUnbacked(
    address,
    uint256,
    address,
    uint16
  ) external pure {
    revert();
  }

  function backUnbacked(address, uint256, uint256) external pure returns (uint256) {
    revert();
  }

  function borrow(
    address,
    uint256,
    uint256,
    uint16,
    address
  ) external pure {
    revert();
  }

  function supplyWithPermit(
    address,
    uint256,
    address,
    uint16,
    uint256,
    uint8,
    bytes32,
    bytes32
  ) external pure {
    revert();
  }

  function repay(
    address,
    uint256,
    uint256,
    address
  ) external pure returns (uint256) {
    revert();
  }

  function repayWithPermit(
    address,
    uint256,
    uint256,
    address,
    uint256,
    uint8,
    bytes32,
    bytes32
  ) external pure returns (uint256) {
    revert();
  }

  function repayWithATokens(
    address,
    uint256,
    uint256
  ) external pure returns (uint256) {
    revert();
  }

  function swapBorrowRateMode(address, uint256) external pure {
    revert();
  }

  function rebalanceStableBorrowRate(address, address) external pure {
    revert();
  }

  function setUserUseReserveAsCollateral(address, bool) external pure {
    revert();
  }

  function liquidationCall(
    address,
    address,
    address,
    uint256,
    bool
  ) external pure {
    revert();
  }

  function flashLoan(
    address,
    address[] calldata,
    uint256[] calldata,
    uint256[] calldata,
    address,
    bytes calldata,
    uint16
  ) external pure {
    revert();
  }

  function flashLoanSimple(
    address,
    address,
    uint256,
    bytes calldata,
    uint16
  ) external pure {
    revert();
  }

  function getUserAccountData(
    address
  )
    external 
    pure
    returns (
      uint256,
      uint256,
      uint256,
      uint256,
      uint256,
      uint256
    ) {
      revert();
    }

  function initReserve(
    address,
    address,
    address,
    address,
    address
  ) external pure {
    revert();
  }

  function dropReserve(address) external pure {
    revert();
  }

  function setReserveInterestRateStrategyAddress(
    address,
    address
  ) external pure {
    revert();
  }

  function setConfiguration(
    address,
    DataTypes.ReserveConfigurationMap calldata
  ) external pure {
    revert();
  }

  function getConfiguration(
    address
  ) external pure returns (DataTypes.ReserveConfigurationMap memory) {
    revert();
  }

  function getUserConfiguration(
    address
  ) external pure returns (DataTypes.UserConfigurationMap memory) {
    revert();
  }

  function getReserveNormalizedIncome(address) external pure returns (uint256) {
    revert();
  }

  function getReserveNormalizedVariableDebt(address) external pure returns (uint256) {
    revert();
  }

  function finalizeTransfer(
    address,
    address,
    address,
    uint256,
    uint256,
    uint256
  ) external pure {
    revert();
  }

  function getReservesList() external pure returns (address[] memory) {
    revert();
  }

  function getReserveAddressById(uint16) external pure returns (address) {
    revert();
  }

  function ADDRESSES_PROVIDER() external pure returns (IPoolAddressesProvider) {
    revert();
  }

  function updateBridgeProtocolFee(uint256) external pure {
    revert();
  }

  function updateFlashloanPremiums(
    uint128,
    uint128
  ) external pure {
    revert();
  }

  function configureEModeCategory(uint8, DataTypes.EModeCategory memory) external pure {
    revert();
  }

  function getEModeCategoryData(uint8) external pure returns (DataTypes.EModeCategory memory) {
    revert();
  }

  function setUserEMode(uint8) external pure {
    revert();
  }

  function getUserEMode(address) external pure returns (uint256) {
    revert();
  }

  function resetIsolationModeTotalDebt(address) external pure {
    revert();
  }

  function MAX_STABLE_RATE_BORROW_SIZE_PERCENT() external pure returns (uint256) {
    revert();
  }

  function FLASHLOAN_PREMIUM_TOTAL() external pure returns (uint128) {
    revert();
  }

  function BRIDGE_PROTOCOL_FEE() external pure returns (uint256) {
    revert();
  }

  function FLASHLOAN_PREMIUM_TO_PROTOCOL() external pure returns (uint128) {
    revert();
  }

  function MAX_NUMBER_RESERVES() external pure returns (uint16) {
    revert();
  }

  function mintToTreasury(address[] calldata) external pure {
    revert();
  }

  function rescueTokens(address, address, uint256) external pure {
    revert();
  }

  function deposit(address, uint256, address, uint16) external pure {
    revert();
  }
}