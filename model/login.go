package model

import "time"

type Login struct {
	Id            int64     `json:"id"`
	UserId        int64     `json:"user_id"`
	ApplicationId int64     `json:"application_id"`
	Token         string    `json:"token"`
	Datetime      time.Time `json:"datetime" xorm:"created"`
	Status        string    `json:"status"`
}
