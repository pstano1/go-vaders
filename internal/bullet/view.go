package bullet

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type IBulletObserver interface {
	UpdatePosition(x, y float32)
	ChangeVisibility(hide bool)
}

type BulletView struct {
	Sprite *canvas.Image
}

func NewBulletView(imagePath string, bullet IBullet) *BulletView {
	sprite := canvas.NewImageFromFile(imagePath)
	sprite.FillMode = canvas.ImageFillOriginal
	sprite.Resize(fyne.NewSize(2, 20))
	sprite.Move(fyne.NewPos(bullet.Position()))

	return &BulletView{
		Sprite: sprite,
	}
}

func (v *BulletView) UpdatePosition(x, y float32) {
	v.Sprite.Move(fyne.NewPos(x, y))
}

func (v *BulletView) ChangeVisibility(hide bool) {
	v.Sprite.Hidden = hide
}
