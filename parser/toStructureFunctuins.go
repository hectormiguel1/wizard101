package parser

import (
	"wizard101/messages"
	"wizard101/messages/game"
	"wizard101/messages/misc"
	"wizard101/messages/wizard"
)

func buildLoginMessage(array []byte, number uint8) messages.ProtocolMessage {
	return BuildLoginStruct(array, LoginMessageNumberLookUp[int(number)].String())
}

func buildPatchMessage(array []byte, number uint8) messages.ProtocolMessage {
	return BuildPatchMessages(array, number)
}

func buildPetMessage(array []byte, number uint8) messages.ProtocolMessage {
	return nil

}

func buildTestManagerMessage(array []byte, number uint8) messages.ProtocolMessage {
	return nil

}

func buildWizardMessage(array []byte, number uint8) messages.ProtocolMessage {
	return wizard.Wizard{int(number)}

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
	return BuildMiniGameMessage(array, number)

}

func buildDoodleDougMessage(array []byte, number uint8) messages.ProtocolMessage {
	return BuildMiniGameMessage(array, number)

}

func buildMG1Message(array []byte, number uint8) messages.ProtocolMessage {
	return BuildMiniGameMessage(array, number)

}

func buildMG2Message(array []byte, number uint8) messages.ProtocolMessage {
	return BuildMiniGameMessage(array, number)

}

func buildMG3Message(array []byte, number uint8) messages.ProtocolMessage {
	return BuildMiniGameMessage(array, number)

}

func buildMG4Message(array []byte, number uint8) messages.ProtocolMessage {
	return BuildMiniGameMessage(array, number)

}

func buildMG5Message(array []byte, number uint8) messages.ProtocolMessage {
	return BuildMiniGameMessage(array, number)

}

func buildMG6Message(array []byte, number uint8) messages.ProtocolMessage {
	return BuildMiniGameMessage(array, number)

}

func buildWizardHousingMessage(array []byte, number uint8) messages.ProtocolMessage {
	return nil

}

func buildDuelMessage(array []byte, number uint8) messages.ProtocolMessage {
	return buildCombatStruct(array, CombatMessageNumberLookUp[number].String())

}

func buildQuestMessage(array []byte, number uint8) messages.ProtocolMessage {
	return buildQuestStruct(array, QuestMessageNumberLookup[number].String())

}

func buildWizard2Message(array []byte, number uint8) messages.ProtocolMessage {
	return wizard.Wizard2{int(number)}

}

func buildCatchAKeyMessage(array []byte, number uint8) messages.ProtocolMessage {
	return BuildMiniGameMessage(array, number)

}

func buildGameMessage(array []byte, number uint8) messages.ProtocolMessage {
	return game.Game{int(number)}

}

func buildExtendedBaseMessage(array []byte, number uint8) messages.ProtocolMessage {
	return nil

}

func buildBaseMessage(array []byte, number uint8) messages.ProtocolMessage {
	switch number {
	case Ping:
		return misc.Ping{}
	case PingRsp:
		return misc.PingResponse{}
	default:
		return nil
	}
}

func buildControlMessage(opCode uint8) messages.ProtocolMessage {
	switch opCode {
	case KeepAlive:
		return misc.KeepAlive{}
	case KeepAliveRsp:
		return misc.KeepAliveResponse{}
	default:
		return nil
	}
}
