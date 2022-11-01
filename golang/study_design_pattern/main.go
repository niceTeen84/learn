package main

import (
	"fmt"

	"renbw.org/learn_desin_pattern/build-pattern"
	"renbw.org/learn_desin_pattern/build-pattern/factory"
)

func main() {
	t := &build.Tool{}
	fmt.Println(t)
	ptr := &factory.Worker{Name: "cofee"}
	fmt.Println(ptr)
}
