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
	items *sync.Map
	file  string
}

func NewCache() *Cache {
	cache := &Cache{file: filepath.Join(UserDir, "cache.json"), items: &sync.Map{}}
	cache.loadFromDisk()
	go cache.autoSave(3 * time.Minute) // 每3分钟自动保存
	AppCache = cache
	return cache
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	now := time.Now()
	expireAt := now.Add(ttl)
	if now.Equal(expireAt) {
		expireAt = now.Add(24 * 30 * 12 * 100 * time.Hour)
	}
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
		c.Set(key, value, 24*30*time.Hour)
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
	if c.items == nil {
		c.items = &sync.Map{}
	}
	jsonData, err := os.ReadFile(c.file)
	if err != nil {
		// 如果文件不存在，创建空文件
		if os.IsNotExist(err) {
			e := os.WriteFile(c.file, []byte("{}"), os.ModePerm)
			if e != nil {
				fmt.Println("Error creating cache file:", e)
			}
			return
		}
		fmt.Println("Error reading cache file:", err)
		return
	}
	if jsonData == nil {
		c.items = &sync.Map{}
	}
	var tempMap map[string]CacheItem
	if err := json.Unmarshal(jsonData, &tempMap); err != nil {
		fmt.Println("Error unmarshaling cache data:", err)
		c.items = &sync.Map{}
		return
	}
	// 将普通 map 的内容加载到 sync.Map
	for key, value := range tempMap {
		c.items.Store(key, value)
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
		fmt.Println("—————————————————————————————自动保存缓存正在执行———————————————————————————")
		err := c.SaveToDisk()
		if err != nil {
			return
		}
	}
}
