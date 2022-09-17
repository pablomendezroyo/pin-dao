package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting pinner DAO")

	// load envs
	ipfs_api_url, eth_rpc_url := load_envs()
	fmt.Println("IPFS API URL:", ipfs_api_url)
	fmt.Println("Ethereum RPC URL:", eth_rpc_url)

	// Get ipfs client instance
	var ipfs_client = get_ipfs_client(ipfs_api_url)
	// Get eth client instance
	var get_eth_client = get_eth_client(eth_rpc_url)

}
