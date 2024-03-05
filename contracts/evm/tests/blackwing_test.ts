import chai, { expect } from "chai";
chai.use(require('chai-as-promised'));

import {
  deployMockContract,
  deployBlackwingContract,
  deployMockCoin,
  uuid,
} from "../utils";
import {
  SwapPayloadStruct,
} from "@/hardhat_artifacts/types/contracts/evm/blackwing.sol/BlackwingMach1";

describe("Blackwing", function () {
  it("Should not allow multiple instances of the same orders to be processed", async function () {
    const blackwing = await deployBlackwingContract();
    const blackwingAddress = await blackwing.getAddress();
    const mockOrderId = uuid();

    const fromCoin = await deployMockCoin();
    const fromCoinAddress = await fromCoin.getAddress();
    const toCoin = await deployMockCoin();
    const toCoinAddress = await toCoin.getAddress();

    const startToBalance = 10000;
    const finalToBalance = 10005;
    const startFromBalance = 10000;
    const finalFromBalance = 9990;

    await (await toCoin.mint(blackwingAddress, startToBalance)).wait();
    await (await fromCoin.mint(blackwingAddress, startFromBalance)).wait();

    const mockContract = await deployMockContract();
    const mockContractAddress = await mockContract.getAddress();
    const mockSwapPayload = mockContract.interface.encodeFunctionData(
      "adj",
      [fromCoinAddress, -10, toCoinAddress, 5],
    );

    const swapPayload: SwapPayloadStruct = {
      fromTokenAddress: fromCoinAddress,
      toTokenAddress: toCoinAddress,
      target: mockContractAddress,
      value: 0,
      data: mockSwapPayload,
    };

    expect(await blackwing.isOrderProcessed(mockOrderId)).to.equal(false);

    const tx1 = await blackwing.processOrder(mockOrderId, swapPayload);
    await tx1.wait();
    expect(tx1).to.emit(blackwing, "OrderProcessed").withArgs(mockOrderId, false, {
      fromTokenGiven: 10,
      toTokenReceived: 5,
    });
    expect(await blackwing.isOrderProcessed(mockOrderId)).to.equal(true);

    expect(await toCoin.balanceOf(blackwingAddress)).to.equal(finalToBalance);
    expect(await fromCoin.balanceOf(blackwingAddress)).to.equal(finalFromBalance);

    const tx2 = await blackwing.processOrder(mockOrderId, swapPayload);
    await tx2.wait();
    expect(await blackwing.isOrderProcessed(mockOrderId)).to.equal(true);
    expect(tx2).to.emit(blackwing, "OrderProcessed")
      .withArgs(mockOrderId, true, {fromTokenGiven: 0, toTokenReceived: 0});

    expect(await toCoin.balanceOf(blackwingAddress)).to.equal(finalToBalance);
    expect(await fromCoin.balanceOf(blackwingAddress)).to.equal(finalFromBalance);
  });
});
