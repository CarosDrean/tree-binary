package main

import (
	"fmt"

	"binary-tree/ordered"
	"binary-tree/tree"
)

func main() {
	orderedInt()
	orderedString()
}

func orderedInt() {
	fmt.Println("======= TREE INT =======")
	treeInt := tree.New(ordered.NewOrderedInt())
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

func orderedString() {
	fmt.Println("======= TREE STRING =======")
	treeString := tree.New(ordered.NewOrderedString())
	_ = treeString.Add("F")
	_ = treeString.Add("B")
	_ = treeString.Add("G")
	_ = treeString.Add("A")
	_ = treeString.Add("D")
	_ = treeString.Add("C")
	_ = treeString.Add("E")
	_ = treeString.Add("I")
	_ = treeString.Add("H")

	fmt.Println("======= IN ORDER =======")
	fmt.Println(treeString.InOrder())

	fmt.Println("======= PRE ORDER =======")
	fmt.Println(treeString.PreOrder())

	fmt.Println("======= POST ORDER =======")
	fmt.Println(treeString.PostOrder())

	fmt.Println("======= SEARCH I =======")
	_, ok := treeString.Search("I")
	if ok {
		fmt.Println("exist")
		return
	}
	fmt.Println("not found")
}
