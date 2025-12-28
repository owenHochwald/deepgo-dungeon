package utils

import "math/rand"

const (
	tileSize         = 14
	hallwayVariation = 4
)

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

func (n *Node) isLeaf() bool {
	return n != nil && n.Left == nil && n.Right == nil
}

func GetCenter(x, y, w, h int) (int, int) {
	return x + w/2, y + h/2
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

func (n *Node) CreateRoom() {
	padding := 4

	w := n.Container.W - padding*2
	h := n.Container.H - padding*2

	if w > 5 {
		w = rand.Intn(w-5) + 5
	}
	if h > 5 {
		h = rand.Intn(h-5) + 5
	}

	cX, cY := GetCenter(
		n.Container.X+padding,
		n.Container.Y+padding,
		n.Container.W-padding*2,
		n.Container.H-padding*2,
	)

	n.Room = &Rect{
		X: 0,
		Y: 0,
		W: 0,
		H: 0,
	}

	n.Room = &Rect{
		X: cX,
		Y: cY,
		W: w,
		H: h,
	}
}

func (n *Node) CreateHallways(hallways *[]Rect) {
	if n.Left == nil || n.Right == nil {
		return
	}

	ax, ay := GetCenter(n.Left.Container.X, n.Left.Container.Y, n.Left.Container.W, n.Left.Container.H)
	bx, by := GetCenter(n.Right.Container.X, n.Right.Container.Y, n.Right.Container.W, n.Right.Container.H)

	width := bx - ax
	if width < 0 {
		*hallways = append(*hallways, Rect{bx, ay, -width + hallwayVariation, tileSize})
	} else {
		*hallways = append(*hallways, Rect{ax, ay, width + hallwayVariation, tileSize})
	}

	height := by - ay
	if height < 0 { // Handle bottom-to-top
		*hallways = append(*hallways, Rect{bx, by, tileSize, -height + hallwayVariation})
	} else {
		*hallways = append(*hallways, Rect{bx, ay, tileSize, height + hallwayVariation})
	}

	n.Left.CreateHallways(hallways)
	n.Right.CreateHallways(hallways)
}
func (n *Node) GetLeaves() []*Rect {
	q := Queue{}
	q.Push(n)

	var children []*Rect

	for len(q) > 0 {
		curr, err := q.Pop()

		if err != nil {
			break
		}

		if curr.isLeaf() {
			curr.CreateRoom()
			children = append(children, curr.Room)
		} else {
			if curr.Left != nil {
				q.Push(curr.Left)
			}
			if curr.Right != nil {
				q.Push(curr.Right)
			}
		}
	}

	return children
}
