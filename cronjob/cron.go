package cronjob

import (
	"log"
	"os"

	"github.com/robfig/cron/v3"
)

var globleCronJob *cron.Cron

// InitCronJob - Initializes the global cron job if it's not already initialized.
func InitCronJob() {
	if globleCronJob != nil {
		return
	}

	// Create a new cron job instance with seconds precision and verbose logging.
	globleCronJob = cron.New(cron.WithSeconds(), cron.WithChain(cron.SkipIfStillRunning(cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags)))))
}

// Add a scheduled function to the cron job.
func AddFunc(spec, name string, cmd func()) (cron.EntryID, error) {
	return globleCronJob.AddFunc(spec, cmd)
}

// Remove a cron job by its EntryID.
func Remove(id cron.EntryID) {
	globleCronJob.Remove(id)
}

// Execute the cron jobs.
func Run() {
	globleCronJob.Run()
}

// Start the cron job.
func Start() {
	globleCronJob.Start()
}

// Stop the cron job.
func Stop() {
	globleCronJob.Stop()
}
