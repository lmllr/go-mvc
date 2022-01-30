package models

import (
	"time"

	"github.com/lmllr/go-mvc/app/helpers"
	"github.com/lmllr/go-mvc/db"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func (usr *User) Create() (err error) {
	// Postgres does not automatically return the last insert id, because it would be wrong to assume
	// you're always using a sequence.You need to use the RETURNING keyword in your insert to get this
	// information from postgres.
	statement := "INSERT INTO users (uuid, name, email, password, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, created_at"
	stmt, err := db.Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	hash, _ := helpers.HashPassword(usr.Password)

	// use QueryRow to return a row and scan the returned id into the User struct
	err = stmt.QueryRow(helpers.CreateUUID(), usr.Name, usr.Email, hash, time.Now()).Scan(&usr.Id, &usr.Uuid, &usr.CreatedAt)
	return
}
