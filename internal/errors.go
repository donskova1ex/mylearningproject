package internal

import (
	"errors"
)

var ErrRecipeNotFound = errors.New("recipe not found")
var ErrEntityNotFound = errors.New("entity not found")
var ErrEntityGetByUUID = errors.New("entity not found")
var ErrReadRows = errors.New("can not read rows")
