package pkg

type IEnemyController interface {
	Move(dx float32, direction int)
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

func (c *EnemyController) Move(dx float32, direction int) {
	c.enemy.Move(dx, direction)

	c.observer.UpdatePosition(c.enemy.Position())
}
