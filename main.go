package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Mike17K/chess-engine/engine"
)

type ChessEngineOption struct {
	Name  string
	Value interface{}
}

type ChessEngine struct {
	// Add engine-specific fields here
	options map[string]interface{}
}

func NewChessEngine() *ChessEngine {
	return &ChessEngine{}
}

func (e *ChessEngine) Run() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		e.handleCommand(input)
	}
}

func (e *ChessEngine) handleCommand(input string) {
	tokens := strings.Fields(input)
	if len(tokens) == 0 {
		// Handle empty input
		return
	}
	// Handle command
}

func main() {
	// Define command-line flags
	hashSize := flag.Int("hash", 64, "Hash table size in MB")
	nalimovPath := flag.String("nalimov-path", "", "Path to Nalimov tablebases (separate multiple paths with ';')")
	nalimovCache := flag.Int("nalimov-cache", 32, "Nalimov cache size in MB")
	ponder := flag.Bool("ponder", true, "Enable pondering")
	ownBook := flag.Bool("own-book", false, "Use engine's own opening book")
	multiPV := flag.Int("multipv", 1, "Number of principal variations (MultiPV)")

	// Custom usage function to provide more detailed help
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	e := engine.NewChessEngine()
	e.SetOption("hash", *hashSize)
	e.SetOption("ponder", *ponder)
	e.SetOption("own-book", *ownBook)
	e.SetOption("multipv", *multiPV)
	e.SetOption("nalimov-path", *nalimovPath)
	e.SetOption("nalimov-cache", *nalimovCache)
	e.Run()
}
