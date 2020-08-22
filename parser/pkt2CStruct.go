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
	combat.CombatCheat{},
	combat.CombatDraw{},
	combat.CombatFlee{},
	combat.CombatHand{},
	combat.CombatHealth{},
	combat.CombatLoaded{},
	combat.CombatMatchResult{},
	combat.CombatMove{},
	combat.CombatMoveSelection{},
	combat.CombatPaused{},
	combat.CombatPhase{},
	combat.CombatPhaseForSpectators{},
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
	case combat.CombatCheat{}.String():
		return buildCombatCheat(packet)
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
	case combat.CombatMove{}.String():
		return buildCombatMove(packet)
	case combat.CombatMoveSelection{}.String():
		return buildCombatMoveSelection(packet)
	case combat.CombatPaused{}.String():
		return buildCombatPaused(packet)
	case combat.CombatPhase{}.String():
		return buildCombatPhase(packet)
	case combat.CombatPhaseForSpectators{}.String():
		return buildCombatPhaseForSpectators(packet)
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

func buildCombatPaused(packet []byte) messages.ProtocolMessage {
	packet, paused := ReadByte(packet)
	return combat.CombatPaused{Paused: paused}
}

func buildCombatCheat(packet []byte) messages.ProtocolMessage {
	packet, cheatFlag := ReadUInt(packet)
	return combat.CombatCheat{CheatFlag: cheatFlag}
}

func buildUpdateDuelTimer(packet []byte) messages.ProtocolMessage {
	packet, duelID := ReadGuid(packet)
	packet, remainTime := ReadUInt(packet)
	return combat.UpdateDuelTimer{DuelID: duelID, RemainingTime: remainTime}
}

func buildShowPetCard(packet []byte) messages.ProtocolMessage {
	packet, length := ReadUShort(packet)
	packet, petData := ReadStringWithLength(packet, length)
	packet, coolDown := ReadInt(packet)
	packet, requirementFailed := ReadInt(packet)

	return combat.ShowPetCard{PetData: petData, CoolDown: coolDown, RequirementFailed: requirementFailed}
}

func buildShowCombatUI(packet []byte) messages.ProtocolMessage {
	packet, duelID := ReadGuid(packet)
	packet, altTurn := ReadInt(packet)
	packet, altTurnTeam := ReadInt(packet)

	return combat.ShowCombatUI{DuelID: duelID, AltTurn: altTurn, AltTurnTeam: altTurnTeam}
}

func buildSetPlanningPhaseTimer(packet []byte) messages.ProtocolMessage {
	packet, duelID := ReadGuid(packet)
	packet, time := ReadInt(packet)
	return combat.SetPlanningPhaseTimer{DuelID: duelID, Time: time}
}

func buildPetWillCast(packet []byte) messages.ProtocolMessage {
	packet, target := ReadInt(packet)
	return combat.PetWillCast{Target: target}
}

func buildEndDuel(packet []byte) messages.ProtocolMessage {
	packet, _ = ReadUShort(packet)
	packet, data := readString(packet)
	return combat.Duel{Data: []byte(data)}
}

func buildDuel(packet []byte) messages.ProtocolMessage {
	packet, length := ReadUShort(packet)
	packet, data := ReadStringWithLength(packet, length)
	return combat.Duel{Data: data}
}

func buildCombatVictory(packet []byte) messages.ProtocolMessage {
	return combat.CombatVictory{}
}

func buildCombatUpFirst(packet []byte) messages.ProtocolMessage {
	packet, duelID := ReadGuid(packet)
	packet, upFirst := ReadUByte(packet)
	packet, roundNumber := ReadUShort(packet)

	return combat.CombatUpFirst{DuelID: duelID, UpFirst: upFirst, RoundNum: roundNumber}
}

