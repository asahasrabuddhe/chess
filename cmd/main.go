package main

import (
	"chess"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		log.Println("Usage: chess <piece type> <position>")
		log.Println("Example: chess Pawn E2")
		log.Println("Valid piece types: Pawn, King, Queen")
		log.Println("Valid positions: A1-H8")
		os.Exit(1)
	}

	var piece chess.Piece
	pieceType, position := strings.ToLower(os.Args[1]), strings.ToUpper(os.Args[2])

	switch pieceType {
	case "pawn":
		piece = chess.NewPawn(position)
	case "king":
		piece = chess.NewKing(position)
	case "queen":
		piece = chess.NewQueen(position)
	default:
		log.Fatalln("Invalid piece type")
	}

	moves := piece.Moves()
	slices.Sort(moves)
	fmt.Println("Possible moves:", strings.Join(moves, ", "))
}
