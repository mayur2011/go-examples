package main

import (
	"fmt"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

func initiLogger() {
	//fpath := "C:\\Users\\Mayur_Upadhyay\\dev\\logs\\access_log%Y%m%d%H%M%S"
	path := "/home/mayur/godev/logs/access.log"
	//f, err := os.OpenFile("n", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	// writer, err := rotatelogs.New(
	// 	fmt.Sprintf("%s.%s", path, "%Y-%m-%d.%H:%M:%S"),
	// 	rotatelogs.WithLinkName("/var/log/service/current"),
	// 	rotatelogs.WithMaxAge(time.Second*10),
	// 	rotatelogs.WithRotationTime(time.Second*1),
	// )
	//

	rl, err := rotatelogs.New(
		fmt.Sprintf("%s%s", path, "%Y%m%d%H%M%S"),
		//rotatelogs.WithLinkName("/home/mayur/godev/logs/current"),
		rotatelogs.WithLinkName(path),
		rotatelogs.WithRotationTime(time.Second*5),
	)

	if err != nil {
		fmt.Println("ERROR -\n", err)
		//log.Fatalf("Failed to Initialize Log File %s", err)
		panic(err)
	}
	fmt.Println("No Error with log writer")
	//log.SetOutput(writer)
	log.SetOutput(rl)
	fmt.Println("DONE with log.SetOutput	")

}

func main() {
	initiLogger()
	for i := 0; i <= 100; i++ {
		fmt.Println("i =", i)
		log.Info("PRINTING i = ", i)
		if i%5 == 0 {
			log.Info("GOING TO SLEEP For 1 Sec")
			time.Sleep(time.Second)
		}
	}
}
