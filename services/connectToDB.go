package services

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq" // IMPORTANT TO orm.RegisterDataBase
)

func ConnectToDB() {
	//orm.RegisterDriver("postgres", 	orm.DRPostgres)
	//orm.RegisterDataBase("default", "postgres", "postgres:123@/orm_test?charset=utf8")

	user := beego.AppConfig.String("user")
	pass := beego.AppConfig.String("pass")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	db_name := beego.AppConfig.String("db_name")

	dbConnStr := "user=" + user + " password=" + pass + " host=" + host + " port=" + port + " dbname=" + db_name + " sslmode=disable"
	//"postgres://postgres:123456@127.0.0.1:5432/postgres?sslmode=disable"

	//dbConnStr := user + ":" + pass + "@" + host + ":" + port + "/" + db_name + "?" + "sslmode=disable"

	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		log.Error(err)
	}
	err = orm.RegisterDataBase("default",
		"postgres",
		dbConnStr)
	if err != nil {
		log.Error(err)
	}
}
