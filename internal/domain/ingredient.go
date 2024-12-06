package domain

type Ingredient struct {
	ID   uint32 `json:"id" db:"id"`
	UUID string `json:"uuid" db:"uuid"` //TODO: поправить схему open api(добавить во все)
	Name string `json:"name" db:"name"`
}
