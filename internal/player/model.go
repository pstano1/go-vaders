package player

import "github.com/pstano1/go-vaders/internal/bullet"

type IPlayer interface {
	Position() (float32, float32)

	Shoot(appendBullet func(bullet.IBulletController)) *bullet.BulletView
	Move(dx float32, direction int) (float32, float32)
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

func (p *Player) Shoot(appendBullet func(bullet.IBulletController)) *bullet.BulletView {
	x, y := p.Position()
	b := bullet.NewBullet(x+23, y)
	v := bullet.NewBulletView("assets/bullet.png", b)
	appendBullet(bullet.NewBulletController(b, v))

	return v
}
