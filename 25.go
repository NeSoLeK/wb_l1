package main

//Реализовать собственную функцию sleep.

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {

	<-time.After(duration)
}

func t25() {
	fmt.Println("Start")
	sleep(2 * time.Second)
	fmt.Println("Stop")
}
