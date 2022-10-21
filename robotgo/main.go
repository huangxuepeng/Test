package main

import (
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	for {

		// 将鼠标移动到屏幕 x:800 y:400 的位置（闪现到指定位置）
		robotgo.MoveMouse(800, 400)

		robotgo.ScrollMouse(2, `down`)
		time.Sleep(time.Second * 2)
	}

}
