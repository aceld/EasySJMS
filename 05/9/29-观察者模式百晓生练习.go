package main

import "fmt"

/*
               百晓生
	[丐帮]               [明教]
    洪七公               张无忌
    黄蓉					韦一笑
    乔峰				    金毛狮王
*/

const (
	PGaiBang string = "丐帮"
	PMingJiao string = "明教"
)


//-------- 抽象层 -------
type Listener interface {
	//当同伴被揍了该怎么办
	OnFriendBeFight()
}

type Notifier interface {
	//添加观察者
	AddListener(listener Listener)
	//删除观察者
	RemoveListener(listener Listener)
	//通知广播
	Notify()
}


//-------- 实现层 -------
//英雄(Listener)
type Hero struct {
	Name string
	Party string
}

func (hero *Hero) Fight(another *Hero, notify Notifier) {
	fmt.Println(hero.Title(), " 将 ", another.Title(), " 揍了，这个消息让交互百晓生知道了...")
}

func (hero *Hero) Title() string {
	return fmt.Sprintf("[%s]:%s", hero.Party, hero.Name)
}

func (hero *Hero) OnFriendBeFight() {

}

//百晓生(Nofifier)
type BaiXiao struct {
	heroList []Listener
}

//添加观察者
func (b *BaiXiao) AddListener(listener Listener) {
	b.heroList = append(b.heroList, listener)
}

//删除观察者
func (b *BaiXiao) RemoveListener(listener Listener) {
	for index, l := range b.heroList {
		//找到要删除的元素位置
		if listener == l {
			//将删除的点前后的元素链接起来
			b.heroList = append(b.heroList[:index], b.heroList[index+1:]...)
			break
		}
	}
}

//通知广播
func (b *BaiXiao) Notify() {
	for _, listener := range b.heroList {
		//依次调用全部观察的具体动作
		listener.OnFriendBeFight()
	}
}



func main() {

}
