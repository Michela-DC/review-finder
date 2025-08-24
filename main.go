package main

import (
	"fmt"

	"github.com/Michela-DC/review-finder-clients/mockclient"
	"github.com/Michela-DC/review-finder/interfaces"
)

func main() {
	var (
		ac interfaces.AnimeClient = &mockclient.Animes{}
		gc interfaces.GameClient  = &mockclient.Games{}
		mc interfaces.MovieClient = &mockclient.Movies{}
	)

	printResponse(ac.FindAnimes("animename"))
	printResponse(gc.FindGames("gamename"))
	printResponse(mc.FindMovies("moviename"))
}

func printResponse[T any](in []*T) {
	for _, i := range in {
		fmt.Printf("%+v\n", i)
	}
}
