package main

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wealdtech/go-ens/v3"
)

func resolve_ens_list(ensList []string, client *ethclient.Client) []string {
	var ipfsHashList []string
	for _, ensName := range ensList {
		// resolve the ENS name to an IPFS hash

		resolver, err := ens.NewResolver(client, ensName)
		if err != nil {
			fmt.Println("Error getting resolver for", ensName, err)
			os.Exit(1)
		}

		bytesHash, err := resolver.Contenthash()
		if err != nil {
			fmt.Println("Error getting contenthash for", ensName, err)
			os.Exit(1)
		}

		ipfsHash, err := ens.ContenthashToString(bytesHash)
		if err != nil {
			fmt.Println("Error converting contenthash to string for", ensName, err)
			os.Exit(1)
		}

		// convert the bytes hash to a string hash using ipfs

		ipfsHashList = append(ipfsHashList, ipfsHash)
	}

	return ipfsHashList
}
