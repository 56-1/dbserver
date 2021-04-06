package controllers

import (
	_ "dbserver/models"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Key string
	Orm orm.Ormer
)

func init(){
	initDatabase()
	Key, _ = beego.AppConfig.String("key")
	Orm = orm.NewOrm()
}

func initDatabase(){
	user, _ := beego.AppConfig.String("mysqluser")
	pass, _ := beego.AppConfig.String("mysqlpass")
	ip, _ := beego.AppConfig.String("mysqlip")
	port, _ := beego.AppConfig.String("mysqlport")
	db, _ := beego.AppConfig.String("mysqldb")

	connectSql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, pass, ip, port, db)
	orm.RegisterDataBase("default", "mysql", connectSql)

	orm.RunSyncdb("default", false, true)
}
