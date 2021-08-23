package gorm

/***
常用教程 ：
	https://www.tizi365.com/archives/8.html
	https://gorm.io/zh_CN/docs/create.html
*/
import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"testing"
	"time"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}


// 钩子函数 在insert之前执行
func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	//fmt.Println("======BeforeCreate====")
	return
}


func TestModelBase(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移 schema
	db.AutoMigrate(&Product{})
	// Create
	db.Create(&Product{Code: "D42", Price: 100})
	// Read
	var product Product
	db.First(&product, 1)                 // 根据整形主键查找
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	// Delete - 删除 product
	db.Delete(&product, 1)
}

func initMysqlDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gotest?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatal("Gorm 连接异常 :", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Gorm - db.DB()异常 :", err)
	}
	sqlDB.SetMaxIdleConns(1000)         //设置最大连接数
	sqlDB.SetMaxOpenConns(1000)         //设置最大空闲连接池
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接可复用的最大时间
	//return db
	return db.Debug() //返回debug模式的，可以打印日志
}

func TestMySQLAutoMigrate(t *testing.T) {
	db := initMysqlDB()
	exist := db.Migrator().HasTable(&Product{}) //判断表是否存在
	if exist == false {
		// 自动建表
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Product{}) //参数可以多个，一次性建多个表
	}

}

func TestMySQLCreate(t *testing.T) {
	db := initMysqlDB()
	p := &Product{
		Code:  "P00001",
		Price: 1000,
	}
	result := db.Create(p)
	fmt.Println("ID = ", p.ID, ",error = ", result.Error, ",RowsAffected = ", result.RowsAffected)
}

func TestMySQLCreate2(t *testing.T) {
	db := initMysqlDB()
	p := &Product{
		Code:  "P00002",
		Price: 2000,
	}
	// 忽略掉price的插入
	result := db.Select("Code").Create(p)
	fmt.Println("ID = ", p.ID, ",error = ", result.Error, ",RowsAffected = ", result.RowsAffected)
}

func TestMySQLCreate3(t *testing.T) {
	db := initMysqlDB()
	//批量插入
	ps := []Product{
		{
			Code:  "P00001",
			Price: 1000,
		},
		{
			Code:  "P00002",
			Price: 2000,
		},
		{
			Code:  "P00002",
			Price: 3000,
		},
		{
			Code:  "P00004",
			Price: 4000,
		},
		{
			Code:  "P00005",
			Price: 5000,
		},
	}
	result := db.Create(ps)
	fmt.Println(",error = ", result.Error, ",RowsAffected = ", result.RowsAffected)
	for _, p := range ps {
		fmt.Println("ID = ", p.ID)
	}
}

func TestMySQLSelect(t *testing.T) {
	db := initMysqlDB()
	p1 := &Product{}
	db.First(p1)  // SELECT * FROM `products` WHERE `products`.`deleted_at` IS NULL ORDER BY `products`.`id` LIMIT 1
	fmt.Println(p1)

	p2 := &Product{}
	db.Take(p2)	 // SELECT * FROM `products` WHERE `products`.`deleted_at` IS NULL LIMIT 1
	fmt.Println(p2)

	p3 := &Product{}
	db.First(p3,20) // SELECT * FROM `products` WHERE `products`.`id` = 10 AND `products`.`deleted_at` IS NULL ORDER BY `products`.`id` LIMIT 1
	fmt.Println(p3)

	var p4 []Product
	db.Find(&p4, []int{20,21,22}) // SELECT * FROM `products` WHERE `products`.`id` IN (20,21,22) AND `products`.`deleted_at` IS NULL
	fmt.Println("p4 = ", p4)

	var p5 []Product
	db.Find(&p5)		// SELECT * FROM `products` WHERE `products`.`deleted_at` IS NULL
	fmt.Println("p5 = ", p5)

	var p6 []Product
	db.Where("code = ?" , "P00005").Find(&p6)  // SELECT * FROM `products` WHERE code = 'P00005' AND `products`.`deleted_at` IS NULL
	fmt.Println("p6 = ", p6)

	// 使用struct作为条件
	var p7 []Product
	db.Where(&Product{Code: "P00005", Price: 5000}).Find(&p7) // SELECT * FROM `products` WHERE `products`.`code` = 'P00005' AND `products`.`price` = 5000 AND `products`.`deleted_at` IS NULL
	fmt.Println("p7 = ", p7)

	var p8 []Product
	db.Where("code = ? And price = ?" , "P00005" , 5000).Find(&p8)
	fmt.Println("p8 = " , p8)

	var p9 []Product
	db.Not("code = ?", "P00005").Find(&p9) // SELECT * FROM `products` WHERE NOT code = 'P00005' AND `products`.`deleted_at` IS NULL
	fmt.Println(p9)

	var p10 []Product
	db.Where("code = ?", "P00005").Or("price = ?" , 3000).Find(&p10)  // SELECT * FROM `products` WHERE (code = 'P00005' OR price = 3000) AND `products`.`deleted_at` IS NULL
	fmt.Println(p10)


	// 选定特定字段输出
	var p11 []Product
	db.Select([]string{"code","price"}).Find(&p11)
	fmt.Println(p11)

	// 排序
	var p12 []Product
	db.Order("code desc").Find(&p12)  // SELECT * FROM `products` WHERE `products`.`deleted_at` IS NULL ORDER BY code desc
	fmt.Println(p12)

}

func TestMySQLSelect_2(t *testing.T) {
	// 原生SQL
	db := initMysqlDB()
	var p1 []Product
	db.Raw("SELECT id,code,price FROM products").Scan(&p1)
	fmt.Println(p1)
}

func TestMySQLUpdate(t *testing.T) {
	db := initMysqlDB()
	var p1 Product
	db.First(&p1) 		// SELECT * FROM `products` WHERE `products`.`deleted_at` IS NULL ORDER BY `products`.`id` LIMIT 1
	p1.Code = "P11111"
	p1.Price = 12345
	db.Save(&p1)		// UPDATE `products` SET `created_at`='2021-08-19 22:07:49.984',`updated_at`='2021-08-23 11:22:09.39',`deleted_at`=NULL,`code`='P11111',`price`=12345 WHERE `id` = 20
	fmt.Println(p1)
}

// 更新单个列
func TestMySQLUpdate_2(t *testing.T) {
	db := initMysqlDB()
	// UPDATE `products` SET `code`='P123456',`updated_at`='2021-08-23 11:23:58.443' WHERE code = 'P11111' AND `products`.`deleted_at` IS NULL
	db.Model(&Product{}).Where("code = ?" , "P11111").Update("code", "P123456")
}

func TestMySQLUpdate_3(t *testing.T) {

	db := initMysqlDB()
	// UPDATE `products` SET `updated_at`='2021-08-23 11:27:27.554',`code`='P111111',`price`=654321 WHERE code = 'P123456' AND `products`.`deleted_at` IS NULL
	db.Model(&Product{}).Where("code = ?" , "P123456").Updates(Product{Code: "P111111",Price: 654321})
}
