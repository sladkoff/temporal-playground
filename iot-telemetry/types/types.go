package types

type Input struct {
	DeviceID    string `json:"device_id"`
	ReceiveTime string `json:"receive_time"`
	Content     string `json:"content"`
}

type Output struct {
	AverageTelemetry float64 `json:"average_telemetry"`
}

type Message struct {
	Telemetry Telemetry `json:"telemetry"`
	Errors    []Error   `json:"errors"`
}

type Telemetry = float64

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
