package main

import (
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	all, err := ioutil.ReadFile("headerless-data,255,253.zx4")
	if err != nil {
		log.Fatalf("%v", err)
	}
	paris := all[0xeb00-0x6f00 : 0xeb00-0x6f00+0x80*0x14]
	width := []int{126, 126, 128, 127}
	show(all, paris, 0, width[0])
	for a := 1; a < 4; a++ {
		start := int(all[0x6022-0x5b00+a*2]) + 0x100*int(all[0x6023-0x5b00+a*2]) - 0x5b00
		m := decompress(all, start)
		show(all, m, a, width[a])
	}
}

type BitStream struct {
	data   []byte
	bitPos int
}

func (b *BitStream) Fetch(bits int) int {
	r := 0
	for i := 0; i < bits; i++ {
		o := b.Bit()
		r |= (o << uint(i))
	}
	return r
}

func (b *BitStream) DecodeFetch(table []byte, f int) byte {
	if f == 30 {
		return byte(b.Fetch(8))
	}
	return table[f]
}

func (b *BitStream) Bit() int {
	index := b.bitPos / 8
	bit := 7 - (b.bitPos % 8)
	b.bitPos++
	return int((b.data[index] >> uint(bit)) & 1)
}

func decompress(all []byte, start int) []byte {
	m := []byte{}
	b := &BitStream{
		data:   all[start+0x1E:],
		bitPos: 0,
	}
	for {
		c := b.Fetch(5)
		if c == 31 {
			p := b.Fetch(5)
			if p == 1 {
				return m
			}
			q := b.Fetch(5)
			x := b.DecodeFetch(all[start:start+30], q)
			for a := 0; a < p; a++ {
				m = append(m, x)
			}
		} else {
			x := b.DecodeFetch(all[start:start+30], c)
			m = append(m, x)
		}
	}
}

func show(all []byte, m []byte, level int, width int) {

	palette := []color.Color{
		color.RGBA{0x00, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xcc, 0xff},
		color.RGBA{0x00, 0xcc, 0x00, 0xff}, color.RGBA{0x00, 0xcc, 0xcc, 0xff},
		color.RGBA{0xcc, 0x00, 0x00, 0xff}, color.RGBA{0xcc, 0x00, 0xcc, 0xff},
		color.RGBA{0xcc, 0xcc, 0x00, 0xff}, color.RGBA{0xcc, 0xcc, 0xcc, 0xff},
		color.RGBA{0x00, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff},
		color.RGBA{0x00, 0xff, 0x00, 0xff}, color.RGBA{0x00, 0xff, 0xff, 0xff},
		color.RGBA{0xff, 0x00, 0x00, 0xff}, color.RGBA{0xff, 0x00, 0xff, 0xff},
		color.RGBA{0xff, 0xff, 0x00, 0xff}, color.RGBA{0xff, 0xff, 0xff, 0xff},
	}

	img := image.NewPaletted(image.Rect(0, 0, int(0x10*width), 0x10*len(m)/0x80), palette)

	for i := 0; i < int(all[0x7707-0x5b00]); i++ {
		all[0xea00-0x6f00+int(all[0xb29f-0x5b00+4+5*i])] = all[0xb29f-0x5b00+level+5*i]
	}

	for y := 0; y < len(m)/0x80; y++ {
		for x := 0; x < width; x++ {
			k := int(m[y*0x80+x])
			offset := 0x8000 - 0x5b00 + k*0x20
			chardata := all[offset : offset+0x20]
			for yi := 0; yi < 0x10; yi++ {
				attr := all[0xea00-0x6f00+k]
				for xil := 0; xil < 2; xil++ {
					ch := chardata[yi*2+xil]
					for xih := 0; xih < 8; xih++ {
						c := (ch >> uint(7-xih)) & 0x01
						l := ((attr>>6)&0x01)*0x33 + 0xcc
						q := color.NRGBA{
							R: l * ((attr >> (4 - c*3)) & 0x1),
							G: l * ((attr >> (5 - c*3)) & 0x1),
							B: l * ((attr >> (3 - c*3)) & 0x1),
							A: 0xff,
						}
						img.Set(int(x*0x10+xil*8+xih), int(y*0x10+yi), q)
					}
				}
			}
		}
	}
	create("level"+strconv.Itoa(level+1)+".png", img)
}

func create(filename string, im image.Image) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, im); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
