package domain

type Ingredient struct {
	InternalID uint64
	UUID       string `json:"uuid" db:"uuid"` //TODO: поправить схему open api(добавить во все)
	Name       string `json:"name" db:"name"`
}
