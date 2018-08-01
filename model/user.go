package model

import "time"

type User struct {
	ID          int64
	UserId      int64
	FirstName   string
	LastName    string
	Picture     string
	Gender      string
	Email       string
	Password    string
	Role        string
	LatestLogin time.Time
	LastIp      string
	Connection  string
	Created     time.Time `xorm:"created"`
	Updated     time.Time `xorm:"updated"`
}
