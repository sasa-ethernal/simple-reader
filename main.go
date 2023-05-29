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

	fmt.Println("RPC Endpoint:", rpcEndpoint)
	fmt.Println("Start block:", startBlock)
	fmt.Println("End block:", endBlock)

	client, err := rpc.Dial(rpcEndpoint)
	if err != nil {
		log.Fatal("Failed to connect to the RPC endpoint:", err)
	}

	for i := startBlock; i <= endBlock; i++ {
		blockNumber := fmt.Sprintf("0x%x", i)

		var block map[string]interface{}
		err := client.Call(&block, "eth_getBlockByNumber", blockNumber, true)
		if err != nil {
			log.Fatal("Error retrieving block", i, ":", err)
		}

		fmt.Println("Block", i, ":", block)
	}

}
