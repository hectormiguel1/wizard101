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

func (m LoginComplete) String() string {
	return fmt.Sprintf("%T", m)
}

/* Client-Initiated Radial Chat request */
type RequestRadialChat struct {
	Message string
}

func (m RequestRadialChat) String() string {
	return fmt.Sprintf("%T", m)
}

/* Client initiated radial owner chat request */
type RequestRadialChatOwner struct {
	Message    string
	SourceID   uint64
	Filter     int8
	SourceName string
	IsOwner    int8
}

func (m RequestRadialChatOwner) String() string {
	return fmt.Sprintf("%T", m)
}

/* Server-initiated radial chat response - send to all in range of player who requested radial chat */
type RadialChat struct {
	SourceName string
	SourceID   uint64
	Message    string
	Filter     int8
}

func (m RadialChat) String() string {
	return fmt.Sprintf("%T", m)
}

/* Client Initiated radial quick chat request */
type RequestRadialQuickChat struct {
	MessageID uint32
}

func (m RequestRadialQuickChat) String() string {
	return fmt.Sprintf("%T", m)
}

/* Server Initiated Radial quick chat response sent to all nearby players */
type RadialQuickChat struct {
	SourceName string
	SourceID   uint64
	MessageID  uint32
	Filter     int8
}

func (m RadialQuickChat) String() string {
	return fmt.Sprintf("%T", m)
}

/* Client-initiated radial quick chat request that uses extended format for message */
type RequestRadialQuickChatText struct {
	Message string
}

func (m RequestRadialQuickChatText) String() string {
	return fmt.Sprintf("%T", m)
}

/* Server Initiated radial quick chat response for messaged that use extended format */
type RadialQuickChatText struct {
	SourceName string
	SourceID   uint64
	Message    string
	Filter     int8
}

func (m RadialQuickChatText) String() string {
	return fmt.Sprintf("%T", m)
}

/* Client initiated Directed chat request */
type RequestDirectedChat struct {
	Message  string
	TargetID uint64
}

func (m RequestDirectedChat) String() string {
	return fmt.Sprintf("%T", m)
}

/* Client initiated Directed chat request by Name*/
type RequestDirectedChatByName struct {
	Message    string
	TargetName string
}

func (m RequestDirectedChatByName) String() string {
	return fmt.Sprintf("%T", m)
}

/* Response for RequestDirectedChatByName */
type DirectedChatByNameResponse struct {
	Message    string
	TargetName string
}

func (m DirectedChatByNameResponse) String() string {
	return fmt.Sprintf("%T", m)
}

/* Response to RequestDirectedChat */
type DirectedChat struct {
	SourceName string
	SourceID   uint64
	Message    string
	TargetID   uint64
	Filter     int8
}

func (m DirectedChat) String() string {
	return fmt.Sprintf("%T", m)
}

/* Server initiated Channel chat Sent from source Name */
type ChannelChat struct {
	SourceName string
	SourceID   uint64
	Message    string
	TargetID   uint64
	Filter     int8
	Flags      int32
}

func (m ChannelChat) String() string {
	return fmt.Sprintf("%T", m)
}

/* Server Notifying client about new channel instance */
type NotifyChannelInstance struct {
	RecipientID uint64
	ParentID    uint64
	ID          uint64
	Name        string
	Create      int8
}

func (m NotifyChannelInstance) String() string {
	return fmt.Sprintf("%T", m)
}

/* Client request Directed Quick chat */
type RequestDirectedQuickChat struct {
	TargetID  uint64
	MessageID uint32
}

func (m RequestDirectedQuickChat) String() string {
	return fmt.Sprintf("%T", m)
}

/* Client request Directed Quick chat (new format)*/
type RequestDirectedQuickChatText struct {
	TargetID uint64
	Message  string
}

func (m RequestDirectedQuickChatText) String() string {
	return fmt.Sprintf("%T", m)
}

/* Server Directed Quick Chat response */
type DirectedQuickChat struct {
	SourceName string
	SourceID   uint64
	MessageID  uint32
	Filter     int8
	Flags      uint32
}

func (m DirectedQuickChat) String() string {
	return fmt.Sprintf("%T", m)
}

/* Server Directed Quick Chat response (new format)*/
type DirectedQuickChatText struct {
	SourceName string
	SourceID   uint64
	Message    string
	Filter     int8
	Flags      uint32
}

func (m DirectedQuickChatText) String() string {
	return fmt.Sprintf("%T", m)
}

/* Server informing client target cant restive direct chat */
type DirectedChatFail struct {
}

func (m DirectedChatFail) String() string {
	return fmt.Sprintf("%T", m)
}

/* Transferring zones in the same server */
type ZoneTransfer struct {
	ZoneName            string
	ZoneID              uint64
	Slot                int32
	DynamicZoneID       uint32
	DynamicServerProcID uint32
	ZoneCounter         int8
	TransitionID        uint32
	SessionID           uint64
	SessionSlot         int32
}

func (m ZoneTransfer) String() string {
	return fmt.Sprintf("%T", m)
}

/* Test if server transfer can be done */
type ServerTransfer struct {
	Ip              string
	TcpPort         int32
	UdpPort         int32
	Key             int32
	UserID          uint64
	CharID          uint64
	ZoneName        string
	ZoneID          uint64
	Location        string
	Slot            int32
	SessionID       uint64
	SessionSlot     int32
	TargetPlayerID  uint64
	FallBackIP      string
	FallBackTCPPort int32
	FallBackUDPPort int32
	FallBackKey     int32
	FallbackZone    string
	FallbackZoneID  uint64
	TransitionID    uint32
}

func (m ServerTransfer) String() string {
	return fmt.Sprintf("%T", m)
}

/* Release character is delayed */
type ReleaseDelay struct {
}

func (m ReleaseDelay) String() string {
	return fmt.Sprintf("%T", m)
}

/* Server has detected an invalid movement, correcting client */
type MoveCorrection struct {
	LocationX float32
	LocationY float32
	LocationZ float32
	Direction float32
}

func (m MoveCorrection) String() string {
	return fmt.Sprintf("%T", m)
}

/* Request code for secure chat with friend */
type RequestChatCode struct {
	ListOwnerGID uint64
	Code         string
	CodeTime     int32
	SpecialCode  int32
	Forwarded    int8
	NameBlob     string
}

func (m RequestChatCode) String() string {
	return fmt.Sprintf("%T", m)
}

/* Send generated code for secure chat */
type SendCode struct {
	ListOwnerID uint64
	Code        string
	Error       uint32
	UseSuccess  uint64
	CreatorName string
}

func (m SendCode) String() string {
	return fmt.Sprintf("%T", m)
}

/* Ask that chat server to use the chat code for secure communication */
type UseChatCode struct {
	ListOwnerGID uint64
	BuddyID      uint64
	Code         string
	Forwarded    int8
}

func (m UseChatCode) String() string {
	return fmt.Sprintf("%T", m)
}

/* Ignore List as a blob */
type IgnoreList struct {
	ListOwnerGID uint64
	ListData     string
	Add          int8
}

func (m IgnoreList) String() string {
	return fmt.Sprintf("%T", m)
}

/* Setting best friend on buddy */
type BestFriend struct {
	LisOwnerGID  uint64
	BuddyID      uint64
	Forwarded    int8
	FriendSymbol int8
}

func (m BestFriend) String() string {
	return fmt.Sprintf("%T", m)
}
