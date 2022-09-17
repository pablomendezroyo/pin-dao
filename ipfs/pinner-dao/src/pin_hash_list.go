package main

import (
	"fmt"

	shell "github.com/ipfs/go-ipfs-api"
)

func pin_hash_list(ipfsHashList []string, ipfs_client *shell.Shell) {
	for _, ipfsHash := range ipfsHashList {
		err := ipfs_client.Pin(ipfsHash)
		if err != nil {
			fmt.Println("Error pinning", ipfsHash, err)
			// Do not exit and continue pinning the rest of the hashes
		}
	}
}
