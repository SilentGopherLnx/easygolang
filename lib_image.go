package easygolang

import (
	"bytes"

	"image"
	"image/color"
	"image/draw"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

const MAX_A = 65535 / 255 //?

func ImageDecode(data *[]byte) image.Image {
	if data == nil || len(*data) == 0 {
		return nil
	}
	img, _, err := image.Decode(bytes.NewReader(*data))
	if err != nil {
		Prln("image.Decode ERROR")
		return nil
	}
	return img
}

func ImageDecodeRGBA(data *[]byte, colorTransperent color.RGBA) *image.RGBA {
	if data == nil || len(*data) == 0 {
		return nil
	}
	img, _, err := image.Decode(bytes.NewReader(*data))
	if err != nil {
		Prln("image.Decode ERROR")
		return nil
	}
	switch rgba := img.(type) {
	case *image.RGBA:
		return rgba
	default: //case *image.NRGBA:
		img2 := image.NewRGBA(img.Bounds())
		wh := img2.Bounds().Max
		for y := 0; y < wh.Y; y++ {
			for x := 0; x < wh.X; x++ {
				img2.SetRGBA(x, y, colorTransperent)
			}
		}
		draw.Draw(img2, img.Bounds(), img, image.Point{}, draw.Over)
		return img2
	}
}

func ImageClone(img image.Image) *image.RGBA {
	if InterfaceNil(img) {
		return nil
	}
	w := img.Bounds().Max.X
	h := img.Bounds().Max.Y
	img2 := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			col := img.At(x, y)
			img2.Set(x, y, col)
		}
	}
	return img2
}

func ImageResizeNearest(img image.Image, w int, h int) *image.RGBA {
	if InterfaceNil(img) {
		return nil
	}
	w1 := img.Bounds().Max.X
	h1 := img.Bounds().Max.Y
	xs := float64(w1) / float64(w)
	ys := float64(h1) / float64(h)
	img2 := image.NewRGBA(image.Rect(0, 0, w, h))
	m := img2.ColorModel()
	for y2 := 0; y2 < h; y2++ {
		y1 := RoundF((float64(y2)+0.5)*ys - 0.5)
		y1 = MAXI(0, MINI(h1, y1))
		for x2 := 0; x2 < w; x2++ {
			x1 := RoundF((float64(x2)+0.5)*xs - 0.5)
			x1 = MAXI(0, MINI(w1, x1))
			col := img.At(x1, y1)
			r0, g0, b0, a0 := m.Convert(col).RGBA()
			img2.Set(x2, y2, color.RGBA{uint8(r0), uint8(g0), uint8(b0), uint8(a0 / MAX_A)})
			// if a0/MAX_A > 255 {
			// 	Prln("!!")
			// }
		}
	}
	return img2
}

func ImageResizeHalfNice(img *image.RGBA) *image.RGBA {
	if img == nil {
		return nil
	}
	w := img.Bounds().Max.X
	h := img.Bounds().Max.Y
	w2 := w / 2
	h2 := h / 2
	//return ImageResizeNearest(img, w2, h2)
	img2 := image.NewRGBA(image.Rect(0, 0, w2, h2))
	mix4 := func(c1, c2, c3, c4 uint32) uint8 {
		v1 := int(c1 / MAX_A)
		v2 := int(c2 / MAX_A)
		v3 := int(c3 / MAX_A)
		v4 := int(c4 / MAX_A)
		r := (v1 + v2 + v3 + v4) / 4
		r = MAXI(0, MINI(255, r))
		return uint8(r)
	}
	m := img2.ColorModel()
	for y2 := 0; y2 < h2; y2++ {
		for x2 := 0; x2 < w2; x2++ {
			x22 := x2 * 2
			y22 := y2 * 2
			r1, g1, b1, a1 := m.Convert(img.At(x22, y22)).RGBA()
			r2, g2, b2, a2 := m.Convert(img.At(x22+1, y22)).RGBA()
			r3, g3, b3, a3 := m.Convert(img.At(x22, y22+1)).RGBA()
			r4, g4, b4, a4 := m.Convert(img.At(x22+1, y22+1)).RGBA()
			a1i := int(a1 / MAX_A)
			a2i := int(a2 / MAX_A)
			a3i := int(a3 / MAX_A)
			a4i := int(a4 / MAX_A)
			//Prln(I2S(int(r1)) + "," + I2S(int(g1)) + "," + I2S(int(b1)) + "," + I2S(int(a1)))
			alpha := int(a1i + a2i + a3i + a4i)
			if alpha > 0 {
				r_new := mix4(r1, r2, r3, r4)
				g_new := mix4(g1, g2, g3, g4)
				b_new := mix4(b1, b2, b3, b4)
				img2.SetRGBA(x2, y2, color.RGBA{uint8(r_new), uint8(g_new), uint8(b_new), uint8(alpha / 4)})
			}
		}
	}
	return img2
}

