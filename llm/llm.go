package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"mcp-llm-client/llm/dto"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

var rotator = &lumberjack.Logger{Filename: "app.log", MaxSize: 10}
var jsonLogger = slog.New(slog.NewJSONHandler(rotator, nil))
var singletonLLM *LLM

type LLM struct {
	ModelName string
}

func (llm *LLM) StartConversation() {

}
func (l *LLM) SendSystemMessage(message string, conversation []dto.Message) (string, []dto.Message, error) {
	sysMsg := dto.Message{
		Role:    "system",
		Content: message,
	}

	conversation = append(conversation, sysMsg)

	chatReq := dto.ChatRequest{
		Model:     os.Getenv("MODEL_NAME"),
		Messages:  conversation,
		MaxTokens: MAXTOKENS,
	}

	reply, err := getChatCompletion(chatReq)

	if err != nil {
		return "", conversation, err
	}

	conversation = append(conversation, dto.Message{Role: "assistant", Content: reply})
	return reply, conversation, nil
}

func GetLLM() *LLM {
	if singletonLLM == nil {
		singletonLLM = &LLM{ModelName: os.Getenv("MODEL_NAME")}
	}

	return singletonLLM
}

func Init() (string, []dto.Message, error) {
	// fmt.Printf("System Message: %v", SystemMsg)
	reply, conversation, err := SendSystemMessage(createSystemMessage(), []dto.Message{})

	if err != nil {
		return "", nil, err
	}

	return reply, conversation, nil
}

func SendSystemMessage(message string, conversation []dto.Message) (string, []dto.Message, error) {
	sysMsg := dto.Message{
		Role:    "system",
		Content: message,
	}

	conversation = append(conversation, sysMsg)

	chatReq := dto.ChatRequest{
		Model:     os.Getenv("MODEL_NAME"),
		Messages:  conversation,
		MaxTokens: MAXTOKENS,
	}

	reply, err := getChatCompletion(chatReq)

	if err != nil {
		return "", conversation, err
	}

	conversation = append(conversation, dto.Message{Role: "assistant", Content: reply})
	return reply, conversation, nil
}

func SendUserMessage(message string, conversation []dto.Message) (string, []dto.Message, error) {
	conversation = append(conversation, dto.Message{Role: "user", Content: message})

	chatReq := dto.ChatRequest{
		Model:     os.Getenv("MODEL_NAME"),
		Messages:  conversation,
		MaxTokens: MAXTOKENS,
	}

	reply, err := getChatCompletion(chatReq)

	conversation = append(conversation, dto.Message{Role: "assistant", Content: reply})
	return reply, conversation, err
}

func GetChatCompletion(messages []dto.Message) (string, error) {
	chatReq := dto.ChatRequest{
		Model:     os.Getenv("MODEL_NAME"),
		Messages:  messages,
		MaxTokens: MAXTOKENS,
	}

	return getChatCompletion(chatReq)
}

func getChatCompletion(chatReq dto.ChatRequest) (string, error) {
	reply, err := getChatCompletionCore(chatReq)

	for pendingRetry := 3; err != nil && pendingRetry > 0; pendingRetry -= 1 {
		reply, err = getChatCompletionCore(chatReq)
	}

	return reply, err
}

func getChatCompletionCore(chatReq dto.ChatRequest) (string, error) {
	// Convert to JSON
	jsonData, err := json.Marshal(chatReq)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	jsonLogger.Info("", "Chat Request", string(jsonData))

	// Send request
	client := &http.Client{Timeout: 30 * time.Second}

	modelApiKey := os.Getenv("MODEL_API_KEY")
	modelUrl := os.Getenv("MODEL_URL")

	req, err := http.NewRequest("POST", modelUrl+"/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Length", strconv.Itoa(len(jsonData)))
	req.Header.Set("Host", "localhost:8080")

	if modelApiKey != "" {
		req.Header.Set("Authorization", "Bearer "+modelApiKey)
	}

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	jsonLogger.Info("", "Chat Response", body)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("server error (%d): %s", resp.StatusCode, string(body))
	}

	// Parse response
	var chatResp dto.ChatResponse
	if err := json.Unmarshal(body, &chatResp); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	if len(chatResp.Choices) == 0 {
		return "", fmt.Errorf("no response choices received")
	}

	replyStr := chatResp.Choices[0].Message.Content

	jsonLogger.Info("", "LLM Response", replyStr)

	replyStr = strings.Trim(replyStr, " ")
	replyStr = strings.Trim(replyStr, "<|Assistant|>")

	return replyStr, nil
}

func createSystemMessage() string {
	return INITMESSAGE
}
