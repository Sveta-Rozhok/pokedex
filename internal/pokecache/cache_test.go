package pokecache

import (
	"testing"
	"time"
)

func TestCache_AddAndGet(t *testing.T) {
	cache := NewCache(10 * time.Second) // создаём кэш с заданным интервалом реапинга

	key := "test_key"
	value := []byte("test_value")

	cache.Add(key, value) // добавляем значение в кэш
	result, found := cache.Get(key)

	if !found {
		t.Errorf("expected key to be found in cache")
	}
	if string(result) != string(value) {
		t.Errorf("expected value %s, but got %s", value, result)
	}
}
