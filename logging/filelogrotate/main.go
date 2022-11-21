package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

func initLogger(logDir *string) {
	fPath := ""
	if *logDir != "" {
		fPath = fmt.Sprintf("%vaccess.log", *logDir)
	} else {
		wd, _ := os.Getwd()
		fPath = wd + "-access.log"
	}
	rl, err := rotatelogs.New(
		fmt.Sprintf("%s.%s", fPath, "%Y%m%d%H%M%S"),
		rotatelogs.WithLinkName(fPath),
		rotatelogs.WithRotationTime(time.Second*5),
		rotatelogs.WithMaxAge(7),
	)

	if err != nil {
		fmt.Println("ERROR -\n", err)
		panic(err)
	}
	log.SetOutput(rl)
}

func main() {
	logDir := flag.String("logdir", "", "logging directory")
	flag.Parse()
	initLogger(logDir)
	for i := 0; i <= 30; i++ {
		fmt.Println("i =", i)
		log.Info("PRINTING i = ", i)
		if i%5 == 0 {
			log.Info("GOING TO SLEEP For 1 Sec")
			time.Sleep(time.Second)
		}
	}
}
