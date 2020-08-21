package misc

import "fmt"

/* Keep Alive */
type KeepAlive struct {

}
func (m KeepAlive) String() string {
	return fmt.Sprintf("%T",m)
}

/* Keep Alive Response */
type KeepAliveResponse struct {

}
func (m KeepAliveResponse) String() string {
	return fmt.Sprintf("%T",m)
}