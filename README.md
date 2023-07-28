  # BitSrunLoginGo

[![Lisense](https://img.shields.io/github/license/Mmx233/BitSrunLoginGo)](https://github.com/Mmx233/BitSrunLoginGo/blob/main/LICENSE)
[![Release](https://img.shields.io/github/v/release/Mmx233/BitSrunLoginGo?color=blueviolet&include_prereleases)](https://github.com/Mmx233/BitSrunLoginGo/releases)
[![GoReport](https://goreportcard.com/badge/github.com/Mmx233/BitSrunLoginGo)](https://goreportcard.com/report/github.com/Mmx233/BitSrunLoginGo)

深澜校园网登录脚本 Go 语言版。GO 语言可以直接交叉编译出 mips 架构可执行程序（路由器）（主流平台更不用说了），从而免除安装环境。

> 主要登录逻辑来自： https://github.com/coffeehat/BIT-srun-login-script

Openwrt 可以参考 [immortalwrt/packages/net/bitsrunlogin-go](https://github.com/immortalwrt/packages/tree/master/net/bitsrunlogin-go) 添加构建配置文件。单独复制文件夹、添加源或直接使用 [immortalwrt](https://github.com/immortalwrt/immortalwrt) 编译固件，选择 application / [luci-app-bitsrunlogin-go](https://github.com/immortalwrt/luci/tree/master/applications/luci-app-bitsrunlogin-go)

## :gear:运行

编译结果为可执行文件，下载 release 或编译后直接运行即可

首次运行将自动生成配置文件，首次使用建议开启调试日志（`settings.log.debug_level`），可以通过添加启动参数 `--config` 指定配置文件路径，默认为当前目录的 `Config.yaml`

支持 `json`、`yaml`、`yml`、`toml`、`hcl`、`tfvars` 等，仅对 `json`和`yaml` 进行了优化与测试

```shell
./bitsrun --config=./another-config.yaml
./bitsrun --config=/demo/i.json
```

其他 Flags:

```text
--interface eth0.1  #指定使用 eth0.1 登录，多网卡模式无效
--debug             #临时开启 debug 模式
--auto-acid         #自动嗅探 acid
--acid 7            #覆写配置文件 acid 值
```

Config.yaml 说明：

```yaml
form:
  domain: www.msftconnecttest.com #登录地址 ip 或域名
  username: "" #账号
  usertype: cmcc #运营商类型，详情看下方
  password: "" #密码
meta: #登录参数
  "n": "200"
  type: "1"
  acid: "5"
  enc: srun_bx1
settings:
  basic: #基础设置
    https: false #访问校园网 API 时使用 https 协议
    skip_cert_verify: false #跳过证书有效校验
    timeout: 5 #网络请求超时时间（秒，正整数）
    interfaces: "" #网卡名称正则（注意转义），如：eth0\.[2-3]，不为空时为多网卡模式
    ignore_public_ip: true # 设定为true时，分配到公网ip将认为已经登录
  guardian: #守护模式（后台常驻）
    enable: false 
    duration: 300 #网络检查周期（秒，正整数）
  log:
    debug_level: false #打印调试日志
    write_file: false #写日志文件
    log_path: ./ #日志文件存放目录路径
    log_name: "" #指定日志文件名
  ddns: # 校园网内网 ip ddns
    enable: false
    domain: www.example.com
    ttl: 600
    provider: "cloudflare" # dns provider 配置见 DDNS 说明
    #zone: "xxxx"
    #token: "xxxx"
```

登录参数从原网页登陆时对 `/srun_portal` 的请求抓取，抓取时请把浏览器控制台的 `preserve log`（保留日志）启用。

运营商类型在原网页会被自动附加在账号后，请把 `@` 后面的部分填入 `user_type`，没有则留空（删掉默认的）

## :bow_and_arrow: DDNS

将 `ddns.enable` 设为 `true` 后，将在登录成功时以及守护模式下首次启动已在线时设置指定域名的解析地址（ipv4，A 记录），用于校园网内网通信

支持的 Provider 及其设置（将额外配置添加到配置文件 DDNS 配置内）：

|  Provider  | 额外配置项                                   |
|:----------:|-----------------------------------------|
|   aliyun   | `access_key_id`<br/>`access_key_secret` |
| cloudflare | `zone` 区域 ID<br/>`token` API 令牌         |
|   dnspod   | `secret_id`<br/>`secret_key`            |

需要注意的是，应该避免在多网卡模式下使用 ddns，如有需要，可以每个网卡创建一个配置文件单独进行登录。此外，校园网内网通信并不是安全的，校方往往会对其施加监管

## :shower: 多拨

登录请求的网卡绑定在 Linux 下稳定生效，但在其它系统中可能无法成功绑定。如果要在 windows 等系统寻求稳定的多拨效果，可以考虑使用 docker 的网桥或系统层面的绑定。

请注意应该避免在多拨中使用同一个账号。因为在大部分学校中即使多拨成功，同一账号在同一个网关下的多个设备也是共享带宽限制的，这意味着多拨没有效果。

你可以通过配置文件中的 `settings.basic.interfaces` 指定网卡，也可以在将该配置留空的情况下使用 `--interface` 指定网卡。

## :anchor: Docker / Kubernetes

镜像：`mmx233/bitsrunlogin-go:latest`

支持 linux/amd64、linux/386、linux/arm64、linux/arm/v7 架构

直接使用：

配置文件挂载至 `/data/Config.yaml`，若需更改配置文件类型，可以使用 `--entrypoint` 覆写启动参数

```shell
docker run -v path_to_config:/data/Config.yaml mmx233/bitsrunlogin-go:latest
```

自行构建：

如果需要在其他系统或架构使用，可能需要更改构建层与底层镜像，目前使用的 alpine 并不支持 linux 之外的系统

```shell
git clone https://github.com/Mmx233/BitSrunLoginGo.git
cd BitSrunLoginGo
docker build . --file Dockerfile --tag mmx233/bitsrunlogin-go:latest
```

## :hammer_and_wrench:构建

请安装最新版 golang

```shell
#直接编译本系统可执行程序：
git clone https://github.com/Mmx233/BitSrunLoginGo.git
cd BitSrunLoginGo
go build ./cmd/bitsrun

#或者使用经过优化的构建命令：
go build -gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH -ldflags "-s -w -extldflags '-static'" ./cmd/bitsrun

```

交叉编译（Linux）：

```shell
export GOGGC=0
export GOOS=windows #系统
export GOARCH=amd64 #架构
go build -gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH -ldflags "-s -w -extldflags '-static'" ./cmd/bitsrun
```

交叉编译（Powershell）：

```shell
$env:GOGGC=0
$env:GOOS='linux' #系统
$env:GOARCH='amd64' #架构
go build -gcflags=-trimpath=$env:GOPATH -asmflags=-trimpath=$env:GOPATH -ldflags "-s -w -extldflags '-static'" ./cmd/bitsrun

#消除命令行窗口
go build -gcflags=-trimpath=$env:GOPATH -asmflags=-trimpath=$env:GOPATH -ldflags "-s -w -extldflags '-static' -H windowsgui" ./cmd/bitsrun
```

golang 支持的系统与架构请自行查询

## :jigsaw: 作为 module 使用

**\*本项目使用了 AGPL V3 许可证，请酌情引用**

示例：

```go
package main

import (
	"github.com/Mmx233/BitSrunLoginGo/pkg/srun"
)

func main() {
    //具体用法请查看 struct 注释
    client:=srun.New(&srun.Conf{
        Https:  false,
        Client: nil,
        LoginInfo: srun.LoginInfo{
            Form: &srun.LoginForm{
                Domain:   "",
                UserName: "",
                UserType: "",
                PassWord: "",
            },
            Meta: &srun.LoginMeta{
                N:    "",
                Type: "",
                Acid: "",
                Enc:  "",
            },
        },
    })

    online, ip, e := client.LoginStatus()
    if e!=nil {
        panic(e)
    }
	
    if !online {
        e=client.DoLogin(ip)
        if e!=nil {
            panic(e)
        }	
    }
}
```
