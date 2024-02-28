package loader

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Loader struct{}

type Envs struct {
	APP_MODE string `mapstructure:"APP_MODE"`

	DBName string `mapstructure:"DB_NAME"`
	DBUri  string `mapstructure:"DB_URI"`

	GIN_SERVER_PORT int `mapstructure:"GIN_SERVER_PORT"`
}

// using Viper to load .env files
func (lDr *Loader) LoadEnvironmetns() (env Envs) {

	fmt.Println("------------ {Loading Environmets} ------------")
	godotenv.Load()
	mode := os.Getenv("APP_MODE")
	envFileName := "dev"

	if mode == "production" {
		envFileName = "prod"
	}
	viper.AddConfigPath("config/")
	viper.SetConfigName(envFileName)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	fmt.Println(err)

	viper.Unmarshal(&env)

	fmt.Println(&env)

	return
}
