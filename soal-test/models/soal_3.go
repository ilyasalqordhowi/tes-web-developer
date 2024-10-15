package models


type Message struct {
	Pesan string `json:"pesan" binding:"required"`
}
