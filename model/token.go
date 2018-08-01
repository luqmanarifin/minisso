package model

import "time"

type Token struct {
	Id        int64     `json:"id"`
	Token     string    `json:"token"`
	UserId    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at" xorm:"created"`
	ExpiredAt time.Time `json:"expired_at"`
}
