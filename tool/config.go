package tool

import mcp "github.com/manthanchauhan/mcp-go-util/mcp"

var TOOLLIST = []mcp.Tool{
	{
		Name:        "get_user_by_mobile",
		Description: "Get user information by mobile number",
		InputSchema: mcp.InputSchema{
			Type:     mcp.OBJECT_TYPE,
			Required: []string{"mobile"},
			Properties: map[string]mcp.InputSchemaProperty{
				"mobile": {
					Type:        mcp.STRING_TYPE,
					Description: "Mobile number of the user",
				},
			},
		},
	},
}
