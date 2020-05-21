package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"log"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "POST"
)


type IUsersRepository struct {
	Ctx context.Context
}

//getConnection obtain a DB connection
func GetConnection() (db *sql.DB) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func GetNUllTime(t time.Time) (n pq.NullTime) {
	if t.IsZero() {
		n.Valid = false
	} else {
		n.Valid = true
		n.Time = t
	}
	return
}

func GetNUllString(s string) (n sql.NullString) {
	if s == "" {
		n.Valid = false
	} else {
		n.Valid = true
		n.String = s
	}
	return
}
func GetNUllInt(i int64) (n sql.NullInt64) {
	if i == 0 {
		n.Valid = false
	} else {
		n.Valid = true
		n.Int64 = int64(i)
	}
	return
}

func (rep IUsersRepository) Store(user Users) error {
	q := `INSERT INTO users (name, password, avatar, create_at)
    VALUES($1, $2, $3, $4)`

	db := GetConnection()
	defer db.Close()

	st, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer st.Close()
	r, err := st.Exec(GetNUllString(user.Name),
		GetNUllString(user.Password),
		GetNUllString(user.Avatar),
		GetNUllTime(user.CreateAt),
	)
	if err != nil {
		return err
	}
	rows, err := r.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("No rows affected")
	}
	return nil
}

func (rep IUsersRepository) GetAll() ([]Users, error) {
	q := `SELECT id, name, password, avatar, create_at
         FROM users`
	db := GetConnection()
	defer db.Close()

	na := sql.NullString{}
	pa := sql.NullString{}
	av := sql.NullString{}
	cr := pq.NullTime{}

	rows, err := db.Query(q)
	if err != nil {
		return []Users{}, err
	}
	defer rows.Close()
	al := []Users{}

	for rows.Next() {
		a := Users{}
		err := rows.Scan(&a.Id,
			&na,
			&pa,
			&av,
			&cr,
		)
		if err != nil {
			return []Users{}, err
		}
		a.Name = na.String
		a.Password = pa.String
		a.Avatar = av.String
		a.CreateAt = cr.Time
		al = append(al, a)
	}
	return al, nil
}

func (rep IUsersRepository) GetById(id int) (Users, error) {
	q := `SELECT id, name, password, avatar, create_at
         FROM users WHERE id = $1`
	db := GetConnection()
	defer db.Close()
	na := sql.NullString{}
	pa := sql.NullString{}
	av := sql.NullString{}
	cr := pq.NullTime{}
	m := Users{}

	err := db.QueryRowContext(rep.Ctx, q, id).Scan(
		&m.Id,
		&na,
		&pa,
		&av,
		&cr)
	if err != nil {
		return m, err
	}
	m.Name = na.String
	m.Password = pa.String
	m.Avatar = pa.String
	m.CreateAt = cr.Time

	return m, nil
}

func (rep IUsersRepository) Delete(id int) error {
	q := `DELETE FROM users WHERE id = $1`
	db := GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rows, err := r.RowsAffected()
	if rows != 1 {
		return errors.New("Error: No rows affected")
	}
	return nil
}

func (rep IUsersRepository) Update(users Users) error {
	q := `UPDATE users
          SET name = $1, password = $2, avatar = $3
          WHERE id = $4`

	db := GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(users.Name, users.Password, users.Avatar, users.Id)
	if err != nil {
		return err
	}

	rows, err := r.RowsAffected()
	if rows != 1 {
		return errors.New("Error: No rows affected")
	}
	return nil
}
