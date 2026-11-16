package combat

type Target interface {
	Name() string
	Armor() int
	TakeDamage(amount int) int
	IsAlive() bool
}

type Stats struct {
	name      string
	hp, maxHP int
	armor     int
}

func NewStats(name string, hp, armor int) Stats {
	return Stats{name: name, hp: hp, maxHP: hp, armor: armor}
}

func (s *Stats) Name() string  { return s.name }
func (s *Stats) HP() int       { return s.hp }
func (s *Stats) MaxHP() int    { return s.maxHP }
func (s *Stats) Armor() int    { return s.armor }
func (s *Stats) IsAlive() bool { return s.hp > 0 }

func (s *Stats) TakeDamage(amount int) int {
	amount = min(max(amount, 0), s.hp)
	s.hp -= amount
	return amount
}

func (s *Stats) Heal(amount int) {
	s.hp = min(s.hp+amount, s.maxHP)
}
