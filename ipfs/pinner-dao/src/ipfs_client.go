package main

import (
	"fmt"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

func get_ipfs_client(ipfs_api_url string) *shell.Shell {
	sh := shell.NewShell(ipfs_api_url)
	var IsUp = sh.IsUp()
	if IsUp == false {
		fmt.Fprintf(os.Stderr, "IPFS is not up, make sure the ipfs endpoint %s is up", ipfs_api_url)
		os.Exit(1)
	}

	return sh
}
