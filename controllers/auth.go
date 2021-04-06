package controllers

import (
	"github.com/beego/beego/v2/server/web/context"
	"judger"
)

func Auth(ctx *context.Context){
	salt := ctx.Input.Query("salt")
	timsStamp := ctx.Input.Query("timestamp")
	token := ctx.Input.Query("token")

	s := judger.Stamp{Salt: salt, TimeStamp: timsStamp, Token: token}
	if !judger.Verify(s, []byte(Key)){
		msg := map[string]interface{}{
			"ret": 0,
			"message": "404",
		}
		ctx.Output.JSON(msg, false, true)
	}
	return
}
