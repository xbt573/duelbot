package types

import "sync"

type UserDB struct {
	mu sync.RWMutex
	kv map[int64]int
}

func (d *UserDB) Set(key int64, value int) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.kv[key] = value
}

func (d *UserDB) Get(key int64) int {
	d.mu.RLock()
	defer d.mu.RUnlock()

	v, ok := d.kv[key]
	if ok {
		return v
	}

	return 0
}

func NewUserDB() *UserDB {
	return &UserDB{
		mu: sync.RWMutex{},
		kv: make(map[int64]int),
	}
}
