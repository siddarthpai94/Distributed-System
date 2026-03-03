# Lab 01 — Single-Leader Key-Value Store

Implement a 3-node KV store with a designated leader. Writes go to leader only; reads can be served by any node.

Steps (starter guidance):
- Build `PUT` and `GET` endpoints (HTTP or simple RPC)
- Make leader forward writes to followers
- Followers should reject direct writes (redirect)
- Use in-memory map for state; persist later if desired

Run: `go run .` (from this directory)
