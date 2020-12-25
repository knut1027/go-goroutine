package main

import (
	"testing"

	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestLeak(t *testing.T) {
	Leak()
}

// func TestRetChan(t *testing.T) {
// 	ctx, _ := context.WithCancel(context.Background())
// 	retChan(ctx)
// }
