package iot_telemetry

import (
	"go.temporal.io/sdk/workflow"
	"time"
)

func GreetingWorkflow(ctx workflow.Context) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	var result string
	err := workflow.ExecuteActivity(ctx, ComposeGreeting, "John").Get(ctx, &result)

	return result, err
}
