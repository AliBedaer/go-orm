package main

import (
	"fmt"
	"github.com/AliBedaer/go-orm/database"
)

func main() {
	orm := database.Orm{}

	orm.Init(5432, "localhost", "postgres", "password", "main_system", "postgres")

	fmt.Println(orm.Port)

	orm.Connect()
	q := orm.Select("name, age,id", "users").Where("asd", "=", "asd")

	fmt.Println(q)

	fmt.Println(orm.Query)

	orm.Execute()
}
