package config

const (
	ActivateReteLimiter = true
	//6 request in 2 mins
	RateLimiterToken = 6  //5 request
	RateLimterTime   = 21 //scs

	DefaultLanguage = "en"

)

var AcceptableLangs = []string{"en", "tr"}