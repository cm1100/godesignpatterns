package main

type Buffer struct {
	width, height int
	buffer        []rune
}

func NewBuffer(width, height int) *Buffer {
	return &Buffer{
		width, height, make([]rune, height*width),
	}
}

func (b *Buffer) At(index int) rune {
	return b.buffer[index]
}

type ViewPort struct {
	buffer *Buffer
	offset int
}

func (v *ViewPort) GetCharacterAt(index int) rune {
	return v.buffer.At(v.offset + index)
}

func NewViewPort(buffer *Buffer) *ViewPort {
	return &ViewPort{buffer: buffer}
}
