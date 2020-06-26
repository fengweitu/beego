package routers

import (
	"beegodemo02/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {

	beego.InsertFilter("/v1/*", beego.BeforeRouter, FilterFunc)
	beego.Router("/register", &controllers.RegisterController{}, "get:ShowRegister;post:HandleRegister")
	beego.Router("/login", &controllers.LoginController{}, "get:ShowLogin;post:HandleLogin")
	beego.Router("/v1/index", &controllers.IndexController{}, "get:ShowIndex;post:HandleSelect")
	beego.Router("/v1/addarticle", &controllers.AddArticleController{}, "get:ShowAddArticle;post:HandleAddArticle")
	beego.Router("/v1/content", &controllers.ContentController{}, "get:ShowContent")
	beego.Router("/v1/deletearticle", &controllers.DeleteArticleController{}, "get:HandleDeleteArticle")
	beego.Router("/v1/updatearticle", &controllers.UpdateArticleController{}, "get:ShowUpdateArticle;post:HandleUpdateArticle")
	beego.Router("/v1/addtype", &controllers.AddArticleType{}, "get:ShowArticleType;post:HandleArticleType")
	beego.Router("/v1/deletetype", &controllers.DeleteType{}, "get:HandleDeleteType")
	beego.Router("/v1/logout", &controllers.LogoutController{}, "get:Logout")
}

var FilterFunc = func(ctx *context.Context) {
	username := ctx.Input.Session("userName")
	if username == nil {
		ctx.Redirect(302, "/login")

	}
}
