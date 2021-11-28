package character

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/x0y14/mega8/pkg/entity"
	"github.com/x0y14/mega8/pkg/game/compute"
)

type Character struct {
	count  int
	Offset *compute.Coordinate
	compute.Direction
	CurrentMotionKind entity.MotionKind
	MotionSet         map[entity.MotionKind]entity.Motion
}

func NewCharacter(offset *compute.Coordinate, defaultDirection compute.Direction, motionSet map[entity.MotionKind]entity.Motion) *Character {
	return &Character{
		count:             0,
		Offset:            offset,
		Direction:         defaultDirection,
		CurrentMotionKind: entity.Idle,
		MotionSet:         motionSet,
	}
}

func (c *Character) CurrentMotion() entity.Motion {
	return c.MotionSet[c.CurrentMotionKind]
}

func (c *Character) Draw(screen *ebiten.Image, offsetAdjust bool) {
	op := &ebiten.DrawImageOptions{}

	offsetX := float64(c.Offset.X)
	offsetY := float64(c.Offset.Y)
	directionX := float64(1)
	directionY := float64(1)

	if compute.Left == c.CurrentMotion().DefaultDirection && compute.Right == c.Direction {
		// モーションの向きが左なんだけど、キャラが右を向いていた場合。
		directionX *= -1.0
		// 位置調整
		if offsetAdjust {
			offsetX += float64(c.CurrentMotion().CurrentFrame().Width)
		}
	} else if compute.Right == c.CurrentMotion().DefaultDirection && compute.Left == c.Direction {
		directionX *= -1.0
		if offsetAdjust {
			offsetX -= float64(c.CurrentMotion().CurrentFrame().Width)
		}
	}

	op.GeoM.Scale(directionX, directionY)
	op.GeoM.Translate(offsetX, offsetY)

	frame := c.CurrentMotion().CurrentFrame()
	screen.DrawImage(frame.Trim(), op)

	return
}

func (c *Character) Update() {
	c.count++
	//fmt.Printf("character-count: %v\n", c.count)
}
