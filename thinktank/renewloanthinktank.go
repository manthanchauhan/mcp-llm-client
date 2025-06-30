package thinktank

import (
	"mcp-llm-client/llm"
	"mcp-llm-client/sessionmanager"
)

type RenewLoanThinkTank struct {
}

func (tt *RenewLoanThinkTank) Converse(userInput string, sid int64) (string, error) {
	sm := sessionmanager.GetOrCreateSessionManager()
	sd := sm.GetSession(sid)
	convHist := sd.GetConversationHistory()
	llmResp, conv, err := llm.SendUserMessage(userInput, convHist)
	if err != nil {
		return "", err
	}
	sd.SetConversationHistory(conv)
	return llmResp, nil
}
