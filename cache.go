package in_memory_cache

type Cache interface {
	New()
}

type Data struct {
	cache map[string]interface{}
}

func New() Data {
	var newCacheObject = make(map[string]interface{})
	return Data{newCacheObject}
}

func (cache Data) Set(key string, value interface{}) {
	cache.cache[key] = value
}

func (cache Data) Get(key string) interface{} {
	return cache.cache[key]
}

func (cache Data) Delete(key string) {
	delete(cache.cache, key)
}
