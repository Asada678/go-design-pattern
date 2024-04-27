package models

import "time"

type DogBreed struct {
	Id               int    `json:"id"`
	Breed            string `json:"bread"`
	WeightLowLbs     int    `json:"weight_low_lds"`
	WeightHighLbs    int    `json:"weight_high_lds"`
	AverageWeight    int    `json:"average_weight"`
	Lifespan         int    `json:"lifespan"`
	Details          string `json:"details"`
	AlternateNames   string `json:"alternate_names"`
	GeographicOrigin string `json:"geographic_origin"`
}

type CatBreed struct {
	Id               int    `json:"id"`
	Breed            string `json:"bread"`
	WeightLowLbs     int    `json:"weight_low_lds"`
	WeightHighLbs    int    `json:"weight_high_lds"`
	AverageWeight    int    `json:"average_weight"`
	Lifespan         int    `json:"lifespan"`
	Details          string `json:"details"`
	AlternateNames   string `json:"alternate_names"`
	GeographicOrigin string `json:"geographic_origin"`
}

type Dog struct {
	Id               int       `json:"id"`
	DogName          string    `json:"dog_name"`
	BreedId          int       `json:"breed_id"`
	BreederId        int       `json:"breeder_id"`
	Color            string    `json:"color"`
	DateOfBirth      time.Time `json:"date_of_birth"`
	SpayedOrNeutered int       `json:"spayed_or_neutered"`
	Description      string    `json:"description"`
	Weight           int       `json:"weight"`
	Breed            DogBreed  `json:"breed"`
	Breeder          Breeder   `json:"breeder"`
}

type Cat struct {
	Id               int       `json:"id"`
	CatName          string    `json:"cat_name"`
	BreedId          int       `json:"breed_id"`
	BreederId        int       `json:"breeder_id"`
	Color            string    `json:"color"`
	DateOfBirth      time.Time `json:"date_of_birth"`
	SpayedOrNeutered int       `json:"spayed_or_neutered"`
	Description      string    `json:"description"`
	Weight           int       `json:"weight"`
	Breed            CatBreed  `json:"breed"`
	Breeder          Breeder   `json:"breeder"`
}

type Breeder struct {
	Id          int         `json:"id"`
	BreederName string      `json:"breeder_name"`
	Address     string      `json:"address"`
	City        string      `json:"city"`
	ProvState   string      `json:"prov_state"`
	Country     string      `json:"country"`
	Zip         string      `json:"zip"`
	Phone       string      `json:"phone"`
	Email       string      `json:"email"`
	Active      int         `json:"active"`
	DogBreeds   []*DogBreed `json:"dog_breeds"`
	CatBreeds   []*CatBreed `json:"cat_breeds"`
}

type Pet struct {
	Species     string `json:"species"`
	Breed       string `json:"breed"`
	MinWeight   int    `json:"min_weight"`
	MaxWeight   int    `json:"max_weight"`
	Description string `json:"description"`
	Lifespan    string `json:"lifespan"`
}
