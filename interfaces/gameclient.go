package interfaces

import "github.com/Michela-DC/review-finder-clients/model"

type GameClient interface {
	FindGames(string) []*model.Game
}
