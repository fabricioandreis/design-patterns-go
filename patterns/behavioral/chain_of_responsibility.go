package behavioral

import (
	"fmt"
	"sync"
)

// Chain of Responsibility is a behavioral design pattern that lets you pass requests along a chain of handlers.
// Upon receiving a request, each handler decides either to process the request or to pass it to the next handler in the chain.
// https://refactoring.guru/design-patterns/chain-of-responsibility

type Creature struct {
	Name            string
	Attack, Defense int
}

func NewCreature(name string, attack, defense int) *Creature {
	return &Creature{name, attack, defense}
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack, c.Defense)
}

type Modifier interface {
	Add(m Modifier)
	Handle()
}

type CreatureModifier struct {
	creature *Creature
	next     Modifier
}

func NewCreatureModifier(c *Creature) *CreatureModifier {
	return &CreatureModifier{creature: c}
}

func (cm *CreatureModifier) Add(m Modifier) {
	if cm.next != nil {
		cm.next.Add(m)
	} else {
		cm.next = m
	}
}

func (cm *CreatureModifier) Handle() {
	if cm.next != nil {
		cm.next.Handle()
	}
}

type DoubleAttackModifier struct {
	CreatureModifier
}

func NewDoubleAttackModifier(c *Creature) *DoubleAttackModifier {
	return &DoubleAttackModifier{CreatureModifier{creature: c}}
}

func (d *DoubleAttackModifier) Handle() {
	fmt.Println("Doubling", d.creature.Name, "\b's attack")
	d.creature.Attack *= 2
	d.CreatureModifier.Handle()
}

type IncreaseDefenseModifier struct {
	CreatureModifier
}

func NewIncreaseDefenseModifier(c *Creature) *IncreaseDefenseModifier {
	return &IncreaseDefenseModifier{CreatureModifier{creature: c}}
}

func (i *IncreaseDefenseModifier) Handle() {
	if i.creature.Attack <= 2 {
		fmt.Println("Increasing", i.creature.Name, "\b's defense")
		i.creature.Defense++
	}
	i.CreatureModifier.Handle()
}

type NoBonusesModifier struct {
	CreatureModifier
}

func NewNoBonusesModifier(c *Creature) *NoBonusesModifier {
	return &NoBonusesModifier{CreatureModifier{creature: c}}
}

func (n *NoBonusesModifier) Handle() {
	fmt.Println("Canceling chain of modifiers")
}

// Chain of Responsibility, Mediator, Observer, Command Query Separation
type Argument int

const (
	Attack Argument = iota
	Defense
)

type Query struct {
	CreatureName string
	WhatToQuery  Argument
	Value        int
}

type BrokerObserver interface { // Observer pattern
	Handle(q *Query)
}

type BrokerObservable interface { // aka Publisher
	Subscribe(o BrokerObserver)
	Unsubscribe(o BrokerObserver)
	Fire(q *Query) // aka Publish
}

// aka Publisher
type BrokerGame struct {
	Observers sync.Map
}

func (g *BrokerGame) Subscribe(o BrokerObserver) {
	// registers BrokerObserver `o` as an observer of this BrokerGame `g`
	g.Observers.Store(o, struct{}{})
}

func (g *BrokerGame) Unsubscribe(o BrokerObserver) {
	// un-registers BrokerObserver `o` as an observer of this BrokerGame `g`
	g.Observers.Delete(o)
}

func (g *BrokerGame) Fire(q *Query) {
	g.Observers.Range(func(key, value any) bool {
		if key == nil { // Reached the end of the range
			return false
		}
		key.(BrokerObserver).Handle(q)
		return true
	})
}

type BrokerCreature struct {
	game            *BrokerGame // Mediator pattern
	Name            string
	attack, defense int
}

func NewBrokerCreature(game *BrokerGame, name string, attach int, defense int) *BrokerCreature {
	return &BrokerCreature{game, name, attach, defense}
}

func (c *BrokerCreature) Attack() int {
	q := &Query{c.Name, Attack, c.attack}
	// Publish the query to be responded
	c.game.Fire(q)
	// After every single element that wanted to process this event
	// has in fact processed it and maybe modified Value,
	// we return the final Value
	return q.Value
}

func (c *BrokerCreature) Defense() int {
	q := &Query{c.Name, Defense, c.defense}
	// Publish the query to be responded
	c.game.Fire(q)
	// After every single element that wanted to process this event
	// has in fact processed it and maybe modified Value,
	// we return the final Value
	return q.Value
}

func (c *BrokerCreature) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack(), c.Defense())
}

type BrokerCreatureModifier struct {
	game     *BrokerGame
	creature *BrokerCreature
}

func (c *BrokerCreatureModifier) Handle(q *Query) {
	// This method turns the object into an observer
	// No operation here because it only exists to compose as part of actual modifiers
}

type BrokerDoubleAttachModifier struct {
	BrokerCreatureModifier
}

func NewBrokerDoubleAttachModifier(g *BrokerGame, c *BrokerCreature) *BrokerDoubleAttachModifier {
	d := &BrokerDoubleAttachModifier{BrokerCreatureModifier{g, c}}
	g.Subscribe(d)
	return d
}

func (d *BrokerDoubleAttachModifier) Handle(q *Query) {
	if q.CreatureName == d.creature.Name && q.WhatToQuery == Attack {
		q.Value *= 2
	}
}

func (d *BrokerDoubleAttachModifier) Close() error {
	d.game.Unsubscribe(d)
	return nil
}
