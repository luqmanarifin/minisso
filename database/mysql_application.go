package database

import (
	"log"

	"github.com/luqmanarifin/minisso/model"
)

func (m *Mysql) CreateApplication(app model.Application) {
	affected, err := m.xorm.Insert(&app)
	if err != nil {
		log.Printf("%d error %s", affected, err)
	}
}

func (m *Mysql) FindAllApplications() []model.Application {
	applications := make([]model.Application, 0)
	return applications
}

func (m *Mysql) FindApplication(id int64) model.Application {
	var app = model.Application{Id: id}
	has, err := m.xorm.Get(&app)
	if err != nil {
		log.Printf("has %v error %v", has, err)
		return model.Application{}
	}
	if !has {
		return model.Application{}
	}
	return app
}

func (m *Mysql) FindApplicationByClientId(clientId string) model.Application {
	var app = model.Application{ClientId: clientId}
	has, err := m.xorm.Get(&app)
	if err != nil {
		log.Printf("has %v error %v", has, err)
		return model.Application{}
	}
	if !has {
		return model.Application{}
	}
	return app
}

func (m *Mysql) UpdateApplication(app model.Application) {
	affected, err := m.xorm.Id(app.Id).Update(&app)
	if err != nil {
		log.Printf("%d error %s", affected, err)
	}
}

func (m *Mysql) DeleteApplication(app model.Application) {
	affected, err := m.xorm.Id(app.Id).Delete(&model.Application{})
	if err != nil {
		log.Printf("affected %v error %s", affected, err.Error())
	}
}
