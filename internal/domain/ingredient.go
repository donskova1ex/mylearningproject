package domain

type Ingredient struct {
	ID   uint32
	UUID string `json:"uuid" db:"uuid"` //TODO: поправить схему open api(добавить во все)
	Name string `json:"name" db:"name"`
}
