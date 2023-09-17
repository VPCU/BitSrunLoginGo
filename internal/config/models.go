package config

type (
	GuardianConf struct {
		Enable   bool `json:"enable" yaml:"enable"`
		Duration uint `json:"duration" yaml:"duration"`
	}

	BasicConf struct {
		Https          bool   `json:"https" yaml:"https"`
		SkipCertVerify bool   `json:"skip_cert_verify" yaml:"skip_cert_verify"`
		Timeout        uint   `json:"timeout" yaml:"timeout"`
		Interfaces     string `json:"interfaces" yaml:"interfaces"`
	}

	LogConf struct {
		DebugLevel bool   `json:"debug_level" yaml:"debug_level"`
		WriteFile  bool   `json:"write_file" yaml:"write_file"`
		FilePath   string `json:"file_path" yaml:"log_path"`
		FileName   string `json:"file_name" yaml:"log_name"`
	}

	DdnsConf struct {
		Enable   bool                   `json:"enable" yaml:"enable"`
		TTL      uint                   `json:"ttl" yaml:"ttl"`
		Domain   string                 `json:"domain" yaml:"domain"`
		Provider string                 `json:"provider" yaml:"provider"`
		Config   map[string]interface{} `json:"config" yaml:"config"`
	}
)

type SettingsConf struct {
	Basic    BasicConf    `json:"basic" yaml:"basic"`
	Guardian GuardianConf `json:"guardian" yaml:"guardian"`
	Log      LogConf      `json:"log" yaml:"log"`
	DDNS     DdnsConf     `json:"ddns" yaml:"ddns"`
}
