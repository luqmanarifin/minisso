package service

import (
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/luqmanarifin/minisso/database"
	"github.com/luqmanarifin/minisso/model"
)

type ApplicationService struct {
	mysql database.Mysql
}

func NewApplicationService(mysqlOpt database.MysqlOption) ApplicationService {
	mysql, err := database.NewXorm(mysqlOpt)
	if err != nil {
		log.Fatal("cant connect to mysql")
	}
	return ApplicationService{
		mysql: mysql,
	}
}

func (a *ApplicationService) CreateApplication(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	EnableCors(&w)
	credential, _, _ := ExtractCredential(r)
	application := credential.Application
	application.ClientId = GenerateString(10)
	application.ClientSecret = GenerateString(15)

	a.mysql.CreateApplication(application)

	HandleResponse(w, nil, "", 200)
}

func (a *ApplicationService) FindAllApplications(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	EnableCors(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
	applications := a.mysql.FindAllApplications()

	HandleResponse(w, applications, "", 200)
}

func (a *ApplicationService) FindApplication(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	EnableCors(&w)
	strId := params.ByName("id")
	id, _ := strconv.ParseInt(strId, 10, 64)
	application := a.mysql.FindApplication(id)

	HandleResponse(w, application, "", 200)
}

func (a *ApplicationService) UpdateApplication(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	EnableCors(&w)
	credential, _, _ := ExtractCredential(r)
	application := credential.Application

	a.mysql.UpdateApplication(application)

	HandleResponse(w, nil, "", 200)
}

func (a *ApplicationService) DeleteApplication(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	EnableCors(&w)
	id, _ := strconv.ParseInt(params.ByName("id"), 10, 64)
	application := model.Application{Id: id}

	a.mysql.DeleteApplication(application)
	HandleResponse(w, nil, "", 200)
}
