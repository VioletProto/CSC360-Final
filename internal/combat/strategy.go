package combat

type AttackStrategy interface {
	Name() string
	Damage(power int, target Target) int
}

type Melee struct{}

func (Melee) Name() string                   { return "Melee" }
func (Melee) Damage(power int, t Target) int { return max(power-t.Armor(), 1) }

type Magic struct{}

func (Magic) Name() string                   { return "Melee" }
func (Magic) Damage(power int, _ Target) int { return power * 3 / 4 }
