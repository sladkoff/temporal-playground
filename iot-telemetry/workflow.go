package iot_telemetry

import (
	"context"
	"go.temporal.io/sdk/workflow"
	"iot-telemetry/types"
	"time"
)

func IngestWorkflow(ctx workflow.Context, input types.Input) (*types.Output, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	ctx2 := context.Background()

	_, err := Parse(ctx2, input.Content)

	if err != nil {
		return nil, err
	}

	//var result string
	//err := workflow.ExecuteActivity(ctx, ComposeGreeting, "John").Get(ctx, &result)

	return nil, err
}
