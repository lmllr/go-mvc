package models

import (
	"errors"
	"time"

	"github.com/lmllr/go-mvc/db"
)

type Message struct {
	Id        int64
	Name      string
	Message   string
	CreatedAt time.Time
}

// Get all messages
func Messages() (msgs []Message, err error) {
	rows, err := db.Db.Query("SELECT * FROM messages ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		msg := Message{}
		if err = rows.Scan(&msg.Id, &msg.Name, &msg.Message, &msg.CreatedAt); err != nil {
			return
		}
		msgs = append(msgs, msg)
	}
	rows.Close()
	return
}

// Create a message
func (msg *Message) Create() (err error) {
	// validate form values
	if msg.Name == "" || msg.Message == "" {
		errors.New("this error just triggers another error")
		return
	}

	statement := `INSERT INTO messages
(name, message, created_at) VALUES ($1, $2, $3)
RETURNING id, name, message, created_at`

	stmt, err := db.Db.Prepare(statement)
	if err != nil {
		return
	}

	// use QueryRow to return a row and scan the returned id into the User struct
	err = stmt.QueryRow(msg.Name, msg.Message, time.Now()).Scan(&msg.Id, &msg.Name, &msg.Message, &msg.CreatedAt)
	return
}

// Show a single message
func (msg *Message) Show() (err error) {
	statement := `SELECT id, name, message, created_at
FROM messages WHERE id=$1`

	stmt, err := db.Db.Prepare(statement)
	if err != nil {
		return
	}

	err = stmt.QueryRow(msg.Id).Scan(&msg.Id, &msg.Name, &msg.Message, &msg.CreatedAt)
	return
}

// Update message
func (msg *Message) Update() (err error) {
	// validate form values
	if msg.Name == "" || msg.Message == "" {
		errors.New("this error just triggers another error")
		return
	}
	statement := "UPDATE messages SET name=$2, message=$3 WHERE id=$1"
	stmt, err := db.Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(msg.Id, msg.Name, msg.Message)
	return
}

// Delete messages
func DeleteAll() (err error) {
	statement := "DELETE FROM messages"
	_, err = db.Db.Exec(statement)
	return
}

// Delete message
func (msg *Message) Delete() (err error) {
	statement := "DELETE FROM messages WHERE id=$1"
	stmt, err := db.Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(msg.Id)
	return
}
