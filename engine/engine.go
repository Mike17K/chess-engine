package engine

import (
	"bufio"
	"fmt"
	"os"
)

type ChessEngineOption struct {
	// Engine parameters
	HashSize int // Size of the transposition table in MB
	Threads  int // Number of CPU threads to use
	MultiPV  int // Number of principal variations to calculate

	// Search options
	Ponder   bool   // Whether to ponder (think on opponent's time)
	OwnBook  bool   // Whether to use the engine's opening book
	BookFile string // Path to the opening book file

	// Tablebases
	SyzygyPath       string // Path to Syzygy endgame tablebases
	SyzygyProbeDepth int    // Minimum remaining depth to probe tablebases

	// Evaluation parameters
	ContemptValue int // Contempt value in centipawns (for playing strength adjustment)

	// UCI options
	UCI_AnalyseMode   bool // Whether the engine is in analysis mode
	UCI_LimitStrength bool // Whether to limit engine strength
	UCI_Elo           int  // Target Elo rating when limiting strength
}

type ChessEngine struct {
	options ChessEngineOption
	// Add fields for game state, search info, etc.
	position    Position
	searchInfo  SearchInfo
	timeManager TimeManager
}

func NewChessEngine() *ChessEngine {
	return &ChessEngine{
		options: ChessEngineOption{
			HashSize:          64, // Default 64 MB hash table
			Threads:           1,  // Default single-threaded
			MultiPV:           1,  // Default single best move
			Ponder:            false,
			OwnBook:           true,
			BookFile:          "book.bin",
			SyzygyPath:        "",
			SyzygyProbeDepth:  1,
			ContemptValue:     0,
			UCI_AnalyseMode:   false,
			UCI_LimitStrength: false,
			UCI_Elo:           2800, // Default to maximum strength
		},
		// Initialize other fields here
	}
}

func (e *ChessEngine) Run() error {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if err := e.handleCommand(input); err != nil {
			return err
		}
	}
	return scanner.Err()
}

func (e *ChessEngine) handleCommand(input string) error {
	fmt.Println("Received command:", input)
	return nil
}

// Utility methods

func (e *ChessEngine) IsReady() bool {
	// TODO: Implement readiness check
	return true
}

func (e *ChessEngine) GetOptions() []string {
	return []string{
		"Hash",
		"Threads",
		"MultiPV",
		"Ponder",
		"OwnBook",
		"BookFile",
		"SyzygyPath",
		"SyzygyProbeDepth",
		"Contempt",
		"UCI_AnalyseMode",
		"UCI_LimitStrength",
		"UCI_Elo",
	}
}

func (e *ChessEngine) SetOption(name string, value interface{}) {
	switch name {
	case "Hash":
		e.options.HashSize = value.(int)
	case "Threads":
		e.options.Threads = value.(int)
	case "MultiPV":
		e.options.MultiPV = value.(int)
	case "Ponder":
		e.options.Ponder = value.(bool)
	case "OwnBook":
		e.options.OwnBook = value.(bool)
	case "BookFile":
		e.options.BookFile = value.(string)
	case "SyzygyPath":
		e.options.SyzygyPath = value.(string)
	case "SyzygyProbeDepth":
		e.options.SyzygyProbeDepth = value.(int)
	case "Contempt":
		e.options.ContemptValue = value.(int)
	case "UCI_AnalyseMode":
		e.options.UCI_AnalyseMode = value.(bool)
	case "UCI_LimitStrength":
		e.options.UCI_LimitStrength = value.(bool)
	case "UCI_Elo":
		e.options.UCI_Elo = value.(int)
	}
}

func (e *ChessEngine) GetInfo() map[string]string {
	output := map[string]string{
		"hash": fmt.Sprintf("%d", e.options.HashSize),
	}
	return output
}
