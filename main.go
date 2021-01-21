package main

import (
	"github.com/razorva/watson"

	"github.com/gin-gonic/gin"
	socket "github.com/razorva/socket"
	"github.com/watson-developer-cloud/go-sdk/v2/assistantv2"
)

const (
	AssistantID = "1eddd8fa-59b6-454c-99c0-62f5789cf816"
	ApiKey      = "lveej99H9_azkg2S4Jk7JIoOH89-1FBck4oHdPTippsE"
	ServiceUrl  = "https://api.eu-gb.assistant.watson.cloud.ibm.com/instances/9a057952-b757-4c75-9498-30b2ff47753b"
)

func main() {
	assistant := watson.SetupWatson()
	// fmt.Println(watson.SendMessage(assistant,"what is the status of pay_2314314","rzp2333"))
	socket.InitializeAndListen(assistant)
	// handleRoutes(assistant)
}

func handleRoutes(assistant *assistantv2.AssistantV2) {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "RazorBot is here to answer all your",
		})
	})

	r.Run(":5000")
}
