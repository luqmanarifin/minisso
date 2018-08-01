package service

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/luqmanarifin/minisso/database"
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
	cookie, err := r.Cookie("kentang")
	if err != nil {
		log.Printf("ambil cookie error")
	} else {
		log.Printf("cookie %s: %s", cookie.Name, cookie.Value)
	}
	setCookie := http.Cookie{Name: "luqman", Value: "ganteng"}
	http.SetCookie(w, &setCookie)
	HandleResponse(w, "", "ok", 200)

	// if there is cookie and valid, 200 ok already loggedin

	// if email already available, 200 ok reject

	// 201 created and give token
}

func (u *UserService) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// if there is cookie and valid, 200 ok already loggedin

	// if user pass invalid, 200 wrong password

	// 200 ok give token
}

func (u *UserService) Validate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// if there is cookie and valid, 200 ok give user info

	// 200 wrong password
}
