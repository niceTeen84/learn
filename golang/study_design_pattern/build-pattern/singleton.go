package build

import "sync"

type Tool struct {
	name,
	id string
}

var instance *Tool
// need a lock avoid multi thread context
var lock sync.Mutex

// 双重检查懒加载模式
func GetInstance() *Tool {
	if instance == nil {
		lock.Lock()
		if instance == nil {
			instance = &Tool{name: "init", id: "12345"}
		}
		lock.Unlock()
	}
	return instance
}



