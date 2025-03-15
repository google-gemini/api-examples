// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package examples

import (
	"testing"
)

// TestCacheCreate verifies that the CacheCreate function successfully creates a cache
// with a document and uses it to generate content. It checks that the function returns
// a valid response without errors.
func TestCacheCreate(t *testing.T) {
	resp, err := CacheCreate()
	if err != nil {
		t.Errorf("CacheCreate returned an error: %v", err)
	}
	if resp == nil {
		t.Errorf("CacheCreate returned nil response")
	}
}

// TestCacheCreateFromName verifies that the CacheCreateFromName function successfully
// retrieves a cache by name and uses it to generate content. It ensures that the
// function returns a valid response without errors.
func TestCacheCreateFromName(t *testing.T) {
	resp, err := CacheCreateFromName()
	if err != nil {
		t.Errorf("CacheCreateFromName returned an error: %v", err)
	}
	if resp == nil {
		t.Errorf("CacheCreateFromName returned nil response")
	}
}

// TestCacheCreateFromChat verifies that the CacheCreateFromChat function successfully
// creates a cache from chat history and uses it to continue a conversation. It checks
// that the function returns a valid response without errors.
func TestCacheCreateFromChat(t *testing.T) {
	resp, err := CacheCreateFromChat()
	if err != nil {
		t.Errorf("CacheCreateFromChat returned an error: %v", err)
	}
	if resp == nil {
		t.Errorf("CacheCreateFromChat returned nil response")
	}
}

// TestCacheDelete verifies that the CacheDelete function successfully creates and
// then deletes a cache. Since this function doesn't return a value, the test only
// checks that it runs without errors.
func TestCacheDelete(t *testing.T) {
	// CacheDelete does not return anything; ensure it runs without error
	err := CacheDelete()
	if err != nil {
		t.Errorf("CacheDelete returned an error: %v", err)
	}
}

// TestCacheGet verifies that the CacheGet function successfully retrieves a cache
// by name. It checks that the function returns a valid cache object with a non-empty
// name and no errors.
func TestCacheGet(t *testing.T) {
	cache, err := CacheGet()
	if err != nil {
		t.Errorf("CacheGet returned an error: %v", err)
	}
	if cache == nil {
		t.Errorf("CacheGet returned nil cache")
	}
	if cache != nil && cache.Name == "" {
		t.Errorf("CacheGet returned cache with empty name")
	}
}

// TestCacheList verifies that the CacheList function successfully lists all available
// caches. Since this function doesn't return a value but prints the results, the test
// only checks that it runs without errors.
func TestCacheList(t *testing.T) {
	// CacheList does not return anything; ensure it runs without error
	err := CacheList()
	if err != nil {
		t.Errorf("CacheList returned an error: %v", err)
	}
}

// TestCacheUpdate verifies that the CacheUpdate function successfully updates a cache's
// time-to-live (TTL) value. It checks that the function returns a valid updated cache
// object with a non-empty name and no errors.
func TestCacheUpdate(t *testing.T) {
	cache, err := CacheUpdate()
	if err != nil {
		t.Errorf("CacheUpdate returned an error: %v", err)
	}
	if cache == nil {
		t.Errorf("CacheUpdate returned nil cache")
	}
	if cache != nil && cache.Name == "" {
		t.Errorf("CacheUpdate returned cache with empty name")
	}
}
