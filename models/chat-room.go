package model

type ChatRoom struct {
	DefaultProps
	Members []string `json:"members"`
	Type    string   `json:"type"`
}
