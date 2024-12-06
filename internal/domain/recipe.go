package domain

type Recipe struct {
	UUID            string `json:"uuid" db:"uuid"`
	ID              uint32 `json:"id"`
	Name            string `json:"name"`
	BrewTimeSeconds int32  `json:"brew_time_seconds,omitempty"`
	//Ingredients     []*Ingredient
}
