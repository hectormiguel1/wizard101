package game

import "fmt"

/* Contains function definitions for Game Messages Protocol 5 */

/* Attach a socket to the game object */
type Attach struct {
	GameObjectID   uint64
	LoginKey       string
	UserID         uint64
	CharID         uint64
	ZoneName       string
	Location       string
	TargetPlayerID uint64
	ZoneID         uint64
	Slot           int32
	SessionID      uint64
	SessionSlot    uint64
	PassKey        string
	ReAttach       int8
	Retry          int8
	Locale         string
}

func (m Attach) String() string {
	return fmt.Sprintf("%T", m)
}

/* Client needs to fallback to its previous location */
type AttachFailed struct {
	Error        uint32
	Rejected     uint32
	NoDisconnect uint32
}

func (m AttachFailed) String() string {
	return fmt.Sprintf("%T", m)
}

/* Notify the client that the zone servers is shutting down */
type ServerShutdown struct {
	Message uint32
}

func (m ServerShutdown) String() string {
	return fmt.Sprintf("%T", m)
}

/* A move update from mobile, transferred from client to server */
type ClientMove struct {
	LocationX   uint16
	LocationY   uint16
	LocationZ   uint16
	Direction   int8
	ZoneCounter int8
}

func (m ClientMove) String() string {
	return fmt.Sprintf("%T", m)
}

/* A move update from mobile, transferred from server to client  */
type ServerMove struct {
	LocationX   uint16
	LocationY   uint16
	LocationZ   uint16
	Direction   int8
	ZoneCounter int8
	MobileID    uint64
}

func (m ServerMove) String() string {
	return fmt.Sprintf("%T", m)
}

/* A force teleport message, transferred from server to client */
type ServerTeleport struct {
	LocationX   uint16
	LocationY   uint16
	LocationZ   uint16
	Direction   int8
	ZoneCounter int8
	MobileID    uint64
}

func (m ServerTeleport) String() string {
	return fmt.Sprintf("%T", m)
}

/* Inform the client to create a new object */
type NewObject struct {
	Data []byte //Serialized object data
}

func (m NewObject) String() string {
	return fmt.Sprintf("%T", m)
}

/* Informs the client to create an already-existing game object at a particular location */
type AddObject struct {
	GameObjectID uint64
	LocationX    float32
	LocationY    float32
	LocationZ    float32
	Direction    float32
	Name         string
	TemplateID   uint64
}

func (m AddObject) String() string {
	return fmt.Sprintf("%T", m)
}

/* Special variation for a custom despawn of a game object */
type DeleteObject struct {
	GameObjectID uint64
	Data         []byte //Serialized object data
}

func (m DeleteObject) String() string {
	return fmt.Sprintf("%T", m)
}

/* Inform a client that an object has moved out of proximity */
type RemoveObject struct {
	GameObjectID uint64
}

func (m RemoveObject) String() string {
	return fmt.Sprintf("%T", m)
}

/* Inform the client that an object has an icon above its head */
type WizBang struct {
	GameObjectID uint64
	WizBangID    uint32
}

func (m WizBang) String() string {
	return fmt.Sprintf("%T", m)
}

/* Informs the client when login is finished, informs of its GID and initial location */
type LoginComplete struct {
	Zone                        string
	Data                        []byte //Serialized Blob
	ServerTime                  uint32
	ZoneID                      uint64
	DynamicZoneID               uint32
	DynamicServerProcID         int32
	Permissions                 uint32
	IsCSR                       int32
	ZoneServer                  string
	TestServer                  int8
	AltMusicFile                uint32
	ShowSubscriberIcons         int8
	SubscriberCrownPricePercent int32
	UseFriendFinder             int32
	RealmName                   string
	IsBossMarkZone              int8
	CriticalObjects             []byte //Object blobs.
}
