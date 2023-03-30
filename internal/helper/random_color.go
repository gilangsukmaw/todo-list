package helper

import (
	"gitlab.com/todo-list-app1/todo-list-backend/internal/consts"
	"math/rand"
)

func RandomColor() string {
	colorList := [25]string{
		consts.Gray,
		consts.Red,
		consts.Orange,
		consts.Amber,
		consts.Yellow,
		consts.Lime,
		consts.Green,
		consts.Emerald,
		consts.Teal,
		consts.Cyan,
		consts.Sky,
		consts.Blue,
		consts.Indigo,
		consts.Violet,
		consts.Purple,
		consts.Fuchsia,
		consts.Pink,
		consts.Rose,
	}

	randomIndex := rand.Intn(len(colorList))
	pick := colorList[randomIndex]

	return pick
}
