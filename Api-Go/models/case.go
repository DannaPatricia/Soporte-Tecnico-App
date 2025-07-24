package models

import "time"

type SupportCase struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	// Para modificar la hora de modificacion
	CreatedAt   time.Time `json:"createdAt"`
	UserID      int `json:"userId"`
	CategoryID  int `json:"categoryId"`
}
