package rating

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

const interval = 5

type Rating struct {
	originalImg image.Image
	marks       []*ebiten.Image
	ops         []*ebiten.DrawImageOptions
	maskImg     *ebiten.Image
	maxVal      int
	fillCount   int
}

func NewRating(img image.Image, x, y int, max int) *Rating {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	maskImg := ebiten.NewImage(width, height)

	marks := []*ebiten.Image{}
	for i := 0; i < max; i++ {
		marks = append(marks, ebiten.NewImageFromImage(img))
	}

	ops := []*ebiten.DrawImageOptions{}
	for i := 0; i < max; i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(x+i*(width+interval)), float64(y))
		ops = append(ops, op)
	}

	return &Rating{originalImg: img, marks: marks, ops: ops, maskImg: maskImg, fillCount: max, maxVal: max}
}

func (r *Rating) SetValue(val float64) {
	r.marks[r.fillCount-1] = ebiten.NewImageFromImage(r.originalImg)

	r.fillCount = int(val)
	if r.fillCount >= int(r.maxVal) {
		return
	}

	op := &ebiten.DrawImageOptions{}
	shiftRate := val - float64(int(val))
	width := float64(r.originalImg.Bounds().Dx())
	op.GeoM.Translate(width*shiftRate, 0.0)
	op.Blend = ebiten.BlendClear
	r.marks[r.fillCount-1].DrawImage(r.maskImg, op)
}

func (r *Rating) Draw(screen *ebiten.Image) {
	for i := 0; i < len(r.marks); i++ {
		if i <= r.fillCount-1 {
			screen.DrawImage(r.marks[i], r.ops[i])
		}
	}
}
