package models

import (
	"crawler/entities"
	"fmt"
	"strings"
)

func (movie MovieDetail) MovieModel2MovieEntity() entities.Movie {
	return entities.Movie{
		Type:          movie.Type,
		Url:           movie.Url,
		Name:          movie.Name,
		Image:         movie.Image,
		Description:   movie.Description,
		Keywords:      movie.Keywords,
		DatePublished: movie.DatePublished,
		AggregateRating: entities.AggregateRating{
			Type:        movie.AggregateRating.Type,
			RatingCount: movie.AggregateRating.RatingCount,
			BestRating:  movie.AggregateRating.BestRating,
			WorstRating: movie.AggregateRating.WorstRating,
			RatingValue: movie.AggregateRating.RatingValue,
		},
		Genre: strings.Join(movie.Genre, ","),
		Actors: func() (result []entities.Actor) {
			for _, v := range movie.Actors {
				result = append(result, entities.Actor{
					Type: v.Type,
					Url:  v.Url,
					Name: v.Name,
				})
			}
			return
		}(),
		Director: func() (result []entities.Director) {
			for _, v := range movie.Directors {
				result = append(result, entities.Director{
					Type: v.Type,
					Url:  v.Url,
					Name: v.Name,
				})
			}
			return
		}(),
		Creator: func() (result []entities.Creator) {
			for _, v := range movie.Creator {
				result = append(result, entities.Creator{
					Type: v.Type,
					Url:  v.Url,
					Name: v.Name,
				})
			}
			return
		}(),
	}
}

func (dBConfig DBConfig) String() string {
	return fmt.Sprintf("[Driver: %s, Host: %s, Port: %s, Username: %s, Passwd: %s, DBName: %s]",
		dBConfig.Driver, dBConfig.Host, dBConfig.Port, dBConfig.Username, dBConfig.Passwd, dBConfig.DBName)
}

func (person Person) String() string {
	return fmt.Sprintf("[Type: %s, Url: %s, Name: %s]", person.Type, person.Url, person.Name)
}

func (rating AggregateRating) String() string {
	return fmt.Sprintf("[Type: %s, RattingCount: %f, BestRating: %f, WorstRating: %f, RatingValue: %f]",
		rating.Type, rating.RatingCount, rating.BestRating, rating.WorstRating, rating.RatingValue)
}

func (movie MovieDetail) String() string {

	return fmt.Sprintf("[Type: %s, Url: %s, Name: %s, Image: %s, Description: %s, AggregateRating: %s, Genre: %s, Keyworks: %s, DataPublished: %s, Actors: %s, Directors: %s, Creator: %s]",
		movie.Type, movie.Url, movie.Name, movie.Image, movie.Description, movie.AggregateRating.String(), movie.Genre, movie.Keywords, movie.DatePublished, getString(movie.Actors), getString(movie.Directors), getString(movie.Creator))
}

func getString(persons []Person) string {
	var lstPersonStr = make([]string, 0)
	for _, person := range persons {
		lstPersonStr = append(lstPersonStr, person.String())
	}

	return strings.TrimSpace(strings.Join(lstPersonStr, ","))
}
