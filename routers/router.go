package routers

import(
	"github.com/astaxie/beego"
	"ispider/controllers/api"
	"ispider/controllers/admin"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/book",
			// /v1/book  获取小说列表
			beego.NSRouter("/", &api.BookController{}, "get:GetAll"),
			// /v1/book/getchapters?bookid=xxx&page=1  获取小说章节列表
			beego.NSRouter("/getchapters", &api.BookController{}, "get:GetChapters"),
			// /v1/book/chapter?id=xxx  获取章节内容
			beego.NSRouter("/chapter", &api.BookController{}, "get:GetChapter"),
		),
	)
	beego.AddNamespace(ns)

	nsAdmin := beego.NewNamespace("/admin",
		beego.NSRouter("/", &admin.IndexController{},"get:Index"),
		beego.NSRouter("/login", &admin.LoginController{},"get:Login"),
		beego.NSRouter("/users", &admin.UserController{},"get:Users"),
		beego.NSRouter("/users/edit", &admin.UserController{},"get:Edit"),

		
		beego.NSNamespace("/service",
			beego.NSRouter("/login", &admin.LoginController{},"post:AjaxLogin"),
			beego.NSRouter("/user_add", &admin.UserController{},"post:AjaxAdd"),
			beego.NSRouter("/user_delete", &admin.UserController{},"post:AjaxDelete"),
			beego.NSRouter("/user_enable", &admin.UserController{},"post:AjaxEnable"),
			beego.NSRouter("/user_disable", &admin.UserController{},"post:AjaxDisable"),
			beego.NSRouter("/user_edit", &admin.UserController{},"post:AjaxEdit"),
		),
	)
	beego.AddNamespace(nsAdmin)
}