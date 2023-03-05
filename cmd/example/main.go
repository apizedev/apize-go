package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"github.com/apizedev/apize-go"
)

func main() {
	// Create a context for the program
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	defer cancel()

	// Create the Apize client
	client := apize.New("token")

	// Send a request to the API
	res, err := client.Measure(ctx, &apize.MeasureRequest{
		Text: "the quick brown fox jumped over the lazy dog",
	})
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	// Print the results
	fmt.Printf("Words: %d\n", res.Words)
}
