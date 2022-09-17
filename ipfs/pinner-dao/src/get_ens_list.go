package main

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const smart_contract_address = "0x0000000000000000000000000000000000000000"

func get_ens_list(eth_client_provider *ethclient.Client) []string {
	sc, err := NewMain(common.HexToAddress(smart_contract_address), eth_client_provider)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to instantiate a smart contract: %v ", err)
		os.Exit(1)
	}

	ensList, err := sc.GetENSList(nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get ENS list: %v ", err)
		os.Exit(1)
	}

	return ensList
}
