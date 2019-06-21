package link

import (
	"errors"
)

var (
	ERROR_DOES_NOT_EXIST = errors.New("Does not exist.")
	ERROR_EMPTY_LINK     = errors.New("It is a empty link")
)
