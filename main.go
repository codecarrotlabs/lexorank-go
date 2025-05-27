package main

import (
	"fmt"
	"lexorank-go/lexorank"
)

func main() {
	rank, err := lexorank.NewLexoRank("0|hzzzzzzzz")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	next, _ := rank.Next()
	fmt.Printf("Current: %s\n", rank.String())
	fmt.Printf("Next: %s\n", next.String())
}