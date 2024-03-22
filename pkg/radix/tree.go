package radix

const (
	DirNodeLength = 3
)

type Tree struct {
	root         *Node
	currentDepth int
}

func SplitMyDirNode(dirNode string, currentDepth int) string {
	dirNodeLengthOnDepth := currentDepth * DirNodeLength
	if dirNodeLengthOnDepth >= len(dirNode) {
		return ""
	}
	return dirNode[dirNodeLengthOnDepth:(dirNodeLengthOnDepth + DirNodeLength)]
}

func SplitDirNodeByDepth(dirNode string, depth int) string {
	dirNodeLengthOnDepth := (depth + 1) * DirNodeLength
	if len(dirNode) < dirNodeLengthOnDepth {
		return ""
	}
	return dirNode[0:dirNodeLengthOnDepth]
}

func MakeTree(rootDir *Dir, dirs []*Dir) *Tree {
	t := &Tree{
		root: &Node{
			depth:    0,
			dir:      *rootDir,
			children: make(map[string]*Node),
		},
		currentDepth: 0,
	}

	for _, dir := range dirs {
		t.Add(dir)
	}

	return t
}

func (t *Tree) Add(dir *Dir) {
	t.root.AddChild(dir, 0, "")
}

func (t *Tree) Dump() {
	t.root.Dump()
}
