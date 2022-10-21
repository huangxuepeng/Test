package main

import (
	"log"
	"sync"

	"github.com/go-vgo/robotgo"
)

var arr = []string{"大傻逼",
	"真傻逼", "脑子有病",
	"你是不是有病", "你就是有病", "真傻逼", "根本无法形容"}

func tt() {

}
func main() {
	var wg sync.WaitGroup
	robotgo.MoveMouse(800, 650)
	robotgo.Click()
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	for {
		for i := 0; i < 6; i++ {
			wg.Add(1)
			go func(i int) {
				robotgo.MoveMouse(800, 650)
				robotgo.Click()
				robotgo.TypeStr(arr[i])
			}(i)
			wg.Done()
		}
		wg.Wait()
		// robotgo.KeyPress("enter")
	}
}
