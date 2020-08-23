package game

import "fmt"

type Game struct {
	MsgOrder int
}

func (m Game) String() string {
	return fmt.Sprintf("%T", m)
}
