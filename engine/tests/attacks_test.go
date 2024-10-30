package engine_test

import (
	"testing"

	"github.com/Mike17K/chess-engine/engine"
)

func TestNightAttacks(t *testing.T) {
	tests := []struct {
		name     string
		pieces   uint64
		expected uint64
	}{
		{
			name:     "White night at a1",
			pieces:   uint64(1),
			expected: uint64(132096),
		},
		{
			name:     "White night at h8",
			pieces:   uint64(9223372036854775808),
			expected: uint64(9077567998918656),
		}, {
			name:     "White night at d4",
			pieces:   uint64(0x0000000008000000),
			expected: uint64(22136263676928),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := engine.GetNightAttacks(tt.pieces)
			if result != tt.expected {
				t.Errorf("GetPosibleMovesNight() = %d; want %d", result, tt.expected)
			}
		})
	}
}

func TestGetPosibleMovesNight(t *testing.T) {
	tests := []struct {
		name     string
		color    engine.Color
		position engine.Position
		expected uint64
	}{
		{
			name:  "White night at a1",
			color: engine.White,
			position: engine.Position{
				Wn: 1,
			},
			expected: 132096,
		},
		{
			name:  "White night at h8",
			color: engine.White,
			position: engine.Position{
				Wn: 9223372036854775808,
			},
			expected: 9077567998918656,
		},
		{
			name:  "White night at d4",
			color: engine.White,
			position: engine.Position{
				Wn: 0x0000000008000000,
			},
			expected: 22136263676928,
		}, {
			name:  "Black night at b1, g1",
			color: engine.Black,
			position: engine.Position{
				Bn: 66,
			},
			expected: 10819584,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := engine.GetPosibleMovesNight(tt.color, tt.position)
			if result != tt.expected {
				t.Errorf("GetPosibleMovesNight() = %d; want %d", result, tt.expected)
			}
		})
	}
}

func TestGetPosibleMovesRook(t *testing.T) {
	tests := []struct {
		name     string
		color    engine.Color
		position engine.Position
		expected uint64
	}{
		{
			name:  "White rook at a1",
			color: engine.White,
			position: engine.Position{
				Wr: 1,
			},
			expected: 72340172838076926,
		},
		{
			name:  "White rook at h8",
			color: engine.White,
			position: engine.Position{
				Wr: 9223372036854775808,
			},
			expected: 9187484529235886208,
		},
		{
			name:  "White rook at d4",
			color: engine.White,
			position: engine.Position{
				Wr: 134217728,
			},
			expected: 578721386714368008,
		}, {
			name:  "Black rook at b1, g1",
			color: engine.Black,
			position: engine.Position{
				Br: 65,
			},
			expected: 4702111234474983935,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := engine.GetPosibleMovesRook(tt.color, tt.position)
			if result != tt.expected {
				t.Errorf("GetPosibleMovesRook() = %d; want %d", result, tt.expected)
			}
		})
	}
}
