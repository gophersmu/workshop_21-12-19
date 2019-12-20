package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/marcusolsson/tui-go"
)

// UI represents the chat UI
type UI struct {
	messages *tui.Box
	users    *tui.Box

	tui.UI
}

// NewUI returns a new UI
func NewUI() (*UI, error) {
	users := tui.NewVBox(
		tui.NewLabel("USERS:"),
		tui.NewLabel(id+" (me)"),
	)

	usersScroll := tui.NewScrollArea(users)

	usersBox := tui.NewVBox(usersScroll)
	usersBox.SetBorder(true)

	messages := tui.NewVBox()

	messagesScroll := tui.NewScrollArea(messages)
	messagesScroll.SetAutoscrollToBottom(true)

	messagesBox := tui.NewVBox(messagesScroll)
	messagesBox.SetBorder(true)

	input := tui.NewEntry()
	input.SetFocused(true)
	input.SetSizePolicy(tui.Expanding, tui.Maximum)

	inputBox := tui.NewHBox(input)
	inputBox.SetBorder(true)
	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)

	chat := tui.NewVBox(messagesBox, inputBox)
	chat.SetSizePolicy(tui.Expanding, tui.Expanding)

	input.OnSubmit(func(e *tui.Entry) {
		// clear input box
		defer input.SetText("")

		msg := e.Text()
		msg = strings.TrimSpace(msg)
		if msg == "" {
			// ignore empty messages
			return
		}

		// send message to peers
		go sendMessage(msg)

		// display message on our side
		ui.addMessage(id, msg)
	})

	root := tui.NewHBox(usersBox, chat)

	ui, err := tui.New(root)
	if err != nil {
		return nil, err
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })

	return &UI{messages, users, ui}, nil
}

func (u *UI) addMessage(author string, message string) {
	u.Update(func() {
		u.messages.Append(tui.NewHBox(
			tui.NewLabel(time.Now().Format("15:04")),
			tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("<%s>", author))),
			tui.NewLabel(message),
			tui.NewSpacer(),
		))
	})
}

func (u *UI) addUsers(user string) {
	u.Update(func() {
		u.users.Append(tui.NewLabel(user))
	})
}
