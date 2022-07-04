package main

import "fmt"

//抽象主题
type BeautyWoman interface {
	//对男人抛媚眼
	MakeEyesWithMan()
	//和男人浪漫的约会
	HappyWithMan()
}


//具体主题
type PanJinLian struct {}

//对男人抛媚眼
func (p *PanJinLian) MakeEyesWithMan() {
	fmt.Println("潘金莲对本官抛了个媚眼")
}

//和男人浪漫的约会
func (p *PanJinLian) HappyWithMan() {
	fmt.Println("潘金莲和本官共度了浪漫的约会。")
}

//代理中介人, 王婆
type WangPo struct {
	woman BeautyWoman
}

func NewProxy(woman BeautyWoman) BeautyWoman {
	return &WangPo{woman}
}

//对男人抛媚眼
func (p *WangPo) MakeEyesWithMan() {
	p.woman.MakeEyesWithMan()
}

//和男人浪漫的约会
func (p *WangPo) HappyWithMan() {
	p.woman.HappyWithMan()
}


//西门大官人
func main() {
	//大官人想找金莲，让王婆来安排
	wangpo := NewProxy(new(PanJinLian))
	//王婆命令潘金莲抛媚眼
	wangpo.MakeEyesWithMan()
	//王婆命令潘金莲和西门庆约会
	wangpo.HappyWithMan()
}
