package character

import (
	"context"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/x0y14/mega8/pkg/entity"
	"github.com/x0y14/mega8/pkg/game/compute"
)

type Character struct {
	count                    int
	internalCount            int
	currentMotionMaintainTo  int
	motionKindUpdateSender   chan<- entity.MotionKind
	motionKindUpdateReceiver <-chan entity.MotionKind

	Name              string
	Offset            *compute.Coordinate
	CurrentDirection  compute.Direction
	CurrentMotionKind entity.MotionKind
	MotionSet         map[entity.MotionKind]entity.Motion
}

func NewCharacter(name string, offset *compute.Coordinate, defaultDirection compute.Direction, motionSet map[entity.MotionKind]entity.Motion) *Character {
	motionKindUpdateCh := make(chan entity.MotionKind)

	return &Character{
		count:                    0,
		internalCount:            0,
		currentMotionMaintainTo:  0,
		Offset:                   offset,
		motionKindUpdateSender:   motionKindUpdateCh,
		motionKindUpdateReceiver: motionKindUpdateCh,
		Name:                     name,
		CurrentDirection:         defaultDirection,
		CurrentMotionKind:        entity.Idle,
		MotionSet:                motionSet,
	}
}

func (c *Character) SetOffset(coordinate *compute.Coordinate) {
	c.Offset = coordinate
}

func (c *Character) CurrentMotion() entity.Motion {
	return c.MotionSet[c.CurrentMotionKind]
}

func (c *Character) IsCurrentDirection(direction compute.Direction) bool {
	return c.CurrentDirection == direction
}

func (c *Character) UpdateCurrentDirection(direction compute.Direction) {
	c.CurrentDirection = direction
}

func (c *Character) UpdateCurrentDirectionIfNotSame(direction compute.Direction) {
	if c.CurrentDirection != direction {
		c.UpdateCurrentDirection(direction)
	}
}

func (c *Character) Draw(screen *ebiten.Image, offsetAdjust bool) {
	op := &ebiten.DrawImageOptions{}

	offsetX := float64(c.Offset.X)
	offsetY := float64(c.Offset.Y)
	directionX := float64(1)
	directionY := float64(1)

	if compute.Left == c.CurrentMotion().DefaultDirection && compute.Right == c.CurrentDirection {
		// モーションの向きが左なんだけど、キャラが右を向いていた場合。
		directionX *= -1.0
		// 位置調整
		if offsetAdjust {
			offsetX += float64(c.CurrentMotion().CurrentFrame().Width)
		}
	} else if compute.Right == c.CurrentMotion().DefaultDirection && compute.Left == c.CurrentDirection {
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
	c.internalCount++
	fmt.Printf("character update: %v\n", c.Name)
	cm := c.CurrentMotion()
	cm.Update()
}

func (c *Character) IsCurrentMotionKind(kind entity.MotionKind) bool {
	return c.CurrentMotionKind == kind
}

func (c *Character) UpdateCurrentMotionKind(kind entity.MotionKind) {
	c.motionKindUpdateSender <- kind
	fmt.Printf("new current motion kind: %v\n", kind)
}

func (c *Character) UpdateCurrentMotionKindIfNotSame(kind entity.MotionKind) {
	if c.CurrentMotionKind != kind {
		c.UpdateCurrentMotionKind(kind)
	}
}

func (c *Character) ListenMotionKindUpdate(ctx context.Context) {
	fmt.Printf("Start ListenMotionKindUpdate: %v\n", c.Name)
	for {
		select {
		case <-ctx.Done():
			return
		case <-c.motionKindUpdateReceiver:
			c.CurrentMotionKind = <-c.motionKindUpdateReceiver
			fmt.Printf("motion kind updated: %v\n", c.CurrentMotionKind.String())
			c.internalCount = 0
		}
	}
}

func (c *Character) IsCurrentMotionAllowInterrupt() bool {
	return c.CurrentMotion().AllowInterrupt
}

func (c *Character) SetMotionMaintainTo(totalLifetime int) {
	c.currentMotionMaintainTo = c.count + totalLifetime
}

func (c *Character) IsMotionEnded() bool {
	return c.count > c.currentMotionMaintainTo
}

func (c *Character) DoIdle() {
	c.UpdateCurrentMotionKindIfNotSame(entity.Idle)
	c.SetMotionMaintainTo(0)
}

func (c *Character) DoHide() {
	c.UpdateCurrentMotionKindIfNotSame(entity.Hide)
	c.SetMotionMaintainTo(0)
}

func (c *Character) DoWalk(direction compute.Direction) {
	if c.IsCurrentMotionKind(entity.Walk) && c.IsCurrentDirection(direction) {
		// 向きを変えたわけでも、動きを変更したわけでもないので終わり
		return
	}
	c.UpdateCurrentDirectionIfNotSame(direction)
	c.UpdateCurrentMotionKindIfNotSame(entity.Walk)
	c.SetMotionMaintainTo(0)
}

func (c *Character) DoAttack() {
	if c.IsCurrentMotionKind(entity.Attack) {
		// 攻撃中モーション中に攻撃することは現段階で想定してないので、終わり
		return
	}
	c.UpdateCurrentMotionKindIfNotSame(entity.Attack)
	c.SetMotionMaintainTo(c.CurrentMotion().TotalLifetime)
}
