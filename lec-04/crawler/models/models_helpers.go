package models

import "crawler/entities"

func (movie MovieDetail) Model2Entity(entityType string) interface{} {
	switch entityType {
	case "AggregateRating":
		return entities.AggregateRating{
			Type:        movie.AggregateRating.Type,
			RatingCount: movie.AggregateRating.RatingCount,
			BestRating:  movie.AggregateRating.BestRating,
			WorstRating: movie.AggregateRating.WorstRating,
			RatingValue: movie.AggregateRating.RatingValue,
		}
	case "Actor":
		var reuslt []entities.Actor
		for _, v := range movie.Actors {
			reuslt = append(reuslt, entities.Actor{
				Type: v.Type,
				Url:  v.Url,
				Name: v.Name,
			})
		}
		return reuslt
	case "Director":
		var reuslt []entities.Director
		for _, v := range movie.Directors {
			reuslt = append(reuslt, entities.Director{
				Type: v.Type,
				Url:  v.Url,
				Name: v.Name,
			})
		}
		return reuslt
	case "Creator":
		var reuslt []entities.Creator
		for _, v := range movie.Creator {
			reuslt = append(reuslt, entities.Creator{
				Type: v.Type,
				Url:  v.Url,
				Name: v.Name,
			})
		}
		return reuslt
	case "Genre":
		return movie.Genre
	default:
		return nil
	}
}
