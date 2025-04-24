package examples

import (
	"testing"
)

func TestCacheCreate(t *testing.T) {
	_, err := CacheCreate()
	if err != nil {
		t.Errorf("CacheCreate returned an error.")
	}
}

func TestCacheCreateFromName(t *testing.T) {
	_, err := CacheCreateFromName()
	if err != nil {
		t.Errorf("CacheCreateFromName returned an error.")
	}
}

func TestCacheCreateFromChat(t *testing.T) {
	_, err := CacheCreateFromChat()
	if err != nil {
		t.Errorf("CacheCreateFromChat returned an error.")
	}
}

func TestCacheDelete(t *testing.T) {
	err := CacheDelete()
	if err != nil {
		t.Errorf("CacheDelete returned an error.")
	}
}

func TestCacheGet(t *testing.T) {
	err := CacheGet()
	if err != nil {
		t.Errorf("CacheGet returned an error.")
	}
}

func TestCacheList(t *testing.T) {
	err := CacheList()
	if err != nil {
		t.Errorf("CacheList returned an error.")
	}
}

func TestCacheUpdate(t *testing.T) {
	err := CacheUpdate()
	if err != nil {
		t.Errorf("CacheUpdate returned an error.")
	}
}
