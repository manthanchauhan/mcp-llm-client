package dto

import functioncall "mcp-llm-client/functioncall/dto"

type AssistantReply struct {
	ReplyToUser    *string                                  `json:"reply_to_human"`
	FunctionCall   *functioncall.FunctionCallAssistantReply `json:"function_call"`
	InfoExtraction *map[string]any                          `json:"info_extraction"`
}
