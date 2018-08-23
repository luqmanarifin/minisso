package database

import (
	"log"

	"github.com/luqmanarifin/minisso/model"
)

func (m *Mysql) CreateLogin(login model.Login) {
	affected, err := m.xorm.Insert(&login)
	if err != nil {
		log.Printf("%d error %s", affected, err)
	}
}
