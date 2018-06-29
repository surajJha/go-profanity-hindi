package go_profanity_hindi

import "errors"

var (
	INPUT_NOT_FOUND           = errors.New("input is missing")
	MULTIPLE_INPUTS_NOT_FOUND = errors.New("one of the inputs is missing")
)
