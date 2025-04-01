package domain

type CacheRepository interface {
	// Get retrieves an item from the cache by its key
	Get(key string) (string, error)
	// Set stores an item in the cache with a key and expiration time
	Set(key string, value string, expiration int) error
	// Delete removes an item from the cache by its key
	Delete(key string) error
	// Delete by prefix removes all items from the cache that start with the given prefix
	DeleteByPrefix(prefix string) error
	// Exists checks if an item exists in the cache by its key
	Exists(key string) (bool, error)
}
