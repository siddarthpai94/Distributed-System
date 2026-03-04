# Distributed Systems Implementation Examples

This repository contains compact, runnable Go programs that model core distributed systems protocols and storage designs.

## Quick Start

```bash
cd distributed-systems-implementation-examples
go run ./01-introduction
```

Compile checks:

```bash
go test ./...
```

## Exercise Index

| Folder | Topic | What the code demonstrates | Key invariant / takeaway |
|---|---|---|---|
| `01-introduction` | Introduction | Timeout and retry behavior | Retries increase availability but need idempotency |
| `02-spanner` | Spanner | Commit timestamp + uncertainty wait | External consistency needs bounded time uncertainty |
| `03-dynamo` | Dynamo | Leaderless quorum plus vector-clock context | `w+r>n` gives overlap; concurrent writes require reconciliation |
| `04-raft` | Raft | Leader commit on majority acknowledgements | Majority commit provides durability under failures |
| `05-zookeeper` | ZooKeeper | Ephemeral sequential node election pattern | Lowest active sequence wins leadership |
| `06-paxos` | Paxos | Majority promises and accepts for one value | Quorum intersection preserves safety |
| `07-pbft` | PBFT | Byzantine prepare/commit quorum checks | Need `2f+1` votes in `n=3f+1` systems |
| `08-stellar` | Stellar | Federated quorum slices and intersection intuition | Safety depends on overlap assumptions |
| `09-cops` | COPS | Dependency-aware visibility for causal writes | Writes become visible after dependencies are satisfied |
| `10-gfs` | GFS | Master metadata and chunk lease ownership | Centralized metadata simplifies replica coordination |
| `11-honeybadger` | HoneyBadger | Asynchronous batch aggregation sketch | Liveness under weak timing assumptions |
| `12-hotstuff` | HotStuff | Chained quorum-certificate pipeline | Pipelining improves BFT steady-state throughput |
| `13-two-phase-commit` | 2PC | Coordinator prepare/commit decision flow | Any reject vote forces abort; blocking remains a risk |
| `14-aurora` | Aurora | Decoupled storage log and compute failover model | Shared durable log simplifies recovery |
| `15-coral` | Coral | Versioned ownership transfer sketch | Monotonic versions prevent stale owners from writing |
| `16-harp` | Harp | Primary-backup/witness recovery intuition | New primary must preserve safe log prefix |

## Notes

- These are educational models, not production-grade implementations.
- Each folder is independent and runnable with `go run ./<folder>`.
