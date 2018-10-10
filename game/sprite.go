package game

import "github.com/veandco/go-sdl2/sdl"

type Sprite interface {
	LoadContent()
	Bounds() (x float64, y float64, w float64, h float64)
	SetX(float64)
	SetY(float64)
	SetAnimation(string)
	GetCurrentFrame() (*sdl.Texture, *sdl.Rect)
	Update()
	Draw(*sdl.Renderer)
}

type sprite struct {
	x, y, width, height                  float64
	animations                           map[string]Animation
	currentAnimation                     Animation
	frameIndex, ticksPerFrame, tickCount int
}

func NewSprite(animations map[string]Animation, startingAnimation string, x, y, width, height float64, frameIndex, ticksPerFrame int) Sprite {
	return &sprite{
		x:                x,
		y:                y,
		width:            width,
		height:           height,
		animations:       animations,
		currentAnimation: animations[startingAnimation],
		frameIndex:       frameIndex,
		ticksPerFrame:    ticksPerFrame,
	}
}

func (s *sprite) LoadContent() {}

func (s *sprite) Bounds() (float64, float64, float64, float64) {
	return s.x, s.y, s.width, s.height
}

func (s *sprite) SetX(x float64) {
	s.x = x
}

func (s *sprite) SetY(y float64) {
	s.y = y
}

func (s *sprite) SetAnimation(name string) {
	animation, ok := s.animations[name]
	if !ok {
		panic("Handle this better! " + "animation '" + name + "' not found")
	}

	s.currentAnimation = animation
	s.frameIndex = 0
}

func (s *sprite) Update() {
	s.tickCount++

	if s.tickCount > s.ticksPerFrame {
		s.tickCount = 0
		s.frameIndex++

		if s.frameIndex >= s.currentAnimation.Len() {
			s.frameIndex = 0
		}
	}
}

func (s *sprite) GetCurrentFrame() (*sdl.Texture, *sdl.Rect) {
	return s.currentAnimation.Image(), s.currentAnimation.GetFrame(s.frameIndex)
}

func (s *sprite) Draw(r *sdl.Renderer) {
	if err := r.CopyEx(
		s.currentAnimation.Image(),
		s.currentAnimation.GetFrame(s.frameIndex),
		&sdl.Rect{100, 100, int32(s.width), int32(s.height)},
		0,
		nil,
		0,
	); err != nil {
		panic("Handle this better! " + err.Error())
	}
}
