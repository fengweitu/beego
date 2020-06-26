package controllers

import (
	"beegodemo02/models"
	"path"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type AddArticleController struct {
	beego.Controller
}

func (this *AddArticleController) ShowAddArticle() {

	o := orm.NewOrm()
	var articleType []models.ArticleType
	qs := o.QueryTable("ArticleType")
	_, err := qs.All(&articleType)
	if err != nil {
		beego.Info("获取articleType类型失败")
		return
	}
	this.Data["articletype"] = articleType
	this.TplName = "add.html"
}

func (this *AddArticleController) HandleAddArticle() {
	articlename := this.GetString("articleName")
	articlecontent := this.GetString("content")

	beego.Info(articlename)
	beego.Info(articlecontent)
	file, header, err := this.GetFile("uploadname")
	defer file.Close()

	if err != nil {
		beego.Info("文件上传失败：", err)
		return
	}

	ext := path.Ext(header.Filename)
	//判断文件格式uploadname
	phototype := []string{".jpg", ".png", ".jpeg"}
	sum := 0
	for _, i := range phototype {
		if ext == i {
			sum++
		}
	}
	if sum != 1 {
		return
	}

	//判断文件大小
	if header.Size > 1000000 {
		beego.Info("文件过大")
		return
	}

	//文件不重名
	filetime := time.Now().Format("2006-01-02 15-04-05")
	path := "./static/img/" + filetime + ext
	beego.Info(path)
	err = this.SaveToFile("uploadname", path)

	if err != nil {
		beego.Info("保存文件出错：", err)
	}

	//插入数据
	//获取orm对象
	o := orm.NewOrm()

	//初始化article对象
	article := models.Article{}

	//对article对象赋值
	article.Title = articlename
	article.Content = articlecontent
	article.Img = "." + path

	//获取articleType对象
	articleType := models.ArticleType{}
	articleType.TypeName = this.GetString("select")
	if articleType.TypeName == "" {
		beego.Info("从视图中获取typename对象失败")
		return
	}
	err = o.Read(&articleType, "TypeName")
	if err != nil {
		beego.Info("获取articletype对象失败", err)
		return
	}
	article.ArticleType = &articleType

	//插入
	_, err = o.Insert(&article)
	if err != nil {
		beego.Info("插入失败：", err)
		return
	}

	this.Redirect("/v1/index", 302)

}
