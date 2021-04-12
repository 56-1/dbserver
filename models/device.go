package models

import (
	"database/sql"
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
)

/*Read 从设备表中读取数据
**Param struct{id, pageIndex, pageSize} id为设备id或者为组织结构id，当id为设备id时， 不指定其它两个字段
**pageIndex为当前页面，pageSize为每页的大小
**return {[]*DevManage, nil} or {nil, err}
 */
func (d *DevManage)Read(Orm orm.Ormer, data []byte)(result interface{}, err error){
	p := struct{
		ID int `json:"id"`
		PageIndex int `json:"pageIndex"`
		PageSize int `json:"pageSize"`
	}{}

	err = json.Unmarshal(data, &p)
	if err != nil {
		return
	}

	var devs []*DevManage
	qs := Orm.QueryTable(d)
	if p.PageSize == 0 { 	//PageSize == 0 时为查询指定的设备本身的数据
		_, err = qs.Filter("dev_id", p.ID).All(devs)
		if err != nil {
			return
		}
		result = devs
		return
	}

	_, err = qs.Filter("organization_id", p.ID).Limit(p.PageSize, (p.PageIndex - 1) * p.PageSize).OrderBy("dev_id").All(devs)
	if err != nil {
		return
	}

	result = devs
	return
}

/*Del 根据需要删除的设备id从数据库中删除该设备及该设备绑定的通道，通道删除失败时设备也删除失败
**Param DevManage{DevID}
**return {nil, err} or {nil, nil}
 */
func (d *DevManage)Del(Orm orm.Ormer, data []byte)(result interface{}, err error){
	err = json.Unmarshal(data, d)
	if err != nil {
		return
	}

	chanel := new(ChanManage)
	qs := Orm.QueryTable(chanel)
	_, err = qs.Filter("device_id", d.DevID).Delete()
	if err != nil {
		if err != sql.ErrNoRows {
			return
		}
	}

	qs = Orm.QueryTable(d)
	_, err = qs.Filter("dev_id", d.DevID).Delete()
	if err != nil {
		return
	}

	return
}

/*
**Add 添加新的设备，支持批量插入
**Param []DevManage 接收的阐述为设备结构体列表
**return (nil, err) or (nil, nil)
 */
func (d *DevManage)Add(Orm orm.Ormer, data []byte)(result interface{}, err error){
	var devs []DevManage
	err = json.Unmarshal(data, &devs)
	if err != nil {
		return
	}

	_, err = Orm.InsertMulti(len(devs), devs)
	if err != nil {
		return
	}
	return
}

/*
**Update 更新指定的设备数据
**Param DevManage{DevID, DevName, AddDevTime, OrgID}
**return (nil, err) or (nil, nil)
 */
func (d *DevManage)Update(Orm orm.Ormer, data []byte)(result interface{}, err error){
	err = json.Unmarshal(data, d)
	if err != nil {
		return
	}

	_, err = Orm.Update(d)
	if err != nil{
		return
	}
	return
}