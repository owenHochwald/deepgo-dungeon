package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertNodeValid(t *testing.T, n *Node) {
	if n == nil {
		return
	}

	assert.Greater(t, n.Container.W, 0)
	assert.Greater(t, n.Container.H, 0)
	assert.LessOrEqual(t, n.Level, 3)

	if n.Left != nil && n.Right != nil {
		assert.GreaterOrEqual(t, n.Container.W, n.Left.Container.W+n.Right.Container.W)
		assert.GreaterOrEqual(t, n.Container.H, n.Left.Container.H+n.Right.Container.H)
	}

	assertNodeValid(t, n.Left)
	assertNodeValid(t, n.Right)
}

func TestNode_Split(t *testing.T) {
	type fields struct {
		Level     int
		Container Rect
		Left      *Node
		Right     *Node
		Room      *Rect
	}
	type args struct {
		minSize  int
		maxLevel int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"Basic 3 level size 10 test", fields{Level: 0, Container: Rect{X: 0, Y: 0, W: 100, H: 100}}, args{minSize: 10, maxLevel: 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				Level:     tt.fields.Level,
				Container: tt.fields.Container,
				Left:      tt.fields.Left,
				Right:     tt.fields.Right,
				Room:      tt.fields.Room,
			}
			n.Split(tt.args.minSize, tt.args.maxLevel)
			assertNodeValid(t, n)

		})
	}
}
