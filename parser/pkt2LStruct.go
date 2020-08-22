package parser

import (
	"wizard101/messages"
	"wizard101/messages/login"
)

var LoginMessageNumberLookUp = [...]messages.ProtocolMessage{
	nil,
	login.CharacterInfo{},
	login.CharacterList{},
	login.CharacterSelected{},
	login.CreateCharacter{},
	login.CreateCharacterResponse{},
	login.DeleteCharacter{},
	login.DeleteCharacterResponse{},
	login.RequestCharacterList{},
	login.RequestServerList{},
	login.SelectCharacter{},
	login.ServerList{},
	login.StartCharacterList{},
	login.UserAuth{},
	login.UserAuthRsp{},
	login.UserValidate{},
	login.UserValidateRsp{},
	login.DisconnectLoginAfk{},
	login.LoginNotAfk{},
	login.LoginServerShutdown{},
	login.UserAdmitInd{},
	login.WebCharacterInfo{},
	login.UserAuthV2{},
	login.SaveCharacter{},
	login.WebAuth{},
	login.WebValidate{},
	login.ChangeCharacterName{},
	login.UserAuthV3{},
	nil,
}

func BuildLoginStruct(packet []byte, message string) messages.ProtocolMessage {
	switch message {
	case login.CharacterInfo{}.String():
		{
			return buildCharacterInfo(packet)
		}
	case login.CharacterList{}.String():
		{
			return buildCharacterList(packet)
		}
	case login.CharacterSelected{}.String():
		{
			return buildCharacterSelected(packet)
		}
	case login.CreateCharacter{}.String():
		{
			return buildCreateCharacter(packet)
		}
	case login.CreateCharacterResponse{}.String():
		{
			return buildCreateCharacterResponse(packet)
		}
	case login.DeleteCharacter{}.String():
		{
			return buildDeleteCharacter(packet)
		}
	case login.DeleteCharacterResponse{}.String():
		{
			return buildDeleteCharacterResponse(packet)
		}
	case login.RequestCharacterList{}.String():
		{
			return buildRequestCharacterList(packet)
		}
	case login.RequestServerList{}.String():
		{
			return buildRequestServerList(packet)
		}
	case login.SelectCharacter{}.String():
		{
			return buildSelectCharacter(packet)
		}
	case login.ServerList{}.String():
		{
			return buildServerList(packet)
		}
	case login.StartCharacterList{}.String():
		{
			return buildStartCharacterList(packet)
		}
	case login.UserAuth{}.String():
		{
			return buildUserAuthen(packet)
		}
	case login.UserAuthRsp{}.String():
		{
			return buildUserAuthenRsp(packet)
		}
	case login.UserValidate{}.String():
		{
			return buildUserValidate(packet)
		}
	case login.UserValidateRsp{}.String():
		{
			return buildUserValidateRsp(packet)
		}
	case login.DisconnectLoginAfk{}.String():
		{
			return buildDisconnectLoginAfk(packet)
		}
	case login.LoginNotAfk{}.String():
		{
			return buildLoginNotAfk(packet)
		}
	case login.LoginServerShutdown{}.String():
		{
			return buildLoginServerShutdown(packet)
		}
	case login.UserAdmitInd{}.String():
		{
			return buildUserAdmitInd(packet)
		}
	case login.WebCharacterInfo{}.String():
		{
			return buildWebCharacterInfo(packet)
		}
	case login.UserAuthV2{}.String():
		{
			return buildUserAuthenV2(packet)
		}
	case login.SaveCharacter{}.String():
		{
			return buildSaveCharacter(packet)
		}
	case login.WebAuth{}.String():
		{
			return buildWebAuthen(packet)
		}
	case login.WebValidate{}.String():
		{
			return buildWebValidate(packet)
		}
	case login.ChangeCharacterName{}.String():
		{
			return buildChangeCharacterName(packet)
		}
	case login.UserAuthV3{}.String():
		{
			return buildUserAuthenV3(packet)
		}
	}
	return nil
}
func buildCharacterInfo(packet []byte) messages.ProtocolMessage {
	packet, stringLength := ReadUShort(packet)
	packet, outString := ReadStringWithLength(packet, stringLength)
	return login.CharacterInfo{CharacterInformation: outString}
}
func buildCharacterList(packet []byte) messages.ProtocolMessage {
	_, errorCode := ReadUInt(packet)

	return login.CharacterList{ErrorCode: errorCode}
}
func buildCharacterSelected(packet []byte) messages.ProtocolMessage {
	packet, stringLength := ReadUShort(packet)
	packet, ip := ReadStringWithLength(packet, stringLength)
	packet, tcpPort := ReadInt(packet)
	packet, udpPort := ReadInt(packet)
	packet, stringLength = ReadUShort(packet)
	packet, key3 := ReadStringWithLength(packet, stringLength)
	packet, userID := ReadGuid(packet)
	packet, charID := ReadGuid(packet)
	packet, zoneID := ReadGuid(packet)
	packet, stringLength = ReadUShort(packet)
	packet, zoneName := ReadStringWithLength(packet, stringLength)
	packet, stringLength = ReadUShort(packet)
	packet, location := ReadStringWithLength(packet, stringLength)
	packet, slot := ReadInt(packet)
	packet, prepPhase := ReadInt(packet)
	packet, errorCode := ReadInt(packet)
	packet, stringLength = ReadUShort(packet)
	packet, loginServer := ReadStringWithLength(packet, stringLength)

	return login.CharacterSelected{
		Ip:          string(ip),
		TcpPort:     int(tcpPort),
		UdpPort:     int(udpPort),
		Key:         string(key3),
		Userid:      userID,
		CharID:      charID,
		ZoneID:      zoneID,
		ZoneName:    string(zoneName),
		Location:    string(location),
		Slot:        int(slot),
		PrepPhase:   int(prepPhase),
		ErrorCode:   int(errorCode),
		LoginServer: string(loginServer),
	}
}
func buildCreateCharacter(packet []byte) messages.ProtocolMessage {
	packet, stringLength := ReadUShort(packet)
	packet, creationInformation := ReadStringWithLength(packet, stringLength)

	return login.CreateCharacter{
		CreationInfo: creationInformation,
	}
}
func buildCreateCharacterResponse(packet []byte) messages.ProtocolMessage {
	packet, errorCode := ReadInt(packet)
	return login.CreateCharacterResponse{
		ErrorCode: errorCode,
	}
}
func buildDeleteCharacter(packet []byte) messages.ProtocolMessage {
	_, characterToDelete := ReadGuid(packet)
	return login.DeleteCharacter{CharID: characterToDelete}
}
func buildDeleteCharacterResponse(packet []byte) messages.ProtocolMessage {
	packet, errorCode := ReadInt(packet)
	return login.DeleteCharacterResponse{ErrorCode: errorCode}
}
func buildRequestCharacterList(packet []byte) messages.ProtocolMessage {
	return login.RequestCharacterList{}
}
func buildRequestServerList(packet []byte) messages.ProtocolMessage {
	return login.RequestServerList{}
}
func buildSelectCharacter(array []byte) messages.ProtocolMessage {
	packet, charID := ReadGuid(array)
	packet, stringLength := ReadUShort(packet)
	_, serverName := ReadStringWithLength(packet, stringLength)
	return login.SelectCharacter{
		CharID:     charID,
		ServerName: string(serverName),
	}
}
func buildServerList(array []byte) messages.ProtocolMessage {
	packet, stringLength := ReadUShort(array)
	_, servers := ReadStringWithLength(packet, stringLength)
	return login.ServerList{Servers: string(servers)}
}
func buildStartCharacterList(array []byte) messages.ProtocolMessage {
	array, stringLength := ReadUShort(array)
	array, loginServer := ReadStringWithLength(array, stringLength)
	_, purshasedSlots := ReadInt(array)
	return login.StartCharacterList{
		LoginServer:             string(loginServer),
		PurchasedCharacterSlots: purshasedSlots,
	}
}
func buildUserAuthen(array []byte) messages.ProtocolMessage {
	array, stringLength := ReadUShort(array)
	array, rec := ReadStringWithLength(array, stringLength)
	array, stringLength = ReadUShort(array)
	array, version := ReadStringWithLength(array, stringLength)
	array, stringLength = ReadUShort(array)
	array, revision := ReadStringWithLength(array, stringLength)
	array, stringLength = ReadUShort(array)
	array, dataRevision := ReadStringWithLength(array, stringLength)
	array, stringLength = ReadUShort(array)
	array, crc := ReadStringWithLength(array, stringLength)
	array, machineID := ReadGuid(array)
	array, stringLength = ReadUShort(array)
	array, patchClientID := ReadStringWithLength(array, stringLength)
	return login.UserAuth{
		Rec1:          string(rec),
		Version:       string(version),
		Revision:      string(revision),
		DataRevision:  string(dataRevision),
		Crc:           string(crc),
		MachineID:     machineID,
		PatchClientID: string(patchClientID),
	}
}
func buildUserAuthenRsp(array []byte) messages.ProtocolMessage {
	array, err := ReadInt(array)
	array, userID := ReadGuid(array)
	array, stringLength := ReadUShort(array)
	array, rec := ReadStringWithLength(array, stringLength)
	array, stringLength = ReadUShort(array)
	array, reason := ReadStringWithLength(array, stringLength)
	array, stringLength = ReadUShort(array)
	array, timeStamp := ReadStringWithLength(array, stringLength)
	array, payingUser := ReadInt(array)
	_, flags := ReadInt(array)
	return login.UserAuthRsp{
		ErrorCode:  err,
		UserID:     userID,
		Rec1:       rec,
		Reason:     string(reason),
		Time:       string(timeStamp),
		PayingUser: payingUser,
		Flags:      flags,
	}
}
func buildUserValidate(array []byte) messages.ProtocolMessage {
	array, userID := ReadGuid(array)
	array, stringLength := ReadUShort(array)
	array, passKey3 := ReadStringWithLength(array, stringLength)
	array, machineID := ReadGuid(array)
	array, stringLength = ReadUShort(array)
	array, locale := ReadStringWithLength(array, stringLength)
	array, stringLength = ReadUShort(array)
	array, pactchClientID := ReadStringWithLength(array, stringLength)
	return login.UserValidate{
		UserID:        userID,
		Passkey3:      string(passKey3),
		MachineID:     machineID,
		Locale:        string(locale),
		PatchClientID: string(pactchClientID),
	}
}
func buildUserValidateRsp(array []byte) messages.ProtocolMessage {
	array, err := ReadInt(array)
	array, stringLength := ReadUShort(array)
	array, reason := ReadStringWithLength(array, stringLength)
	array, userID := ReadGuid(array)
	array, stringLength = ReadUShort(array)
	array, timeStamp := ReadStringWithLength(array, stringLength)
	array, payingUser := ReadInt(array)
	_, flags := ReadInt(array)
	return login.UserValidateRsp{
		ErrorCode:  err,
		Reason:     string(reason),
		UserID:     userID,
		TimeStamp:  string(timeStamp),
		PayingUser: payingUser,
		Flags:      flags,
	}
}
func buildDisconnectLoginAfk(array []byte) messages.ProtocolMessage {
	_, warning := ReadByte(array)
	return login.DisconnectLoginAfk{
		Warning: warning,
	}
}
func buildLoginNotAfk(array []byte) messages.ProtocolMessage {
	_, badgeNameID := ReadUInt(array)
	return login.LoginNotAfk{
		BadgeNameID: badgeNameID,
	}
}
func buildLoginServerShutdown(array []byte) messages.ProtocolMessage {
	_, message := ReadUInt(array)
	return login.LoginServerShutdown{
		Message: message,
	}
}
func buildUserAdmitInd(array []byte) messages.ProtocolMessage {
	array, status := ReadInt(array)
	_, posInQueue := ReadUInt(array)
	return login.UserAdmitInd{
		Status:        status,
		PositionInQue: posInQueue,
	}
}
func buildWebCharacterInfo(array []byte) messages.ProtocolMessage {
	array, name := ReadInt(array)
	array, gender := ReadInt(array)
	_, school := ReadInt(array)
	return login.WebCharacterInfo{
		Name:   name,
		Gender: gender,
		School: school,
	}
}
func buildUserAuthenV2(array []byte) messages.ProtocolMessage {
	array, stringLength := ReadUShort(array)
	array, rec := ReadStringWithLength(array, stringLength)
	array, stringLength = ReadUShort(array)
	array, version := ReadStringWithLength(array, stringLength)
	array, stringLength = ReadUShort(array)
	array, revision := ReadStringWithLength(array, stringLength)
	array, stringLength = ReadUShort(array)
	array, dataRevision := ReadStringWithLength(array, stringLength)
	array, stringLength = ReadUShort(array)
	array, crc := ReadStringWithLength(array, stringLength)
	array, machineID := ReadGuid(array)
	array, stringLength = ReadUShort(array)
	array, locale := ReadStringWithLength(array, stringLength)
	array, stringLength = ReadUShort(array)
	array, patchClientID := ReadStringWithLength(array, stringLength)
	return login.UserAuthV2{
		Rec1:          rec,
		Version:       string(version),
		Revision:      string(revision),
		DataRevision:  string(dataRevision),
		Crc:           string(crc),
		MachineID:     machineID,
		Locale:        string(locale),
		PatchClientID: string(patchClientID),
	}
}
func buildSaveCharacter(array []byte) messages.ProtocolMessage {
	array, charID := ReadGuid(array)
	_, success := ReadUByte(array)
	return login.SaveCharacter{
		CharID:  charID,
		Success: success,
	}
}
func buildWebAuthen(array []byte) messages.ProtocolMessage {
	array, stringLength := ReadUShort(array)
	array, rec := ReadStringWithLength(array, stringLength)
	array, stringLength = ReadUShort(array)
	array, version := ReadStringWithLength(array, stringLength)
	array, stringLength = ReadUShort(array)
	array, revision := ReadStringWithLength(array, stringLength)
	array, stringLength = ReadUShort(array)
	array, dataRevision := ReadStringWithLength(array, stringLength)
	array, stringLength = ReadUShort(array)
	array, crc := ReadStringWithLength(array, stringLength)
	array, machineID := ReadGuid(array)

	return login.WebAuth{
		Rec1:         rec,
		Version:      string(version),
		Revision:     string(revision),
		DataRevision: string(dataRevision),
		Crc:          string(crc),
		MachineID:    machineID,
	}

}
func buildWebValidate(array []byte) messages.ProtocolMessage {
	array, userID := ReadGuid(array)
	array, stringLength := ReadUShort(array)
	array, passKey3 := ReadStringWithLength(array, stringLength)
	array, machineID := ReadGuid(array)
	array, stringLength = ReadUShort(array)
	array, locale := ReadStringWithLength(array, stringLength)
	return login.WebValidate{
		UserID:    userID,
		Passkey3:  passKey3,
		MachineID: machineID,
		Locale:    string(locale),
	}
}
func buildChangeCharacterName(array []byte) messages.ProtocolMessage {
	array, charID := ReadGuid(array)
	array, stringLength := ReadUShort(array)
	array, newName := ReadStringWithLength(array, stringLength)
	array, stringLength = ReadUShort(array)
	array, serverName := ReadStringWithLength(array, stringLength)
	return login.ChangeCharacterName{
		CharID:     charID,
		NewName:    string(newName),
		ServerName: string(serverName),
	}
}
func buildUserAuthenV3(array []byte) messages.ProtocolMessage {
	array, stringLength := ReadUShort(array)
	array, rec := ReadStringWithLength(array, stringLength)
	array, stringLength = ReadUShort(array)
	array, version := ReadStringWithLength(array, stringLength)
	array, stringLength = ReadUShort(array)
	array, revision := ReadStringWithLength(array, stringLength)
	array, stringLength = ReadUShort(array)
	array, dataRevision := ReadStringWithLength(array, stringLength)
	array, stringLength = ReadUShort(array)
	array, crc := ReadStringWithLength(array, stringLength)
	array, machineID := ReadGuid(array)
	array, stringLength = ReadUShort(array)
	array, locale := ReadStringWithLength(array, stringLength)
	array, stringLength = ReadUShort(array)
	array, patchClientID := ReadStringWithLength(array, stringLength)
	_, isSteamPatcher := ReadUInt(array)
	return login.UserAuthV3{
		Rec1:           rec,
		Version:        string(version),
		Revision:       string(revision),
		DataRevision:   string(dataRevision),
		Crc:            string(crc),
		MachineID:      machineID,
		Locale:         string(locale),
		PatchClientID:  string(patchClientID),
		IsSteamPatcher: isSteamPatcher,
	}
}
