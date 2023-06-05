package mysql

import (
	"database/sql"
	"fmt"
	"hitss/pkg/helper/logger"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Open() *sql.DB {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_SERVER"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err := sql.Open("mysql", connString)
	if err != nil {
		logger.Write(err)
		panic(err)
	}

	if err = db.Ping(); err != nil {
		logger.Write(err)
		panic(err)
	}

	return db
}

func Close(db *sql.DB) {
	err := db.Close()
	if err != nil {
		logger.Write(err)
		panic(err)
	}
}
