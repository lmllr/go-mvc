package models

import (
	"errors"
	"time"

	"github.com/lmllr/go-mvc/db"
)

type MessageJSON struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Message   string    `json:"msg"`
	CreatedAt time.Time `json:"created_at"`
}

func (msg *MessageJSON) CreateFromJSON() (err error) {
	// validate form values
	if msg.Name == "" || msg.Message == "" {
		errors.New("this error just triggers another error")
		return
	}

	statement := `INSERT INTO messages (name, message, created_at) VALUES ($1, $2, $3) RETURNING id, name, message, created_at`

	stmt, err := db.Db.Prepare(statement)
	if err != nil {
		return
	}

	// use QueryRow to return a row and scan the returned id into the User struct
	err = stmt.QueryRow(msg.Name, msg.Message, time.Now()).Scan(&msg.Id, &msg.Name, &msg.Message, &msg.CreatedAt)
	return
}
