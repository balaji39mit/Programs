package main

import (
	"fmt"
)

func getChessBoard(chessBoard [][]string) [][]string {
	fmt.Println("Please enter the initial position of the pieces in chess board.")
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Scanln(&chessBoard[i][j])
		}
	}
	return chessBoard
}

func processMove(piece, start, end string, chessBoard [][]string) [][]string {
	x1, y1, x2, y2 := GetIntFromString(start, end)
	chessBoard[x2][y2] = chessBoard[x1][y1]
	chessBoard[x1][y1] = "--"
	return chessBoard
}

func print(board [][]string) {
	fmt.Println("Displaying the board")
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print(fmt.Sprintf("%v ", board[i][j]))
		}
		fmt.Print(fmt.Sprintf("\n"))
	}
}

func getMovesAndValidate(chessBoard [][]string) {
	fmt.Println("Please enter the moves")
	var piece, start, end string
GetMove:
	fmt.Scanln(&piece, &start, &end)
	fmt.Println(piece, start, end)
	if ValidateMove(piece, start, end, chessBoard) {
		fmt.Println("Valid Move")
		chessBoard = processMove(piece, start, end, chessBoard)
		print(chessBoard)
	} else {
		fmt.Println("Invalid Move")
	}
	goto GetMove
}

func main() {
	Init()
	chessBoard := make([][]string, 8)
	for i := range chessBoard {
		chessBoard[i] = make([]string, 8)
	}
	board := [][]string{
		{"WR", "WH", "WB", "WQ", "WK", "WB", "WH", "WR"},
		{"WP", "WP", "WP", "WP", "WP", "WP", "WP", "WP"},
		{"--", "--", "--", "--", "--", "--", "--", "--"},
		{"--", "--", "--", "--", "--", "--", "--", "--"},
		{"--", "--", "--", "--", "--", "--", "--", "--"},
		{"--", "--", "--", "--", "--", "--", "--", "--"},
		{"BP", "BP", "BP", "BP", "BP", "BP", "BP", "BP"},
		{"BR", "BH", "BB", "BQ", "BK", "BB", "BH", "BR"},
	}
	var input string
	fmt.Println("Start the game.")
	fmt.Scanln(&input)
	processInput := true
	for processInput {
		switch input {
		case BoardName:
			chessBoard = getChessBoard(chessBoard)
		case Moves:
			getMovesAndValidate(board)
			processInput = false
		}
	}
}
