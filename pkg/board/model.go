package pkg

import (
	"fyne.io/fyne/v2"
	e "github.com/pstano1/go-vaders/pkg/enemy"
)

type IBoard interface {
}

type Board struct {
	Enemies []e.IEnemyController
}

func NewBoard(c *fyne.Container) IBoard {
	enemies := make([]e.IEnemyController, 55)
	red := 22
	yellow := 22
	green := 11
	counter := 5

	for i := 0; i < len(enemies); i++ {
		if i%11 == 0 {
			counter--
		}
		var enemy e.IEnemy
		var view *e.EnemyView

		if green != 0 {
			enemy = e.NewEnemy(float32(i%11*60), float32(counter*50), 40, "green")
			view = e.NewEnemyView("assets/green.png", enemy)
			green--
		} else if yellow != 0 {
			enemy = e.NewEnemy(float32(i%11*60), float32(counter*50), 20, "yellow")
			view = e.NewEnemyView("assets/yellow.png", enemy)
			yellow--
		} else {
			enemy = e.NewEnemy(float32(i%11*60), float32(counter*50), 10, "red")
			view = e.NewEnemyView("assets/red.png", enemy)
			red--
		}
		controller := e.NewEnemyController(enemy, view)
		enemies[i] = controller
		c.Add(view.Sprite)
	}

	return &Board{
		Enemies: enemies,
	}
}
