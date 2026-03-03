# Distributed Systems Study Notes and Labs

This repository focuses on replication and consensus fundamentals through concise notes and runnable Go labs.

## Scope

- Distributed systems foundations (failure models, coordination, and consistency)
- Replication models (leader-based, multi-leader, and leaderless)
- Consensus workflows (election, log replication, ordering, locking, and fault injection)

## Repository Layout

- `Excercises/replication-labs/`: progressive Go labs
- `Excercises/replication-exercises.md`: exercise progression and tasks
- `Replication-and-Consensus/replication-and-consensus-notes.md`: conceptual notes
- `replication-exercises.md`: top-level copy of exercise plan

## Lab Progression

- Labs 1 to 5: leader-based replication and read behavior
- Labs 6 to 7: multi-leader replication and conflict handling
- Labs 8 to 10: leaderless replication, quorum, anti-entropy, vector clocks
- Labs 11 to 16: election, replicated logs, ordering, locks, and chaos testing

## Quick Start

Prerequisites:
- Go 1.20+

Run a lab:
```bash
cd Excercises/replication-labs
go run ./lab-01-single-leader
```

Run tests:
```bash
go test ./...
```

## Notes

- The `Excercises` directory name is preserved to match current repository structure.
