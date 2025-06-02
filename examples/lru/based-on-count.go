/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"cachito/lru"
	"log"
)

// Example01__CacheMetadataT represents TODO
type Example01__CacheMetadataT struct {
	MaxItems     int
	CurrentCount int
}

func main() {
	/*
		##################################################################
		Example 1:
		LRU with eviction based on the amount of entries
		##################################################################
	*/

	// 1. Create your LRU injecting the information needed for later usage
	customCacheMetadata := Example01__CacheMetadataT{
		MaxItems:     3,
		CurrentCount: 0,
	}
	cache := lru.New(customCacheMetadata)

	// 2. Define what to do with that information to perform evictions when needed
	cache.ShouldEvict(func(metadata *Example01__CacheMetadataT, entry lru.Entry) bool {
		log.Printf("Items count currently stored: %v", metadata.CurrentCount)
		return metadata.CurrentCount > metadata.MaxItems
	})

	cache.OnInsert(func(metadata *Example01__CacheMetadataT, entry lru.Entry) error {
		metadata.CurrentCount++
		return nil
	})

	cache.OnDelete(func(metadata *Example01__CacheMetadataT, entry lru.Entry) error {
		metadata.CurrentCount--
		return nil
	})

	// 3. Use your LRU
	cache.CreateElement("a", 1)
	cache.CreateElement("b", 2)
	cache.CreateElement("c", 3)
	cache.CreateElement("d", 4)  // Eviction will happen here
	cache.CreateElement("e", 5)  // Eviction will happen here
	cache.CreateElement("f", 6)  // Eviction will happen here
	cache.CreateElement("g", 7)  // Eviction will happen here
	cache.CreateElement("h", 8)  // Eviction will happen here
	cache.CreateElement("i", 9)  // Eviction will happen here
	cache.CreateElement("j", 10) // Eviction will happen here

}
