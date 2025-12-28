package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateGrid(t *testing.T) {
	type args struct {
		width    int
		height   int
		Rooms    []*Rect
		Hallways []Rect
	}
	tests := []struct {
		name string
		args args
		want [][]TileType
	}{
		{
			name: "empty grid with no rooms or hallways",
			args: args{
				width:    5,
				height:   3,
				Rooms:    []*Rect{},
				Hallways: []Rect{},
			},
			want: [][]TileType{
				{Void, Void, Void, Void, Void},
				{Void, Void, Void, Void, Void},
				{Void, Void, Void, Void, Void},
			},
		},
		{
			name: "grid with room and hallway creating door at intersection",
			args: args{
				width:  10,
				height: 8,
				Rooms: []*Rect{
					{X: 2, Y: 2, W: 3, H: 3},
				},
				Hallways: []Rect{
					{X: 4, Y: 3, W: 4, H: 1},
				},
			},
			want: [][]TileType{
				{Void, Void, Void, Void, Void, Void, Void, Void, Void, Void},
				{Void, Void, Void, Void, Void, Void, Void, Void, Void, Void},
				{Void, Void, Floor, Floor, Floor, Void, Void, Void, Void, Void},
				{Void, Void, Floor, Floor, Door, Floor, Floor, Floor, Void, Void},
				{Void, Void, Floor, Floor, Floor, Void, Void, Void, Void, Void},
				{Void, Void, Void, Void, Void, Void, Void, Void, Void, Void},
				{Void, Void, Void, Void, Void, Void, Void, Void, Void, Void},
				{Void, Void, Void, Void, Void, Void, Void, Void, Void, Void},
			},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GenerateGrid(tt.args.width, tt.args.height, tt.args.Rooms, tt.args.Hallways), "GenerateGrid(%v, %v, %v, %v)", tt.args.width, tt.args.height, tt.args.Rooms, tt.args.Hallways)
		})
	}
}
