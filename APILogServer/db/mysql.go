package db

import (
	"database/sql"
	"fmt"
	"log"
	"sixshop/apilog/configuration"

	_ "github.com/go-sql-driver/mysql"
)

var DB Db

type Db struct {
	D *sql.DB
}

func makeDbConfig() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		configuration.Conf.Mysql.User,
		configuration.Conf.Mysql.Password,
		configuration.Conf.Mysql.Host,
		configuration.Conf.Mysql.Port,
		configuration.Conf.Mysql.Name)
}

func InitDB() {
	db, err := sql.Open("mysql", makeDbConfig())
	if err != nil {
		log.Fatal(err)
	}
	DB = Db{db}
}
