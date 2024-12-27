package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSlice(t *testing.T) {
	var i []int
	require.Panics(t, func() {
		fmt.Printf("%v", i[:1])
	})
}
