package auth

type Grants struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	RoomID    string `json:"roomId"`
	RoomAdmin bool   `json:"roomAdmin"`
}
