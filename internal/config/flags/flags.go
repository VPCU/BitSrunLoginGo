package flags

import (
	"flag"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

var (
	// Path 配置文件路径
	Path string

	Interface string
	Debug     bool
	AutoAcid  bool
	Acid      string
)

func init() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("[init] 获取用户目录失败：", err)
	}
	default_path := filepath.Join(homedir, ".config", "bitsrun.yaml")
	flag.StringVar(&Path, "config", default_path, "config path")

	flag.StringVar(&Interface, "interface", "", "specify the eth name")
	flag.BoolVar(&Debug, "debug", false, "enable debug mode")
	flag.BoolVar(&AutoAcid, "auto-acid", false, "auto detect acid")
	flag.StringVar(&Acid, "acid", "", "specify acid value")

	flag.Parse()
}
