package dns

import (
	"github.com/Mmx233/BitSrunLoginGo/dns/aliyun"
	"github.com/Mmx233/BitSrunLoginGo/dns/cloudflare"
	log "github.com/sirupsen/logrus"
)

func Run(c *Config) error {
	if c.TTL == 0 {
		c.TTL = 600
	}

	// 配置解析

	var dns Provider
	var e error
	switch c.Provider {
	case "aliyun":
		dns, e = aliyun.New(c.TTL, c.Conf, c.Http)
	case "cloudflare":
		dns, e = cloudflare.New(c.TTL, c.Conf, c.Http)
	default:
		log.Warnf("DDNS 模块 dns 运营商 %s 不支持", c.Provider)
		return nil
	}
	if e != nil {
		log.Warnf("解析 DDNS 配置失败：%v", e)
		return e
	}

	// 修改 dns 记录

	if e = dns.SetDomainRecord(c.Domain, c.IP); e != nil {
		log.Warnf("设置 dns 解析记录失败：%v", e)
		return e
	}

	return nil
}
