package unicon

import (
	"strings"
	"time"

	"github.com/spf13/cast"
)

// MemoryConfig is a simple abstraction to map[]interface{} for in process memory backed configuration
// only implements Configurable use JsonConfig to save/load if needed
type MemoryConfig struct {
	data   map[string]interface{}
	casing map[string]string
}

// NewMemoryConfig returns a new memory backed Configurable
// The most basic Configurable simply backed by a map[string]interface{}
func NewMemoryConfig() *MemoryConfig {
	cfg := &MemoryConfig{
		data:   make(map[string]interface{}),
		casing: make(map[string]string),
	}
	return cfg
}

func (mem *MemoryConfig) init() {
	mem.data = make(map[string]interface{})
	mem.casing = make(map[string]string)
}

// Reset if no arguments are provided Reset() re-creates the underlaying map
func (mem *MemoryConfig) Reset(datas ...map[string]interface{}) {
	mem.data = make(map[string]interface{})
	if len(datas) >= 1 {
		for key, value := range datas[0] {
			mem.Set(key, value)
		}
	}
}

// Get key from map
func (mem *MemoryConfig) Get(key string) interface{} {
	if mem.data == nil {
		mem.init()
	}
	return mem.data[strings.ToLower(key)]
}

// GetString casts the value as a string.  If value is nil, it returns ""
func (mem *MemoryConfig) GetString(key string) string {
	return cast.ToString(mem.Get(key))
}

// GetBool casts the value as a bool.  If value is nil, it returns false
func (mem *MemoryConfig) GetBool(key string) bool {
	return cast.ToBool(mem.Get(key))
}

// GetInt casts the value as an int.  If the value is nil, it returns 0
func (mem *MemoryConfig) GetInt(key string) int {
	return cast.ToInt(mem.Get(key))
}

// GetInt64 casts the value as an int64.  If the value is nil, it returns 0
func (mem *MemoryConfig) GetInt64(key string) int64 {
	return cast.ToInt64(mem.Get(key))
}

// GetFloat64 casts the value as a float64.  If the value is nil, it returns 0.0
func (mem *MemoryConfig) GetFloat64(key string) float64 {
	return cast.ToFloat64(mem.Get(key))
}

// GetTime casts the value as a time.Time.  If the value is nil, it returns the 0 time
func (mem *MemoryConfig) GetTime(key string) time.Time {
	return cast.ToTime(mem.Get(key))
}

// GetDuration casts the value as a time.Duration.  If the value is nil, it returns the 0 duration
func (mem *MemoryConfig) GetDuration(key string) time.Duration {
	return cast.ToDuration(mem.Get(key))
}

// All returns all keys
func (mem *MemoryConfig) All() map[string]interface{} {
	if mem.data == nil {
		mem.init()
	}
	allMap := make(map[string]interface{})
	for key, value := range mem.data {
		allMap[mem.casing[key]] = value
	}
	return allMap
}

// Set a key to value
func (mem *MemoryConfig) Set(key string, value interface{}) {
	if mem.data == nil {
		mem.init()
	}
	mem.casing[strings.ToLower(key)] = key
	mem.data[strings.ToLower(key)] = value
}

// BulkSet overwrites the overrides with items in the provided map
func (mem *MemoryConfig) BulkSet(items map[string]interface{}) {
	for k, v := range items {
		mem.Set(k, v)
	}
}
