// Package engine contains core slot engine files
package engine

import (
	"fmt"
	"math/rand"

	"casinoBackend/pkg/parparser"
)

// ReelSet holds pre-built virtual reel strips for fast sampling during simulation
// Keeping RNG encapuslated makes simulations reporducible and thread-safe
type ReelSet struct {
	Strips [][]int
	rng    *rand.Rand
}

// BuildReelSet converts a PAR sheet's reelInstance data into virtual strips
// This is a one-time setup cost. All subsequent spins sample from these arrays
// Flow: Extract -> Expand -> Shuffle -> Ready for Sampling
func BuildReelSet(sheet *parparser.Parsheet, rng *rand.Rand) *ReelSet {
	reelCount := sheet.Matrix.X
	strips := make([][]int, reelCount)

	for r := 0; r < reelCount; r++ {
		reelKey := fmt.Sprintf("%d", r)
		var strip []int

		// EXPAND: Populate strip with symbols IDs based on their frequency
		for _, sym := range sheet.Symbols {
			if count, exists := sym.ReelsInstances[reelKey]; exists {
				for i := 0; i < count; i++ {
					strip = append(strip, sym.ID)
				}
			}
		}

		// SHUFFLE: Explicit Fisher-Yates using our RNG instance
		n := len(strip)
		for i := n - 1; i > 0; i-- {
			j := rng.Intn(i + 1)
			strip[i], strip[j] = strip[j], strip[i]
		}
		strips[r] = strip
	}
	return &ReelSet{strips, rng}
}

// GenerateWindow returns a random window of symbols for each reel.
// The 'rows' parameter comes from sheet.Matrix.Y, keeping this method generic
func (rs *ReelSet) GenerateWindow(rows int) [][]int {
	result := make([][]int, len(rs.Strips))

	for r, strip := range rs.Strips {
		// pick a random starting position on the virtual strip
		startpos := rs.rng.Intn(len(strip))

		// Extract a vertical window, wrapping around if we hit the end
		window := make([]int, rows)
		for i := 0; i < rows; i++ {
			idx := (startpos + i) % len(strip)
			window[i] = strip[idx]
		}
		result[r] = window
	}
	return result
}
