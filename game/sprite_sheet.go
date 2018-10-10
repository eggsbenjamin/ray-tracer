package game

import (
	"github.com/veandco/go-sdl2/sdl"
)

type SpriteSheet interface {
	Image() *sdl.Texture
	GetFrames(int, *sdl.Rect) []*sdl.Rect
}

type spritesheet struct {
	image               *sdl.Texture
	imgWidth, imgHeight int32
}

func NewSpriteSheet(image *sdl.Texture) SpriteSheet {
	_, _, imgWidth, imgHeight, err := image.Query()
	if err != nil {
		panic("Handle this better! " + err.Error())
	}

	return &spritesheet{
		image:     image,
		imgWidth:  imgWidth,
		imgHeight: imgHeight,
	}
}

func (s *spritesheet) Image() *sdl.Texture {
	return s.image
}

func (s *spritesheet) GetFrames(n int, frame *sdl.Rect) []*sdl.Rect {
	frames := []*sdl.Rect{}
	for i := 0; i < n; i++ {
		frames = append(frames, &sdl.Rect{frame.X + int32(i)*frame.W, frame.Y, frame.W, frame.H})
	}

	return frames
}
