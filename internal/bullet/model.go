package bullet

type IBullet interface {
	Position() (float32, float32)

	Move(dy float32) (float32, float32)
}

type Bullet struct {
	x float32
	y float32
}

func NewBullet(startX, startY float32) IBullet {
	return &Bullet{
		x: startX,
		y: startY,
	}
}

func (b *Bullet) Position() (float32, float32) {
	return b.x, b.y
}

func (b *Bullet) Move(dy float32) (float32, float32) {
	b.y = b.y - dy

	return b.x, b.y
}
