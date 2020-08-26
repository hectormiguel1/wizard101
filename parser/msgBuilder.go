package parser

import (
	"wizard101/messages"
)

const (
	NoControl            = 0
	StartSignal          = 61453
	KeepAlive            = 3
	KeepAliveRsp         = 4
	Ping                 = 1
	PingRsp              = 2
	BaseMessages         = 1
	ExtendedBaseMessages = 2
	GameMessages         = 5
	LoginMessages        = 7
	PatchMessages        = 8
	PetMessages          = 9
	ScriptMessages       = 10
	TestManagerMessages  = 11
	WizardMessages       = 12
	MoveBehaviorMessages = 15
	PhysicsMessages      = 16
	SobBlocsMessages     = 25
	SkullRidersMessages  = 40
	DoodleDougMessages   = 41
	MG1Messages          = 42
	MG2Messages          = 43
	MG3Messages          = 44
	MG4Messages          = 45
	MG5Messages          = 46
	MG6Messages          = 47
	WizardHousing        = 50
	DuelMessages         = 51
	QuestMessages        = 52
	Wizard2Messages      = 53
	CatchAKeyMessages    = 54
)

type Message struct {
	sourceIP        string
	DestIP          string
	ProtocolMessage messages.ProtocolMessage
}

/* Function goes through the packet looking for the Food Signal */
func TestPacket(packet []byte) ([]byte, bool) {
	packet, startSignal := ReadUShort(packet)
	for startSignal != StartSignal && len(packet) > SizeOfShort {
		packet, startSignal = ReadUShort(packet)
	}
	if len(packet) > SizeOfShort {
		return packet, true
	} else {
		return packet, false
	}
}

func BuildMessage(packet []byte, sourceIP string, destIP string) Message {
	packet, _ = ReadUShort(packet)
	packet, isControl := ReadUByte(packet)
	packet, opCode := ReadUByte(packet)
	//Unknown bits 1 and 2  (are always zero)
	packet, _ = ReadUByte(packet)
	packet, _ = ReadUByte(packet)

	packet, protocol := ReadUByte(packet)
	packet, protocolMessageNumber := ReadUByte(packet)
	packet, _ = ReadUShort(packet)

	if isControl > NoControl {
		controlMessageTyped := buildControlMessage(opCode)
		return Message{sourceIP: sourceIP, DestIP: destIP, ProtocolMessage: controlMessageTyped}

	} else {
		decodedMessage := DecodeMessage(packet, protocol, protocolMessageNumber)
		return Message{sourceIP: sourceIP, DestIP: destIP, ProtocolMessage: decodedMessage}
	}

}

func DecodeMessage(byteArray []byte, protocolNumber uint8, messageNumber uint8) messages.ProtocolMessage {
	switch protocolNumber {
	case BaseMessages:
		return buildBaseMessage(byteArray, messageNumber)
	case ExtendedBaseMessages:
		return buildExtendedBaseMessage(byteArray, messageNumber)
	case GameMessages:
		return buildGameMessage(byteArray, messageNumber)
	case LoginMessages:
		return buildLoginMessage(byteArray, messageNumber)
	case PatchMessages:
		return buildPatchMessage(byteArray, messageNumber)
	case PetMessages:
		return buildPetMessage(byteArray, messageNumber)
	case TestManagerMessages:
		return buildTestManagerMessage(byteArray, messageNumber)
	case WizardMessages:
		return buildWizardMessage(byteArray, messageNumber)
	case MoveBehaviorMessages:
		return buildMoveBehaviorMassage(byteArray, messageNumber)
	case PhysicsMessages:
		return buildPhysicsMessage(byteArray, messageNumber)
	case SobBlocsMessages:
		return buildSobBlockMessage(byteArray, messageNumber)
	case SkullRidersMessages:
		return buildSkullRiderMessage(byteArray, messageNumber)
	case DoodleDougMessages:
		return buildDoodleDougMessage(byteArray, messageNumber)
	case MG1Messages:
		return buildMG1Message(byteArray, messageNumber)
	case MG2Messages:
		return buildMG2Message(byteArray, messageNumber)
	case MG3Messages:
		return buildMG3Message(byteArray, messageNumber)
	case MG4Messages:
		return buildMG4Message(byteArray, messageNumber)
	case MG5Messages:
		return buildMG5Message(byteArray, messageNumber)
	case MG6Messages:
		return buildMG6Message(byteArray, messageNumber)
	case WizardHousing:
		return buildWizardHousingMessage(byteArray, messageNumber)
	case DuelMessages:
		return buildDuelMessage(byteArray, messageNumber)
	case QuestMessages:
		return buildQuestMessage(byteArray, messageNumber)
	case Wizard2Messages:
		return buildWizard2Message(byteArray, messageNumber)
	case CatchAKeyMessages:
		return buildCatchAKeyMessage(byteArray, messageNumber)
	default:
		return nil
	}
}
