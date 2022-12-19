package cache

import "errors"

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

func (cache Data) Set(key string, value interface{}) error {
	if len(key) == 0 {
		return errors.New("the key length is zero")
	} else if value == nil {
		return errors.New("the value length is zero")
	}

	cache.cache[key] = value
	return nil
}

func (cache Data) Get(key string) (interface{}, error) {
	if len(key) == 0 {
		return nil, errors.New("the key length is zero")
	}

	return cache.cache[key], nil
}

func (cache Data) Delete(key string) error {
	if len(key) == 0 {
		return errors.New("the key length is zero")
	}

	delete(cache.cache, key)
	return nil
}
