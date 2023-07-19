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

/*
func switchPlayer() {
	if currentPlayer == player1 {
		currentPlayer = player2
	} else {
		currentPlayer = player1
	}
}*/
/*ますが全て埋まっているかどうかの確認*/
func (b *Board) checkfull() bool {
	var check int
	check = 1
	for p := 18; p < 22; p++ {
		check *= b.tokens[p]
	}
	if check == 0 {
		return false
	}
	for p := 26; p < 30; p++ {
		check *= b.tokens[p]
	}
	if check == 0 {
		return false
	}
	for p := 34; p < 38; p++ {
		check *= b.tokens[p]
	}
	if check == 0 {
		return false
	}
	return true
}

func (b *Board) inbound(x, y int) bool {
	if x >= 0 && x < 8 && y >= 0 && y < 8 {
		return true
	} else {
		return false
	}
}

func (b *Board) reverse(x, y, u int, check bool) bool {
	if b.get(x, y) != "-" {
		return false
	}

	var my, enemy string
	if u == 1 {
		my = "o"
		enemy = "●"
	} else {
		my = "●"
		enemy = "o"
	}

	var ret = false

	//x+
	if b.get(x+1, y) == enemy {
		var p int
		for p = 1; b.inbound(x+p, y); p++ {
			if b.get(x+p, y) == enemy {
				continue
			} else if b.get(x+p, y) == "-" {
				break
			} else {
				break
			}
		}
		if b.inbound(x+p, y) && b.get(x+p, y) == my {
			ret = true
			if !check {
				for ; p >= 0; p-- {
					b.put(x+p, y, u)
				}
			}
		}
	}
	//x-
	if b.get(x-1, y) == enemy {
		var p int
		for p = 1; b.inbound(x-p, y); p++ {
			if b.get(x-p, y) == enemy {
				continue
			} else if b.get(x-p, y) == "-" {
				break
			} else {
				break
			}
		}
		if b.inbound(x-p, y) && b.get(x-p, y) == my {
			ret = true
			if !check {
				for ; p >= 0; p-- {
					b.put(x-p, y, u)
				}
			}
		}
	}
	//y+
	if b.get(x, y+1) == enemy {
		var p int
		for p = 1; b.inbound(x, y+p); p++ {
			if b.get(x, y+p) == enemy {
				continue
			} else if b.get(x, y+p) == "-" {
				break
			} else {
				break
			}
		}
		if b.inbound(x, y+p) && b.get(x, y+p) == my {
			ret = true
			if !check {
				for ; p >= 0; p-- {
					b.put(x, y+p, u)
				}
			}
		}
	}
	//y-
	if b.get(x, y-1) == enemy {
		var p int
		for p = 1; b.inbound(x, y-p); p++ {
			if b.get(x, y-p) == enemy {
				continue
			} else if b.get(x, y-p) == "-" {
				break
			} else {
				break
			}
		}
		if b.inbound(x, y-p) && b.get(x, y-p) == my {
			ret = true
			if !check {
				for ; p >= 0; p-- {
					b.put(x, y-p, u)
				}
			}
		}
	}

	//x+ y+
	if b.get(x+1, y+1) == enemy {
		var p int
		for p = 1; b.inbound(x+p, y+p); p++ {
			if b.get(x+p, y+p) == enemy {
				continue
			} else if b.get(x+p, y+p) == "-" {
				break
			} else {
				break
			}
		}
		if b.inbound(x+p, y+p) && b.get(x+p, y+p) == my {
			ret = true
			if !check {
				for ; p >= 0; p-- {
					b.put(x+p, y+p, u)
				}
			}
		}
	}
	//x- y-
	if b.get(x-1, y-1) == enemy {
		var p int
		for p = 1; b.inbound(x-p, y-p); p++ {
			if b.get(x-p, y-p) == enemy {
				continue
			} else if b.get(x-p, y-p) == "-" {
				break
			} else {
				break
			}
		}
		if b.inbound(x-p, y-p) && b.get(x-p, y-p) == my {
			ret = true
			if !check {
				for ; p >= 0; p-- {
					b.put(x-p, y-p, u)
				}
			}
		}
	}
	//x+ y-
	if b.get(x+1, y-1) == enemy {
		var p int
		for p = 1; b.inbound(x+p, y-p); p++ {
			if b.get(x+p, y-p) == enemy {
				continue
			} else if b.get(x+p, y-p) == "-" {
				break
			} else {
				break
			}
		}
		if b.inbound(x+p, y-p) && b.get(x+p, y-p) == my {
			ret = true
			if !check {
				for ; p >= 0; p-- {
					b.put(x+p, y-p, u)
				}
			}
		}
	}
	//x- y+
	if b.get(x-1, y+1) == enemy {
		var p int
		for p = 1; b.inbound(x-p, y+p); p++ {
			if b.get(x-p, y+p) == enemy {
				continue
			} else if b.get(x-p, y+p) == "-" {
				break
			} else {
				break
			}
		}
		if b.inbound(x-p, y+p) && b.get(x-p, y+p) == my {
			ret = true
			if !check {
				for ; p >= 0; p-- {
					b.put(x-p, y+p, u)
				}
			}
		}
	}

	return ret
}

func (b *Board) checkWinner() {

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

		/*
			tokens: []int{
				0, 0, 0, 0,
				0, 1, -1, 0,
				0, -1, 1, 0,
				0, 0, 0, 0,
			},
		*/
	}
	b.print()
	b.scanput()
	b.print()
}
