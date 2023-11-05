package cronjob

import (
	"log"
	"os"

	"github.com/robfig/cron/v3"
)

var globleCronJob *cron.Cron

func InitCronJob() {
	if globleCronJob != nil {
		return
	}

	globleCronJob = cron.New(cron.WithSeconds(), cron.WithChain(cron.SkipIfStillRunning(cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags)))))
}

func AddFunc(spec, name string, cmd func()) (cron.EntryID, error) {
	return globleCronJob.AddFunc(spec, cmd)
}

func Remove(id cron.EntryID) {
	globleCronJob.Remove(id)
}

func Run() {
	globleCronJob.Run()
}

func Start() {
	globleCronJob.Start()
}

func Stop() {
	globleCronJob.Stop()
}
