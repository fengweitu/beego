package controllers

import (
	"beegodemo02/models"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) ShowLogin() {
	cookie := this.Ctx.GetCookie("username")
	if cookie != "" {
		checked := "checked"
		this.Data["checked"] = checked
	} else {
		checked := ""
		this.Data["checked"] = checked
	}
	this.Data["cookie"] = cookie
	this.TplName = "login.html"
}

func (this *LoginController) HandleLogin() {
	username := this.GetString("userName")
	password := this.GetString("password")
	message := fmt.Sprintf("username:%s,password:%s", username, password)
	beego.Info(message)
	if username == "" && password == "" {
		//this.Ctx.WriteString("用户名和密码不能为空")
		this.TplName = "login.html"
		return
	}

	//创建orm对象
	o := orm.NewOrm()

	//创建结构体对象
	user := models.User{}

	//给结构体对象赋值
	user.Username = username

	//查询
	err := o.Read(&user, "username")

	if err != nil {
		beego.Info("查询失败，密码错误", err)
		//this.Ctx.WriteString("用户名错误")
		this.TplName = "login.html"
		return
	}

	if user.Password != password {
		beego.Info("查询失败，密码错误")
		//this.Ctx.WriteString("密码错误")
		this.TplName = "login.html"
		return
	}

	// this.Ctx.SetCookie("username", username, time.Second*3600)
	checked := this.GetString("remember")
	if checked == "on" {
		this.Ctx.SetCookie("username", username, time.Second*3600)
	} else {
		this.Ctx.SetCookie("username", username, -1)
	}

	this.SetSession("userName", username)

	// this.Data["json"] = map[string]interface{}{"username": username, "password": password}
	// this.ServeJSON()
	this.Redirect("/v1/index", 302)
}
