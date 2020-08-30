package quests

type MadLibs struct {
	QuestName string
	Level     int32
	GoalType  string
	Location  string
	IconFile  string
}

type GoalData struct {
	GoalType        string
	Location        string
	IconFile        string
	TellyText       string
	TellyText2      string
	Count           int32
	Total           int32
	SubscriberTotal int32
	Goals           string
}

type Rewards struct {
	Gold int32
	Xp   int32
}
