### Usage Example:
<pre>
<code>
package main

import (
	"fmt"

	"github.com/freiny/go-datastructures/tree/bintree"
)

func main() {

	// Tree Format: parent.leftChild.rightChild
	root := bintree.BuildTreeInt(`
			8.5.4
			5.9.7 4..11
			9.. 7.1.12 11.3.
			1.. 12.2. 3..
			2..
		`)

	root.LevelOrder(printAll, func() { fmt.Println() })
	fmt.Print("\n\n")
	// OUTPUT:
	// 8.5.4
	// 5.9.7 4..11
	// 9.. 7.1.12 11.3.
	// 1.. 12.2. 3..
	// 2..

	root.LevelOrder(print)
	fmt.Println()
	// OUTPUT: 8 5 4 9 7 11 1 12 3 2

	root.PreOrder(print)
	fmt.Println()
	// OUTPUT: 8 5 9 7 1 12 2 4 11 3

	root.InOrder(print)
	fmt.Println()
	// OUTPUT: 9 5 1 7 2 12 8 4 3 11

	root.PostOrder(print)
	fmt.Println()
	// OUTPUT: 9 1 2 12 7 5 3 11 4 8
}

func print(t *bintree.Tree) {
	if t != nil {
		fmt.Print(t.Value, " ")
	}
}

func printAll(t *bintree.Tree) {
	fmt.Print(t.Value)
	fmt.Print(".")
	if t.Left() != nil {
		fmt.Print(t.Left().Value)
	}
	fmt.Print(".")
	if t.Right() != nil {
		fmt.Print(t.Right().Value)
	}
	fmt.Print(" ")
}
</code>
</pre>
