package main

import (
	"log"
	"testing"
)

func Test_Sum(t *testing.T) {
	str := sum("1", "999")
	log.Println(str)
}
