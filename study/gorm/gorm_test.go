package gormStudy

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func initMySqlGorm() *gorm.DB {
	if db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用自动创建外键约束
	}); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(20)
		sqlDB.SetMaxOpenConns(20)
		db.AutoMigrate(&Product{})
		return db
	}
}

func TestDemo0(t *testing.T) {
	db := initMySqlGorm()
	// 创建
	//db.Create(&Product{Code: "code00001"})
	// 批量插入
	products := []*Product{{Code: "code00002"}, {Code: "code00003"}, {Code: "code00004"}}
	result := db.Create(products)
	fmt.Printf("result.Error = %s", result.Error)
	fmt.Printf("result.RowsAffected = %d", result.RowsAffected)
	for _, value := range products {
		fmt.Println(value.ID.ID)
	}
}
