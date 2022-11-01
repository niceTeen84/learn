package build

import (
	"fmt"
	"sync"
)

var _instance *Tool
var once sync.Once

// sync once 多线程下没有并发问题
// 只被加载一次，并且是懒加载
func LazyInit() *Tool {
	once.Do(func() {
		_instance = &Tool{name: "lazy", id: "12345"}
	})
	return _instance
}

func Iter() {
	var list []*Tool
	list = append(list, &Tool{}, &Tool{})
	for _, v := range list {
		fmt.Println(v)
	}
}
