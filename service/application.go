package service

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/luqmanarifin/minisso/database"
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
	credential, _, _ := ExtractCredential(r)
	application := credential.Application

	a.mysql.CreateApplication(application)
}

func (a *ApplicationService) FindAllApplications(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

func (a *ApplicationService) FindApplication(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

func (a *ApplicationService) UpdateApplication(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	credential, _, _ := ExtractCredential(r)
	application := credential.Application

	a.mysql.UpdateApplication(application)
}

func (a *ApplicationService) DeleteApplication(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	credential, _, _ := ExtractCredential(r)
	application := credential.Application

	a.mysql.DeleteApplication(application)
}
