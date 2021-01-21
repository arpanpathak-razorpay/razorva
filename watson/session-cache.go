package watson

import "github.com/watson-developer-cloud/go-sdk/v2/assistantv2"

// In memory map to hold session merchant_id

var Session map[string]string = map[string]string{}

func GetSession(merchantId string, assistant *assistantv2.AssistantV2) *string {
	// If merchant session already present then return sessionId
	if sessionId, ok := Session[merchantId]; ok {
		return &sessionId
	}

	result, _, responseErr := assistant.CreateSession(assistant.
		NewCreateSessionOptions(AssistantID))
	if responseErr != nil {
		panic(responseErr)
	}
	// b, _ := json.MarshalIndent(result, "", "  ")
	// fmt.Println(string(b))
	// fmt.Println(response)

	// Store the newly created session id in Map
	Session[merchantId] = *result.SessionID
	return result.SessionID
}
