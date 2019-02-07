package mnist

import (
	"fmt"
	"image"
	_ "image/jpeg" //register jpg codec
	_ "image/png"  //register the png codec
	"os"
	"strings"
)

//FromImage will read a PNG file and extract a DigitImage with prior conversion from rgba to gray scale
func FromImage(filepath string) (*DigitImage, error) {
	if !strings.HasSuffix(filepath, ".png") && !strings.HasSuffix(filepath, ".jpg") && !strings.HasSuffix(filepath, ".jpeg") {
		return nil, fmt.Errorf("file '%s' does not have .png .jpg or .jpeg extension", filepath)
	}
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	//png.Decode
	defer f.Close()
	im, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	bounds := im.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	res := &DigitImage{}
	data := [][]uint8{}
	// Convert color to grayscale
	//gray := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{w, h}})
	gray := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{1, 1}})
	for y := 0; y < h; y++ {
		row := make([]uint8, w)
		for x := 0; x < w; x++ {
			//gray.Set(x, y, im.At(x, y))
			gray.Set(0, 0, im.At(x, y))
			row[x] = gray.GrayAt(0, 0).Y
		}
		data = append(data, row)
	}
	res.Image = data
	return res, nil
}
