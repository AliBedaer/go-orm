package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/AliBedaer/go-orm/database"
	"github.com/joho/godotenv"
)

type user struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	orm := database.Orm{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port, _ := strconv.ParseInt(os.Getenv("DATABASE_PORT"), 10, 64)

	orm.Init(port, os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME"), os.Getenv("DATABASE_TYPE"))

	db := orm.Connect()

	orm.Select("*", "users")

	rows, err := db.Query(orm.Query)

	if err != nil {
		panic(err.Error())
	}
	users := make([]user, 0)
	var user user
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		users = append(users, user)
		if err != nil {
			panic(err.Error())
		}
	}

	fmt.Println(users)

	orm.Execute()
	defer db.Close()
}
