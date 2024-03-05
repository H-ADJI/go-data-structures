package main

import (
	"fmt"

	linkedList "github.com/H-ADJI/go-data-structures/linked_list"
)

const Hello = "lol"

func main() {
	array := []int{1, 2, 1}
	list := linkedList.NewLinkedList()
	list.FromArray(array)
	list.Delete(1)
	fmt.Println(list)
}
