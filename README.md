# treap

Package treap implements a persistent (immutable) treap structure for
associating integer keys with arbitrary values.

The implementation is based on [Treaps: The Magical Awesome
BBT](http://pavpanchekha.com/blog/treap.html).

    import "github.com/bernerdschaefer/treap"

    store := treap.New()
    store = treap.Set(1, "one")
    store = treap.Set(2, "two")
    store = treap.Set(3, "three")

    for key := range store.WalkKeys() {
      fmt.Println(key)
      // Prints 1, 2, 3
    }

    for node := range store.Walk() {
      fmt.Println(node.Value)
      // Prints one, two, three
    }

There's also a simple wrapper around the treap, MutableTreap, which can
be used for quickly initializing a treap.

    mutable := treap.NewMutable()

    for i := 0; i < 1000; i++ {
      mutable.Set(i, i)
    }

    store := mutable.Treap()

    for key := range store.WalkKeys() {
      fmt.Println(key)
      // Prints 0, 1, 2, ..., 999
    }

