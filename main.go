package main

import (
	"fmt"
	"log"
	"mcp-llm-client/cli"
	"mcp-llm-client/llm"
)

func main() {
	fmt.Println("Staring mcp client ...")

	greet, conversation, err := llm.Init()
	if err != nil {
		log.Fatal("Failed to connect to llama.cpp server:", err)
	}

	fmt.Println("Successfully connected to llama.cpp!")

	cli.StartChat(greet, conversation)
}
