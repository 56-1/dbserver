package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"log"
)

// Operations about object
type ObjectController struct {
	beego.Controller
}

// @Title Create
// @Description create instance in object
// @Success 200
// @Failure 403 body is empty
// @router /:who/:dowhat [post]
func (o *ObjectController) Post() {
	who := o.Ctx.Input.Param(":who")
	doWhat := o.Ctx.Input.Param(":dowhat")

	u, operate := GetInstance(who, doWhat)
	if u == nil  || operate == nil {
		msg := map[string]interface{}{
			"ret": 0,
			"message": "module or direct method error",
		}
		o.Data["json"] = &msg
		o.ServeJSON()
		return
	}

	result, err := operate(u, Orm, o.Ctx.Input.RequestBody)
	if err != nil {
		log.Printf("in object post %v\n", err.Error())
		msg := map[string]interface{}{
			"ret": 0,
			"message": "server error",
		}
		o.Data["json"] = &msg
		o.ServeJSON()
		return
	}

	msg := map[string]interface{}{
		"ret": 1,
		"data": result,
	}
	o.Data["json"] = &msg
	o.ServeJSON()
}
