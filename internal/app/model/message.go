package model

import "time"

type Message struct {
	ID      string    `db:"id"`
	Message string    `db:"message"`
	Created time.Time `db:"created_at"`
	Updated time.Time `db:"updated_at"`
}
