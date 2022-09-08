package main

import (
	// .env variables init first
	"context"
	"database/sql"
	"fmt"
	"local/study/sql/db"
	"log"
	"math/rand"
	"os"
	"runtime/trace"
	"time"

	"github.com/jmoiron/sqlx"
)

type Component struct {
	Id        int        `db:"id"`
	CreatedAt *time.Time `db:"created_at"`
	UpdateAt  *time.Time `db:"updated_at"`
	DelAt     *time.Time `db:"deleted_at"`
	Name      *string    `db:"name"`
	Number    int        `db:"num"`
}

func Trx() {
	db.Db.MustBegin()
	// tx.Exec()

}

func InsertMany() {
	compoents := []*Component{}
	for i := 0; i < 400; i++ {
		// 这里为了处理空值，结构体的类型定义为指针
		name, now := randomString(4), time.Now()
		elm := &Component{
			CreatedAt: &now,
			Name:      &name,
			Number:    128}
		compoents = append(compoents, elm)
	}
	s := time.Now()
	_, err := db.Db.NamedExec(`insert into components (created_at, name, num)
					values (:created_at, :name, :num)`, compoents)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("insert done cost ", (time.Now().UnixMilli() - s.UnixMilli()))
}

func ExecOtherCmd() {
	// check the version
	row := db.Db.QueryRowx(`select version() version, (now() + interval ? hour) time`, 5)
	info := R{}
	row.MapScan(info)
	fmt.Println()
}

func ModifyVariable() {
	res := db.Db.MustExec("set global innodb_flush_log_at_trx_commit = 0")
	if _, err := res.RowsAffected(); err != nil {
		fmt.Println("set global variables failed ", err.Error())
	}
	// fmt.Println(num, " row changes")
}

func ModifyContext() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, _ := db.Db.Connx(ctx)
	tx, err := conn.BeginTxx(context.Background(), &sql.TxOptions{Isolation: sql.LevelRepeatableRead, ReadOnly: false})
	if err != nil {
		log.Fatal("begin trx failed ", err.Error())
	}
	query, args, _ := sqlx.In("select * from components where id in (?)", []int{1, 3, 6})
	compos := []Component{}
	err = tx.Select(&compos, query, args...)
	if err != nil {
		log.Fatal("query rows failed ", err.Error())
	}

	fmt.Println(query, args)
	<-ctx.Done()
	tx.Commit()
}

type R map[string]interface{}

func SqlxMapScan() {
	// test cmd
	rows, _ := db.Db.Queryx("show full processlist")
	defer rows.Close()
	list := []R{}
	for rows.Next() {
		row := R{}
		rows.MapScan(row)
		list = append(list, row)
	}
	rows.Close()
	if len(list) == 0 {
		fmt.Println("scan failed")
		return
	}
}

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func main() {
	// go run main.go 2> trace.out
	// go tool trace trace.out
	// tarce info redirect to std err
	defer db.Db.Close()
	trace.Start(os.Stderr)
	defer trace.Stop()
	ModifyContext()
}
