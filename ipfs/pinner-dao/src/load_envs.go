package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func load_envs() (string, string) {
	var environment = os.Getenv("ENVIRONMENT")
	if environment == "" {
		/* 		fmt.Println("Error loading environment")
		   		os.Exit(1) */
		environment = "development"
	}

	if environment == "production" {
		fmt.Println("Loading production environment")
		err := godotenv.Load("prod.env")
		if err != nil {
			fmt.Println("Error loading prod environment", err)
			os.Exit(1)
		}
	} else if environment == "development" {
		fmt.Println("Loading development environment")
		err := godotenv.Load("dev.env")
		if err != nil {
			fmt.Println("Error loading dev environment", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("ENVIRONMENT not set, please set ENVIRONMENT to production or development")
		os.Exit(1)
	}

	var ipfs_api_url = os.Getenv("IPFS_API_URL")
	if ipfs_api_url == "" {
		fmt.Println("IPFS_API_URL environment variable is not set")
		os.Exit(1)
	}

	var eth_rpc_url = os.Getenv("ETH_RPC_URL")
	if eth_rpc_url == "" {
		fmt.Println("ETH_RPC_URL environment variable is not set")
		os.Exit(1)
	}

	return ipfs_api_url, eth_rpc_url
}
