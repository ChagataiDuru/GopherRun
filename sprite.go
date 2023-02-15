package main

import (
	"image"
	"io/fs"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct {
	Texture          *ebiten.Image // full texture
	Image            *ebiten.Image // displayed portion of image
	MirrorX, MirrorY bool
}

func NewSprite(img *ebiten.Image) *Sprite {
	return &Sprite{
		Texture: img,
		Image:   img,
	}
}

const useEmbeddedFiles = false

func OpenFile(path string) (fs.File, error) {
	if useEmbeddedFiles {
		return data.Open(path)
	}

	return os.Open(path)
}

func NewImageFromFile(path string) (*ebiten.Image, image.Image, error) {
	file, err := OpenFile(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	return ebitenutil.NewImageFromReader(file)
}

func (s *Sprite) SetTextureRect(r image.Rectangle) {
	s.Image = s.Texture.SubImage(r).(*ebiten.Image)
}

func (s *Sprite) GetTextureRect() image.Rectangle {
	return s.Image.Bounds()
}

func (s *Sprite) GetSize() Vec2f {
	b := s.Image.Bounds()
	return Vec2f{float64(b.Dx()), float64(b.Dy())}
}

func (s *Sprite) Draw(screen *ebiten.Image, parentM ebiten.GeoM) {
	var m ebiten.GeoM
	m.Translate(-float64(s.Image.Bounds().Dx())/2, -float64(s.Image.Bounds().Dy())/2)
	if s.MirrorX {
		m.Scale(-1, 1)
	}
	if s.MirrorY {
		m.Scale(1, -1)
	}
	m.Concat(parentM)
	op := &ebiten.DrawImageOptions{}
	op.GeoM = m
	screen.DrawImage(s.Image, op)
}
