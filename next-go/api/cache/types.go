package cache

type CacheKey interface {
	Key() string
}

type CacheValue interface {
	Value() string
	FromStr(string)
}
