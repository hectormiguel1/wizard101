package misc

import "fmt"

/*
Structures for Base (System) Messages (protocol 1)
These are the ping messages.
*/

type Ping struct {
}

func (m Ping) String() string {
	return fmt.Sprintf("%T", m)
}

type PingResponse struct {
}

func (m PingResponse) String() string {
	return fmt.Sprintf("%T", m)
}
