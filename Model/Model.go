package Model

import (
	"github.com/sunbelife/Prelook_Strobe_Backend/Config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

const TIME_LAYOUT = "2006-01-02 15:04:05"
const FALSE = "false"
const TRUE = "true"

var db *gorm.DB

type Data struct {
	IDs []int
	Mode string
}

func Init() {
	MySQLConfig, err := Config.GetMySQLConfig()

	if err != nil {
		log.Println("mysql config err")
	}

	dsn := MySQLConfig.MySQL.UserName + ":" + MySQLConfig.MySQL.PassWord + "@tcp(127.0.0.1:3306)/" + MySQLConfig.MySQL.Database + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// 如果数据库链接异常则抛出
	if err != nil {
		log.Println(err)
	}
}
