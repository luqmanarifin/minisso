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