func BlendColors(argb_back [4]float64, argb_over [4]float64, maxv float64) [4]int {
	r := [4]int{0, 0, 0, 0}
	b_a := argb_back[0]
	o_a := argb_over[0]
	r_a := maxv - maxv*(maxv-o_a)*(maxv-b_a)
	if r_a < 0.001 {
		return r
	}
	r[0] = int(r_a)
	r[1] = int((argb_over[1]*o_a + argb_back[1]*b_a*(maxv-o_a)) / r_a / maxv)
	r[2] = int((argb_over[2]*o_a + argb_back[2]*b_a*(maxv-o_a)) / r_a / maxv)
	r[3] = int((argb_over[3]*o_a + argb_back[3]*b_a*(maxv-o_a)) / r_a / maxv)
	maxv_int := int(maxv)
	r[0] = MAXI(0, MINI(maxv_int, r[0]))
	r[1] = MAXI(0, MINI(maxv_int, r[1]))
	r[2] = MAXI(0, MINI(maxv_int, r[2]))
	r[3] = MAXI(0, MINI(maxv_int, r[3]))
	return r
}

func ImageAddOver(img1 *image.RGBA, img2 image.Image, x int, y int) {
	if img1 != nil && !InterfaceNil(img2) {
		// wh2 := img2.Bounds().Max
		// if wh2.X > 0 && wh2.Y > 0 {
		draw.Draw(img1, img1.Bounds(), img2, image.Point{X: -x, Y: -y}, draw.Over)
		//}
	}
	//h := img2.Bounds().Max.Y
	//image.Point{X: x, Y: h - y}
	/*if img1 != nil && img2 != nil {
		w1 := img1.Bounds().Max.X
		h1 := img1.Bounds().Max.Y
		w2 := img2.Bounds().Max.X
		h2 := img2.Bounds().Max.Y
		for y2 := 0; y2 < h2; y2++ {
			for x2 := 0; x2 < w2; x2++ {
				x1 := x + x2
				y1 := y + y2
				if x1 > 0 && y1 > 0 && x1 < w1 && y1 < h1 {
					r, g, b, a := img1.ColorModel().Convert(img2.At(x2, y2)).RGBA()
					a0 := int(a) / MAX_A
					if a0 > 0 {
						if a0 == 255 {
							col2 := color.RGBA{uint8(r), uint8(g), uint8(b), 255}
							img1.Set(x1, y1, col2)
						} else {
							col1 := img1.RGBAAt(x1, y1)
							new_col := BlendColors(
								[4]float64{float64(col1.A), float64(col1.R), float64(col1.G), float64(col1.B)},
								[4]float64{float64(a0), float64(r), float64(g), float64(b)},
								255.0)
							col2 := color.RGBA{uint8(new_col[1]), uint8(new_col[2]), uint8(new_col[3]), uint8(new_col[0])}
							img1.Set(x1, y1, col2)
						}
					}
				}
			}
		}
	}*/
}
