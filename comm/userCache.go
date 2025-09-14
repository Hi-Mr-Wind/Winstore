package comm

import (
	"bytes"
	"encoding/gob"
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
	Items *sync.Map
	File  string
	mu    sync.RWMutex
}

// RegisterType 注册类型到gob，避免循环引用
func RegisterType(value interface{}) {
	gob.Register(value)
}

func NewCache() *Cache {
	cache := &Cache{
		File:  filepath.Join(UserDir, "cache.gob"),
		Items: &sync.Map{},
	}
	cache.loadFromDisk()
	go cache.autoSave(3 * time.Minute)
	AppCache = cache
	return cache
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	expireAt := now.Add(ttl)
	if now.Equal(expireAt) {
		expireAt = now.Add(24 * 30 * 12 * 100 * time.Hour)
	}
	c.Items.Store(key, CacheItem{Value: value, ExpireAt: expireAt})

	go func() {
		err := c.SaveToDisk()
		if err != nil {
			fmt.Println("Error saving cache:", err)
		}
	}()
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.Items.Load(key)
	if !found {
		return nil, false
	}
	cacheItem := item.(CacheItem)
	if time.Now().After(cacheItem.ExpireAt) {
		c.Items.Delete(key)
		return nil, false
	}
	return cacheItem.Value, true
}

func (c *Cache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, found := c.Items.Load(key)
	if !found {
		return fmt.Errorf("key not found")
	}
	c.Items.Delete(key)
	err := c.SaveToDisk()
	if err != nil {
		return fmt.Errorf("delete cache error:%v", err)
	}
	return nil
}

func (c *Cache) GetAll() map[string]interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()

	allItems := make(map[string]interface{})
	c.Items.Range(func(key, value interface{}) bool {
		cacheItem := value.(CacheItem)
		if !time.Now().After(cacheItem.ExpireAt) {
			allItems[key.(string)] = cacheItem.Value
		}
		return true
	})
	return allItems
}

func (c *Cache) loadFromDisk() {
	c.mu.Lock()
	defer c.mu.Unlock()

	file, err := readFromFile(c.File)
	if err != nil {
		fmt.Println("Error loading cache:", err)
		return
	}

	// 将加载的数据复制到当前缓存中
	file.Items.Range(func(key, value interface{}) bool {
		c.Items.Store(key, value)
		return true
	})
}

func (c *Cache) SaveToDisk() error {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return writeToTheFile(c)
}

func (c *Cache) autoSave(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Println("—————————————————————————————自动保存缓存———————————————————————————")
		err := c.SaveToDisk()
		if err != nil {
			fmt.Printf("自动保存缓存失败: %v\n", err)
		}
	}
}

func writeToTheFile(cache *Cache) error {
	dir := filepath.Dir(cache.File)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("创建目录失败: %w", err)
	}

	file, err := os.Create(cache.File)
	if err != nil {
		return fmt.Errorf("创建文件失败: %w", err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			fmt.Printf("关闭文件时出错: %v\n", closeErr)
		}
	}()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(cache)
	if err != nil {
		return fmt.Errorf("编码失败: %w", err)
	}
	return nil
}

func readFromFile(filePath string) (*Cache, error) {
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return &Cache{Items: &sync.Map{}}, nil
		}
		return nil, fmt.Errorf("打开文件失败: %w", err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			fmt.Printf("关闭文件时出错: %v\n", closeErr)
		}
	}()

	decoder := gob.NewDecoder(file)
	var cache Cache
	err = decoder.Decode(&cache)
	if err != nil {
		// 处理空文件或损坏文件的情况
		if err.Error() == "unexpected EOF" {
			fmt.Println("缓存文件为空或已损坏，将创建新的缓存")
			return &Cache{Items: &sync.Map{}}, nil
		}
		return nil, fmt.Errorf("解码失败: %w", err)
	}

	if cache.Items == nil {
		cache.Items = &sync.Map{}
	}
	return &cache, nil
}

func (c *Cache) GobEncode() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(c.File); err != nil {
		return nil, err
	}

	items := make(map[interface{}]interface{})
	c.Items.Range(func(key, value interface{}) bool {
		items[key] = value
		return true
	})

	if err := enc.Encode(items); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c *Cache) GobDecode(data []byte) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err := dec.Decode(&c.File); err != nil {
		return err
	}

	var items map[interface{}]interface{}
	if err := dec.Decode(&items); err != nil {
		return err
	}

	c.Items = &sync.Map{}
	for key, value := range items {
		c.Items.Store(key, value)
	}

	return nil
}
