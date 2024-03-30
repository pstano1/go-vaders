package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	board "github.com/pstano1/go-vaders/pkg/board"
	// enemy "github.com/pstano1/go-vaders/pkg/enemy"
	player "github.com/pstano1/go-vaders/pkg/player"
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
	w.Canvas().SetOnTypedKey(pc.HandleKey)

	currentDirection := 1
	ticker := time.Tick(1 * time.Second)
	go func() {
		for {
			<-ticker

			direction := b.GetDirection(currentDirection)
			if direction != currentDirection {
				b.MoveEnemiesVertically(50, direction)
				currentDirection = direction
			} else {
				b.MoveEnemiesHorizontally(50, direction)
			}
		}
	}()

	a.Run()
}
