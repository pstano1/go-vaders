package pkg

type IPlayer interface {
	Position() (float32, float32)

	Move(dx float32, direction int) (float32, float32)
}

type Player struct {
	Lifes int
	X     float32
	Y     float32
}

func NewPlayer(startX, startY float32) IPlayer {
	return &Player{
		Lifes: 3,
		X:     startX,
		Y:     startY,
	}
}

func (p *Player) Position() (float32, float32) {
	return p.X, p.Y
}

func (p *Player) Move(dx float32, direction int) (float32, float32) {
	if direction == -1 {
		p.X = p.X - dx
	}
	if direction == 1 {
		p.X = p.X + dx
	}

	return p.X, p.Y
}
