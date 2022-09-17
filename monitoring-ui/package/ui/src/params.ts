import { ethers } from "ethers";

export const contractInfo = {
  address: "0xc8BDAe8196ad92dAc0db3B3e1C74a9ac409fB93D",
  abi: ["function getENSList() public pure override returns (string[] memory)"],
};

const providerUrl = "https://rpc.ankr.com/eth_rinkeby";
export const provider = new ethers.providers.JsonRpcProvider(providerUrl);
