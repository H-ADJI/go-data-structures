package main

import (
	"fmt"

	linkedList "github.com/H-ADJI/go-data-structures/linked_list"
)

const Hello = "lol"

func main() {
	array := []int{2, 3, 5, 3, 2, 1, 6, 9, 4, 232, 43, 2, 0}
	list := linkedList.New()
	list.FromArray(array)
	fmt.Println(list)
	iterator := list.Iterator()
	if iterator.HasNext() {
		preceding := iterator.GetNext()
		current := iterator.GetNext()
		following := iterator.GetNext()
		current.Swap(preceding, following)

	}
	fmt.Println(list)
}
