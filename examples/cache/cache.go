package cache

//START-CACHE OMIT
type Cache interface {
	Get(key string) (string, error)
	Set(key, value string) error
}

//END-CACHE OMIT
