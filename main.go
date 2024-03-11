package main

import "fmt"

func main() {
	hey := "i am happy 😊"
	fmt.Println(len(hey))
	fmt.Println(len([]rune(hey)))

	for i, v := range hey {
		fmt.Println(i)
		fmt.Printf("%c, %d, %X, %U\n", v, v, v, v)

	}
}
