package config

import "time"

const (
	ActivateReteLimiter = true

	RateLimiterToken = 6 //5 request
	RateLimterTime   = 5 //scs
	DefaultLanguage  = "en"

	AccessTokenExpiryTime  = time.Hour * 24     //24 hours
	RefreshTokenExpiryTime = time.Hour * 24 * 7 //7 days
)

var AcceptableLangs = []string{"en", "tr"}
var AllowedHost = []string{"localhost:8080"}

var SSLHost = ""

// Request Permissions
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
		"write":  true,
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

type Envs struct {
	APP_MODE              string `mapstructure:"APP_MODE"`
	DBName                string `mapstructure:"DB_NAME"`
	DBUri                 string `mapstructure:"DB_URI"`
	DB_PASSWORD           string `mapstructure:"DB_PASSWORD"`
	DB_USER               string `mapstructure:"DB_USER"`
	GIN_SERVER_PORT       int    `mapstructure:"GIN_SERVER_PORT"`
	BASE_URL              string `mapstructure:"BASE_URL"`
	STORAGE               string `mapstructure:"STORAGE"`
	SESSION_KEY_NAME      string `mapstructure:"SESSION_KEY_NAME"`
	REDIS_DRIVER          string `mapstructure:"REDIS_DRIVER"`
	AUTH_TOKEN_SECRET_KEY string `mapstructure:"AUTH_TOKEN_SECRET_KEY"`
}
