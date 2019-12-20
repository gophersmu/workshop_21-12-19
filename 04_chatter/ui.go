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
	// Users widget
	users := tui.NewVBox(
		tui.NewLabel("USERS:"),
		tui.NewLabel(id+" (me)"),
	)

	usersScroll := tui.NewScrollArea(users)

	usersBox := tui.NewVBox(usersScroll)
	usersBox.SetBorder(true)

	// Messages widget
	messages := tui.NewVBox()

	messagesScroll := tui.NewScrollArea(messages)
	messagesScroll.SetAutoscrollToBottom(true)

	messagesBox := tui.NewVBox(messagesScroll)
	messagesBox.SetBorder(true)

	// Input widget
	input := tui.NewEntry()
	input.SetFocused(true)
	input.SetSizePolicy(tui.Expanding, tui.Maximum)

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
		go broadcastMessage(msg)

		// display message on our side
		messages.Append(ui.newMessageBox(id, msg))
	})

	inputBox := tui.NewHBox(input)
	inputBox.SetBorder(true)
	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)

	// Chat group widget
	chat := tui.NewVBox(messagesBox, inputBox)
	chat.SetSizePolicy(tui.Expanding, tui.Expanding)

	// Root widget
	root := tui.NewHBox(usersBox, chat)

	ui, err := tui.New(root)
	if err != nil {
		return nil, err
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })

	return &UI{messages, users, ui}, nil
}

// AddMessage adds a message in the message box
func (u *UI) AddMessage(author string, message string) {
	u.Update(func() {
		u.messages.Append(u.newMessageBox(author, message))
	})
}

// AddUser adds a user in the users box
func (u *UI) AddUser(user string) {
	u.Update(func() {
		u.users.Append(tui.NewLabel(user))
	})
}

// newMessageBox returns a new message box
func (u *UI) newMessageBox(author string, message string) *tui.Box {
	return tui.NewHBox(
		tui.NewLabel(time.Now().Format("15:04")),
		tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("<%s>", author))),
		tui.NewLabel(message),
		tui.NewSpacer(),
	)
}
