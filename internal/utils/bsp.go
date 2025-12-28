package utils

import "math/rand"

type Rect struct {
	X, Y, W, H int
}
type Node struct {
	Level       int   // what level of the tree is this node?
	Container   Rect  // Content
	Left, Right *Node // subchildren
	Room        *Rect // Leaf nodes contains room
}

func NewNode(x, y, w, h, level int) *Node {
	return &Node{
		Level: level,
		Container: Rect{
			X: x,
			Y: y,
			W: w,
			H: h,
		},
		Left:  nil,
		Right: nil,
		Room:  nil,
	}
}

// Split splits the node into two children using random directions and sizes
func (n *Node) Split(minSize, maxLevel int) bool {
	// Base case
	if n.Level >= maxLevel || n.Container.W < minSize*2 || n.Container.H < minSize*2 {
		return false
	}

	// find split direction or force split
	splitH := rand.Float64() > 0.5
	if float64(n.Container.W) >= 1.25*float64(n.Container.H) {
		splitH = false
	} else if float64(n.Container.H) >= 1.25*float64(n.Container.W) {
		splitH = true
	}

	percent := .3 + rand.Float32()*.4

	// split between 30-70% of the container
	if splitH {
		splitPos := int(float32(n.Container.H) * percent)
		n.Left = NewNode(n.Container.X, n.Container.Y, n.Container.W, splitPos, n.Level+1)
		n.Right = NewNode(n.Container.X, n.Container.Y+splitPos, n.Container.W, n.Container.H-splitPos, n.Level+1)
	} else {
		splitPos := int(float32(n.Container.W) * percent)
		n.Left = NewNode(n.Container.X, n.Container.Y, splitPos, n.Container.H, n.Level+1)
		n.Right = NewNode(n.Container.X+splitPos, n.Container.Y, n.Container.W-splitPos, n.Container.H, n.Level+1)
	}

	// recursively split children
	n.Left.Split(minSize, maxLevel)
	n.Right.Split(minSize, maxLevel)

	return true
}

func (n *Node) isLeaf() bool {
	return n != nil && n.Left == nil && n.Right == nil
}

func CreateRoom(n *Node) {
	padding := 4

	w := n.Container.W - padding*2
	h := n.Container.H - padding*2

	n.Room = &Rect{
		X: n.Container.X + padding,
		Y: n.Container.Y + padding,
		W: w,
		H: h,
	}
}

func (n *Node) GetLeaves() []*Node {
	q := Queue{}
	q.Push(n)

	var children []*Node

	for len(q) > 0 {
		curr, err := q.Pop()

		if err != nil {
			break
		}

		if curr.isLeaf() {
			children = append(children, curr)
		} else {
			if curr.Left != nil {
				q.Push(curr.Left)
			}
			if curr.Right != nil {
				q.Push(curr.Right)
			}
		}
	}

	for _, child := range children {
		CreateRoom(child)
	}

	return children
}
