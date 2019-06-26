package main

import "math"

type utilValidator func(string, string, Color) (bool, int, int)

var IsValidMove map[Pieces]utilValidator

//can choose any diagonal direction and can move any number of steps.
var isValidBishopMove = func(start string, end string, color Color) (bool, int, int) {
	x1, y1, x2, y2 := GetIntFromString(start, end)
	if x2-x1 > 0 && y2-y1 > 0 {
		return true, 1, 1
	}
	if x2-x1 < 0 && y2-y1 > 0 {
		return true, -1, 1
	}
	if x2-x1 < 0 && y2-y1 < 0 {
		return true, -1, -1
	}
	if x2-x1 > 0 && y2-y1 < 0 {
		return true, 1, -1
	}
	return false, 0, 0
}

//rook can move either horizontally or vertically.
var isValidRookMove = func(start string, end string, color Color) (bool, int, int) {
	x1, y1, x2, y2 := GetIntFromString(start, end)
	if x1 == x2 {
		if y2 > y1 {
			return true, 0, 1
		}
		if y2 < y1 {
			return true, 0, -1
		}
		return false, 0, 0
	}
	if y1 == y2 {
		if x1 < x2 {
			return true, 1, 0
		}
		if x1 > x2 {
			return true, -1, 0
		}
		return false, 0, 0
	}
	return false, 0, 0
}

var isValidQueenMove = func(start string, end string, color Color) (bool, int, int) {
	//can move like a rook or bishop.

	if ok, row, col := isValidRookMove(start, end, color); ok {
		return true, row, col
	}
	return isValidBishopMove(start, end, color)
}

var isValidHorseMove = func(start string, end string, color Color) (bool, int, int) {
	x1, y1, x2, y2 := GetIntFromString(start, end)
	x, y := int(math.Abs(float64(x2-x1))), int(math.Abs(float64(y2-y1)))
	if (x == 1 && y == 2) || (x == 2 && y == 1) {
		return true, 0, 0
	}
	return false, 0, 0
}

var isValidPawnMove = func(start string, end string, color Color) (bool, int, int) {
	x1, y1, x2, y2 := GetIntFromString(start, end)
	if color == White {
		if x2 > x1 {
			return true, 1, y2 - y1
		}
		return false, 0, 0
	} else {
		if x2 < x1 {
			return true, -1, y2 - y1
		}
		return false, 0, 0
	}

}

func GetIntFromString(start, end string) (int, int, int, int) {
	x1, y1 := int(start[0]-48), int(start[1]-48)
	x2, y2 := int(end[0]-48), int(end[1]-48)
	return x1, y1, x2, y2
}

func Init() {
	IsValidMove = make(map[Pieces]utilValidator, 10)
	IsValidMove[QUEEN] = isValidQueenMove
	IsValidMove[PAWN] = isValidPawnMove
	IsValidMove[HORSE] = isValidHorseMove
}
