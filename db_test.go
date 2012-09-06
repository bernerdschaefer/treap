package treap

import "testing"

func Test_DBGetSet(t *testing.T) {
	db := NewDB(New())
	db.Set(1, "one")
	if _, ok := db.Get(1); !ok {
		t.Error("expected database to contain key 1")
	}
	db.Set(1, "two")
	if value, _ := db.Get(1); value != "two" {
		t.Error("expected database to contain value 'two'")
	}
}

func Test_DBDelete(t *testing.T) {
	db := NewDB(New())
	db.Set(1, "one")
	db.Delete(1)
	if _, ok := db.Get(1); ok {
		t.Error("expected database not to contain key 1")
	}
}

func Test_DBKeys(t *testing.T) {
	db := NewDB(New())
	db.Set(1, nil)
	db.Set(2, nil)
	db.Set(3, nil)

	keys := db.Keys()
	for i := 1; i < 4; i++ {
		if i != <-keys {
			t.Error("keys not in order")
		}
	}
}

func Test_DBWalk(t *testing.T) {
	db := NewDB(New())
	db.Set(1, 1)
	db.Set(2, 2)
	db.Set(3, 3)

	nodes := db.Walk()
	for i := 1; i < 4; i++ {
		if i != (<-nodes).Value {
			t.Error("nodes not in order")
		}
	}
}

func Test_DBSnapshot(t *testing.T) {
	db := NewDB(New())
	db.Set(1, "one")
	db.Set(2, "two")
	db.Set(3, "three")

	snap := db.Snapshot()

	db.Set(4, "four")
	db.Set(5, "five")
	db.Set(6, "six")

	for key := range snap.WalkKeys() {
		if key > 3 {
			t.Error("snapshot contains keys after snapshot")
		}
	}
}
