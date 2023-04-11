package comom

import (
	"github.com/patrickmn/go-cache"
	"log"
	"time"
)

var memoryCache *cache.Cache = nil

func init() {
	// 默认30分钟过期，每10min清理一次过期缓存
	memoryCache = cache.New(30*time.Second, 10*time.Minute)
	log.Printf("initial cache success.\n")
}

func GetInstance() *cache.Cache {
	return memoryCache
}

func Set(key string, obj interface{}, expireSeconds time.Duration) {
	memoryCache.Set(key, obj, expireSeconds*time.Second)
}

func Get(key string) interface{} {
	out, found := memoryCache.Get(key)
	if found {
		return out
	}
	return nil
}

func Delete(key string) {
	memoryCache.Delete(key)
}
