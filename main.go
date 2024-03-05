package main

import (
	"fmt"

	linkedList "github.com/H-ADJI/go-data-structures/linked_list"
)

const Hello = "lol"

func main() {
	array := []int{1, 3, 2, 4}
	list := linkedList.New()
	list.FromArray(array)
	list.Sort()
	fmt.Println(list)
}
