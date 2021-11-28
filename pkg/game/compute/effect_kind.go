package compute

type EffectKind int

const (
	_ EffectKind = iota
	Move
	Heal
	Damage
)

var effectKinds = [...]string{
	Move:   "Move",
	Heal:   "Heal",
	Damage: "Damage",
}

func (e EffectKind) String() string {
	return effectKinds[e]
}
