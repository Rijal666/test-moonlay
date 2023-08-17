package database

import (
	"fmt"
	"todo-list/models"
	"todo-list/pkg/mysql"
)

func RunMigration() {
	err := mysql.ConnDB.AutoMigrate(
		&models.Todo{},
		&models.SubList{},
	)
	if err != nil {
		fmt.Println(err)
		panic("Migration failed")
	}
	fmt.Println("Migration Successfully")
}
