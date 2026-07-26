package main

import (
	"fmt"

	"casinoBackend/pkg/configs"
	"casinoBackend/pkg/engine"
	"casinoBackend/pkg/parparser"
	"casinoBackend/pkg/utils"
)

func linesEvaluator() {
}

func spin() {}

func simulation() {}

func analytics() {}

func main() {
	// Fututre Flow: load config, load parsheet, validate, build engine, simulation, analytics, export

	// 1. validate parsheet path
	parsheet, err := parparser.Load("vik.json")
	if err != nil {
		panic(err)
	}

	// 2. Configure the System
	cfg := configs.DefaultConfig()
	cfg.RngType = configs.RNGCrypto
	cfg.Seed = 42

	rng := utils.NewRNG(cfg.RngType, cfg.Seed)

	// 3. Build Reel Set
	reels := engine.BuildReelSet(parsheet, rng)
	fmt.Printf("Reels: %v\n", reels.Strips)

	// Result
	rows := parsheet.Matrix.Y
	result := reels.GenerateWindow(rows)

	// Print the Grid
	fmt.Println("Result Matrix: ", result)
	for row := 0; row < parsheet.Matrix.Y; row++ {
		for reel := 0; reel < parsheet.Matrix.X; reel++ {
			fmt.Printf("%2d ", result[reel][row])
		}
		fmt.Println()
	}

	// Evaluator
	lineEvaluator := engine.NewLineEvaluator(parsheet)

	// create contex
	ctx := &engine.SpinContext{
		Grid:       result,
		SpinNumber: 1,
		TotalBet:   0,
		BetPerLine: 0.01,
		Multiplier: 1,
		IsFreeSpin: false,
	}

	// Evaluate
	wins := lineEvaluator.Evaluate(ctx)

	// Print winnings
	if len(wins) == 0 {
		fmt.Println("No Win")
	} else {
		for _, win := range wins {
			fmt.Printf(
				"Path=%s Symbol=%d Match=%d Payout=%.2f\n",
				win.PathID,
				win.SymbolID,
				win.MatchCount,
				win.Payout,
			)
		}
	}
}
