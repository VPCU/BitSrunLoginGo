package models

import (
	"github.com/Mmx233/BitSrunLoginGo/pkg/srun"
)

type Guardian struct {
	Enable   bool `json:"enable" yaml:"enable" mapstructure:"enable"`
	Duration uint `json:"duration" yaml:"duration" mapstructure:"duration"`
}

type Basic struct {
	Https          bool   `json:"https" yaml:"https" mapstructure:"https"`
	SkipCertVerify bool   `json:"skip_cert_verify" yaml:"skip_cert_verify" mapstructure:"skip_cert_verify"`
	Timeout        uint   `json:"timeout" yaml:"timeout" mapstructure:"timeout"`
	Interfaces     string `json:"interfaces" yaml:"interfaces" mapstructure:"interfaces"`
	IgnorePublicIP bool   `json:"ignore_public_ip" yaml:"ignore_public_ip" mapstructure:"ignore_public_ip"`
}

type Log struct {
	DebugLevel bool   `json:"debug_level" yaml:"debug_level" mapstructure:"debug_level"`
	WriteFile  bool   `json:"write_file" yaml:"write_file" mapstructure:"write_file"`
	FilePath   string `json:"log_path" yaml:"log_path" mapstructure:"log_path"`
	FileName   string `json:"log_name" yaml:"log_name" mapstructure:"log_name"`
}

type DDNS struct {
	Enable   bool                   `json:"enable" yaml:"enable" mapstructure:"enable"`
	TTL      uint                   `json:"ttl" yaml:"ttl" mapstructure:"ttl"`
	Domain   string                 `json:"domain" yaml:"domain" mapstructure:"domain"`
	Provider string                 `json:"provider" yaml:"provider" mapstructure:"provider"`
	Config   map[string]interface{} `mapstructure:",remain"`
}

type Settings struct {
	Basic    Basic    `json:"basic" yaml:"basic" mapstructure:"basic"`
	Guardian Guardian `json:"guardian" yaml:"guardian" mapstructure:"guardian"`
	Log      Log      `json:"log" yaml:"log" mapstructure:"log"`
	DDNS     DDNS     `json:"ddns" yaml:"ddns" mapstructure:"ddns"`
}

type Config struct {
	Form     srun.LoginForm `json:"form" yaml:"form" mapstructure:"form"`
	Meta     srun.LoginMeta `json:"meta" yaml:"meta" mapstructure:"meta"`
	Settings Settings       `json:"settings" yaml:"settings" mapstructure:"settings"`
}
