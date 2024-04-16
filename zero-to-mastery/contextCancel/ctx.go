package main

import (
	"context"
	"fmt"
	"time"
)

func sampleOp(ctx context.Context, msg string, msDelay time.Duration) <-chan string {
	out := make(chan string)

	go func() {
		for {
			select {
			case <-time.After(msDelay * time.Millisecond):
				out <- fmt.Sprintf("%v operation completed", msg)
				return
			case <-ctx.Done():
				out <- fmt.Sprintf("%v aborted", msg)
				return
			}
		}
	}()

	return out
}

func main() {
	ctx := context.Background()
	ctx, cancelCtx := context.WithCancel(ctx)

	webServ := sampleOp(ctx, "webserver", 200)
	microServ := sampleOp(ctx, "microservice", 500)
	dbServ := sampleOp(ctx, "database", 900)

MainLoop:
	for {
		select {
		case val := <-webServ:
			fmt.Println(val)
		case val := <-microServ:
			fmt.Println(val)
			fmt.Println("cancel context")
			cancelCtx()
			break MainLoop
		case val := <-dbServ:
			fmt.Println(val)
		}
	}
	// just to see if anything happens
	fmt.Println(<-dbServ)
}
