/*
 * @Author: bowen ren armyknife84@163.com
 * @Date: 2022-10-10 15:58:10
 * @LastEditors: bowen ren armyknife84@163.com
 * @LastEditTime: 2022-10-25 11:35:42
 * @FilePath: \study_cmd\main_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"database/sql"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type Men struct {
	Name   sql.NullString `json:"name"`
	Age    sql.NullInt16  `json:"age"`
	Salary float32        `json:"salary"`
}

type MenVo struct {
	Name string `json:"name"`
	Age  int16  `json:"age"`
}

func (men *Men) ToVo() (ret *MenVo) {
	ret = &MenVo{}
	if men.Name.Valid {
		ret.Name = men.Name.String
	}
	if men.Age.Valid {
		ret.Age = men.Age.Int16
	}
	return
}

type ISOfmt time.Time

func TestCase(t *testing.T) {
	// 测试 am struct
	corrds := []struct {
		x,
		y,
		z float32
	}{
		{1., 2., 3.},
		{4., 5., 6.},
		{7., 8., 9.},
		{10., 11., 12.},
	}

	for _, elm := range corrds {
		fmt.Println(elm.x)
	}
}

func TestCaseTwo(t *testing.T) {
	men := Men{Name: sql.NullString{String: "tom", Valid: true}}
	vo := men.ToVo()
	byts, _ := json.Marshal(vo)
	fmt.Println(string(byts))
}

//go:embed templates/index.html
var template string

func TestStartEngine(t *testing.T) {
	engine := startEngine()
	record := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	engine.ServeHTTP(record, req)
	assert.Equal(t, 200, record.Code)
	assert.Equal(t, "PONG", record.Body.String())
	fmt.Println(template)
	re, err := regexp.Compile("^/[\\d]{3}-[\\d]{4}-[\\d]{4}/$")
	if err != nil {
		log.Fatal("compile regex failed ")
	}
	fmt.Println(re)
}

func BenchmarkExample(b *testing.B) {
	// go tool pprof -http=:8080 heap.out
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive("i am a benchmark")
	}
}
