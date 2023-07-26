package main

import (
	"fmt"
)

type Board struct {
	tokens []int // tokens[0] -> (0,0), tokens[1] -> (0,1), ...
	size   int
	player int
}

func (b *Board) put(x, y int, u int) {
	if u == 1 || u == -1 || u == 0 || u == 2 || u == -2 {
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
	if b.tokens[x+8*y] == 2 {
		// ○がおける場所
		return "△"
	}
	if b.tokes[x+8*y] == -2 {
		// ●がおける場所
		return "▲"
	}
	return "-"
}

func (b *Board) erasetriangle() {
	//▲と△を消す関数
	for i := 0; i < len(b.tokes); i++ {
		if b.tokens[i] == 2 || b.tokens[i] == -2 {
			b.tokens[i] = 0
		}
	}
}

func (b *Board) prediction(u int) {
	ret := false
	if b.size == 8 {
		for y := 0; y < 8; y++ {
			for x := 0; x < 8; x++ {
				if b.reverse(x, y, u, true) == true {
					if u == 1 {
						b.put(x, y, 2)
					} else {
						b.put(x, y, -2)
					}
				}
			}
		}
	} else if b.size == 6 {
		for y := 1; y < 7; y++ {
			for x := 1; x < 7; x++ {
				if b.reverse(x, y, u, true) == true {
					if u == 1 {
						b.put(x, y, 2)
					} else {
						b.put(x, y, -2)
					}
				}
			}
		}
	} else if b.size == 4 {
		for y := 2; y < 6; y++ {
			for x := 2; x < 6; x++ {
				if b.reverse(x, y, u, true) == true {
					if u == 1 {
						b.put(x, y, 2)
					} else {
						b.put(x, y, -2)
					}
				}
			}
		}
	}
}

func (b *Board) print() {
	if b.size == 8 {
		for y := 0; y < 8; y++ {
			for x := 0; x < 8; x++ {
				fmt.Printf("%s", b.get(x, y))
			}
			fmt.Printf("\n")
		}
	} else if b.size == 6 {
		for y := 1; y < 7; y++ {
			for x := 1; x < 7; x++ {
				fmt.Printf("%s", b.get(x, y))
			}
			fmt.Printf("\n")
		}
	} else if b.size == 4 {
		for y := 2; y < 6; y++ {
			for x := 2; x < 6; x++ {
				fmt.Printf("%s", b.get(x, y))
			}
			fmt.Printf("\n")
		}
	} else {
		fmt.Printf("Error size\n")
	}
}

func (b *Board) scanput() {
	var x, y int
	var s int

	if b.player == -1 {
		fmt.print("The turn of player ●\n")
	} else {
		fmt.print("The turn of player ○\n")
	}

	fmt.Scanf("%d %d %d", &x, &y, &b.player)

}

func (b *Board) switchPlayer() {
	b.player *= -1
}

/*ますが全て埋まっているかどうかの確認*/
func (b *Board) checkfull() bool {
	var check int
	check = 1

	if b.size == 8 {
		for y := 0; y < 8; y++ {
			for x := 0; x < 8; x++ {
				check *= b.tokens[x+8*y]
			}
		}
	} else if b.size == 6 {
		for y := 1; y < 7; y++ {
			for x := 1; x < 7; x++ {
				check *= b.tokens[x+8*y]
			}
		}
	} else if b.size == 4 {
		for y := 2; y < 6; y++ {
			for x := 2; x < 6; x++ {
				check *= b.tokens[x+8*y]
			}
		}
	} else {
		fmt.Printf("Error size\n")
	}

	return check != 0
}

func (b *Board) inbound(x, y int) bool {
	if x >= 0 && x < 8 && y >= 0 && y < 8 {
		return true
	} else {
		return false
	}
}

func (b *Board) checkPlaceable(u int) bool {
	ret := false
	if b.size == 8 {
		for y := 0; y < 8; y++ {
			for x := 0; x < 8; x++ {
				ret |= b.reverse(x, y, u, true)
			}
		}
	} else if b.size == 6 {
		for y := 1; y < 7; y++ {
			for x := 1; x < 7; x++ {
				ret |= b.reverse(x, y, u, true)
			}
		}
	} else if b.size == 4 {
		for y := 2; y < 6; y++ {
			for x := 2; x < 6; x++ {
				ret |= b.reverse(x, y, u, true)
			}
		}
	}
	return ret
}

// winner -1 or 1, not finished 0, draw 2
func (b *Board) checkWinner() int {
	if !b.checkfull() {
		return 0
	}
	count1 := 0
	count2 := 0
	for i := 0; i < len(b.tokens); i++ {
		if b.tokens[i] == 1 {
			count1++
		} else if b.tokens[i] == -1 {
			count2++
		}
	}
	if count1 > count2 {
		return 1
	} else if count1 < count2 {
		return -1
	} else if !b.checkPlaceable(-1) && !b.checkPlaceable(1) {
		return 2
	}
	return 0
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
			0, 0, 0, 0, 0, 0, 0, 0,
		},

		size:   4,
		player: 1, //初期値
	}
	for b.checkWinner() == 0 {
		b.prediction(b.player)
		b.print()
		b.scanput()
		if b.size == 4 && b.checkfull() {
			b.size = 8
		}
		b.switchPlayer()
		b.erasetriangle()
	}
}
