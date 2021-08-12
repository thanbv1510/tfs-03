package models

type MovieInfo struct {
	Name     string
	UrlMovie string
	UrlImage string
	Rating   string
}

type DBConfig struct {
	Driver   string `mapstructure:"DB_DRIVER"`
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	Username string `mapstructure:"DB_USERNAME"`
	Passwd   string `mapstructure:"DB_PASSWORD"`
	DBName   string `mapstructure:"DB_NAME"`
}

type MovieDetail struct {
	Context         string          `json:"@context"`
	Type            string          `json:"@type"`
	Url             string          `json:"url"`
	Name            string          `json:"name"`
	Image           string          `json:"image"`
	Description     string          `json:"description"`
	AggregateRating AggregateRating `json:"aggregateRating"`
	Genre           []string        `json:"genre"`
	Keywords        string          `json:"keywords"`
	DatePublished   string          `json:"datePublished"`
	Actors          []Person        `json:"actor"`
	Directors       []Person        `json:"director"`
	Creator         []Person        `json:"creator"`
}

type AggregateRating struct {
	Type        string  `json:"@type"`
	RatingCount float64 `json:"ratingCount"`
	BestRating  float64 `json:"bestRating"`
	WorstRating float64 `json:"worstRating"`
	RatingValue float64 `json:"ratingValue"`
}

type Person struct {
	Type string `json:"@type"`
	Url  string `json:"url"`
	Name string `json:"name"`
}
