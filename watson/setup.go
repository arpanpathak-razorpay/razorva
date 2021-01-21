/***
 *  Watson Wrapper package for interacting with IBM Watson Assistant
***/
package watson

import (
	"github.com/IBM/go-sdk-core/core"
	"github.com/watson-developer-cloud/go-sdk/v2/assistantv2"
)

// Hardcoded constants =P
const (
	AssistantID = "1eddd8fa-59b6-454c-99c0-62f5789cf816"
	ApiKey      = "lveej99H9_azkg2S4Jk7JIoOH89-1FBck4oHdPTippsE"
	ServiceUrl  = "https://api.eu-gb.assistant.watson.cloud.ibm.com/instances/9a057952-b757-4c75-9498-30b2ff47753b"
)

// SetupWatson Setup and get an instance of IBM Watson
func SetupWatson() *assistantv2.AssistantV2 {
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
