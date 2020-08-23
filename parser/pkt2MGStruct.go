package parser

import (
	"wizard101/messages"
	"wizard101/messages/miniGames"
)

const (
	Connect = 1
	Moved   = 2
	Rewards = 3
)

func BuildMiniGameMessage(pkt []byte, msgOrder uint8) messages.ProtocolMessage {
	switch msgOrder {
	case Connect:
		return miniGames.Connect{}
	case Moved:
		return miniGames.Moved{}
	case Rewards:
		{
			pkt, score := ReadInt(pkt)
			pkt, _ = ReadUShort(pkt)
			pkt, gameName := readString(pkt)
			return miniGames.Rewards{Score: score, GameName: gameName}
		}
	default:
		return nil
	}
}
