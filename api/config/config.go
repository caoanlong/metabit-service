package config

type Conf struct {
	Name  string `json:"name" yaml:"name"`
	Env   string `json:"env" yaml:"env"`
	Port  string `json:"port" yaml:"port"`
	Mysql Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}
