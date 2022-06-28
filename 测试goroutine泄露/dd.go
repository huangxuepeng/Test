package main

import (
	"fmt"
	"runtime"
	"time"
)

func getData() {
	ch := make(chan int, 1)
	go func() {
		<-ch
	}()
}

func main() {
	defer func() {
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}()
	getData()
	time.Sleep(time.Second * 5)
}
