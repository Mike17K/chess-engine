package engine

import "math/bits"

type move struct {
	From int8
	To   int8
}

func (m *move) ToString(p Position) string {
	// Convert square numbers to algebraic notation
	files := "abcdefgh"
	ranks := "12345678"

	fromFile := files[m.From%8]
	fromRank := ranks[m.From/8]
	toFile := files[m.To%8]
	toRank := ranks[m.To/8]

	return string(fromFile) + string(fromRank) + string(toFile) + string(toRank)
}

// Helper function to get all pieces of a given color
func getAllPieces(color Color, p Position) uint64 {
	if color == White {
		return p.Wp | p.Wn | p.Wb | p.Wr | p.Wq | p.Wk
	}
	return p.Bp | p.Bn | p.Bb | p.Br | p.Bq | p.Bk
}

// GetPosibleMovesNight generates knight attack bitboard
func GetPosibleMovesNight(color Color, p Position) uint64 {
	var knights uint64
	var attackingSq uint64

	// Select knights based on color
	if color == White {
		knights = p.Wn
	} else {
		knights = p.Bn
	}

	// Process each knight
	for knights != 0 {
		from := bits.TrailingZeros64(knights)
		pos := uint64(1) << from

		// Calculate new positions
		newAttacks := GetNightAttacks(pos)

		// Add new attacks to the attacking square
		attackingSq |= newAttacks

		// Remove the processed knight
		knights &= knights - 1
	}

	return attackingSq
}

func GetNightAttacks(from uint64) uint64 {
	attacks := uint64(0)
	if from < 0x0000000000040000 {
		attacks = (from * H_ATTACKS) >> 18
	} else {
		attacks = (from >> 18) * H_ATTACKS
	}

	if (from & 0x00003C3C3C3C0000) == 0 {
		if (from & 0x0303030303030303) != 0 {
			attacks &= 0x3F3F3F3F3F3F3F3F
		} else if (from & 0xC0C0C0C0C0C0C0C0) != 0 {
			attacks &= 0xFCFCFCFCFCFCFCFC
		}
	}

	return attacks
}

// GetPosibleMovesRook generates rook attack bitboard
func GetPosibleMovesRook(color Color, p Position) uint64 {
	var rooks uint64
	var attackingSq uint64

	// Select rooks based on color
	if color == White {
		rooks = p.Wr
	} else {
		rooks = p.Br
	}

	// Sliding directions for rook (horizontal and vertical)
	directions := []int8{-8, 1, 8, -1}
	limitBoards := []uint64{Rank1BB, FileHBB, Rank8BB, FileABB}

	// Get all pieces for collision detection
	blackPieces := p.Bp | p.Bn | p.Bb | p.Br | p.Bq | p.Bk
	whitePieces := p.Wp | p.Wn | p.Wb | p.Wr | p.Wq | p.Wk
	allPieces := blackPieces | whitePieces

	// For each direction
	for dir := 0; dir < 4; dir++ {
		currentRooks := rooks

		// Process each rook
		for currentRooks != 0 {
			// Clear if limits are reached
			currentRooks &= ^limitBoards[dir]

			// Calculate new positions
			var newPos uint64
			if directions[dir] < 0 {
				newPos = currentRooks >> uint(-directions[dir])
			} else {
				newPos = currentRooks << uint(directions[dir])
			}

			// Add new attacks to the attacking squares
			attackingSq |= newPos

			// Find intersections with all pieces
			intersections := allPieces & newPos
			currentRooks = newPos ^ intersections

			// Remove processed rook
			currentRooks &= currentRooks - 1
		}
	}

	return attackingSq
}

func GetBishopAttacks(color Color, p Position) uint64 {
	var attacks uint64
	var bishops uint64

	// Select bishops based on color
	if color == White {
		bishops = p.Wb
	} else {
		bishops = p.Bb
	}

	// Diagonal directions for bishop
	directions := []int8{9, 7, -9, -7}
	// Combined edge masks for diagonal
	limitBoards := []uint64{
		Rank8BB | FileHBB, //
	}

	// Get all pieces for collision detection
	blackPieces := p.Bp | p.Bn | p.Bb | p.Br | p.Bq | p.Bk
	whitePieces := p.Wp | p.Wn | p.Wb | p.Wr | p.Wq | p.Wk
	allPieces := blackPieces | whitePieces

	// For each direction
	for dir := 0; dir < 4; dir++ {
		currentBishops := bishops

		// Process each bishop
		for currentBishops != 0 {
			// Clear if limits are reached
			currentBishops &= ^limitBoards[dir]

			// Calculate new positions
			var newPos uint64
			if directions[dir] < 0 {
				newPos = currentBishops >> uint(-directions[dir])
			} else {
				newPos = currentBishops << uint(directions[dir])
			}

			// Find intersections with all pieces
			intersections := allPieces & newPos

			// Store attacking squares before removing blocked squares
			attacks |= newPos

			// Remove intersections for next iteration
			newPos -= intersections

			// Update position for next iteration
			currentBishops = newPos
		}
	}

	return attacks
}

func GetKingAttacks(color Color, p Position) uint64 {
	var attacks uint64
	var king uint64

	// Select king based on color
	if color == White {
		king = p.Wk
	} else {
		king = p.Bk
	}

	// All eight directions for king movement
	directions := []int8{9, 7, -9, -7, -8, 1, 8, -1}
	// Combined edge masks for all directions
	limitBoards := []uint64{
		Rank8BB | FileHBB, // Upper-right limits
		Rank8BB | FileABB, // Upper-left limits
		Rank1BB | FileABB, // Lower-left limits
		Rank1BB | FileHBB, // Lower-right limits
		Rank1BB,           // Down limit
		FileHBB,           // Right limit
		Rank8BB,           // Up limit
		FileABB,           // Left limit
	}

	// For each direction
	for dir := 0; dir < 8; dir++ {
		currentKing := king

		// Clear if limits are reached
		currentKing &= ^limitBoards[dir]

		// Calculate new position
		var newPos uint64
		if directions[dir] < 0 {
			newPos = currentKing >> uint(-directions[dir])
		} else {
			newPos = currentKing << uint(directions[dir])
		}

		// Generate moves for all attacking squares
		attacks |= newPos
	}

	return attacks
}

func GetPawnAttacks(color Color, p Position) uint64 {
	var attacks uint64
	var pawns uint64

	// Select pawns based on color
	if color == White {
		pawns = p.Wp
	} else {
		pawns = p.Bp
	}

	// Get all pieces for collision detection
	blackPieces := p.Bp | p.Bn | p.Bb | p.Br | p.Bq | p.Bk
	whitePieces := p.Wp | p.Wn | p.Wb | p.Wr | p.Wq | p.Wk
	allPieces := blackPieces | whitePieces

	// Calculate attacking squares
	if color == White {
		attacks = ((pawns &^ FileHBB) << 9) | ((pawns &^ FileABB) << 7)
	} else {
		attacks = ((pawns &^ FileHBB) >> 7) | ((pawns &^ FileABB) >> 9)
	}

	// Remove blocked squares
	attacks &= allPieces

	return attacks
}

func GetQueenAttacks(color Color, p Position) uint64 {
	return GetBishopAttacks(color, p) | GetPosibleMovesRook(color, p)
}
