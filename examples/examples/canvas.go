// +build !ci

package examples

import "runtime"
import "image/color"
import "path"

import "github.com/fyne-io/fyne/api/app"
import "github.com/fyne-io/fyne/api/ui"
import "github.com/fyne-io/fyne/api/ui/canvas"
import "github.com/fyne-io/fyne/api/ui/layout"

func rgbGradient(x, y, w, h int) color.RGBA {
	g := int(float32(x) / float32(w) * float32(255))
	b := int(float32(y) / float32(h) * float32(255))

	return color.RGBA{uint8(255 - b), uint8(g), uint8(b), 0xff}
}

func Canvas(app app.App) {
	w := app.NewWindow("Main")

	_, filename, _, _ := runtime.Caller(0)
	exampleDir := path.Dir(path.Dir(filename))
	content := ui.NewContainer(
		canvas.NewText("Resize me", color.RGBA{0, 0, 0x80, 0xff}),
		&canvas.Rectangle{FillColor: color.RGBA{0x80, 0, 0, 0xff},
			StrokeColor: color.RGBA{0xff, 0xff, 0xff, 0xff},
			StrokeWidth: 1},
		canvas.NewRaster(rgbGradient),
		canvas.NewImageFromFile(path.Join(exampleDir, "fyne.png")),
		&canvas.Line{StrokeColor: color.RGBA{0, 0x80, 0, 0xff}, StrokeWidth: 5},
		&canvas.Circle{StrokeColor: color.RGBA{0, 0, 0x80, 0xff},
			FillColor:   color.RGBA{0x30, 0x30, 0x30, 0x60},
			StrokeWidth: 2})
	content.Layout = layout.NewFixedGridLayout(ui.NewSize(93, 93))

	w.Canvas().SetContent(content)
	w.Show()
}