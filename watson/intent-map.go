// Role of this class is to hold all intents and api end points for that intent
package watson

// TODO create generic interface for API Method body and Response format
// type MethodBody interface {

// }
type RazorpayAPI struct {
	Method string `json:"method"` // Type of API method such as GET,POST,PUT,DELETE
	Url string `json:"url"`
}

var IntentMap[String] string = {
	"LastFewTransactions": "https://dummy"
}