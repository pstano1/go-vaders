package bullet

import "sync"

type IBulletController interface {
	Bullet() IBullet

	Destroy()
	Move(dy float32, wg *sync.WaitGroup)
}

type BulletController struct {
	bullet   IBullet
	observer IBulletObserver
}

func NewBulletController(bullet IBullet, observer IBulletObserver) IBulletController {
	return &BulletController{
		bullet:   bullet,
		observer: observer,
	}
}

func (c *BulletController) Bullet() IBullet {
	return c.bullet
}

func (c *BulletController) Move(dy float32, wg *sync.WaitGroup) {
	defer wg.Done()
	c.bullet.Move(dy)

	c.observer.UpdatePosition(c.bullet.Position())
}

func (c *BulletController) Destroy() {
	c.observer.ChangeVisibility(true)
}
