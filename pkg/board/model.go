package board

import (
	"sync"

	"fyne.io/fyne/v2"
	b "github.com/pstano1/go-vaders/pkg/bullet"
	e "github.com/pstano1/go-vaders/pkg/enemy"
)

type IBoard interface {
	MoveEnemiesVertically(dy float32, direction int)
	MoveEnemiesHorizontally(dx float32, direction int)
	GetDirection(current int) int
	MoveBullets(dy float32)
	AppendBullet(bullet b.IBulletController)
	CheckForHits()

	edgeMostEnemyTouchesBoundary(direction int) bool
}

type Board struct {
	Height              float32
	Width               float32
	EdgeMostColumnLeft  int
	EdgeMostColumnRight int
	Enemies             []e.IEnemyController
	Bullets             []b.IBulletController
}

func NewBoard(c *fyne.Container, width, height float32) IBoard {
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

		if red != 0 {
			enemy = e.NewEnemy(float32(i%11*60), float32(counter*50), 10, i, "red")
			view = e.NewEnemyView("assets/red.png", enemy)
			red--
		} else if yellow != 0 {
			enemy = e.NewEnemy(float32(i%11*60), float32(counter*50), 20, i, "yellow")
			view = e.NewEnemyView("assets/yellow.png", enemy)
			yellow--
		} else {
			enemy = e.NewEnemy(float32(i%11*60), float32(counter*50), 40, i, "green")
			view = e.NewEnemyView("assets/green.png", enemy)
			green--
		}
		controller := e.NewEnemyController(enemy, view)
		enemies[i] = controller
		c.Add(view.Sprite)
	}

	bullets := make([]b.IBulletController, 0)
	return &Board{
		Height:              height,
		Width:               width,
		Enemies:             enemies,
		Bullets:             bullets,
		EdgeMostColumnLeft:  0,
		EdgeMostColumnRight: 10,
	}
}

func (b *Board) MoveEnemiesVertically(dy float32, direction int) {
	var wg sync.WaitGroup
	wg.Add(len(b.Enemies))

	for _, enemy := range b.Enemies {
		enemy.Move(0, dy, direction, &wg)
	}
}

func (b *Board) MoveEnemiesHorizontally(dx float32, direction int) {
	var wg sync.WaitGroup
	wg.Add(len(b.Enemies))

	for _, enemy := range b.Enemies {
		enemy.Move(dx, 0, direction, &wg)
	}
}

func (b *Board) CheckForHits() {
	for _, enemy := range b.Enemies {
		for index, bullet := range b.Bullets {
			x, y := bullet.Bullet().Position()
			if enemy.CheckForCollision(x, y) {
				bullet.Destroy()
				b.Bullets = append(b.Bullets[:index], b.Bullets[index+1:]...)
				index--
				enemy.Destroy()
				break
			}
		}
	}
}

func (b *Board) GetDirection(current int) int {
	direction := current

	if current == -1 && b.edgeMostEnemyTouchesBoundary(current) {
		direction = 1
	}
	if current == 1 && b.edgeMostEnemyTouchesBoundary(current) {
		direction = -1
	}
	return direction
}

func (b *Board) edgeMostEnemyTouchesBoundary(direction int) bool {
	for _, enemyController := range b.Enemies {
		enemy := enemyController.Enemy()
		column := enemy.GetColumn()
		x, _ := enemy.Position()

		if direction == -1 && column == b.EdgeMostColumnLeft && x == 0 {
			return true
		}
		if direction == 1 && column == b.EdgeMostColumnRight && x == b.Width {
			return true
		}
	}
	return false
}

func (b *Board) MoveBullets(dy float32) {
	var wg sync.WaitGroup
	wg.Add(len(b.Bullets))

	for _, bullet := range b.Bullets {
		bullet.Move(dy, &wg)
	}
}

func (b *Board) AppendBullet(bullet b.IBulletController) {
	b.Bullets = append(b.Bullets, bullet)
}
