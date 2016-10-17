package types

// Message is...
type Message struct {
	MessageType string                 `json:"message_type"`
	Data        map[string]interface{} `json:"data"`
}
