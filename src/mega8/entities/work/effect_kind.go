package work

type EffectKind int

const (
	_ EffectKind = iota
	Move
	Damage
	Attack
)
