import { ethers } from "ethers";

const provider = new ethers.providers.JsonRpcProvider();

export default async function getPinnedDomains() {
  let contract = new ethers.Contract(
    "0x5FbDB2315678afecb367f032d93F642f64180aa3",
    ["function getDomains() public view returns (string[] memory)"],
    provider
  );
}
