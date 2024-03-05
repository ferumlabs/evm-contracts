import { task } from "hardhat/config";
import { MockCoinContractName } from "@/contracts/evm/constants";

task("fund", "Fund account with assets")
  .addPositionalParam("account", "Address of Account to fund")
  .addPositionalParam("amount", "Amount to fund account with, denominated in asset decimals.")
  .addOptionalPositionalParam("asset", "Address of MockCoin Asset to fund with. If not provided, will fund with native currency of the chain.")
  .setAction(async (taskArgs, hre) => {
    const {account, amount, asset} = taskArgs;
    if (asset) {
      const assetInstance = await hre.ethers.getContractAt(MockCoinContractName, asset);
      const assetSymbol = await assetInstance.symbol();
      await assetInstance.mint(account, amount);
      const assetBalance = await assetInstance.balanceOf(account);
      console.log(`Account ${account} has ${assetBalance.toString()} ${assetSymbol}`);
    } else {
      const [_, ...accounts] = await hre.ethers.getSigners();
      for (let fromAcc of accounts) {
        const balance = await hre.ethers.provider.getBalance(await fromAcc.getAddress());
        if (balance > amount) {
          await fromAcc.sendTransaction({
            to: account,
            value: amount,
          });
          const accountBalance = await hre.ethers.provider.getBalance(account);
          console.log(`Account ${account} has ${accountBalance.toString()} native tokens (ETH, BNB, etc)`);
          break;
        }
      }
    }
  });
