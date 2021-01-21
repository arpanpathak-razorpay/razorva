package main

import (
	"fmt"

	"github.com/IBM/go-sdk-core/core"
	"github.com/gin-gonic/gin"
	"github.com/watson-developer-cloud/go-sdk/v2/assistantv2"
)

func main() {
	assistant := setupWatson()
	fmt.Print(assistant.Version)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hi I'm RazorVa",
		})
	})

	r.Run(":5000")
}

func setupWatson() *assistantv2.AssistantV2 {
	authenticator := &core.IamAuthenticator{
		ApiKey: "<api-key>",
	}

	options := &assistantv2.AssistantV2Options{
		Version:       "{version}",
		Authenticator: authenticator,
	}

	assistant, assistantErr := assistantv2.NewAssistantV2(options)

	if assistantErr != nil {
		panic(assistantErr)
	}

	assistant.SetServiceURL("https://api.eu-gb.assistant.watson.cloud.ibm.com/instances/9a057952-b757-4c75-9498-30b2ff47753b")

	return assistant
}
