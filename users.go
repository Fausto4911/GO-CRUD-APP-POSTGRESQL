package main

import "time"

type Users struct {
	Id       uint32
	Name     string
	Password string
	Avatar   string
	CreateAt time.Time
}

type UsersRepository interface {
	Store(users Users) error
	GetAll() ([]Users, error)
	GetById(id int) (Users, error)
	Delete(id int) error
	Update(users Users) error
}
