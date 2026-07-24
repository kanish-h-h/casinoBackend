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
	return &ReelSet{strips}
}
