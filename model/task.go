package model

import "time"

type Task struct {
	Id           string    `json:"id"`
	Description  string    `json:"description"`
	Status       string    `json:"status"`
	CreationTime time.Time `json:"createdAt"`
	UpdationTime time.Time `json:"updatedAt"`
}
