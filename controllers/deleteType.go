package controllers

import (
	"beegodemo02/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type DeleteType struct {
	beego.Controller
}

func (this *DeleteType) HandleDeleteType() {
	id, err := this.GetInt("id")
	if err != nil {
		beego.Info("获取id对象失败：", err)
		return
	}
	o := orm.NewOrm()
	articletype := models.ArticleType{Id: id}
	o.Delete(&articletype)
	this.Redirect("/v1/addtype", 302)

}
