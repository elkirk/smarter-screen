package main

import (
	"errors"
	"testing"
)

func TestParseArg_Valid(t *testing.T) {
	val, err := parseArg("42.5", "width")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != 42.5 {
		t.Errorf("expected 42.5, got %f", val)
	}
}

func TestParseArg_NotANumber(t *testing.T) {
	_, err := parseArg("abc", "width")
	var ve *ValidationError
	if !errors.As(err, &ve) {
		t.Fatalf("expected ValidationError, got %v", err)
	}
	if ve.Field != "width" || ve.Message != "must be a number" {
		t.Errorf("unexpected error fields: %+v", ve)
	}
}

func TestParseArg_Negative(t *testing.T) {
	_, err := parseArg("-5", "height")
	var ve *ValidationError
	if !errors.As(err, &ve) {
		t.Fatalf("expected ValidationError, got %v", err)
	}
	if ve.Field != "height" || ve.Message != "must be positive" {
		t.Errorf("unexpected error fields: %+v", ve)
	}
}

func TestParseArg_Zero(t *testing.T) {
	_, err := parseArg("0", "length")
	var ve *ValidationError
	if !errors.As(err, &ve) {
		t.Fatalf("expected ValidationError, got %v", err)
	}
	if ve.Message != "must be positive" {
		t.Errorf("expected 'must be positive', got %q", ve.Message)
	}
}

func TestValidationError_Format(t *testing.T) {
	err := &ValidationError{Field: "mass", Value: "abc", Message: "must be a number"}
	expected := `invalid mass "abc": must be a number`
	if err.Error() != expected {
		t.Errorf("expected %q, got %q", expected, err.Error())
	}
}
