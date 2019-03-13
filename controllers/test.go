package controllers

import (
	"github.com/astaxie/beego"
	"kuakuaAi/models"
)

type TestController struct {
	beego.Controller
}

// @Title Test
// @Description test
// @Success 200 {object} []models.Blog
// @Param tag query string false "xx"
// @Param sort query string false "xx"
// @Param offset query int false "xx"
// @Param limit query int false "xx"
// @Failure 400 params error
// @router / [get]
func (this *TestController) Get() {
	this.Data["json"] = models.Test{Name: "jiangshiyi"}
	this.ServeJSON()
}
