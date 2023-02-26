package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.WithValue(context.Background(), 1, 2)
	go work(ctx)
	fmt.Scanln()
	//cancel()
	fmt.Scanln()
	//ctx := context.TODO()
}
func work(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Work cancelled")
		return
	case <-time.NewTimer(time.Second * 3).C:
		fmt.Println("work", ctx.Value(1), " done")
	}
}
