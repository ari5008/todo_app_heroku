package main

import (
	"fmt"
	
	"todo_postgres_app/app/controllers"
	"todo_postgres_app/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}
