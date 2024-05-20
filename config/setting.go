package config

const (
	ActivateReteLimiter = true
	//6 request in 2 mins
	RateLimiterToken = 6  //5 request
	RateLimterTime   = 21 //scs

	DefaultLanguage = "en"
)

var AcceptableLangs = []string{"en", "tr"}

var DefinedPermissions = map[string]map[string]bool{
	"admin": {
		"write":  true,
		"read":   true,
		"delete": true,
		"update": true,
	},

	"moderator": {
		"write":  true,
		"read":   true,
		"delete": false,
		"update": true,
	},

	"user": {
		"read":   true,
		"write":  false,
		"delete": false,
		"update": false,
	},
}

var PermissionLookUp = map[string]string{
	"get":    "read",
	"post":   "write",
	"delete": "delete",
	"put":    "update",
}
