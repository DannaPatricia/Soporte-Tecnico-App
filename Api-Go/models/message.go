package models

type Message struct {
	ID            int    `json:"id"`
	Content       string `json:"content"`
	Sent_at       string `json:"sent_at"`
	SupportCaseId int    `json:"supportCaseId"`
	UserId        int    `json:"userId"`
	Sender        string `json:"sender"`
}
