package models

// Counter : Counter DB model
type Counter struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
	Name   string `json:"name"`
}
