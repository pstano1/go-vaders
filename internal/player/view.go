package player

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type IPlayerObserver interface {
	UpdatePosition(x, y float32)
}

type PlayerView struct {
	Sprite *canvas.Image
}

func NewPlayerView(imagePath string, player IPlayer) *PlayerView {
	sprite := canvas.NewImageFromFile(imagePath)
	sprite.FillMode = canvas.ImageFillOriginal
	sprite.Resize(fyne.NewSize(50, 50))
	sprite.Move(fyne.NewPos(player.Position()))

	return &PlayerView{
		Sprite: sprite,
	}
}

func (v *PlayerView) UpdatePosition(x, y float32) {
	v.Sprite.Move(fyne.NewPos(x, y))
}
