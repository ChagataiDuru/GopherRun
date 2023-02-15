package main

import "github.com/hajimehoshi/ebiten/v2"

type Entity struct {
	Transform *TransformComponent
	Movement  *MovementComponent
	Graphics  *GraphicsComponent
}

func NewEntity() Entity {
	return Entity{
		Transform: &TransformComponent{},
		Movement:  &MovementComponent{},
		Graphics:  &GraphicsComponent{},
	}
}

func (e *Entity) Draw(screen *ebiten.Image) {
	var m ebiten.GeoM
	m.Translate(
		e.Transform.Position.X,
		e.Transform.Position.Y,
	)
	e.Graphics.Sprite.Draw(screen, m)
}
