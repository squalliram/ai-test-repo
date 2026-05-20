package main

import (
	"fmt"
	"os"
	"time"

	"github.com/splitio/go-client/v6/splitio/client"
	"github.com/splitio/go-client/v6/splitio/conf"
)

func main() {

	// Original example with Split.io API key
	apiKey := os.Getenv("SPLIT_API_KEY")
	if apiKey == "" {
		fmt.Println("Please set SPLIT_API_KEY environment variable")
		os.Exit(1)
	}

	// Create the Split Factory
	cfg := conf.Default()
	factory, err := client.NewSplitFactory(apiKey, cfg)
	if err != nil {
		fmt.Printf("SDK init error: %v\n", err)
		os.Exit(1)
	}

	// Get the Split client
	split := Wrapper{factory.Client()}

	// Wait until the SDK is ready
	if err := split.fmeClient.BlockUntilReady(10); err != nil {
		fmt.Printf("SDK not ready: %v\n", err)
		os.Exit(1)
	}

	// Define the Feature Flag name and user key
	userKey := "user123"

	counter := 0

	// Get the treatment
	treatment := split.Evaluate(userKey, featureFlagName)

	// Check the treatment and show different messages
	switch treatment {
	case "on":
		fmt.Println("Feature flag is ON: Hello, Split.io World!", "printing from switch statement")
		counter += 1
	case "off":
		fmt.Println("Feature flag is OFF: Hello, Regular World!")
		counter -= 1
	default:
		fmt.Printf("Feature flag returned treatment: %s\n", treatment)
	}

	if treatment == "on" {
		fmt.Println("Feature flag is ON: Hello, Split.io World!", "printing from if statement")
	} else if treatment == "off" {
		fmt.Println("Feature flag is OFF: Hello, Regular World!")
	} else {
		fmt.Printf("Feature flag returned treatment: %s\n", treatment)
	}

	// Print the treatment
	fmt.Printf("Treatment: %s\n", treatment)

	// Print the counter
	fmt.Printf("Counter: %d\n", counter)

	// Sleep for a moment to ensure all impressions are sent
	time.Sleep(1 * time.Second)

	// Gracefully shutdown the Split client
	split.fmeClient.Destroy()
}
