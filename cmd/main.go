package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"

	dispatcher "smarter-screen"
)

func run(args []string) error {
	if len(args) != 4 {
		return errors.New("usage: sort <width> <height> <length> <mass>")
	}

	width, err := parseArg(args[0], "width")
	if err != nil {
		return err
	}

	height, err := parseArg(args[1], "height")
	if err != nil {
		return err
	}

	length, err := parseArg(args[2], "length")
	if err != nil {
		return err
	}

	mass, err := parseArg(args[3], "mass")
	if err != nil {
		return err
	}

	result := dispatcher.Sort(width, height, length, mass)

	slog.Info("package sorted",
		"width", width,
		"height", height,
		"length", length,
		"mass", mass,
		"stack", result,
	)

	fmt.Println(result)
	return nil
}

func main() {
	if err := run(os.Args[1:]); err != nil {
		var ve *ValidationError
		if errors.As(err, &ve) {
			slog.Error("validation failed", "field", ve.Field, "value", ve.Value, "reason", ve.Message)
		} else {
			slog.Error(err.Error())
		}
		os.Exit(1)
	}
}
