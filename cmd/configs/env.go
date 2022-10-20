package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	// Env : export environment variables
	Env = GetConfig()
)

// Config : struct of config. variables
type Config struct {
	CronTime     string
	Localization string
	LogLevel     string
}

// GetConfig : get values of environment variables
func GetConfig() Config {
	var config Config

	if err := godotenv.Load(); err != nil {
		fmt.Println("Runnning the application without a .env file.")
	}

	config.CronTime = os.Getenv("CRON_TIME")
	config.Localization = os.Getenv("LOCALIZATION")
	config.LogLevel = os.Getenv("LOG_LEVEL")

	return config
}
