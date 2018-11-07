package suffixtree

import (
	"fmt"
	"testing"
)

func TestSuffixTree(t *testing.T) {
	words := []string{"banana", "apple", "中文app"}
	tree := NewGeneralizedSuffixTree()
	for k, word := range words {
		tree.Put(word, k)
	}
	indexs := tree.Search("a", -1)

	if len(indexs) != 3 {
		t.Error("indexs len should be 3,but ", len(indexs))
	}
	fmt.Println(indexs)
	for _, index := range indexs {
		fmt.Println(words[index])
	}

	indexs = tree.Search("文", 0)

	if len(indexs) != 1 && indexs[0] != 2 {
		t.Error("indexs len should be 1 and indexs[0] must be 2,but ", len(indexs))
	}

	printnode("\t", tree.root)
}

func printnode(flag string, n *node) {
	for _, e := range n.edges {
		fmt.Printf("%s %s %v \n", flag, string(e.label), e.node.data)
		printnode(flag+"\t-", e.node)
	}
}
