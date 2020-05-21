package repository

import (
	"database/sql"
	"go-crud-app-postgresql/model"
)

type ClientRepository interface {
	Store(client model.Client, db *sql.DB) error
	GetAll(db *sql.DB) ([]model.Client, error)
	GetById(id int, db *sql.DB) (model.Client, error)
	Delete(id int, db *sql.DB) error
	Update(users model.Client, db *sql.DB) error
}
