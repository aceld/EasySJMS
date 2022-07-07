package main

import "fmt"

type SubSystemA struct {}

func (sa *SubSystemA) MethodA() {
	fmt.Println("子系统方法A")
}

type SubSystemB struct {}

func (sb *SubSystemB) MethodB() {
	fmt.Println("子系统方法B")
}

type SubSystemC struct {}

func (sc *SubSystemC) MethodC() {
	fmt.Println("子系统方法C")
}

type SubSystemD struct {}

func (sd *SubSystemD) MethodD() {
	fmt.Println("子系统方法D")
}

//外观模式，提供了一个外观类， 简化成一个简单的接口供使用
type Facade struct {
	a *SubSystemA
	b *SubSystemB
	c *SubSystemC
	d *SubSystemD
}

func (f *Facade) MethodOne() {
	f.a.MethodA()
	f.b.MethodB()
}


func (f *Facade) MethodTwo() {
	f.c.MethodC()
	f.d.MethodD()
}

func main() {
	//如果不用外观模式实现MethodA() 和 MethodB()
	sa := new(SubSystemA)
	sa.MethodA()
	sb := new(SubSystemB)
	sb.MethodB()


	fmt.Println("-----------")
	//使用外观模式
	f := Facade{
		a: new(SubSystemA),
		b: new(SubSystemB),
		c: new(SubSystemC),
		d: new(SubSystemD),
	}

	//调用外观包裹方法
	f.MethodOne()
}
