package iot_telemetry

import (
	"context"
	"encoding/json"
	"fmt"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
	"iot-telemetry/types"
	"math/rand"
	"os"
	"time"
)

func Parse(ctx workflow.Context, raw string) (*types.Message, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Parsing message")
	var message types.Message

	// Parse the raw message into the message struct
	err := json.Unmarshal([]byte(raw), &message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}

func PersistTelemetry(ctx context.Context, telemetry types.Telemetry) error {
	logger := activity.GetLogger(ctx)

	// Fake a delay between 100ms and 2s
	delay := 100 + rand.Intn(1900)
	time.Sleep(time.Duration(delay) * time.Millisecond)

	// 10% chance of failure
	if rand.Intn(10) == 0 {
		logger.Info("Failed to persist telemetry")
		return fmt.Errorf("failed to persist telemetry")
	}

	serialized, err := json.Marshal(telemetry)
	if err != nil {
		return err
	}

	err = appendToFile("telemetry.log", serialized)
	if err != nil {
		return err
	}

	logger.Info("Persisted telemetry")

	return nil
}

func PersistErrors(ctx context.Context, errors []types.Error) error {
	logger := activity.GetLogger(ctx)
	// Fake a delay between 100ms and 2s
	delay := 100 + rand.Intn(1900)
	time.Sleep(time.Duration(delay) * time.Millisecond)

	// 10% chance of failure
	if rand.Intn(10) == 0 {
		logger.Info("Failed to persist errors")
		return fmt.Errorf("failed to persist errors")
	}

	serialized, err := json.Marshal(errors)
	if err != nil {
		return err
	}

	err = appendToFile("errors.log", serialized)
	if err != nil {
		return err
	}

	logger.Info("Persisted errors")

	return nil
}

func appendToFile(fileName string, serialized []byte) error {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	newLine := "\n"

	serialized = append(serialized, newLine...)
	if _, err := f.Write(serialized); err != nil {
		return err
	}
	return nil
}
