package main

import (
	"testing"
)

func TestPutToken01(t *testing.T) {
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
		size:   8,
		player: 1,
	}
	b.put(4, 2, 1)
	//b.print()
	if b.get(4, 2) != "o" {
		t.Errorf("Mistake")
	}
}

func TestPutToken02(t *testing.T) {
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
		size:   8,
		player: -1,
	}
	b.put(2, 3, -1)
	//b.print()
	if b.get(2, 3) != "●" {
		t.Errorf("Mistake")
	}
}

func TestReverseCheck01(t *testing.T) {
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
		size:   8,
		player: -1,
	}
	if b.reverse(2, 3, -1, true) == false {
		t.Errorf("Mistake")
	}
}

func TestReverseCheck02(t *testing.T) {
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
		size:   8,
		player: 1,
	}
	if b.reverse(2, 3, 1, true) == true {
		t.Errorf("Mistake")
	}
}

func TestReverseDo01(t *testing.T) {
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
		size:   4,
		player: -1,
	}

	b.reverse(2, 3, -1, false)
	b.print()

	if b.get(2, 3) != "●" {
		t.Errorf("Mistake")
	}
	if b.get(3, 3) != "●" {
		t.Errorf("Mistake")
	}
}

func TestReverseDo02(t *testing.T) {
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
		size:   6,
		player: 1,
	}

	b.reverse(4, 2, 1, false)
	b.print()

	if b.get(4, 2) != "o" {
		t.Errorf("Mistake")
	}
	if b.get(4, 3) != "o" {
		t.Errorf("Mistake")
	}
}

func TestCheckWinner1(t *testing.T) {
	b := &Board{
		tokens: []int{
			1, -1, 1, 1, -1, 1, -1, 1,
			1, -1, 1, 1, -1, 1, -1, 1,
			-1, 1, -1, -1, 1, -1, 1, -1,
			-1, 1, -1, 1, 1, -1, 1, -1,
			1, -1, 1, -1, -1, 1, -1, 1,
			1, -1, 1, 1, -1, 1, -1, 1,
			-1, 1, -1, -1, 1, -1, 1, -1,
			-1, 1, -1, 1, 1, -1, 1, -1},
		size:   8,
		player: -1,
	}
	if b.checkWinner() != 1 {
		t.Error("Mistake")
	}

}

func TestCheckWinner2(t *testing.T) {
	b := &Board{
		tokens: []int{
			-1, 1, -1, -1, 1, -1, 1, -1,
			-1, 1, -1, -1, 1, -1, 1, -1,
			1, -1, 1, 1, -1, 1, -1, 1,
			1, -1, 1, -1, -1, 1, -1, 1,
			-1, 1, -1, 1, 1, -1, 1, -1,
			-1, 1, -1, -1, 1, -1, 1, -1,
			1, -1, 1, 1, -1, 1, -1, 1,
			1, -1, 1, -1, -1, 1, -1, 1},
		size:   8,
		player: 1,
	}
	if b.checkWinner() != -1 {
		t.Error("Mistake")
	}

}

func TestCheckWinner3(t *testing.T) {
	b := &Board{
		tokens: []int{
			1, -1, 1, -1, 1, -1, 1, -1,
			-1, 1, -1, 1, -1, 1, -1, 1,
			1, -1, 1, -1, 1, -1, 1, -1,
			-1, 1, -1, 1, -1, 1, -1, 1,
			1, -1, 1, -1, 1, -1, 1, -1,
			-1, 1, -1, 1, -1, 1, -1, 1,
			1, -1, 1, -1, 1, -1, 1, -1,
			-1, 1, -1, 1, -1, 1, -1, 1},
		size:   8,
		player: 1,
	}
	if b.checkWinner() != 2 {
		t.Error("Mistake")
	}

}
