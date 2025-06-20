package cli

import (
	"bufio"
	"fmt"
	"mcp-llm-client/llm"
	"mcp-llm-client/llm/dto"
	"os"
)

func StartChat(greet string, conversation []dto.Message) {
	fmt.Println("---------------------------------------------------------------")
	fmt.Printf("AI: %v\n", greet)
	fmt.Print("\n")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("User: ")

		if !scanner.Scan() {
			break
		}

		userInput := scanner.Text()

		if userInput == "quit" || userInput == "exit" {
			fmt.Println("Goodbye! 👋")
			break
		}

		if userInput == "" {
			continue
		}

		aiResponse := ""
		var err error
		aiResponse, conversation, err = llm.SendUserMessage(userInput, conversation)

		if err != nil {
			panic(err)
		}

		fmt.Printf("\nAI: %v", aiResponse)
		fmt.Print("\n\n")
	}
}
