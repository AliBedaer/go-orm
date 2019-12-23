package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// Orm : base database orm
type Orm struct {
	Query    string
	Host     string
	Port     int64
	User     string
	Password string
	DBname   string
	DBDriver string
}

func (or *Orm) Init(port int64, host, user, password, dbname, dbDriver string) {
	or.Host = host
	or.Port = port
	or.User = user
	or.Password = password
	or.DBname = dbname
	or.DBDriver = dbDriver
}

func (or *Orm) Connect() *sql.DB {

	var psqlInfo string

	if or.DBDriver == "postgres" {
		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", or.Host, or.Port, or.User, or.Password, or.DBname)
	} else if or.DBDriver == "mysql" {
		// user:password@tcp(127.0.0.1:3306)/hello
		psqlInfo = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", or.User, or.Password, or.Host, or.Port, or.DBname)
	}

	db, err := sql.Open(or.DBDriver, psqlInfo)

	if err != nil {
		panic(err)
	}

	return db
}

func (or *Orm) Where(param1, operator, param2 string) {
	or.Query = fmt.Sprintf("where %s %s %s", param1, operator, param2)
}

func (or *Orm) Select(fileds, table string) *Orm {
	or.Query = fmt.Sprintf("select %s from %s", fileds, table)
	return or
}

func (or *Orm) SelectRaw() {

}

func (or *Orm) Execute() {
	or.Connect().Exec(or.Query)
}
