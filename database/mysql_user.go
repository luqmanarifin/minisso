package database

import (
	"log"
	"net/http"
	"time"

	"github.com/luqmanarifin/minisso/model"
	"github.com/tomasen/realip"
)

func (m *Mysql) CreateUser(user model.User) {
	affected, err := m.xorm.Insert(&user)
	if err != nil {
		log.Printf("%d error %s", affected, err)
	}
}

func (m *Mysql) IsEmailExist(email string) bool {
	user := new(model.User)
	total, err := m.xorm.Where("email=?", email).Count(user)
	if err != nil {
		log.Printf("Error when query whether email exist or not")
	}
	return total > 0
}

func (m *Mysql) FindUserById(id int64) model.User {
	var user = model.User{Id: id}
	has, _ := m.xorm.Get(&user)
	if has {
		return user
	} else {
		return model.User{}
	}
}

func (m *Mysql) FindUserByEmail(email string) model.User {
	var user = model.User{Email: email}
	has, _ := m.xorm.Get(&user)
	if has {
		return user
	} else {
		return model.User{}
	}
}

func (m *Mysql) TouchUserLogin(r *http.Request, user model.User) model.User {
	user.LatestLogin = time.Now()
	user.LastIp = realip.FromRequest(r)
	m.xorm.Id(user.Id).Update(user)
	return user
}
