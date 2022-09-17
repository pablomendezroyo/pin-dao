package main

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

func get_eth_client(eth_rpc_url string) *ethclient.Client {
	client, err := ethclient.Dial(eth_rpc_url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ethereum node is not up, make sure the endpoint %s is up", eth_rpc_url)
		os.Exit(1)
	}
	return client
}
