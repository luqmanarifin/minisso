package service

import (
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/luqmanarifin/minisso/database"
	"github.com/luqmanarifin/minisso/model"
)

type SessionService struct {
	mysql database.Mysql
}

func NewSessionService(mysqlOpt database.MysqlOption) SessionService {
	mysql, err := database.NewXorm(mysqlOpt)
	if err != nil {
		log.Fatal("can't connect to mysql")
	}
	return SessionService{
		mysql: mysql,
	}
}

func (a *SessionService) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	EnableCors(&w)
	credential, _, _ := ExtractCredential(r)
	user := credential.User

	a.mysql.CreateUser(user)

	HandleResponse(w, nil, "", 200)
}

func (a *SessionService) FindAllUsers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	EnableCors(&w)
	users := a.mysql.FindAllUsers()

	HandleResponse(w, users, "", 200)
}

func (a *SessionService) FindUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	EnableCors(&w)
	strId := params.ByName("id")
	id, _ := strconv.ParseInt(strId, 10, 64)
	user := a.mysql.FindUser(id)

	HandleResponse(w, user, "", 200)
}

func (a *SessionService) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	EnableCors(&w)
	credential, _, _ := ExtractCredential(r)
	user := credential.User

	a.mysql.UpdateUser(user)

	HandleResponse(w, nil, "", 200)
}

func (a *SessionService) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	EnableCors(&w)
	id, _ := strconv.ParseInt(params.ByName("id"), 10, 64)
	user := model.User{Id: id}

	a.mysql.DeleteUser(user)
	HandleResponse(w, nil, "", 200)
}
