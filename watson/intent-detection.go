package watson

import (
	"fmt"
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
func ProcessIntent(intent string, merchantId string, message string, pageSize int) string {
	switch intent {
	case AskingForPaymentStatus:
		return PaymentStatusQuery(message)
	case LastFewTransactions:
		return FetchLastPayments(message, merchantId, pageSize)
	default:
		return "Pardon human I didn't understand your complex language. I'm still learning!"
	}

}

func PaymentStatusQuery(message string) string {
	// reply := "Here is the details of your payment"

	if strings.Index(message, "pay_") == -1 {
		return "Hey hooman you forgot to give me your payment id"
	}

	// if len(message) != 28 {
	// 	return "Hey! please enter a valid payment id"
	// }

	reply, err := utils.SendGetRequest(IntentMap[AskingForPaymentStatus] + "pay_60095e759b9d2b38ef061d5a")

	if err != nil {
		panic(err)
	}
	return reply
}

func FetchLastPayments(message string, merchantId string, pageSize int) string {

	reply, err := utils.SendGetRequest(IntentMap[LastFewTransactions] + merchantId)

	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Here are those last %d payments \n %s", pageSize, reply)
}

// func FetchSuccessRate(gateway string) string {

// }
