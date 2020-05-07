package payloads

import "time"

type MessageResponse struct {
	ID      string    `json:"id"`
	Message string    `json:"message"`
	Created time.Time `json:"created_at"`
	Updated time.Time `json:"updated_at"`
}
