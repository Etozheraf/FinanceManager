package Config

type DataBaseConfig struct {
	Host         string `yaml:"host"`
	UserName     string `yaml:"username"`
	Password     string `yaml:"password"`
	DatabaseName string `yaml:"dbname"`
	Port         string `yaml:"port"`
}
