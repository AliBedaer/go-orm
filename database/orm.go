package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Orm : base database orm
type Orm struct {
	query    string
	host     string
	port     int8
	user     string
	password string
	dbname   string
	dbDriver string
}

func (or *Orm) Init(port int8, host, user, password, dbname, dbDriver string) {
	or.host = host
	or.port = port
	or.user = user
	or.password = password
	or.dbname = dbname
	or.dbDriver = dbDriver
}

func (or *Orm) Connect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", or.host, or.port, or.user, or.password, or.dbname)

	db, err := sql.Open(or.dbDriver, psqlInfo)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	return db
}

func (or *Orm) Where(param1, operator, param2 string) {
	or.query = fmt.Sprintf("where %s %s %s", param1, operator, param2)
}

func (or *Orm) Select(fileds, table string) string {
	return fmt.Sprintf("select %s from %s", fileds, table)
}

func (or *Orm) Execute() {
	or.Connect().Exec(or.query)
}
