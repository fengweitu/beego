package controllers

import (
	"beegodemo02/models"
	"path"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UpdateArticleController struct {
	beego.Controller
}

func (this *UpdateArticleController) ShowUpdateArticle() {

	id, err := this.GetInt("id")
	if err != nil {
		beego.Info("获取id出错：", err)
		return
	}

	//获取orm对象
	o := orm.NewOrm()

	//获取article对象
	article := models.Article{Id: id}
	err = o.Read(&article)
	if err != nil {
		beego.Info("读取数据库失败：", err)
	}

	//将article数据填充到页面
	this.Data["article"] = article

	this.TplName = "update.html"
}

func (this *UpdateArticleController) HandleUpdateArticle() {

	// this.Data["json"] = map[string]interface{}{"id": id}
	// this.ServeJSON()

	//获取id
	id, err := this.GetInt("id")
	if err != nil {
		beego.Info("获取id失败")
	}

	articlename := this.GetString("articleName")
	content := this.GetString("content")
	if articlename == "" && content == "" {
		beego.Info("获取内容出错：", err)
	}

	file, header, err := this.GetFile("uploadname")
	defer file.Close()
	if err != nil {
		beego.Info("文件上传失败：", err)
		return
	}

	ext := path.Ext(header.Filename)
	sum := 0
	phototype := []string{".png", ".jpg", ".jepg"}
	for _, i := range phototype {
		if i == ext {
			sum++
		}
	}
	if sum != 1 {
		beego.Info("图片格式错误")
		return
	}

	if header.Size > 1000000 {
		beego.Info("文件过大")
		return
	}

	time := time.Now().Format("2006-01-02 03-04-05")
	path := "./static/img" + time + ext
	this.SaveToFile("uploadname", path)

	//获取orm对象
	o := orm.NewOrm()

	//获取article对象
	article := models.Article{Id: id}

	//获取article值
	err = o.Read(&article)
	if err != nil {
		beego.Info("获取article值失败：", err)
		return
	}

	//更新article的值
	article.Title = articlename
	article.Content = content
	article.Img = path

	//更新数据
	_, err = o.Update(&article)
	if err != nil {
		beego.Info("更新数据出错：", err)
		return
	}

	//this.Ctx.WriteString("更新成功")
	//跳转页面
	this.Redirect("/v1/index", 302)
}
