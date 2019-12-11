package hashtable

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := NewLRUCache(2)
	assertCorrectValue := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("test put 1,2 then get 1, cache size 2", func(t *testing.T) {
		cache.Put(1, 10)
		cache.Put(2, 20)
		got := cache.Get(1)
		want := 10
		assertCorrectValue(t, got, want)
	})

	t.Run("test putting 3 evicts 2", func(t *testing.T) {
		cache.Put(3, 3)
		got := cache.Get(2)
		want := -1
		assertCorrectValue(t, got, want)
	})

	t.Run("test putting 4 evicts 1", func(t *testing.T) {
		cache.Put(4, 4)
		got := cache.Get(1)
		want := -1
		assertCorrectValue(t, got, want)
	})

	t.Run("test 3 & 4 still present in cache", func(t *testing.T) {
		got := cache.Get(3)
		want := 3
		assertCorrectValue(t, got, want)

		got = cache.Get(4)
		want = 4
		assertCorrectValue(t, got, want)
	})

}

func TestLRUCacheLeetCode(t *testing.T) {

	assertCorrectValue := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("leetcode test", func(t *testing.T) {
		cache := NewLRUCache(2) //null
		cache.Get(2)            // -1
		cache.Put(2, 6)         // null
		cache.Get(1)            // -1
		cache.Put(1, 5)         //null
		cache.Put(1, 2)         //null
		cache.Get(1)            // 2
		cache.Get(2)            // 6
		got := cache.Get(2)
		want := 6
		assertCorrectValue(t, got, want)
	})

}
