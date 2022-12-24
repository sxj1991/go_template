package tool

import (
	"encoding/json"
	"github.com/patrickmn/go-cache"
	"time"
)

// localCache 本地缓存/*
var localCache *cache.Cache

func init() {
	//初始化缓存时间
	localCache = cache.New(5*time.Minute, 10*time.Minute)
}

// SetCacheKey 不限时间 缓存数据/**
func SetCacheKey[T any](key string, t T) {
	localCache.Set(key, t, cache.NoExpiration)
}

func SetCacheKeyTime[T any](key string, t T, d time.Duration) {
	localCache.Set(key, t, d)
}

// GetCacheStruct 获取数据/**
func GetCacheStruct(key string) []byte {
	//返回布尔值和缓存值(接口需断言)
	cache, bool := localCache.Get(key)
	jsonByte, _ := json.Marshal(cache)
	if bool {
		return jsonByte
	} else {
		return nil
	}
}

// RemoveCache 删除数据/**
func RemoveCache(key string) {
	localCache.Delete(key)
}
