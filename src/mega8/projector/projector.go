package projector

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/x0y14/mega8/src/mega8/entities"
	"github.com/x0y14/mega8/src/mega8/entities/character"
	"image"
)

func DrawCharacter(screen *ebiten.Image, entity character.Character, offsetAdjust bool) {
	op := &ebiten.DrawImageOptions{}

	offsetX := float64(entity.OffsetX)
	offsetY := float64(entity.OffsetY)

	dX := 1.0
	dY := 1.0

	if entities.Left == entity.NowMotion().Direction && entities.Right == entity.Direction {
		// モーションの向きが左なんだけど、キャラが右を向いていた場合。
		dX *= -1.0
		// 位置調整
		if offsetAdjust {
			offsetX += float64(entity.NowMotion().NowFrame().Width)
		}
	} else if entities.Right == entity.NowMotion().Direction && entities.Left == entity.Direction {
		// モーションの向きが右なんだけど、キャラが左を向いていた場合。
		dX *= -1.0
		if offsetAdjust {
			offsetX -= float64(entity.NowMotion().NowFrame().Width)
		}
	}

	op.GeoM.Scale(dX, dY)
	op.GeoM.Translate(offsetX, offsetY)

	nowMot := entity.NowMotion()
	nowFrame := nowMot.Frames[nowMot.NowFrameNo]
	nowSheet := nowFrame.Sheet
	img := nowSheet.SubImage(image.Rect(
		nowFrame.OriginX,
		nowFrame.OriginY,
		nowFrame.OriginX+nowFrame.Width,
		nowFrame.OriginY+nowFrame.Height)).(*ebiten.Image)

	screen.DrawImage(img, op)
}

func DrawCharacterWithOp(screen *ebiten.Image, entity character.Character, options *ebiten.DrawImageOptions) {
	nowMot := entity.NowMotion()
	nowFrame := nowMot.Frames[nowMot.NowFrameNo]
	nowSheet := nowFrame.Sheet
	img := nowSheet.SubImage(image.Rect(
		nowFrame.OriginX,
		nowFrame.OriginY,
		nowFrame.OriginX+nowFrame.Width,
		nowFrame.OriginY+nowFrame.Height)).(*ebiten.Image)

	screen.DrawImage(img, options)
}
