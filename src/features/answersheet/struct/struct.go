package answersheet_struct

import "time"

type StartTestInput struct {
	JobId   int `json:"job_id"`
	Payload struct {
		UserId    int        `json:"user_id"`
		TestId    int        `json:"test_id"`
		Event     string     `json:"event"`
		CreatedAt *time.Time `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at"`
	} `json:"payload"`
}
