package models

import (
	"encoding/json"
	"errors"
	"github.com/beego/beego/v2/client/orm"
)

/*Read 从通道表中读取数据
**Param struct{id, pageIndex, pageSize} id为通道id或者为组织结构id，当id为通道id时， 不指定其它两个字段
**pageIndex为当前页面，pageSize为每页的大小
**return {[]*ChanManage, nil} or {nil, err}
 */
func (c *ChanManage)Read(Orm orm.Ormer, data []byte)(result interface{}, err error){
	p := struct{
		ID int `json:"id"`
		PageIndex int `json:"pageIndex"`
		PageSize int `json:"pageSize"`
	}{}

	err = json.Unmarshal(data, &p)
	if err != nil {
		return
	}

	var chans []*ChanManage
	qs := Orm.QueryTable(c)
	if p.PageSize == 0 { 	//PageSize == 0 时为查询指定的通道本身的数据
		_, err = qs.Filter("dev_id", p.ID).All(chans)
		if err != nil {
			return
		}
		result = chans
		return
	}

	_, err = qs.Filter("organization_id", p.ID).Limit(p.PageSize, (p.PageIndex - 1) * p.PageSize).OrderBy("chan_id").All(chans)
	if err != nil {
		return
	}

	result = chans
	return
}

/*Del 通道没有删除选项，启用、停用的功能使用Update字段操作
**return {nil, err} or {nil, nil}
 */
func (c *ChanManage)Del(_ orm.Ormer, _ []byte)(result interface{}, err error){
	err = errors.New("chan don`t support delete")
	return
}

/*
**Add 添加新的通道，支持批量插入
**Param []ChanManage 接收的阐述为通道结构体列表
**return (nil, err) or (nil, nil)
 */
func (c *ChanManage)Add(Orm orm.Ormer, data []byte)(result interface{}, err error){
	var devs []ChanManage
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
**Update 更新指定的通道数据
**Param ChanManage{DevID, DevName, AddDevTime, OrgID}
**return (nil, err) or (nil, nil)
 */
func (c *ChanManage)Update(Orm orm.Ormer, data []byte)(result interface{}, err error){
	err = json.Unmarshal(data, c)
	if err != nil {
		return
	}

	_, err = Orm.Update(c)
	if err != nil{
		return
	}
	return
}