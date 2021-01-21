package socket

import (
	"fmt"

	f "github.com/ambelovsky/gosf"
	"github.com/razorva/watson"
	"github.com/watson-developer-cloud/go-sdk/v2/assistantv2"
)

func Init(assistant *assistantv2.AssistantV2) {
	// Listen on an endpoint
	f.Listen("receive", func(client *f.Client, request *f.Request) *f.Message {
		fmt.Println(request.Message)
		return f.NewSuccessMessage(watson.SendMessage(assistant, request.Message.Text, request.Message.GUID))
	})
}

func InitGoSf(assistant *assistantv2.AssistantV2) {
	Init(assistant)
	// Start the server using a basic configuration
	f.Startup(map[string]interface{}{
		"port": 8080})

}
