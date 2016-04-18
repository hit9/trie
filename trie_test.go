// Copyright 2016 Chao Wang <hit9@icloud.com>

package trie

import (
	"math/rand"
	"runtime"
	"testing"
)

// Must asserts the given value is True for testing.
func Must(t *testing.T, v bool) {
	if !v {
		_, fileName, line, _ := runtime.Caller(1)
		t.Errorf("\n unexcepted: %s:%d", fileName, line)
	}
}

const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// randKey returns a random trie key with given number of segments.
func randKey(n int) string {
	if n == 0 {
		n = 1 // force n >= 1
	}
	b := make([]byte, 2*n-1)
	for i := 0; i < n; i++ {
		j := rand.Intn(len(letters))
		b[i] = letters[j]
		if i+1 < n {
			b[i+1] = '.'
		}
	}
	return string(b)
}

func TestPut(t *testing.T) {
	tr := New(".")
	// Case simple
	tr.Put("a.b.c.d", 4)
	tr.Put("a.b.c.d", 99) // case reset
	tr.Put("a.b.c.d.e", 5)
	tr.Put("a.b.c.d.e.f", 6)
	tr.Put("a.b.c.d.e.f.g", 7)
	tr.Put("a.b.c.d.e.f.g.h", 8)
	Must(t, tr.Len() == 5)
	Must(t, tr.Get("a.b.c.d").(int) == 99)
	Must(t, tr.Get("a.b.c.d.e").(int) == 5)
	Must(t, tr.Get("a.b.c.d.e.f").(int) == 6)
	Must(t, tr.Get("a.b.c.d.e.f.g").(int) == 7)
	Must(t, tr.Get("a.b.c.d.e.f.g.h").(int) == 8)
	// Case larger number.
	n := 1024 * 5
	for i := 0; i < n; i++ {
		key := randKey(rand.Intn(128))
		tr.Put(key, i)
		Must(t, tr.Get(key).(int) == i)
	}
}

func TestGet(t *testing.T) {
	tr := New(".")
	// Case not found.
	Must(t, tr.Get("not.exist") == nil)
	// Case simple.
	tr.Put("a.b.c.d", 43)
	tr.Put("b.c.d.a", 34)
	tr.Put("m.n.o.p.q", 52)
	Must(t, tr.Get("a.b.c.d").(int) == 43)
	Must(t, tr.Get("b.c.d.a").(int) == 34)
	Must(t, tr.Get("m.n.o.p.q").(int) == 52)
	Must(t, tr.Get("a.b.c") == nil)
	// Case Has.
	Must(t, tr.Has("a.b.c.d"))
	Must(t, !tr.Has("a.b.c.d.e"))
}

func TestPop(t *testing.T) {
	tr := New(".")
	// Case not found.
	Must(t, tr.Pop("not.exist") == nil)
	Must(t, tr.Len() == 0)
	// Case simple.
	tr.Put("a.b.c.d", 4)
	tr.Put("a.b.c.d.e", 5)
	tr.Put("a.b.c.d.e.f", 6)
	Must(t, tr.Len() == 3)
	Must(t, tr.Pop("a.b.c") == nil)
	Must(t, tr.Len() == 3)
	Must(t, tr.Pop("a.b.c.d").(int) == 4)
	Must(t, tr.Len() == 2)
	Must(t, tr.Pop("a.b.c.d.e").(int) == 5)
	Must(t, tr.Len() == 1)
	Must(t, tr.Pop("a.b.c.d.e.f").(int) == 6)
	Must(t, tr.Len() == 0)
}

func TestClear(t *testing.T) {
	tr := New(".")
	// Case simple.
	tr.Put("a.b.c.d", 4)
	tr.Put("a.b.c.d.e", 5)
	tr.Put("a.b.c.d.e.f", 6)
	Must(t, tr.Len() == 3)
	tr.Clear()
	Must(t, tr.Len() == 0)
	Must(t, !tr.Has("a.b.c.d"))
}

func TestMatch(t *testing.T) {
	tr := New(".")
	// Case simple.
	tr.Put("a.b.c.d", 4)
	tr.Put("a.b.c.f", 9)
	tr.Put("a.b.c.d.e", 5)
	tr.Put("a.b.c.d.e.f", 6)
	tr.Put("m.n.o.p", 43)
	tr.Put("m.n.o.p.q", 53)
	tr.Put("m.n.o.p.q.r", 63)
	var m map[string]interface{}
	// Case x.*
	m = tr.Match("a.b.*.*")
	Must(t, len(m) == 2)
	Must(t, m["a.b.c.d"].(int) == 4)
	Must(t, m["a.b.c.f"].(int) == 9)
	// Case x
	m = tr.Match("a.b.c.d")
	Must(t, len(m) == 1)
	Must(t, m["a.b.c.d"].(int) == 4)
	// Case ""
	m = tr.Match("")
	Must(t, len(m) == 0)
	// Case *.x
	m = tr.Match("*.n.o.p")
	Must(t, len(m) == 1)
	Must(t, m["m.n.o.p"].(int) == 43)
	// Case *.*
	m = tr.Match("*.b.c.*")
	Must(t, len(m) == 2)
	Must(t, m["a.b.c.d"].(int) == 4)
	Must(t, m["a.b.c.f"].(int) == 9)
	// Case *...*
	m = tr.Match("*.*.*.*")
	Must(t, len(m) == 3)
	// Case x.*.x
	m = tr.Match("a.*.*.d")
	Must(t, len(m) == 1)
}

func TestMap(t *testing.T) {
	tr := New(".")
	// Case empty.
	Must(t, len(tr.Map()) == 0)
	// Case simple.
	tr.Put("a.b.c.d", 41)
	tr.Put("a.b.c.d.e", 51)
	tr.Put("a.b.c.d.e.f", 61)
	m := tr.Map()
	Must(t, len(m) == 3)
	Must(t, m["a.b.c.d"].(int) == 41)
	Must(t, m["a.b.c.d.e"].(int) == 51)
	Must(t, m["a.b.c.d.e.f"].(int) == 61)
}

func BenchmarkPutRandKeys(b *testing.B) {
	tr := New(".")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tr.Put(randKey(128), i)
	}
}

func BenchmarkPutPrefixedKeys(b *testing.B) {
	tr := New(".")
	m := 63
	n := 16
	prefix := randKey(m)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%n == 0 {
			prefix = randKey(m)
		}
		key := prefix + "." + randKey(63)
		tr.Put(key, i)
	}
}
func BenchmarkPutAndGetRandKeys(b *testing.B) {
	tr := New(".")
	b.SetParallelism(8)
	for i := 0; i < b.N; i++ {
		tr.Put(randKey(128), i)
		tr.Get(randKey(128))
	}
}

func BenchmarkPutAndGetPrefixedKeys(b *testing.B) {
	tr := New(".")
	m := 63
	n := 16
	prefix := randKey(m)
	for i := 0; i < b.N; i++ {
		if i%n == 0 {
			prefix = randKey(m)
		}
		key := prefix + "." + randKey(63)
		tr.Put(key, i)
		tr.Get(key)
	}
}
