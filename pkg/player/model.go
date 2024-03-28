package pkg

type IPlayer interface {
	Position() (float32, float32)

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
