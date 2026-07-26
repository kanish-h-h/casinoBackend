package engine

import (
	"fmt"

	"casinoBackend/pkg/parparser"
)

// Win represents single winning combination data (result)
type Win struct {
	PathID     string  // "Line-3", "Ways-12", "Cluster-A", etc
	LineIndex  int     // Which payline from sheet.Lines triggered
	SymbolID   int     // The symbol that formed the win
	MatchCount int     // How many consecutive matches (3, 4 or 5)
	Payout     float64 // Final payout = betPerLine(or totalbet) x multiplier
}

// SpinContext represents context of each spin (input)
type SpinContext struct {
	Grid       [][]int
	SpinNumber int64
	TotalBet   float64
	BetPerLine float64
	Multiplier int
	IsFreeSpin bool
}

// Evaluator gives me wins. (Start of polymorphism) (behaviour)
type Evaluator interface {
	Evaluate(ctx *SpinContext) []Win
	Type() string
}

// LineEvaluator implementation
type LineEvaluator struct {
	sheet *parparser.Parsheet
}

func (le *LineEvaluator) Type() string {
	return "LineEvaluator"
}

// Evaluate checks every payline in the sheet against the grid
func (le *LineEvaluator) Evaluate(ctx *SpinContext) []Win {
	var wins []Win

	// Pre-Build symbol map for O(1) lookups during evaluation
	symbolMap := make(map[int]*parparser.Symbol)
	for i := range le.sheet.Symbols {
		symbolMap[le.sheet.Symbols[i].ID] = &le.sheet.Symbols[i]
	}

	// Determine pay direction
	directions := []string{"ltr"}

	switch le.sheet.PayDirection {
	case "rtl":
		directions = []string{"rtl"}
	case "both":
		directions = []string{"ltr", "rtl"}
	}

	// Iterate through every defined payline
	for lineIdx, linePath := range le.sheet.Lines {
		if len(linePath) != len(ctx.Grid) {
			continue
		}

		// 1. Extract symbols across this specific payline
		lineSymbols := make([]int, len(ctx.Grid))
		for reel := 0; reel < len(ctx.Grid); reel++ {
			row := linePath[reel]
			lineSymbols[reel] = ctx.Grid[reel][row]
		}

		// 2. Check for win in required direction
		for _, dir := range directions {
			symbols := lineSymbols
			if dir == "rtl" {
				// Reverse symbols for RTL evaluation
				symbols = make([]int, len(lineSymbols))
				for i, s := range lineSymbols {
					symbols[len(lineSymbols)-1-i] = s
				}
			}

			win := le.checkLineWin(symbols, symbolMap, le.sheet.MinMatchCount, ctx.BetPerLine, lineIdx, dir)
			if win != nil {
				wins = append(wins, *win)
			}
		}
	}

	return wins
}

// checkLineWin evaluates a single array of symbols for a consecutive match.
func (le *LineEvaluator) checkLineWin(symbols []int, symbolMap map[int]*parparser.Symbol, minMatch int, betPerLine float64, lineIdx int, direction string) *Win {
	if len(symbols) == 0 {
		return nil
	}

	matchSymbolID := -1
	matchCount := 0

	// count consecutive match count
	for _, symID := range symbols {
		_ /*symbol*/, exists := symbolMap[symID]
		if !exists {
			break
		}

		if matchSymbolID == -1 {
			if matchCount > 0 {
				// The win for this symbol cannot start from reel 1.
				break
			}
			matchSymbolID = symID
			matchCount++
		} else if symID == matchSymbolID {
			matchCount++
		} else {
			break // First mismatch
		}
	}

	// Did we meet the minimum requirement?
	if matchCount < minMatch {
		return nil
	}

	// Get symbol configuration
	cfg, exists := symbolMap[matchSymbolID]
	if !exists || len(cfg.Multiplier) == 0 {
		return nil // Symbol has no payout
	}

	// Payout index mapping: assumes multiplier[0] is max match (Reels)
	// Example 5 reels: 5x -> idx 0, 4x -> idx 1, 3x -> idx 2
	maxReels := len(symbols)
	idx := maxReels - matchCount

	if idx < 0 || idx >= len(cfg.Multiplier) {
		return nil
	}

	multiplier := cfg.Multiplier[idx]
	payout := betPerLine * float64(multiplier)

	pathID := fmt.Sprintf("L%d", lineIdx+1)
	if direction == "rtl" {
		pathID += "-RTL"
	}

	return &Win{
		pathID,
		lineIdx,
		matchSymbolID,
		matchCount,
		payout,
	}
}

func NewLineEvaluator(sheet *parparser.Parsheet) *LineEvaluator {
	return &LineEvaluator{
		sheet: sheet,
	}
}

// WaysEvaluator implementation
type WaysEvaluator struct {
	sheet *parparser.Parsheet
}

func (we *WaysEvaluator) Type() string {
	return "WaysEvaluator"
}

func (we *WaysEvaluator) Evaluate(ctx *SpinContext) []Win {
	var wins []Win
	return wins
}

func NewWaysEvaluator(sheet *parparser.Parsheet) *WaysEvaluator {
	return &WaysEvaluator{
		sheet: sheet,
	}
}
