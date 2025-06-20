package dto

type FunctionCallAssistantReply struct {
	Name      string                 `json:"name"`
	Arguments map[string]interface{} `json:"arguments"`
}

func (*FunctionCallAssistantReply) Execute() {

}
