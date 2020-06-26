package controllers

import (
	"beegodemo02/models"
	"math"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) ShowIndex() {

	// username := this.GetSession("username")
	// if username == nil {
	// 	beego.Info("username为空")
	// 	this.Redirect("/login", 302)
	// 	return
	// }

	//var article []models.Article
	//定义一个orm对象
	o := orm.NewOrm()

	//查询
	qs := o.QueryTable("article")
	// _, err := qs.All(&article)
	// if err != nil {
	// 	beego.Info("查询出错：", err)
	// }

	//获取文章总数
	// count, err := qs.Count()
	var count int64
	var err error
	articleType01 := this.GetString("select")
	if articleType01 == "" {
		count01, err01 := qs.Count()
		count = count01
		err = err01
	} else {
		count01, err01 := qs.RelatedSel("ArticleType").Filter("ArticleType__TypeName", articleType01).Count()
		count = count01
		err = err01
	}
	//count, err := qs.RelatedSel("ArticleType").Filter("ArticleType__TypeName", articleType01).Count()
	if err != nil {
		beego.Info("获取文章总数失败：", err)
		return
	}
	this.Data["Count"] = count

	//设置文章页数
	pageSize := 2
	pageCount := float64(count) / float64(pageSize)
	pageCount = math.Ceil(pageCount)
	this.Data["pageCount"] = pageCount

	//设置当前页文章数量
	pageNum := this.GetString("pageNum")
	beego.Info("当前页：", pageNum)
	index, err := strconv.Atoi(pageNum)
	if err != nil {
		beego.Info("获取页码失败：", err)
		//return
	}
	if index == 0 {
		index = 1
	}
	indexNum := pageSize * (index - 1)
	beego.Info("开始条数：", indexNum)
	// qs = qs.Limit(pageSize, indexNum)
	// _, err = qs.All(&article)
	// if err != nil {
	// 	beego.Info("查询数据失败：", err)
	// 	return
	// }

	firstPage := false
	if index == 1 {
		firstPage = true
	}
	lasterPage := false
	if index == int(pageCount) {
		lasterPage = true
	}

	//给文章下拉列表传值
	var articleType []models.ArticleType
	qs = o.QueryTable("articleType")
	_, err = qs.All(&articleType)
	if err != nil {
		beego.Info("获取类型失败：", err)
	}

	//获取下拉列表得值
	var articleWithType []models.Article
	articleSelectType := this.GetString("select")
	if articleSelectType == "" {
		qs = o.QueryTable("article")
		qs = qs.Limit(pageSize, indexNum)
		_, err = qs.RelatedSel("ArticleType").All(&articleWithType)
		if err != nil {
			beego.Info("查询数据失败：", err)
			return
		}
	} else {
		_, err = o.QueryTable("article").Limit(pageSize, indexNum).RelatedSel("ArticleType").Filter("ArticleType__TypeName", articleSelectType).All(&articleWithType)
		if err != nil {
			beego.Info("查询对象失败")
		}
	}

	username := this.GetSession("userName")

	//把数据传递给视图
	this.Data["userName"] = username
	this.Data["typeName"] = articleSelectType
	this.Data["articleType"] = articleType
	this.Data["FirstPage"] = firstPage
	this.Data["LasterPage"] = lasterPage
	this.Data["article"] = articleWithType
	this.Data["index"] = index

	this.TplName = "index.html"
}

func (this *IndexController) HandleSelect() {
	// articleType := this.GetString("select")
	// //beego.Info(articleType)

	// o := orm.NewOrm()

	// var article []models.Article
	// _, err := o.QueryTable("article").RelatedSel("ArticleType").Filter("ArticleType__TypeName", articleType).All(&article)
	// if err != nil {
	// 	beego.Info("查询对象失败")
	// }
	// beego.Info(article)
	// this.Ctx.WriteString("article")

}
