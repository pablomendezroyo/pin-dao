package main

import (
	"fmt"
	"os"

	"github.com/pablomendezroyo/pinner-dao/contract"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const smart_contract_address = "0xc8BDAe8196ad92dAc0db3B3e1C74a9ac409fB93D"

func get_ens_list(eth_client_provider *ethclient.Client) []string {
	sc, err := contract.NewContract(common.HexToAddress(smart_contract_address), eth_client_provider)
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
