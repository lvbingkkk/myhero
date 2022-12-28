package main

import (
	"image/color"
)

type Hero struct {
	Name    string  `json:"name"`
	Role    string  `json:"role"`
	Hp      float32 `json:"hp"`      //生命
	Mp      float32 `json:"mp"`      //魔法值
	Level   int     `json:"level"`   //等级
	Damage  float32 `json:"damage"`  //伤害
	Defence float32 `json:"defence"` //防御
}

type Role struct {
	Profession string //职业
	Color      color.RGBA
	Hp         float32 //该职业基础血量
	Mp         float32
	Level      int
	Damage     float32
	Defence    float32
}

// func main() {
// 	zhanShi := Role{"zhanShi", color.RGBA{220, 100, 100, 100}}
// 	fmt.Println(zhanShi.color)
// 	colour.Red("hello")
// }

var zhanShi Role = Role{"战士", color.RGBA{140, 120, 100, 100}, 5000, 100, 1, 120, 150}
var siQi Role = Role{"死骑", color.RGBA{120, 20, 40, 100}, 4800, 100, 1, 140, 130}

var shuShi Role = Role{"术士", color.RGBA{90, 90, 150, 100}, 2800, 220, 1, 240, 50}
var faShi Role = Role{"法师", color.RGBA{70, 160, 175, 100}, 3000, 200, 1, 220, 60}
var lieRen Role = Role{"猎人", color.RGBA{150, 180, 100, 100}, 2900, 100, 1, 180, 60}
var lieMo Role = Role{"猎魔", color.RGBA{150, 70, 160, 100}, 3000, 100, 1, 200, 60}

var muShi Role = Role{"牧师", color.RGBA{188, 188, 188, 100}, 2800, 220, 1, 220, 60}
var xiaoDe Role = Role{"小德", color.RGBA{200, 120, 40, 100}, 6000, 150, 1, 100, 100}
var shengQi Role = Role{"圣骑", color.RGBA{170, 100, 130, 100}, 4000, 200, 1, 160, 80}

var rolePro = []Role{zhanShi, faShi, muShi, lieRen, xiaoDe, siQi, shengQi, shuShi, lieMo}
