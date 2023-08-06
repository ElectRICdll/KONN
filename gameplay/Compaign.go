package gameplay

import "konn/entity/basic"

var (
	CurrentCompaign Compaign
)

type Compaign struct {
	campaignID string
	manager    *Manager
	chessboard *basic.Chessboard
}
