# Distributed Systems Implementation Examples

This repository contains compact, runnable Go programs that model core distributed systems protocols and storage designs.

The goal is practical understanding: each example isolates one concept, keeps the control flow visible, and is small enough to reason about in a single reading session.

## Engineering Intent

- Keep examples minimal but not vague: enough structure to teach invariants.
- Prefer explicit state transitions over framework abstractions.
- Make failure/consistency tradeoffs visible in stdout behavior.
- Treat these as teaching models, not production-ready implementations.

## Quick Start

```bash
cd distributed-systems-implementation-examples
go run ./01-introduction
```

Run all compile checks:

```bash
go test ./...
```

## Example Index

| Folder | Topic | What the code demonstrates | Key invariant / takeaway |
|---|---|---|---|
| `01-introduction` | Retry + timeout | Client retries request after transient timeout | Retry logic improves availability but requires idempotency |
| `02-zookeeper` | Leader election with ephemeral sequential nodes | Lowest sequence znode becomes leader | Session liveness controls leadership |
| `03-two-phase-commit` | 2PC coordinator flow | Prepare votes then global commit/abort | Any `NO` vote forces abort; blocking risk remains |
| `04-paxos` | Single-decree Paxos | Proposer obtains majority promises then value is accepted | Safety from quorum intersection |
| `05-raft` | Log replication and commit rule | Leader commits entry after majority acknowledgements | Committed entries are durable under majority availability |
| `06-harp` | Primary-backup style replication | Primary/backup log alignment and witness view tracking | Failover must preserve latest safe log prefix |
| `07-pbft` | Byzantine quorum voting | Prepare/commit quorum checks for `n=3f+1` | Need `2f+1` votes to make progress safely |
| `08-honeybadger` | Asynchronous batch agreement sketch | Nodes propose transaction batches and merge into agreed batch | Liveness without synchrony assumptions (high-level model) |
| `09-hotstuff` | Chained QC pipeline | Consecutive blocks each carry quorum certificate | Pipelined BFT reduces view-change complexity |
| `10-coral` | Ownership transfer with metadata versioning | Region ownership changes using monotonic version | Version monotonicity prevents stale ownership writes |
| `11-dynamo` | Leaderless quorum + vector clock concept | Quorum overlap (`w+r>n`) and sibling/conflict context | Availability via quorums, conflicts via concurrent writes |
| `12-cops` | Causal dependency enforcement | Write is visible only after dependency keys are present | Causal consistency requires dependency tracking |
| `13-gfs` | Master + chunk metadata | Master tracks primary lease and chunk replicas | Metadata centralization simplifies placement/lease control |
| `14-aurora` | Decoupled storage/compute model | Storage log is replicated independently of compute node | Fast failover through shared durable log |
| `15-spanner` | Commit-wait with time uncertainty | Commit timestamp + epsilon wait before external visibility | External consistency needs bounded-clock uncertainty handling |
| `16-stellar` | Federated quorum slices | Nodes define local trust slices; intersection enables agreement | Safety depends on quorum intersection assumptions |

## Reading Guide (Per Example)

For each folder, read in this order:

1. Problem framing in the first `fmt.Println` line.
2. State variables (proposal number, term, vector clock, quorum sizes, etc.).
3. Decision condition (majority, quorum, dependency satisfaction, version check).
4. Output interpretation: what success means and where failure/limitations are implied.

## What These Examples Are Not

- Not wire-compatible with real systems.
- Not complete protocol proofs.
- Not hardened for partitions, retries, persistence, or adversarial behavior.

They are intentionally focused slices that preserve the core mental model behind each system.

## Suggested Extensions

1. Add deterministic tests for safety/liveness properties in each folder.
2. Replace in-memory state with an append-only log + recovery path.
3. Inject message loss/reordering to make edge cases observable.
4. Add metrics (`latency`, `quorum success`, `retries`, `staleness window`) per run.

