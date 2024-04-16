package bullet

const (
	EnemysBullet  = "enemys-bullet"
	PlayersBullet = "players-bullet"
)

type IBullet interface {
	Position() (float32, float32)
	Owner() string

	Move(dy float32) (float32, float32)
}

type Bullet struct {
	x     float32
	y     float32
	owner string
}

func NewBullet(startX, startY float32, owner string) IBullet {
	return &Bullet{
		x:     startX,
		y:     startY,
		owner: owner,
	}
}

func (b *Bullet) Position() (float32, float32) {
	return b.x, b.y
}

func (b *Bullet) Move(dy float32) (float32, float32) {
	b.y = b.y - dy

	return b.x, b.y
}

func (b *Bullet) Owner() string {
	return b.owner
}
