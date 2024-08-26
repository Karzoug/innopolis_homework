package model

type Message struct {
	Token  string `json:"token"`
	FileID string `json:"file_id"`
	Data   string `json:"data,omitempty"`
}
