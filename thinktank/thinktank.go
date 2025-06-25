package thinktank

import (
	"encoding/json"
	"fmt"
	"mcp-llm-client/llm"
	"mcp-llm-client/llm/dto"
	"mcp-llm-client/sessionmanager"
)

var singletonThinkTank *ThinkTank

type ThinkTank struct {
}

func (tt *ThinkTank) Converse(userInput string, sid int64) (string, error) {
	sm := sessionmanager.GetOrCreateSessionManager()
	sd := sm.GetSession(sid)

	if sd.CustomerRequestCategory == nil {
		rcFound, llmReply, err := tt.converseToIdentifyCustomerRequestCategory(userInput, sd)
		if err != nil {
			return "", err
		}

		if !rcFound {
			return llmReply, nil
		}

		userInput = ""
	}

	if sd.UserData == nil || sd.UserData.Id == nil {
		udFound, llmReply, err := tt.converseToIdentifyUser(userInput, sd)
		if err != nil {
			return "", err
		}

		if !udFound {
			return llmReply, nil
		}
	}

	return "", nil
}

func (tt *ThinkTank) StartConversation(sid int64) (string, error) {
	initMsg := llm.INITMESSAGE
	llmInst := llm.GetLLM()
	sm := sessionmanager.GetOrCreateSessionManager()

	sd := sm.GetSession(sid)
	sd.ResetConversationHistory()

	convHist := sd.GetConversationHistory()

	greet, conv, err := llmInst.SendSystemMessage(initMsg, convHist)
	if err != nil {
		return "", err
	}

	sd.SetConversationHistory(conv)
	return greet, nil
}

func (tt *ThinkTank) converseToIdentifyCustomerRequestCategory(userInput string, sd *sessionmanager.SessionData) (bool, string, error) {
	convHist := sd.ConversationHistory

	userMsg := dto.Message{Role: "user", Content: userInput}

	if convHist == nil {
		convHist = []dto.Message{userMsg}
	} else {
		convHist = append(convHist, userMsg)
	}

	systMsg := `From the conversation identify ONE request category for which the USER needs support. 
	
	These are the ONLY VALID CATEGORIES:
		1. ABOUT_INDIAGOLD_COMPANY
		2. RENEW_EXISTING_LOAN
		3. CLOSE_EXISTING_LOAN
		4. BOOK_NEW_LOAN

	If VALID request category is identified: Respond in this EXACT JSON: {"info_extraction": {"request_category": <request-category>}}
	Else: Ask the user what he needs help with and respond in this EXACT JSON: {"reply_to_human": <reply-to-human>}

	If user is unable to decide, give options to the user
	`

	llmResp, convHist, err := llm.GetLLM().SendSystemMessage(systMsg, convHist)

	if err != nil {
		return false, "", err
	}

	sd.SetConversationHistory(convHist)

	var assitantReply dto.AssistantReply

	if err := json.Unmarshal([]byte(llmResp), &assitantReply); err == nil {
		info := assitantReply.InfoExtraction

		if info != nil {
			sessionmanager.GetOrCreateSessionManager().EnrichSessionData(sd.SessionId, info)
			return true, "", nil
		}

		userReply := assitantReply.ReplyToUser

		if userReply != nil {
			return false, *userReply, nil
		}

		return false, "", fmt.Errorf("no reply for user form llm")
	}

	return false, "", err
}

func (tt *ThinkTank) converseToIdentifyUser(userInput string, sd *sessionmanager.SessionData) (bool, string, error) {
	convHist := sd.ConversationHistory

	if userInput != "" {
		userMsg := dto.Message{Role: "user", Content: userInput}

		if convHist == nil {
			convHist = []dto.Message{userMsg}
		} else {
			convHist = append(convHist, userMsg)
		}
	}

	systMsg := `From the conversation extract the following fields,
	
	1. USER_MOBILE_NUMBER

	If ANY of above fields is identified: Respond in this EXACT JSON: {"info_extraction": {"user_data": {<field-name-1>: <field-value-1>}}}
	Else: Ask the user to provide any of the above data and respond in this EXACT JSON: {"reply_to_human": <reply-to-human>}
	`

	// filteredConvHist := llm.FilterSystemMessagesExceptInitMsg(convHist)
	llmResp, convHist, err := llm.GetLLM().SendSystemMessage(systMsg, convHist)

	if err != nil {
		return false, "", err
	}

	sd.SetConversationHistory(convHist)

	var assitantReply dto.AssistantReply

	if err := json.Unmarshal([]byte(llmResp), &assitantReply); err == nil {
		info := assitantReply.InfoExtraction

		if info != nil {
			sessionmanager.GetOrCreateSessionManager().EnrichSessionData(sd.SessionId, info)
			return true, "", nil
		}

		userReply := assitantReply.ReplyToUser

		if userReply != nil {
			return false, *userReply, nil
		}

		return false, "", fmt.Errorf("no reply for user form llm")
	}

	return false, "", err
}

func GetThinkTank() *ThinkTank {
	if singletonThinkTank == nil {
		singletonThinkTank = &ThinkTank{}
	}

	return singletonThinkTank
}
