import { EnsHash } from "../types/types";
import getHashFromEns from "./getHashFromEns";
import getPinnedDomains from "./getPinnedDomains";

export default async function getAllEnsHash(): Promise<EnsHash[]> {
  //1.Read from SC the list of ENS domains
  const ensDomains = await getPinnedDomains();

  //2. Look for the hash of each domain
  let ensHashes: EnsHash[] = [];
  for (const domain of ensDomains) {
    const hash = await getHashFromEns(domain);
    if (hash) {
      //3. Add the domain and the hash to the array
      ensHashes.push({ domain, hash });
    }
  }

  return ensHashes;
}
