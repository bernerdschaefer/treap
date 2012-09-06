package treap

import "sync"

// A simple wrapper around a treap for providing synchronized immutable
// modifications to the tree, with read operations operating on a snapshot.
type DB struct {
	sync.Mutex
	treap *Node
}

func NewDB(treap *Node) *DB {
	db := new(DB)
	db.treap = treap
	return db
}

func (db *DB) Set(key int, value interface{}) {
	db.Lock()
	defer db.Unlock()

	db.treap = db.treap.Set(key, value)
}

func (db *DB) Delete(key int) {
	db.Lock()
	defer db.Unlock()

	db.treap = db.treap.Delete(key)
}

func (db *DB) Get(key int) (interface{}, bool) {
	return db.treap.Get(key)
}

func (db *DB) Keys() (<-chan int) {
	return db.treap.WalkKeys()
}

func (db *DB) Walk() (<-chan *Node) {
	return db.treap.Walk()
}

func (db *DB) Snapshot() *Node {
	return db.treap
}
