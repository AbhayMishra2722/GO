package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("MyTime")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2020, time.August, 23, 10, 10, 10, 10, time.Local)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))
}
