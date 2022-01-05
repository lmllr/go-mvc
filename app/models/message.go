package models

import (
	"errors"
	"time"

	"github.com/lmllr/go-mvc/db"
)

type Message struct {
	Id        int
	Name      string
	Message   string
	CreatedAt time.Time
}

func Messages() (msgs []Message, err error) {
	rows, err := db.Db.Query("SELECT id, name, created_at FROM messages")
	if err != nil {
		return
	}
	for rows.Next() {
		msg := Message{}
		if err = rows.Scan(&msg.Id, &msg.Name, &msg.CreatedAt); err != nil {
			return
		}
		msgs = append(msgs, msg)
	}
	rows.Close()
	return
}

func (msg *Message) Create() (err error) {
	// validate form values
	if msg.Name == "" || msg.Message == "" {
		return errors.New("this error just triggers another error")
	}

	statement := `INSERT INTO messages
(name, message, created_at) values ($1, $2, $3)
returning id, name, message, created_at`

	stmt, err := db.Db.Prepare(statement)
	if err != nil {
		return
	}

	// use QueryRow to return a row and scan the returned id into the User struct
	err = stmt.QueryRow(msg.Name, msg.Message, time.Now()).Scan(&msg.Id, &msg.Name, &msg.Message, &msg.CreatedAt)
	return
}
