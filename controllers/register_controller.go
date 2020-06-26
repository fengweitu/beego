package controllers

import (
	"beegodemo02/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) ShowRegister() {
	this.TplName = "register.html"
}

func (this *RegisterController) HandleRegister() {
	username := this.GetString("userName")
	password := this.GetString("password")
	message := fmt.Sprintf("username:%s,password:%s", username, password)
	beego.Info(message)

	//创建orm对象
	o := orm.NewOrm()

	//创建结构体对象
	user := models.User{}

	//对结构体对象进行赋值
	user.Username = username
	user.Password = password

	//将数据插入表中
	if user.Username != "" && user.Password != "" {
		o.Insert(&user)
	} else {
		beego.Info("用户名和密码不能为空")
		this.Ctx.WriteString("用户名和密码不能为空")
	}

	// this.Data["json"] = map[string]interface{}{"username": username, "password": password}
	// this.ServeJSON()

	this.Redirect("/login", 302)
}
