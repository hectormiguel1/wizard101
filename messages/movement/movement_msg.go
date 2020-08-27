package movement

import "fmt"

/*
Structure definitions for Movement Messages (protocol 15)
MoveBehavior Messages
*/

/*
A move update for a mobile, transferred from client to server.
*/
type MobileMove struct {
	LocationX   uint16
	LocationY   uint16
	LocationZ   uint16
	Direction   uint16
	ZoneCounter uint16
	GlobalID    uint64
}

func (m MobileMove) String() string {
	return fmt.Sprintf("%T", m)
}

/*
A move update for a mobile, transferred from client to server.
*/
type MobileMoveTime struct {
	LocationX   uint16
	LocationY   uint16
	LocationZ   uint16
	Direction   uint16
	ZoneCounter uint16
	GlobalID    uint64
	Time        float32
}

func (m MobileMoveTime) String() string {
	return fmt.Sprintf("%T", m)
}

/*
A forced teleport message, transferred from server to client (contains GID)
*/
type MobileTeleport struct {
	LocationX   uint16
	LocationY   uint16
	LocationZ   uint16
	Direction   uint16
	ZoneCounter uint16
	GlobalID    uint64
}

func (m MobileTeleport) String() string {
	return fmt.Sprintf("%T", m)
}

/* The server is changing the movement state on a mobile */
type MobileMoveState struct {
	GlobalID uint64
	NewState int8
}

func (m MobileMoveState) String() string {
	return fmt.Sprintf("%T", m)
}
