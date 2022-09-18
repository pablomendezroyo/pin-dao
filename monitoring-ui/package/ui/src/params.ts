import { ethers } from "ethers";

export const contractInfo = {
  address: "0x8aF9F4Bc25Fd5e8B8c3D41D6bDFFB80306918f71",
  abi: ["function getENSList() public pure override returns (string[] memory)"],
};

const providerUrl = "https://rpc.ankr.com/eth_rinkeby";
export const provider = new ethers.providers.JsonRpcProvider(providerUrl);
