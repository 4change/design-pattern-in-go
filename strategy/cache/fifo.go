package main

import "fmt"

type Fifo struct {
}

// 实现 evictionAlgo 接口 evict 方法
func (l *Fifo) evict(c *Cache) {
    fmt.Println("Evicting by fifo strtegy")
}