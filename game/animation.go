package game

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Animation interface {
	Len() int
	Image() *sdl.Texture
	GetFrame(int) *sdl.Rect
}

type animation struct {
	spritesheet SpriteSheet
	frames      []*sdl.Rect
}

func NewAnimation(spritesheet SpriteSheet, frames []*sdl.Rect) Animation {
	return &animation{
		spritesheet: spritesheet,
		frames:      frames,
	}
}

func (a *animation) Len() int {
	return len(a.frames)
}

func (a *animation) Image() *sdl.Texture {
	return a.spritesheet.Image()
}

func (a *animation) GetFrame(i int) *sdl.Rect {
	return a.frames[i]
}
