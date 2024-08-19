package main

import (
	mobs "Game-style/internal"
	"fmt"
)

func main() {
	Ivan := mobs.Hero{
		MainPlayer: true,
		Level:      1.0,
	}
	Slime := mobs.Enemy{
		Name:     "Слизень",
		Level:    3.0,
		Strength: 1.0,
		Health:   5.0,
	}
	Slime.Hit(Ivan.Strength)
	fmt.Println(Slime.Health)

}
