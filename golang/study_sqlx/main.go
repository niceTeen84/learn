package main

import (
	// .env variables init first
	_ "github.com/joho/godotenv/autoload"
	"fmt"
	"os"
	"local/study/sql/db"
	"math/rand"
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
		compoents = append(compoents, &Component{CreatedAt: time.Now(), Name: randomString(4), Number: 128})
	}
	s := time.Now()
	_, err := db.Db.NamedExec(`insert into components (created_at, name, num)
					values (:created_at, :name, :num)`, compoents)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("insert done cost ", (time.Now().UnixMilli() - s.UnixMilli()))
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
	trace.Start(os.Stderr)
    defer trace.Stop()
	InsertMany()
}
