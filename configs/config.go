package configs

var Cfg Config

type Config struct {
	AccessKey    string `yaml:"access_key" mapstructure:"access_key"`
	DefaultReply string `yaml:"default_reply" mapstructure:"default_reply"`
	BaseUrl      string `yaml:"base_url" mapstructure:"base_url"`
	ConfigFile   string `yaml:"config_file" mapstructure:"config_file"`
}
