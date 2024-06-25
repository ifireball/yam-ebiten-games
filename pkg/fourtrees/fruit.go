package fourtrees

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/ifireball/yam-ebiten-games/resources"
)

const (
	kinds = 4
	amount = 20

	fruitWidth = 25 * 2 * 3 / 4
	fruitHeight = 25 * 2 * 3 / 4
)

var (
	minX, maxX = 74, 887
	treeEnds = [kinds - 1]int{247, 428, 696}
	minY, maxY = 38, 198

	imageNames = [kinds]string{"orange", "lemon", "apple", "pear"}
)

type Fruit struct {
	initialized bool
	locations [amount]struct{
		position ebiten.GeoM
		kind int
	}
	images [kinds]*ebiten.Image
}

func (f *Fruit) Update(screen *ebiten.Image) error {
	if !f.initialized {
		err := f.loadImages()
		if err != nil {
			return err
		}
		f.randomizeLocations()
		f.initialized = true
	}
	return nil
}

func (f *Fruit) loadImages() error {
	for i := 0; i < kinds; i++ {
		image, err := resources.ImageFromSVG(imageNames[i], fruitWidth, fruitHeight)
		if err != nil {
			return err
		}
		f.images[i], _ = ebiten.NewImageFromImage(image, ebiten.FilterDefault)
	}
	return nil
}

func (f *Fruit) randomizeLocations() {
	for i := 0; i < amount; i++ {
		x := rand.Intn(maxX-minX)+minX
		y := rand.Intn(maxY-minY)+minY
		top := float64(x) * 2 * 3 / 4 - fruitWidth / 2
		left := float64(y) * 2 * 3 / 4 - fruitHeight / 2
		f.locations[i].position.Translate(float64(top), float64(left))
		ki := 0
		for ; ki < kinds - 1 && x > treeEnds[ki]; ki++ {}
		f.locations[i].kind = ki

		println(x, y, top, left, ki)
	}
}

func (f *Fruit) Draw(screen *ebiten.Image) {
	dio := ebiten.DrawImageOptions{}
	for i := 0; i < amount; i++ {
		dio.GeoM = f.locations[i].position
		screen.DrawImage(f.images[f.locations[i].kind], &dio)
	}
}
