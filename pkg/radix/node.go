package radix

import "fmt"

type Dir struct {
	Name    string
	DirNode string
	Path    string
}

type Node struct {
	depth  int
	myNode string
	dir    Dir

	parent   *Node
	children map[string]*Node
}

func (n *Node) AddChild(targetDir *Dir, depth int, parentPath string) {
	// skip root npde
	if depth == 0 {
		path := n.dir.Path + "/"
		n.AddChild(targetDir, depth+1, path)
		return
	}

	// find target node
	if n.dir.DirNode == targetDir.DirNode {
		n.dir.Name = targetDir.Name
		n.dir.Path = n.parent.dir.Path + n.dir.Name
		n.updateChildrenPath()
		return
	}

	// get node name on current node from target dir node
	myDirNode := SplitMyDirNode(targetDir.DirNode, depth)
	if len(myDirNode) < 3 {
		return
	}

	// get current dir node from target dir name
	targetCurrentDirNode := SplitDirNodeByDepth(targetDir.DirNode, depth)

	// find current dir node on children
	child, exist := n.children[myDirNode]
	if !exist {
		node := &Node{
			depth:  n.depth + 1,
			myNode: myDirNode,
			dir: Dir{
				Name:    "",
				DirNode: targetCurrentDirNode,
				Path:    "",
			},
			parent:   n,
			children: map[string]*Node{},
		}

		n.children[myDirNode] = node
		child = node
	}

	path := n.dir.Path
	if len(n.dir.Path) == 0 {
		path = parentPath
	}
	path += "/"

	child.AddChild(targetDir, depth+1, path)
}

func (n *Node) updateChildrenPath() {
	for _, child := range n.children {
		child.dir.Path = n.dir.Path + "/" + child.dir.Name
	}
}

func (n *Node) Dump() {
	fmt.Printf("[%d]", n.depth)
	fmt.Printf("[%s]", n.dir.DirNode)
	for i := n.depth + 1; i < 30; i = i * DirNodeLength {
		fmt.Printf("\t")
	}
	fmt.Printf("%s\n", n.dir.Path)
	for _, child := range n.children {
		child.Dump()
	}
}
