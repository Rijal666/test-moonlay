package main

import (
	"fmt"
	"todo-list/database"
	"todo-list/pkg/mysql"
	"todo-list/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	mysql.DataBaseInit()
	database.RunMigration()

	routes.Routes(e.Group("api/v1"))

	port := "5002"
	fmt.Println("server running on port", port)
	e.Logger.Fatal(e.Start("localhost:" + port))
}
