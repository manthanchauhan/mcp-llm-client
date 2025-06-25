package main

import (
	"fmt"
	"log"
	"mcp-llm-client/cli"
	"mcp-llm-client/llm"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Staring mcp client ...")

	_, _, err = llm.Init()
	if err != nil {
		log.Fatal("Failed to connect to LLM server:", err)
	}

	modelName := llm.GetLLM().ModelName
	fmt.Println("Successfully connected to LLM model " + modelName)

	cli.StartChat()
}
