package wizard

import "fmt"

type Wizard struct {
	MsgOrder int
}

func (m Wizard) String() string {
	return fmt.Sprintf("%T", m)
}
