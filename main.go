package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type LRUCache struct {
	capacity int
	cache    map[string]string
	order    []string
}

type DeletedItem struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	Timestamp string `json:"timestamp"`
}

var deletedHistory []DeletedItem

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]string),
		order:    []string{},
	}
}

func (l *LRUCache) Get(key string) (string, bool) {
	value, exists := l.cache[key]
	if !exists {
		return "", false
	}
	l.updateOrder(key)
	return value, true
}

func (l *LRUCache) Put(key, value string) string {
	var removedKey string

	if _, exists := l.cache[key]; exists {
		l.cache[key] = value
		l.updateOrder(key)
		return removedKey
	}

	if len(l.cache) >= l.capacity {
		oldest := l.order[0]
		l.order = l.order[1:]
		deletedHistory = append(deletedHistory, DeletedItem{
			Key:       oldest,
			Value:     l.cache[oldest],
			Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		})
		removedKey = oldest
		delete(l.cache, oldest)
	}

	l.cache[key] = value
	l.order = append(l.order, key)
	return removedKey
}

func (l *LRUCache) updateOrder(key string) {
	for i, k := range l.order {
		if k == key {
			l.order = append(l.order[:i], l.order[i+1:]...)
			break
		}
	}
	l.order = append(l.order, key)
}

func (l *LRUCache) View() map[string]string {
	return l.cache
}

func main() {
	r := gin.Default()
	cache := NewLRUCache(5)

	r.Static("/static", "./static")

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	r.GET("/api/lru/view", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"cache": cache.View()})
	})

	r.GET("/api/lru/get/:key", func(c *gin.Context) {
		key := c.Param("key")
		value, exists := cache.Get(key)
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"message": "Cache miss"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Cache hit", "value": value})
	})

	r.POST("/api/lru/put", func(c *gin.Context) {
		var json struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		}
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
			return
		}

		removedKey := cache.Put(json.Key, json.Value)
		c.JSON(http.StatusOK, gin.H{"message": "Item added to cache", "removedKey": removedKey})
	})

	r.GET("/api/lru/history", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"history": deletedHistory})
	})

	r.Run(":8080")
}
