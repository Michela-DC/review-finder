package interfaces

import "github.com/Michela-DC/review-finder-clients/model"

type AnimeClient interface {
	FindAnimes(string) []*model.Anime
}
