package linked_list_test

import (
	"testing"

	linkedList "github.com/H-ADJI/go-data-structures/linked_list"
)

func TestNew(t *testing.T) {
	list := linkedList.New()
	if list == nil {
		t.Fatalf(`Wanted non-nil pointer for linked list, got %p pointer`, list)
	}
}
func TestSearch(t *testing.T) {
	list := linkedList.New()
	if list.Search(0) {
		t.Fatalf(`List should be emnpty`)
	}
	if list == nil {
		t.Fatalf(`Wanted non-nil pointer for linked list, got %p pointer`, list)
	}
}
func TestFromArray(t *testing.T) {
	array := []int{2, 3, 5, 3, 2, 1, 6, 9, 4, 232, 43, 2, 0}
	list := linkedList.New().FromArray(array)
	for _, element := range array {
		if !list.Search(element) {
			t.Fatalf("Element %v should be in the list", element)
		}
	}
}
