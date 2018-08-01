package model

import "time"

type Login struct {
	Id            int64     `json:"id"`
	UserId        int64     `json:"user_id"`
	ApplicationId int64     `json:"application_id"`
	Token         string    `json:"token"`
	CreatedAt     time.Time `json:"created_at" xorm:"created"`
	Status        string    `json:"status"`
}
