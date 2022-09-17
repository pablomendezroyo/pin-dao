package main

import (
	"encoding/json"
	"fmt"
	"io"

	shell "github.com/ipfs/go-ipfs-api"
)

type Manifest struct {
	GH_REPO        string
	COMMIT         string
	IPFS_HASH_REPO string
	ENS            string
}

func resolve_manifest_list(ipfsHashList []string, ipfs_client *shell.Shell) []string {
	var ipfsManifestRepoHashList []string
	for _, ipfsHash := range ipfsHashList {
		files, err := ipfs_client.List(ipfsHash)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		for _, file := range files {
			if file.Name == "manifest.json" {
				manifestReader, err := ipfs_client.Cat(file.Hash)
				if err != nil {
					fmt.Println("Error on ipfs.cat:", err)
					break
				}

				manifestBytes, err := io.ReadAll(manifestReader)
				if err != nil {
					fmt.Println("Error on io.ReadAll:", err)
					break
				}

				manifestString := string(manifestBytes[:])
				manifest := Manifest{}
				json.Unmarshal([]byte(manifestString), &manifest)

				fmt.Println("Manifest:", manifest)
				// TODO: check it is a valid ipfs hash the manifest.IPFS_HASH_REPO
				ipfsManifestRepoHashList = append(ipfsManifestRepoHashList, manifest.IPFS_HASH_REPO)
			}
		}
	}

	return ipfsManifestRepoHashList
}
