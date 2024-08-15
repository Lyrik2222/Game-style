package mobs

type Mob interface {
	Hit()
}

type Hero struct {
	MainPlayer bool
	Name       string
	Level      float64
	Strength   float64
	Health     float64
}

func (h *Hero) Hit(enemydamage float64) {
	h.Health -= enemydamage
}

func (h *Hero) Healing(heal float64) {
	h.Health += heal
}

type Enemy struct {
	Name     string
	Level    float64
	Strength float64
	Health   float64
}

func (m *Enemy) Hit(herodamage float64) {
	m.Health -= herodamage
}
