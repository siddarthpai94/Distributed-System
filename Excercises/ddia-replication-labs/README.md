# DDIA Replication Labs (Go)

This repository contains progressive hands-on labs for Chapter 5 (Replication & Consensus) of "Designing Data-Intensive Applications" implemented as Go starter scaffolds.

Structure (each lab contains a minimal `main.go` starter and a short lab README):

```
ddia-replication-labs/
├── README.md
├── lab-01-single-leader/
├── lab-02-replication-log/
├── lab-03-async-vs-sync/
├── lab-04-follower-recovery/
├── lab-05-read-consistency/
├── lab-06-multi-leader/
├── lab-07-conflict-resolution/
├── lab-08-leaderless-quorum/
├── lab-09-anti-entropy/
├── lab-10-vector-clocks/
├── lab-11-leader-election/
├── lab-12-raft-log-replication/
├── lab-13-raft-full/
├── lab-14-total-order-broadcast/
├── lab-15-distributed-lock/
├── lab-16-chaos-testing/
├── shared/
│   ├── network.go
│   ├── node.go
│   ├── kvstore.go
│   └── test_harness.go
└── solutions/
```

Tech: Go (recommended `go 1.20+`).

Each lab's `main.go` is a minimal, well-commented scaffold to implement the exercise. The `shared` package provides lightweight simulation utilities (network, node, kvstore, test harness) intended for educational experimentation.

Use `go run ./lab-01-single-leader` to run a lab's starter program.

Enjoy the labs — design thoughtfully and test under faults!
