package main

import (
	"context"
	"fmt"
	app "iot-telemetry"
	"iot-telemetry/types"
	"log"
	"os"

	"go.temporal.io/sdk/client"
)

func main() {

	// arguments
	args := os.Args[1:]

	// Create the client object just once per process
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "ingest-workflow",
		TaskQueue: app.DeviceMessagesQueue,
	}

	input := types.Input{
		DeviceID:    "abc",
		ReceiveTime: "2021-01-01T00:00:00Z",
		Content:     args[0],
	}

	// Start the Workflow
	we, err := c.ExecuteWorkflow(context.Background(), options, app.IngestWorkflow, input)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}

	// Get the results
	var greeting string
	err = we.Get(context.Background(), &greeting)
	if err != nil {
		log.Fatalln("unable to get Workflow result", err)
	}

	printResults(greeting, we.GetID(), we.GetRunID())
}

func printResults(greeting string, workflowID, runID string) {
	fmt.Printf("\nWorkflowID: %s RunID: %s\n", workflowID, runID)
	fmt.Printf("\n%s\n\n", greeting)
}
