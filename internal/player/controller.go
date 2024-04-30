package player

import (
	"fyne.io/fyne/v2"
	"github.com/pstano1/go-vaders/internal/bullet"
)

type IPlayerController interface {
	CheckForCollision(x, y float32) bool
	UpdateLifes(difference int)
	Move(direction fyne.KeyName, width float32)
	Shoot() (bullet.IBulletController, *bullet.BulletView)

	GetLifes() int
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

func (p *PlayerController) CheckForCollision(x, y float32) bool {
	isColliding := p.player.CheckForCollision(x, y)
	if isColliding {
		p.player.UpdateLifes(p.player.Lifes() - 1)
		p.player.ResetPosition()
	}
	return isColliding
}

func (c *PlayerController) Move(direction fyne.KeyName, width float32) {
	switch direction {
	case "Right":
		x, _ := c.player.Position()
		if x < width-10 {
			c.player.Move(10, 1)
		}
	case "Left":
		x, _ := c.player.Position()
		if x > 10 {
			c.player.Move(10, -1)
		}
	}

	c.observer.UpdatePosition(c.player.Position())
}

func (c *PlayerController) Shoot() (bullet.IBulletController, *bullet.BulletView) {
	return c.player.Shoot()
}

func (c *PlayerController) UpdateLifes(difference int) {
	newLives := c.player.Lifes() + difference
	if newLives < 0 {
		newLives = 0
	} else if newLives > 3 {
		newLives = 3
	}
	c.player.UpdateLifes(newLives)
}

func (c *PlayerController) GetLifes() int {
	return c.player.Lifes()
}
