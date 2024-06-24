package resources

import (
	"embed"
	"fmt"
	"image"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/rasterizer"
)

//go:embed svg
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