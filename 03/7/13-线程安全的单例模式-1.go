package main

import (
	"fmt"
	"sync"
)

//定义锁
var lock sync.Mutex

type singelton struct {}

var instance *singelton

func GetInstance() *singelton {
	//为了线程安全，增加互斥
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		return new(singelton)
	} else {
		return instance
	}
}

func (s *singelton) SomeThing() {
	fmt.Println("单例对象的某方法")
}


func main() {
	s := GetInstance()
	s.SomeThing()
}
