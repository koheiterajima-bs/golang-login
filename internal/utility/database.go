package utility

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	user := os.Getenv("MYSQL_USER")
	pw := os.Getenv("MYSQL_PASSWORD")
	db_name := os.Getenv("MYSQL_DATABASE")
	var path string = fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true", user, pw, db_name)
	var err error
	if Db, err = sql.Open("mysql", path); err != nil {
		log.Fatal("Db open error:", err.Error())
	}
	checkConnect(100)

	fmt.Println("db connected!!")
}

func checkConnect(count uint) {
	var err error
	if err = Db.Ping(); err != nil {
		time.Sleep(time.Second * 2)
		count--
		fmt.Printf("retry... count:%v\n", count)
		checkConnect(count)
	}
}
