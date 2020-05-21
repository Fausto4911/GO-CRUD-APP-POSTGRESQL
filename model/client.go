package model

import "time"

type Client struct {
	Id       uint32
	Email    string
	Password string
	NickName string
	CreateAt time.Time
	UpdateAt time.Time
}
