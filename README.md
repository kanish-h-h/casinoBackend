# Casino Backend (Go)

A modular slot game backend written in Go, designed for deterministic simulations, game mathematics, and future analytics/ML integration.

## Current Architecture

```
casinoBackend/
├── pkg/
│   ├── configs/      # Engine and simulation configuration
│   ├── engine/       # Core slot engine logic
│   ├── parparser/    # PAR sheet parser
│   └── utils/        # Shared utilities (RNG, helpers)
├── crz.json          # PAR sheet
├── go.mod
└── main.go
```

## Current Pipeline

```
Load PAR Sheet
        │
        ▼
Load Configuration
        │
        ▼
Initialize RNG
        │
        ▼
Build Virtual Reel Strips
        │
        ▼
Ready for Simulation
```

## Features Implemented

- JSON PAR sheet loader
- Configurable simulation settings
- Multiple RNG modes
  - `crypto`
  - `seed`
  - `none`

- Deterministic RNG for reproducible simulations
- Virtual reel strip generation from PAR sheet
- Fisher-Yates shuffle using injected RNG
- Modular package structure

## RNG Modes

| Mode     | Description                              |
| -------- | ---------------------------------------- |
| `crypto` | Seed generated using `crypto/rand`       |
| `seed`   | Fixed seed for deterministic simulations |
| `none`   | Seed from current system time            |

## Current Flow

```go
Load PAR Sheet
      ↓
Configure Engine
      ↓
Create RNG
      ↓
Build Reel Set
```

## Planned Roadmap

### Phase 1 — Core Engine

- [x] PAR sheet parser
- [x] Configuration system
- [x] RNG abstraction
- [x] Reel strip generation
- [ ] Spin generation
- [ ] Visible screen generation
- [ ] Payline evaluation
- [ ] Win calculation

### Phase 2 — Simulation

- [ ] Monte Carlo simulation
- [ ] RTP calculation
- [ ] Hit frequency
- [ ] Volatility metrics
- [ ] Symbol statistics

### Phase 3 — Features

- [ ] Wild symbols
- [ ] Scatter symbols
- [ ] Bonus games
- [ ] Free spins
- [ ] Multipliers

### Phase 4 — Analytics

- [ ] Simulation reports
- [ ] CSV export
- [ ] JSON reports
- [ ] Visualization support

### Phase 5 — AI/ML

- [ ] Spin data generation
- [ ] Dataset exporter
- [ ] Python integration
- [ ] Prediction experiments

## Design Principles

- Modular package structure
- Deterministic simulations
- Dependency injection where appropriate
- Separation of parsing, configuration, engine, and utilities
- Reproducible game mathematics

## Status

Current development milestone:

```
PAR Sheet
      │
      ▼
Virtual Reel Set
      │
      ▼
Next: Spin Generation
```
