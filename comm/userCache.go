package comm

import (
	"encoding/json/v2"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type CacheItem struct {
	Value    interface{}
	ExpireAt time.Time
}

type Cache struct {
	items sync.Map
	file  string
}

func NewCache() *Cache {
	cache := &Cache{file: filepath.Join(UserDir, "cache.json")}
	cache.loadFromDisk()
	go cache.autoSave(3 * time.Minute) // 每3分钟自动保存
	return cache
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	expireAt := time.Now().Add(ttl)
	c.items.Store(key, CacheItem{Value: value, ExpireAt: expireAt})
	go func() {
		err := c.SaveToDisk()
		if err != nil {
			fmt.Println("Error saving cache:", err)
		}
	}()
}

func (c *Cache) Get(key string) (interface{}, bool) {
	item, found := c.items.Load(key)
	if !found {
		return nil, false
	}
	cacheItem := item.(CacheItem)
	if time.Now().After(cacheItem.ExpireAt) {
		c.items.Delete(key)
		return nil, false
	}
	return cacheItem.Value, true
}

func (c *Cache) Delete(key string) error {
	value, found := c.Get(key)
	if !found {
		return fmt.Errorf("key not found")
	}
	c.items.Delete(key)
	err := c.SaveToDisk()
	if err != nil {
		c.Set(key, value, 30*time.Minute)
		return fmt.Errorf("delete cache error:%v", err)
	}
	return nil
}

// GetAll 获取所有缓存项
func (c *Cache) GetAll() map[string]interface{} {
	allItems := make(map[string]interface{})
	c.items.Range(func(key, value interface{}) bool {
		cacheItem := value.(CacheItem)
		if !time.Now().After(cacheItem.ExpireAt) {
			allItems[key.(string)] = cacheItem.Value
		}
		return true
	})
	return allItems
}

// 从文件加载数据
func (c *Cache) loadFromDisk() {
	jsonData, err := os.ReadFile(c.file)
	if err != nil {
		fmt.Println("Error loading cache:", err)
		return
	}
	jsonerr := json.Unmarshal(jsonData, &c.items)
	if jsonerr != nil {
		fmt.Println("Error loading cache:", jsonerr)
		return
	}
}

// SaveToDisk 保存数据到文件
func (c *Cache) SaveToDisk() error {
	data := make(map[string]CacheItem)
	c.items.Range(func(key, value interface{}) bool {
		cacheItem := value.(CacheItem)
		if !time.Now().After(cacheItem.ExpireAt) {
			data[key.(string)] = cacheItem
		}
		return true
	})
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	oserr := os.WriteFile(c.file, jsonData, os.ModePerm)
	if oserr != nil {
		return oserr
	}
	return nil
}

// 自动保存数据
func (c *Cache) autoSave(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		err := c.SaveToDisk()
		if err != nil {
			return
		}
	}
}
