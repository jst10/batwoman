package main

import (
	"fmt"
	"made.by.jst10/celtra/batwoman/cmd/api"
	"made.by.jst10/celtra/batwoman/cmd/database"
)

func main(){
	fmt.Println("It works")
	database.InitDatabase()
	api.StartApi()
}