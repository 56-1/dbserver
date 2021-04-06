package controllers

import (
	"dbserver/models"
	"github.com/beego/beego/v2/client/orm"
)

type Union interface {
	Add(ormer orm.Ormer, data []byte)(result interface{}, err error)
	Del(ormer orm.Ormer, data []byte)(result interface{}, err error)
	Read(ormer orm.Ormer, data []byte)(result interface{}, err error)
	Update(ormer orm.Ormer, data []byte)(result interface{}, err error)
}

var operate = map[string]func(union Union, ormer orm.Ormer, data []byte)(result interface{}, err error){
	"add": Union.Add,
	"del": Union.Del,
	"read": Union.Read,
	"update": Union.Update,
}

func GetInstance(who, doWhat string)(u Union, method func(u Union, ormer orm.Ormer, data []byte)(result interface{}, err error)){
	switch who {
	case "user":
		u = new(models.User)
	}

	method = operate[doWhat]
	return
}
