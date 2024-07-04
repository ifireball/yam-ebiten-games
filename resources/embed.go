package resources

import (
	"embed"
	"encoding/json"
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/rasterizer"
)

//go:embed svg
//go:embed data
var Resources embed.FS

func ImageFromSVG(svgName string, w, h int) (image.Image, error) {
	file, err := Resources.Open(fmt.Sprintf("svg/%s.svg", svgName))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	cvs, err := canvas.ParseSVG(file)
	if err != nil {
		return nil, err
	}

	cw, ch := cvs.Size()
	scaleX, scaleY := float64(w)/cw, float64(h)/ch
	cvs.Transform(canvas.Identity.Scale(scaleX, scaleY))
	cvs.Clip(canvas.Rect{X: 0, Y: 0, W: float64(w), H: float64(h)})
	image := rasterizer.Draw(cvs, 1.0, canvas.DefaultColorSpace)
	return image, nil
}

func EbitenImageFromSVG(svgName string, w, h int) (*ebiten.Image, error) {
	image, err := ImageFromSVG(svgName, w, h)
	if err != nil {
		return nil, err
	}
	return ebiten.NewImageFromImage(image, ebiten.FilterDefault)
}

func DecodeData(jsonName string, dest interface{}) error {
	file, err := Resources.Open(fmt.Sprintf("data/%s.json", jsonName))
	if err != nil {
		return err
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	return dec.Decode(dest)
}
