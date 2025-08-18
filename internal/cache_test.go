package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("test-data"),
		},
		{
			key: "http://example.com/path",
			val: []byte("moretestData"),
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("unexpected to find the key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expecetd to find value")
				return
			}

		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("http://example.com", []byte("testData"))
	_, ok := cache.Get("http://example.com")
	if !ok {
		t.Errorf("expected to find the key")
		return
	}
	time.Sleep(waitTime)
	_, ok = cache.Get("http://example.com")
	if ok {
		t.Errorf("expected to not find key")
	}

}
