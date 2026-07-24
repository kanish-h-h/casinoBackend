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
	parsheet, err := parparser.Load("crz.json")
	if err != nil {
		panic(err)
	}

	// 2. Configure the System
	cfg := configs.DefaultConfig()
	cfg.RngType = configs.RNGSeed
	cfg.Seed = 42

	rng := utils.NewRNG(cfg.RngType, cfg.Seed)

	// 3. Build Reel Set
	reels := engine.BuildReelSet(parsheet, rng)
	fmt.Printf("Reels: %v\n", reels.Strips)
}
