package combat

import "fmt"

/*
This File defines the structures for Combat Messages (51).
*/

/* Informs the client of the contents of a Duel via serialization */
type Duel struct {
	Data []byte
}

func (m Duel) String() string {
	return fmt.Sprintf("%T", m)
}

/* Inform the client to destroy the duel */
type EndDuel struct {
	Data []byte
}

func (m EndDuel) String() string {
	return fmt.Sprintf("%T", m)
}

/* Inform the client to allow the player to leave duel before its destroyed */
type AllowLeavePvp struct {
}

func (m AllowLeavePvp) String() string {
	return fmt.Sprintf("%T", m)
}

/* Telling the client the phase of combat has changed.  */
type CombatPhase struct {
	DuelID   uint64
	NewPhase uint8
	Data     []byte
}

func (m CombatPhase) String() string {
	return fmt.Sprintf("%T", m)
}

/* Client selecting move of Combat: 0 for Pass, 1 for attack, 2 for cast on spell, 3 for flee */
type CombatMove struct {
	MoveType       uint8
	SpellSelection uint8
	SpellTarget    uint8
}

func (m CombatMove) String() string {
	return fmt.Sprintf("%T", m)
}

/*Server informing client of cards in hand */
type CombatHand struct {
	HandData          []byte //Serialized
	DeckCount         uint16
	TotalDeckCount    uint16
	TreasureCardCount uint16
	ParticipantID     uint64
}

func (m CombatHand) String() string {
	return fmt.Sprintf("%T", m)
}

/* Server informing Client of all the actions that have taken place */
type CombatActions struct {
	DuelID     uint64
	ActionData []byte
}

func (m CombatActions) String() string {
	return fmt.Sprintf("%T", m)
}

/* Server informing client new participant has joined */
type CombatAdd struct {
	DuelID          uint64
	ParticipantData []byte //Serialized character data?
}

func (m CombatAdd) String() string {
	return fmt.Sprintf("%T", m)
}

/* Inform client a participant has fled */
type CombatRemove struct {
	DuelID        uint64
	ParticipantID uint64
}

func (m CombatRemove) String() string {
	return fmt.Sprintf("%T", m)
}

/* Server informing client about teammates move selection */
type CombatMoveSelection struct {
	DuelID           uint64
	ParticipantID    uint64
	MoveType         uint8
	SpellID          int32
	SpellTargetIndex int8
	EnchantmentID    int32
	IsItemCard       uint8
	IsTreasureCard   uint8
	IsBattleCard     uint8
}

func (m CombatMoveSelection) String() string {
	return fmt.Sprintf("%T", m)
}

/* Client requesting to draw a card from sideboard */
type CombatDraw struct {
}

func (m CombatDraw) String() string {
	return fmt.Sprintf("%T", m)
}

/* Server informing client of number of pips/powerpips each participant has */
type CombatPips struct {
	DuelID  uint64
	PipData []byte //Serialized blob
}

func (m CombatPips) String() string {
	return fmt.Sprintf("%T", m)
}

/* Server informing client of health each participant has */
type CombatHealth struct {
	DuelID     uint64
	HealthData []byte //Serialized blob
}

func (m CombatHealth) String() string {
	return fmt.Sprintf("%T", m)
}

/* Server informing client the stats of each teammate */
type CombatStats struct {
	DualID    uint64
	PartID    uint64
	StatsData []byte
}

func (m CombatStats) String() string {
	return fmt.Sprintf("%T", m)
}

/* Client informing server it is done with a match */
type CombatVictory struct {
}

func (m CombatVictory) String() string {
	return fmt.Sprintf("%T", m)
}

/* Server informing client who wont the match */
type CombatMatchResult struct {
	DuelID      uint64
	WinningTeam int32
}

func (m CombatMatchResult) String() string {
	return fmt.Sprintf("%T", m)
}

/* Server informing client that a cloaked hanging effect has been revealed */
type CombatRevealHanging struct {
	PartID          uint64
	CloakEffectType int32
	SpellTemplateID int32
	EffectType      int32
	EffectAmount    int32
	DamageType      int32
	ActNum          uint8
}

func (m CombatRevealHanging) String() string {
	return fmt.Sprintf("%T", m)
}

/* Server informing Participant that someone has fled */
type CombatFlee struct {
	PartID uint64
}

func (m CombatFlee) String() string {
	return fmt.Sprintf("%T", m)
}

/* Client asking Server for Participant to be invincible */
type CombatCheat struct {
	CheatFlag uint32
}

func (m CombatCheat) String() string {
	return fmt.Sprintf("%T", m)
}

/* Notify who is first in a duel */
type CombatUpFirst struct {
	DuelID   uint64
	UpFirst  uint8
	RoundNum uint16
}

func (m CombatUpFirst) String() string {
	return fmt.Sprintf("%T", m)
}

/* Server update combat participant information */
type UpdateCombatParticipant struct {
	ObjectID         uint64
	HidePvPEnemyChat uint8
}

func (m UpdateCombatParticipant) String() string {
	return fmt.Sprintf("%T", m)
}

/* Force duel to refresh its view */
type CombatLoaded struct {
	DuelID          uint64
	RoundNum        int32
	FirstTeamToAct  int32
	ParticipantList []byte
}

func (m CombatLoaded) String() string {
	return fmt.Sprintf("%T", m)
}

/* Server telling client to pause/resume combat */
type CombatPaused struct {
	Paused int8
}

func (m CombatPaused) String() string {
	return fmt.Sprintf("%T", m)
}

/* Server telling clients (spectators) combat phase has changed */
type CombatPhaseForSpectators struct {
	DuelID     uint64
	NewPhase   uint8
	Time       uint8
	PartName1  string
	PartName2  string
	PartName3  string
	PartName4  string
	PartName5  string
	PartName6  string
	PartName7  string
	PartName8  string
	SubCircled uint32
	TeamName0  uint32
	TeamName1  uint32
}

func (m CombatPhaseForSpectators) String() string {
	return fmt.Sprintf("%T", m)
}

/*Send the remaining time (in seconds) for duel */
type UpdateDuelTimer struct {
	DuelID        uint64
	RemainingTime uint32
}

func (m UpdateDuelTimer) String() string {
	return fmt.Sprintf("%T", m)
}

/*Server informing client of players combat afk status */
type CombatAFK struct {
	DuelID      uint64
	IsCombatAFK uint8
}

func (m CombatAFK) String() string {
	return fmt.Sprintf("%T", m)
}

/* Server informing Client to show combat UI */
type ShowCombatUI struct {
	DuelID      uint64
	AltTurn     int32
	AltTurnTeam int32
}

func (m ShowCombatUI) String() string {
	return fmt.Sprintf("%T", m)
}

/* Server telling cline t the current planning phase timer */
type SetPlanningPhaseTimer struct {
	DuelID uint64
	Time   int32
}

func (m SetPlanningPhaseTimer) String() string {
	return fmt.Sprintf("%T", m)
}

/* Server tell Client to shot Pet card */
type ShowPetCard struct {
	PetData           []byte
	CoolDown          int32
	RequirementFailed int32
}

func (m ShowPetCard) String() string {
	return fmt.Sprintf("%T", m)
}

/* Client Request pet will cast */
type PetWillCast struct {
	Target int32
}

func (m PetWillCast) String() string {
	return fmt.Sprintf("%T", m)
}
