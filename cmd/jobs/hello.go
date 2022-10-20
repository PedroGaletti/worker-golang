package jobs

import (
	"sync"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

// Hello: Struct of Hello job
type Hello struct {
	wg        sync.WaitGroup
	Scheduler *gocron.Scheduler
}

// NewHello: Create Hello job
func NewHello(scheduler *gocron.Scheduler) IJob {
	return &Hello{Scheduler: scheduler}
}

// Start: Start the job and execute the task
func (h *Hello) Start() {
	log.Info().Msg("Starting Hello job...")

	h.Scheduler.Cron(env.CronTime).Do(func() {
		h.wg.Add(1)

		defer h.wg.Done()

		log.Info().Msg("Hello World!")

		h.wg.Wait()
	})
}

// Stop: Stop the job when the current task has finished
func (h *Hello) Stop() {
	log.Info().Msg("Stopping Hello job...")

	h.wg.Wait()

	log.Info().Msg("Hello job stopped!")
}
