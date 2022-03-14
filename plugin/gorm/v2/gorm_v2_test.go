package v2

import (
	"example/plugin/gorm/v2/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
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

func TestCreate_001(t *testing.T) {
	tm := &model.TokenMapper{
		RootToken:  "bc",
		ChildToken: "bc",
		Mintable:   "N",
		MapType:    "POS",
		Name:       "XZY",
		Symbol:     "X",
		Deleted:    "N",
		ChainId:    15001,
		CreatedAt:  time.Now(),
	}
	db := initDb()
	result := db.Create(tm) // 会自动填充主键
	fmt.Println(tm)
	fmt.Println(result.Error)        // 返回 error
	fmt.Println(result.RowsAffected) // 返回插入记录的条数
}

func TestCreate_002(t *testing.T) {
	db := initDb()
	tm := &model.TokenMapper{
		RootToken:  "bc",
		ChildToken: "bc",
		Mintable:   "N",
	}
	// INSERT INTO `token_mapper`(`root_token`,`child_token`,`mintable`) VALUES ('','','')
	result := db.Select("RootToken", "ChildToken", "Mintable").Create(tm)

	// db.omit与 db.select是相反的，表示添加的是除了RootToken,ChildToken,Mintable这些其他的字段
	// db.Omit("RootToken","ChildToken","Mintable").Create(tm)

	fmt.Println(tm.ID)
	fmt.Println(result.Error)        // 返回 error
	fmt.Println(result.RowsAffected) // 返回插入记录的条数
}

func TestCreate_003(t *testing.T) {
	// 批量插入
	var tms = []model.TokenMapper{
		{
			RootToken:  "bc",
			ChildToken: "bc",
			Mintable:   "N",
			MapType:    "POS",
			Name:       "XZY",
			Symbol:     "X",
			Deleted:    "N",
			ChainId:    15001,
			CreatedAt:  time.Now(),
		},
		{
			RootToken:  "dc",
			ChildToken: "dc",
			Mintable:   "N",
			MapType:    "POS",
			Name:       "XZY",
			Symbol:     "X",
			Deleted:    "N",
			ChainId:    15002,
			CreatedAt:  time.Now(),
		},
	}
	db := initDb()
	result := db.Create(&tms)
	// db.CreateInBatches可以指定批量插入的数量，但是一般不用，有多少数量自己构建的数组知道
	// result := db.CreateInBatches(&tms, 100)

	fmt.Println(result.Error)        // 返回 error
	fmt.Println(result.RowsAffected) // 返回插入记录的条数
}

// map方式创建
func TestCreate_004(t *testing.T) {
	db := initDb()
	result1 := db.Model(&model.TokenMapper{}).Create(map[string]interface{}{
		"RootToken":  "dc",
		"ChildToken": "bc",
		"Mintable":   "N",
	})
	fmt.Println(result1.Error)        // 返回 error
	fmt.Println(result1.RowsAffected) // 返回插入记录的条数

	db.Model(&model.TokenMapper{}).Create([]map[string]interface{}{
		{
			"RootToken":  "dc",
			"ChildToken": "bc",
			"Mintable":   "N",
		},
		{
			"RootToken":  "dc",
			"ChildToken": "bc",
			"Mintable":   "N",
		},
	})
}

func TestQuery_001(t *testing.T) {
	db := initDb()
	tm1 := &model.TokenMapper{}
	// SELECT * FROM `token_mapper` ORDER BY `token_mapper`.`id` LIMIT 1
	result := db.First(tm1)
	fmt.Println(result.RowsAffected) // 找到的记录数

	tm2 := &model.TokenMapper{}
	// SELECT * FROM `token_mapper` LIMIT 1
	db.Take(tm2)

	tm3 := &model.TokenMapper{}
	// SELECT * FROM `token_mapper` ORDER BY `token_mapper`.`id` DESC LIMIT 1
	db.Last(tm3)

	data := map[string]interface{}{}
	db.Table("token_mapper").Take(&data)

}

func TestQuery_002(t *testing.T) {
	db := initDb()
	tm1 := &model.TokenMapper{}
	db.First(&tm1, 2) // SELECT * FROM `token_mapper` WHERE `token_mapper`.`id` = 2 ORDER BY `token_mapper`.`id` LIMIT 1

	var tms []*model.TokenMapper
	db.Find(&tms, []int{9, 10, 11}) // SELECT * FROM `token_mapper` WHERE `token_mapper`.`id` IN (9,10,11)
}

func TestQuery_003(t *testing.T) {
	db := initDb()

	var tm1s []*model.TokenMapper
	db.Where("map_type = ?", "POS").Find(&tm1s) // SELECT * FROM `token_mapper` WHERE map_type = 'POS'

	var tm2s []*model.TokenMapper
	db.Where("map_type IN ?", []string{"POS", "POW"}).Find(&tm2s) //SELECT * FROM `token_mapper` WHERE map_type IN ('POS','POW')

	var tm3s []*model.TokenMapper
	db.Where("map_type = ? AND deleted = ?" , "POS" , "N").Find(&tm3s)
}

// 原生 SQL
func TestQuery_004(t *testing.T){
	type Result struct {
		RootChain string
		ChildToken string
	}
	var results []*Result
	db := initDb()
	db.Raw("SELECT root_token AS RootChain,child_token AS ChildToken from token_mapper").Scan(&results)

}


func initDb() *gorm.DB {
	username := "root"   //账号
	password := "123456" //密码
	host := "127.0.0.1"  //数据库地址，可以是Ip或者域名
	port := 3306         //数据库端口
	Dbname := "gotest"   //数据库名
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
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
