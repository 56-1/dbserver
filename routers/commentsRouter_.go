package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["dbserver/controllers:ObjectController"] = append(beego.GlobalControllerRouter["dbserver/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/:who/:dowhat",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
