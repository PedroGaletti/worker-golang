package jobs

import (
	"time"
	"worker-golang/cmd/configs"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

var (
	env = configs.Env
)

type IJob interface {
	Start()
	Stop()
}

type Jobs struct {
	Hello IJob
}

// Setting: Setting the cron application
func Setting() *gocron.Scheduler {
	log.Info().Msg("Setting cron...")
	location, _ := time.LoadLocation(env.Localization)
	return gocron.NewScheduler(location)
}

// New: Start the jobs of the application and return then to the system
func New() *Jobs {
	scheduler := Setting() // Call setting function

	// Initiate Hello job
	hello := NewHello(scheduler)
	go hello.Start()

	scheduler.StartAsync() // Start the cron scheduler

	return &Jobs{Hello: NewHello(scheduler)} // return the jobs for the system
}
