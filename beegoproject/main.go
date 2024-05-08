package main

import (
	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"

	"beegoproject/models"
)

func init() {

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "core:core@/core?charset=utf8")

	orm.RegisterModel(new(models.Cars))

}

func main() {

	orm.RunSyncdb("default", false, true)

	beego.Run()
}
