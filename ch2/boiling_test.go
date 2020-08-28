package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBoilingPoint(t *testing.T) {
	t.Run("水の沸騰点を表示すること", func(t *testing.T) {
		expect := "boiling point = 212F or 100C"
		actual := getBoilingPoint
		assert.Equal(t, expect, actual)
	})
}
