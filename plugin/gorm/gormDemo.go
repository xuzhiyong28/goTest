package gorm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// https://blog.csdn.net/yoyogu/article/details/109318626
type User struct {
	Id   uint
	Name string
	Age  uint
}

type Staff struct {
	Id       	int64  `gorm:"column:username;not null;type:int(4) primary key auto_increment;comment:'用户名'"`
	Password 	string `gorm:"column:password;type:varchar(30);index:idx_name"`
	CreateTime 	int64 `gorm:"column:createtime"`
}

func connDb() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/gotest?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	return db, nil
}

func InsertToTable() {
	db, _ := connDb()
	defer db.Close()
	u := User{
		Name: "xuzhiyong",
		Age:  12,
	}
	db.Table("user").Create(&u)
}

//根据ID查询
func FindById() {
	db, _ := connDb()
	defer db.Close()
	u := User{Id: 6}
	db_temp := db.Table("user").First(&u)
	defer db_temp.Close()
	if db_temp.RecordNotFound() {
		fmt.Println("没找到")
	} else {
		fmt.Println("找到", u)
	}
}

func FindByWhere() {
	db, _ := connDb()
	defer db.Close()
	var u User
	db.Where("name=?", "许志勇").First(&u)
	fmt.Println(u)

	var us []User
	db.Where("name=?", "许志勇").Find(&us)
	fmt.Println(us)

	var u2 User
	db.Where("name = ? and age = ? ", "许志勇", 3).First(&u2)
	fmt.Println(u2)
}

func CrateTable(){
	db, _ := connDb()
	defer db.Close()
	db.Table("staff").CreateTable(&Staff{})
}
