package domain

type Recipe struct {
	InternalID      uint64
	Id              uint32 `json:"id"`
	Name            string `json:"name"`
	BrewTimeSeconds int32  `json:"brew_time_seconds,omitempty"`
	Ingredients     []*Ingredient
}
