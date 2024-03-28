package pkg

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type IEnemyObserver interface {
	UpdatePosition(x, y float32)
}

type EnemyView struct {
	Sprite *canvas.Image
}

func NewEnemyView(imagePath string, enemy IEnemy) *EnemyView {
	sprite := canvas.NewImageFromFile(imagePath)
	sprite.FillMode = canvas.ImageFillOriginal
	sprite.Resize(fyne.NewSize(50, 50))
	sprite.Move(fyne.NewPos(enemy.Position()))

	return &EnemyView{
		Sprite: sprite,
	}
}

func (v *EnemyView) UpdatePosition(x, y float32) {
	v.Sprite.Move(fyne.NewPos(x, y))
}
