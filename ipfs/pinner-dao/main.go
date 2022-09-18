package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting pinner DAO")

	// Load envs
	ipfs_api_url, eth_rpc_url := load_envs()
	fmt.Println("IPFS API URL:", ipfs_api_url)
	fmt.Println("Ethereum RPC URL:", eth_rpc_url)

	// Get ipfs client instance
	var ipfs_client = get_ipfs_client(ipfs_api_url)
	// Get eth client instance
	var eth_client = get_eth_client(eth_rpc_url)
	// Get the smart contract instance
	var ensList = get_ens_list(eth_client)
	fmt.Println("ENS List:", ensList)

	// Resolve the list of ENS names to IPFS hashes
	var ipfsBuildHashList = resolve_ens_list(ensList, eth_client)
	fmt.Println("Ipfs build hashes:", ipfsBuildHashList)

	// Resolve the list of IPFS repository hashes from the manifest to IPFS content
	var ipfsManifestHashList = resolve_manifest_list(ipfsBuildHashList, ipfs_client)
	fmt.Println("Ipfs manifest hashes:", ipfsManifestHashList)

	// JOIN the two lists
	var ipfsHashList = append(ipfsBuildHashList, ipfsManifestHashList...)

	// Pin the list of IPFS hashes
	pin_hash_list(ipfsHashList, ipfs_client)
}
