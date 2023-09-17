package controllers

import (
	"net/http"
	"errors"
	
	"github.com/VPCU/BitSrunLoginGo/internal/config"
	"github.com/VPCU/BitSrunLoginGo/internal/config/flags"
	"github.com/VPCU/BitSrunLoginGo/internal/pkg/dns"
	"github.com/VPCU/BitSrunLoginGo/pkg/srun"
	"github.com/VPCU/BitSrunLoginGo/tools"
	log "github.com/sirupsen/logrus"
)

// Login 登录逻辑
func Login(eth *tools.Eth, debugOutput bool) error {
	// 登录配置初始化
	httpClient := tools.HttpPackSelect(eth).Client
	srunClient := srun.New(&srun.Conf{
		Https: config.Settings.Basic.Https,
		LoginInfo: srun.LoginInfo{
			Form: *config.Form,
			Meta: *config.Meta,
		},
		Client: httpClient,
	})

	// 嗅探 acid
	if flags.AutoAcid {
		log.Debugln("开始嗅探 acid")
		acid, err := srunClient.DetectAcid()
		if err != nil {
			if errors.Is(err, srun.ErrAcidCannotFound) {
				log.Errorln("找不到 acid，使用配置 acid")
			} else {
				log.Errorf("嗅探 acid 失败，使用配置 acid: %v", err)
			}
		} else {
			log.Debugf("使用嗅探 acid: %s", acid)
			srunClient.LoginInfo.Meta.Acid = acid
		}
	}

	// 选择输出函数
	var output func(args ...interface{})
	if debugOutput {
		output = log.Debugln
	} else {
		output = log.Infoln
	}

	output("正在获取登录状态")

	online, ip, err := srunClient.LoginStatus()
	if err != nil {
		return err
	}

	log.Debugln("认证客户端 ip: ", ip)

	// 登录执行

	if online {
		output("已登录~")

		if config.Settings.DDNS.Enable && config.Settings.Guardian.Enable && ipLast != ip {
			if ddns(ip, httpClient) == nil {
				ipLast = ip
			}
		}

		return nil
	} else {
		log.Infoln("检测到用户未登录，开始尝试登录...")

		if err = srunClient.DoLogin(ip); err != nil {
			return err
		}

		log.Infoln("登录成功~")

		if config.Settings.DDNS.Enable {
			_ = ddns(ip, httpClient)
		}
	}

	return nil
}

var ipLast string

func ddns(ip string, httpClient *http.Client) error {
	return dns.Run(&dns.Config{
		Provider: config.Settings.DDNS.Provider,
		IP:       ip,
		Domain:   config.Settings.DDNS.Domain,
		TTL:      config.Settings.DDNS.TTL,
		Conf:     config.Settings.DDNS.Config,
		Http:     httpClient,
	})
}
