package iot_telemetry

import (
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"iot-telemetry/types"
	"time"
)

func IngestWorkflow(ctx workflow.Context) (*types.Output, error) {

	logger := workflow.GetLogger(ctx)
	logger.Info("Ingesting message")

	options := workflow.ActivityOptions{
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 5,
		},
		StartToCloseTimeout: time.Second * 5,
	}

	lastTenTelemetryValues := make([]float64, 0)
	avg := 0.0

	ctx = workflow.WithActivityOptions(ctx, options)

	err := workflow.SetQueryHandler(ctx, "average", func(input []byte) (float64, error) {
		return avg, nil
	})

	if err != nil {
		logger.Info("SetQueryHandler failed: " + err.Error())
		return nil, err
	}

	// while true
	for {
		var rawMessage types.Input

		// get signal channel
		workflow.GetSignalChannel(ctx, "message-signal").Receive(ctx, &rawMessage)

		message, err := Parse(ctx, rawMessage.Content)

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

		lastTenTelemetryValues = append(lastTenTelemetryValues, telemetry)
		if len(lastTenTelemetryValues) > 10 {
			lastTenTelemetryValues = lastTenTelemetryValues[1:]
		}

		avg = 0
		for _, v := range lastTenTelemetryValues {
			avg += v
		}
		avg /= float64(len(lastTenTelemetryValues))

	}

	return nil, nil
}
