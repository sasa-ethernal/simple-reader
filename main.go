package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <RPC endpoint URI> <start block> <end block>")
		return
	}

	rpcEndpoint := os.Args[1]
	startBlockStr := os.Args[2]
	endBlockStr := os.Args[3]

	startBlock, err := strconv.Atoi(startBlockStr)
	if err != nil {
		log.Fatal("Invalid start block:", err)
	}

	endBlock, err := strconv.Atoi(endBlockStr)
	if err != nil {
		log.Fatal("Invalid end block:", err)
	}

	if endBlock < startBlock {
		log.Fatal("End block must be larger than start block")
	}

	fmt.Println("RPC Endpoint:", rpcEndpoint)
	fmt.Println("Start block:", startBlock)
	fmt.Println("End block:", endBlock)
	fmt.Println()

	client, err := rpc.Dial(rpcEndpoint)
	if err != nil {
		log.Fatal("Failed to connect to the RPC endpoint:", err)
	}

	for i := startBlock; i <= endBlock; i++ {
		blockNumber := fmt.Sprintf("0x%x", i) // Convert block number to hexadecimal string

		var block map[string]interface{}
		err := client.Call(&block, "eth_getBlockByNumber", blockNumber, true)
		if err != nil {
			log.Fatal("Error retrieving block", i, ":", err)
		}

		transactions, ok := block["transactions"].([]interface{})

		fmt.Println("Block", i, "Information:")
		fmt.Println("  Block Hash:", block["hash"])
		fmt.Println("  Parent Hash:", block["parentHash"])

		if !ok || len(transactions) == 0 {
			fmt.Println("  No transactions in this block")
		} else {
			fmt.Println("  Transactions:")
			for _, tx := range transactions {
				txData := tx.(map[string]interface{})
				fmt.Println("    Transaction Hash:", txData["hash"])
				fmt.Println("    From:", txData["from"])
				fmt.Println("    To:", txData["to"])
				fmt.Println("    Value:", txData["value"])
				fmt.Println()
			}
		}

		fmt.Println()
	}

}
