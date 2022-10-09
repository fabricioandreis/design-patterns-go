package behavioral

import "fmt"

// Mediator (Also known as: Intermediary, Controller) is a behavioral design pattern that lets you reduce chaotic dependencies between objects. The pattern restricts direct communications between the objects and forces them to collaborate only via a mediator object.
// Applicable when components may go in and out of a system at any time: chat room participants, players in an online game, and so on.
// https://refactoring.guru/design-patterns/mediator

type ChatUser struct {
	Name    string
	Room    *ChatRoom
	chatLog []string
}

func NewChatUser(name string) *ChatUser {
	return &ChatUser{Name: name}
}

func (u *ChatUser) Receive(sender, message string) {
	s := fmt.Sprintf("%s: %s", sender, message)
	fmt.Printf("[%s's chat session]: %s\n", u.Name, s)
	u.chatLog = append(u.chatLog, s)
}

func (u *ChatUser) Say(message string) {
	u.Room.Broadcast(u.Name, message)
}

func (u *ChatUser) PrivateMessage(who, message string) {
	u.Room.Unicast(u.Name, who, message)
}

func (u *ChatUser) ChatLog() <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for _, log := range u.chatLog {
			out <- log
		}
	}()
	return out
}

type ChatRoom struct {
	users []*ChatUser
}

func (r *ChatRoom) Broadcast(source, message string) {
	for _, u := range r.users {
		if u.Name != source {
			u.Receive(source, message)
		}
	}
}

func (r *ChatRoom) Unicast(src, dst, msg string) {
	for _, u := range r.users {
		if u.Name == dst {
			u.Receive(src, msg)
		}
	}
}

func (c *ChatRoom) Join(u *ChatUser) {
	joinMsg := u.Name + " joins the chat"
	c.Broadcast("Room", joinMsg)

	u.Room = c
	c.users = append(c.users, u)
}
