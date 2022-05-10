package main

import (
	"fmt"

	"binary-tree/tree"
)

func main() {
	treeInt := tree.New(NewOrderedInt())
	_ = treeInt.Add(5)
	_ = treeInt.Add(8)
	_ = treeInt.Add(4)
	_ = treeInt.Add(7)
	_ = treeInt.Add(1)

	fmt.Println("======= IN ORDER =======")
	fmt.Println(treeInt.InOrder())

	fmt.Println("======= PRE ORDER =======")
	fmt.Println(treeInt.PreOrder())

	fmt.Println("======= POST ORDER =======")
	fmt.Println(treeInt.PostOrder())

	fmt.Println("======= SEARCH 4 =======")
	_, ok := treeInt.Search(4)
	if ok {
		fmt.Println("exist")
		return
	}
	fmt.Println("not found")
}
