package entities

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model

	ID            uint `gorm:"AUTO_INCREMENT;column:id primaryKey;"`
	Type          string
	Url           string
	Name          string
	Image         string
	Description   string
	Keywords      string
	DatePublished string
	Genre         string

	AggregateRating AggregateRating `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Actors          []Actor         `gorm:"many2many:movie_actor;"`
	Director        []Director      `gorm:"many2many:movie_director;"`
	Creator         []Creator       `gorm:"many2many:movie_creator;"`
}

type AggregateRating struct {
	gorm.Model

	ID          uint `gorm:"AUTO_INCREMENT"`
	Type        string
	RatingCount float64
	BestRating  float64
	WorstRating float64
	RatingValue float64
}

type Actor struct {
	gorm.Model

	ID    uint `gorm:"AUTO_INCREMENT"`
	Type  string
	Url   string
	Name  string
	Movie []Movie `gorm:"many2many:movie_actor"`
}

type Director struct {
	gorm.Model

	ID    uint `gorm:"AUTO_INCREMENT"`
	Type  string
	Url   string
	Name  string
	Movie []Movie `gorm:"many2many:movie_director"`
}

type Creator struct {
	gorm.Model

	ID    uint `gorm:"AUTO_INCREMENT"`
	Type  string
	Url   string
	Name  string
	Movie []Movie `gorm:"many2many:movie_creator"`
}
