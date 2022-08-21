package structural

import (
	"fmt"
)

// Proxy is a structural design pattern that lets you provide a substitute or placeholder for another object. A proxy controls access to the original object, allowing you to perform something either before or after the request gets through to the original object.
// https://refactoring.guru/design-patterns/proxy

// Same interface, different behavior
// A type that functions as an interface to a particular resource.
// That resource may be remote, expensive to construct, or may require
// logging or some other added functionality

// Protection proxy: provides access control

type Driven interface {
	Drive() bool
}

type Car struct {
}

func (c *Car) Drive() bool {
	fmt.Println("Car is being driven")
	return true
}

type Driver struct {
	Age int
}

// This is the proxy object, which intercepts calls to the main object
type SafeCar struct {
	car    Car
	driver *Driver // Proxy could have any properties
}

// Drive method is also implemented by the proxy
func (c *SafeCar) Drive() bool {
	if c.driver.Age < 18 {
		fmt.Println("Driver cannot drive the car (age < 18)")
		return false
	}
	return c.car.Drive()
}

func NewSafeCar(driver *Driver) *SafeCar {
	return &SafeCar{Car{}, driver}
}

// Virtual proxy: pretends it is really there when in fact it isn't
// Only performs costly operation when it is actually needed
type Image interface {
	Draw() bool
}

type Bitmap struct {
	filename string
}

func NewBitmap(filename string) *Bitmap {
	fmt.Println("Loading image from", filename)
	return &Bitmap{filename}
}

func (b *Bitmap) Draw() bool {
	fmt.Println("Drawing image", b.filename)
	return true
}

func DrawImage(image Image) {
	fmt.Println("About to draw the image")
	image.Draw()
	fmt.Println("Done drawing the image")
}

type LazyBitmap struct {
	filename string
	bitmap   *Bitmap
}

func NewLazyBitmap(filename string) *LazyBitmap {
	return &LazyBitmap{filename: filename}
}

func (l *LazyBitmap) Draw() bool {
	if l.bitmap == nil {
		l.bitmap = NewBitmap(l.filename)
	}
	l.bitmap.Draw()
	return true
}
