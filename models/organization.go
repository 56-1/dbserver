package models

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/beego/beego/v2/client/orm"
)

/*Read 查询设备和通道所属的组织结构
**Param struct{ID, PageIndex, PageSize, tree}
**ID == 0 为查询所有的组织结构数据；ID == 组织结构Id，查询指定组织结构的数据；ID == 父级Id时，查询该Id下的所有子组织结构
**tree 用于指定查询设备所属的组织结构或是查询用户所属的组织结构
**return {[]*DevOrgManage, nil} or {nil, err}
 */
func (o *DevOrgManage)Read(Orm orm.Ormer, data []byte)(result interface{}, err error){
	p := struct {
		ID int `json:"id"`
		PageIndex int `json:"pageIndex"`
		PageSize int `json:"pageSize"`
		Tree string `json:"tree"`
	}{}

	err = json.Unmarshal(data, &p)
	if err != nil {
		return
	}

	table := ""
	switch p.Tree {
	case "userOrgTree":
		table = "user_org_manage"
	case "devOrgTree":
		table = "dev_org_manage"
	default:
		err = errors.New("choice the org tree that you want, please")
		return
	}

	switch {
	case p.ID == 0:
		var orgs []*DevOrgManage
		qs := Orm.QueryTable(table)
		_, err = qs.All(orgs)
		if err != nil {
			return
		}
		result = orgs
	case p.PageSize == 0:
		var orgs []*DevOrgManage
		qs := Orm.QueryTable(table)
		_, err = qs.Filter("org_id", p.ID).All(orgs)
		if err != nil {
			return
		}
		result = orgs
	default:
		var orgs []*DevOrgManage
		qs := Orm.QueryTable(table)
		_, err = qs.Filter("parent_id", p.ID).Limit(p.PageSize, (p.PageIndex - 1) * p.PageSize).All(orgs)
		if err != nil {
			return
		}
		result = orgs
	}
	return
}

/*
**Del 用于删除一个组织结构，没有父节点的节点不支持删除，删除时会把子节点移向父节点
**Param DevOrgManage{OrgID}
**return {nil, err} or {nil, nil}
 */
func (o *DevOrgManage)Del(Orm orm.Ormer, data []byte)(result interface{}, err error){
	err = json.Unmarshal(data, o)
	if err != nil {
		return
	}

	err = Orm.Read(o)
	if err != nil {
		return
	}

	if o.ParentID == 0 {
		err = errors.New("this node don`t support delete")
		return
	}

	var r orm.RawSeter
	r = Orm.Raw("UPDATE org_manage SET parent_id=? WHERE parent_id=?", o.ParentID, o.OrgID)
	_, err = r.Exec()
	if err != nil {
		if err != sql.ErrNoRows {
			return
		}
	}

	_, err = Orm.Delete(o)
	if err != nil {
		return
	}
	return
}

/*
**Add 新增组织结构
**Param DevOrgManage{OrgID, OrgName, ParentID}
**return {nil, err} or {nil, nil}
 */
func (o *DevOrgManage)Add(Orm orm.Ormer, data []byte)(result interface{}, err error){
	err = json.Unmarshal(data, o)
	if err != nil {
		return
	}

	_, err = Orm.Insert(o)
	if err != nil {
		return
	}
	return
}

/*
**Update 用来更新组织结构数据
**Param DevOrgManage{OrgID, OrgName, ParentID}
**return {nil, err} or {nil, nil}
 */
func (o *DevOrgManage)Update(Orm orm.Ormer, data []byte)(result interface{}, err error){
	err = json.Unmarshal(data, o)
	if err != nil {
		return
	}

	_, err = Orm.Update(o)
	if err != nil {
		return
	}
	return
}
