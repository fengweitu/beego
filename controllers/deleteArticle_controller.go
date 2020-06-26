package controllers

import (
	"beegodemo02/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type DeleteArticleController struct {
	beego.Controller
}

func (this *DeleteArticleController) HandleDeleteArticle() {

	//获取要删除对象的id
	id, err := this.GetInt("id")
	if err != nil {
		beego.Info("获取id失败：", err)
		return
	}

	//获取orm对象
	o := orm.NewOrm()

	//获取article对象
	article := models.Article{Id: id}

	//从数据库中删除对象
	_, err = o.Delete(&article)
	if err != nil {
		beego.Info("从数据库中删除数据出错：", err)
	}

	//重定向
	this.Redirect("/v1/index", 302)
}
