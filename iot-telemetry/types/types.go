package types

type Input struct {
	DeviceID    string `json:"device_id"`
	ReceiveTime string `json:"receive_time"`
	Content     string `json:"content"`
}

type Output struct {
}

type Message struct {
	Telemetry []Telemetry `json:"telemetry"`
	Errors    []Error     `json:"errors"`
}

type Telemetry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
