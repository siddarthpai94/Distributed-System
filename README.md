# Replication & Consensus — Hands-On Exercises

A progressive lab series for implementing the core concepts  Recommended language: **Python or Go** (networking primitives are clean in both).

---

## Repo Structure

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
│   ├── network.py          # Simulated unreliable network
│   ├── node.py             # Base node class
│   ├── kvstore.py          # Simple key-value store
│   └── test_harness.py     # Test framework with fault injection
└── solutions/
```

---

## Part 1: Leader-Based Replication (Labs 1–5)

### Lab 01 — Single-Leader Key-Value Store

**Concept:** Leader-follower architecture basics

**Task:** Build a 3-node KV store where one node is the designated leader. Writes go only to the leader; reads can go to any node.

**Requirements:**
- Implement `PUT(key, value)` and `GET(key)` over HTTP or TCP
- Leader accepts writes and forwards them to followers
- Followers reject write requests with a redirect to the leader
- All nodes serve reads from local state

**Test:** Write to leader, read from follower after a short delay. Verify eventual consistency.

**Starter code:** Provided node skeleton with networking.

---

### Lab 02 — Write-Ahead Log Replication

**Concept:** Replication log / change stream

**Task:** Add a sequential write-ahead log (WAL) to the leader. Followers replicate by consuming this log.

**Requirements:**
- Leader maintains an append-only log: `[(seq=1, PUT, k1, v1), (seq=2, PUT, k2, v2), ...]`
- Followers track their current log position and pull new entries
- Followers apply log entries in order to rebuild state
- Expose `/status` endpoint showing each node's current log position

**Test:** Write 1000 keys. Verify all followers eventually have identical state and log length.

---

### Lab 03 — Synchronous vs. Asynchronous Replication

**Concept:** Sync, async, and semi-sync replication tradeoffs

**Task:** Make the replication mode configurable and measure the impact.

**Requirements:**
- `SYNC` mode: leader waits for ALL followers to ACK before responding to client
- `ASYNC` mode: leader responds immediately, followers catch up in background
- `SEMI_SYNC` mode: leader waits for 1 follower ACK, rest are async
- Measure and report write latency for each mode
- Simulate a follower going slow (add 500ms delay) — observe effect on each mode

**Test:** Benchmark 1000 writes under each mode. Kill the leader after a write in async mode — show data loss.

---

### Lab 04 — Follower Recovery & Catch-Up

**Concept:** Handling node failures and rejoining

**Task:** Implement follower crash recovery using the replication log.

**Requirements:**
- Follower persists its last-applied log position to disk
- On restart, follower resumes replication from its last position
- If follower is too far behind (log compacted), it requests a full snapshot from leader
- Leader periodically compacts the log and maintains a snapshot

**Test:**
1. Write 100 keys, kill a follower
2. Write 100 more keys
3. Restart follower — verify it catches up to the full 200 keys
4. Compact the log, kill another follower, write more — verify snapshot-based recovery

---

### Lab 05 — Read Consistency Guarantees

**Concept:** Read-after-write, monotonic reads, consistent prefix reads

**Task:** Implement three read consistency levels that clients can request.

**Requirements:**
- `read-after-write`: If a client wrote key X, their next read of X must reflect the write. (Implementation: route reads to leader for recently-written keys, or track write timestamps.)
- `monotonic-read`: A client must never see older data than they've already seen. (Implementation: pin client to a single replica, or track high-water marks.)
- `consistent-prefix`: Reads must respect causal ordering. (Implementation: tag causally related writes and enforce ordering.)
- Default `eventual` mode for comparison.

**Test:** Script a client that writes, then reads from random replicas. Count anomalies under each consistency mode. `eventual` should show anomalies; others should not.

---

## Part 2: Multi-Leader Replication (Labs 6–7)

### Lab 06 — Multi-Leader Setup

**Concept:** Multi-datacenter / multi-leader replication

**Task:** Extend the system to support two leaders, each with one follower. Leaders replicate to each other asynchronously.

**Requirements:**
- Both leaders accept writes independently
- Each leader forwards its log to the other leader (cross-replication)
- Detect and skip already-applied entries (avoid infinite replication loops) using origin tagging
- Simulate "datacenter" latency between leaders (200ms delay)

**Test:** Write different keys to each leader simultaneously. Verify all 4 nodes converge.

---

### Lab 07 — Conflict Detection & Resolution

**Concept:** Write conflicts in multi-leader systems

**Task:** Handle concurrent conflicting writes to the same key on different leaders.

**Implement four resolution strategies:**
1. **Last-Write-Wins (LWW):** Use wall-clock timestamps; highest timestamp wins. Show how this silently drops data.
2. **Highest-Replica-ID Wins:** Deterministic but arbitrary. Show data loss.
3. **Merge:** For a "shopping cart" use case, merge the values (union of sets).
4. **Application-Level Callback:** On conflict, store both values and expose a `/conflicts` endpoint for manual resolution.

**Test:** Concurrently write `PUT(x, "A")` on leader-1 and `PUT(x, "B")` on leader-2. Verify each strategy's behavior. Count lost writes for LWW.

---

## Part 3: Leaderless Replication (Labs 8–10)

### Lab 08 — Quorum Reads and Writes

**Concept:** Dynamo-style quorum consistency (w + r > n)

**Task:** Build a 5-node leaderless KV store with configurable quorum parameters.

**Requirements:**
- Client sends writes to `w` nodes and reads from `r` nodes in parallel
- Configurable `n`, `w`, `r` (default: n=5, w=3, r=3)
- On read, return the value with the highest version number
- Implement read repair: if a read finds a stale node, send it the latest value

**Test:**
- With w=3, r=3: write a value, kill 2 nodes, read — should still succeed
- With w=1, r=1: show stale reads are possible
- With w=3, r=3: kill 3 nodes — show that writes fail (no quorum)

---

### Lab 09 — Anti-Entropy & Merkle Trees

**Concept:** Background consistency repair

**Task:** Implement a background anti-entropy process that detects and repairs inconsistencies between replicas.

**Requirements:**
- Each node maintains a Merkle tree over its key-value data
- Periodically, nodes exchange Merkle tree roots with peers
- If roots differ, walk the tree to find divergent key ranges
- Sync only the differing keys (efficient bandwidth)
- Log all repairs

**Test:** Manually corrupt data on one node (change a value directly). Verify anti-entropy detects and fixes it within one cycle.

---

### Lab 10 — Version Vectors (Vector Clocks)

**Concept:** Detecting concurrent writes without a leader

**Task:** Replace simple version numbers with version vectors to properly detect concurrent vs. sequential writes.

**Requirements:**
- Each node maintains a vector clock: `{node_id: counter, ...}`
- On write, increment local counter in the vector
- On read, return all causally concurrent versions (siblings)
- Client must resolve siblings on next write (read-modify-write)
- Implement "descends from" check: version A descends from B if all A's counters ≥ B's

**Test:**
1. Client-1 writes `x=A` via node-1
2. Client-2 (not having seen x=A) writes `x=B` via node-3
3. Read x — should return BOTH `{A, B}` as siblings
4. Client-3 reads both, writes `x=C` (resolving conflict) — should have vector that dominates both

---

## Part 4: Consensus (Labs 11–15)

### Lab 11 — Leader Election

**Concept:** Distributed leader election with terms/epochs

**Task:** Implement leader election for a 3-node cluster using the Raft election mechanism.

**Requirements:**
- Nodes start as followers with randomized election timeouts
- On timeout, a node becomes a candidate: increments term, votes for itself, requests votes from peers
- A node grants a vote only if it hasn't voted in this term and candidate's log is at least as up-to-date
- Candidate receiving majority becomes leader; sends periodic heartbeats
- If a follower receives a heartbeat with a higher term, it steps down

**Test:**
1. Start 3 nodes — exactly one should become leader within a few seconds
2. Kill the leader — a new leader should be elected
3. Restart the old leader — it should rejoin as a follower (its term is stale)
4. Network-partition the leader from both followers — followers elect a new leader; old leader steps down when partition heals

---

### Lab 12 — Raft Log Replication

**Concept:** Replicated state machine via log consensus

**Task:** Extend the Raft leader election to include log replication.

**Requirements:**
- Leader receives client commands and appends them to its log with the current term
- Leader sends `AppendEntries` RPCs to followers
- Followers accept entries if their log matches the leader's previous entry (consistency check)
- Leader tracks `matchIndex` per follower; commits entry when replicated to majority
- Committed entries are applied to the KV store state machine

**Test:**
1. Write 50 keys via the leader
2. Verify all nodes' logs are identical
3. Kill a follower, write 20 more keys, restart follower — verify it catches up
4. Simulate log divergence (old leader with uncommitted entries) — verify Raft resolves it

---

### Lab 13 — Full Raft Implementation

**Concept:** Complete consensus protocol

**Task:** Combine election + log replication + safety properties into a full Raft implementation.

**Additional requirements beyond Labs 11-12:**
- Log compaction via snapshots (install snapshot RPC)
- Cluster membership changes (adding/removing nodes)
- Client request deduplication (idempotency via client IDs + sequence numbers)
- Linearizable reads (leader confirms it's still leader before responding to reads)

**Test:** Run the full Jepsen-style test suite (provided):
- Random network partitions
- Node crashes and restarts
- Concurrent client writes
- Verify: no lost committed writes, no stale reads, logs always consistent

---

### Lab 14 — Total Order Broadcast

**Concept:** Using consensus for total order broadcast / replicated log

**Task:** Build a total order broadcast primitive on top of your Raft implementation.

**Requirements:**
- `broadcast(message)` — delivers message to all nodes in the same order
- Every node sees every message, and in the same sequence
- Build two applications on top:
  1. **Replicated counter:** All nodes maintain a counter; `increment` messages are broadcast; all nodes converge to same count
  2. **Unique ID generator:** Broadcast ID claims; reject duplicates

**Test:** 5 clients concurrently broadcast 100 messages each. Verify all nodes deliver the same 500 messages in the same order.

---

### Lab 15 — Distributed Lock Service

**Concept:** Using consensus for coordination (like ZooKeeper/etcd)

**Task:** Build a simple lock service on top of your consensus layer.

**Requirements:**
- `acquire(lock_name, client_id, ttl)` — acquire a lock with timeout
- `release(lock_name, client_id)` — release a lock
- `renew(lock_name, client_id, ttl)` — extend lock lease
- Locks auto-expire after TTL (fencing against dead clients)
- Fencing tokens: each lock acquisition returns a monotonically increasing token; resources can reject requests with old tokens

**Test:**
1. Two clients try to acquire the same lock — only one succeeds
2. Lock holder crashes — lock is released after TTL
3. Old lock holder wakes up, tries to write with stale fencing token — rejected
4. Run under network partitions — verify no two clients hold the lock simultaneously

---

## Part 5: Chaos Engineering (Lab 16)

### Lab 16 — Chaos Testing Your Distributed System

**Concept:** Validating system correctness under real failure conditions

**Task:** Build a chaos testing framework and run it against your Raft-backed KV store.

**Implement these fault injectors:**
- **Network partition:** Isolate a subset of nodes from the rest
- **Message delay:** Add random latency (10ms–2s) to all messages
- **Message drop:** Randomly drop X% of messages
- **Node crash:** Kill and restart a random node
- **Clock skew:** Offset a node's logical clock
- **Slow disk:** Add latency to persistence operations

**Implement these correctness checkers:**
- **Linearizability checker:** Given a history of operations, verify it is linearizable (use a simplified Wing & Gong algorithm)
- **Convergence checker:** After chaos stops, do all nodes eventually reach the same state?
- **Liveness checker:** Are client requests eventually served (within a timeout)?

**Test:** Run 60-second chaos sessions with random fault injection. Report pass/fail for each checker.

---

## Bonus Challenges

### Bonus A — Sloppy Quorums & Hinted Handoff
Implement sloppy quorums: when designated nodes are unavailable, allow writes to temporary stand-in nodes. Implement hinted handoff to route data back when original nodes recover.

### Bonus B — CRDTs (Conflict-Free Replicated Data Types)
Implement a G-Counter (grow-only counter) and an OR-Set (observed-remove set) that can be replicated across nodes without coordination and always converge.

### Bonus C — Dynamo-Style Consistent Hashing
Implement a hash ring with virtual nodes for key distribution. Handle node joins/leaves with minimal data movement.

### Bonus D — Chain Replication
Implement chain replication as an alternative to Raft: writes go to the head, propagate down the chain, and are acknowledged by the tail. Compare throughput and latency to your Raft implementation.

### Bonus E — Byzantine Fault Tolerance (Stretch Goal)
Modify your consensus algorithm to handle one Byzantine (malicious) node in a 4-node cluster using a simplified PBFT protocol.

---

## Suggested Progression

| Week | Labs | Focus |
|------|------|-------|
| 1 | 01–03 | Leader-follower basics, replication modes |
| 2 | 04–05 | Recovery, read consistency |
| 3 | 06–07 | Multi-leader, conflict resolution |
| 4 | 08–10 | Leaderless, quorums, vector clocks |
| 5 | 11–12 | Raft election and log replication |
| 6 | 13 | Full Raft (the big one) |
| 7 | 14–15 | Building on consensus: TOB, locks |
| 8 | 16 + Bonus | Chaos testing, CRDTs, advanced topics |

---

## Tech Stack Recommendations

**For learning (simplicity first):**
- Python + Flask/FastAPI for HTTP-based nodes
- Threading or asyncio for concurrent operations
- SQLite or dict-in-memory for storage
- Docker Compose for running multi-node clusters locally

**For production-grade practice:**
- Go with gRPC for node communication
- BoltDB or BadgerDB for persistent storage
- Kubernetes for orchestration
- Prometheus for metrics

---

## Grading / Evaluation Rubric (if used for a course)

| Criteria | Points |
|----------|--------|
| Correctness under normal operation | 30 |
| Correctness under failures (node crash, network partition) | 30 |
| Performance (latency, throughput benchmarks) | 15 |
| Code quality and documentation | 15 |
| Chaos test pass rate | 10 |

---

*Exercises designed to accompany Chapter 5 of Designing Data-Intensive Applications by Martin Kleppmann.*
