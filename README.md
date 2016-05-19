Trie
====

Package trie implements a in-memory trie tree.

https://godoc.org/github.com/hit9/trie

Example
-------

```go
package main

import (
	"fmt"
	"github.com/hit9/trie"
)

func main() {
	tr := trie.New(".")
	tr.Put("a.b.c", "data1")
	tr.Put("a.b.c.d", "data2")
	tr.Put("a.b.c.d.e", "data3")
	tr.Put("a.b.m.n.p", "data4")
	tr.Put("a.b.m.s.t", "data5")
	fmt.Println(tr.Get("a.b.c")) // "data1"
	m := tr.Match("a.*.m.*.*")
	fmt.Println(m) // map[a.b.m.s.t:data5 a.b.m.n.p:data4]

	tr1 := trie.New(".")
	tr1.Put("a.*.c.*", "data1")
	tr1.Put("a.b.c.*", "data2")
	// Used as wildcard like patterns to match a string.
	m = tr1.Matched("a.b.c.d")
	fmt.Println(m) // map[a.*.c.*:data1 a.b.c.*:data2]
}
```



License
-------

BSD.
