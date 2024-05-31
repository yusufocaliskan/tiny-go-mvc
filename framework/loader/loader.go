package loader

import (
	"fmt"
	"os"

	"github.com/yusufocaliskan/tiny-go-mvc/config"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Loader struct{}

// using Viper to load .env files
func (lDr *Loader) LoadEnvironments() (env config.Envs) {

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
	if err != nil {
		fmt.Println(err)
	}

	viper.Unmarshal(&env)

	return
}

func (lDr *Loader) LoadFile() {

}
