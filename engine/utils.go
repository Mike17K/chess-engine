package engine

import (
	"math/bits"
)

// abs returns the absolute value of an int8.
func Abs(x int8) int8 {
	if x < 0 {
		return -x
	}
	return x
}

// lsb returns the least significant '1' of the given number (64-bit integer).
func Lsb(b uint64) uint64 {
	// returns the least significant bit set to '1'
	return (b^(b-1))>>1 + 1
}

// msb returns the most significant '1' of the given number (64-bit integer).
func Msb(b uint64) uint64 {
	// calculates log2(b) to find the position of the most significant bit
	if b == 0 {
		return 0
	}
	return 1 << (63 - bits.LeadingZeros64(b))
}

// popLSB removes the least significant '1' from the given number and returns the result.
func PopLSB(b uint64) uint64 {
	// removes the least significant bit
	return b & (b - 1)
}

// _rank returns the rank of the '1' in the bitboard.
func Rank(b uint64) uint64 {
	// Rank1BB represents the bitmask for rank 1 (0xFF) shifted left if necessary.
	const Rank1BB uint64 = 0xFF
	const FileHBB uint64 = 0x0101010101010101

	return (((b * Rank1BB) & FileHBB) >> 7) * Rank1BB
}

// _file returns the file of the '1' in the bitboard.
func File(b uint64) uint64 {
	// FileABB represents the bitmask for file A (0x0101010101010101).
	const FileABB uint64 = 0x0101010101010101
	const Rank8BB uint64 = 0xFF << 56

	return (((b * FileABB) & Rank8BB) >> 56) * FileABB
}
