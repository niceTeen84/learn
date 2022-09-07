package db

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	USER_NAME = "USER_NAME"
	PASSWORD  = "PASSWORD"
	DATABASE  = "DATABASE"
	PORT      = "PORT"
	HOST      = "HOST"
)

var Db *sqlx.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		os.Getenv(USER_NAME),
		os.Getenv(PASSWORD),
		os.Getenv(HOST),
		os.Getenv(PORT),
		os.Getenv(DATABASE))
	db, err := sqlx.Connect("mysql", dsn)

	if err != nil {
		log.Fatal("init database failed \n", err.Error())
	}

	db.SetMaxIdleConns(32)
	db.SetMaxOpenConns(256)
	db.SetConnMaxLifetime(30 * time.Second)
	db.SetConnMaxIdleTime(10 * time.Second)

	if err := db.Ping(); err == nil {
		fmt.Println("db init success")
	}

	Db = db
}
