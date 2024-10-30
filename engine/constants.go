package engine

const (
	H_ATTACKS uint64 = 0xA1100110A // Pattern for knight moves

	// Board rank and file masks
	Rank1BB uint64 = 0x00000000000000FF
	Rank8BB uint64 = 0xFF00000000000000
	FileABB uint64 = 0x0101010101010101
	FileHBB uint64 = 0x8080808080808080
)
