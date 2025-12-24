package config

type Nacos struct {
	Host      string `yaml:"Host"`
	Port      int    `yaml:"Port"`
	NameSpace string `yaml:"NameSpace"`
	Group     string `yaml:"Group"`
	DataId    string `yaml:"DataId"`
}

type Mysql struct {
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	Database string `yaml:"Database"`
}
type Redis struct {
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	Password string `yaml:"Password"`
	Database int    `yaml:"Database"`
}
type Config struct {
	Nacos Nacos `yaml:"Nacos"`
	Mysql Mysql `yaml:"Mysql"`
	Redis Redis `yaml:"Redis"`
}
