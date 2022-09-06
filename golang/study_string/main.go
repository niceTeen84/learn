package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

func Empty(s string) bool {
	ret := true
	if s = strings.Trim(s, " "); len(s) > 0 {
		ret = false
	}
	return ret
}

func Join() {
	fmt.Println(strings.Join([]string{"aaaa", "bbbb"}, ","))
	s := "keep calm and carry on"
	fmt.Println(s[3:])
	fmt.Println(strings.ToUpper(s))
	fmt.Println(strings.ToLower(s))
	fmt.Println(strings.Count(s, "e"))
}

func string2Int(str string) int {
	// ret, _ = strconv.Atoi(s)
	// 第二个参数是进制 2 8 10 16
	// 第三个参数是适配的位大小 0 ，8 ，16， 32 64
	r, _ := strconv.ParseInt(str, 10, 0)
	// 10 进制转换, bitszie 0
	// Atoi array to integer
	// Itoa Integer to array
	// iota 希腊字母 go 中只能用于常量
	// ref https://en.wikipedia.org/wiki/Iota
	fmt.Println(strconv.Itoa(123))
	if v, err := strconv.Atoi(str); err == nil {
		fmt.Printf("%T %v \n", v, v)
	}

	// E e 代表十进制
	// float64 转换 string
	strFloat := strconv.FormatFloat(math.Pi, 'E', -1, 64)
	if val, err := strconv.ParseFloat(strFloat, 32); err == nil {
		fmt.Printf("%T %v \n", val, val)
	}
	// int64 转为 int32 会被截取
	return int(r)
}

func Calc() {
	// 精确计算浮点数
	// 需要引入类似 Java BigDecimal 一样的三方库
	a, _ := decimal.NewFromString("0.1")
	b, _ := decimal.NewFromString("0.2")
	fmt.Println(a.Add(b))
	c, d := 0.1, 0.2
	fmt.Println(c + d)
}

func main() {
	str := "hello 你好"
	// golang 默认采用 UTF-8 编码, 一个汉字占用 3 个字节
	fmt.Println("string length is ", len(str))
	// 测试 emoji, 通常 emoji 占用 4 个字节
	// 异常 emoji 👩‍❤️‍💋‍👨
	str = "hello 😊 "
	fmt.Println("string length is ", len(str))
	result := string2Int("123456")
	fmt.Printf("%T %d \n", result, result)
	Join()
	Calc()
}
