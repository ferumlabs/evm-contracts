// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.23;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {AccessControlUpgradeable} from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import {IERC20} from "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";

// The data needed to make a call to another function. Used to arbitrarily call 
// swap functions.
struct SwapPayload {
  address fromTokenAddress;
  address toTokenAddress;
  address target;
  uint value;
  bytes data;
}

// Describes the result of a swap.
struct SwapResult {
  uint fromTokenGiven;
  uint toTokenReceived;
}

error TransactionFailed(address target, bytes inputdata, bytes returndata);

contract BlackwingMach1 is Initializable, AccessControlUpgradeable {
  bytes32 public constant OWNER_ROLE = keccak256("OWNER_ROLE");

  mapping(bytes16 => bool) processedOrders;

  event OrderProcessed(bytes16 id, bool duplicate, SwapResult result);

  function initialize() public initializer {
    AccessControlUpgradeable.__AccessControl_init();
    AccessControlUpgradeable._grantRole(OWNER_ROLE, msg.sender);
  }

  function processOrder(bytes16 id, SwapPayload memory swapPayload) public returns (bool, SwapResult memory) {
    require(
      hasRole(OWNER_ROLE, msg.sender),
      "unauthorized"
    );
    
    if (processedOrders[id]) {
      emit OrderProcessed(id, true, SwapResult(0, 0));
      return (false, SwapResult(0, 0));
    }
    processedOrders[id] = true;

    uint preSwapFromTokenBalance = IERC20(swapPayload.fromTokenAddress).balanceOf(address(this));
    uint preSwapToTokenBalance = IERC20(swapPayload.toTokenAddress).balanceOf(address(this));

    (bool success, bytes memory returndata) = swapPayload.target.call{value: swapPayload.value}(swapPayload.data);
    if (!success) revert TransactionFailed(swapPayload.target, swapPayload.data, returndata);

    uint postSwapFromTokenBalance = IERC20(swapPayload.fromTokenAddress).balanceOf(address(this));
    uint postSwapToTokenBalance = IERC20(swapPayload.toTokenAddress).balanceOf(address(this));

    // This bakes in the assumption that the swap decreases the from token balance and increases the to 
    // token balance.
    SwapResult memory res = SwapResult(
      sub(preSwapFromTokenBalance, postSwapFromTokenBalance),
      sub(postSwapToTokenBalance, preSwapToTokenBalance)
    );
    
    emit OrderProcessed(id, false, res);
    
    return (true, res);
  }

  // Subtract b from a.
  function sub(uint a, uint b) internal pure returns (uint) {
    assert(b <= a);
    return a - b;
  }

  function isOrderProcessed(bytes16 id) public view returns (bool) {
    return processedOrders[id];
  }
}