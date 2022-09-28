package main

import (
	"fmt"
	"sync"
)

var (
	BookDict map[uint32]*Book
	OnlyOne  sync.Once
)

type Book struct {
	Id uint32
	Name,
	Isbn string
}

func loadBooks() {
	fmt.Println("init book dict")
	BookDict = make(map[uint32]*Book, 1)
	BookDict[1] = &Book{Id: 1, Name: "chinese", Isbn: "123456"}
	BookDict[2] = &Book{Id: 2, Name: "math", Isbn: "123456"}
	BookDict[3] = &Book{Id: 3, Name: "english", Isbn: "123456"}
	BookDict[4] = &Book{Id: 4, Name: "chemy", Isbn: "123456"}
	BookDict[5] = &Book{Id: 5, Name: "art", Isbn: "123456"}

}

func GetBookById(id uint32) *Book {
	// 这里使用 sync 包下的同步器，内部的互斥锁机制确保只执行一次
	OnlyOne.Do(loadBooks)
	return BookDict[id]
}

func main() {
	fmt.Println(GetBookById(2))
	fmt.Println(GetBookById(3))
}
