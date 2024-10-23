package config

type Config struct {
	Application Application `yaml:"application"`
	Server      Server      `yaml:"server"`
	// Redis       Redis       `yaml:"redis"`
	EnvCtrl EnvCtrl   `yaml:"envctrl"`
	Api     Api       `yaml:"api"`
	Service []Service `yaml:"service"`
}

type Service struct {
	Name              string `yaml:"name"`
	Endpoint          string `yaml:"endpoint"`
	System            string `yaml:"system"`
	User              string `yaml:"user"`
	Password          string `yaml:"password"`
	ConnectionTimeout string `yaml:"connectionTimeout"`
	ReadTimeout       string `yaml:"readTimeout"`
	ApiKey            string `yaml:"apiKey"`
	ApiValue          string `yaml:"apiValue"`
}

type Application struct {
	Name    string `yaml:"name"`
	Profile string `yaml:"profile"`
}

type Server struct {
	Port string `yaml:"port"`
}

type Redis struct {
	Address string `yaml:"address"`
}

type EnvCtrl struct {
	MaxThread int `yaml:"maxThread"`
}

type Api struct {
	GetRedisCacheValues   GetRedisCacheValues   `yaml:"getrediscachevalues"`
	SetRedisCacheMultiple SetRedisCacheMultiple `yaml:"setrediscachemultiple"`
	DelRedisCacheMultiple DelRedisCacheMultiple `yaml:"delrediscachemultiple"`
}

type GetRedisCacheValues struct {
	MaxInputCacheKeys int `yaml:"maxinputcachekeys"`
}

type SetRedisCacheMultiple struct {
	MaxInputCacheKeys int `yaml:"maxinputcachekeys"`
}

type DelRedisCacheMultiple struct {
	MaxInputCacheKeys int `yaml:"maxinputcachekeys"`
}
