package service

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/luqmanarifin/minisso/database"
	"github.com/luqmanarifin/minisso/model"
)

type UserService struct {
	mysql database.Mysql
}

func NewUserService(mysqlOption database.MysqlOption) UserService {
	mysql, err := database.NewXorm(mysqlOption)
	if err != nil {
		log.Fatal("cant connect to mysql")
	}
	return UserService{
		mysql: mysql,
	}
}

func (u *UserService) Signup(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	user := model.User{
		FirstName: "Luqman",
		LastName:  "Arifin",
	}
	u.mysql.CreateUser(user)
}

func Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

func Validate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}
