package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Tiempo terminado")
			default:
				fmt.Println("Trabajando")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(2 * time.Second)

	cancel()

	time.Sleep(1 * time.Second)

}
