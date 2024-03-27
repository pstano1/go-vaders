package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	player "github.com/pstano1/go-vaders/pkg/player"
)

func main() {
	a := app.New()
	w := a.NewWindow("go-vaders")
	w.Resize(fyne.NewSize(800, 600))
	content := container.NewWithoutLayout()
	p := player.NewPlayer(375, 550)
	v := player.NewPlayerView("assets/player.png", 375, 550)
	pc := player.NewPlayerController(p, v)

	content.Add(v.Sprite)

	w.SetContent(content)
	w.Show()
	w.Canvas().SetOnTypedKey(pc.HandleKey)

	a.Run()
}
