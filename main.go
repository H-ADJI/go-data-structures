package main

import (
	"fmt"

	linkedList "github.com/H-ADJI/go-data-structures/linked_list"
)

const Hello = "lol"

func main() {
	array := []int{1, 2}
	list := linkedList.NewLinkedList()
	list.FromArray(array)
	fmt.Println(list.Search(1))
	fmt.Println(list)
}
