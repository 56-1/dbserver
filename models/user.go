package models

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/beego/beego/v2/client/orm"
)


/*Read 使用用户名从数据库总读取对应的数据
**Param User{Name, Pwd}
**return {nil, "account not exist"/"password error"/err} or {nil, nil}
 */
func (u *User)Read(Orm orm.Ormer, data []byte)(result interface{}, err error){
	err = json.Unmarshal(data, u)
	if err != nil {
		return
	}

	user := User{Name: u.Name}
	err = Orm.Read(&user, "Name")
	if err != nil {
		if err != sql.ErrNoRows {
			return
		}
		err = errors.New("account not exist")
		return
	}

	if user.Pwd != u.Pwd {
		err = errors.New("password error")
		return
	}

	return
}

/*Del 根据用户名删除指定的用户
**Param User{Name, Pwd}
**return {nil, "account not exist"/err} or {nil, nil}
 */
func (u *User)Del(Orm orm.Ormer, data []byte)(result interface{}, err error){
	err = json.Unmarshal(data, u)
	if err != nil {
		return
	}

	user := User{Name: u.Name}
	err = Orm.Read(&user, "Name")
	if err != nil {
		if err != sql.ErrNoRows {
			return
		}
		err = errors.New("account not exist")
		return
	}

	_, err = Orm.Delete(&user)
	if err != nil {
		return
	}
	return
}

/*Add 添加新用户
**Param User{Name, Pwd}
**return {nil, err} or {nil, nil}
 */
func (u *User)Add(Orm orm.Ormer, data []byte)(result interface{}, err error){
	err = json.Unmarshal(data, u)
	if err != nil {
		return
	}

	_, err = Orm.Insert(u)
	if err != nil {
		return
	}
	return
}

/*Update 根据用户名更新用户数据
**Param User{Name, Pwd}
**return {nil, err} or {nil, nil}
 */
func (u *User)Update(Orm orm.Ormer, data []byte)(result interface{}, err error){
	err = json.Unmarshal(data, u)
	if err != nil {
		return
	}

	_, err = Orm.Update(u, "Pwd")
	if err != nil {
		return
	}
	return
}
