package main

import "fmt"

type Image interface {
	Draw()
}

type Bitmap struct {
	filename string
}

func NewBitMap(filename string) *Bitmap {
	fmt.Println("Loading the Image from", filename)
	return &Bitmap{filename: filename}
}

func (b *Bitmap) Draw() {
	fmt.Println("Drawing the Image", b.filename)
}

func DrawImage(image Image) {
	fmt.Println("About to draw the image")
	image.Draw()
	fmt.Println("Done Drawing the Image")
}

type LazyBitMap struct {
	filename string
	bitmap   *Bitmap
}

func (l *LazyBitMap) Draw() {
	if l.bitmap == nil {
		l.bitmap = NewBitMap(l.filename)
	}
	l.bitmap.Draw()
}

func NewLaxyBitMap(filename string) *LazyBitMap {
	return &LazyBitMap{filename: filename}
}

func main() {

	_ = NewBitMap("demo.png")
	//bmb.Draw()

	//DrawImage(bmb)
	_ = NewLaxyBitMap("demo.png")
}
