package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	Leak()
}

func Leak() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		ch := retChan(ctx)
		for s := range ch {
			fmt.Printf("s: %s\n", s)

			if s == "20" {
				cancel()
				return
			}
		}
	}()
	// 擬似サーバー
	time.Sleep(3 * time.Second)
}

func retChan(ctx context.Context) <-chan string {
	ss := []string{}
	for i := 0; i < 100; i++ {
		ss = append(ss, fmt.Sprintf("%d", i))
	}
	ch := make(chan string, 10)

	// goroutine
	go func() {
		defer close(ch)

		// ここでcancelを伝搬
		for _, s := range ss {
			select {
			case <-ctx.Done():
				fmt.Println("detect cancel")
				return
			default:
			}

			ch <- s
		}
	}()
	return ch
}
