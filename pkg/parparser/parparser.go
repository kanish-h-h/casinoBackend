// Package parparser parses json parsheet
package parparser

import (
	"encoding/json"
	"os"
)

type Parsheet struct {
	ID            string    `json:"id"`
	Matrix        Matrix    `json:"matrix"`
	Bets          []float32 `json:"bets"`
	Lines         [][]int   `json:"lines"`
	Symbols       []Symbol  `json:"symbols"`
	MinMatchCount int       `json:"minMatchCount"`
	PayDirection  string    `json:"payDirection"`
}

type Matrix struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Symbol struct {
	ID             int            `json:"id"`
	Name           string         `json:"name"`
	ReelsInstances map[string]int `json:"reelsInstance"`
	UseWildSub     bool           `json:"useWildSub"`
	Multiplier     []int          `json:"multiplier"`
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
