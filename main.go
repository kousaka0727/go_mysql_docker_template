package main

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)


func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("err : %v", err)
		panic("Error loading .env")
	}
}


func main() {
	fmt.Printf("#%v", "hello world")

	db := sqlConnect()
	defer db.Close()

}


func sqlConnect() (database *gorm.DB) {
	DBMS := "mysql"
	PROTOCOL := "tcp(mysql:3306)"
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	DBNAME := os.Getenv("MYSQL_DATABASE")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	count := 0
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		for {
		if err == nil {
			fmt.Println("")
			break
		}
		fmt.Print(".")
		time.Sleep(time.Second)
		count++
		if count > 180 {
			fmt.Println("")
			fmt.Println("DB接続失敗")
			panic(err)
		}
		db, err = gorm.Open(DBMS, CONNECT)
		}
	}
	fmt.Println("DB接続成功")

	return db
}
