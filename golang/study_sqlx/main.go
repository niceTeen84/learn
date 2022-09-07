package main

import (
	// .env variables init first
	_ "github.com/joho/godotenv/autoload"
	"fmt"
	"local/study/sql/db"
	"math/rand"
	"os"
	"runtime/trace"
	"time"
)

type Component struct {
	Id        int       `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdateAt  time.Time `db:"updated_at"`
	Name      string    `db:"name"`
	Number    int       `db:"num"`
}

func Trx() {
	db.Db.MustBegin()
	// tx.Exec()

}

func InsertMany() {
	compoents := []*Component{}
	for i := 0; i < 400; i++ {
		elm := &Component{
			CreatedAt: time.Now(),
			Name:      randomString(4),
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
	trace.Start(os.Stderr)
	defer trace.Stop()
	ExecOtherCmd()
}
