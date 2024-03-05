package main

//go:generate npx hardhat compile
//go:generate abigen -abi ./hardhat_abi/contracts/evm/blackwing.sol/BlackwingMach1.json -out ./contracts/evm/blackwing_gen.go -pkg evm -type Blackwing
//go:generate abigen -abi ./hardhat_abi/contracts/evm/launch_vault/vault.sol/BlackwingVault.json -out ./contracts/evm/launch_vault_gen.go -pkg evm -type LaunchVault
//go:generate abigen -abi ./hardhat_abi/@openzeppelin/contracts/token/ERC20/ERC20.sol/ERC20.json -out ./contracts/evm/ierc20_gen.go -pkg evm -type ERC20
