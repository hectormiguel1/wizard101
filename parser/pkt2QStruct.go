package parser

import (
	"wizard101/messages"
	"wizard101/messages/quest"
)

var QuestMessageNumberLookup = [...]messages.ProtocolMessage{
	nil,
	quest.AcceptQuest{},
	quest.CompleteGoal{},
	quest.CompletePersona{},
	quest.CompleteQuest{},
	quest.DeclineQuest{},
	quest.InteractNpc{},
	quest.NpcInfo{},
	quest.PersonaInfo{},
	quest.QuestOffer{},
	quest.RemoveGoal{},
	quest.RemoveQuest{},
	quest.SendGoal{},
	quest.SendNpcOptions{},
	quest.SendQuest{},
}

func buildQuestStruct(packet []byte, structToBuild string) messages.ProtocolMessage {
	switch structToBuild {
	case quest.AcceptQuest{}.String():
		return buildAcceptQuest(packet)
	case quest.CompleteGoal{}.String():
		return buildCompleteGoal(packet)
	case quest.CompletePersona{}.String():
		return buildCompletePersona(packet)
	case quest.CompleteQuest{}.String():
		return buildCompleteQuest(packet)
	case quest.DeclineQuest{}.String():
		return buildDeclineQuest(packet)
	case quest.InteractNpc{}.String():
		return buildInteractNpc(packet)
	case quest.NpcInfo{}.String():
		return buildNpcInfo(packet)
	case quest.PersonaInfo{}.String():
		return buildPersonaInfo(packet)
	case quest.QuestOffer{}.String():
		return buildQuestOffer(packet)
	case quest.RemoveGoal{}.String():
		return buildRemoveGoal(packet)
	case quest.RemoveQuest{}.String():
		return buildRemoveQuest(packet)
	case quest.SendGoal{}.String():
		return buildSendGoal(packet).(quest.SendGoal)
	case quest.SendNpcOptions{}.String():
		return buildSendNpcOptions(packet)
	case quest.SendQuest{}.String():
		return buildSendQuest(packet)
	default:
		return nil
	}
}

func buildSendQuest(packet []byte) messages.ProtocolMessage {
	packet, questID := ReadGuid(packet)
	packet, questNameID := ReadUInt(packet)
	packet, questType := ReadUInt(packet)
	packet, length := ReadUShort(packet)
	packet, questTitle := readString(packet)
	packet, length = ReadUShort(packet)
	packet, questInfo := ReadStringWithLength(packet, length)
	packet, New := ReadUByte(packet)
	packet, length = ReadUShort(packet)
	packet, questMadLibs := ReadStringWithLength(packet, length)
	packet, length = ReadUShort(packet)
	packet, goalData := ReadStringWithLength(packet, length)
	packet, length = ReadUShort(packet)
	packet, rewards := ReadStringWithLength(packet, length)
	packet, length = ReadUShort(packet)
	packet, clientTags := ReadStringWithLength(packet, length)
	packet, noQuestHelper := ReadUByte(packet)
	packet, mainline := ReadUByte(packet)
	packet, skipAutoSelect := ReadUByte(packet)
	packet, petOnly := ReadUByte(packet)

	return quest.SendQuest{QuestID: questID, QuestNameID: questNameID, QuestType: questType, QuestTitle: questTitle,
		QuestInfo: questInfo, New: New, QuestMadLibs: string(questMadLibs), GoalData: string(goalData), Rewards: rewards, ClientTags: clientTags,
		NoQuestHelper: noQuestHelper, MainLine: mainline, SkipQHAutoSelect: skipAutoSelect, PetOnlyQuest: petOnly}

}

func buildSendNpcOptions(packet []byte) messages.ProtocolMessage {
	packet, mobileID := ReadGuid(packet)
	packet, length := ReadUShort(packet)
	packet, Options := ReadStringWithLength(packet, length)
	packet, reInteract := ReadInt(packet)

	return quest.SendNpcOptions{MobileID: mobileID, Options: []byte(string(Options)), ReInteract: reInteract}
}

func buildSendGoal(packet []byte) interface{} {
	packet, questID := ReadGuid(packet)
	packet, goalID := ReadGuid(packet)
	packet, goalName := ReadUInt(packet)
	packet, length := ReadUShort(packet)
	packet, goalTitle := ReadStringWithLength(packet, length)
	packet, length = ReadUShort(packet)
	packet, goalLocation := ReadStringWithLength(packet, length)
	packet, length = ReadUShort(packet)
	packet, goalImage := ReadStringWithLength(packet, length)
	packet, length = ReadUShort(packet)
	packet, goalImage2 := ReadStringWithLength(packet, length)
	packet, length = ReadUShort(packet)
	packet, personaName := ReadStringWithLength(packet, length)
	packet, goalType := ReadUByte(packet)
	packet, goalStatus := ReadUByte(packet)
	packet, goalCount := ReadInt(packet)
	packet, goalTotal := ReadInt(packet)
	packet, subscriberGoalTotal := ReadInt(packet)
	packet, useTally := ReadUByte(packet)
	packet, length = ReadUShort(packet)
	packet, tallyString := ReadStringWithLength(packet, length)
	packet, sendType := ReadInt(packet)
	packet, length = ReadUShort(packet)
	packet, goalMadLibs := readString(packet)
	packet, length = ReadUShort(packet)
	packet, clientTags := readString(packet)
	packet, length = ReadUShort(packet)
	packet, patronIcon := readString(packet)
	packet, noQuestHelper := ReadUByte(packet)
	packet, petOnly := ReadUByte(packet)

	return quest.SendGoal{QuestID: questID, GoalID: goalID, GoalNameID: goalName, GoalTitle: string(goalTitle),
		GoalLocation: string(goalLocation), GoalImage: string(goalImage), GoalImage2: string(goalImage2), PersonaName: string(personaName),
		GoalType: goalType, GoalStatus: goalStatus, GoalCount: goalCount, GoalTotal: goalTotal, SubscriberGoalTotal: subscriberGoalTotal,
		UseTally: useTally, TallyText: string(tallyString), SendType: sendType, GoalMadLibs: []byte(goalMadLibs), ClientTags: []byte(clientTags), PatronIcon: patronIcon,
		NoQuestHelper: noQuestHelper, PetOnlyQuest: petOnly}
}

