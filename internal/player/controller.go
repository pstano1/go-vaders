package player

import (
	"fyne.io/fyne/v2"
	"github.com/pstano1/go-vaders/internal/board"
)

type IPlayerController interface {
	HandleKey(e *fyne.KeyEvent, b board.IBoard, container *fyne.Container)
}

type PlayerController struct {
	player   IPlayer
	observer IPlayerObserver
}

func NewPlayerController(player IPlayer, observer IPlayerObserver) IPlayerController {
	return &PlayerController{
		player:   player,
		observer: observer,
	}
}

func (c *PlayerController) HandleKey(e *fyne.KeyEvent, b board.IBoard, container *fyne.Container) {
	switch e.Name {
	case fyne.KeyRight:
		w, _ := b.Size()
		x, _ := c.player.Position()
		if x < w-10 {
			c.player.Move(10, 1)
		}
	case fyne.KeyLeft:
		x, _ := c.player.Position()
		if x > 10 {
			c.player.Move(10, -1)
		}
	case fyne.KeySpace:
		v := c.player.Shoot(b.AppendBullet)
		container.Add(v.Sprite)
	}

	c.observer.UpdatePosition(c.player.Position())
}
