package lru

import (
	"reflect"
	"testing"
)

type String string

func (d String) Len() int {
	return len(d)
}

func TestGet(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("key1", String("123"))
	if v, ok := lru.Get("key1"); !ok || string(v.(String)) != "123" {
		t.Fatal("Cache hit key1=123 failed!")
	}
	if _, ok := lru.Get("key2"); ok {
		t.Fatal("Cache miss key2 failed")
	}
}

func TestRemoveOldest(t *testing.T) {
	k1, k2, k3 := "1", "2", "3"
	v1, v2, v3 := "v2", "v2", "v3"
	cap := len(k1) + len(k2) + len(v1) + len(v2)
	lru := New(int64(cap), nil)
	lru.Add(k1, String(v1))
	lru.Add(k2, String(v2))
	lru.Add(k3, String(v3))

	if _, ok := lru.Get("1"); ok || lru.Len() != 2 {
		t.Fatal("Remove last element failed")
	}
}

func TestOnEvicted(t *testing.T) {
	keys := make([]string, 0)
	callback := func(key string, value Value) {
		keys = append(keys, key)
	}
	lru := New(int64(10), callback)
	lru.Add("key1", String("123456"))
	lru.Add("k2", String("v2"))
	lru.Add("k3", String("v3"))
	lru.Add("k4", String("v4"))

	expected := []string{"key1", "k2"}

	if !reflect.DeepEqual(expected, keys) {
		t.Fatal("Vall OnEvicted func failed!")
	}
}
