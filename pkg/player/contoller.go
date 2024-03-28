package pkg

import "fyne.io/fyne/v2"

type IPlayerController interface {
	HandleKey(e *fyne.KeyEvent)
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

func (c *PlayerController) HandleKey(e *fyne.KeyEvent) {
	switch e.Name {
	case fyne.KeyRight:
		c.player.Move(10, 1)
	case fyne.KeyLeft:
		c.player.Move(10, -1)
	}

	c.observer.UpdatePosition(c.player.Position())
}
