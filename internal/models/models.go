package models

type Item struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type ItemLog struct {
	ID        int    `json:"id"`
	Action    string `json:"action"`
	OcorredAt string `json:"ocorred_at"`
	User      string `json:"user"`
	UserEmail string `json:"user_email"`
	Device    string `json:"device"`
	IP        string `json:"ip"`
}
