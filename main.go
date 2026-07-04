package main

import (
	cryptorand "crypto/rand"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	mathrand "math/rand"
	"os"
	"time"
)

type Parsheet struct {
	ID      string    `json:"id"`
	Matrix  Matrix    `json:"matrix"`
	Bets    []float32 `json:"bets"`
	Symbols []Symbol  `json:"symbols"`
}

type Matrix struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Symbol struct {
	Name           string         `json:"name"`
	ID             int            `json:"id"`
	IsSpecialCrz   bool           `json:"isSpecialCrz"`
	Payout         int            `json:"payout"`
	SpecialType    string         `json:"specialType"`
	ReelsInstances map[string]int `json:"reelsInstance"`
}

// parParser loads json parsheet
func parParser(filePath string) (*Parsheet, error) {
	var sheet Parsheet

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &sheet); err != nil {
		return nil, err
	}
	// fmt.Println("Parsheet loaded succesfully!!")
	return &sheet, nil
}

func rngSelector() *mathrand.Rand {
	const (
		ModeCrypto = "crypto"
		ModeSeed   = "seed"
		ModeNone   = "none"
	)

	rngMode := ModeCrypto

	switch rngMode {
	case ModeCrypto:
		var b [8]byte

		if _, err := cryptorand.Read(b[:]); err != nil {
			log.Fatal(err)
		}

		seed := int64(binary.LittleEndian.Uint64(b[:]))
		return mathrand.New(mathrand.NewSource(seed))

	case ModeSeed:
		customSeed := int64(43)
		return mathrand.New(mathrand.NewSource(customSeed))

	default: // ModeNone
		return mathrand.New(mathrand.NewSource(time.Now().UnixNano()))
	}
}

func linesEvaluator(sheet *Parsheet) {
}

func spin() {}

func simulation() {}

func analytics() {}

func main() {
	// Loading Parsheet
	sheet, err := parParser("crz.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ID", sheet.ID)
	fmt.Println("MATRIX: ", sheet.Matrix)

	var symbolPool []int
	for _, sym := range sheet.Symbols {
		symbolPool = append(symbolPool, sym.ID)
	}

	rng := rngSelector()

	shuffled := make([]int, len(symbolPool))

	// 1. Copy the elements from symbolPool into shuffled
	copy(shuffled, symbolPool)

	// 2. SHUFFLE: Explicit Fisher-Yates using our RNG instance
	for i := len(shuffled) - 1; i > 0; i-- {
		// Pick a random index from 0 to i (inclusive)
		j := rng.Intn(i + 1)
		// Swap the elements
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}

	fmt.Println("Shuffled Pool:", shuffled)

	rows := sheet.Matrix.Y
	cols := sheet.Matrix.X
	grid := make([][]int, rows)

	for i := 0; i < rows; i++ {
		start := i * cols
		end := start + cols

		// If shuffled array runs out of elements prevents crash
		if end > len(shuffled) {
			break
		}
		grid[i] = shuffled[start:end]
	}
	fmt.Println("--- Generated 2D Grid Screen ---")
	for _, row := range grid {
		fmt.Println(row)
	}
}
