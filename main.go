package main

import (
	"github.com/AliBedaer/go-orm/database"
	orm "github.com/AliBedaer/go-orm/database"
)

func main() {
	db := database.Orm.Init()
	con := orm.Orm.Init(5432, "localhost", "postgres", "password", "main_system", "postgres")
}
