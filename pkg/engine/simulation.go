package engine

import "math/rand"

type Engine struct {
	reels     *ReelSet
	rng       *rand.Rand
	evaluator Evaluator
}

func New() *Engine {
	return &Engine{}
}

func (e *Engine) Spin() {}

func (e *Engine) GenRTP() {}
