package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func initLogger(logDir *string) {
	fPath := ""
	if *logDir != "" {
		fmt.Println("IF IF IF")
		fPath = fmt.Sprintf("%vaccess.log", *logDir)
	} else {
		wd, _ := os.Getwd()
		fPath = wd + "-access.log"
	}
	lumberjackLogger := &lumberjack.Logger{
		Filename:   filepath.ToSlash(fPath),
		MaxSize:    1, // MB
		MaxBackups: 10,
		MaxAge:     10, // days
	}

	// Fork writing into two outputs
	multiWriter := io.MultiWriter(os.Stderr, lumberjackLogger)

	logFormatter := new(log.TextFormatter)
	logFormatter.TimestampFormat = time.RFC1123Z // or RFC3339
	logFormatter.FullTimestamp = true

	log.SetFormatter(logFormatter)
	log.SetLevel(log.InfoLevel)
	log.SetOutput(multiWriter)
}

func main() {
	logDir := flag.String("logdir", "", "logging directory")
	flag.Parse()
	initLogger(logDir)
	for i := 0; i <= 1000; i++ {
		fmt.Println("i =", i)
		log.Info("PRINTING i = ", i)
		if i%5 == 0 {
			log.Info("GOING TO SLEEP For 1 Sec")
			//time.Sleep(time.Second)
		}
	}
}
