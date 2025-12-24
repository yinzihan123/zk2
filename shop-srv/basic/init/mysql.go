package init

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"shop-srv/basic/globals"
	"shop-srv/handler/models"
)

func InitMysql() {
	var err error
	mysqlConfig := globals.AppConfig.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.User,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.Database)
	globals.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("数据库连接成功")
	}
	err = globals.DB.AutoMigrate(&models.User{}, &models.Banner{}, &models.Good{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("数据库迁移成功")
	}
}
