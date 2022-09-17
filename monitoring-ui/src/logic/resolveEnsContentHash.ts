import { ENS } from "@ensdomains/ensjs";
import { ethers } from "ethers";

const providerUrl = "http://fullnode.dappnode:8545";

const provider = new ethers.providers.JsonRpcProvider(providerUrl);

const ENSInstance = new ENS();

export default async function resolveEnsContentHash(ensDomain: string) {
  await ENSInstance.setProvider(provider);
}
