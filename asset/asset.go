package asset

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	fontSize   = 1000
	fontsPath  = "./assets/fonts/"
	imagesPath = "./assets/images/"
)

func init() {
	if err := ttf.Init(); err != nil {
		panic("Handle this better! " + err.Error())
	}

	if retVal := img.Init(img.INIT_PNG); retVal == 0 {
		panic("Handle this better! " + img.GetError().Error())
	}
}

type Manager interface {
	LoadContent()
	UnloadContent()
	GetImage(string) *sdl.Texture
	GetFont(string) *ttf.Font
}

type manager struct {
	renderer              *sdl.Renderer
	fontsPath, imagesPath string
	fontsSize             int
	fonts                 map[string]*ttf.Font
	images                map[string]*sdl.Texture
}

func NewManager(renderer *sdl.Renderer, fontsPath, imagesPath string, fontsSize int) *manager {
	return &manager{
		renderer:   renderer,
		fontsPath:  fontsPath,
		imagesPath: imagesPath,
		fontsSize:  fontsSize,
		fonts:      make(map[string]*ttf.Font),
		images:     make(map[string]*sdl.Texture),
	}
}

func (m *manager) LoadContent() {
	fontFiles, err := ioutil.ReadDir(fontsPath)
	if err != nil {
		panic("Handle this better! " + err.Error())
	}

	fmt.Println("Loading font files...")
	for _, f := range fontFiles {
		fnt, err := ttf.OpenFont(fontsPath+f.Name(), fontSize)
		if err != nil {
			panic("Handle this better! " + err.Error())
		}

		m.fonts[strings.Split(f.Name(), ".")[0]] = fnt
		fmt.Printf("\t%s%s\n", fontsPath, f.Name())
	}

	imageFiles, err := ioutil.ReadDir(imagesPath)
	if err != nil {
		panic("Handle this better! " + err.Error())
	}

	fmt.Println("Loading image files...")
	for _, f := range imageFiles {
		image, err := img.LoadTexture(m.renderer, imagesPath+f.Name())
		if err != nil {
			panic("Handle this better! " + err.Error())
		}

		m.images[strings.Split(f.Name(), ".")[0]] = image
		fmt.Printf("\t%s%s\n", imagesPath, f.Name())
	}
}

func (m *manager) UnloadContent() {}

func (m *manager) GetImage(name string) *sdl.Texture {
	image, ok := m.images[name]
	if !ok {
		panic("image '" + name + "' doesn't exist")
	}

	return image
}

func (m *manager) GetFont(name string) *ttf.Font {
	fnt, ok := m.fonts[name]
	if !ok {
		panic("font '" + name + "' doesn't exist")
	}

	return fnt
}
