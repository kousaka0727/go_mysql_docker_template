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
	fmt.Println("hello world !!")

	db := sqlConnect()
	defer db.Close()

}


func sqlConnect() (database *gorm.DB) {
	DBMS := "mysql"
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		"mysql",
		"3306",
		os.Getenv("MYSQL_DATABASE"),
	)

	count := 0
	db, err := gorm.Open(DBMS, dsn)
	if err != nil {
		for {
		if err == nil {
			fmt.Println("")
			break
		}
		fmt.Print(".")
		time.Sleep(time.Second)
		count++
		if count > 50 {
			fmt.Println("")
			fmt.Println("Connection to DB failed...")
			panic(err)
		}
		db, err = gorm.Open(DBMS, dsn)
		}
	}
	fmt.Println("Successful connection to DB !! ")

	return db
}
