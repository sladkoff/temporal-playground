package iot_telemetry

import (
	"context"
	"encoding/json"
	"fmt"
	"iot-telemetry/types"
	"log"
	"math/rand"
	"os"
	"time"
)

func Parse(ctx context.Context, raw string) (*types.Message, error) {
	log.Print("Parsing message")
	var message types.Message

	// Parse the raw message into the message struct
	err := json.Unmarshal([]byte(raw), &message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}

func PersistTelemetry(ctx context.Context, telemetry []types.Telemetry) error {
	// Fake a delay between 100ms and 2s
	delay := 100 + rand.Intn(1900)
	time.Sleep(time.Duration(delay) * time.Millisecond)

	// 10% chance of failure
	if rand.Intn(10) == 0 {
		log.Print("Failed to persist telemetry")
		return fmt.Errorf("failed to persist telemetry")
	}

	// serialize telemetry as json
	serialized, err := json.Marshal(telemetry)
	if err != nil {
		return err
	}

	// append to a file
	f, err := os.OpenFile("telemetry.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(serialized); err != nil {
		return err
	}

	return nil
}

func PersistErrors(ctx context.Context, errors []types.Error) error {
	// Fake a delay between 100ms and 2s
	delay := 100 + rand.Intn(1900)
	time.Sleep(time.Duration(delay) * time.Millisecond)

	// 10% chance of failure
	if rand.Intn(10) == 0 {
		log.Print("Failed to persist errors")
		return fmt.Errorf("failed to persist errors")
	}
	return nil
}