func buildRemoveQuest(packet []byte) messages.ProtocolMessage {
	packet, questID := ReadGuid(packet)
	packet, npcId := ReadGuid(packet)

	return quest.RemoveQuest{QuestID: questID, NpcID: npcId}
}

func buildRemoveGoal(packet []byte) messages.ProtocolMessage {
	packet, questID := ReadGuid(packet)
	packet, goalID := ReadGuid(packet)

	return quest.RemoveGoal{QuestID: questID, GoalID: goalID}
}

func buildQuestOffer(packet []byte) messages.ProtocolMessage {
	packet, mobileID := ReadGuid(packet)
	packet, length := ReadUShort(packet)
	packet, questName := ReadStringWithLength(packet, length)
	packet, length = ReadUShort(packet)
	packet, questTitle := ReadStringWithLength(packet, length)
	packet, length = ReadUShort(packet)
	packet, questInfo := ReadStringWithLength(packet, length)
	packet, level := ReadInt(packet)
	packet, length = ReadUShort(packet)
	packet, rewards := ReadStringWithLength(packet, length)
	packet, length = ReadUShort(packet)
	packet, goalData := ReadStringWithLength(packet, length)
	packet, isMainLine := ReadUByte(packet)

	return quest.QuestOffer{MobileID: mobileID, QuestName: string(questName), QuestTitle: string(questTitle), QuestInfo: questInfo,
		Level: level, Rewards: rewards, GoalData: goalData, MainLine: isMainLine}

}

func buildPersonaInfo(packet []byte) messages.ProtocolMessage {
	packet, mobileID := ReadGuid(packet)
	packet, questID := ReadGuid(packet)
	packet, goalID := ReadGuid(packet)
	packet, _ = ReadUShort(packet)
	packet, goalHyperlink := readString(packet)

	return quest.PersonaInfo{MobileID: mobileID, QuestID: questID, GoalID: goalID, GoalHyperlink: goalHyperlink}
}

func buildNpcInfo(packet []byte) messages.ProtocolMessage {
	packet, mobileID := ReadGuid(packet)
	packet, _ = ReadUShort(packet)
	packet, name := readString(packet)
	packet, _ = ReadUShort(packet)
	packet, greeting := readString(packet)
	packet, length := ReadUShort(packet)
	packet, personaMadLibs := ReadStringWithLength(packet, length)

	return quest.NpcInfo{MobileID: mobileID, Name: name, Greeting: greeting, PersonaMadLibs: personaMadLibs}
}

func buildInteractNpc(packet []byte) messages.ProtocolMessage {
	packet, globalID := ReadGuid(packet)
	packet, _ = ReadUShort(packet)
	packet, serviceName := readString(packet)
	packet, reinteract := ReadInt(packet)
	packet, serviceIndex := ReadUInt(packet)

	return quest.InteractNpc{GlobalID: globalID, ServiceName: serviceName, ReInteract: reinteract, ServiceIndex: serviceIndex}
}

func buildDeclineQuest(packet []byte) messages.ProtocolMessage {
	packet, _ = ReadUShort(packet)
	packet, questName := readString(packet)

	return quest.DeclineQuest{QuestName: questName}
}

func buildCompleteQuest(packet []byte) messages.ProtocolMessage {
	packet, questID := ReadGuid(packet)
	packet, _ = ReadUShort(packet)
	packet, completeText := readString(packet)

	return quest.CompleteQuest{QuestID: questID, CompleteText: completeText}
}

func buildCompletePersona(packet []byte) messages.ProtocolMessage {
	packet, mobileID := ReadGuid(packet)
	packet, questID := ReadGuid(packet)
	packet, goalID := ReadGuid(packet)

	return quest.CompletePersona{MobileID: mobileID, QuestID: questID, GoalID: goalID}
}

func buildCompleteGoal(packet []byte) messages.ProtocolMessage {
	packet, questID := ReadGuid(packet)
	packet, goalID := ReadGuid(packet)
	packet, _ = ReadUShort(packet)
	packet, completeText := readString(packet)
	return quest.CompleteGoal{QuestID: questID, GoalID: goalID, CompleteText: completeText}
}

func buildAcceptQuest(packet []byte) messages.ProtocolMessage {
	packet, mobileID := ReadGuid(packet)
	packet, _ = ReadUShort(packet)
	packet, questName := readString(packet)

	return quest.AcceptQuest{MobileID: mobileID, QuestName: questName}
}
