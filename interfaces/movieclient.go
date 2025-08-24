package interfaces

import "github.com/Michela-DC/review-finder-clients/model"

type MovieClient interface {
	FindMovies(string) []*model.Movie
}
