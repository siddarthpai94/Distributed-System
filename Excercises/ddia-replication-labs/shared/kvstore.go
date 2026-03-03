package shared
package shared

import (
    "sync"
)

// KVStore is a minimal in-memory key-value store with mutex protection.
// Keep it intentionally simple for labs; replace with persistence as needed.
type KVStore struct {
    mu    sync.RWMutex
    store map[string]string
}

// NewKVStore constructs a new KVStore.
func NewKVStore() *KVStore {
    return &KVStore{store: make(map[string]string)}
}

// Put stores a value.
func (kv *KVStore) Put(key, value string) {
    kv.mu.Lock()
    defer kv.mu.Unlock()
    kv.store[key] = value
}

// Get retrieves a value and a boolean indicating presence.
func (kv *KVStore) Get(key string) (string, bool) {
    kv.mu.RLock()
    defer kv.mu.RUnlock()
    v, ok := kv.store[key]
    return v, ok
}

// Snapshot returns a shallow copy of the current store.
func (kv *KVStore) Snapshot() map[string]string {
    kv.mu.RLock()
    defer kv.mu.RUnlock()
    copy := make(map[string]string, len(kv.store))
    for k, v := range kv.store {
        copy[k] = v
    }
    return copy
}
