package configs

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Configure: setting zerolog
func Configure() {
	zerolog.TimeFieldFormat = "02/01/2006 15:04:05"

	logLevel, err := zerolog.ParseLevel(Env.LogLevel)
	if err != nil {
		log.Panic().Msg("log level undefined, check .env file")
	}

	zerolog.SetGlobalLevel(logLevel)
}
