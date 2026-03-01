package tree_sitter_make_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_make "github.com/tree-sitter-grammars/tree-sitter-make/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_make.Language())
	if language == nil {
		t.Errorf("Error loading make grammar")
	}
}
