package main

import "fmt"

type Mob interface {
}

type Hero struct {
	Name     string
	Level    float64
	Strength float64
	Health   float64
}

func (h *Hero) Hit(Enemyhealth float64) {
	Enemyhealth = Enemyhealth - h.Strength
}

type Enemy struct {
	Name     string
	Level    float64
	Strength float64
}

func main() {
	Ivan := Enemy{}
	Ivan.Name = "Ivan"
	fmt.Println(Ivan.Name)

}
