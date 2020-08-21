package parser

import (
	"wizard101/messages"
	"wizard101/messages/combat"
)

var CombatMessageNumberLookUp = [...]messages.ProtocolMessage{
	nil,
	combat.AllowLeavePvp{},
	combat.CombatAFK{},
	combat.CombatActions{},
	combat.CombatAdd{},
	combat.CombatDraw{},
	combat.CombatFlee{},
	combat.CombatHand{},
	combat.CombatHealth{},
	combat.CombatLoaded{},
	combat.CombatMatchResult{},
	combat.CombatMoveSelection{},
	combat.CombatMove{},
	combat.CombatPhaseForSpectators{},
	combat.CombatPhase{},
	combat.CombatPips{},
	combat.CombatRemove{},
	combat.CombatRevealHanging{},
	combat.CombatStats{},
	combat.CombatUpFirst{},
	combat.CombatVictory{},
	combat.Duel{},
	combat.EndDuel{},
	combat.PetWillCast{},
	combat.SetPlanningPhaseTimer{},
	combat.ShowCombatUI{},
	combat.ShowPetCard{},
	combat.UpdateCombatParticipant{},
	combat.UpdateDuelTimer{},
}

func buildCombatStruct(packet []byte, message string) messages.ProtocolMessage {
	switch message {
	case combat.AllowLeavePvp{}.String():
		return buildAllowLeavePvp(packet)
	case combat.CombatAFK{}.String():
		return buildCombatAfk(packet)
	case combat.CombatActions{}.String():
		return buildCombatActions(packet)
	case combat.CombatAdd{}.String():
		return buildCombatAdd(packet)
	case combat.CombatDraw{}.String():
		return buildCombatDraw(packet)
	case combat.CombatFlee{}.String():
		return buildCombatFlee(packet)
	case combat.CombatHand{}.String():
		return buildCombatHand(packet)
	case combat.CombatHealth{}.String():
		return buildCombatHealth(packet)
	case combat.CombatLoaded{}.String():
		return buildCombatLoaded(packet)
	case combat.CombatMatchResult{}.String():
		return buildCombatMatchResult(packet)
	case combat.CombatMoveSelection{}.String():
		return buildCombatMoveSelection(packet)
	case combat.CombatMove{}.String():
		return buildCombatMove(packet)
	case combat.CombatPhaseForSpectators{}.String():
		return buildCombatPhaseForSpectators(packet)
	case combat.CombatPhase{}.String():
		return buildCombatPhase(packet)
	case combat.CombatPips{}.String():
		return buildCombatPips(packet)
	case combat.CombatRemove{}.String():
		return buildCombatRemove(packet)
	case combat.CombatRevealHanging{}.String():
		return buildCombatRevealHanging(packet)
	case combat.CombatStats{}.String():
		return buildCombatStats(packet)
	case combat.CombatUpFirst{}.String():
		return buildCombatUpFirst(packet)
	case combat.CombatVictory{}.String():
		return buildCombatVictory(packet)
	case combat.Duel{}.String():
		return buildDuel(packet)
	case combat.EndDuel{}.String():
		return buildEndDuel(packet)
	case combat.PetWillCast{}.String():
		return buildPetWillCast(packet)
	case combat.SetPlanningPhaseTimer{}.String():
		return buildSetPlanningPhaseTimer(packet)
	case combat.ShowCombatUI{}.String():
		return buildShowCombatUI(packet)
	case combat.ShowPetCard{}.String():
		return buildShowPetCard(packet)
	case combat.UpdateCombatParticipant{}.String():
		return buildUpdateCombatParticipants(packet)
	case combat.UpdateDuelTimer{}.String():
		return buildUpdateDuelTimer(packet)
	default:
		return nil
	}
}

func buildUpdateDuelTimer(packet []byte) messages.ProtocolMessage {

}

func buildShowPetCard(packet []byte) messages.ProtocolMessage {

}

func buildShowCombatUI(packet []byte) messages.ProtocolMessage {

}

func buildSetPlanningPhaseTimer(packet []byte) messages.ProtocolMessage {

}

func buildPetWillCast(packet []byte) messages.ProtocolMessage {

}

func buildEndDuel(packet []byte) messages.ProtocolMessage {
	packet, length := ReadUShort(packet)
	packet, data := ReadStringWithLength(packet, length)
	return combat.Duel{Data: data}
}

func buildDuel(packet []byte) messages.ProtocolMessage {
	packet, length := ReadUShort(packet)
	packet, data := ReadStringWithLength(packet, length)
	return combat.Duel{Data: data}
}

func buildCombatVictory(packet []byte) messages.ProtocolMessage {

}

func buildCombatUpFirst(packet []byte) messages.ProtocolMessage {

}

func buildCombatStats(packet []byte) messages.ProtocolMessage {

}

func buildCombatRevealHanging(packet []byte) messages.ProtocolMessage {

}

func buildCombatRemove(packet []byte) messages.ProtocolMessage {

}

func buildCombatPips(packet []byte) messages.ProtocolMessage {

}

func buildCombatPhase(packet []byte) messages.ProtocolMessage {
	packet, duelID := ReadGuid(packet)
	packet, newPhase := ReadUByte(packet)
	packet, length := ReadUShort(packet)
	packet, data := ReadStringWithLength(packet, length)
	return combat.CombatPhase{DuelID: duelID, NewPhase: newPhase, Data: data}
}

func buildCombatPhaseForSpectators(packet []byte) messages.ProtocolMessage {

}

func buildCombatMove(packet []byte) messages.ProtocolMessage {
	packet, moveType := ReadUByte(packet)
	packet, spellSelection := ReadUByte(packet)
	packet, spellTarget := ReadUByte(packet)

	return combat.CombatMove{MoveType: moveType, SpellSelection: spellSelection, SpellTarget: spellTarget}
}

func buildUpdateCombatParticipants(packet []byte) messages.ProtocolMessage {

}

func buildCombatMoveSelection(packet []byte) messages.ProtocolMessage {

}

func buildCombatMatchResult(packet []byte) messages.ProtocolMessage {

}

func buildCombatLoaded(packet []byte) messages.ProtocolMessage {

}

func buildCombatHealth(packet []byte) messages.ProtocolMessage {

}

func buildCombatHand(packet []byte) messages.ProtocolMessage {

}

func buildCombatFlee(packet []byte) messages.ProtocolMessage {

}

func buildCombatDraw(packet []byte) messages.ProtocolMessage {

}

func buildCombatAdd(packet []byte) messages.ProtocolMessage {

}

func buildCombatActions(packet []byte) messages.ProtocolMessage {
	packet, duelID := ReadGuid(packet)
	packet, _ = ReadUShort(packet)
	packet, actionData := readString(packet)
	return combat.CombatActions{DuelID: duelID, ActionData: []byte(actionData)}
}

func buildCombatAfk(packet []byte) messages.ProtocolMessage {

}

func buildAllowLeavePvp(packet []byte) messages.ProtocolMessage {
	return combat.AllowLeavePvp{}
}
