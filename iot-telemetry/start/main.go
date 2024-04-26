package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	app "iot-telemetry"
	"iot-telemetry/types"
	"log"
	"math/rand"
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
		TaskQueue: app.DeviceMessagesQueue,
	}

	inputs := []types.Input{}

	if len(args) == 0 {
		inputs = generateInputs(uuid.New().String(), 10)
	} else {
		inputs = append(inputs, types.Input{
			DeviceID:    "device-1",
			ReceiveTime: "2021-01-01T00:00:00Z",
			Content:     args[0],
		})
	}

	// iterate over inputs
	var we client.WorkflowRun
	for _, input := range inputs {
		// Start the Workflow
		we, err = c.SignalWithStartWorkflow(context.Background(), fmt.Sprintf("ingest-%s", input.DeviceID), "message-signal", input, options, app.IngestWorkflow)
		if err != nil {
			log.Fatalln("unable to complete Workflow", err)
		}
	}

	// Get the results
	average, err := c.QueryWorkflow(context.Background(), we.GetID(), we.GetRunID(), "average", nil)
	if err != nil {
		return
	}

	var value float64
	err = average.Get(&value)

	if err != nil {
		log.Fatalln("unable to get average", err)
	}

	printResults(value, we.GetID(), we.GetRunID())
}

func generateInputs(device string, count int) []types.Input {
	var inputs []types.Input
	for i := 0; i < count; i++ {
		inputs = append(inputs, types.Input{
			DeviceID:    device,
			ReceiveTime: "2021-01-01T00:00:00Z",
			Content:     fmt.Sprintf(`{"telemetry": %f, "errors": []}`, rand.Float64()),
		})
	}
	return inputs
}

func printResults(average float64, workflowID, runID string) {
	fmt.Printf("\nWorkflowID: %s RunID: %s\n", workflowID, runID)
	fmt.Printf("\n%f\n\n", average)
}
