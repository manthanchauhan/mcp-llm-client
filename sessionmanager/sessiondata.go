package sessionmanager

import (
	"mcp-llm-client/llm/dto"
)

type SessionData struct {
	SessionId               int64            `json:"session_data"`
	UserData                *SessionUserData `json:"session_user_data"`
	ConversationHistory     []dto.Message    `json:"conversation_history"`
	CustomerRequestCategory *string          `json:"request_category"`
	LoanIDs                 []int            `json:"loan_ids"`
}

func (sd *SessionData) ResetConversationHistory() {
	sd.ConversationHistory = []dto.Message{}
}

func (sd *SessionData) SetConversationHistory(ch []dto.Message) {
	sd.ConversationHistory = ch
}

func (sd *SessionData) GetConversationHistory() []dto.Message {
	return sd.ConversationHistory
}
