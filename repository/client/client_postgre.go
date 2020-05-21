package client

import (
	"context"
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"go-crud-app-postgresql/model"
	"time"
)

type IClientRepository struct {
	Ctx context.Context
}

func (rep IClientRepository) Store(client model.Client, db *sql.DB) error {
	q := `INSERT INTO client (email, password, nick_name, create_at)
    VALUES($1, $2, $3, $4)`

	defer db.Close()
	st, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer st.Close()
	r, err := st.Exec(GetNUllString(client.Email),
		GetNUllString(client.Password),
		GetNUllString(client.NickName),
		GetNUllTime(client.CreateAt),
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

func (rep IClientRepository) GetAll(db *sql.DB) ([]model.Client, error) {
	q := `SELECT id, email, password, nick_name, create_at, update_at
         FROM client`
	defer db.Close()
	email := sql.NullString{}
	password := sql.NullString{}
	nick_name := sql.NullString{}
	create_at := pq.NullTime{}
	update_at := pq.NullTime{}

	rows, err := db.Query(q)
	if err != nil {
		return []model.Client{}, err
	}
	defer rows.Close()
	al := []model.Client{}

	for rows.Next() {
		a := model.Client{}
		err := rows.Scan(&a.Id,
			&email,
			&password,
			&nick_name,
			&create_at,
			&update_at,
		)
		if err != nil {
			return []model.Client{}, err
		}
		a.Email = email.String
		a.Password = password.String
		a.NickName = nick_name.String
		a.CreateAt = create_at.Time
		a.UpdateAt = update_at.Time
		al = append(al, a)
	}
	return al, nil
}

func (rep IClientRepository) GetById(id int, db *sql.DB) (model.Client, error) {
	q := `SELECT id, email, password, nick_name, create_at, update_at
         FROM client WHERE id = $1`
	defer db.Close()
	email := sql.NullString{}
	password := sql.NullString{}
	nick_name := sql.NullString{}
	create_at := pq.NullTime{}
	update_at := pq.NullTime{}
	m := model.Client{}
	err := db.QueryRowContext(rep.Ctx, q, id).Scan(
		&m.Id,
		&email,
		&password,
		&nick_name,
		&create_at,
		&update_at,
		)
	if err != nil {
		return m, err
	}
	m.Email = email.String
	m.Password = password.String
	m.NickName = nick_name.String
	m.CreateAt = create_at.Time
	m.UpdateAt = update_at.Time

	return m, nil
}

func (rep IClientRepository) Delete(id int, db *sql.DB) error {
	q := `DELETE FROM client WHERE id = $1`
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

func (rep IClientRepository) Update(client model.Client, db *sql.DB) error {
	q := `UPDATE client
          SET email = $1, password = $2, nick_name = $3, update_at = $4
          WHERE id = $5`
	defer db.Close()
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(
		GetNUllString(client.Email),
		GetNUllString(client.Password),
		GetNUllString(client.NickName),
		GetNUllTime(client.UpdateAt),
		client.Id,
		)
	if err != nil {
		return err
	}

	rows, err := r.RowsAffected()
	if rows != 1 {
		return errors.New("Error: No rows affected")
	}
	return nil
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
