package entities

type Animation struct {
	FrameNum      int
	NowFrameNo    int
	LifetimeCount int
	Frames        []Frame
	Direction
}

func (a *Animation) SetDirection(direction Direction) {
	a.Direction = direction
}

func (a *Animation) NowFrame() *Frame {
	return &a.Frames[a.NowFrameNo]
}

func (a *Animation) NextFrame() {
	a.NowFrameNo++
	if a.FrameNum <= a.NowFrameNo {
		a.NowFrameNo = 0
	}
}

func (a *Animation) Update() {
	a.LifetimeCount++
	if a.NowFrame().Lifetime < a.LifetimeCount {
		a.NextFrame()
		a.LifetimeCount = 0
	}
}
