package entities

import (
	"fmt"
	"strings"
)

func (actor Actor) String() string {
	return fmt.Sprintf("[Type: %s, Url: %s, Name: %s]", actor.Type, actor.Url, actor.Name)
}

func (director Director) String() string {
	return fmt.Sprintf("[Type: %s, Url: %s, Name: %s]", director.Type, director.Url, director.Name)
}

func (creator Creator) String() string {
	return fmt.Sprintf("[Type: %s, Url: %s, Name: %s]", creator.Type, creator.Url, creator.Name)
}

func (rating AggregateRating) String() string {
	return fmt.Sprintf("[Type: %s, RattingCount: %f, BestRating: %f, WorstRating: %f, RatingValue: %f]",
		rating.Type, rating.RatingCount, rating.BestRating, rating.WorstRating, rating.RatingValue)
}

func (movie Movie) String() string {

	return fmt.Sprintf("[Type: %s, Url: %s, Name: %s, Image: %s, Description: %s, AggregateRating: %s, Genre: %s, Keyworks: %s, DataPublished: %s, Actors: %s, Directors: %s, Creator: %s]",
		movie.Type, movie.Url, movie.Name, movie.Image, movie.Description, movie.AggregateRating.String(), movie.Genre, movie.Keywords, movie.DatePublished, getActorStr(movie.Actors), getDirectorStr(movie.Director), getCreatorStr(movie.Creator))
}

func getActorStr(persons []Actor) string {
	var lstPersonStr = make([]string, 0)
	for _, person := range persons {
		lstPersonStr = append(lstPersonStr, person.String())
	}

	return strings.TrimSpace(strings.Join(lstPersonStr, ","))
}

func getCreatorStr(persons []Creator) string {
	var lstPersonStr = make([]string, 0)
	for _, person := range persons {
		lstPersonStr = append(lstPersonStr, person.String())
	}

	return strings.TrimSpace(strings.Join(lstPersonStr, ","))
}

func getDirectorStr(persons []Director) string {
	var lstPersonStr = make([]string, 0)
	for _, person := range persons {
		lstPersonStr = append(lstPersonStr, person.String())
	}

	return strings.TrimSpace(strings.Join(lstPersonStr, ","))
}
