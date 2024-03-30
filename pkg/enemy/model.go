package pkg

type IEnemy interface {
	IsAlive() bool
	GetColumn() int
	Position() (float32, float32)

	Move(dx, dy float32, direction int) (float32, float32)
}

const (
	Yellow = "yellow"
	Green  = "green"
	Red    = "red"
)

type Enemy struct {
	Type    string
	points  int
	X       float32
	Y       float32
	Column  int
	isAlive bool
}

func NewEnemy(startX, startY float32, points, column int, enemyType string) IEnemy {
	return &Enemy{
		Type:    enemyType,
		points:  points,
		X:       startX,
		Y:       startY,
		Column:  column,
		isAlive: true,
	}
}

func (p *Enemy) Position() (float32, float32) {
	return p.X, p.Y
}

func (p *Enemy) IsAlive() bool {
	return p.isAlive
}

func (p *Enemy) GetColumn() int {
	return p.Column
}

func (p *Enemy) Move(dx, dy float32, direction int) (float32, float32) {
	if direction == -1 {
		p.X = p.X - dx
	}
	if direction == 1 {
		p.X = p.X + dx
	}
	p.Y = p.Y + dy

	return p.X, p.Y
}
