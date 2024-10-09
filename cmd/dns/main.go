package main

import (
	"github.com/miekg/dns"

	"github.com/jianlu8023/go-example/pkg/logger"
)

var (
	log = logger.GetAppLogger()
)

func main() {
	// 设置要查询的域名和 DNS 服务器地址
	domain := "baidu.com"
	server := "223.5.5.5:53"

	// 创建一个 DNS 消息
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(domain), dns.TypeA)

	// 发送 DNS 查询
	c := new(dns.Client)
	in, _, err := c.Exchange(m, server)
	if err != nil {
		log.Errorf("发送DNS查询失败 %v", err)
		return
	}

	// 处理 DNS 响应
	if len(in.Answer) > 0 {
		for _, ans := range in.Answer {
			if a, ok := ans.(*dns.A); ok {
				log.Infof("a %v", a.A)
			}
		}
	}
}
