package radix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree(t *testing.T) {
	rootDir := &Dir{"root", "1{2", "root"}

	dirs := []*Dir{
		{"AAA", "1{2AAA", ""},
		{"BBB", "1{2BBB", ""},
		{"CCC", "1{2CCC", ""},
		{"CCC", "1{2AAACCC", ""},
		{"BBB", "1{2AAABBB", ""},
		{"AAA", "1{2CCCAAA", ""},
		{"CCC", "1{2AAABBBCCC", ""},
		{"EEE", "1{2AAABBBCCCEEE", ""},
		{"DDD", "1{2AAABBBCCCDDD", ""},
		{"EEE", "1{2AAABBBCCCDDDEEE", ""},
	}

	tree, err := MakeTree(rootDir, dirs)
	if err == nil {
		tree.Dump()
	}

	assert.Nil(t, err)
	assert.NotNil(t, tree)

	dir, err := tree.FindByDirNode("1{2CCC")
	assert.Nil(t, err)
	if err != nil {
		assert.Equal(t, dir.Path, "root/CCC")
	}
}
