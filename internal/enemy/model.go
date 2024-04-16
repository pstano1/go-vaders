package enemy

import "github.com/pstano1/go-vaders/internal/bullet"

type IEnemy interface {
	IsAlive() bool
	Position() (float32, float32)
	Points() int

	Destroy()
	CheckForCollision(x, y float32) bool
	Shoot() (bullet.IBulletController, *bullet.BulletView)
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
	if x >= e.X+70 && x <= e.X+50+70 && e.Y <= y && e.Y+50+70 >= y {
		isColliding = true
	}

	return isColliding
}

func (e *Enemy) Points() int {
	return e.points
}

func (e *Enemy) Shoot() (bullet.IBulletController, *bullet.BulletView) {
	x, y := e.Position()
	b := bullet.NewBullet(x+23+70, y, bullet.EnemysBullet)
	v := bullet.NewBulletView("assets/bullet.png", b)
	c := bullet.NewBulletController(b, v)

	return c, v
}
