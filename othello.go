package main

import (
	"fmt"
)

type Board struct {
	tokens []int // tokens[0] -> (0,0), tokens[1] -> (0,1), ...
}

func (b *Board) put(x, y int, u int) {
	if u == 1 || u == -1 || u == 0 {
		b.tokens[x+8*y] = u
	}
}

func (b *Board) get(x, y int) string {
	if b.tokens[x+8*y] == 1 {
		return "o"
	}
	if b.tokens[x+8*y] == -1 {
		return "●"
	}
	return "-"
}

func (b *Board) print() {
	/*
		fmt.Printf("%s%s%s\n", b.get(0, 0), b.get(1, 0), b.get(2, 0))
		fmt.Printf("%s%s%s\n", b.get(0, 1), b.get(1, 1), b.get(2, 1))
		fmt.Printf("%s%s%s\n", b.get(0, 2), b.get(1, 2), b.get(2, 2))
	*/

	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			fmt.Printf("%s", b.get(x, y))
			if x == 7 {
				fmt.Printf("\n")
			}
		}
	}
}

func (b *Board) scanput() {
	var x, y int
	var s int
	fmt.Scanf("%d %d %d", &x, &y, &s)
	b.put(x, y, s)
}

func main() {
	b := &Board{
		tokens: []int{
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 1, -1, 0, 0, 0,
			0, 0, 0, -1, 1, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.print()
	b.scanput()
	b.print()
}
