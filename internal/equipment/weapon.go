package equipment

type Weapon interface {
	Description() string
	Power() int
}

type BasicWeapon struct {
	name  string
	power int
}

func NewSword() Weapon                       { return BasicWeapon{"Sword", 10} }
func NewBasic(name string, power int) Weapon { return BasicWeapon{name, power} }

func (w BasicWeapon) Description() string { return w.name }
func (w BasicWeapon) Power() int          { return w.power }

type Upgrade interface {
	Weapon
	Unwrap() Weapon
}

type wrapper struct{ inner Weapon }

func (u wrapper) Description() string { return u.inner.Description() }
func (u wrapper) Power() int          { return u.inner.Power() }
func (u wrapper) Unwrap() Weapon      { return u.inner }

type sharpened struct{ wrapper }

func Sharpen(w Weapon) Upgrade { return sharpened{wrapper{w}} }

func (s sharpened) Description() string { return "Sharpened" + s.inner.Description() }
func (s sharpened) Power() int          { return s.inner.Power() + 3 }

type enchanted struct{ wrapper }

func Enchant(w Weapon) Upgrade { return enchanted{wrapper{w}} }

func (e enchanted) Description() string { return e.inner.Description() + " of Flame" }
func (e enchanted) Power() int          { return e.inner.Power() + 5 }
