import { ethers } from "ethers";

const providerUrl = "https://rpc.ankr.com/eth_rinkeby";
const provider = new ethers.providers.JsonRpcProvider(providerUrl);

export default async function getHashFromEns(ensName: string) {
  const resolver = await provider.getResolver(ensName);
  const hash = await resolver!.getContentHash();
  return hash;
}
