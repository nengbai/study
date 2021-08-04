package config

//Nginx nginx  配置
type Nginx struct {
	Port    int    `yaml:"Port"`
	LogPath string `yaml:"LogPath"`
	Path    string `yaml:"Path"`
}

//Database Config
type Database struct {
	Port     int      `yaml:"Port"`
	Hostname []string `yaml:"Hostname"`
	Username string   `yaml:"Username"`
	Password string   `yaml:"Password"`
}

//Config   系统配置配置
type Config struct {
	Name         string   `yaml:"SiteName"`
	Addr         string   `yaml:"SiteAddr"`
	HTTPS        bool     `yaml:"Https"`
	SiteNginx    Nginx    `yaml:"Nginx"`
	SiteDatabase Database `yaml:"Database"`
}
