package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

type Game struct {
	player           *Player
	meteorSpawnTimer *Timer
	meteors          []*Meteor
	score            int
}

func (g *Game) Update() error {
	g.player.Update()

	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()

		m := NewMeteor()
		g.meteors = append(g.meteors, m)
	}

	for _, m := range g.meteors {
		m.Update()
	}

	for i, m := range g.meteors {
		for j, b := range g.player.bullets {
			if m.Collider().Intersects(b.Collider()) {
				g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
				g.player.bullets = append(g.player.bullets[:j], g.player.bullets[i+j:]...)
				g.score += 1
			}

		}
	}

	for _, m := range g.meteors {
		if m.Collider().Intersects(g.player.Collider()) {
			g.Reset()
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
	for _, m := range g.meteors {
		m.Draw(screen)
	}

	msg := fmt.Sprintf("%06d", g.score)
	op := &text.DrawOptions{}
	op.GeoM.Translate(ScreenWidth/2-100, 50)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, msg, text.NewGoXFace(FontFace), op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Reset() {
	g.player = NewPlayer()
	g.meteors = nil
	g.score = 0
}

func main() {
	g := &Game{
		player:           NewPlayer(),
		meteorSpawnTimer: NewTimer(5 * time.Second),
	}

	err := ebiten.RunGame(g)

	if err != nil {
		panic(err)
	}
}
