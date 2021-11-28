package display

import "github.com/x0y14/mega8/pkg/game/compute"

type Animation struct {
	currentFrameNo int
	lifetimeCount  int

	FrameNum         int
	Frames           []Frame
	DefaultDirection compute.Direction
}

func NewAnimation(frames []Frame, defaultDirection compute.Direction) *Animation {
	return &Animation{
		currentFrameNo:   0,
		lifetimeCount:    0,
		FrameNum:         len(frames),
		Frames:           frames,
		DefaultDirection: defaultDirection,
	}
}

func (a *Animation) CurrentFrame() Frame {
	return a.Frames[a.currentFrameNo]
}
