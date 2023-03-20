package main

import (
	"fmt"
	
	"todo_postgresql_app/app/controllers"
	"todo_postgresql_app/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}
