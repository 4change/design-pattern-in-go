package main

import "fmt"

type Lfu struct {
}

// 实现 evictionAlgo 接口 evict 方法
func (l *Lfu) evict(c *Cache) {
    fmt.Println("Evicting by lfu strtegy")
}