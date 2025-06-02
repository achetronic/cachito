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
	"strconv"
)

// Example02__CacheMetadataT represents TODO
type Example02__CacheMetadataT struct {
	MaxDiskUtilizationBytes     int
	CurrentDiskUtilizationBytes int
}

// Example02_CustomValueRepresentation represents user-defined TODO
type Example02__CustomValueRepresentation struct {
	FilePath      string
	FileSizeBytes int
}

func main() {
	/*
		##################################################################
		Example 2:
		LRU with eviction based on the size of files stored in disk
		##################################################################
	*/

	// 1. Create your LRU injecting the information needed for later usage
	customCacheMetadata := Example02__CacheMetadataT{
		MaxDiskUtilizationBytes:     30000,
		CurrentDiskUtilizationBytes: 0,
	}
	cache := lru.New(customCacheMetadata)

	// 2. Define what to do with that information to perform evictions when needed
	cache.ShouldEvict(func(metadata *Example02__CacheMetadataT, entry lru.Entry) bool {

		entryValue := entry.Value.(Example02__CustomValueRepresentation)
		futureDiskUtilizationBytes := metadata.CurrentDiskUtilizationBytes + entryValue.FileSizeBytes

		log.Printf("Current total size: %v", metadata.CurrentDiskUtilizationBytes)

		return futureDiskUtilizationBytes > metadata.MaxDiskUtilizationBytes
	})

	cache.OnInsert(func(metadata *Example02__CacheMetadataT, entry lru.Entry) error {
		entryValue := entry.Value.(Example02__CustomValueRepresentation)
		metadata.CurrentDiskUtilizationBytes += entryValue.FileSizeBytes
		return nil
	})

	cache.OnDelete(func(metadata *Example02__CacheMetadataT, entry lru.Entry) error {
		entryValue := entry.Value.(Example02__CustomValueRepresentation)
		metadata.CurrentDiskUtilizationBytes -= entryValue.FileSizeBytes
		return nil
	})

	// 3. Use your LRU
	fileEntryList := []Example02__CustomValueRepresentation{
		{FilePath: "/tmp/sample", FileSizeBytes: 10000},
		{FilePath: "/tmp/sample", FileSizeBytes: 20000},
		{FilePath: "/tmp/sample", FileSizeBytes: 5000}, // Eviction will happen here
		{FilePath: "/tmp/sample", FileSizeBytes: 300},
		{FilePath: "/tmp/sample", FileSizeBytes: 8000}, // Eviction will happen here
		{FilePath: "/tmp/sample", FileSizeBytes: 1500},
	}

	for itemIndex, item := range fileEntryList {
		cache.CreateElement("my-cdn.com/path/to/picture/"+strconv.Itoa(itemIndex), item)
	}
}
