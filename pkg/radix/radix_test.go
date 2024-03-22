package radix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// {"CCC", "1{2AAABBBCCC", ""},
// {"EEE", "1{2AAABBBCCCEEE", ""},
// {"AAA", "1{2AAA", ""},
// {"EEE", "1{2AAABBBCCCDDDEEE", ""},
// {"BBB", "1{2BBB", ""},
// {"CCC", "1{2CCC", ""},
// {"DDD", "1{2AAABBBCCCDDD", ""},
// {"CCC", "1{2AAACCC", ""},

func TestTree(t *testing.T) {
	rootDir := &Dir{"root", "1{2", "root"}

	dirs := []*Dir{
		{"CCC", "1{2AAABBBCCC", ""},
		{"DDD", "1{2AAABBBCCCDDD", ""},
		{"AAA", "1{2AAA", ""},
		// {"BBB", "1{2BBB", ""},
		// {"CCC", "1{2CCC", ""},
		{"BBB", "1{2AAABBB", ""},
		// {"CCC", "1{2AAACCC", ""},

		// {"EEE", "1{2AAABBBCCCEEE", ""},
		// {"EEE", "1{2AAABBBCCCDDDEEE", ""},
	}

	tree := MakeTree(rootDir, dirs)
	tree.Dump()
	assert.NotNil(t, tree)
}
