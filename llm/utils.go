package llm

import "mcp-llm-client/llm/dto"

func FilterSystemMessagesExceptInitMsg(convHist []dto.Message) []dto.Message {
	if len(convHist) == 0 {
		return convHist
	}

	filteredConvHist := []dto.Message{convHist[0]}

	for i := 1; i < len(convHist); i++ {
		msg := convHist[i]

		if msg.Role != "system" {
			filteredConvHist = append(filteredConvHist, msg)
		}
	}

	return filteredConvHist
}
