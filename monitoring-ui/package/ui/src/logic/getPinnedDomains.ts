import { ethers } from "ethers";
import { contractInfo, provider } from "../params";

export default async function getPinnedDomains(): Promise<string[]> {
  const contract = new ethers.Contract(
    contractInfo.address,
    contractInfo.abi,
    provider
  );

  return await contract.getENSList();
}
