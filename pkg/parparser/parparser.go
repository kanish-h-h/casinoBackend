// Package parparser parses json parsheet
package parparser

import (
	"encoding/json"
	"os"
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

// Load json parsheet
func Load(filePath string) (*Parsheet, error) {
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
