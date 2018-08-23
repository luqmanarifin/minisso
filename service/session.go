package service

import (
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/luqmanarifin/minisso/database"
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
	credential, _, _ := ExtractCredential(r)
	user := credential.User

	a.mysql.CreateUser(user)

	HandleResponse(w, nil, "", 200)
}

func (a *SessionService) FindAllUsers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	users := a.mysql.FindAllUsers()

	HandleResponse(w, users, "", 200)
}

func (a *SessionService) FindUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	strId := params.ByName("id")
	id, _ := strconv.ParseInt(strId, 10, 64)
	user := a.mysql.FindUser(id)

	HandleResponse(w, user, "", 200)
}

func (a *SessionService) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	credential, _, _ := ExtractCredential(r)
	user := credential.User

	a.mysql.UpdateUser(user)

	HandleResponse(w, nil, "", 200)
}

func (a *SessionService) DeleteSession(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	credential, _, _ := ExtractCredential(r)
	user := credential.User

	a.mysql.DeleteUser(user)
	HandleResponse(w, nil, "", 200)
}
