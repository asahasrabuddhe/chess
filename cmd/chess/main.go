package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/asahasrabuddhe/chess"
)

func main() {
	log.SetFlags(0)

	if len(os.Args) != 2 {
		log.Println("Usage: chess <piece type, position>")
		log.Println("Example: chess 'Pawn, E2'")
		log.Println("Valid piece types: Pawn, King, Queen")
		log.Println("Valid positions: A1-H8")
		os.Exit(1)
	}

	var (
		piece chess.Piece
		err   error
	)
	inputs := strings.Split(os.Args[1], ", ")
	pieceType, position := strings.ToLower(inputs[0]), strings.ToUpper(inputs[1])

	switch pieceType {
	case "pawn":
		piece, err = chess.NewPawn(position)
		if err != nil {
			log.Fatalln(err)
		}
	case "king":
		piece, err = chess.NewKing(position)
		if err != nil {
			log.Fatalln(err)
		}
	case "queen":
		piece, err = chess.NewQueen(position)
		if err != nil {
			log.Fatalln(err)
		}
	default:
		log.Fatalln("Invalid piece type")
	}

	moves := piece.Moves()
	fmt.Println("Possible moves:", strings.Join(moves, ", "))
}
