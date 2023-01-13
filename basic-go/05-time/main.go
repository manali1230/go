package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("This is Time !!")

	// current time
	currentTime := time.Now()
	fmt.Println(currentTime)

	fmt.Println(currentTime.Format("01-02-2006 15:04:05 Monday"))
	fmt.Println(currentTime.Format("Today's date is 01-02-2006"))
	fmt.Println(currentTime.Format("Current Time is 15:04:05"))
	fmt.Println(currentTime.Format("Today's day is Monday"))

	// different time
	diffTime := time.Date(1999, time.August, 12, 23, 23, 5, 0, time.UTC)
	year, month, date := diffTime.Date()
	fmt.Println(diffTime)
	fmt.Println(diffTime.Format("01-02-2006 Monday"))
	fmt.Println("year is ", year)
	fmt.Println("month is ", month)
	fmt.Println("date is ", date)
}
