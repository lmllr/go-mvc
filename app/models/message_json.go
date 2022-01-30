package models

import (
	"errors"
	"time"

	"github.com/lmllr/go-mvc/db"
)

type MessageJSON struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Message   string    `json:"msg"`
	CreatedAt time.Time `json:"created_at"`
}

func MessagesJSON() (msgs []MessageJSON, err error) {
	rows, err := db.Db.Query("SELECT * FROM messages ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		msg := MessageJSON{}
		if err = rows.Scan(&msg.Id, &msg.Name, &msg.Message, &msg.CreatedAt); err != nil {
			return
		}
		msgs = append(msgs, msg)
	}
	rows.Close()
	return
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

// Show a single message
func (msg *MessageJSON) Show() (err error) {
	statement := `SELECT id, name, message, created_at
FROM messages WHERE id=$1`

	stmt, err := db.Db.Prepare(statement)
	if err != nil {
		return
	}

	err = stmt.QueryRow(msg.Id).Scan(&msg.Id, &msg.Name, &msg.Message, &msg.CreatedAt)
	return
}

// Update
// Delete