func buildCombatStats(packet []byte) messages.ProtocolMessage {
	packet, duelID := ReadGuid(packet)
	packet, partID := ReadGuid(packet)
	packet, _ = ReadUShort(packet)
	packet, statData := readString(packet)
	return combat.CombatStats{DuelID: duelID, PartID: partID, StatsData: []byte(statData)}
}

func buildCombatRevealHanging(packet []byte) messages.ProtocolMessage {
	packet, partID := ReadGuid(packet)
	packet, cloakedEffectType := ReadInt(packet)
	packet, spellTemplateID := ReadInt(packet)
	packet, effectType := ReadInt(packet)
	packet, effectAmount := ReadInt(packet)
	packet, damageType := ReadInt(packet)
	packet, actNum := ReadUByte(packet)
	return combat.CombatRevealHanging{PartID: partID, CloakEffectType: cloakedEffectType, SpellTemplateID: spellTemplateID,
		EffectType: effectType, EffectAmount: effectAmount, DamageType: damageType, ActNum: actNum}
}

func buildCombatRemove(packet []byte) messages.ProtocolMessage {
	packet, duelID := ReadGuid(packet)
	packet, partID := ReadGuid(packet)
	return combat.CombatRemove{DuelID: duelID, ParticipantID: partID}
}

func buildCombatPips(packet []byte) messages.ProtocolMessage {
	packet, duelID := ReadGuid(packet)
	packet, _ = ReadUShort(packet)
	packet, pipData := readString(packet)
	return combat.CombatPips{DuelID: duelID, PipData: []byte(pipData)}
}

func buildCombatPhase(packet []byte) messages.ProtocolMessage {
	packet, duelID := ReadGuid(packet)
	packet, newPhase := ReadUByte(packet)
	packet, length := ReadUShort(packet)
	packet, data := ReadStringWithLength(packet, length)
	return combat.CombatPhase{DuelID: duelID, NewPhase: newPhase, Data: data}
}

func buildCombatPhaseForSpectators(packet []byte) messages.ProtocolMessage {
	packet, duelID := ReadGuid(packet)
	packet, newPhase := ReadUByte(packet)
	packet, time := ReadUByte(packet)
	packet, _ = ReadUShort(packet)
	packet, partName1 := readString(packet)
	packet, _ = ReadUShort(packet)
	packet, partName2 := readString(packet)
	packet, _ = ReadUShort(packet)
	packet, partName3 := readString(packet)
	packet, _ = ReadUShort(packet)
	packet, partName4 := readString(packet)
	packet, _ = ReadUShort(packet)
	packet, partName5 := readString(packet)
	packet, _ = ReadUShort(packet)
	packet, partName6 := readString(packet)
	packet, _ = ReadUShort(packet)
	packet, partName7 := readString(packet)
	packet, _ = ReadUShort(packet)
	packet, partName8 := readString(packet)
	packet, subCircled := ReadUInt(packet)
	packet, TeamName0 := ReadUInt(packet)
	packet, TeamName1 := ReadUInt(packet)

	return combat.CombatPhaseForSpectators{DuelID: duelID,
		NewPhase: newPhase, Time: time, PartName1: partName1, PartName2: partName2, PartName3: partName3,
		PartName4: partName4, PartName5: partName5, PartName6: partName6, PartName7: partName7, PartName8: partName8,
		SubCircled: subCircled, TeamName0: TeamName0, TeamName1: TeamName1}
}

func buildCombatMove(packet []byte) messages.ProtocolMessage {
	packet, moveType := ReadUByte(packet)
	packet, spellSelection := ReadUByte(packet)
	packet, spellTarget := ReadUByte(packet)

	return combat.CombatMove{MoveType: moveType, SpellSelection: spellSelection, SpellTarget: spellTarget}
}

func buildUpdateCombatParticipants(packet []byte) messages.ProtocolMessage {
	packet, objectID := ReadGuid(packet)
	packet, hidePvPEnemyChat := ReadUByte(packet)

	return combat.UpdateCombatParticipant{ObjectID: objectID, HidePvPEnemyChat: hidePvPEnemyChat}
}

