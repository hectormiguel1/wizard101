package quest

import (
	"fmt"
)

/*
This file provides the structure implementations for the Quest Messages (protocol 52).
 */
/*
Contains information for the given Quest
 */
type SendQuest struct {
	QuestID uint64
	QuestNameID uint32
	QuestType uint32
	QuestTitle string
	QuestInfo []byte
	New uint8
	QuestMadLibs []byte
	GoalData []byte
	Rewards []byte
	ClientTags []byte
	NoQuestHelper uint8
	MainLine uint8
	SkipQHAutoSelect uint8
	PetOnlyQuest uint8
}
func (m SendQuest) String() string {
	return fmt.Sprintf("%T",m)
}

/*
Info for the given goal
 */
type SendGoal struct {
	QuestID uint64
	GoalID uint64
	GoalName uint32
	GoalTitle string
	GoalLocation string
	GoalImage string
	GoalImage2 string
	PersonaName string
	GoalType uint8
	GoalStatus uint8
	GoalCount int32
	GoalTotal int32
	SubscriberGoalTotal int32
	UseTally uint8
	TallyText string
	SendType int32
	GoalMadLibs []byte
	ClientTags []byte
	PatronIcon string
	NoQuestHelper uint8
	PetOnlyQuest uint8
}
func (m SendGoal) String() string {
	return fmt.Sprintf("%T",m)
}
/* Tell the Client to dump the quest  */
type RemoveQuest struct {
	QuestID uint64
	NpcID uint64
}
func (m RemoveQuest) String() string {
	return fmt.Sprintf("%T",m)
}
/* Tell the Client to dump the goal*/
type RemoveGoal struct {
	QuestID uint64
	GoalID uint64
}
func (m RemoveGoal) String() string {
	return fmt.Sprintf("%T",m)
}

/* The given goal has just become complete */
type CompleteGoal struct {
	QuestID uint64
	GoalID uint64
	CompleteText string
}
func (m CompleteGoal) String() string {
	return fmt.Sprintf("%T",m)
}

/* Header information on NPC */
type NpcInfo struct {
	MobileID uint64
	Name string
	Greeting string
	PersonaMadLibs []byte
}
func (m NpcInfo) String() string {
	return fmt.Sprintf("%T",m)
}

/* Info block for quest being offered by NPC */ 
type QuestOffer struct {
	MobileID uint64
	QuestName string 
	QuestTitle string 
	QuestInfo []byte
	Level int32 
	Rewards []byte
	GoalData []byte
	MainLine uint8
}
func (m QuestOffer) String() string {
	return fmt.Sprintf("%T",m)
}

/* Option to complete a person goal, sent to client */
type PersonaInfo struct {
	MobileID uint64 
	QuestID uint64 
	GoalID uint64 
	GoalHyperlink string 
}
func (m PersonaInfo) String() string {
	return fmt.Sprintf("%T",m)
}

/* Tell the server to accept offered Quest */
type AcceptQuest struct {
	MobileID uint64 
	QuestName string 
}
func (m AcceptQuest) String() string {
	return fmt.Sprintf("%T",m)
}
/* Tell the server about declined Quest */
type DeclineQuest struct {
	QuestName string 
}
func (m DeclineQuest) String() string {
	return fmt.Sprintf("%T",m)
}

/*Tell the server to complete talking to NPC */
type CompletePersona struct {
	MobileID uint64 
	QuestID uint64 
	GoalID uint64
}
func (m CompletePersona) String() string {
	return fmt.Sprintf("%T",m)
}

/* Notify the client that the quest has been completed */ 
type CompleteQuest struct {
	QuestID uint64 
	CompleteText string 
}
func (m CompleteQuest) String() string {
	return fmt.Sprintf("%T",m)
}
/* Some kind of interaction with an NPC */
type InteractNpc struct {
	GlobalID uint64 
	ServiceName string 
	ReInteract int32 
	ServiceIndex uint32 
}
func (m InteractNpc) String() string {
	return fmt.Sprintf("%T",m)
}

/* Informing the client of a set of things an NPC can do */
type SendNpcOptions struct {
	MobileID uint64 
	Options []byte
	ReInteract int32 
}
func (m SendNpcOptions) String() string {
	return fmt.Sprintf("%T",m)
}