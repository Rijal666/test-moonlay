package mysql

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ConnDB *gorm.DB

func DataBaseInit() {
	var err error
	// DBurl := "root:@tcp(localhost:3306)/todo-list?charset=utf8mb4&parseTime=True&loc=Local"
	// ConnDB, err = gorm.Open(mysql.Open(DBurl), &gorm.Config{})
	DBurl := "host=localhost user=postgres password=123 dbname=test-moonlay sslmode=disable"
	ConnDB, err = gorm.Open(postgres.Open(DBurl), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	fmt.Println("connected to database")
}
