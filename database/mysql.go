package database

import (
	"fmt"
	"log"

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
	return model.User{}
}

func (m *Mysql) FindUserByEmail(email string) model.User {
	return model.User{}
}

func (m *Mysql) FindUserByEmailAndPass(user model.User) (model.User, error) {
	return model.User{}, nil
}

func (m *Mysql) CreateToken(token model.Token) model.Token {
	return model.Token{}
}

func (m *Mysql) CreateLogin(login model.Login) {

}
