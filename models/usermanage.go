package models

import (
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
)

/*Read 从用户表中读取数据
**Param struct{id, pageIndex, pageSize} id为用户id或者为组织结构id，当id为用户id时， 不指定其它两个字段
**pageIndex为当前页面，pageSize为每页的大小
**return {[]*UserManage, nil} or {nil, err}
 */
func (u *UserManage)Read(Orm orm.Ormer, data []byte)(result interface{}, err error){
	p := struct{
		ID int `json:"id"`
		PageIndex int `json:"pageIndex"`
		PageSize int `json:"pageSize"`
	}{}

	err = json.Unmarshal(data, &p)
	if err != nil {
		return
	}

	var users []*UserManage
	qs := Orm.QueryTable(u)
	if p.PageSize == 0 { 	//PageSize == 0 时为查询指定的用户本身的数据
		_, err = qs.Filter("dev_id", p.ID).All(users)
		if err != nil {
			return
		}
		result = users
		return
	}

	_, err = qs.Filter("organization_id", p.ID).Limit(p.PageSize, (p.PageIndex - 1) * p.PageSize).OrderBy("user_id").All(users)
	if err != nil {
		return
	}

	result = users
	return
}

/*Del 根据需要删除的用户id从数据库中删除该用户
**Param UserManage{UserID}
**return {nil, err} or {nil, nil}
 */
func (u *UserManage)Del(Orm orm.Ormer, data []byte)(result interface{}, err error){
	err = json.Unmarshal(data, u)
	if err != nil {
		return
	}

	_, err = Orm.Delete(u)
	if err != nil {
		return
	}

	return
}

/*
**Add 添加新的用户，支持批量插入
**Param []UserManage 接收的参数为用户结构体列表
**return (nil, err) or (nil, nil)
 */
func (u *UserManage)Add(Orm orm.Ormer, data []byte)(result interface{}, err error){
	var users []UserManage
	err = json.Unmarshal(data, &users)
	if err != nil {
		return
	}

	_, err = Orm.InsertMulti(len(users), users)
	if err != nil {
		return
	}
	return
}

/*
**Update 更新指定的用户数据
**Param UserManage{UserID, UserName, RegistTime, DeadLine, OrgID}
**return (nil, err) or (nil, nil)
 */
func (u *UserManage)Update(Orm orm.Ormer, data []byte)(result interface{}, err error){
	err = json.Unmarshal(data, u)
	if err != nil {
		return
	}

	_, err = Orm.Update(u)
	if err != nil{
		return
	}
	return
}