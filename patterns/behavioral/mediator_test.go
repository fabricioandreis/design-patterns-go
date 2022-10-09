package behavioral_test

import (
	"testing"

	"github.com/fabricioandreis/design-patterns-go/patterns/behavioral"
	"github.com/stretchr/testify/assert"
)

func TestMediator(t *testing.T) {
	t.Run("Should be able to use a mediator in a chat room", func(t *testing.T) {
		room := behavioral.ChatRoom{}
		john := behavioral.NewChatUser("John")
		jane := behavioral.NewChatUser("Jane")
		room.Join(john)
		room.Join(jane)
		john.Say("hi room")
		jane.Say("oh, hey john")
		simon := behavioral.NewChatUser("Simon")
		room.Join(simon)
		simon.Say("hi everyone!")
		jane.PrivateMessage("Simon", "glad you could join us!")

		expected := map[*behavioral.ChatUser][]string{
			john:  {"Room: Jane joins the chat", "Jane: oh, hey john", "Room: Simon joins the chat", "Simon: hi everyone!"},
			jane:  {"John: hi room", "Room: Simon joins the chat", "Simon: hi everyone!"},
			simon: {"Jane: glad you could join us!"},
		}
		for u, e := range expected {
			i := 0
			for log := range u.ChatLog() {
				assert.Equal(t, e[i], log)
				i++
			}
		}
	})
}
