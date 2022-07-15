package entity

type Message struct {
	Id          string `json:"id"`
	SenderId    string `json:"sender_id"`
	RecipientId string `json:"receiver_id"`
	Message     string `json:"message"`
}
