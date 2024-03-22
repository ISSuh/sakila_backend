package radix

import (
	"errors"
	"fmt"
)

type node struct {
	depth  int
	myNode string
	dir    *Dir

	parent   *node
	children map[string]*node
}

func (n *node) addChild(targetDir *Dir, depth int) error {
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
	currentDirNodeOnDepth := SplitMyDirNode(targetDir.DirNode, depth)
	if len(currentDirNodeOnDepth) < 3 {
		return errors.New("invalid myDirNode")
	}

	// get current dir node from target dir name
	// ex) node : 1{2AAABBB , depth : 1 => 1{2AAA
	targetCurrentDirNode := SplitDirNodeByDepth(targetDir.DirNode, depth)
	if len(targetCurrentDirNode) < 3 {
		return errors.New("invalid targetCurrentDirNode")
	}

	// find current dir node on children
	child, exist := n.children[currentDirNodeOnDepth]
	if !exist {
		node := &node{
			depth:  depth + 1,
			myNode: currentDirNodeOnDepth,
			dir: &Dir{
				Name:    "",
				DirNode: targetCurrentDirNode,
				Path:    "",
			},
			parent:   n,
			children: map[string]*node{},
		}

		n.children[currentDirNodeOnDepth] = node
		child = node
	}

	return child.addChild(targetDir, depth+1)
}

func (n *node) updateChildrenPath() {
	for _, child := range n.children {
		child.dir.Path = n.dir.Path + "/" + child.dir.Name
		child.updateChildrenPath()
	}
}

func (n *node) findByDirNode(dirNode string, depth int) (*Dir, error) {
	if n.dir.DirNode == dirNode {
		return n.dir, nil
	}

	currentDirNodeOnDepth := SplitMyDirNode(dirNode, depth)
	if len(currentDirNodeOnDepth) < 3 {
		return nil, errors.New("invalid myDirNode")
	}

	// find current dir node on children
	child, exist := n.children[currentDirNodeOnDepth]
	if !exist {
		return nil, errors.New("not exist dir")
	}
	return child.findByDirNode(dirNode, depth)
}

func (n *node) dump() {
	fmt.Printf("[%d]", n.depth)
	fmt.Printf("[%s]", n.dir.DirNode)
	for i := n.depth + 1; i < 30; i = i * DirNodeLength {
		fmt.Printf("\t")
	}
	fmt.Printf("%s\n", n.dir.Path)
	for _, child := range n.children {
		child.dump()
	}
}
