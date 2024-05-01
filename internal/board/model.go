package board

import (
	"fmt"
	"image/color"
	"math/rand"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	bt "github.com/pstano1/go-vaders/internal/bullet"
	e "github.com/pstano1/go-vaders/internal/enemy"
	p "github.com/pstano1/go-vaders/internal/player"
)

type IBoard interface {
	Size() (float32, float32)
	GetScore() int

	MoveEnemiesVertically(dy float32, direction int)
	MoveEnemiesHorizontally(dx float32, direction int)
	GetDirection(current int) int
	MoveBullets(dy float32)
	AppendBullet(bullet bt.IBulletController)
	MakeEnemiesShoot()
	CheckForHits()
	clearEnemiesBullets()

	edgeMostEnemyTouchesBoundary(direction int) bool
	EdgeMostEnemyReachesPlayer() bool
	IsPlayerAlive() bool
	CreateGameOverOverlay()
	HandleInput(e *fyne.KeyEvent, container *fyne.Container)
}

type Board struct {
	Height  float32
	Width   float32
	Score   int
	content *fyne.Container
	Player  p.IPlayerController
	Enemies []e.IEnemyController
	Bullets []bt.IBulletController
}

func NewBoard(c *fyne.Container, pc p.IPlayerController, width, height float32) IBoard {
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

	bullets := make([]bt.IBulletController, 0)
	return &Board{
		Height:  height,
		Width:   width,
		Score:   0,
		content: c,
		Player:  pc,
		Enemies: enemies,
		Bullets: bullets,
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
			if bullet.Bullet().Owner() == bt.PlayersBullet && enemy.CheckForCollision(x, y) {
				bullet.Destroy()
				b.Bullets = append(b.Bullets[:index], b.Bullets[index+1:]...)
				index--
				b.Score += enemy.Destroy()
				break
			}
			if b.Player.CheckForCollision(x, y) {
				bullet.Destroy()
				b.Bullets = append(b.Bullets[:index], b.Bullets[index+1:]...)
				index--
				b.Player.UpdateLifes(b.Player.GetLifes() - 1)
				b.clearEnemiesBullets()
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
		x, _ := enemy.Position()

		if direction == -1 && x <= 0 && enemy.IsAlive() {
			return true
		}
		if direction == 1 && x >= b.Width && enemy.IsAlive() {
			return true
		}
	}
	return false
}

func (b *Board) EdgeMostEnemyReachesPlayer() bool {
	for _, enemyController := range b.Enemies {
		enemy := enemyController.Enemy()
		_, y := enemy.Position()

		if y >= b.Height-100 && enemy.IsAlive() {
			return true
		}
	}
	return false
}

func (b *Board) MoveBullets(dy float32) {
	var wg sync.WaitGroup
	wg.Add(len(b.Bullets))

	for _, bullet := range b.Bullets {
		if bullet.Bullet().Owner() == bt.PlayersBullet {
			bullet.Move(dy, &wg)
			continue
		}
		bullet.Move(-dy, &wg)
	}
}

func (b *Board) AppendBullet(bullet bt.IBulletController) {
	b.Bullets = append(b.Bullets, bullet)
}

func (b *Board) CreateGameOverOverlay() {
	gameOver := canvas.NewText("GAME OVER", color.White)
	gameOver.Alignment = fyne.TextAlignCenter
	gameOver.TextSize = 48
	score := canvas.NewText(fmt.Sprintf("Your score is: %d", b.Score), color.White)
	score.Alignment = fyne.TextAlignCenter
	score.TextSize = 24

	overlayBackground := canvas.NewRectangle(color.Black)
	overlayBackground.Resize(fyne.NewSize(b.Width+60, b.Height))
	overlayBackground.FillColor = color.RGBA{R: 0, G: 0, B: 0, A: 127}

	UI := container.NewVBox(gameOver, score)
	UI.Move(fyne.NewPos(b.Width/2, 0))
	overlayContent := container.NewWithoutLayout()
	overlayContent.Resize(overlayBackground.Size())
	overlayContent.Move(fyne.NewPos(0, 0))
	overlayContent.Add(overlayBackground)
	overlayContent.Add(UI)
	b.content.Add(overlayContent)
}

func (b *Board) Size() (float32, float32) {
	return b.Width, b.Height
}

func (b *Board) GetScore() int {
	return b.Score
}

func (b *Board) HandleInput(e *fyne.KeyEvent, container *fyne.Container) {
	switch e.Name {
	case fyne.KeyRight:
		b.Player.Move(fyne.KeyRight, b.Width)
	case fyne.KeyLeft:
		b.Player.Move(fyne.KeyLeft, b.Width)
	case fyne.KeySpace:
		c, v := b.Player.Shoot()
		b.AppendBullet(c)
		container.Add(v.Sprite)
	}
}

func (b *Board) MakeEnemiesShoot() {
	var wg sync.WaitGroup
	wg.Add(len(b.Enemies))

	for _, enemy := range b.Enemies {
		if rand.Float64() < 0.05 && enemy.Enemy().IsAlive() {
			go func(e e.IEnemyController) {
				defer wg.Done()
				c, v := e.Shoot()
				b.AppendBullet(c)
				b.content.Add(v.Sprite)
			}(enemy)
		} else {
			wg.Done()
		}
	}

	wg.Wait()
}

func (b *Board) IsPlayerAlive() bool {
	return b.Player.GetLifes() != 0
}

func (b *Board) clearEnemiesBullets() {
	var bullets []bt.IBulletController
	for _, bullet := range b.Bullets {
		if bullet.Bullet().Owner() == bt.PlayersBullet {
			bullets = append(bullets, bullet)
			continue
		}
		bullet.Destroy()
	}
	b.Bullets = bullets
}
