package main

type Nginx struct {
    application       *Application		// 组合实现继承，继承自 Application 类
    maxAllowedRequest int
    rateLimiter       map[string]int
}

// 构造器
func newNginxServer() *Nginx {
    return &Nginx{
        application:       &Application{},
        maxAllowedRequest: 2,
        rateLimiter:       make(map[string]int),
    }
}

// 实现 Server 接口 handleRequest 方法
func (n *Nginx) handleRequest(url, method string) (int, string) {
    allowed := n.checkRateLimiting(url)
    if !allowed {
        return 403, "Not Allowed"
    }
    return n.application.handleRequest(url, method)
}

func (n *Nginx) checkRateLimiting(url string) bool {
    if n.rateLimiter[url] == 0 {
        n.rateLimiter[url] = 1
    }
    if n.rateLimiter[url] > n.maxAllowedRequest {
        return false
    }
    n.rateLimiter[url] = n.rateLimiter[url] + 1
    return true
}