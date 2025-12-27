package utils

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

func (n *Node) Split(minSize, maxLevel int) {}
