package domain

type Ingredient struct {
	InternalID uint64
	Id         string `json:"id"`
	Name       string `json:"name"`
}
