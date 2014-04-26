package routers

import (
	. "blog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &MainController{}, "GET:Index")
	beego.Router("/about", &AboutController{}, "GET:Index")
	beego.Router("/contact", &ContactController{}, "GET:Index")
	beego.Router("/articles/:id:int", &ArticlesController{}, "GET:Index")
	//beego.Router("/tags/:id:int", &MainController{})
	//beego.Router("/conmments/:id:int", &MainController{})
	//beego.Router("/date/:id:int", &MainController{})
	//beego.Router("/admin", &MainController{})
}
