package main

import "fmt"

func ValidateMove(piece string, startingPos string, endingPos string, chessBoard [][]string) bool {
	//basic validations for bound checks.
	if !basicValidation(startingPos, endingPos) {
		fmt.Println("Basic validation fails")
		return false
	}

	//verify the current piece is in starting position.
	if !verifySource(piece, startingPos, chessBoard) {
		fmt.Println("Source validation fails")
		return false
	}

	//verify the ending position is empty or opponent piece.
	if !verifyTarget(piece, endingPos, chessBoard) {
		fmt.Println("Target validation fails")
		return false
	}

	if !verifyPieceMove(piece, startingPos, endingPos, chessBoard) {
		fmt.Println("Piece validation fails")
		return false
	}

	return true
}

//check for the bound checks in the starting and ending position.
func basicValidation(start string, end string) bool {
	x1, y1, x2, y2 := GetIntFromString(start, end)
	if x1 < 0 || x2 < 0 || y1 < 0 || y2 < 0 || x1 > 7 || x2 > 7 || y1 > 7 || y2 > 7 {
		return false
	}
	return true
}

//Verify that given piece is present in the correct position.
func verifySource(piece string, source string, chessBoard [][]string) bool {
	x1, y1 := int(source[0]-48), int(source[1]-48)
	return chessBoard[x1][y1] == piece
}

//verify that target position is either empty or the opponent's piece.
func verifyTarget(piece string, target string, chessBoard [][]string) bool {
	x1, y1 := int(target[0]-48), int(target[1]-48)
	if chessBoard[x1][y1] == "--" {
		return true
	}
	color := Color(string(piece[0]))
	return color != Color(chessBoard[x1][y1][0])
}

var PathValidator = func(start string, end string, rowIncrement int, colIncrement int, color Color, board [][]string) bool {
	x1, y1, x2, y2 := GetIntFromString(start, end)
	//the error comes here.
	fmt.Println(x1, y1, x2, y2, rowIncrement, colIncrement)
	for i, j := x1, y1; i != x2 && j != y2; i, j = i+rowIncrement, j+colIncrement {
		fmt.Println(board[i][j], !(i == x1 && j == y1))
		if board[i][j] != "--" && !(i == x1 && j == y1) {
			fmt.Println(board[i][j])
			return false
		}
	}
	if board[x2][y2] == "--" {
		return true
	}
	if color == White {
		return Color(board[x2][y2][0]) == Black
	} else {
		return Color(board[x2][y2][0]) == White
	}
}

func verifyPieceMove(piece string, start string, end string, chessBoard [][]string) bool {
	color := Color(string(piece[0]))
	coin := Pieces(string(piece[1]))
	validator, ok := IsValidMove[coin]
	if !ok {
		return true
	}
	var row, col int
	if ok, row, col = validator(start, end, color); !ok {
		return false
	}
	switch coin {
	case QUEEN:
		return PathValidator(start, end, row, col, color, chessBoard)
	case PAWN:
		return PathValidator(start, end, row, col, color, chessBoard)
	case HORSE: // horse does not need to worry about intermediate paths.
		return true
	}
	return true
}
