package misc

import "fmt"

/*
These are structure definitions for extended base messages, used for passsing DML
and raw strings. Protocol 2
*/

/* Raw text message */
type RawText struct {
	Message string
}

func (m RawText) String() string {
	return fmt.Sprintf("%T", m)
}

/* Custom Dictionary Message - Prep for CUSTOMRECORD */
type CustomDict struct {
	UnknownData string
}

func (m CustomDict) String() string {
	return fmt.Sprintf("%T", m)
}

/* Custom Record Message (using format specified by CUSTOMDICT) */
type CustomRecord struct {
	UnknownData string
}

func (m CustomRecord) String() string {
	return fmt.Sprintf("%T", m)
}

/* DMLRecord streamed with dynamic structure. */
type RawRecord struct {
	UnknownData string
}

func (m RawRecord) String() string {
	return fmt.Sprintf("%T", m)
}

/* A generic messaging service. */
type ServerMessage struct {
	Modal   uint8
	Message string
}

func (m ServerMessage) String() string {
	return fmt.Sprintf("%T", m)
}

/* The server is forcing the player to disconnect */
type ForceDisconnect struct {
	Type      uint32
	TimeStamp string
	Message   string
}

func (m ForceDisconnect) String() string {
	return fmt.Sprintf("%T", m)
}
