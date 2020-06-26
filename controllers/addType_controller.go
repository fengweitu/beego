package controllers

import (
	"beegodemo02/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type AddArticleType struct {
	beego.Controller
}

func (this *AddArticleType) ShowArticleType() {

	//创建orm对象
	o := orm.NewOrm()

	//创建ArticleType对象
	var articleType []models.ArticleType

	//查询数据库对articleTypr对象赋值
	qs := o.QueryTable("article_type")
	_, err := qs.All(&articleType)
	if err != nil {
		beego.Info("查询数据出错：", err)
	}
	this.Data["articleType"] = articleType

	//将值传给视图
	this.TplName = "addType.html"
}

func (this *AddArticleType) HandleArticleType() {

	//创建orm对象
	o := orm.NewOrm()

	//创建articleType对象
	articleType := models.ArticleType{}

	//将articleType对象赋值
	typeName := this.GetString("typeName")
	if typeName == "" {
		beego.Info("插入的数据为空")
		return
	}
	articleType.TypeName = typeName
	beego.Info(typeName)

	//将数据插入数据库
	_, err := o.Insert(&articleType)
	if err != nil {
		beego.Info("插入数据库失败：", err)
		return
	}

	//跳转
	this.Redirect("/v1/addtype", 302)
}
