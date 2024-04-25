package iot_telemetry

import (
	"context"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"iot-telemetry/types"
	"log"
	"time"
)

func IngestWorkflow(ctx workflow.Context, input types.Input) (*types.Output, error) {
	log.Print("Ingesting message")
	options := workflow.ActivityOptions{
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 5,
		},
		StartToCloseTimeout: time.Second * 5,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	ctx2 := context.Background()

	message, err := Parse(ctx2, input.Content)

	if err != nil {
		return nil, err
	}

	telemetry := message.Telemetry
	errors := message.Errors

	err = workflow.ExecuteActivity(ctx, PersistTelemetry, telemetry).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	err = workflow.ExecuteActivity(ctx, PersistErrors, errors).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	return nil, err
}
