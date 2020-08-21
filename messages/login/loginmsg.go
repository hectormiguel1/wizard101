package login

import "fmt"

/*
This File contains the definition of the login Protocol (7), and all its Structures.
*/
type CharacterInfo struct {
	CharacterInformation []byte
}

func (m CharacterInfo) String() string {
	return fmt.Sprintf("%T",m)
}

type CharacterList struct {
	ErrorCode uint32
}

func (m CharacterList) String() string {
	return fmt.Sprintf("%T",m)
}

type CharacterSelected struct {
	Ip          string
	TcpPort     int
	UdpPort     int
	Key         string
	Userid      uint64
	CharID      uint64
	ZoneID      uint64
	ZoneName    string
	Location    string
	Slot        int
	PrepPhase   int
	ErrorCode   int
	LoginServer string
}

func (m CharacterSelected) String() string {
	return fmt.Sprintf("%T",m)
}

type CreateCharacter struct {
	CreationInfo []byte
}

func (m CreateCharacter) String() string {
	return fmt.Sprintf("%T",m)

}

type CreateCharacterResponse struct {
	ErrorCode int32
}

func (m CreateCharacterResponse) String() string {
	return fmt.Sprintf("%T",m)

}

type DeleteCharacter struct {
	CharID uint64
}

func (m DeleteCharacter) String() string {
	return fmt.Sprintf("%T",m)

}

type DeleteCharacterResponse struct {
	ErrorCode int32
}

func (m DeleteCharacterResponse) String() string {
	return fmt.Sprintf("%T",m)

}

type RequestCharacterList struct {
}

func (m RequestCharacterList) String() string {
	return fmt.Sprintf("%T",m)

}

type RequestServerList struct {
}

func (m RequestServerList) String() string {
	return fmt.Sprintf("%T",m)

}

type SelectCharacter struct {
	CharID     uint64
	ServerName string
}

func (m SelectCharacter) String() string {
	return fmt.Sprintf("%T",m)

}

type ServerList struct {
	Servers string
}

func (m ServerList) String() string {
	return fmt.Sprintf("%T",m)

}

type StartCharacterList struct {
	LoginServer             string
	PurchasedCharacterSlots int32
}

func (m StartCharacterList) String() string {
	return fmt.Sprintf("%T",m)

}

type UserAuth struct {
	Rec1          string
	Version       string
	Revision      string
	DataRevision  string
	Crc           string
	MachineID     uint64
	PatchClientID string
}

func (m UserAuth) String() string {
	return fmt.Sprintf("%T",m)
}

type UserAuthRsp struct {
	ErrorCode  int32
	UserID     uint64
	Rec1       []byte
	Reason     string
	Time       string
	PayingUser int32
	Flags      int32
}

func (m UserAuthRsp) String() string {
	return fmt.Sprintf("%T",m)

}

type UserValidate struct {
	UserID        uint64
	Passkey3      string
	MachineID     uint64
	Locale        string
	PatchClientID string
}

func (m UserValidate) String() string {
	return fmt.Sprintf("%T",m)

}

type UserValidateRsp struct {
	ErrorCode  int32
	Reason     string
	UserID     uint64
	TimeStamp  string
	PayingUser int32
	Flags      int32
}

func (m UserValidateRsp) String() string {
	return fmt.Sprintf("%T",m)

}

type DisconnectLoginAfk struct {
	Warning int8
}

func (m DisconnectLoginAfk) String() string {
	return fmt.Sprintf("%T",m)

}

type LoginNotAfk struct {
	BadgeNameID uint32
}

func (m LoginNotAfk) String() string {
	return fmt.Sprintf("%T",m)

}

type LoginServerShutdown struct {
	Message uint32
}

func (m LoginServerShutdown) String() string {
	return fmt.Sprintf("%T",m)

}

type UserAdmitInd struct {
	Status        int32
	PositionInQue uint32
}

func (m UserAdmitInd) String() string {
	return fmt.Sprintf("%T",m)

}

type WebCharacterInfo struct {
	Name   int32
	Gender int32
	School int32
}

func (m WebCharacterInfo) String() string {
	return fmt.Sprintf("%T",m)

}

type UserAuthV2 struct {
	Rec1          []byte
	Version       string
	Revision      string
	DataRevision  string
	Crc           string
	MachineID     uint64
	Locale        string
	PatchClientID string
}

func (m UserAuthV2) String() string {
	return fmt.Sprintf("%T",m)

}

type SaveCharacter struct {
	CharID  uint64
	Success uint8
}

func (m SaveCharacter) String() string {
	return fmt.Sprintf("%T",m)

}

type WebAuth struct {
	Rec1         []byte
	Version      string
	Revision     string
	DataRevision string
	Crc          string
	MachineID    uint64
}

func (m WebAuth) String() string {
	return fmt.Sprintf("%T",m)

}

type WebValidate struct {
	UserID    uint64
	Passkey3  []byte
	MachineID uint64
	Locale    string
}

func (m WebValidate) String() string {
	return fmt.Sprintf("%T",m)

}

type ChangeCharacterName struct {
	CharID     uint64
	NewName    string
	ServerName string
}

func (m ChangeCharacterName) String() string {
	return fmt.Sprintf("%T",m)

}

type UserAuthV3 struct {
	Rec1           []byte
	Version        string
	Revision       string
	DataRevision   string
	Crc            string
	MachineID      uint64
	Locale         string
	PatchClientID  string
	IsSteamPatcher uint32
}

func (m UserAuthV3) String() string {
	return fmt.Sprintf("%T",m)

}