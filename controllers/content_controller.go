package controllers

import (
	"beegodemo02/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ContentController struct {
	beego.Controller
}

func (this *ContentController) ShowContent() {

	//获取id
	id, err := this.GetInt("id")
	if err != nil {
		beego.Info("获取id失败：", err)
	}

	//获取orm对象
	o := orm.NewOrm()

	//获取article对象
	article := models.Article{Id: id}

	//根据id查询数据库对象
	err = o.Read(&article)
	if err != nil {
		beego.Info("查询数据库失败：", err)
	}

	//实现每查询一次阅读量加一
	article.Count += 1

	//更新Count字段数据
	o.Update(&article, "Count")

	userName := this.GetSession("userName")
	m2m := o.QueryM2M(&article, "User")
	user := models.User{}
	user.Username = userName.(string)
	err = o.Read(&user, "Username")
	if err != nil {
		beego.Info("查询user失败", err)
		return
	}
	_, err = m2m.Add(&user)
	if err != nil {
		beego.Info("插入失败：", err)
	}

	//查询读者
	//_, err = o.LoadRelated(&article, "User") //会显示重复数据
	var users []models.User
	o.QueryTable("User").Filter("Article__Article__Id", id).Distinct().All(&users)

	this.Data["User"] = users

	this.Data["article"] = article

	this.TplName = "content.html"
}
