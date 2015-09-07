package routers

import (
	"deploy_sys/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/hook", &controllers.HookController{}, "post:PostHook")
	beego.Router("/log", &controllers.HookController{}, "*:GetLog")
	//beego.
}
