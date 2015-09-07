package controllers

import (
	"deploy_sys/service"
	"encoding/json"
	"github.com/astaxie/beego"
	"strings"
)

type HookController struct {
	beego.Controller
}

var (
	TOKEN = beego.AppConfig.String("git_hook_token")
)

func (this *HookController) PostHook() {

	// cmdLine := []string{"pull", "origin", "master"}
	// git.RunCmd("git", cmdLine)
	var ob map[string]interface{}
	ob = make(map[string]interface{})
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	token, ok := ob["token"].(string)
	if ok {
		if strings.EqualFold(token, TOKEN) {
			println("token ok")
			go git.Pull("/Users/luxingmin/Documents/go/src/blog", "origin", "master")
		} else {
			println("token error")
		}
	} else {
		println("token not exist")
	}
	this.Data["json"] = ob
	this.ServeJson()
}
func (this *HookController) GetLog() {

}
func (this *HookController) test() {
	cmdLine := []string{"pull", "origin", "master"}
	git.RunCmd("git", cmdLine)
}
