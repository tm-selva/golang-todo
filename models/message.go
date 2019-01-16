package model

type Message struct {
	DefaultProps
	AttachementIds []string `json:"attachment_ids"`
	ChatRoom       string   `json:"chat_room_id"`
	DeliveredTo    []string `json:"delivered_to"`
	User           string   `json:"user_id"`
	SeenBy         []string `json:"seen_by"`
	Text           string   `json:"text"`
	Type           string   `json:"type"`
}
