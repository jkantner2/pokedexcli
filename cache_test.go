package main

import (
	"testing"
	"time"
	"fmt"

	"github.com/jkantner2/pokedexcli/internal/pokecache"
)

func TestCache(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key	string
		val	[]byte
}{
	{
		key: "www.example.com",
		val: []byte("testdata"),
	},
	{
		key: "https://example.com/path",
		val: []byte("evenmoredata"),
	},

}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := pokecache.NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("Expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("Expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = 5 * time.Millisecond
	cache := pokecache.NewCache(baseTime)
	cache.Add("http://example.com", []byte("somedata"))

	_, ok := cache.Get("http://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)
	
	_, ok = cache.Get("http://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
