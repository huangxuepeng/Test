package main

import (
	"testing"

	"go.uber.org/goleak"
)

// func TestGetData(t *testing.T) {
// 	defer goleak.VerifyNone(t)
// 	getData()
// }

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
	// getData()
}
