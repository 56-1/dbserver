package models

import "github.com/beego/beego/v2/client/orm"

type User struct {
	Account string	`orm:"column(account);pk" json:"account"`
	PassWord string `orm:"column(password)" json:"passwd"`
}

func init(){
	orm.RegisterModel(new(User))
}
