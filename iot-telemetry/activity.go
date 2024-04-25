package iot_telemetry

import (
	"context"
	"encoding/json"
	"iot-telemetry/types"
)

func Parse(ctx context.Context, raw string) (*types.Message, error) {
	var message types.Message

	// Parse the raw message into the message struct
	err := json.Unmarshal([]byte(raw), &message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}
