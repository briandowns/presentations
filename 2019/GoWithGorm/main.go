package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=mysecretpassword")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	_ = db
}
