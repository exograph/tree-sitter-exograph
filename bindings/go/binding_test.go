package tree_sitter_exograph_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-exograph"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_exograph.Language())
	if language == nil {
		t.Errorf("Error loading Exograph grammar")
	}
}
