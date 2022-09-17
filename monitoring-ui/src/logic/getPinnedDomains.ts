import { ethers } from "ethers";

const providerUrl = "https://rpc.ankr.com/eth_rinkeby";
const provider = new ethers.providers.JsonRpcProvider(providerUrl);

export default async function getPinnedDomains(): Promise<string[]> {
  let contract = new ethers.Contract(
    "0xd2d2AFb9CDAc195d0a665F526B6Ef137118F1291",
    ["function getENSList() public pure override returns (string[] memory)"],
    provider
  );

  let pinnedEns = await contract.getENSList();
  return pinnedEns;
}
