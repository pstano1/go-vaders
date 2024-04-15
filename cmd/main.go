package main

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

	"github.com/pstano1/go-vaders/internal/board"
	"github.com/pstano1/go-vaders/internal/player"
)

func main() {
	a := app.New()
	w := a.NewWindow("go-vaders")
	w.Resize(fyne.NewSize(1000, 800))

	content := container.NewWithoutLayout()
	boardContainer := createBoard()
	p := player.NewPlayer(375, 550)
	v := player.NewPlayerView("assets/player.png", p)
	pc := player.NewPlayerController(p, v)

	b := board.NewBoard(boardContainer, 800, 600)

	boardContainer.Add(v.Sprite)

	score := b.GetScore()
	lifes := p.Lifes()
	scoreText := canvas.NewText(fmt.Sprint(score), color.White)
	scoreText.TextSize = 24
	lifesText := canvas.NewText(fmt.Sprintf("Lifes: %d", lifes), color.White)
	lifesText.TextSize = 24
	topBar := container.NewHBox(scoreText, lifesText)
	topBar.Move(fyne.NewPos(70, 50))
	content.Add(topBar)
	content.Add(boardContainer)

	w.SetContent(content)
	w.Show()
	w.Canvas().SetOnTypedKey(func(e *fyne.KeyEvent) {
		pc.HandleKey(e, b, content)
	})

	currentDirection := 1
	ticker := make(chan struct{})
	bulletTicker := make(chan struct{})
	go func() {
		for {
			<-time.After(2 * time.Second)

			direction := b.GetDirection(currentDirection)
			if direction != currentDirection {
				b.MoveEnemiesVertically(50, direction)
				currentDirection = direction
				if ok := b.EdgeMostEnemyReachesPlayer(); ok {
					close(ticker)
					close(bulletTicker)
					b.CreateGameOverOverlay()
					break
				}
			} else {
				b.MoveEnemiesHorizontally(50, direction)
			}
		}
	}()

	go func() {
		for {
			<-time.After(200 * time.Millisecond)

			b.MoveBullets(50)
			b.CheckForHits()

			score = b.GetScore()
			updateScore(scoreText, score)
		}
	}()

	a.Run()
}

func updateScore(sc *canvas.Text, newValue int) {
	sc.Text = fmt.Sprintf("score: %d", newValue)
	sc.Refresh()
}

func updateLifes(lf *canvas.Text, newValue int) {
	return
}

func createBoard() *fyne.Container {
	boardContainer := container.NewWithoutLayout()
	boardContainer.Move(fyne.NewPos(70, 100))
	border := canvas.NewRectangle(color.Transparent)
	border.Resize(fyne.NewSize(870, 610))
	border.Move(fyne.NewPos(-5, -5))
	border.FillColor = color.Transparent
	border.StrokeColor = color.White
	border.StrokeWidth = 5
	boardContainer.Add(border)

	return boardContainer
}
