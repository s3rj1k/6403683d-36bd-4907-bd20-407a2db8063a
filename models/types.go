package models

import "time"

type Sign struct {
	Questions []string `json:"questions"`
	Answers   []string `json:"answers"`
}

type Verify struct {
	UserID    string `json:"user_id"`
	Signature string `json:"signature"`
}

type Data struct {
	UserID    string    `json:"user_id"`
	Questions []string  `json:"questions"`
	Answers   []string  `json:"answers"`
	Signature string    `json:"signature"`
	Timestamp time.Time `json:"timestamp"`
}
