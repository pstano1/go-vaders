package enemy

import "sync"

type IEnemyController interface {
	Enemy() IEnemy

	Destroy()
	CheckForCollision(x, y float32) bool
	Move(dx, dy float32, direction int, wg *sync.WaitGroup)
}

type EnemyController struct {
	enemy    IEnemy
	observer IEnemyObserver
}

func NewEnemyController(enemy IEnemy, observer IEnemyObserver) IEnemyController {
	return &EnemyController{
		enemy:    enemy,
		observer: observer,
	}
}

func (c *EnemyController) Enemy() IEnemy {
	return c.enemy
}

func (c *EnemyController) Move(dx, dy float32, direction int, wg *sync.WaitGroup) {
	defer wg.Done()
	c.enemy.Move(dx, dy, direction)

	c.observer.UpdatePosition(c.enemy.Position())
}

func (c *EnemyController) Destroy() {
	c.enemy.Destroy()
	c.observer.ChangeVisibility(true)
}

func (c *EnemyController) CheckForCollision(x, y float32) bool {
	return c.enemy.CheckForCollision(x, y)
}
