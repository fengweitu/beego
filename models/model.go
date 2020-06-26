package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int        `orm:"pk;auto"`
	Username string     `size(50)`
	Password string     `size(100)`
	Article  []*Article `orm:"rel(m2m)"`
}

type Article struct {
	Id          int          `orm:"pk;auto"`
	Title       string       `orm:"size(20)"`
	Content     string       `orm:"size(200)"`
	Img         string       `orm:"size(50)"`
	Time        time.Time    `orm:"type(datetime);auto_now_add"`
	Count       int          `orm:"default(0)"`
	User        []*User      `orm:"reverse(many)"`
	ArticleType *ArticleType `orm:"rel(fk)"`
}

type ArticleType struct {
	Id       int
	TypeName string     `orm:"size(20)`
	Article  []*Article `orm:"reverse(many)"`
}

func init() {
	drivername := beego.AppConfig.String("drivername")

	username := beego.AppConfig.String("username")
	password := beego.AppConfig.String("password")
	port := beego.AppConfig.String("post")
	host := beego.AppConfig.String("host")
	dbname := beego.AppConfig.String("dbname")

	dbCon := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	orm.RegisterDataBase("default", drivername, dbCon)

	orm.RegisterModel(new(User), new(Article), new(ArticleType))
	orm.RunSyncdb("default", false, true)
}
