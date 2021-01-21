package watson

import (
	"strings"

	"github.com/razorva/utils"
)

// TODO

const (
	AskingWhyPaymentsAreGettingFailed = "AskingWhyPaymentsAreGettingFailed"
	AskingForPaymentStatus            = "AskingForPaymentStatus"
	LastFewTransactions               = "LastFewTransactions"
	GatewaySuccessRate                = "GatewaySuccessRate"
)

// Process intent and return result
func ProcessIntent(intent string, merchantId string, message string) string {
	switch intent {
	case AskingForPaymentStatus:
		return PaymentStatusQuery(message)
	default:
		return "Pardon human I didn't understand your complex language. I'm still learning!"
	}

}

func PaymentStatusQuery(message string) string {
	// reply := "Here is the details of your payment"

	if strings.Index(message, "pay_") == -1 {
		return "Hey hooman you forgot to give me your payment id"
	}

	if !isPaymentIdValid(message){
		return "Hey! please enter a valid payment id"
	}

	reply, err := utils.SendGetRequest(IntentMap[AskingForPaymentStatus] + "pay_60095e759b9d2b38ef061d5a")

	if err != nil {
		panic(err)
	}
	return reply
}


func isPaymentIdValid(message string) bool {

	lastIndex := strings.LastIndex(message,"pay_")

	if (len(message) - lastIndex) == 23 {
		return true
	}

	return false
}
