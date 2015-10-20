package main

type messages struct {
	lastMessage string
}

func (m *messages) getMessages() string {
	return m.lastMessage
}
