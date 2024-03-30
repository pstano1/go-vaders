package pkg

import "sync"

type IEnemyController interface {
	Enemy() IEnemy

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
