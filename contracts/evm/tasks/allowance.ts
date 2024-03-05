import { task } from "hardhat/config";

task("allowance", "Set ERC20 allowance")
  .addPositionalParam("account", "Address of Account to configure allowance for")
  .addPositionalParam("erc20", "ERC20 asset to modify allowance for")
  .addPositionalParam("spender", "Spender allowance is being approved for")
  .addPositionalParam("amount", "Amount allowance is being approved for, in asset decimals")
  .setAction(async (taskArgs, hre) => {
    const {account, erc20, spender, amount} = taskArgs;
    await hre.network.provider.request({
      method: "hardhat_impersonateAccount",
      params: [account],
    });
    const accountSigner = await hre.ethers.getSigner(account);
    const contract = await hre.ethers.getContractAt("ERC20", erc20);
    await contract.connect(accountSigner).approve(spender, amount);
    const allowance = await contract.allowance(account, spender);
    const symbol = await contract.symbol();
    console.log(`Allowance for ${spender} set to ${allowance.toString()} for ${symbol}`);
  });
