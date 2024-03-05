// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.23;

import { MockCoin } from "./mock_coin.sol";

contract MockContract {
  constructor() {}

  function fail() public pure {
    revert("fail");
  }
  
  function success() public pure {
    return;
  }

  function adj(MockCoin tokenA, int tokenAAmt, MockCoin tokenB, int tokenBAmt) public {
    if (tokenAAmt < 0) {
      tokenA.burn(msg.sender, uint(-tokenAAmt));
    } else {
      tokenA.mint(msg.sender, uint(tokenAAmt));
    }
    if (tokenBAmt < 0) {
      tokenB.burn(msg.sender, uint(-tokenBAmt));
    } else {
      tokenB.mint(msg.sender, uint(tokenBAmt));
    }
  }
}