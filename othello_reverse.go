package main

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
