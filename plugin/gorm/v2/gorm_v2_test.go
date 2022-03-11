package v2

import (
	"example/plugin/gorm/v2/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
	"time"
)

func TestInitDB(t *testing.T) {
	initDb()
}

func TestAutoMigrate(t *testing.T) {
	db := initDb()
	// 迁移 schema - 主要是自动在数据库端创建表
	// 一般不用，我们最好在数据库端创建表
	err := db.AutoMigrate(&model.TokenMapper{})
	if err != nil {
		log.Fatal(err)
	}
}

func initDb() *gorm.DB {
	username := "root"   //账号
	password := "123456" //密码
	host := "127.0.0.1"  //数据库地址，可以是Ip或者域名
	port := 3306         //数据库端口
	Dbname := "gotest"   //数据库名
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{})
	// gorm使用的是database/sql维护的连接池
	if err != nil {
		log.Fatal(err)
		return nil
	}
	sqlDB, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}
