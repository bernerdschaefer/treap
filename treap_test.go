package treap

import "math/rand"
import "testing"

func Test_Sanity(t *testing.T) {
	treap := New()
	treap = treap.Set(5, "a")
	treap = treap.Set(7, "b")
	if val, _ := treap.Get(5); val != "a" {
		t.Error("fail!")
	}

	if val, _ := treap.Get(7); val != "b" {
		t.Error("fail!")
	}

	treap = treap.Set(2, "c")
	if val, _ := treap.Get(2); val != "c" {
		t.Error("fail")
	}

	treap = treap.Set(2, "d")
	if val, _ := treap.Get(2); val != "d" {
		t.Error("fail")
	}

	treap = treap.Delete(5)
	if treap.Contains(5) {
		t.Error("fail")
	}
}

func Test_TreapIsBalanced(t *testing.T) {
	treap := New()
	for i := 0; i < 1000; i++ {
		treap = treap.Set(i, i)
	}
	if treap.Depth() >= 40 {
		t.Error("treap is unbalanced")
	}
}

func Test_TreapIsBalancedReverse(t *testing.T) {
	treap := New()
	for i := 999; i >= 0; i-- {
		treap = treap.Set(i, i)
	}
	if treap.Depth() >= 40 {
		t.Error("treap is unbalanced")
	}
}

func Test_TreapIsBalancedRandom(t *testing.T) {
	treap := New()
	for i := 0; i < 1000; i++ {
		treap = treap.Set(rand.Int(), nil)
	}
	if treap.Depth() >= 40 {
		t.Error("treap is unbalanced")
	}
}

func Test_TreapWalkKeys(t *testing.T) {
	treap := New()
	for i := 0; i < 100; i++ {
		treap = treap.Set(i, i)
	}

	keys := treap.WalkKeys()
	for i := 0; i < 100; i++ {
		key := <-keys
		if key != i {
			t.Error("keys out of order")
		}
	}
}

func Test_TreapWalk(t *testing.T) {
	treap := New()
	for i := 0; i < 100; i++ {
		treap = treap.Set(i, i)
	}

	nodes := treap.Walk()
	for i := 0; i < 100; i++ {
		node := <-nodes
		if node.Value != i {
			t.Error("keys out of order")
		}
	}
}
func Benchmark_PersistentTreapSet(b *testing.B) {
	treap := New()
	for i := 0; i < b.N; i++ {
		treap = treap.Set(i, string(i))
	}
}
