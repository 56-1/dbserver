package models

import (
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
)

func (u *User)Read(Orm orm.Ormer, data []byte)(result interface{}, err error){
	err = json.Unmarshal(data, u)
	if err != nil {
		return
	}

	err = Orm.Read(u)
	if err != nil {
		return
	}

	result = u
	return
}

func (u *User)Del(Orm orm.Ormer, data []byte)(result interface{}, err error){
	//TODO
	return
}

func (u *User)Add(Orm orm.Ormer, data []byte)(result interface{}, err error){
	//TODO
	return
}

func (u *User)Update(Orm orm.Ormer, data []byte)(result interface{}, err error){
	//TODO
	return
}
