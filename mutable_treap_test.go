package treap

import "math/rand"
import "testing"

func Test_MutableNodeGet(t *testing.T) {
	mutableTreap := NewMutable()
	mutableTreap.Set(1, "a")

	treap := mutableTreap.Treap()
	if _, ok := treap.Get(1); !ok {
		t.Error("expected treap to contain key 1 but didn't")
	}
	if _, ok := treap.Get(2); ok {
		t.Error("expected treap not to contain key 2 but it did")
	}
}

func Test_MutableNodeSet(t *testing.T) {
	mutable := NewMutable()
	mutable.Set(1, "a")
	mutable.Set(1, "b")
	mutable.Set(0, "c")
	treap := mutable.Treap()

	value, ok := treap.Get(1)
	if !ok {
		t.Error("expected treap to contain key 1 but didn't")
	}
	if value != "b" {
		t.Errorf("expected treap[1] to equal 'b' but was '%s'", value)
	}

	value, ok = treap.Get(0)
	if !ok {
		t.Error("expected treap to contain key 0 but didn't")
	}
	if value != "c" {
		t.Errorf("expected treap[0] to equal 'c' but was '%s'", value)
	}
}

func Test_MutableTreapIsBalanced(t *testing.T) {
	treap := NewMutable()
	for i := 0; i < 1000; i++ {
		treap.Set(i, i)
	}
	if treap.Treap().Depth() >= 40 {
		t.Error("treap is unbalanced")
	}
}

func Test_MutableTreapIsBalancedReverse(t *testing.T) {
	treap := NewMutable()
	for i := 999; i >= 0; i-- {
		treap.Set(i, i)
	}
	if treap.Treap().Depth() >= 40 {
		t.Error("treap is unbalanced")
	}
}

func Test_MutableTreapIsBalancedRandom(t *testing.T) {
	treap := NewMutable()
	for i := 0; i < 1000; i++ {
		treap.Set(rand.Int(), nil)
	}
	if treap.Treap().Depth() >= 40 {
		t.Error("treap is unbalanced")
	}
}

func Benchmark_MutableTreapSet(b *testing.B) {
	treap := NewMutable()
	for i := 0; i < b.N; i++ {
		treap.Set(i, string(i))
	}
}
