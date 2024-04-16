package player

import "github.com/pstano1/go-vaders/internal/bullet"

type IPlayer interface {
	Position() (float32, float32)
	Lifes() int

	Shoot() (bullet.IBulletController, *bullet.BulletView)
	Move(dx float32, direction int) (float32, float32)
	CheckForCollision(x, y float32) bool
	UpdateLifes(newValue int)
	ResetPosition()
}

type Player struct {
	lifes int
	x     float32
	y     float32
}

func NewPlayer(startX, startY float32) IPlayer {
	return &Player{
		lifes: 3,
		x:     startX,
		y:     startY,
	}
}

func (p *Player) Position() (float32, float32) {
	return p.x, p.y
}

func (p *Player) Move(dx float32, direction int) (float32, float32) {
	if direction == -1 {
		p.x = p.x - dx
	}
	if direction == 1 {
		p.x = p.x + dx
	}

	return p.x, p.y
}

func (p *Player) Shoot() (bullet.IBulletController, *bullet.BulletView) {
	x, y := p.Position()
	b := bullet.NewBullet(x+23+70, y, bullet.PlayersBullet)
	v := bullet.NewBulletView("assets/bullet.png", b)
	c := bullet.NewBulletController(b, v)

	return c, v
}

func (p *Player) Lifes() int {
	return p.lifes
}

func (p *Player) CheckForCollision(x, y float32) bool {
	if x >= p.x && x <= p.x+50 && p.y <= y && p.y >= y {
		return true
	}

	return false
}

func (p *Player) UpdateLifes(newValue int) {
	p.lifes = newValue
}

func (p *Player) ResetPosition() {
	p.x = 375
}
