// Package configs holds configuration of engine
package configs

// Config holds parameters for a Monte Carlo run.
// Generic enough for any line-based slot variant.
type Config struct {
	TotalSpins  int64   // Number of spins to simulate
	RngType     string  // type of rng we gonna use
	Seed        int64   // RNG seed for reproducibility (0 = random)
	BetStrategy string  // "min", "max", "random", or "fixed"
	FixedBet    float64 // Used if BetStrategy == "fixed"
	ReportEvery int64   // Print progress every N spins (0 = silent)
}

const (
	RNGCrypto = "crypto"
	RNGSeed   = "seed"
	RNGNone   = "none"
)

// DefaultConfig returns sane defaults for development testing.
func DefaultConfig() Config {
	return Config{
		TotalSpins:  10000,
		RngType:     "seed",
		Seed:        40,
		ReportEvery: 1000,
	}
}
