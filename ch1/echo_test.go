package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEcho2(t *testing.T) {
	t.Run("ブランク区切りで渡された文字列をそのままreturnすること", func(t *testing.T) {
		name := "g o l a n g"
		expect := "g o l a n g"
		actual := echo2(name)
		assert.Equal(t, expect, actual)
	})
}

func TestEcho3(t *testing.T) {
	t.Run("ブランク区切りで渡された文字列をそのままreturnすること", func(t *testing.T) {
		name := "g o l a n g"
		expect := "g o l a n g"
		actual := echo2(name)
		assert.Equal(t, expect, actual)
	})
}
