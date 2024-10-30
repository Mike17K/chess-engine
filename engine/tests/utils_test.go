package engine_test

import (
	"testing"

	"github.com/Mike17K/chess-engine/engine"
)

// https://bitmap-tool.mikekaipis.com/ - bitmap tool to visualize bitboards

func TestMsb(t *testing.T) {
	tests := []struct {
		name     string
		pieces   uint64
		expected uint64
	}{
		{
			name:     "White night at a1",
			pieces:   uint64(1),
			expected: uint64(1),
		},
		{
			name:     "White night at h8",
			pieces:   uint64(9223372036854775808),
			expected: uint64(9223372036854775808),
		}, {
			name:     "White night at a8, b8",
			pieces:   uint64(216172782113783808),
			expected: uint64(144115188075855872),
		}, {
			name:     "White night at a8, b8, d4",
			pieces:   uint64(216172782248001536),
			expected: uint64(144115188075855872),
		}, {
			name:     "White night at a8, d4",
			pieces:   uint64(72057594172145664),
			expected: uint64(72057594037927936),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := engine.Msb(tt.pieces)
			if result != tt.expected {
				t.Errorf("GetPosibleMovesNight() = %d; want %d", result, tt.expected)
			}
		})
	}
}
