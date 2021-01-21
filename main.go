package main

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/go-sdk-core/core"
	"github.com/gin-gonic/gin"
	"github.com/watson-developer-cloud/go-sdk/v2/assistantv2"
)

const (
	AssistantID = "1eddd8fa-59b6-454c-99c0-62f5789cf816"
	ApiKey      = "lveej99H9_azkg2S4Jk7JIoOH89-1FBck4oHdPTippsE"
	ServiceUrl  = "https://api.eu-gb.assistant.watson.cloud.ibm.com/instances/9a057952-b757-4c75-9498-30b2ff47753b"
)

func main() {
	assistant := setupWatson()
	testMessage(assistant)
	handleRoutes(assistant)
}

func setupWatson() *assistantv2.AssistantV2 {
	authenticator := &core.IamAuthenticator{
		ApiKey: "lveej99H9_azkg2S4Jk7JIoOH89-1FBck4oHdPTippsE",
	}
	version := "2020-09-24"
	options := &assistantv2.AssistantV2Options{
		Version:       &version,
		Authenticator: authenticator,
	}

	assistant, assistantErr := assistantv2.NewAssistantV2(options)

	if assistantErr != nil {
		panic(assistantErr)
	}

	assistant.SetServiceURL(ServiceUrl)

	return assistant
}

func handleRoutes(assistant *assistantv2.AssistantV2) {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":5000")
}

func testMessage(assistant *assistantv2.AssistantV2) {

	sessionID := createSession(assistant)
	result, _, responseErr := assistant.Message(
		&assistantv2.MessageOptions{
			AssistantID: core.StringPtr(AssistantID),
			SessionID:   sessionID,
			Input: &assistantv2.MessageInput{
				MessageType: core.StringPtr("text"),
				Text:        core.StringPtr("Fetch last 10 payments"),
			},
		},
	)
	if responseErr != nil {
		panic(responseErr)
	}
	b, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(b))
}

func createSession(assistant *assistantv2.AssistantV2) *string {
	result, _, responseErr := assistant.CreateSession(assistant.
		NewCreateSessionOptions(AssistantID))
	if responseErr != nil {
		panic(responseErr)
	}
	// b, _ := json.MarshalIndent(result, "", "  ")
	// fmt.Println(string(b))
	// fmt.Println(response)

	return result.SessionID
}
