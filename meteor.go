package main

import (
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	rotationSpeedMin = -0.02
	rotationSpeedMax = 0.02
)

type Meteor struct {
	position      Vector
	sprite        *ebiten.Image
	movement      Vector
	rotation      float64
	rotationSpeed float64
}

func NewMeteor() *Meteor {
	var target = Vector{
		X: ScreenWidth / 2,
		Y: ScreenHeight / 2,
	}

	var r = ScreenWidth / 2.0

	var angle = rand.Float64() * 2 * math.Pi

	var pos = Vector{
		X: target.X + math.Cos(angle)*r,
		Y: target.Y + math.Sin(angle)*r,
	}

	var velocity = 0.25 + rand.Float64()*1.5

	var direction = Vector{
		X: target.X - pos.X,
		Y: target.Y - pos.Y,
	}

	var normalizedDirection = direction.Normalize()

	movement := Vector{
		X: normalizedDirection.X * velocity,
		Y: normalizedDirection.Y * velocity,
	}

	sprite := MeteorSprites[rand.Intn(len(MeteorSprites))]

	return &Meteor{
		position:      pos,
		sprite:        sprite,
		movement:      movement,
		rotationSpeed: rotationSpeedMin + rand.Float64()*(rotationSpeedMax-rotationSpeedMin),
	}
}

func (m *Meteor) Update() {
	m.position.X += m.movement.X
	m.position.Y += m.movement.Y
	m.rotation += m.rotationSpeed
}

func (m *Meteor) Draw(screen *ebiten.Image) {
	bounds := m.sprite.Bounds()

	halfH := float64(bounds.Dy()) / 2
	halfW := float64(bounds.Dx()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(m.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(m.position.X, m.position.Y)

	screen.DrawImage(m.sprite, op)

}

func (m *Meteor) Collider() Rect {

	bounds := m.sprite.Bounds()

	return NewRect(
		m.position.X,
		m.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
