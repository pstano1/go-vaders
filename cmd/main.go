package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"github.com/pstano1/go-vaders/internal/board"
	"github.com/pstano1/go-vaders/internal/player"
)

func main() {
	a := app.New()
	w := a.NewWindow("go-vaders")
	w.Resize(fyne.NewSize(800, 600))

	content := container.NewWithoutLayout()
	p := player.NewPlayer(375, 550)
	v := player.NewPlayerView("assets/player.png", p)
	pc := player.NewPlayerController(p, v)

	b := board.NewBoard(content, 800, 600)

	content.Add(v.Sprite)

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
		}
	}()

	a.Run()
}
