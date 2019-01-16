package model

type User struct {
	DefaultProps
	About     string `json:"about"`
	AvatarURL string `json:"avatar_url"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}
