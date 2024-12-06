package domain

type Witch struct {
	UUID string `json:"uuid" db:"uuid"`
	ID   uint32 `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
