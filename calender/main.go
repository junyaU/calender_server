package main

import (
	"calender/models"
	_ "calender/routers"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	beego.Run()
}

func setupDB() {
	db := beego.AppConfig.String("driver")

	orm.RegisterDriver(db, orm.DRMySQL)
	orm.RegisterDataBase("default", db, beego.AppConfig.String("sqlconn")+"?charset=utf8")
	orm.RegisterModel(
		new(models.User),
	)
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println(err)
	}
}
