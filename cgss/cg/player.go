package cg

import (
	"fmt"
)

type Player struct {
	Name  string
	Level int
	Exp   int
	Room  int

	mq chan *Message
	rq chan string
}

func NewPlayer() *Player {
	m := make(chan *Message, 1024)
	rq := make(chan string)
	player := &Player{"", 0, 0, 0, m, rq}

	go func(p *Player) {
		for {
			msg := <-p.mq
			fmt.Println(p.Name, "received message:", msg.Content)
			rq <- "receive"
		}
	}(player)

	return player
}
