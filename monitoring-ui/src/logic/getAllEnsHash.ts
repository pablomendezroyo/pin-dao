import { EnsHash } from "../types/types";
import getHashFromEns from "./getHashFromEns";
import getPinnedDomains from "./getPinnedDomains";

export default async function getAllEnsHash(): Promise<EnsHash[]> {
  //1.Read from SC the list of ENS domains
  //TODO
  //const ensDomains = await getPinnedDomains();
  const ensDomains = [
    "subdomain1.hehehehehe.eth",
    "pinnerdao.eth",
    "subdomain1.hehehehehe.eth",
  ];

  //2. Look for the hash of each domain
  let ensHashes: EnsHash[] = [];
  for (const domain of ensDomains) {
    const hash = await getHashFromEns(domain);
    if (hash) {
      ensHashes.push({ domain, hash });
    }
  }

  return ensHashes;
}
