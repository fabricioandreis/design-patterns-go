package principles

// Liskov Substitution Principle

type Sized interface {
	Width() int
	SetWidth(width int)
	Height() int
	SetHeight(height int)
}

type Rectangle struct {
	W, H int
}

func (r *Rectangle) Width() int {
	return r.W
}

func (r *Rectangle) SetWidth(width int) {
	r.W = width
}

func (r *Rectangle) Height() int {
	return r.H
}

func (r *Rectangle) SetHeight(height int) {
	r.H = height
}

type WrongSquare struct {
	Rectangle
}

func NewSquare(size int) *WrongSquare {
	sq := WrongSquare{}
	sq.W = size
	sq.H = size
	return &sq
}

func (s *WrongSquare) SetWidth(width int) {
	// Setting both width and height violates LSP because WrongSquare
	// has a different behavior from its composed (super) class
	s.W = width
	s.H = width
}

func (s *WrongSquare) SetHeight(height int) {
	s.H = height
	s.W = height
}

type RightSquare struct {
	size int
}

func (s *RightSquare) Rectangle() Rectangle {
	return Rectangle{s.size, s.size}
}
