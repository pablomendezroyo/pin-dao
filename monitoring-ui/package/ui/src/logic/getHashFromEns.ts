import { provider } from "../params";

export default async function getHashFromEns(ensName: string) {
  const resolver = await provider.getResolver(ensName);
  const hash = await resolver?.getContentHash();
  return hash;
}
