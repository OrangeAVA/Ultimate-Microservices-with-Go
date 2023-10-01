package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	childContext := context.WithValue(ctx, "apiKey", 123456)
	printAPIKey(childContext)

	contextWithCancel, cancel := context.WithCancel(ctx)
	cancel()
	fmt.Println("What happened ?", contextWithCancel.Err()) // What happened ? context canceled

	contextWithTimeout, _ := context.WithTimeout(ctx, 30*time.Second)
	fmt.Println("What happened ?", contextWithTimeout.Err()) // What happened ? <nil>
	time.Sleep(35 * time.Second)
	fmt.Println("What happened ?", contextWithTimeout.Err()) // What happened ? context deadline exceeded

}

func printAPIKey(ctx context.Context) {
	apiKey := ctx.Value("apiKey")
	fmt.Println("API Key:", apiKey) // API Key: 123456
}
