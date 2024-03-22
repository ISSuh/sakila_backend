package radix

import (
	"errors"
	"fmt"
)

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

func (n *Node) AddChild(targetDir *Dir, depth int) error {
	// current node is same to target node
	if n.dir.DirNode == targetDir.DirNode {
		if len(n.parent.dir.Path) == 0 {
			return errors.New("should need parent dir info")
		}

		n.dir.Name = targetDir.Name
		n.dir.Path = n.parent.dir.Path + "/" + n.dir.Name
		n.updateChildrenPath()
		return nil
	}

	// get node name on current node from target dir node
	// ex) node : 1{2AAABBB , depth : 1 => AAA
	myDirNode := SplitMyDirNode(targetDir.DirNode, depth)
	if len(myDirNode) < 3 {
		return errors.New("invalid myDirNode")
	}

	// get current dir node from target dir name
	// ex) node : 1{2AAABBB , depth : 1 => 1{2AAA
	targetCurrentDirNode := SplitDirNodeByDepth(targetDir.DirNode, depth)
	if len(targetCurrentDirNode) < 3 {
		return errors.New("invalid targetCurrentDirNode")
	}

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

	return child.AddChild(targetDir, depth+1)
}

func (n *Node) updateChildrenPath() {
	for _, child := range n.children {
		child.dir.Path = n.dir.Path + "/" + child.dir.Name
		child.updateChildrenPath()
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
