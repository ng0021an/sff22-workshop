# Rocksolid back-end app for SFF Web3 Workshop 
This repo is the RockSolid back-end app presented during the SFF Web3 Workshop.
The role of this back-end server is to connect your application to an EVM blockchain node (Ethereum, Polygon, ...), and send assets from a wallet to a customer's wallet.

- [Rocksolid back-end app for SFF Web3 Workshop](#rocksolid-back-end-app-for-sff-web3-workshop)
  - [Pre-requisite](#pre-requisite)
    - [MacOs Users](#macos-users)
    - [Ubuntu/Debian Users](#ubuntudebian-users)
  - [1. Setup the back-end service](#1-setup-the-back-end-service)
    - [Install the dependencies](#install-the-dependencies)
    - [Create the .env config file](#create-the-env-config-file)
  - [2. Build and run the server](#2-build-and-run-the-server)
  - [3. Test the server](#3-test-the-server)
    - [Make a request to airdrop ERC1155 tokens](#make-a-request-to-airdrop-erc1155-tokens)
- [Appendix](#appendix)
  - [Appendix 1: Deploy your ERC-1155 Contract](#appendix-1-deploy-your-erc-1155-contract)
    - [1. Install npm dependencies](#1-install-npm-dependencies)
    - [2. (Optional) Edit the Solidity contracts](#2-optional-edit-the-solidity-contracts)
    - [3. Compile Solidity contracts](#3-compile-solidity-contracts)
    - [4. Deploy contracts](#4-deploy-contracts)
    - [Note](#note)

## Pre-requisite
You will need on your computer: 
* An EVM Wallet (Private key and Mnemonic). We recommend to use a temporary wallet. DO NOT use an important wallet.
  * Suggestion: [Coinbase Wallet](https://www.coinbase.com/wallet)
  * The wallet mnemonic
* Golan ([Install page](https://go.dev/doc/install))
* Node ([Install page](https://nodejs.org/en/download/))
* NPM

### MacOs Users
```bash
# Install Brew
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
brew update && brew doctor
brew install node
brew install go
```

### Ubuntu/Debian Users
```bash
sudo apt install nodejs npm golang-go
```

## 1. Setup the back-end service

### Install the dependencies
```bash
make deps
```

### Create the .env config file
The `.env` file contains all sensitive information to connect to your blockchain node and manage the wallet.
Create the `.env` file at the root of this project
```bash
cp env.sample .env
```

Edit the `.env` file values
```bash
vi .env
```

```ini
USERNAME=<Username of Coinbase Cloud account>
PASSWORD=<Password of Coinbase Cloud account>
NODE_URI="goerli.ethereum.coinbasecloud.net"
MNEMONIC=<Mnemonic of the wallet holding the ERC1155 tokens>
CONTRACT_ADDRESS=<Contract address of the ERC1155 tokens>
MAX_GOLD_BADGE_TOTAL_QUANTITY=<Max number of gold badges owned by one user>
MAX_GOLD_BADGE_TRANSFER_QUANTITY=<Max number of gold badges per transfer>
MAX_POINT_TOTAL_QUANTITY=<Max number of points owned by one user>
MAX_POINT_TRANSFER_QUANTITY=<Max number of points per transfer>

```

## 2. Build and run the server

```bash
make build && make run
```

## 3. Test the server
### Make a request to airdrop ERC1155 tokens
```
curl --url 'http://localhost:8081/gettoken?to=<the address to airdrop tokens to>&id=<id of the nft item>&quantity=<amount of the nft item>'
```
Example
```
curl --url 'http://localhost:8081/gettoken?to=0xF820cf368b4a798b676DE9DEA90f637A9CdEE572&id=2&quantity=3'
```

# Appendix
## Appendix 1: Deploy your ERC-1155 Contract

You can choose to use the Remix-IDE to deploy the contract manually. Please follow the instruction [here](https://remix-ide.readthedocs.io/en/latest/create_deploy.html#deploy-the-contract)

You might want to use the same MNEMONIC that you specify in the .env file so that you can directly transfer from the same wallet.

