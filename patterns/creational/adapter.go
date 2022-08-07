package creational

import (
	"crypto/md5"
	"encoding/json"
	"strings"
)

// Adapter is a structural design pattern that allows objects with incompatible interfaces to collaborate.
// Gives us the interface that we require from the interface that we have.
// An adapter is a special object that converts the interface of one object so that another object can understand it.

// https://refactoring.guru/design-patterns/adapter

// API for rendering graphical objects where all the images are defined as lines (vector based)
//
// The structs and function bellow are the interface given
type Line struct {
	X1, Y1, X2, Y2 int
}

type VectorImage struct {
	Lines []Line
}

func NewRectangle(width, height int) *VectorImage {
	width -= 1 // zero based coordinates
	height -= 1
	return &VectorImage{
		[]Line{
			{0, 0, width, 0},
			{0, 0, 0, height},
			{width, 0, width, height},
			{0, height, width, height},
		},
	}
}

// The structs and function bellow are the interface we have
type Point struct {
	X, Y int
}

type RasterImage interface {
	GetPoints() []Point
}

func DrawPoints(owner RasterImage) string {
	maxX, maxY := 0, 0
	points := owner.GetPoints()
	for _, pixel := range points {
		if pixel.X > maxX {
			maxX = pixel.X
		}
		if pixel.Y > maxY {
			maxY = pixel.Y
		}
	}
	maxX += 1
	maxY += 1

	// preallocate
	data := make([][]rune, maxY)
	for i := 0; i < maxY; i++ {
		data[i] = make([]rune, maxX)
		for j := range data[i] {
			data[i][j] = ' '
		}
	}

	for _, point := range points {
		data[point.Y][point.X] = '*'
	}

	b := strings.Builder{}
	for _, line := range data {
		b.WriteString(string(line))
		b.WriteRune('\n')
	}
	return b.String()
}

type DefaultImage struct {
	Points []Point
}

func (d *DefaultImage) GetPoints() []Point {
	return d.Points
}

func (d *DefaultImage) addLine(line Line) {
	if d.processFromCache(line) {
		return
	}

	left, right := minMax(line.X1, line.X2)
	top, bottom := minMax(line.Y1, line.Y2)
	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			d.Points = append(d.Points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			d.Points = append(d.Points, Point{x, top})
		}
	}
	pointCache[d.hash(line)] = d.Points
}

// processFromCache returns true if the line was found in the cache and the points added to the DefaultImage.
// False otherwise.
func (d *DefaultImage) processFromCache(line Line) bool {
	h := d.hash(line)
	pts, ok := pointCache[h]
	if !ok {
		return false
	}
	for _, pt := range pts {
		d.Points = append(d.Points, pt)
	}
	return true
}

func (d *DefaultImage) hash(obj interface{}) [16]byte {
	bytes, _ := json.Marshal(obj)
	return md5.Sum(bytes)
}

func minMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

var pointCache = map[[16]byte][]Point{}

// VectorToRaster is an adapter that returns an interface that we want from an interface that we have
func VectorToRaster(vector *VectorImage) RasterImage {
	img := DefaultImage{}
	if vector == nil {
		return &img
	}

	for _, l := range vector.Lines {
		img.addLine(l)
	}
	return &img
}
