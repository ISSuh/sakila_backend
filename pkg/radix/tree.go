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

func MakeTree(rootDir *Dir, dirs []*Dir) (*Tree, error) {
	t := &Tree{
		root: &Node{
			depth:    0,
			dir:      *rootDir,
			children: make(map[string]*Node),
		},
		currentDepth: 0,
	}

	for _, dir := range dirs {
		if err := t.Add(dir); err != nil {
			return nil, err
		}
	}

	return t, nil
}

func (t *Tree) Add(dir *Dir) error {
	return t.root.AddChild(dir, 1)
}

func (t *Tree) Dump() {
	t.root.Dump()
}
