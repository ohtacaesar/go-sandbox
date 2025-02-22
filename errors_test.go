package main

import (
	"errors"
	"fmt"
	"testing"

	"golang.org/x/xerrors"
)

func TestErrorsJoin(t *testing.T) {
	parentErr := errors.New("parent error")
	childErr := fmt.Errorf("%w: child error", parentErr)

	err := xerrors.Errorf("何かに失敗しました: %w", childErr)
	fmt.Printf("%+v\n\n", err)

	err = errors.Join(err, errors.New("another error"))
	fmt.Printf("%+v\n\n", err)
}
