package main

import "fmt"

func main() {
	str := "hello 你好"
	// golang 默认采用 UTF-8 编码, 一个汉字占用 3 个字节
	fmt.Println("string length is ", len(str))
	// 测试 emoji, 通常 emoji 占用 4 个字节
	// 异常 emoji 👩‍❤️‍💋‍👨
	str = "hello 😊 "
	fmt.Println("string length is ", len(str))
}
