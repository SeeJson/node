package algo

import (
	"testing"
)

type String string

func (s String) len() int {
	return len(s)
}

func TestCache_Add(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("xfasfsa", String("xsds"))
	lru.Add("key1", String("1234"))
	if v, ok := lru.Get("key1"); !ok || string(v.(String)) != "1234" {
		t.Fatalf("cache hit key1=1234 failed")
	}
	if _, ok := lru.Get("key2"); ok {
		t.Fatalf("cache miss key2 failed")
	}
	lru.Add("key1", String("123456565"))
	println(lru)
}
