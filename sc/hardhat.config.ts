import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import "@nomicfoundation/hardhat-chai-matchers";
import "@nomiclabs/hardhat-ethers";
require('dotenv').config({path: './process.env'});

const { MAINNET_URL, GOERLI_URL, RINKEBY_URL, PRIVATE_KEY, ETHERSCAN_API_KEY } = process.env;

const config: HardhatUserConfig = {
  networks: {
    hardhat: {
    },
    mainnet: {
      url: `${MAINNET_URL}`,
      accounts: [`${PRIVATE_KEY}`],
    },
    goerli: {
      url: `${GOERLI_URL}`,
      accounts: [`${PRIVATE_KEY}`],
    },
    rinkeby: {
      url: `${RINKEBY_URL}`,
      accounts: [`${PRIVATE_KEY}`],
    }
  },
  etherscan: {
    apiKey: `${ETHERSCAN_API_KEY}`,
  },
  solidity: "0.8.17",
};

export default config;
