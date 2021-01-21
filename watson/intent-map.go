// Role of this class is to hold all intents and api end points for that intent
package watson

// TODO create generic interface for API Method body and Response format
// type MethodBody interface {

// }
type RazorpayAPI struct {
	Method string `json:"method"` // Type of API method such as GET,POST,PUT,DELETE
	Url    string `json:"url"`
	Body   string `json:"body"`
}

const ()

var IntentMap map[string]string = map[string]string{
	LastFewTransactions:    "https://zcp1h9zwwf.execute-api.ap-south-1.amazonaws.com/lab/payments?merchantId=",
	AskingForPaymentStatus: "https://zcp1h9zwwf.execute-api.ap-south-1.amazonaws.com/lab/payment-status?paymentId=",
	MerchantSuccessRate:    "https://zcp1h9zwwf.execute-api.ap-south-1.amazonaws.com/lab/successrate?merchantId=",
}
