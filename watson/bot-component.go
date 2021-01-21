package watson

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/go-sdk-core/core"
	"github.com/watson-developer-cloud/go-sdk/v2/assistantv2"
)

// SendMessage send message to Watson and get the intent
func SendMessage(assistant *assistantv2.AssistantV2, message string, merchantId string) string {

	// Get the
	sessionID := GetSession(merchantId, assistant)
	result, _, responseErr := assistant.Message(
		&assistantv2.MessageOptions{
			AssistantID: core.StringPtr(AssistantID),
			SessionID:   sessionID,
			Input: &assistantv2.MessageInput{
				MessageType: core.StringPtr("text"),
				Text:        core.StringPtr(message),
			},
		},
	)

	if responseErr != nil {
		panic(responseErr)
	}
	b, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(b))

	if len(result.Output.Intents) > 0 {
		return ProcessIntent(*result.Output.Intents[0].Intent, merchantId, message)
	}
	return "Pardon human I didn't understand! I need some more training to address your query. Please forgive me for now!" //Change this
}
