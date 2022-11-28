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

func (p *Publisher) publish(data any) {
	for z := p.subs.Front(); z != nil; z = z.Next() {
		z.Value.(Observer).notify(data)
	}
}

type Observer interface {
	notify(data any)
}

type Patient struct {
	Publisher
	Name string
	age  int
}

func NewPatient(name string, age int) *Patient {
	return &Patient{
		Publisher: Publisher{&list.List{}},
		Name:      name,
		age:       age,
	}
}

func (p *Patient) CatchACold() {
	p.publish(p.Name)
}

type DoctorService struct {
	Messages []string
}

func (d *DoctorService) notify(data any) {
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

type Client struct {
	Publisher
	age int
}

func NewClient(age int) *Client {
	return &Client{
		Publisher: Publisher{new(list.List)},
		age:       age,
	}
}

func (p *Client) Age() int {
	return p.age
}

func (p *Client) SetAge(age int) {
	if age == p.age {
		return
	}
	p.age = age
	p.publish(PropertyChange{"age", p.age})
}

type PropertyChange struct {
	Name  string
	Value any
}

type TrafficManagement struct {
	Publisher
	Messages []string
}

func (t *TrafficManagement) notify(data any) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Value.(int) >= 16 {
			msg := "Congrats, you can drive now!"
			t.Messages = append(t.Messages, msg)
			fmt.Println(msg)
			t.Unsubscribe(t) // no longer observe age of this publisher
		}
	}
}
