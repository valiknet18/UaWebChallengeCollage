package twitterApi

var (
	CounsumerApiKey string
	CounsumerSecret string
)

func Init(apiKey, secret string) {
	CounsumerApiKey = apiKey
	CounsumerSecret = secret
}
