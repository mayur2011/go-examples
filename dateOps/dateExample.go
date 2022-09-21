package main

import (
	"fmt"
	"time"
)

func main() {

	//Getting the current system time
	myTimeNow := time.Now()
	//myTimeZone, _ := myTimeNow.Zone()
	fmt.Println("My Time :", myTimeNow) // local time
	setUTCLocation, _ := time.LoadLocation("UTC")
	fmt.Printf("UTC time: %v\n", time.Now().UTC())

	cdtTime := myTimeNow.In(setUTCLocation).Add(-time.Hour * 5)
	fmt.Printf("CDT Time: %v\n", cdtTime)
	fmt.Println(cdtTime.Year(), "-", cdtTime.YearDay(), "-", cdtTime.Day())

	//fmt.Println(time.Now().UTC())
	fmt.Println(time.Now().UTC().Format("2006-01-02T15:04:05.999"))

	//fmt.Printf("Current-time: %v \nType:%T\n", myTimeNow, myTimeNow)
	//fmt.Println("\nTime String:", myTimeNow.String())

	/*
		//Formatting time to a string
		fmt.Println()
		//fmt.Println(myTimeNow.Format("2 Jan 06 03:04PM"))
		//fmt.Println(myTimeNow.Format("2006-01-02T15:04:05.999"))
		//fmt.Println(myTimeNow.Format(time.RFC3339))
		fmt.Println(myTimeNow.Format(time.RFC3339Nano))

		oneHourPastTime := myTimeNow.Add(-time.Hour * 1)
		fmt.Println("One hour past time:\n", oneHourPastTime)

		fmt.Println()
		inputTm := myTimeNow.Format(time.RFC3339Nano)
		//inputTm := myTimeNow.Format("2 Jan 06 03:04PM")
		parseTime, err := time.Parse(time.RFC3339Nano, inputTm)
		if err != nil {
			log.Println("Error while parsing time..", err.Error())
		}
		fmt.Println("parsedTime\n", parseTime)
	*/
}
