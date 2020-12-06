package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseSeatID(t *testing.T) {
	testString := "FBFBBFFRLR"
	result := getSeatID(strings.Split(testString, ""))
	require.Equal(t, 357, result)
}
