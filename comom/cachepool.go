package comom

import (
	"sync"
	"time"
)

var cachePools *cachePool

type cachePool struct {
	m      sync.RWMutex
	cache  map[string]time.Time
	expire time.Duration
}

func NewCachePool(expire time.Duration) *cachePool {
	pool := &cachePool{
		m:      sync.RWMutex{},
		cache:  map[string]time.Time{},
		expire: expire, // 缓存过期时间为 2 秒
	}
	if cachePools != nil {
		pool.cache = cachePools.cache
	}
	cachePools = pool
	return pool
}

func (p *cachePool) Check(key string) bool {
	p.m.Lock()
	defer p.m.Unlock()

	if _, ok := p.cache[key]; !ok {
		// 如果key不存在，表示该数据没有被限流过，直接返回true,通行
		p.cache[key] = time.Now()
		return true
	}
	lastTime := p.cache[key]
	// 如果key存在，判断当前时间与上一次限流时间的差值是否大于等于限制的时间
	if time.Since(lastTime) < p.expire {
		p.cache[key] = time.Now()
		return false
	}
	return true
}
