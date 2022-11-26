package behavioral

import (
	"container/list"
	"fmt"
)

// Observer is a behavioral design pattern that lets you define a subscription mechanism to notify multiple objects about any events that happen to the object they’re observing.

// Motivation: we need to be informed when certain things happen:
// - object's fields changes
// - object does something
// - some external event occurs

// https://refactoring.guru/design-patterns/observer

type Publisher struct {
	subs *list.List
}

func (p *Publisher) Subscribe(o Observer) {
	p.subs.PushBack(o)
}

func (p *Publisher) Unsubscribe(o Observer) {
	for z := p.subs.Front(); z != nil; z = z.Next() {
		if z.Value.(Observer) == o {
			p.subs.Remove(z)
		}
	}
}

func (p *Publisher) Publish(data string) {
	for z := p.subs.Front(); z != nil; z = z.Next() {
		z.Value.(Observer).Notify(data)
	}
}

type Observer interface {
	Notify(data string)
}

type Patient struct {
	Publisher
	Name string
}

func NewPatient(name string) *Patient {
	return &Patient{
		Publisher: Publisher{&list.List{}},
		Name:      name,
	}
}

func (p *Patient) CatchACold() {
	p.Publish(p.Name)
}

type DoctorService struct {
	Messages []string
}

func (d *DoctorService) Notify(data string) {
	msg := fmt.Sprintf("A doctor has been called for %s", data)
	d.Messages = append(d.Messages, msg)
	fmt.Println(msg)
}

func (d *DoctorService) LastMessage() string {
	if len(d.Messages) <= 0 {
		return ""
	}
	return d.Messages[len(d.Messages)-1]
}
