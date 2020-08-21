package parser

import (
	"wizard101/messages"
	"wizard101/messages/misc"
)

func buildLoginMessage(array []byte, number uint8) messages.ProtocolMessage {
	return BuildLoginStruct(array, MessageMap[int(number)])
}

func buildPatchMessage(array []byte, number uint8) messages.ProtocolMessage {
	return nil
}

func buildPetMessage(array []byte, number uint8) messages.ProtocolMessage {
	return nil

}

func buildTestManagerMessage(array []byte, number uint8) messages.ProtocolMessage {
	return nil

}

func buildWizardMessage(array []byte, number uint8) messages.ProtocolMessage {
	return nil

}

func buildMoveBehaviorMassage(array []byte, number uint8) messages.ProtocolMessage {
	return nil

}

func buildPhysicsMessage(array []byte, number uint8) messages.ProtocolMessage {
	return nil

}

func buildSobBlockMessage(array []byte, number uint8) messages.ProtocolMessage {
	return nil

}

func buildSkullRiderMessage(array []byte, number uint8) messages.ProtocolMessage {
	return nil

}

func buildDoodleDougMessage(array []byte, number uint8) messages.ProtocolMessage {
	return nil

}

func buildMG1Message(array []byte, number uint8) messages.ProtocolMessage {
	return nil

}

func buildMG2Message(array []byte, number uint8) messages.ProtocolMessage {
	return nil

}

func buildMG3Message(array []byte, number uint8) messages.ProtocolMessage {
	return nil

}

func buildMG4Message(array []byte, number uint8) messages.ProtocolMessage {
	return nil

}

func buildMG5Message(array []byte, number uint8) messages.ProtocolMessage {
	return nil

}

func buildMG6Message(array []byte, number uint8) messages.ProtocolMessage {
	return nil

}

func buildWizardHousingMessage(array []byte, number uint8) messages.ProtocolMessage {
	return nil

}

func buildDuelMessage(array []byte, number uint8) messages.ProtocolMessage {
	return nil

}

func buildQuestMessage(array []byte, number uint8) messages.ProtocolMessage {
	return buildQuestStruct(array,MessageNumberLookUp[number].String())

}

func buildWizard2Message(array []byte, number uint8) messages.ProtocolMessage {
	return nil

}

func buildCatchAKeyMessage(array []byte, number uint8) messages.ProtocolMessage {
	return nil

}

func buildGameMessage(array []byte, number uint8) messages.ProtocolMessage {
	return nil

}

func buildExtendedBaseMessage(array []byte, number uint8) messages.ProtocolMessage {
	return nil

}

func buildBaseMessage(array []byte, number uint8) messages.ProtocolMessage {
	return nil
}

func buildControlMessage(opCode uint8) messages.ProtocolMessage {
	switch opCode {
	case KeepAlive: return misc.KeepAlive{}
	case KeepAliveRsp: return misc.KeepAliveResponse{}
	default: return nil
	}
}
