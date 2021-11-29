package display

import (
	"github.com/x0y14/mega8/pkg/game/compute"
)

type Animation struct {
	currentFrameNo int
	lifetimeCount  int

	FrameNum         int
	Frames           []Frame
	DefaultDirection compute.Direction
	TotalLifetime    int
}

func NewAnimation(frames []Frame, defaultDirection compute.Direction) *Animation {
	totalLifetime := 0
	for _, frame := range frames {
		totalLifetime += frame.Lifetime
	}

	return &Animation{
		currentFrameNo:   0,
		lifetimeCount:    0,
		FrameNum:         len(frames),
		Frames:           frames,
		DefaultDirection: defaultDirection,
		TotalLifetime:    totalLifetime,
	}
}

func (a *Animation) CurrentFrame() Frame {
	return a.Frames[a.currentFrameNo]
}

func (a *Animation) MoveNextFrame() {
	a.currentFrameNo++
	if a.currentFrameNo >= a.FrameNum {
		a.currentFrameNo = 0
	}
}

func (a *Animation) Update() {
	//fmt.Printf("frame no: %v, lifetimeCount: %v\n", a.currentFrameNo, a.lifetimeCount)
	a.lifetimeCount++
	if a.lifetimeCount >= a.CurrentFrame().Lifetime {
		a.lifetimeCount = 0
		a.MoveNextFrame()
	}
}
