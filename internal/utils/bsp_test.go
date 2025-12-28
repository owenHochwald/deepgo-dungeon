package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertNodeValid(t *testing.T, n *Node, maxLevel int) {
	if n == nil {
		return
	}

	assert.Greater(t, n.Container.W, 0)
	assert.Greater(t, n.Container.H, 0)
	assert.LessOrEqual(t, n.Level, 3)

	if n.Left != nil && n.Right != nil {
		assert.Equal(t, n.Level+1, n.Left.Level)
		assert.Equal(t, n.Level+1, n.Right.Level)

		widthsMatch := n.Left.Container.W+n.Right.Container.W == n.Container.W
		heightsMatch := n.Left.Container.H+n.Right.Container.H == n.Container.H
		// sneaky xor trick from CPSC 121
		assert.True(t, (widthsMatch || heightsMatch) && !(widthsMatch && heightsMatch))

		if widthsMatch {
			assert.Equal(t, n.Container.H, n.Left.Container.H)
			assert.Equal(t, n.Container.H, n.Right.Container.H)
			assert.Equal(t, n.Left.Container.X+n.Left.Container.W, n.Right.Container.X)
		} else {
			assert.Equal(t, n.Container.W, n.Left.Container.W)
			assert.Equal(t, n.Container.W, n.Right.Container.W)
			assert.Equal(t, n.Left.Container.Y+n.Left.Container.H, n.Right.Container.Y)
		}

	}

	assertNodeValid(t, n.Left, maxLevel)
	assertNodeValid(t, n.Right, maxLevel)
}

func TestNode_Split(t *testing.T) {
	tests := []struct {
		name     string
		w, h     int
		minSize  int
		maxLevel int
	}{
		{"Basic 3 level", 100, 100, 10, 3},
		{"Tall room", 50, 150, 10, 3},
		{"Wide room", 150, 50, 10, 3},
		{"Small room", 25, 25, 10, 2}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NewNode(0, 0, tt.w, tt.h, 0)
			n.Split(tt.minSize, tt.maxLevel) // Just one call!
			assertNodeValid(t, n, tt.maxLevel)
		})
	}
}
