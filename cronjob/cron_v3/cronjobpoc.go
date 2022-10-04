package main

import (
	"fmt"
	"time"

	cron "github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}

func runJob(name string, ts string) {
	message := fmt.Sprintf("%v %v\n", name, ts)
	fmt.Println(message)
}

func printCronEntries(cronEntries []cron.Entry) {
	log.Infof("Cron Info: %+v\n", cronEntries)
}

func runCronJobs(ts string) {
	log.Info("Create new cron")
	c := cron.New()
	//c.AddFunc("*/1 * * * *", func() { log.Info("[Job 1] Every minute job\n") })
	c.AddFunc("@every 3s", func() {
		runJob("[Job 1] Every minute job", ts)
	})

	// Start cron with one scheduled job
	log.Info("Start cron")
	c.Start()
	printCronEntries(c.Entries())
}

func main() {
	startTimeUTC := time.Now().UTC().Add(-time.Minute * 90).Format("2006-01-02T15:04:05:999")
	runCronJobs(startTimeUTC)
	fmt.Scanln()
	//time.Sleep(10 * time.Second)
}
