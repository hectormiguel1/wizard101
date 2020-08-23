package wizard

import "fmt"

type Wizard2 struct {
	MsgOrder int
}

func (m Wizard2) String() string {
	return fmt.Sprintf("%T", m)
}
