package main

import "fmt"

type Lru struct {
}

// 实现 evictionAlgo 接口 evict 方法
func (l *Lru) evict(c *Cache) {
    fmt.Println("Evicting by lru strtegy")
}