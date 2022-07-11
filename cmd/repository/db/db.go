package db

import (
	"database/sql"
	"fmt"

	"github.com/armzerpa/simple-rest-api-mysql-example/cmd/config"
	_ "github.com/go-sql-driver/mysql" //mysql driver
)

//GetDatabase - returns a Database object
func InitDatabase(config config.DBConfig) (*sql.DB, error) {
	db, err := sql.Open(config.Dialect, fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", config.Username, config.Password, config.Host, config.Port, config.Name))
	if err == nil {
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(20)
		err = db.Ping()
	}
	return db, err
}

func Close(db *sql.DB) {
	if db != nil {
		defer db.Close()
	}
}
