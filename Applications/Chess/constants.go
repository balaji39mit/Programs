package main

//input-constants.
const (
	BoardName = "Board"
	Moves     = "Moves"
)

type Pieces string

const (
	KING   Pieces = "K"
	QUEEN         = "Q"
	ROOK          = "R"
	HORSE         = "H"
	BISHOP        = "B"
	PAWN          = "P"
)

type Color string

const (
	White Color = "W"
	Black       = "B"
)
