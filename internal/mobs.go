package mobs

import (
	"fmt"
)

type Hero struct {
	Live       bool
	MainPlayer bool
	Name       string
	Level      float64
	Strength   float64
	Health     float64
	Protection float64
	Money      int
	NerbEnemy  bool
}

func (h *Hero) Rename() {
	fmt.Println("Введите имя:")
	fmt.Scanln(h.Name)
}

func (h *Hero) Income(money int) {
	h.Money += money
}

func (h *Hero) Hit(enemydamage float64) {
	h.Health -= enemydamage
}

func (h *Hero) Healing(heal float64) {
	h.Health += heal
}

type Enemy struct {
	Live     bool
	Name     string
	Level    float64
	Strength float64
	Health   float64
}

func (m *Enemy) Hit(herodamage float64) {
	m.Health -= herodamage
}
