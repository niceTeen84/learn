package main

import "fmt"

type Person struct {
	Name string
	Age  uint8
}

func main() {
	caseOne()
	caseTwo()
	// 两次结果如下
	// case one slice addr 0xc000114480 first element addr 0xc000004078 
	// BypassSliceOne slice addr 0xc000114480 first element addr 0xc0000040a8
	// slice addr 0xc00005c250 first element addr 0xc0000040d8 
	// BypassSliceTwo slice addr 0xc00005c250 first element addr 0xc0000040d8

	// slice 声明后就是一个引用类型不需要再次取址
	// 如果 slice 元素类型不是指针类型，参数传递时会发生拷贝，通常元素类型为引用类型
}

func caseOne() {
	var slice []Person
	slice = append(slice, Person{Name: "tom", Age: 18}, Person{Name: "jerry", Age: 17})
	elm := slice[0]
	pts := &(slice[0])
	// the same output
	fmt.Println(elm.Age)
	fmt.Println(pts.Age)
	fmt.Println((*pts).Age)

	fmt.Printf("case one slice addr %p first element addr %p \n", slice, &elm)
	BypassSliceOne(slice)
}

func BypassSliceOne(slice []Person) {
	elm := slice[0]
	fmt.Printf("BypassSliceOne slice addr %p first element addr %p\n", slice, &elm)
}

func caseTwo() {
	var slice []*Person
	slice = append(slice, &Person{Name: "tom", Age: 18}, &Person{Name: "jerry", Age: 17})
	elm := slice[0]
	fmt.Printf("slice addr %p first element addr %p \n", slice, elm)
	BypassSliceTwo(slice)
}

func BypassSliceTwo(slice []*Person) {
	elm := slice[0]
	fmt.Printf("BypassSliceTwo slice addr %p first element addr %p\n", slice, elm)
}


