package cli

import (
	"bufio"
	"fmt"
	"mcp-llm-client/sessionmanager"
	"mcp-llm-client/thinktank"
	"os"
)

var sessionId int64

func StartChat() {
	fmt.Println("---------------------------------------------------------------")

	scanner := bufio.NewScanner(os.Stdin)
	sessionData := sessionmanager.GetOrCreateSessionManager().CreateSession()
	sessionId = sessionData.SessionId

	thinkTank := thinktank.GetThinkTank()

	greet, err := thinkTank.StartConversation(sessionId)
	if err != nil {
		panic(err)
	}

	fmt.Println("AI: ", greet)

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
		aiResponse, err = thinkTank.Converse(userInput, sessionData.SessionId)

		if err != nil {
			panic(err)
		}

		fmt.Printf("\nAI: %v", aiResponse)
		fmt.Print("\n\n")
	}
}
