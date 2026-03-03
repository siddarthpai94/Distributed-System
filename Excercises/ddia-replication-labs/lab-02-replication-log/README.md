# Lab 02 — Write-Ahead Log Replication

Leader maintains an append-only log. Followers replicate by pulling entries and applying them in order.

Focus:
- WAL structure
- Followers track log position and pull
- `/status` endpoint to show log positions
