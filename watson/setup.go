// package watson

// import (
// 	"github.com/IBM/go-sdk-core/core"
// 	"github.com/watson-developer-cloud/go-sdk/v2/assistantv2"
// )

// func SetupWatson() {
// 	authenticator := &core.IamAuthenticator{
// 		ApiKey: "{apikey}",
// 	}

// 	options := &assistantv2.AssistantV2Options{
// 		Version:       "{version}",
// 		Authenticator: authenticator,
// 	}

// 	assistant, assistantErr := assistantv2.NewAssistantV2(options)

// 	if assistantErr != nil {
// 		panic(assistantErr)
// 	}

// 	assistant.SetServiceURL("{url}")
// }
