package main

import (
	"fmt"
	"strconv"
)

type ValidationError struct {
	Field   string
	Value   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("invalid %s %q: %s", e.Field, e.Value, e.Message)
}

func parseArg(raw string, field string) (float64, error) {
	val, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return 0, &ValidationError{Field: field, Value: raw, Message: "must be a number"}
	}
	if val <= 0 {
		return 0, &ValidationError{Field: field, Value: raw, Message: "must be positive"}
	}
	return val, nil
}
