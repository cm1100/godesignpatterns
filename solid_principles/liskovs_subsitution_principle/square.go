package main

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(width int) {
	s.width = width
	s.height = width
}

func (s *Square) GetHeight() int {
	return s.height
}

func (s *Square) GetWidth() int {
	return s.width
}
