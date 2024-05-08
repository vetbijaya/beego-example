package routers

import (
	"beegoproject/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/cars", &controllers.CarController{})
}
