package mem

import (
	"encoding/json"
	"math/big"

	"github.com/coocood/freecache"
	"github.com/pkg/errors"
	"skatechain.org/lib/crypto/ecdsa"
)

type MemCache struct {
	*freecache.Cache
}

func NewCache(size int) *MemCache {
	return &MemCache{Cache: freecache.NewCache(size)}
}

type Operator struct {
	Address  string
	Strategy string
	Amount   string // base 10 value
}

func GenKey(op Operator) []byte {
	return ecdsa.Keccak256([]byte(op.Address), []byte(op.Strategy))
}

func (cache *MemCache) CacheOperator(key []byte, operator Operator) error {
	cacheKey := ecdsa.Keccak256([]byte("skateavs:operator:"), key)
	data, err := json.Marshal(operator)
	if err != nil {
		return err
	}
	return cache.Set(cacheKey, data, 0)
}

func (cache *MemCache) GetOperator(key []byte) (*Operator, error) {
	cacheKey := ecdsa.Keccak256([]byte("skateavs:operator:"), key)
	data, err := cache.Get(cacheKey)
	if err != nil {
		return nil, errors.Wrap(err, "Cache.GetOperator/Get")
	}
	if len(data) == 0 {
		return nil, nil
	}
	var result Operator
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, errors.Wrap(err, "Cache.GetOperator/UnMarshal")
	}

	return &result, nil
}

func (cache *MemCache) EvictOperator(key []byte) bool {
	cacheKey := ecdsa.Keccak256([]byte("skateavs:operator:"), key)
	return cache.Del(cacheKey)
}

func (cache *MemCache) CacheStake(strategy string, value big.Int) error {
	cacheKey := ecdsa.Keccak256([]byte("skateavs:stake:"), []byte(strategy))
	return cache.Set(cacheKey, value.Bytes(), 0)
}

func (cache *MemCache) GetStake(strategy string) (*big.Int, error) {
	cacheKey := ecdsa.Keccak256([]byte("skateavs:stake:"), []byte(strategy))
	data, err := cache.Get(cacheKey)
	if err != nil {
		return nil, errors.Wrap(err, "Cache.GetStake/Get")
	}
	if len(data) == 0 {
		return nil, nil
	}

  result := new(big.Int)

	return result.SetBytes(data), nil
}
