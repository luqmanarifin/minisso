package database

import (
	"log"

	"github.com/luqmanarifin/minisso/model"
)

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
