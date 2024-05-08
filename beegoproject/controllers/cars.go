package controllers

import (
	"beegoproject/models"

	//"github.com/astaxie/beego"
	"github.com/beego/beego/v2/adapter/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type CarController struct {
	beego.Controller
}

func (c *CarController) Get() {
	o := orm.NewOrm()
	var cars []*models.Cars
	num, err := o.QueryTable("cars").All(&cars)
	if err != nil {
		c.Ctx.WriteString(err.Error())
	} else {
		c.Data["cars"] = cars
		c.Data["num"] = num
		c.TplName = "car/index.html"
	}
}
