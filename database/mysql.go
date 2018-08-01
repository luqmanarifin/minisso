package database

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/tomasen/realip"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/luqmanarifin/minisso/model"
)

// Option holds all necessary options for database.
type MysqlOption struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
	Charset  string
}

type Mysql struct {
	xorm *xorm.Engine
}

func NewXorm(opt MysqlOption) (Mysql, error) {
	command := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		opt.User,
		opt.Password,
		opt.Host,
		opt.Port,
		opt.Database,
		opt.Charset,
	)
	log.Println(command)
	db, err := xorm.NewEngine("mysql", command)

	if err != nil {
		return Mysql{}, err
	}

	// syncing table
	db.Sync2(new(model.Application))
	db.Sync2(new(model.Login))
	db.Sync2(new(model.Token))
	db.Sync2(new(model.User))

	log.Printf("Success connecting MySQL to %s:%s with pass %s\n", opt.Host, opt.Port, opt.Password)
	return Mysql{xorm: db}, nil
}

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

func (m *Mysql) FindToken(tokenString string) model.Token {
	var token = model.Token{Token: tokenString}
	has, _ := m.xorm.Get(&token)
	if has {
		return token
	} else {
		return model.Token{}
	}
}

func (m *Mysql) CreateToken(token model.Token) model.Token {
	affected, err := m.xorm.Insert(&token)
	if err != nil {
		log.Printf("%d error %s", affected, err)
	}
	return m.FindToken(token.Token)
}

func (m *Mysql) CreateLogin(login model.Login) {
	affected, err := m.xorm.Insert(&login)
	if err != nil {
		log.Printf("%d error %s", affected, err)
	}
}

func (m *Mysql) TouchUserLogin(r *http.Request, user model.User) model.User {
	user.LatestLogin = time.Now()
	user.LastIp = realip.FromRequest(r)
	m.xorm.Id(user.Id).Update(user)
	return user
}
