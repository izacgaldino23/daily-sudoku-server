package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArray(t *testing.T) {

	t.Run("TestMap", func(t *testing.T) {})

	t.Run("TestHas", func(t *testing.T) {})

	t.Run("TestShuffle", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5, 6}
		b := []int{1, 2, 3, 4, 5, 6}

		Shuffle(b)

		var hasDiference bool
		for i := range a {
			if a[i] != b[i] {
				hasDiference = true
			}
		}

		assert.True(t, hasDiference)
	})

}
