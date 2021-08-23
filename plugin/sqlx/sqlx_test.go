package sqlx

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	// 必须要导入mysql 不然报错
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

func initDataBase() *sqlx.DB {
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gotest")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return nil
	}
	return database
}

func TestInsertDemo(t *testing.T) {
	db := initDataBase()
	defer db.Close()
	r, err := db.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("insert succ:", id)
}

func TestSelectDemo(t *testing.T) {
	db := initDataBase()
	defer db.Close()
	var person []Person
	err := db.Select(&person, "select user_id, username, sex, email from person where user_id=?", 1)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("select succ:", person)
}

func TestUpdateDemo(t *testing.T) {
	db := initDataBase()
	defer db.Close()
	res, err := db.Exec("update person set username=? where user_id=?", "stu0003", 1)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
	}
	fmt.Println("update succ:", row)
}

func TestDeleteDemo(t *testing.T) {
	db := initDataBase()
	defer db.Close()
	res, err := db.Exec("delete from person where user_id=?", 1)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
	}

	fmt.Println("delete succ: ", row)
}

func TestTxDemo(t *testing.T) {
	db := initDataBase()
	defer db.Close()
	conn, err := db.Begin()
	if err != nil {
		fmt.Println("begin failed :", err)
		return
	}
	r, err := conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert succ:", id)
	r, err = conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	id, err = r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert succ:", id)
	conn.Commit()
}
