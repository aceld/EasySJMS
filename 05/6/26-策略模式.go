package main

import "fmt"

//武器策略(抽象的策略)
type WeaponStrategy interface {
	UseWeapon()  //使用武器
}


//具体的策略
type Ak47 struct {}

func (ak *Ak47) UseWeapon() {
	fmt.Println("使用Ak47 去战斗")
}

//具体的策略
type Knife struct {}

func (k *Knife) UseWeapon() {
	fmt.Println("使用匕首 去战斗")
}

//环境类
type Hero struct {
	strategy WeaponStrategy //拥有一个抽象的策略
}

//设置一个策略
func (h *Hero) SetWeaponStrategy(s WeaponStrategy) {
	h.strategy = s
}

func (h *Hero) Fight() {
	h.strategy.UseWeapon() //调用策略
}

func main() {
	hero := Hero{}
	//更换策略1
	hero.SetWeaponStrategy(new(Ak47))
	hero.Fight()

	hero.SetWeaponStrategy(new(Knife))
	hero.Fight()
}
