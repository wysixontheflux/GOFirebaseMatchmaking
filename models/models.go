package models

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Room struct {
	ID      string   `json:"id"`
	Players []string `json:"players"`
}
