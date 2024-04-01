package enemy

type IEnemy interface {
	IsAlive() bool
	GetColumn() int
	Position() (float32, float32)

	Destroy()
	CheckForCollision(x, y float32) bool
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

func (e *Enemy) Position() (float32, float32) {
	return e.X, e.Y
}

func (e *Enemy) IsAlive() bool {
	return e.isAlive
}

func (e *Enemy) GetColumn() int {
	return e.Column
}

func (e *Enemy) Move(dx, dy float32, direction int) (float32, float32) {
	if direction == -1 {
		e.X = e.X - dx
	}
	if direction == 1 {
		e.X = e.X + dx
	}
	e.Y = e.Y + dy

	return e.X, e.Y
}

func (e *Enemy) Destroy() {
	e.isAlive = false
}

func (e *Enemy) CheckForCollision(x, y float32) bool {
	var isColliding bool
	if !e.isAlive {
		return false
	}
	if x >= e.X && x <= e.X+50 && e.Y <= y && e.Y+50 >= y {
		isColliding = true
	}

	return isColliding
}
