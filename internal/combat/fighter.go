package combat

import "github.com/VioletProto/CSC360-Final/internal/equipment"

type AttackResult struct {
	Attacker, Defender, Style string
	Damage                    int
	Defeated                  bool
}

type Fighter struct {
	Stats
	weapon   equipment.Weapon
	strategy AttackStrategy
}

func NewFighter(stats Stats, w equipment.Weapon, s AttackStrategy) *Fighter {
	return &Fighter{Stats: stats, weapon: w, strategy: s}
}

func (f *Fighter) Attack(t Target) AttackResult {
	dmg := f.strategy.Damage(f.weapon.Power(), t)
	dealt := t.TakeDamage(dmg)
	return AttackResult{
		Attacker: f.Name(),
		Defender: t.Name(),
		Style:    f.strategy.Name(),
		Damage:   dealt,
		Defeated: !t.IsAlive(),
	}
}
