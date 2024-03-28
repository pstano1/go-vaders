package pkg

type IEnemy interface {
	Position() (float32, float32)

	Move(dx float32, direction int) (float32, float32)
}

const (
	Yellow = "yellow"
	Green  = "green"
	Red    = "red"
)

type Enemy struct {
	Type   string
	points int
	x      float32
	y      float32
}

func NewEnemy(startX, startY float32, points int, enemyType string) IEnemy {
	return &Enemy{
		Type:   enemyType,
		points: points,
		x:      startX,
		y:      startY,
	}
}

func (p *Enemy) Position() (float32, float32) {
	return p.x, p.y
}

func (p *Enemy) Move(dx float32, direction int) (float32, float32) {
	if direction == -1 {
		p.x = p.x - dx
	}
	if direction == 1 {
		p.x = p.x + dx
	}

	return p.x, p.y
}
