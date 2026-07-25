package engine

import (
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

func (le *LineEvaluator) Evaluate(ctx *SpinContext) []Win {
	var wins []Win
	return wins
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
