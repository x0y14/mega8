package projector

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/x0y14/mega8/src/mega8/character"
	"image"
)

func DrawCharacter(screen *ebiten.Image, entity character.Character, flipHorizon bool) {
	op := &ebiten.DrawImageOptions{}
	if flipHorizon {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(entity.OffsetX)+float64(entity.NowMotion().NowFrame().Width), float64(entity.OffsetY))
	} else {
		op.GeoM.Translate(float64(entity.OffsetX), float64(entity.OffsetY))
	}

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
