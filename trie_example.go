// Copyright 2016 Chao Wang <hit9@icloud.com>

// +build ignore

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
	fmt.Println(tr.Get("a.b.c"))
	fmt.Println(tr.Get("a.b.c.d"))
	fmt.Println(tr.Get("a.b.c.d.e"))
	fmt.Println(tr.Get("a.b.m.n.p"))
	fmt.Println(tr.Get("a.b.m.s.t"))
	m := tr.Match("a.*.m.*.*")
	fmt.Println(m) // map[a.b.m.s.t:data5 a.b.m.n.p:data4]
}
