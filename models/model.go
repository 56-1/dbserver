package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

//管理员
type User struct {
	Id int	`json:"-"`
	Name string `json:"name"`
	Pwd string `json:"pwd"`
}

//用户
type UserManage struct {
	UserID string `orm:"column(user_id);pk" json:"user_id"`
	UserName string `json:"user_name"`
	RegistTime time.Time `orm:"auto_now_add" json:"regist_time"`
	DeadLine time.Time `json:"dead_line"`
	OrgID int 	`orm:"column(organization_id)" json:"org_id"`
}

//设备
type DevManage struct {
	DevID int `orm:"column(dev_id);pk" json:"dev_id"`
	DevName string `json:"dev_name"`
	AddDevTime time.Time `orm:"auto_now_add" json:"add_dev_time"`
	OrgID int	`orm:"column(organization_id)" json:"org_id"`
}

type ChanManage struct {
	ChanID string `orm:"column(chan_id);pk" json:"chan_id"`
	ChanName string `json:"chan_name"`
	AddChanTime time.Time `orm:"auto_now_add" json:"add_chan_time"`
	DevID int	`orm:"column(device_id)" json:"dev_id"`
	OrgID int	`orm:"column(organization_id)" json:"org_id"`
}

type DevOrgManage struct {
	OrgID int `orm:"column(organization_id);pk" json:"org_id"` 
	OrgName string `json:"org_name"`
	ParentID int `orm:"column(parent_id)" json:"parent_id"`
}

type UserOrgManage struct {
	OrgID int `orm:"column(organization_id);pk" json:"org_id"`
	OrgName string `json:"org_name"`
	ParentID int `orm:"column(parent_id)" json:"parent_id"`
}

func init(){
	orm.RegisterModel(new(User), new(UserManage), new(DevManage), new(ChanManage), new(DevOrgManage), new(UserOrgManage))
}
