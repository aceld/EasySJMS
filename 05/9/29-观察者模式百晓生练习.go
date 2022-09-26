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
	OnFriendBeFight(event *Event)
	GetName() string
	GetParty() string
	Title() string
}

type Notifier interface {
	//添加观察者
	AddListener(listener Listener)
	//删除观察者
	RemoveListener(listener Listener)
	//通知广播
	Notify(event *Event)
}

type Event struct {
	Noti Notifier //被知晓的通知者
	One Listener  //事件主动发出者
	Another Listener //时间被动接收者
	Msg string		//具体消息
}


//-------- 实现层 -------
//英雄(Listener)
type Hero struct {
	Name string
	Party string
}

func (hero *Hero) Fight(another Listener, baixiao Notifier) {
	msg := fmt.Sprintf("%s 将 %s 揍了...", hero.Title(), another.Title(),)

	//生成事件
	event := new(Event)
	event.Noti = baixiao
	event.One = hero
	event.Another = another
	event.Msg = msg

	baixiao.Notify(event)
}

func (hero *Hero) Title() string {
	return fmt.Sprintf("[%s]:%s", hero.Party, hero.Name)
}

func (hero *Hero) OnFriendBeFight(event *Event) {
	//判断是否为当事人
	if hero.Name == event.One.GetName() || hero.Name == event.Another.GetName() {
		return
	}

	//本帮派同伴将其他门派揍了，要拍手叫好!
	if hero.Party == event.One.GetParty() {
		fmt.Println(hero.Title(), "得知消息，拍手叫好！！！")
		return
	}

	//本帮派同伴被其他门派揍了，要主动报仇反击!
	if hero.Party == event.Another.GetParty() {
		fmt.Println(hero.Title(), "得知消息，发起报仇反击！！！")
		hero.Fight(event.One, event.Noti)
		return
	}
}

func (hero *Hero) GetName() string {
	return hero.Name
}

func (hero *Hero) GetParty() string {
	return hero.Party
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
func (b *BaiXiao) Notify(event *Event) {
	fmt.Println("【世界消息】 百晓生广播消息: ", event.Msg)
	for _, listener := range b.heroList {
		//依次调用全部观察的具体动作
		listener.OnFriendBeFight(event)
	}
}

func main() {
	hero1 := Hero{
		"黄蓉",
		PGaiBang,
	}

	hero2 := Hero{
		"洪七公",
		PGaiBang,
	}

	hero3 := Hero{
		"乔峰",
		PGaiBang,
	}

	hero4 := Hero{
		"张无忌",
		PMingJiao,
	}

	hero5 := Hero{
		"韦一笑",
		PMingJiao,
	}

	hero6 := Hero{
		"金毛狮王",
		PMingJiao,
	}

	baixiao := BaiXiao{}

	baixiao.AddListener(&hero1)
	baixiao.AddListener(&hero2)
	baixiao.AddListener(&hero3)
	baixiao.AddListener(&hero4)
	baixiao.AddListener(&hero5)
	baixiao.AddListener(&hero6)

	fmt.Println("武林一片平静.....")
	hero1.Fight(&hero5, &baixiao)
}