func buildCombatMoveSelection(packet []byte) messages.ProtocolMessage {
	packet, duelID := ReadGuid(packet)
	packet, partID := ReadGuid(packet)
	packet, moveType := ReadUByte(packet)
	packet, spellID := ReadInt(packet)
	packet, spellTargetIndex := ReadByte(packet)
	packet, enchantmentID := ReadInt(packet)
	packet, isItemCard := ReadUByte(packet)
	packet, isTreasureCard := ReadUByte(packet)
	packet, isBattleCard := ReadUByte(packet)

	return combat.CombatMoveSelection{DuelID: duelID, ParticipantID: partID, MoveType: moveType, SpellID: spellID,
		SpellTargetIndex: spellTargetIndex, EnchantmentID: enchantmentID, IsItemCard: isItemCard, IsTreasureCard: isTreasureCard,
		IsBattleCard: isBattleCard}
}

func buildCombatMatchResult(packet []byte) messages.ProtocolMessage {
	packet, duelID := ReadGuid(packet)
	packet, winningTeam := ReadInt(packet)

	return combat.CombatMatchResult{DuelID: duelID, WinningTeam: winningTeam}
}

func buildCombatLoaded(packet []byte) messages.ProtocolMessage {
	packet, duelID := ReadGuid(packet)
	packet, roundNum := ReadInt(packet)
	packet, firstTeamToAct := ReadInt(packet)
	packet, _ = ReadUShort(packet)
	packet, participantList := readString(packet)

	return combat.CombatLoaded{DuelID: duelID, RoundNum: roundNum, FirstTeamToAct: firstTeamToAct, ParticipantList: []byte(participantList)}
}

func buildCombatHealth(packet []byte) messages.ProtocolMessage {
	packet, duelID := ReadGuid(packet)
	packet, _ = ReadUShort(packet)
	packet, healthData := readString(packet)
	return combat.CombatHealth{DuelID: duelID, HealthData: []byte(healthData)}
}

func buildCombatHand(packet []byte) messages.ProtocolMessage {
	packet, _ = ReadUShort(packet)
	packet, handData := readString(packet)
	packet, deckCount := ReadUShort(packet)
	packet, totalDeckCount := ReadUShort(packet)
	packet, treasureCardCount := ReadUShort(packet)
	packet, partID := ReadGuid(packet)

	return combat.CombatHand{HandData: []byte(handData), DeckCount: deckCount, TotalDeckCount: totalDeckCount,
		TreasureCardCount: treasureCardCount, ParticipantID: partID}

}

func buildCombatFlee(packet []byte) messages.ProtocolMessage {
	packet, partID := ReadGuid(packet)

	return combat.CombatFlee{PartID: partID}
}

func buildCombatDraw(packet []byte) messages.ProtocolMessage {
	return combat.CombatDraw{}
}

func buildCombatAdd(packet []byte) messages.ProtocolMessage {
	packet, duelID := ReadGuid(packet)
	packet, length := ReadUShort(packet)
	packet, participantData := ReadStringWithLength(packet, length)
	return combat.CombatAdd{DuelID: duelID, ParticipantData: participantData}

}

func buildCombatActions(packet []byte) messages.ProtocolMessage {
	packet, duelID := ReadGuid(packet)
	packet, _ = ReadUShort(packet)
	packet, actionData := readString(packet)
	return combat.CombatActions{DuelID: duelID, ActionData: []byte(actionData)}
}

func buildCombatAfk(packet []byte) messages.ProtocolMessage {
	packet, duelID := ReadGuid(packet)
	packet, isCombatAFK := ReadUByte(packet)

	return combat.CombatAFK{DuelID: duelID, IsCombatAFK: isCombatAFK}
}

func buildAllowLeavePvp(packet []byte) messages.ProtocolMessage {
	return combat.AllowLeavePvp{}
}
