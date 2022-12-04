package behavioral

import "fmt"

// State is a behavioral design pattern that lets an object alter its behavior when its internal state changes. It appears as if the object changed its class.
// https://refactoring.guru/design-patterns/state

// An object transitions from one state to another (something needs to trigger a transition)
// A formalized construct which managers state and transitions is called a state machine.

// 1: States as concrete types implementing an interface
type Switch struct {
	State
}

func NewSwitch() *Switch {
	return &Switch{NewOffState()}
}

func (s *Switch) On() {
	s.State.On(s)
}

func (s *Switch) Off() {
	s.State.Off(s)
}

type State interface {
	On(sw *Switch)
	Off(sw *Switch)
}

type BaseState struct{}

func (b *BaseState) On(sw *Switch) {
	fmt.Println("Light is already on")
}

func (b *BaseState) Off(sw *Switch) {
	fmt.Println("Light is already off")
}

type OnState struct {
	BaseState
}

func NewOnState() *OnState {
	fmt.Println("Light turned on")
	return &OnState{BaseState{}}
}

func (o *OnState) Off(sw *Switch) {
	fmt.Println("Turning light off...")
	sw.State = NewOffState() // Double dispatch: https://en.wikipedia.org/wiki/Double_dispatch
}

type OffState struct {
	BaseState
}

func NewOffState() *OffState {
	fmt.Println("Light turned off")
	return &OffState{BaseState{}}
}

func (o *OffState) On(sw *Switch) {
	fmt.Println("Turning light on...")
	sw.State = NewOnState()
}

// 2: States as enums and transitions encoded in a map
type PhoneState int

const (
	OffHook PhoneState = iota
	Connecting
	Connected
	OnHold
	OnHook
)

func (s PhoneState) String() string {
	switch s {
	case OffHook:
		return "OffHook"
	case Connecting:
		return "Connecting"
	case Connected:
		return "Connected"
	case OnHold:
		return "OnHold"
	case OnHook:
		return "OnHook"
	}
	return "Unknown"
}

type PhoneTrigger int

const (
	CallDialed PhoneTrigger = iota
	HungUp
	CallConnected
	PlacedOnHold
	TakenOffHold
	LeftMessage
)

func (t PhoneTrigger) String() string {
	switch t {
	case CallDialed:
		return "CallDialed"
	case HungUp:
		return "HungUp"
	case CallConnected:
		return "CallConnected"
	case PlacedOnHold:
		return "PlacedOnHold"
	case TakenOffHold:
		return "TakenOffHold"
	case LeftMessage:
		return "LeftMessage"
	}
	return "Unknown"
}

type PhoneTriggerResult struct {
	PhoneTrigger
	PhoneState
}

var Rules = map[PhoneState][]PhoneTriggerResult{
	OffHook: {
		{CallDialed, Connecting},
	},
	Connecting: {
		{HungUp, OnHook},
		{CallConnected, Connected},
	},
	Connected: {
		{LeftMessage, OnHook},
		{HungUp, OnHook},
		{PlacedOnHold, OnHold},
	},
	OnHold: {
		{TakenOffHold, Connected},
		{HungUp, OnHook},
	},
}

// 3: States as enums and transitions encoded in a switch statement
type SystemState int

const (
	Locked SystemState = iota
	Failed
	Unlocked
)

type SystemSecret struct {
	code  string
	state SystemState
}

func NewSystemSecret(code string) *SystemSecret {
	return &SystemSecret{code: code, state: Locked}
}

func (s *SystemSecret) Unlock(code string) bool {
	switch s.state {
	case Locked:
		if s.code == code {
			s.state = Unlocked
			return true
		}

		s.state = Failed
		return false
	case Failed:
		fmt.Println("A previous attempt to unlock was not successful, resetting the state do Locked")
		s.state = Locked
		return false
	case Unlocked:
		fmt.Println("Unlocked")
		return true
	}
	return false
}

func (s *SystemSecret) State() SystemState {
	return s.state
}
