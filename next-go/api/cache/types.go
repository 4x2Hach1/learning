package cache

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

type CacheKey interface {
	Key() string
}

type CacheValue interface {
	Value() string
	FromStr(string)
}

// CacheHeavy ///////////////////////

type CacheHeavyKey struct {
	User int
}

func (c *CacheHeavyKey) Key() string {
	key := sha256.Sum256([]byte(fmt.Sprintf("user:%d", c.User)))
	hash := hex.EncodeToString(key[:])
	return hash
}

type CacheHeavyRawKey struct {
	Hash string
}

func (c *CacheHeavyRawKey) Key() string {
	return c.Hash
}

type CacheHeavyValue struct {
	Status int
}

func (c *CacheHeavyValue) Value() string {
	strStatus := strconv.Itoa(c.Status)
	return strStatus
}

func (c *CacheHeavyValue) FromStr(strStatus string) {
	val, _ := strconv.Atoi(strStatus)
	c.Status = val
}
