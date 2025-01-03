package main

type PlayerStore interface {
	Get(name string) (int, bool)
}

type InMemoryDb struct {
	store map[string]int
}

func (db *InMemoryDb) Get(name string) (int, bool) {
	score, ok := db.store[name]
	return score, ok
}

func (db *InMemoryDb) Add(name string) {
	db.store[name]++
}
