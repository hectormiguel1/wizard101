package miniGames

import "fmt"

/* Structure for MiniGame Definitions
(protocol 54, 45, 42, 41, 47, 43, 46,44, 40) */

type Rewards struct {
	Score    int32
	GameName string
}

func (m Rewards) String() string {
	return fmt.Sprintf("%T", m)
}

type Moved struct {
}

func (m Moved) String() string {
	return fmt.Sprintf("%T", m)
}

type Connect struct {
}

func (m Connect) String() string {
	return fmt.Sprintf("%T", m)
}
