package global

import (
	"flag"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

var Flags struct {
	// 配置文件路径
	Path string

	// settings overwrite
	Interface string
	Debug     bool
	AutoAcid  bool
	Acid      string
}

func initFlags() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("[init] 获取用户目录失败：", err)
	}
	default_path := filepath.Join(homedir, ".config", "bitsrun.yaml")
	flag.StringVar(&Flags.Path, "config", default_path, "config path")

	flag.StringVar(&Flags.Interface, "interface", "", "specify the eth name")
	flag.BoolVar(&Flags.Debug, "debug", false, "enable debug mode")
	flag.BoolVar(&Flags.AutoAcid, "auto-acid", false, "auto detect acid")
	flag.StringVar(&Flags.Acid, "acid", "", "specify acid value")

	flag.Parse()
}
