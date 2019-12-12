package unique_binary_search_trees_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateTrees(t *testing.T) {
	trees := generateTrees(3)
	assert.Equal(t, 5, len(trees))
}
