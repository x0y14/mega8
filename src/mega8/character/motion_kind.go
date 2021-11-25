package character

type MotionKind int

const (
	_          MotionKind = iota
	Idle                  // 待機
	Walk                  // 歩く
	Run                   // 走る
	Jump                  // ジャンプ
	DoubleJump            // 二段ジャンプ

	Attack        // 攻撃
	RunningAttack // 走りながら攻撃

	Hurt // ダメージ受けた
	Die  // 死亡
)

var motionKinds = [...]string{
	Idle:          "Idle",
	Walk:          "Walk",
	Run:           "Run",
	Jump:          "Jump",
	DoubleJump:    "DoubleJump",
	Attack:        "Attack",
	RunningAttack: "RunningAttack",
	Hurt:          "Hurt",
	Die:           "Die",
}

func (mk MotionKind) String() string {
	return motionKinds[mk]
}
