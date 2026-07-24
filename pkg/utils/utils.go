// Package utils contains standalone functions
package utils

import (
	cryptorand "crypto/rand"
	"encoding/binary"
	"log"
	mathrand "math/rand"
	"time"
)

// NewRNG selects rng type, crypto, seed or none
func NewRNG(rngType string, seed int64) *mathrand.Rand {
	const (
		ModeCrypto = "crypto"
		ModeSeed   = "seed"
		ModeNone   = "none"
	)

	switch rngType {
	case ModeCrypto:
		var b [8]byte

		if _, err := cryptorand.Read(b[:]); err != nil {
			log.Fatal(err)
		}

		seed := int64(binary.LittleEndian.Uint64(b[:]))
		return mathrand.New(mathrand.NewSource(seed))

	case ModeSeed:
		customSeed := seed
		return mathrand.New(mathrand.NewSource(customSeed))

	default: // ModeNone
		return mathrand.New(mathrand.NewSource(time.Now().UnixNano()))
	}
}
