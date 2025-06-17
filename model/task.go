package model

import "time"

// Task represents a single task in the task management system
type Task struct {
	// Id is the unique identifier for the task
	Id string `json:"id"`
	// Description contains the task details
	Description string `json:"description"`
	// Status represents the current state of the task.
	// Possible values: "pending", "in-progress", "dome"
	Status string `json:"status"`
	// CreationTime is when the task was created
	CreationTime time.Time `json:"createdAt"`
	// UpdatingTime is when the task was last modified
	UpdatingTime time.Time `json:"updatedAt"`
}
