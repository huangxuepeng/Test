package main

import (
	"log"
	"testing"
)

func Test_numIslands(t *testing.T) {
	arr := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	c := numIslands(arr)
	log.Println(c)
}
